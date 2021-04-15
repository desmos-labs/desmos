package v0130

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	ModuleName = "relationships"
)

// GenesisState contains the data of a v0.15.0 genesis state for the relationships module
type GenesisState struct {
	UsersRelationships map[string][]Relationship `json:"users_relationships"`
	UsersBlocks        []UserBlock               `json:"users_blocks"`
}

type Relationship struct {
	Recipient sdk.AccAddress `json:"recipient" yaml:"recipient"`
	Subspace  string         `json:"subspace" yaml:"subspace"`
}

type UserBlock struct {
	Blocker  sdk.AccAddress `json:"blocker" yaml:"blocker"`
	Blocked  sdk.AccAddress `json:"blocked" yaml:"blocked"`
	Reason   string         `json:"reason,omitempty" yaml:"reason,omitempty"`
	Subspace string         `json:"subspace" yaml:"subspace"`
}
