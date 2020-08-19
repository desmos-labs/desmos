package types

import "fmt"

// GenesisState contains the data of the genesis state for the profile module
type GenesisState struct {
	Profiles           []Profile                  `json:"profiles" yaml:"profiles"`
	Params             Params                     `json:"params" yaml:"params"`
	Relationships      Relationships              `json:"relationships"`
	UsersRelationships map[string]RelationshipIDs `json:"users_relationships"`
}

// NewGenesisState creates a new genesis state
func NewGenesisState(profiles []Profile, params Params, relationships Relationships, usersRelationships map[string]RelationshipIDs) GenesisState {
	return GenesisState{
		Profiles:           profiles,
		Params:             params,
		Relationships:      relationships,
		UsersRelationships: usersRelationships,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Profiles:           Profiles{},
		Params:             DefaultParams(),
		Relationships:      Relationships{},
		UsersRelationships: map[string]RelationshipIDs{},
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

	for _, rel := range data.Relationships {
		if err := rel.Validate(); err != nil {
			return err
		}
	}

	for _, ids := range data.UsersRelationships {
		for _, id := range ids {
			if !id.Valid() {
				return fmt.Errorf("invalid relationshipID %s", id)
			}
		}
	}

	return nil
}
