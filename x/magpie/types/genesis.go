package types

import (
	"fmt"
)

// NewGenesisState allows to create a new genesis state containing the given default session length and sessions
func NewGenesisState(defaultSessionLength uint64, sessions []Session) *GenesisState {
	return &GenesisState{
		DefaultSessionLength: defaultSessionLength,
		Sessions:             sessions,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(240, nil)
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(state *GenesisState) error {
	if state.DefaultSessionLength == 0 {
		return fmt.Errorf("invalid default session length: %d", state.DefaultSessionLength)
	}

	return nil
}
