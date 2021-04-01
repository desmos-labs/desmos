package types

import (
	fmt "fmt"

	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// NewGenesisState creates a new genesis state
func NewGenesisState(portId string, links []Link) *GenesisState {
	return &GenesisState{
		PortId: portId,
		Links:  links,
	}
}

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(PortID, nil)
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}

	// Check for duplicated source in links
	linksSourceMap := make(map[string]bool)

	for _, elem := range gs.Links {
		if _, ok := linksSourceMap[elem.Source]; ok {
			return fmt.Errorf("duplicated source for links")
		}
		linksSourceMap[elem.Source] = true
	}

	return nil
}
