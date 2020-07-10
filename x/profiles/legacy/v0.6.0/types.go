package v060

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "profile"
)

// GenesisState contains the data of a v0.6.0 genesis state for the profile module
type GenesisState struct {
	Profiles []Profile `json:"profiles"`
}

// Profile is a struct for a Profile
type Profile struct {
	Moniker  string         `json:"moniker" yaml:"moniker"`
	Name     *string        `json:"name,omitempty" yaml:"name,omitempty"`
	Surname  *string        `json:"surname,omitempty" yaml:"surname,omitempty"`
	Bio      *string        `json:"bio,omitempty" yaml:"bio,omitempty"`
	Pictures *Pictures      `json:"pictures,omitempty" yaml:"pictures,omitempty"`
	Creator  sdk.AccAddress `json:"creator" yaml:"creator"`
}

// Pictures is a struct for Profile Pictures
type Pictures struct {
	Profile *string `json:"profile,omitempty" yaml:"profile,omitempty"`
	Cover   *string `json:"cover,omitempty" yaml:"cover,omitempty"`
}
