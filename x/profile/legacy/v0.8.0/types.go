package v0_8_0

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Profile struct {
	DTag     string         `json:"dtag" yaml:"dtag"`
	Moniker  *string        `json:"moniker,omitempty" yaml:"moniker,omitempty"`
	Bio      *string        `json:"bio,omitempty" yaml:"bio,omitempty"`
	Pictures *Pictures      `json:"pictures,omitempty" yaml:"pictures,omitempty"`
	Creator  sdk.AccAddress `json:"creator" yaml:"creator"`
}

type Pictures struct {
	Profile *string `json:"profile,omitempty" yaml:"profile,omitempty"`
	Cover   *string `json:"cover,omitempty" yaml:"cover,omitempty"`
}
