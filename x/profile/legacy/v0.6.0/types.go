package v060

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "profiles"
)

type Profile struct {
	Moniker  string         `json:"moniker" yaml:"moniker"`
	Name     *string        `json:"name,omitempty" yaml:"name,omitempty"`
	Surname  *string        `json:"surname,omitempty" yaml:"surname,omitempty"`
	Bio      *string        `json:"bio,omitempty" yaml:"bio,omitempty"`
	Pictures *Pictures      `json:"pictures,omitempty" yaml:"pictures,omitempty"`
	Creator  sdk.AccAddress `json:"creator" yaml:"creator"`
}

type Pictures struct {
	Profile *string `json:"profile,omitempty" yaml:"profile,omitempty"`
	Cover   *string `json:"cover,omitempty" yaml:"cover,omitempty"`
}
