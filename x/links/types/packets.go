package types

import (
	"encoding/hex"
	fmt "fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewIBCAccountConnectionPacketData(
	sourceChainPrefix string,
	sourceAddress string,
	sourcePubKey string,
	destinationAddress string,
	sourceSignature string,
	destinationSignature string,
) IBCAccountConnectionPacketData {
	return IBCAccountConnectionPacketData{
		SourceChainPrefix:    sourceChainPrefix,                              // Bech32 prefix of the source chain
		SourceAddress:        sourceAddress,                                  // Bech32-encoded address
		SourcePubKey:         sourcePubKey,                                   // Hex-encoded public key related to the address
		DestinationAddress:   hex.EncodeToString([]byte(destinationAddress)), // hex of destination address
		SourceSignature:      sourceSignature,                                // Hex-encoded signature by source key
		DestinationSignature: destinationSignature,                           // Hex-encoded signature by destination key
	}
}

// Validate is used for validating the packet
func (p IBCAccountConnectionPacketData) Validate() error {

	if p.SourceChainPrefix == "" {
		return fmt.Errorf("invalid chain prefix")
	}

	sourceAddressBytes, err := sdk.GetFromBech32(p.SourceAddress, p.SourceChainPrefix)
	if err != nil {
		return fmt.Errorf("source address parse failed")
	}
	sourceAccAddress := sdk.AccAddress(sourceAddressBytes)

	sourcePubKeyBytes, err := hex.DecodeString(p.SourcePubKey)
	if err != nil {
		return fmt.Errorf("source pubkey decode failed")
	}

	sourcePubKey := &secp256k1.PubKey{Key: sourcePubKeyBytes}
	if !sourceAccAddress.Equals(sdk.AccAddress(sourcePubKey.Address().Bytes())) {
		return fmt.Errorf("source pubkey mismatch with source address")
	}

	destinationAddress, err := hex.DecodeString(p.DestinationAddress)
	if err != nil {
		return fmt.Errorf("destination address decode failed")
	}

	_, err = sdk.AccAddressFromBech32(string(destinationAddress))
	if err != nil {
		return fmt.Errorf("destination address parse failed")
	}

	_, err = hex.DecodeString(p.SourceSignature)
	if err != nil {
		return fmt.Errorf("source signature decode failed")
	}

	_, err = hex.DecodeString(p.DestinationSignature)
	if err != nil {
		return fmt.Errorf("destination signature decode failed")
	}

	return nil
}

func VerifySignature(msg []byte, sig []byte, pubKey cryptotypes.PubKey) bool {
	if !pubKey.VerifySignature(msg, sig) {
		return false
	}
	return true
}

// GetBytes is a helper for serialising
func (p IBCAccountConnectionPacketData) GetBytes() ([]byte, error) {
	var modulePacket LinksPacketData

	modulePacket.Packet = &LinksPacketData_IbcAccountConnectionPacket{&p}

	return modulePacket.Marshal()
}
