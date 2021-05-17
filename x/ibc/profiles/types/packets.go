package types

import (
	"encoding/hex"
	fmt "fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewIBCAccountConnectionPacketData(
	sourceChainPrefix string,
	sourceChainID string,
	sourceAddress string,
	sourcePubKey string,
	destinationAddress string,
	sourceSignature string,
	destinationSignature string,
) IBCAccountConnectionPacketData {
	return IBCAccountConnectionPacketData{
		SourceChainPrefix:    sourceChainPrefix,    // Bech32 prefix of the source chain
		SourceChainID:        sourceChainID,        // Chain ID of source chain
		SourceAddress:        sourceAddress,        // Bech32-encoded address
		SourcePubKey:         sourcePubKey,         // Hex-encoded public key related to the address
		DestinationAddress:   destinationAddress,   // Bech32-encoded  destination address
		SourceSignature:      sourceSignature,      // Hex-encoded signature by source key
		DestinationSignature: destinationSignature, // Hex-encoded signature by destination key
	}
}

// Validate is used for validating the packet
func (p IBCAccountConnectionPacketData) Validate() error {

	if p.SourceChainPrefix == "" {
		return fmt.Errorf("chain prefix cannot be empty")
	}

	if p.SourceChainID == "" {
		return fmt.Errorf("chain id cannot be empty")
	}

	srcAddrBz, err := sdk.GetFromBech32(p.SourceAddress, p.SourceChainPrefix)
	if err != nil {
		return fmt.Errorf("failed to parse source address")
	}
	srcAccAddr := sdk.AccAddress(srcAddrBz)

	srcPubKeyBz, err := hex.DecodeString(p.SourcePubKey)
	if err != nil {
		return fmt.Errorf("failed to decode source pubkey")
	}

	srcPubKey := &secp256k1.PubKey{Key: srcPubKeyBz}
	if !srcAccAddr.Equals(sdk.AccAddress(srcPubKey.Address().Bytes())) {
		return fmt.Errorf("source pubkey and source address are mismatched")
	}

	_, err = sdk.AccAddressFromBech32(p.DestinationAddress)
	if err != nil {
		return fmt.Errorf("failed to parse destination address")
	}

	srcSig, err := hex.DecodeString(p.SourceSignature)
	if err != nil {
		return fmt.Errorf("failed to parse decode source signature")
	}

	_, err = hex.DecodeString(p.DestinationSignature)
	if err != nil {
		return fmt.Errorf("failed to decode destination signature")
	}

	packetProof := []byte(p.SourceAddress)

	if !srcPubKey.VerifySignature(packetProof, srcSig) {
		return fmt.Errorf("failed to verify source signature")
	}

	return nil
}

// GetBytes is a helper for serialising
func (p IBCAccountConnectionPacketData) GetBytes() ([]byte, error) {
	var modulePacket IBCProfilesPacketData
	modulePacket.Packet = &IBCProfilesPacketData_IbcAccountConnectionPacket{&p}
	return sdk.SortJSON(ProtoCdc.MustMarshalJSON(&modulePacket))
}

// ___________________________________________________________________________________________________________________

func NewIBCAccountLinkPacketData(
	sourceChainPrefix string,
	sourceChainID string,
	sourceAddress string,
	sourcePubKey string,
	signature string,
) IBCAccountLinkPacketData {
	return IBCAccountLinkPacketData{
		SourceChainPrefix: sourceChainPrefix, // Bech32 prefix of the source chain
		SourceChainID:     sourceChainID,     // Chain ID of source chain
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

	if p.SourceChainID == "" {
		return fmt.Errorf("chain id cannot be empty")
	}

	srcAddrBz, err := sdk.GetFromBech32(p.SourceAddress, p.SourceChainPrefix)
	if err != nil {
		return fmt.Errorf("failed to parse source address")
	}
	srcAccAddr := sdk.AccAddress(srcAddrBz)

	srcPubKeyBz, err := hex.DecodeString(p.SourcePubKey)
	if err != nil {
		return fmt.Errorf("failed to decode source pubkey")
	}

	srcPubKey := &secp256k1.PubKey{Key: srcPubKeyBz}
	if !srcAccAddr.Equals(sdk.AccAddress(srcPubKey.Address().Bytes())) {
		return fmt.Errorf("source pubkey and source address are mismatched")
	}

	_, err = hex.DecodeString(p.Signature)
	if err != nil {
		return fmt.Errorf("failed to decode source signature")
	}

	return nil
}

// GetBytes is a helper for serialising
func (p IBCAccountLinkPacketData) GetBytes() ([]byte, error) {
	var modulePacket IBCProfilesPacketData
	modulePacket.Packet = &IBCProfilesPacketData_IbcAccountLinkPacket{&p}
	return sdk.SortJSON(ProtoCdc.MustMarshalJSON(&modulePacket))
}
