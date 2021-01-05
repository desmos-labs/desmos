package types

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
		err := rel.Validate()
		if err != nil {
			return err
		}
	}

	for _, ub := range data.Blocks {
		err := ub.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}
