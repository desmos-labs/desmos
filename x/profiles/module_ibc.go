package profiles

import (
	"fmt"
	"math"

	oracletypes "github.com/bandprotocol/chain/x/oracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	porttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/05-port/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"

	"github.com/desmos-labs/desmos/x/profiles/keeper"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

// OnChanOpenInit implements the IBCModule interface
func (am AppModule) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) error {
	if err := ValidateProfilesChannelParams(ctx, am.keeper, order, portID, channelID, version); err != nil {
		return err
	}

	// Claim channel capability passed back by IBC module
	if err := am.keeper.ClaimCapability(ctx, chanCap, host.ChannelCapabilityPath(portID, channelID)); err != nil {
		return err
	}

	return nil
}

// OnChanOpenTry implements the IBCModule interface
func (am AppModule) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version,
	counterpartyVersion string,
) error {

	if err := ValidateProfilesChannelParams(ctx, am.keeper, order, portID, channelID, version); err != nil {
		return err
	}

	if counterpartyVersion != types.IBCVersion {
		return sdkerrors.Wrapf(types.ErrInvalidVersion, "invalid counterparty version: got: %s, expected %s", counterpartyVersion, types.IBCVersion)
	}

	// Module may have already claimed capability in OnChanOpenInit in the case of crossing hellos
	// (ie chainA and chainB both call ChanOpenInit before one of them calls ChanOpenTry)
	// If module can already authenticate the capability then module already owns it so we don't need to claim
	// Otherwise, module does not have channel capability and we must claim it from IBC
	if !am.keeper.AuthenticateCapability(ctx, chanCap, host.ChannelCapabilityPath(portID, channelID)) {
		// Only claim channel capability passed back by IBC module if we do not already own it
		if err := am.keeper.ClaimCapability(ctx, chanCap, host.ChannelCapabilityPath(portID, channelID)); err != nil {
			return err
		}
	}

	return nil
}

// OnChanOpenAck implements the IBCModule interface
func (am AppModule) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	counterpartyVersion string,
) error {
	if counterpartyVersion != types.IBCVersion {
		return sdkerrors.Wrapf(types.ErrInvalidVersion, "invalid counterparty version: %s, expected %s", counterpartyVersion, types.IBCVersion)
	}
	return nil
}

// OnChanOpenConfirm implements the IBCModule interface
func (am AppModule) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

// OnChanCloseInit implements the IBCModule interface
func (am AppModule) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	// Disallow user-initiated channel closing for channels
	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "user cannot close channel")
}

// OnChanCloseConfirm implements the IBCModule interface
func (am AppModule) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

// OnRecvPacket implements the IBCModule interface
func (am AppModule) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
) (*sdk.Result, []byte, error) {
	var ack channeltypes.Acknowledgement
	var err error

	// Try handling the chain link packet data
	ack, err = am.handleLinkChainAccountPacketData(ctx, packet)
	if types.ErrInvalidPacketData.Is(err) {
		// Try handling the oracle request packet data
		ack, err = am.handleOracleRequestPacketData(ctx, packet)
	}

	// If packet data is still invalid, return an error
	if types.ErrInvalidPacketData.Is(err) {
		return nil, nil, err
	}

	// Encode acknowledgement
	ackBytes, err := sdk.SortJSON(types.ProtoCdc.MustMarshalJSON(&ack))
	if err != nil {
		return nil, []byte{}, sdkerrors.Wrap(sdkerrors.ErrInvalidType, err.Error())
	}

	// NOTE: acknowledgement will be written synchronously during IBC handler execution.
	return &sdk.Result{
		Events: ctx.EventManager().Events().ToABCIEvents(),
	}, ackBytes, nil
}

// handleLinkChainAccountPacketData tries handling a LinkChainAccountPacketData packet.
// Returns ErrInvalidPacketData if the given packet data is not of such type.
func (am AppModule) handleLinkChainAccountPacketData(
	ctx sdk.Context, packet channeltypes.Packet,
) (channeltypes.Acknowledgement, error) {
	var packetData types.LinkChainAccountPacketData
	if err := types.ProtoCdc.UnmarshalJSON(packet.GetData(), &packetData); err != nil {
		return channeltypes.Acknowledgement{}, sdkerrors.Wrapf(types.ErrInvalidPacketData, "%T", packet)
	}

	var acknowledgement channeltypes.Acknowledgement

	packetAck, err := am.keeper.OnRecvLinkChainAccountPacket(ctx, packetData)
	if err != nil {
		acknowledgement = channeltypes.NewErrorAcknowledgement(err.Error())
	} else {
		// Encode packet acknowledgment
		packetAckBytes, err := packetAck.Marshal()
		if err != nil {
			return channeltypes.Acknowledgement{}, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
		acknowledgement = channeltypes.NewResultAcknowledgement(packetAckBytes)
	}

	address, err := types.UnpackAddressData(am.cdc, packetData.SourceAddress)
	if err != nil {
		return channeltypes.Acknowledgement{}, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeLinkChainAccountPacket,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeChainLinkSourceAddress, address.GetAddress()),
			sdk.NewAttribute(types.AttributeChainLinkSourceChainName, packetData.SourceChainConfig.Name),
			sdk.NewAttribute(types.AttributeChainLinkDestinationAddress, packetData.DestinationAddress),
			sdk.NewAttribute(types.AttributeKeyAckSuccess, fmt.Sprintf("%t", true)),
		),
	)

	return acknowledgement, nil
}

