package v0150

import (
	"time"

	v080 "github.com/desmos-labs/desmos/x/posts/legacy/v0.8.0"
	v0130 "github.com/desmos-labs/desmos/x/profiles/legacy/v0.13.0"
)

// GenesisState contains the data of a v0.15.0 genesis state for the profiles module
type GenesisState struct {
	Profiles             []Profile             `json:"profiles" yaml:"profiles"`
	DtagTransferRequests []DTagTransferRequest `json:"dtag_transfer_requests" yaml:"dtag_transfer_requests"`
	Params               v080.Params           `json:"params" yaml:"params"`
}

type Profile struct {
	Dtag         string    `json:"dtag,omitempty" yaml:"dtag"`
	Moniker      string    `json:"moniker,omitempty" yaml:"moniker,omitempty"`
	Bio          string    `json:"bio,omitempty" yaml:"bio,omitempty"`
	Pictures     Pictures  `json:"pictures" yaml:"pictures"`
	Creator      string    `json:"creator,omitempty" yaml:"creator,omitempty"`
	CreationDate time.Time `json:"creation_date" yaml:"creation_date"`
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
	Profile string `json:"profile,omitempty" yaml:"profile,omitempty"`
	Cover   string `json:"cover,omitempty" yaml:"cover,omitempty"`
}

type DTagTransferRequest struct {
	DtagToTrade string `json:"dtag_to_trade"    yaml:"dtag_to_trade"`
	Sender      string `json:"sender,omitempty" yaml:"sender,omitempty"`
	Receiver    string `json:"receiver,omitempty" yaml:"receiver,omitempty"`
}
