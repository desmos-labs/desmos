package keeper

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	"github.com/desmos-labs/desmos/x/links/types"
)

// TransmitIBCAccountConnectionPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitIBCAccountConnectionPacket(
	ctx sdk.Context,
	packetData types.IBCAccountConnectionPacketData,
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

	// get the next sequence
	sequence, found := k.channelKeeper.GetNextSequenceSend(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", sourcePort, sourceChannel,
		)
	}

	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: "+err.Error())
	}

	packet := channeltypes.NewPacket(
		packetBytes,
		sequence,
		sourcePort,
		sourceChannel,
		destinationPort,
		destinationChannel,
		timeoutHeight,
		timeoutTimestamp,
	)

	if err := k.channelKeeper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}

// OnRecvIBCAccountConnectionPacket processes packet reception
func (k Keeper) OnRecvIBCAccountConnectionPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IBCAccountConnectionPacketData) (packetAck types.IBCAccountConnectionPacketData, err error) {

	// validate packet data upon receiving
	if err := data.Validate(); err != nil {
		return packetAck, err
	}

	sourcePubkeyBytes, _ := hex.DecodeString(data.SourcePubKey)
	destinationAddress, _ := hex.DecodeString(data.DestinationAddress)
	sourceSignature, _ := hex.DecodeString(data.SourceSignature)
	destSignature, _ := hex.DecodeString(data.DestinationSignature)
	link := types.NewLink(data.SourceAddress, string(destinationAddress))

	linkBytes, _ := link.Marshal()
	sourcePubkey := &secp256k1.PubKey{Key: sourcePubkeyBytes}
	destinationAccAddress, _ := sdk.AccAddressFromBech32(string(destinationAddress))

	destinationPubkey, err := k.GetLinkPubKey(ctx, destinationAccAddress)
	if err != nil {
		return packetAck, err
	}

	if !types.VerifySignature(sourceSignature, destSignature, destinationPubkey) {
		return packetAck, fmt.Errorf("verify destination failed")
	}

	if !types.VerifySignature(linkBytes, sourceSignature, sourcePubkey) {
		return packetAck, fmt.Errorf("verify source Signature failed")
	}

	k.StoreLink(ctx, link)

	return packetAck, nil
}

// OnAcknowledgementIBCAccountConnectionPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementIBCAccountConnectionPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IBCAccountConnectionPacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.IBCAccountConnectionPacketData
		err := packetAck.Unmarshal(dispatchedAck.Result)
		if err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutIBCAccountConnectionPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutIBCAccountConnectionPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IBCAccountConnectionPacketData) error {

	// TODO: packet timeout logic

	return nil
}

// ___________________________________________________________________________________________________________________

// TransmitIBCAccountLinkPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitIBCAccountLinkPacket(
	ctx sdk.Context,
	packetData types.IBCAccountLinkPacketData,
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

	// get the next sequence
	sequence, found := k.channelKeeper.GetNextSequenceSend(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", sourcePort, sourceChannel,
		)
	}

	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: "+err.Error())
	}

	packet := channeltypes.NewPacket(
		packetBytes,
		sequence,
		sourcePort,
		sourceChannel,
		destinationPort,
		destinationChannel,
		timeoutHeight,
		timeoutTimestamp,
	)

	if err := k.channelKeeper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}
