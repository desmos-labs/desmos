package types

import (
	fmt "fmt"

	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId: PortID,
		Links:  []*Link{},
	}
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
