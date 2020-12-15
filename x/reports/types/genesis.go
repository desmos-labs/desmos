package types

// NewGenesisState creates a new genesis state
func NewGenesisState(reports []Report) *GenesisState {
	return &GenesisState{
		Reports: reports,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() *GenesisState {
	return NewGenesisState([]Report{})
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data *GenesisState) error {
	for _, reports := range data.Reports {
		if err := reports.Validate(); err != nil {
			return err
		}
	}
	return nil
}
