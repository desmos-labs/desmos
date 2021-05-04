package types

import (
	fmt "fmt"

	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// NewGenesisState creates a new genesis state
func NewGenesisState(portID string, links []Link) *GenesisState {
	return &GenesisState{
		PortId: portID,
		Links:  links,
	}
}

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(PortID, nil)
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(gs *GenesisState) error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}

	// Check for duplicated source in links
	linksSourceMap := make(map[string]bool)

	for _, elem := range gs.Links {
		if _, ok := linksSourceMap[elem.SourceAddress]; ok {
			return fmt.Errorf("duplicated source for links")
		}
		linksSourceMap[elem.SourceAddress] = true
	}

	return nil
}
