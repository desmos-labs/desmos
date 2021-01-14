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
	Profiles             []Profile             `json:"profiles"`
	Params               Params                `json:"params"`
	DTagTransferRequests []DTagTransferRequest `json:"dtag_transfer_requests"`
}

// ----------------------------------------------------------------------------------------------------------------

type Profile struct {
	DTag         string         `json:"dtag"`
	Moniker      *string        `json:"moniker,omitempty"`
	Bio          *string        `json:"bio,omitempty"`
	Pictures     *Pictures      `json:"pictures,omitempty"`
	Creator      sdk.AccAddress `json:"creator"`
	CreationDate time.Time      `json:"creation_date"`
}

type Pictures struct {
	Profile *string `json:"profile,omitempty"`
	Cover   *string `json:"cover,omitempty"`
}

// ----------------------------------------------------------------------------------------------------------------

type DTagTransferRequest struct {
	DTagToTrade string         `json:"dtag_to_trade"`
	Receiver    sdk.AccAddress `json:"current_owner"`
	Sender      sdk.AccAddress `json:"receiving_user"`
}

// ----------------------------------------------------------------------------------------------------------------

type Params struct {
	MonikerParams MonikerParams `json:"moniker_params"`
	DtagParams    DtagParams    `json:"dtag_params"`
	MaxBioLen     sdk.Int       `json:"max_bio_length"`
}

type MonikerParams struct {
	MinMonikerLen sdk.Int `json:"min_length"`
	MaxMonikerLen sdk.Int `json:"max_length"`
}

type DtagParams struct {
	RegEx      string  `json:"reg_ex"`
	MinDtagLen sdk.Int `json:"min_length"`
	MaxDtagLen sdk.Int `json:"max_length"`
}
