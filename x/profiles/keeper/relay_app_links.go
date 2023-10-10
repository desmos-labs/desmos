package keeper

import (
	"encoding/hex"
	"fmt"
	"strings"

	"cosmossdk.io/errors"
	"github.com/armon/go-metrics"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"

	"github.com/desmos-labs/desmos/v6/pkg/obi"

	"github.com/desmos-labs/desmos/v6/x/profiles/types"

	oracletypes "github.com/desmos-labs/desmos/v6/x/oracle/types"
)

// OracleScriptCallData represents the data that should be OBI-encoded and sent to perform an oracle request
type OracleScriptCallData struct {
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
//
//nolint:interfacer
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
	sourceChannelEnd, found := k.ChannelKeeper.GetChannel(ctx, sourcePort, sourceChannel)
	if !found {
		return errors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", sourcePort, sourceChannel)
	}

	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

	// Begin createOutgoingPacket logic
	channelCap, ok := k.ScopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return errors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	// Create the call data to be used
	data := OracleScriptCallData{
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

	// Send the IBC packet
	_, err = k.ChannelKeeper.SendPacket(ctx, channelCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, packetData.GetBytes())
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
		ctx.BlockTime().Add(k.GetParams(ctx).AppLinks.ValidityDuration),
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
