package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewLinkChainAccountPacketData(
	sourceAddress string,
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

// Validate is used for validating the packet
func (p LinkChainAccountPacketData) Validate() error {
	if strings.TrimSpace(p.SourceAddress) == "" {
		return fmt.Errorf("source address cannot be empty or blank")
	}
	if err := p.SourceProof.Validate(); err != nil {
		return err
	}
	if err := p.SourceChainConfig.Validate(); err != nil {
		return err
	}
	if _, err := sdk.AccAddressFromBech32(p.DestinationAddress); err != nil {
		return fmt.Errorf("invalid destination address: %s", p.DestinationAddress)
	}
	if err := p.DestinationProof.Validate(); err != nil {
		return err
	}

	return nil
}

// GetBytes is a helper for serialising
func (p LinkChainAccountPacketData) GetBytes() ([]byte, error) {
	var modulePacket LinkChainAccountPacketData
	return sdk.SortJSON(ProtoCdc.MustMarshalJSON(&modulePacket))
}
