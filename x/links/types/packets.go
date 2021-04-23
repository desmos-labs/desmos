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
		return fmt.Errorf("chain prefix cannot be empty")
	}

	sourceAddressBytes, err := sdk.GetFromBech32(p.SourceAddress, p.SourceChainPrefix)
	if err != nil {
		return fmt.Errorf("failed to source address")
	}
	sourceAccAddress := sdk.AccAddress(sourceAddressBytes)

	sourcePubKeyBytes, err := hex.DecodeString(p.SourcePubKey)
	if err != nil {
		return fmt.Errorf("failed to source pubkey")
	}

	sourcePubKey := &secp256k1.PubKey{Key: sourcePubKeyBytes}
	if !sourceAccAddress.Equals(sdk.AccAddress(sourcePubKey.Address().Bytes())) {
		return fmt.Errorf("source pubkey and source address are mismatched")
	}

	destinationAddress, err := hex.DecodeString(p.DestinationAddress)
	if err != nil {
		return fmt.Errorf("failed to destination address")
	}

	_, err = sdk.AccAddressFromBech32(string(destinationAddress))
	if err != nil {
		return fmt.Errorf("failed to parse destination address")
	}

	_, err = hex.DecodeString(p.SourceSignature)
	if err != nil {
		return fmt.Errorf("failed to parse decode source signature")
	}

	_, err = hex.DecodeString(p.DestinationSignature)
	if err != nil {
		return fmt.Errorf("failed to decode destination signature")
	}

	return nil
}

// GetBytes is a helper for serialising
func (p IBCAccountConnectionPacketData) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&p))
}

// ___________________________________________________________________________________________________________________

func NewIBCAccountLinkPacketData(
	sourceChainPrefix string,
	sourceAddress string,
	sourcePubKey string,
	signature string,
) IBCAccountLinkPacketData {
	return IBCAccountLinkPacketData{
		SourceChainPrefix: sourceChainPrefix, // Bech32 prefix of the source chain
		SourceAddress:     sourceAddress,     // Bech32-encoded address
		SourcePubKey:      sourcePubKey,      // Hex-encoded public key related to the address
		Signature:         signature,         // Hex-encoded signature by source key
	}
}

// Validate is used for validating the packet
func (p IBCAccountLinkPacketData) Validate() error {

	if p.SourceChainPrefix == "" {
		return fmt.Errorf("chain prefix cannot be empty")
	}

	sourceAddressBytes, err := sdk.GetFromBech32(p.SourceAddress, p.SourceChainPrefix)
	if err != nil {
		return fmt.Errorf("failed to source address")
	}
	sourceAccAddress := sdk.AccAddress(sourceAddressBytes)

	sourcePubKeyBytes, err := hex.DecodeString(p.SourcePubKey)
	if err != nil {
		return fmt.Errorf("failed to source pubkey")
	}

	sourcePubKey := &secp256k1.PubKey{Key: sourcePubKeyBytes}
	if !sourceAccAddress.Equals(sdk.AccAddress(sourcePubKey.Address().Bytes())) {
		return fmt.Errorf("source pubkey and source address are mismatched")
	}

	_, err = hex.DecodeString(p.Signature)
	if err != nil {
		return fmt.Errorf("failed to parse decode source signature")
	}

	return nil
}

// GetBytes is a helper for serialising
func (p IBCAccountLinkPacketData) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&p))
}

// ___________________________________________________________________________________________________________________

func VerifySignature(msg []byte, sig []byte, pubKey cryptotypes.PubKey) bool {
	if !pubKey.VerifySignature(msg, sig) {
		return false
	}
	return true
}
