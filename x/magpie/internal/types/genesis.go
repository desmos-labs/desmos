package types

import (
	"fmt"
)

// GenesisState represents the genesis state for the magpie module
type GenesisState struct {
	DefaultSessionLength int64    `json:"default_session_length"`
	Sessions             Sessions `json:"sessions"`
}

// NewGenesisState allows to create a new genesis state containing the given default session length and sessions
func NewGenesisState(defaultSessionLength int64, sessions Sessions) GenesisState {
	return GenesisState{
		DefaultSessionLength: defaultSessionLength,
		Sessions:             sessions,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() GenesisState {
	return GenesisState{
		DefaultSessionLength: 240, // 24 hours, counting a 6 secs block interval
	}
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(state GenesisState) error {
	if state.DefaultSessionLength <= 0 {
		return fmt.Errorf("invalid default session length: %d", state.DefaultSessionLength)
	}

	return nil
}
