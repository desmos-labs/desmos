package types

import "fmt"

// NewGenesisState creates a new genesis state
func NewGenesisState(subspaces []Subspace) *GenesisState {
	return &GenesisState{
		Subspaces: subspaces,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil)
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data *GenesisState) error {
	for _, subspace := range data.Subspaces {
		if err := subspace.Validate(); err != nil {
			return err
		}
	}

	for _, subspace := range data.Subspaces {
		if containDuplicates(data.Subspaces, subspace) {
			return fmt.Errorf("duplicated subspace: %s", subspace.ID)
		}
	}

	return nil
}

// containDuplicates tells whether the given subspaces slice contain duplicates of the provided subspace
func containDuplicates(subspaces []Subspace, subspace Subspace) bool {
	var count = 0
	for _, s := range subspaces {
		if s.Equal(subspace) {
			count++
		}
	}
	return count > 1
}
