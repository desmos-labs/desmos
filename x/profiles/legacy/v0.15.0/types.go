package v0150

import (
	v0130profiles "github.com/desmos-labs/desmos/x/profiles/legacy/v0.13.0"
	v0150profiles "github.com/desmos-labs/desmos/x/profiles/types"
)

// GenesisState contains the data of a v0.15.0 genesis state for the profiles module
type GenesisState = v0150profiles.GenesisState

// ----------------------------------------------------------------------------------------------------------------

type Profile = v0150profiles.Profile

func newProfile(profile v0130profiles.Profile) Profile {
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

type Pictures = v0150profiles.Pictures

// ----------------------------------------------------------------------------------------------------------------

type DTagTransferRequest = v0150profiles.DTagTransferRequest

// ----------------------------------------------------------------------------------------------------------------

type Params = v0150profiles.Params
type MonikerParams = v0150profiles.MonikerParams
type DTagParams = v0150profiles.DTagParams
