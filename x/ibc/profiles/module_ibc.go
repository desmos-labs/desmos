package ibcprofiles

import (
	"fmt"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	porttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/05-port/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	"github.com/desmos-labs/desmos/x/ibc/profiles/keeper"
	"github.com/desmos-labs/desmos/x/ibc/profiles/types"
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
	if err := ValidateIBCProfilesChannelParams(ctx, am.keeper, order, portID, channelID, version); err != nil {
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

	if err := ValidateIBCProfilesChannelParams(ctx, am.keeper, order, portID, channelID, version); err != nil {
		return err
	}

	if counterpartyVersion != types.Version {
		return sdkerrors.Wrapf(types.ErrInvalidVersion, "invalid counterparty version: got: %s, expected %s", counterpartyVersion, types.Version)
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
	if counterpartyVersion != types.Version {
		return sdkerrors.Wrapf(types.ErrInvalidVersion, "invalid counterparty version: %s, expected %s", counterpartyVersion, types.Version)
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
	modulePacket channeltypes.Packet,
) (*sdk.Result, []byte, error) {
	var modulePacketData types.IBCProfilesPacketData
	if err := types.ProtoCdc.UnmarshalJSON(modulePacket.GetData(), &modulePacketData); err != nil {
		return nil, nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal packet data: %s", err.Error())
	}

	var ack channeltypes.Acknowledgement

	// Dispatch packet
	switch packet := modulePacketData.Packet.(type) {
	case *types.IBCProfilesPacketData_IbcAccountConnectionPacket:
		packetAck, err := am.keeper.OnRecvIBCAccountConnectionPacket(ctx, modulePacket, *packet.IbcAccountConnectionPacket)
		if err != nil {
			ack = channeltypes.NewErrorAcknowledgement(err.Error())
		} else {
			// Encode packet acknowledgment
			packetAckBytes, err := packetAck.Marshal()
			if err != nil {
				return nil, []byte{}, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
			}
			ack = channeltypes.NewResultAcknowledgement(packetAckBytes)
		}
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeIBCAccountConnectionPacket,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
				sdk.NewAttribute(types.AttributeKeyAckSuccess, fmt.Sprintf("%t", err != nil)),
			),
		)
	case *types.IBCProfilesPacketData_IbcAccountLinkPacket:
		packetAck, err := am.keeper.OnRecvIBCAccountLinkPacket(ctx, modulePacket, *packet.IbcAccountLinkPacket)
		if err != nil {
			ack = channeltypes.NewErrorAcknowledgement(err.Error())
		} else {
			// Encode packet acknowledgment
			packetAckBytes, err := packetAck.Marshal()
			if err != nil {
				return nil, []byte{}, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
			}
			ack = channeltypes.NewResultAcknowledgement(packetAckBytes)
		}
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeIBCAccountLinkPacket,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
				sdk.NewAttribute(types.AttributeKeyAckSuccess, fmt.Sprintf("%t", err != nil)),
			),
		)
	default:
		errMsg := fmt.Sprintf("unrecognized %s packet type: %T", types.ModuleName, packet)
		return nil, []byte{}, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
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

// OnAcknowledgementPacket implements the IBCModule interface
func (am AppModule) OnAcknowledgementPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	acknowledgement []byte,
) (*sdk.Result, error) {
	var ack channeltypes.Acknowledgement
	if err := types.ProtoCdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal packet acknowledgement: %v", err)
	}
	var modulePacketData types.IBCProfilesPacketData
	if err := types.ProtoCdc.UnmarshalJSON(modulePacket.GetData(), &modulePacketData); err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal packet data: %s", err.Error())
	}

	var eventType string

	// Dispatch packet
	switch packet := modulePacketData.Packet.(type) {
	case *types.IBCProfilesPacketData_IbcAccountConnectionPacket:
		err := am.keeper.OnAcknowledgementIBCAccountConnectionPacket(ctx, modulePacket, *packet.IbcAccountConnectionPacket, ack)
		if err != nil {
			return nil, err
		}
		eventType = types.EventTypeIBCAccountConnectionPacket
	case *types.IBCProfilesPacketData_IbcAccountLinkPacket:
		err := am.keeper.OnAcknowledgementIBCAccountLinkPacket(ctx, modulePacket, *packet.IbcAccountLinkPacket, ack)
		if err != nil {
			return nil, err
		}
		eventType = types.EventTypeIBCAccountLinkPacket
	default:
		errMsg := fmt.Sprintf("unrecognized %s packet type: %T", types.ModuleName, packet)
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			eventType,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeKeyAck, fmt.Sprintf("%v", ack)),
		),
	)

	switch resp := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Result:
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				eventType,
				sdk.NewAttribute(types.AttributeKeyAckSuccess, string(resp.Result)),
			),
		)
	case *channeltypes.Acknowledgement_Error:
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				eventType,
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
	modulePacket channeltypes.Packet,
) (*sdk.Result, error) {
	var modulePacketData types.IBCProfilesPacketData
	if err := types.ProtoCdc.UnmarshalJSON(modulePacket.GetData(), &modulePacketData); err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal packet data: %s", err.Error())
	}

	// Dispatch packet
	switch packet := modulePacketData.Packet.(type) {
	// Do nothing when timeout is reached
	case *types.IBCProfilesPacketData_IbcAccountConnectionPacket:
	case *types.IBCProfilesPacketData_IbcAccountLinkPacket:
	default:
		errMsg := fmt.Sprintf("unrecognized %s packet type: %T", types.ModuleName, packet)
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
	}

	return &sdk.Result{
		Events: ctx.EventManager().Events().ToABCIEvents(),
	}, nil
}

// ValidateIBCProfilesChannelParams does validation of a newly created ibcprofiles channel. A ibcprofiles
// channel must be UNORDERED, use the correct port (by default 'ibcprofiles'), and use the current
// supported version. Only 2^32 channels are allowed to be created.
func ValidateIBCProfilesChannelParams(
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
		return sdkerrors.Wrapf(types.ErrMaxIBCProfilesChannels, "channel sequence %d is greater than max allowed ibcprofiles channels %d", channelSequence, uint64(math.MaxUint32))
	}
	if order != channeltypes.UNORDERED {
		return sdkerrors.Wrapf(channeltypes.ErrInvalidChannelOrdering, "expected %s channel, got %s ", channeltypes.UNORDERED, order)
	}

	// Require portID is the portID ibcprofiles module is bound to
	boundPort := keeper.GetPort(ctx)
	if boundPort != portID {
		return sdkerrors.Wrapf(porttypes.ErrInvalidPort, "invalid port: %s, expected %s", portID, boundPort)
	}

	if version != types.Version {
		return sdkerrors.Wrapf(types.ErrInvalidVersion, "got %s, expected %s", version, types.Version)
	}
	return nil
}
