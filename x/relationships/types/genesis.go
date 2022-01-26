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
	return NewGenesisState(nil, nil)
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data *GenesisState) error {
	for i, rel := range data.Relationships {
		if containDuplicates(data.Relationships, rel) {
			return fmt.Errorf("duplicated relationship: %s", &data.Relationships[i])
		}

		err := rel.Validate()
		if err != nil {
			return err
		}
	}

	for _, ub := range data.Blocks {
		// TODO: Check duplicates

		err := ub.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

// containDuplicates tells whether the given relationships slice contain duplicates of the provided relationship
func containDuplicates(relationships []Relationship, relationship Relationship) bool {
	var count = 0
	for _, r := range relationships {
		if r.Equal(relationship) {
			count++
		}
	}
	return count > 1
}
