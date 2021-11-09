package keeper

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/armon/go-metrics"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/modules/core/24-host"

	"github.com/desmos-labs/desmos/v2/pkg/obi"

	"github.com/desmos-labs/desmos/v2/x/profiles/types"

	oracletypes "github.com/desmos-labs/desmos/v2/x/oracle/types"
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
	Username  string `obi:"username"`
}

// StartProfileConnection creates and sends an IBC packet containing the proper data allowing to call
// the Band Protocol oracle script so that the sender account can be verified using the given verification data.
// nolint:interfacer
func (k Keeper) StartProfileConnection(
	ctx sdk.Context,
	applicationData types.Data,
	dataSourceCallData string,
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
		Application: strings.ToLower(applicationData.Application),
		CallData:    dataSourceCallData,
	}

	params := k.GetParams(ctx)
	oraclePrams := params.Oracle

	// Serialize the call data using the OBI encoding
	callDataBz, err := obi.Encode(data)
	if err != nil {
		return err
	}

	// Create the Oracle request packet data
	clientID := sender.String() + "-" + applicationData.Application + "-" + applicationData.Username
	packetData := oracletypes.NewOracleRequestPacketData(
		clientID,
		oracletypes.OracleScriptID(oraclePrams.ScriptID),
		callDataBz,
		oraclePrams.AskCount,
		oraclePrams.MinCount,
		oraclePrams.FeeAmount,
		oraclePrams.PrepareGas,
		oraclePrams.ExecuteGas,
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
	err = k.SaveApplicationLink(ctx, types.NewApplicationLink(
		sender.String(),
		applicationData,
		types.ApplicationLinkStateInitialized,
		types.NewOracleRequest(
			0,
			oraclePrams.ScriptID,
			types.NewOracleRequestCallData(applicationData.Application, dataSourceCallData),
			clientID,
		),
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
	link, found, err := k.GetApplicationLinkByClientID(ctx, data.ClientID)
	if err != nil {
		return err
	}

	// If the link is not found, do nothing (it might have been deleted by the user in the meanwhile)
	if !found {
		return nil
	}

	// If the link has already been verified, do nothing
	if link.IsVerificationCompleted() {
		return nil
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
		if !strings.EqualFold(result.Username, link.Data.Username) {
			link.State = types.AppLinkStateVerificationError
			link.Result = types.NewErrorResult(types.ErrInvalidAppUsername)
			return k.SaveApplicationLink(ctx, link)
		}

		// Verify the signature to make sure it's from the same user (avoid identity theft)
		addr, err := sdk.AccAddressFromBech32(link.User)
		if err != nil {
			return err
		}
		acc := k.ak.GetAccount(ctx, addr)

		valueBz, err := hex.DecodeString(result.Value)
		if err != nil {
			return err
		}

		sigBz, err := hex.DecodeString(result.Signature)
		if err != nil {
			return err
		}

		if !acc.GetPubKey().VerifySignature(valueBz, sigBz) {
			link.State = types.AppLinkStateVerificationError
			link.Result = types.NewErrorResult(types.ErrInvalidSignature)
			return k.SaveApplicationLink(ctx, link)
		}

		link.State = types.AppLinkStateVerificationSuccess
		link.Result = types.NewSuccessResult(result.Value, result.Signature)
	}

	return k.SaveApplicationLink(ctx, link)
}

func (k Keeper) OnOracleRequestAcknowledgementPacket(
	ctx sdk.Context,
	data oracletypes.OracleRequestPacketData,
	ack channeltypes.Acknowledgement,
) error {
	// Get the request by the client ID
	link, found, err := k.GetApplicationLinkByClientID(ctx, data.ClientID)
	if err != nil {
		return err
	}

	// If the link is not found, do nothing (it might have been deleted by the user in the meanwhile)
	if !found {
		return nil
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
		link.OracleRequest.ID = uint64(packetAck.RequestID)

	}

	return k.SaveApplicationLink(ctx, link)
}

// OnOracleRequestTimeoutPacket handles the OracleRequestPacketData instance that is sent when a request times out
func (k Keeper) OnOracleRequestTimeoutPacket(
	ctx sdk.Context,
	data oracletypes.OracleRequestPacketData,
) error {
	// Get the request by the client ID
	link, found, err := k.GetApplicationLinkByClientID(ctx, data.ClientID)
	if err != nil {
		return err
	}

	// If the link is not found, do nothing (it might have been deleted by the user in the meanwhile)
	if !found {
		return nil
	}

	link.State = types.AppLinkStateVerificationTimedOut

	return k.SaveApplicationLink(ctx, link)
}
