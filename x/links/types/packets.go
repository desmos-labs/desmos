package types

import (
	"encoding/hex"
	fmt "fmt"

	"github.com/cosmos/cosmos-sdk/codec/legacy"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewIBCAccountConnectionPacketData(
	sourceChainPrefix string,
	sourceAddress string,
	sourcePubkey string,
	destinationAddress string,
	sourceSignature string,
	destinationSignature string,
) IBCAccountConnectionPacketData {
	return IBCAccountConnectionPacketData{
		SourceChainPrefix:    sourceChainPrefix,
		SourceAddress:        sourceAddress,
		SourcePubkey:         hex.EncodeToString([]byte(sourcePubkey)),
		DestinationAddress:   hex.EncodeToString([]byte(destinationAddress)),
		SourceSignature:      hex.EncodeToString([]byte(sourceSignature)),
		DestinationSignature: hex.EncodeToString([]byte(destinationSignature)),
	}
}

// Validate is used for validating the packet
func (p IBCAccountConnectionPacketData) Validate() error {

	destinationPubkey, err := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeAccPub, p.DestinationAddress)
	if err != nil {
		return err
	}

	destinationSignature, err := hex.DecodeString(p.DestinationSignature)
	if err != nil {
		return err
	}

	sourceSignature, err := hex.DecodeString(p.SourceSignature)
	if err != nil {
		return err
	}

	if !destinationPubkey.VerifySignature(sourceSignature, destinationSignature) {
		return fmt.Errorf("verify failed with destination pubkey: %s", p.DestinationAddress)
	}

	sourcePubkeyBytes, err := hex.DecodeString(p.SourcePubkey)
	if err != nil {
		return err
	}

	sourcePubkey, err := legacy.PubKeyFromBytes(sourcePubkeyBytes)
	if err != nil {
		return err
	}

	link := NewLink(p.SourceAddress, p.DestinationAddress)
	linkBytes, err := link.Marshal()
	if err != nil {
		return err
	}

	if !sourcePubkey.VerifySignature(linkBytes, destinationSignature) {
		return fmt.Errorf("verify failed with destination pubkey: %s", destinationSignature)
	}

	return nil
}

// GetBytes is a helper for serialising
func (p IBCAccountConnectionPacketData) GetBytes() ([]byte, error) {
	var modulePacket LinksPacketData

	modulePacket.Packet = &LinksPacketData_IbcAccountConnectionPacket{&p}

	return modulePacket.Marshal()
}
