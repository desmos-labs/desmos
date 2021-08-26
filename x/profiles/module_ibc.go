package profiles

import (
	"fmt"
	"math"

	oracletypes "github.com/desmos-labs/desmos/x/oracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/modules/core/05-port/types"
	host "github.com/cosmos/ibc-go/modules/core/24-host"

	ibcexported "github.com/cosmos/ibc-go/modules/core/exported"

	"github.com/desmos-labs/desmos/x/profiles/keeper"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

var (
	_ porttypes.IBCModule = AppModule{}
)

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

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

// OnChanOpenInit implements the IBCModule interface
func (am AppModule) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	channelCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) error {
	if err := ValidateProfilesChannelParams(ctx, am.keeper, order, portID, channelID, version); err != nil {
		return err
	}

	// Claim channel capability passed back by IBC module
	if err := am.keeper.ClaimCapability(ctx, channelCap, host.ChannelCapabilityPath(portID, channelID)); err != nil {
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
	channelCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version,
	counterpartyVersion string,
) error {

	if err := ValidateProfilesChannelParams(ctx, am.keeper, order, portID, channelID, version); err != nil {
		return err
	}

	// Module may have already claimed capability in OnChanOpenInit in the case of crossing hellos
	// (ie chainA and chainB both call ChanOpenInit before one of them calls ChanOpenTry)
	// If module can already authenticate the capability then module already owns it so we don't need to claim
	// Otherwise, module does not have channel capability and we must claim it from IBC
	if !am.keeper.AuthenticateCapability(ctx, channelCap, host.ChannelCapabilityPath(portID, channelID)) {
		// Only claim channel capability passed back by IBC module if we do not already own it
		err := am.keeper.ClaimCapability(ctx, channelCap, host.ChannelCapabilityPath(portID, channelID))
		if err != nil {
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

// -------------------------------------------------------------------------------------------------------------------

// OnRecvPacket implements the IBCModule interface
func (am AppModule) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	ack := channeltypes.NewResultAcknowledgement([]byte{byte(1)})

	// Try handling the chain link packet data
	ack, err := am.HandlePacket(ctx, packet, handleOracleRequestPacketData, handleLinkChainAccountPacketData)
	if err != nil {
		ack = channeltypes.NewErrorAcknowledgement(err.Error())
	}

	// NOTE: acknowledgement will be written synchronously during IBC handler execution.
	return ack
}

// PacketHandler represents a method that tries handling a packet.
// If the packet has been handled properly, handled will be true and the acknowledgment and error will
// tell how the execution has ended.
// If the packet cannot be handled properly, false will be returned instead as first value.
type PacketHandler = func(
	am AppModule, ctx sdk.Context, packet channeltypes.Packet,
) (handled bool, ack channeltypes.Acknowledgement, err error)

// HandlePacket handles the given packet by passing it to the provided packet handlers one by one until
// at least one of them can handle it.
// If no handler supports the given packet, it returns types.ErrInvalidPacketData.
func (am AppModule) HandlePacket(
	ctx sdk.Context, packet channeltypes.Packet, packetHandlers ...PacketHandler,
) (channeltypes.Acknowledgement, error) {
	for _, handler := range packetHandlers {
		handled, ack, err := handler(am, ctx, packet)
		if handled {
			return ack, err
		}
	}
	return channeltypes.Acknowledgement{}, sdkerrors.Wrapf(types.ErrInvalidPacketData, "%T", packet)
}

// handleLinkChainAccountPacketData tries handling athe given packet by deserializing the inner data
// as a LinkChainAccountPacketData instance.
func handleLinkChainAccountPacketData(
	am AppModule, ctx sdk.Context, packet channeltypes.Packet,
) (handled bool, ack channeltypes.Acknowledgement, err error) {
	var packetData types.LinkChainAccountPacketData
	if err := types.ModuleCdc.UnmarshalJSON(packet.GetData(), &packetData); err != nil {
		return false, channeltypes.Acknowledgement{}, nil
	}

	var acknowledgement channeltypes.Acknowledgement

	packetAck, err := am.keeper.OnRecvLinkChainAccountPacket(ctx, packetData)
	if err != nil {
		acknowledgement = channeltypes.NewErrorAcknowledgement(err.Error())
	} else {
		// Encode packet acknowledgment
		packetAckBytes, err := packetAck.Marshal()
		if err != nil {
			return true, channeltypes.Acknowledgement{}, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
		acknowledgement = channeltypes.NewResultAcknowledgement(packetAckBytes)
	}

	address, err := types.UnpackAddressData(am.cdc, packetData.SourceAddress)
	if err != nil {
		return true, channeltypes.Acknowledgement{}, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeLinkChainAccountPacket,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeChainLinkSourceAddress, address.GetValue()),
			sdk.NewAttribute(types.AttributeChainLinkSourceChainName, packetData.SourceChainConfig.Name),
			sdk.NewAttribute(types.AttributeChainLinkDestinationAddress, packetData.DestinationAddress),
			sdk.NewAttribute(types.AttributeKeyAckSuccess, fmt.Sprintf("%t", true)),
		),
	)

	return true, acknowledgement, nil
}

// handleOracleRequestPacketData tries handling athe given packet by deserializing the inner data
// as an OracleResponsePacketData instance.
func handleOracleRequestPacketData(
	am AppModule, ctx sdk.Context, packet channeltypes.Packet,
) (handled bool, ack channeltypes.Acknowledgement, err error) {
	var data oracletypes.OracleResponsePacketData
	if err := types.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return false, channeltypes.Acknowledgement{}, nil
	}

	acknowledgement := channeltypes.NewResultAcknowledgement([]byte{byte(1)})

	err = am.keeper.OnRecvApplicationLinkPacketData(ctx, data)
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
	return true, acknowledgement, nil
}

// -------------------------------------------------------------------------------------------------------------------

// OnAcknowledgementPacket implements the IBCModule interface
func (am AppModule) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) (*sdk.Result, error) {
	var ack channeltypes.Acknowledgement
	err := types.ModuleCdc.UnmarshalJSON(acknowledgement, &ack)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
			"cannot unmarshal oracle packet acknowledgement: %v", err)
	}

	var data oracletypes.OracleRequestPacketData
	err = types.ModuleCdc.UnmarshalJSON(packet.GetData(), &data)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
			"cannot unmarshal oracle request packet data: %s", err.Error())
	}

	err = am.keeper.OnOracleRequestAcknowledgementPacket(ctx, data, ack)
	if err != nil {
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

// -------------------------------------------------------------------------------------------------------------------

// OnTimeoutPacket implements the IBCModule interface
func (am AppModule) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) (*sdk.Result, error) {
	var data oracletypes.OracleRequestPacketData
	err := types.ModuleCdc.UnmarshalJSON(packet.GetData(), &data)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
			"cannot unmarshal oracle request packet data: %s", err.Error())
	}

	err = am.keeper.OnOracleRequestTimeoutPacket(ctx, data)
	if err != nil {
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
