package types

import (
	"fmt"
)

// NewGenesisState creates a new genesis state
func NewGenesisState(relationships []Relationship, blocks []UserBlock) *GenesisState {
	return &GenesisState{
		Relationships: relationships,
		Blocks:        blocks,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() *GenesisState {
	return NewGenesisState([]Relationship{}, []UserBlock{})
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data *GenesisState) error {
	for _, rel := range data.Relationships {
		if len(rel.Recipient) == 0 {
			return fmt.Errorf("invalid relationship's recipient address %s", rel)
		}
	}

	for _, ub := range data.Blocks {
		if err := ub.Validate(); err != nil {
			return err
		}
	}

	return nil
}
