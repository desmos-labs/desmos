package types

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

	return nil
}

// doesSubspaceExists check if the subspaces array contains a subspace with the given id
func doesSubspaceExists(subspaces []Subspace, id string) bool {
	for _, sub := range subspaces {
		if sub.ID == id {
			return true
		}
	}
	return false
}
