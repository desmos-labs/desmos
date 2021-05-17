package types

import (
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// NewGenesisState creates a new genesis state
func NewGenesisState(portID string) *GenesisState {
	return &GenesisState{
		PortID: portID,
	}
}

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(PortID)
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(gs *GenesisState) error {
	if err := host.PortIdentifierValidator(gs.PortID); err != nil {
		return err
	}
	return nil
}
