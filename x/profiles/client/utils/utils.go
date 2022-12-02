package utils

// DONTCOVER

import (
	"io/ioutil"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/desmos-labs/desmos/v4/x/profiles/types"

	"github.com/cosmos/cosmos-sdk/codec"
)

// NewChainLinkJSON allows to build a new ChainLinkJSON instance
//nolint:interfacer
func NewChainLinkJSON(address types.Address, proof types.Proof, chainConfig types.ChainConfig) ChainLinkJSON {
	return ChainLinkJSON{
		Address:     address,
		Proof:       proof,
		ChainConfig: chainConfig,
	}
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (link *ChainLinkJSON) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	err := link.Address.UnpackInterfaces(unpacker)
	if err != nil {
		return err
	}

	err = link.Proof.UnpackInterfaces(unpacker)
	if err != nil {
		return err
	}

	return nil
}

// ParseChainLinkJSON reads and parses a ChainLinkJSON from file.
func ParseChainLinkJSON(cdc codec.Codec, dataFile string) (ChainLinkJSON, error) {
	var data ChainLinkJSON

	contents, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return data, err
	}

	if err := cdc.UnmarshalJSON(contents, &data); err != nil {
		return data, err
	}

	return data, nil
}
