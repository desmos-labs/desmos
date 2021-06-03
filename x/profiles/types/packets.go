package types

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewLinkChainAccountPacketData(
	sourceAddress AddressData,
	sourceProof Proof,
	sourceChainConfig ChainConfig,
	destinationAddress string,
	destinationProof Proof,
) LinkChainAccountPacketData {
	addressAny, err := codectypes.NewAnyWithValue(sourceAddress)
	if err != nil {
		panic("failed to pack public key to any type")
	}
	return LinkChainAccountPacketData{
		SourceAddress:      addressAny,
		SourceProof:        sourceProof,
		SourceChainConfig:  sourceChainConfig,
		DestinationAddress: destinationAddress,
		DestinationProof:   destinationProof,
	}
}

// Validate is used for validating the packet
func (p LinkChainAccountPacketData) Validate() error {
	if p.SourceAddress == nil {
		return fmt.Errorf("source address cannot be nil")
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
