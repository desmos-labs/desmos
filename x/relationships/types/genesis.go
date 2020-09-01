package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState contains the data of the genesis state for the profile module
type GenesisState struct {
	UsersRelationships map[string][]sdk.AccAddress `json:"users_relationships"`
	UsersBlocks        []UserBlock                 `json:"users_blocks"`
}

// NewGenesisState creates a new genesis state
func NewGenesisState(usersRelationships map[string][]sdk.AccAddress, usersBlocks []UserBlock) GenesisState {
	return GenesisState{
		UsersRelationships: usersRelationships,
		UsersBlocks:        usersBlocks,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() GenesisState {
	return GenesisState{
		UsersRelationships: map[string][]sdk.AccAddress{},
		UsersBlocks:        []UserBlock{},
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

	for _, ub := range data.UsersBlocks {
		if err := ub.Validate(); err != nil {
			return err
		}
	}

	return nil
}
