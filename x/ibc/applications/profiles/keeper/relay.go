package keeper

import (
	"github.com/armon/go-metrics"
	obi "github.com/bandprotocol/chain/pkg/obi"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"

	"github.com/desmos-labs/desmos/x/ibc/applications/profiles/types"

	oracletypes "github.com/bandprotocol/chain/x/oracle/types"
)

// TODO: Make the following parameter
const (
	// OracleScriptID represents the oracle script to be called on Band Protocol
	OracleScriptID = oracletypes.OracleScriptID(32)

	OracleAskCount   = 10
	OracleMinCount   = 6
	OraclePrepareGas = 200_000
	OracleExecuteGas = 600_000
)

var (
	FeeCoins = sdk.NewCoins(sdk.NewCoin("band", sdk.NewInt(0)))
)

// StartProfileConnection creates and sends an IBC packet containing the proper data allowing to call
// the Band Protocol oracle script so that the sender account can be verified using the given verification data.
// nolint:interfacer
func (k Keeper) StartProfileConnection(
	ctx sdk.Context,
	sourcePort,
	sourceChannel string,
	verificationData types.VerificationData,
	sender sdk.AccAddress,
	feePayer string,
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
	// See spec for this logic: https://github.com/cosmos/ics/tree/master/spec/ics-020-fungible-token-transfer#packet-relay
	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	// Get the oracle call data by OBI-encoding the verification data
	callData, err := obi.Encode(verificationData)
	if err != nil {
		return err
	}

	// Create the Oracle request packet data
	packetData := oracletypes.NewOracleRequestPacketData(
		sender.String()+"-"+verificationData.Method,
		OracleScriptID,
		callData,
		OracleAskCount,
		OracleMinCount,
		FeeCoins,
		feePayer,
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
