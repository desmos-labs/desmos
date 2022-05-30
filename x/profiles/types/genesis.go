package types

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
)

// NewGenesisState creates a new genesis state
func NewGenesisState(
	requests []DTagTransferRequest,
	params Params, portID string,
	chainLinks []ChainLink, applicationLinks []ApplicationLink,
) *GenesisState {
	return &GenesisState{
		Params:               params,
		DTagTransferRequests: requests,
		IBCPortID:            portID,
		ChainLinks:           chainLinks,
		ApplicationLinks:     applicationLinks,
	}
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (g GenesisState) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	for _, link := range g.ChainLinks {
		err := link.UnpackInterfaces(unpacker)
		if err != nil {
			return err
		}
	}
	return nil
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, DefaultParams(), IBCPortID, nil, nil)
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data *GenesisState) error {
	err := data.Params.Validate()
	if err != nil {
		return err
	}

	for _, req := range data.DTagTransferRequests {
		err = req.Validate()
		if err != nil {
			return err
		}
	}

	err = host.PortIdentifierValidator(data.IBCPortID)
	if err != nil {
		return err
	}

	for _, l := range data.ChainLinks {
		err := l.Validate()
		if err != nil {
			return err
		}
	}

	for _, link := range data.ApplicationLinks {
		err = link.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}
