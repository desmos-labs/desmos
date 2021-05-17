package keeper

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	"github.com/desmos-labs/desmos/x/ibc/profiles/types"
)

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

	packetProof := []byte(data.SourceAddress)

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
