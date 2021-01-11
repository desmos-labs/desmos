package v0130

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "profiles"
)

// GenesisState contains the data of a v0.13.0 genesis state for the profiles module
type GenesisState struct {
	Profiles             []Profile             `json:"profiles" yaml:"profiles"`
	Params               Params                `json:"params" yaml:"params"`
	DTagTransferRequests []DTagTransferRequest `json:"dtag_transfer_requests" yaml:"dtag_transfer_requests"`
}

// ----------------------------------------------------------------------------------------------------------------

type Profile struct {
	DTag         string         `json:"dtag" yaml:"dtag"`
	Moniker      *string        `json:"moniker,omitempty" yaml:"moniker,omitempty"`
	Bio          *string        `json:"bio,omitempty" yaml:"bio,omitempty"`
	Pictures     *Pictures      `json:"pictures,omitempty" yaml:"pictures,omitempty"`
	Creator      sdk.AccAddress `json:"creator" yaml:"creator"`
	CreationDate time.Time      `json:"creation_date" yaml:"creation_date"`
}

type Pictures struct {
	Profile *string `json:"profile,omitempty" yaml:"profile,omitempty"`
	Cover   *string `json:"cover,omitempty" yaml:"cover,omitempty"`
}

// ----------------------------------------------------------------------------------------------------------------

type DTagTransferRequest struct {
	DTagToTrade string         `json:"dtag_to_trade" yaml:"dtag_to_trade"`
	Receiver    sdk.AccAddress `json:"receiver" yaml:"receiver"`
	Sender      sdk.AccAddress `json:"sender" yaml:"sender"`
}

// ----------------------------------------------------------------------------------------------------------------

type Params struct {
	MonikerParams MonikerParams `json:"moniker_params" yaml:"moniker_params"`
	DtagParams    DtagParams    `json:"dtag_params" yaml:"dtag_params"`
	MaxBioLen     sdk.Int       `json:"max_bio_length" yaml:"max_bio_length"`
}

type MonikerParams struct {
	MinMonikerLen sdk.Int `json:"min_length" yaml:"min_length"`
	MaxMonikerLen sdk.Int `json:"max_length" yaml:"max_length"`
}

type DtagParams struct {
	RegEx      string  `json:"reg_ex" yaml:"reg_ex"`
	MinDtagLen sdk.Int `json:"min_length" yaml:"min_length"`
	MaxDtagLen sdk.Int `json:"max_length" yaml:"max_length"`
}
