package types

import (
	"fmt"
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
)

// GenesisState contains the data of the genesis state for the profile module
type GenesisState struct {
	Profiles             []models.Profile            `json:"profiles"`
	NameSurnameLenParams models.NameSurnameLenParams `json:"name_surname_len_params"`
	MonikerLenParams     models.MonikerLenParams     `json:"moniker_len_params"`
	BioLenParams         models.BioLenParams         `json:"bio_len_params"`
}

// NewGenesisState creates a new genesis state
func NewGenesisState(profiles []models.Profile, nsp models.NameSurnameLenParams, mp models.MonikerLenParams, bp models.BioLenParams) GenesisState {
	return GenesisState{
		Profiles:             profiles,
		NameSurnameLenParams: nsp,
		MonikerLenParams:     mp,
		BioLenParams:         bp,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Profiles:             models.Profiles{},
		NameSurnameLenParams: models.DefaultNameSurnameLenParams(),
		MonikerLenParams:     models.DefaultMonikerLenParams(),
		BioLenParams:         models.DefaultBioLenParams(),
	}
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data GenesisState) error {
	for _, profile := range data.Profiles {
		if err := profile.Validate(); err != nil {
			return err
		}
	}

	// name/surname params validity checks
	if data.NameSurnameLenParams.MinNameSurnameLen.IsNegative() || data.NameSurnameLenParams.MinNameSurnameLen.LT(models.DefaultMinNameSurnameLength) {
		return fmt.Errorf("invalid minimum name/surname length param: %s", data.NameSurnameLenParams.MinNameSurnameLen)
	}

	if data.NameSurnameLenParams.MaxNameSurnameLen.IsNegative() || data.NameSurnameLenParams.MaxNameSurnameLen.GT(models.DefaultMaxNameSurnameLength) {
		return fmt.Errorf("invalid max name/surname length param: %s", data.NameSurnameLenParams.MaxNameSurnameLen)
	}

	// moniker validity checks
	if data.MonikerLenParams.MinMonikerLen.IsNegative() || data.MonikerLenParams.MinMonikerLen.LT(models.DefaultMinMonikerLength) {
		return fmt.Errorf("invalid minimum moniker length param: %s", data.MonikerLenParams.MinMonikerLen)
	}

	if data.MonikerLenParams.MaxMonikerLen.IsNegative() {
		return fmt.Errorf("invalid max moniker length param: %s", data.MonikerLenParams.MaxMonikerLen)
	}

	if data.BioLenParams.MaxBioLen.IsNegative() {
		return fmt.Errorf("invalid max bio length param: %s", data.BioLenParams.MaxBioLen)
	}

	return nil
}
