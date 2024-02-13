package utils

// DONTCOVER

import (
	"os"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/desmos-labs/desmos/v7/x/profiles/types"

	"github.com/cosmos/cosmos-sdk/codec"
)

// NewChainLinkJSON allows to build a new ChainLinkJSON instance
//
//nolint:interfacer
func NewChainLinkJSON(data types.AddressData, proof types.Proof, chainConfig types.ChainConfig) ChainLinkJSON {
	any, err := codectypes.NewAnyWithValue(data)
	if err != nil {
		panic(err)
	}

	return ChainLinkJSON{
		Address:     any,
		Proof:       proof,
		ChainConfig: chainConfig,
	}
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (link *ChainLinkJSON) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	if link.Address != nil {
		var address types.AddressData
		err := unpacker.UnpackAny(link.Address, &address)
		if err != nil {
			return err
		}
	}

	err := link.Proof.UnpackInterfaces(unpacker)
	if err != nil {
		return err
	}

	return nil
}

// ParseChainLinkJSON reads and parses a ChainLinkJSON from file.
func ParseChainLinkJSON(cdc codec.Codec, dataFile string) (ChainLinkJSON, error) {
	var data ChainLinkJSON

	contents, err := os.ReadFile(dataFile)
	if err != nil {
		return data, err
	}

	if err := cdc.UnmarshalJSON(contents, &data); err != nil {
		return data, err
	}

	return data, nil
}
