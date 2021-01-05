package v0130

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	v080 "github.com/desmos-labs/desmos/x/posts/legacy/v0.8.0"
)

// GenesisState contains the data of a v0.13.0 genesis state for the profiles module
type GenesisState struct {
	Profiles             []Profile             `json:"profiles" yaml:"profiles"`
	Params               v080.Params           `json:"params" yaml:"params"`
	DTagTransferRequests []DTagTransferRequest `json:"dtag_transfer_requests" yaml:"dtag_transfer_requests"`
}

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

type DTagTransferRequest struct {
	DTagToTrade string         `json:"dtag_to_trade" yaml:"dtag_to_trade"`
	Receiver    sdk.AccAddress `json:"receiver" yaml:"receiver"`
	Sender      sdk.AccAddress `json:"sender" yaml:"sender"`
}
