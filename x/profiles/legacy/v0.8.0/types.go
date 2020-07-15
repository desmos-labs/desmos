package v080

// DONTCOVER

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "profiles"
)

// GenesisState contains the data of a v0.8.0 genesis state for the profile module
type GenesisState struct {
	Profiles []Profile `json:"profiles"`
	Params   Params    `json:"params" yaml:"params"`
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

type Params struct {
	MonikerParams MonikerParams `json:"moniker_params" yaml:"moniker_params"`
	DtagParams    DtagParams    `json:"dtag_params" yaml:"dtag_params"`
	MaxBioLen     sdk.Int       `json:"max_bio_length" yaml:"max_bio_length"`
}

// MonikerParams defines the paramsModule around moniker len
type MonikerParams struct {
	MinMonikerLen sdk.Int `json:"min_length" yaml:"min_length"`
	MaxMonikerLen sdk.Int `json:"max_length" yaml:"max_length"`
}

// DtagParams defines the paramsModule around profiles' dtag
type DtagParams struct {
	RegEx      string  `json:"reg_ex" yaml:"reg_ex"`
	MinDtagLen sdk.Int `json:"min_length" yaml:"min_length"`
	MaxDtagLen sdk.Int `json:"max_length" yaml:"max_length"`
}
