package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState contains the data of the genesis state for the profile module
type GenesisState struct {
	Profiles           []Profile                   `json:"profiles" yaml:"profiles"`
	Params             Params                      `json:"params" yaml:"params"`
	UsersRelationships map[string][]sdk.AccAddress `json:"users_relationships"`
}

// NewGenesisState creates a new genesis state
func NewGenesisState(profiles []Profile, params Params, usersRelationships map[string][]sdk.AccAddress) GenesisState {
	return GenesisState{
		Profiles:           profiles,
		Params:             params,
		UsersRelationships: usersRelationships,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Profiles:           Profiles{},
		Params:             DefaultParams(),
		UsersRelationships: map[string][]sdk.AccAddress{},
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

	for _, relationships := range data.UsersRelationships {
		for _, address := range relationships {
			if !address.Empty() {
				return fmt.Errorf("invalid address %s", address)
			}
		}
	}

	return nil
}
