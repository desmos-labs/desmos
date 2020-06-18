package types

import (
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
)

// GenesisState contains the data of the genesis state for the profile module
type GenesisState struct {
	Profiles []models.Profile `json:"profiles" yaml:"profiles"`
	Params   models.Params    `json:"params" yaml:"params"`
}

// NewGenesisState creates a new genesis state
func NewGenesisState(profiles []models.Profile, params models.Params) GenesisState {
	return GenesisState{
		Profiles: profiles,
		Params:   params,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Profiles: models.Profiles{},
		Params:   DefaultParams(),
	}
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data GenesisState) error {
	for _, profile := range data.Profiles {
		if err := profile.Validate(); err != nil {
			return err
		}
	}

	if err := data.Params.Validate(); err != nil {
		return err
	}

	return nil
}
