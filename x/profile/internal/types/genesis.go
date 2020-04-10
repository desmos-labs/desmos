package types

// GenesisState contains the data of the genesis state for the profile module
type GenesisState struct {
	Profiles []Profile `json:"profiles"`
}

// NewGenesisState creates a new genesis state
func NewGenesisState(profiles []Profile) GenesisState {
	return GenesisState{
		Profiles: profiles,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data GenesisState) error {
	for _, profile := range data.Profiles {
		if err := profile.Validate(); err != nil {
			return err
		}
	}
	return nil
}
