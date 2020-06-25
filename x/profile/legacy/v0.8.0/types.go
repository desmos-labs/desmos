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
	MonikerLengths MonikerLengths `json:"moniker_lengths" yaml:"moniker_lengths"`
	DtagLengths    DtagLengths    `json:"dtag_lengths" yaml:"dtag_lengths"`
	MaxBioLen      sdk.Int        `json:"max_bio_len" yaml:"max_bio_len"`
}

// MonikerLengths defines the paramsModule around moniker len
type MonikerLengths struct {
	MinMonikerLen sdk.Int `json:"min_moniker_len" yaml:"min_moniker_len"`
	MaxMonikerLen sdk.Int `json:"max_moniker_len" yaml:"max_moniker_len"`
}

// DtagLengths defines the paramsModule around profiles' dtag
type DtagLengths struct {
	MinDtagLen sdk.Int `json:"min_dtag_len" yaml:"min_dtag_len"`
	MaxDtagLen sdk.Int `json:"max_dtag_len" yaml:"max_dtag_len"`
}
