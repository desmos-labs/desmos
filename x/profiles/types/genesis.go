package types

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
)

// NewGenesisState creates a new genesis state
func NewGenesisState(
	requests []DTagTransferRequest,
	params Params, portID string,
	chainLinks []ChainLink, defaultExternalAddresses []DefaultExternalAddressEntry,
	applicationLinks []ApplicationLink,
) *GenesisState {
	return &GenesisState{
		Params:                   params,
		DTagTransferRequests:     requests,
		IBCPortID:                portID,
		ChainLinks:               chainLinks,
		DefaultExternalAddresses: defaultExternalAddresses,
		ApplicationLinks:         applicationLinks,
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
	return NewGenesisState(nil, DefaultParams(), IBCPortID, nil, nil, nil)
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

	for _, entry := range data.DefaultExternalAddresses {
		err := entry.Validate()
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

// --------------------------------------------------------------------------------------------------------------------

// NewDefaultExternalAddressEntry returns a new DefaultExternalAddressEntry instance
func NewDefaultExternalAddressEntry(owner, chainName, target string) DefaultExternalAddressEntry {
	return DefaultExternalAddressEntry{
		Owner:     owner,
		ChainName: chainName,
		Target:    target,
	}
}

// Validate implements fmt.Validator
func (data DefaultExternalAddressEntry) Validate() error {
	_, err := sdk.AccAddressFromBech32(data.Owner)
	if err != nil {
		return fmt.Errorf("invalid owner: %s", data.Owner)
	}

	if strings.TrimSpace(data.ChainName) == "" {
		return fmt.Errorf("invalid chain name: %s", data.ChainName)
	}

	if strings.TrimSpace(data.Target) == "" {
		return fmt.Errorf("invalid external address target: %s", data.Target)
	}
	return nil
}
