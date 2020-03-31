package types

// GenesisState contains the data of the genesis state for the profile module
type GenesisState struct {
	Profiles []Profile `json:"accounts"`
}

// NewGenesisState creates a new genesis state
func NewGenesisState(accounts []Profile) GenesisState {
	return GenesisState{
		Profiles: accounts,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data GenesisState) error {
	for _, account := range data.Profiles {
		if err := account.Validate(); err != nil {
			return err
		}
	}
	return nil
}
