package types

import (
	"fmt"

	"github.com/desmos-labs/desmos/x/relationships/types/models"
)

// GenesisState contains the data of the genesis state for the profile module
type GenesisState struct {
	UsersRelationships map[string]models.Relationships `json:"users_relationships"`
}

// NewGenesisState creates a new genesis state
func NewGenesisState(usersRelationships map[string]models.Relationships) GenesisState {
	return GenesisState{
		UsersRelationships: usersRelationships,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() GenesisState {
	return GenesisState{
		UsersRelationships: map[string]models.Relationships{},
	}
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data GenesisState) error {
	for _, relationships := range data.UsersRelationships {
		for _, rel := range relationships {
			if rel.Recipient.Empty() {
				return fmt.Errorf("invalid relationship's recipient address %s", rel)
			}
		}
	}

	return nil
}
