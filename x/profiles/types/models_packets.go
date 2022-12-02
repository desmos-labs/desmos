// Package types
//nolint:interfacer
package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewLinkChainAccountPacketData returns a new LinkChainAccountPacketData instance
func NewLinkChainAccountPacketData(
	sourceAddress Address,
	sourceProof Proof,
	sourceChainConfig ChainConfig,
	destinationAddress string,
	destinationProof Proof,
) LinkChainAccountPacketData {
	return LinkChainAccountPacketData{
		SourceAddress:      sourceAddress,
		SourceProof:        sourceProof,
		SourceChainConfig:  sourceChainConfig,
		DestinationAddress: destinationAddress,
		DestinationProof:   destinationProof,
	}
}

// Validate validates the LinkChainAccountPacketData
func (p LinkChainAccountPacketData) Validate() error {
	err := p.SourceAddress.Validate()
	if err != nil {
		return fmt.Errorf("invalid source address: %s", err)
	}

	err = p.SourceProof.Validate()
	if err != nil {
		return fmt.Errorf("invalid source proof: %s", err)
	}

	err = p.SourceChainConfig.Validate()
	if err != nil {
		return fmt.Errorf("invalid source chain config: %s", err)
	}

	_, err = sdk.AccAddressFromBech32(p.DestinationAddress)
	if err != nil {
		return fmt.Errorf("invalid destination address: %s", p.DestinationAddress)
	}

	err = p.DestinationProof.Validate()
	if err != nil {
		return fmt.Errorf("invalid destination proof: %s", err)
	}

	return nil
}

// GetBytes is a helper for serialising
func (p LinkChainAccountPacketData) GetBytes() ([]byte, error) {
	var modulePacket LinkChainAccountPacketData
	return sdk.SortJSON(ModuleCdc.MustMarshalJSON(&modulePacket))
}
