package types

// GenesisState contains the data of the genesis state for the profile module
type GenesisState struct {
	Profiles             []Profile            `json:"profiles"`
	NameSurnameLenParams NameSurnameLenParams `json:"name_surname_len_params"`
	MonikerLenParams     MonikerLenParams     `json:"moniker_len_params"`
	BioLenParams         BioLenParams         `json:"bio_len_params"`
}

// NewGenesisState creates a new genesis state
func NewGenesisState(profiles []Profile, nsp NameSurnameLenParams, mp MonikerLenParams, bp BioLenParams) GenesisState {
	return GenesisState{
		Profiles:             profiles,
		NameSurnameLenParams: nsp,
		MonikerLenParams:     mp,
		BioLenParams:         bp,
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
