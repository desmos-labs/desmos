package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState contains the data of the genesis state for the profile module
type GenesisState struct {
	UsersRelationships map[string][]sdk.AccAddress `json:"users_relationships"`
}

// NewGenesisState creates a new genesis state
func NewGenesisState(usersRelationships map[string][]sdk.AccAddress) GenesisState {
	return GenesisState{
		UsersRelationships: usersRelationships,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() GenesisState {
	return GenesisState{
		UsersRelationships: map[string][]sdk.AccAddress{},
	}
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data GenesisState) error {
	for _, relationships := range data.UsersRelationships {
		for _, address := range relationships {
			if address.Empty() {
				return fmt.Errorf("invalid address %s", address)
			}
		}
	}

	return nil
}
