package v0150

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	v0130 "github.com/desmos-labs/desmos/x/profiles/legacy/v0.13.0"
)

// GenesisState contains the data of a v0.15.0 genesis state for the profiles module
type GenesisState struct {
	Profiles             []Profile             `json:"profiles"`
	DtagTransferRequests []DTagTransferRequest `json:"dtag_transfer_requests"`
	Params               Params                `json:"params"`
}

// ----------------------------------------------------------------------------------------------------------------

type Profile struct {
	Dtag         string    `json:"dtag,omitempty"`
	Moniker      string    `json:"moniker,omitempty"`
	Bio          string    `json:"bio,omitempty"`
	Pictures     Pictures  `json:"pictures"`
	Creator      string    `json:"creator,omitempty"`
	CreationDate time.Time `json:"creation_date"`
}

func newProfile(profile v0130.Profile) Profile {
	moniker := ""
	bio := ""
	var pictures Pictures

	if profile.Moniker != nil {
		moniker = *profile.Moniker
	}

	if profile.Bio != nil {
		bio = *profile.Bio
	}

	if profile.Pictures != nil {
		pictures = Pictures{
			Profile: "",
			Cover:   "",
		}
		if profile.Pictures.Profile != nil {
			pictures.Profile = *profile.Pictures.Profile
		}
		if profile.Pictures.Cover != nil {
			pictures.Cover = *profile.Pictures.Cover
		}
	}

	return Profile{
		Dtag:         profile.DTag,
		Moniker:      moniker,
		Bio:          bio,
		Pictures:     pictures,
		Creator:      profile.Creator.String(),
		CreationDate: profile.CreationDate,
	}
}

type Pictures struct {
	Profile string `json:"profile,omitempty" `
	Cover   string `json:"cover,omitempty" `
}

// ----------------------------------------------------------------------------------------------------------------

type DTagTransferRequest struct {
	DtagToTrade string `json:"dtag_to_trade"`
	Sender      string `json:"sender,omitempty"`
	Receiver    string `json:"receiver,omitempty"`
}

// ----------------------------------------------------------------------------------------------------------------

type Params struct {
	MonikerParams MonikerParams `json:"moniker_params"`
	DtagParams    DTagParams    `json:"dtag_params"`
	MaxBioLength  sdk.Int       `json:"max_bio_length"`
}

type MonikerParams struct {
	MinMonikerLength sdk.Int `json:"min_length"`
	MaxMonikerLength sdk.Int `json:"max_length"`
}

type DTagParams struct {
	RegEx         string  `json:"reg_ex"`
	MinDtagLength sdk.Int `json:"min_length"`
	MaxDtagLength sdk.Int `json:"max_length"`
}