// handleOracleRequestPacketData tries handling a OracleResponsePacketData packet.
// Returns ErrInvalidPacketData if the given packet data is not of such type.
func (am AppModule) handleOracleRequestPacketData(
	ctx sdk.Context, packet channeltypes.Packet,
) (channeltypes.Acknowledgement, error) {
	var data oracletypes.OracleResponsePacketData
	if err := oracletypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return channeltypes.Acknowledgement{}, sdkerrors.Wrapf(types.ErrInvalidPacketData, "%T", packet)
	}

	acknowledgement := channeltypes.NewResultAcknowledgement([]byte{byte(1)})

	err := am.keeper.OnRecvApplicationLinkPacketData(ctx, data)
	if err != nil {
		acknowledgement = channeltypes.NewErrorAcknowledgement(err.Error())
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypePacket,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeKeyClientID, data.ClientID),
			sdk.NewAttribute(types.AttributeKeyRequestID, fmt.Sprintf("%d", data.RequestID)),
			sdk.NewAttribute(types.AttributeKeyResolveStatus, data.ResolveStatus.String()),
			sdk.NewAttribute(types.AttributeKeyAckSuccess, fmt.Sprintf("%t", true)),
		),
	)

	// NOTE: acknowledgement will be written synchronously during IBC handler execution.
	return acknowledgement, nil
}

// OnAcknowledgementPacket implements the IBCModule interface
func (am AppModule) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
) (*sdk.Result, error) {
	var ack channeltypes.Acknowledgement
	if err := types.AminoCdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
			"cannot unmarshal oracle packet acknowledgement: %v", err)
	}

	var data oracletypes.OracleRequestPacketData
	if err := oracletypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
			"cannot unmarshal oracle request packet data: %s", err.Error())
	}

	if err := am.keeper.OnOracleRequestAcknowledgementPacket(ctx, data, ack); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypePacket,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeKeyClientID, data.ClientID),
			sdk.NewAttribute(types.AttributeKeyAck, fmt.Sprintf("%v", ack)),
		),
	)

	switch resp := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Result:
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypePacket,
				sdk.NewAttribute(types.AttributeKeyAckSuccess, string(resp.Result)),
			),
		)
	case *channeltypes.Acknowledgement_Error:
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypePacket,
				sdk.NewAttribute(types.AttributeKeyAckError, resp.Error),
			),
		)
	}

	return &sdk.Result{
		Events: ctx.EventManager().Events().ToABCIEvents(),
	}, nil
}

// OnTimeoutPacket implements the IBCModule interface
func (am AppModule) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
) (*sdk.Result, error) {
	var data oracletypes.OracleRequestPacketData
	if err := types.AminoCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
			"cannot unmarshal oracle request packet data: %s", err.Error())
	}
	// refund tokens
	if err := am.keeper.OnOracleRequestTimeoutPacket(ctx, data); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeTimeout,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeKeyOracleID, fmt.Sprintf("%d", data.OracleScriptID)),
			sdk.NewAttribute(types.AttributeKeyClientID, data.ClientID),
			sdk.NewAttribute(types.AttributeKeyRequestKey, data.RequestKey),
		),
	)

	return &sdk.Result{
		Events: ctx.EventManager().Events().ToABCIEvents(),
	}, nil
}

// ValidateProfilesChannelParams does validation of a newly created profiles channel. A profiles
// channel must be UNORDERED, use the correct port (by default 'profiles'), and use the current
// supported version. Only 2^32 channels are allowed to be created.
func ValidateProfilesChannelParams(
	ctx sdk.Context,
	keeper keeper.Keeper,
	order channeltypes.Order,
	portID string,
	channelID string,
	version string,
) error {
	// NOTE: for escrow address security only 2^32 channels are allowed to be created
	// Issue: https://github.com/cosmos/cosmos-sdk/issues/7737
	channelSequence, err := channeltypes.ParseChannelSequence(channelID)
	if err != nil {
		return err
	}
	if channelSequence > uint64(math.MaxUint32) {
		return sdkerrors.Wrapf(types.ErrMaxProfilesChannels, "channel sequence %d is greater than max allowed profiles channels %d", channelSequence, uint64(math.MaxUint32))
	}
	if order != channeltypes.UNORDERED {
		return sdkerrors.Wrapf(channeltypes.ErrInvalidChannelOrdering, "expected %s channel, got %s ", channeltypes.UNORDERED, order)
	}

	// Require portID is the portID profiles module is bound to
	boundPort := keeper.GetPort(ctx)
	if boundPort != portID {
		return sdkerrors.Wrapf(porttypes.ErrInvalidPort, "invalid port: %s, expected %s", portID, boundPort)
	}

	if version != types.IBCVersion {
		return sdkerrors.Wrapf(types.ErrInvalidVersion, "got %s, expected %s", version, types.IBCVersion)
	}
	return nil
}
