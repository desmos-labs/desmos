package keeper

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/armon/go-metrics"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"

	"github.com/desmos-labs/desmos/pkg/obi"

	"github.com/desmos-labs/desmos/x/profiles/types"

	oracletypes "github.com/desmos-labs/desmos/x/oracle/types"
)

// TODO: Make the following parameter
const (
	// OracleScriptID represents the oracle script to be called on Band Protocol
	OracleScriptID = 32

	OracleAskCount   = 10
	OracleMinCount   = 6
	OraclePrepareGas = 50_000
	OracleExecuteGas = 200_000

	FeePayer = "desmos-ibc-profiles"
)

var (
	FeeCoins = sdk.NewCoins(sdk.NewCoin("band", sdk.NewInt(0)))
)

// oracleScriptCallData represents the data that should be OBI-encoded and sent to perform an oracle request
type oracleScriptCallData struct {
	Application string `obi:"application"`
	CallData    string `obi:"call_data"`
}

// resultData represents the data that is returned by the oracle script
type resultData struct {
	Signature string `obi:"signature"`
	Value     string `obi:"value"`
}

// StartProfileConnection creates and sends an IBC packet containing the proper data allowing to call
// the Band Protocol oracle script so that the sender account can be verified using the given verification data.
// nolint:interfacer
func (k Keeper) StartProfileConnection(
	ctx sdk.Context,
	applicationData types.Data,
	oracleRequestCallData types.OracleRequest_CallData,
	sender sdk.AccAddress,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) error {
	sourceChannelEnd, found := k.channelKeeper.GetChannel(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", sourcePort, sourceChannel)
	}

	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

	// Get the next sequence
	sequence, found := k.channelKeeper.GetNextSequenceSend(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", sourcePort, sourceChannel,
		)
	}

	// Begin createOutgoingPacket logic
	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	// Create the call data to be used
	data := oracleScriptCallData{
		Application: oracleRequestCallData.Application,
		CallData:    oracleRequestCallData.CallData,
	}

	// Serialize the call data using the OBI encoding
	callDataBz, err := obi.Encode(data)
	if err != nil {
		return err
	}

	// Create the Oracle request packet data
	clientID := sender.String() + "-" + applicationData.Application + "-" + applicationData.Username
	packetData := oracletypes.NewOracleRequestPacketData(
		clientID,
		OracleScriptID,
		callDataBz,
		OracleAskCount,
		OracleMinCount,
		FeeCoins,
		FeePayer,
		OraclePrepareGas,
		OracleExecuteGas,
	)

	// Create the IBC packet
	packet := channeltypes.NewPacket(
		packetData.GetBytes(),
		sequence,
		sourcePort,
		sourceChannel,
		destinationPort,
		destinationChannel,
		timeoutHeight,
		timeoutTimestamp,
	)

	// Send the IBC packet
	err = k.channelKeeper.SendPacket(ctx, channelCap, packet)
	if err != nil {
		return err
	}

	// Store the connection
	err = k.SaveApplicationLink(ctx, sender.String(), types.NewApplicationLink(
		applicationData,
		types.ApplicationLinkStateInitialized,
		types.NewOracleRequest(-1, int64(OracleScriptID), oracleRequestCallData, clientID),
		nil,
		ctx.BlockTime(),
	))
	if err != nil {
		return err
	}

	labels := []metrics.Label{
		telemetry.NewLabel("destination-port", destinationPort),
		telemetry.NewLabel("destination-channel", destinationChannel),
	}

	defer func() {
		telemetry.IncrCounterWithLabels(
			[]string{"ibc", types.ModuleName, "connect-profile"},
			1,
			labels,
		)
	}()

	return nil
}

func (k Keeper) OnRecvApplicationLinkPacketData(
	ctx sdk.Context,
	data oracletypes.OracleResponsePacketData,
) error {
	// Get the request by the client ID
	user, link, err := k.GetApplicationLinkByClientID(ctx, data.ClientID)
	if err != nil {
		return err
	}

	switch data.ResolveStatus {
	case oracletypes.RESOLVE_STATUS_EXPIRED:
		link.State = types.AppLinkStateVerificationError
		link.Result = types.NewErrorResult(types.ErrRequestExpired)

	case oracletypes.RESOLVE_STATUS_FAILURE:
		link.State = types.AppLinkStateVerificationError
		link.Result = types.NewErrorResult(types.ErrRequestFailed)

	case oracletypes.RESOLVE_STATUS_SUCCESS:
		var result resultData
		err = obi.Decode(data.Result, &result)
		if err != nil {
			return fmt.Errorf("error while decoding request result: %s", err)
		}

		// Verify the application username to make sure it's the same that is returned (avoid replay attacks)
		if !strings.EqualFold(result.Value, link.Data.Username) {
			link.State = types.AppLinkStateVerificationError
			link.Result = types.NewErrorResult(types.ErrInvalidAppUsername)
			return k.SaveApplicationLink(ctx, user, link)
		}

		// Verify the signature to make sure it's from the same user (avoid identity theft)
		addr, err := sdk.AccAddressFromBech32(user)
		if err != nil {
			return err
		}
		acc := k.ak.GetAccount(ctx, addr)

		sigBz, err := hex.DecodeString(result.Signature)
		if err != nil {
			return err
		}

		if !acc.GetPubKey().VerifySignature([]byte(result.Value), sigBz) {
			link.State = types.AppLinkStateVerificationError
			link.Result = types.NewErrorResult(types.ErrInvalidSignature)
			return k.SaveApplicationLink(ctx, user, link)
		}

		link.State = types.AppLinkStateVerificationSuccess
		link.Result = types.NewSuccessResult(result.Value, result.Signature)
	}

	return k.SaveApplicationLink(ctx, user, link)
}

func (k Keeper) OnOracleRequestAcknowledgementPacket(
	ctx sdk.Context,
	data oracletypes.OracleRequestPacketData,
	ack channeltypes.Acknowledgement,
) error {
	// Get the request by the client ID
	user, link, err := k.GetApplicationLinkByClientID(ctx, data.ClientID)
	if err != nil {
		return err
	}

	switch res := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:
		// The acknowledgment failed on the receiving chain.
		// Update the state to ERROR and the result to an error one
		link.State = types.AppLinkStateVerificationError
		link.Result = types.NewErrorResult(res.Error)

	case *channeltypes.Acknowledgement_Result:
		// The acknowledgement succeeded on the receiving chain
		// Set the state to STARTED
		link.State = types.AppLinkStateVerificationStarted

		var packetAck oracletypes.OracleRequestPacketAcknowledgement
		err = types.ModuleCdc.UnmarshalJSON(res.Result, &packetAck)
		if err != nil {
			return fmt.Errorf("cannot unmarshal oracle request packet acknowledgment: %s", err)
		}

		// Set the oracle request ID returned from BAND
		link.OracleRequest.ID = int64(packetAck.RequestID)

	}

	return k.SaveApplicationLink(ctx, user, link)
}

func (k Keeper) OnOracleRequestTimeoutPacket(
	ctx sdk.Context,
	data oracletypes.OracleRequestPacketData,
) error {
	// Get the request by the client ID
	user, connection, err := k.GetApplicationLinkByClientID(ctx, data.ClientID)
	if err != nil {
		return err
	}

	connection.State = types.AppLinkStateVerificationTimedOut

	return k.SaveApplicationLink(ctx, user, connection)
}
