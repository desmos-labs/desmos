package types

// GenesisState contains the data of the genesis state for the posts module
type GenesisState struct {
	Reports map[string]Reports `json:"reports" yaml:"reports"`
}

// NewGenesisState creates a new genesis state
func NewGenesisState(reports map[string]Reports) GenesisState {
	return GenesisState{
		Reports: reports,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data GenesisState) error {
	for _, reports := range data.Reports {
		if err := reports.Validate(); err != nil {
			return err
		}
	}
	return nil
}
