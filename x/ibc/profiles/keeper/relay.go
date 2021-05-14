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
	"github.com/desmos-labs/desmos/x/ibc/profiles/types"
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
func (k Keeper) OnRecvIBCAccountConnectionPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	data types.IBCAccountConnectionPacketData,
) (packetAck types.IBCAccountConnectionPacketAck, err error) {

	// validate packet data upon receiving
	if err := data.Validate(); err != nil {
		return packetAck, err
	}

	srcSig, _ := hex.DecodeString(data.SourceSignature)
	destSig, _ := hex.DecodeString(data.DestinationSignature)
	destAccAddr, _ := sdk.AccAddressFromBech32(data.DestinationAddress)

	destPubkey, err := k.GetAccountPubKey(ctx, destAccAddr)
	if err != nil {
		return packetAck, err
	}

	// Signature should be verified here because source chain doesn't know the pubkey on the destination chain
	if !destPubkey.VerifySignature(srcSig, destSig) {
		return packetAck, fmt.Errorf("failed to verify destination signature")
	}

	packetAck.SourceAddress = data.SourceAddress

	return packetAck, nil
}

// OnAcknowledgementIBCAccountConnectionPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementIBCAccountConnectionPacket(ctx sdk.Context,
	packet channeltypes.Packet,
	data types.IBCAccountConnectionPacketData,
	ack channeltypes.Acknowledgement,
) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:
		return errors.New(dispatchedAck.Error)
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.IBCAccountConnectionPacketAck
		err := packetAck.Unmarshal(dispatchedAck.Result)
		if err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}
		// the acknowledgement succeeded on the receiving chain so nothing
		// needs to be executed and no error needs to be returned
		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutIBCAccountConnectionPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutIBCAccountConnectionPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	data types.IBCAccountConnectionPacketData,
) error {
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

	srcChannelEnd, found := k.channelKeeper.GetChannel(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", sourcePort, sourceChannel)
	}

	destPort := srcChannelEnd.GetCounterparty().GetPortID()
	destChannel := srcChannelEnd.GetCounterparty().GetChannelID()

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

	packetBz, _ := packetData.GetBytes()
	packet := channeltypes.NewPacket(
		packetBz,
		sequence,
		sourcePort,
		sourceChannel,
		destPort,
		destChannel,
		timeoutHeight,
		timeoutTimestamp,
	)

	if err := k.channelKeeper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}

// OnRecvIBCAccountLinkPacket processes packet reception
func (k Keeper) OnRecvIBCAccountLinkPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	data types.IBCAccountLinkPacketData,
) (packetAck types.IBCAccountLinkPacketAck, err error) {

	// validate packet data upon receiving
	if err := data.Validate(); err != nil {
		return packetAck, err
	}

	srcPubkeyBz, _ := hex.DecodeString(data.SourcePubKey)
	sig, _ := hex.DecodeString(data.Signature)
	srcPubKey := &secp256k1.PubKey{Key: srcPubkeyBz}

	destAddr := sdk.AccAddress(srcPubKey.Address().Bytes()).String()

	packetProof := []byte(data.SourceAddress + "-" + destAddr)

	// Signature should be verified here because source chain doesn't know the destination of packet
	if !srcPubKey.VerifySignature(packetProof, sig) {
		return packetAck, fmt.Errorf("failed to verify source signature")
	}

	packetAck.SourceAddress = data.SourceAddress

	return packetAck, nil
}

// OnAcknowledgementIBCAccountLinkPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementIBCAccountLinkPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	data types.IBCAccountLinkPacketData,
	ack channeltypes.Acknowledgement,
) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:
		return errors.New(dispatchedAck.Error)
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.IBCAccountLinkPacketAck
		err := packetAck.Unmarshal(dispatchedAck.Result)
		if err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// the acknowledgement succeeded on the receiving chain so nothing
		// needs to be executed and no error needs to be returned
		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutIBCAccountLinkPacket responds to the case where a packet has not been transmitted because of a timeout
// No error needs to be returned
func (k Keeper) OnTimeoutIBCAccountLinkPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	data types.IBCAccountLinkPacketData,
) error {
	return nil
}
