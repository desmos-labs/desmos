package v0150

import (
	"time"

	v0130profiles "github.com/desmos-labs/desmos/x/profiles/legacy/v0.13.0"
	v0150profiles "github.com/desmos-labs/desmos/x/profiles/types"
)

// GenesisState contains the data of a v0.15.0 genesis state for the profiles module
type GenesisState struct {
	Profiles             []Profile             `protobuf:"bytes,1,rep,name=profiles,proto3" json:"profiles" yaml:"profiles"`
	DTagTransferRequests []DTagTransferRequest `protobuf:"bytes,2,rep,name=dtag_transfer_requests,json=dtagTransferRequests,proto3" json:"dtag_transfer_requests" yaml:"dtag_transfer_requests"`
	Params               Params                `protobuf:"bytes,3,opt,name=params,proto3" json:"params" yaml:"params"`
}

// ----------------------------------------------------------------------------------------------------------------

type Profile struct {
	DTag         string    `protobuf:"bytes,1,opt,name=dtag,proto3" json:"dtag,omitempty" yaml:"dtag"`
	Moniker      string    `protobuf:"bytes,2,opt,name=moniker,proto3" json:"moniker,omitempty" yaml:"moniker"`
	Bio          string    `protobuf:"bytes,3,opt,name=bio,proto3" json:"bio,omitempty" yaml:"bio"`
	Pictures     Pictures  `protobuf:"bytes,4,opt,name=pictures,proto3" json:"pictures" yaml:"pictures"`
	Creator      string    `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
	CreationDate time.Time `protobuf:"bytes,6,opt,name=creation_date,json=creationDate,proto3,stdtime" json:"creation_date" yaml:"creation_date"`
}

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
		DTag:         profile.DTag,
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
