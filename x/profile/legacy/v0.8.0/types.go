package v080

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "profiles"
)

// GenesisState contains the data of a v0.4.0 genesis state for the profile module
type GenesisState struct {
	Profiles []Profile `json:"profiles"`
}

// Profile is a struct for a Profile
type Profile struct {
	DTag         string         `json:"dtag" yaml:"dtag"`
	Moniker      *string        `json:"moniker,omitempty" yaml:"moniker,omitempty"`
	Bio          *string        `json:"bio,omitempty" yaml:"bio,omitempty"`
	Pictures     *Pictures      `json:"pictures,omitempty" yaml:"pictures,omitempty"`
	Creator      sdk.AccAddress `json:"creator" yaml:"creator"`
	CreationDate time.Time      `json:"creation_date" yaml:"creation_date"`
}

// Pictures is a struct for Profile Pictures
type Pictures struct {
	Profile *string `json:"profile,omitempty" yaml:"profile,omitempty"`
	Cover   *string `json:"cover,omitempty" yaml:"cover,omitempty"`
}
