package types

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/commons"
)

// Profile represents a generic account on Desmos, containing the information of a single user
type Profile struct {
	DTag         string         `json:"dtag" yaml:"dtag"`
	Moniker      *string        `json:"moniker,omitempty" yaml:"moniker,omitempty"`
	Bio          *string        `json:"bio,omitempty" yaml:"bio,omitempty"`
	Pictures     *Pictures      `json:"pictures,omitempty" yaml:"pictures,omitempty"`
	Creator      sdk.AccAddress `json:"creator" yaml:"creator"`
	CreationDate time.Time      `json:"creation_date" yaml:"creation_date"`
}

func NewProfile(dtag string, creator sdk.AccAddress, creationDate time.Time) Profile {
	return Profile{DTag: dtag, Creator: creator, CreationDate: creationDate}
}

// WithMoniker updates profile's moniker with the given one
func (profile Profile) WithMoniker(moniker *string) Profile {
	profile.Moniker = moniker
	return profile
}

// WithBio updates profile's bio with the given one
func (profile Profile) WithBio(bio *string) Profile {
	profile.Bio = bio
	return profile
}

// WithPicture updates profile's pictures with the given one
func (profile Profile) WithPictures(profilePic, coverPic *string) Profile {
	profile.Pictures = NewPictures(profilePic, coverPic)
	return profile
}

// String implements fmt.Stringer
func (profile Profile) String() string {
	bytes, err := json.Marshal(&profile)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// Equals allows to check whether the contents of acc are the same of other
func (profile Profile) Equals(other Profile) bool {
	var arePicturesEquals bool
	if profile.Pictures == nil || other.Pictures == nil {
		arePicturesEquals = profile.Pictures == other.Pictures
	} else {
		arePicturesEquals = profile.Pictures.Equals(other.Pictures)
	}

	return profile.DTag == other.DTag &&
		commons.StringPtrsEqual(profile.Bio, other.Bio) &&
		arePicturesEquals &&
		profile.Creator.Equals(other.Creator)
}

// Validate check the validity of the Profile
func (profile Profile) Validate() error {
	if profile.Creator.Empty() {
		return fmt.Errorf("profile creator cannot be empty or blank")
	}
	if strings.TrimSpace(profile.DTag) == "" {
		return fmt.Errorf("profile dtag cannot be empty or blank")
	}

	if profile.Pictures != nil {
		if err := profile.Pictures.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Profiles represents a slice of profile objects
type Profiles []Profile

// NewProfiles allows to easily create a Profiles object from a list of profiles
func NewProfiles(profiles ...Profile) Profiles {
	return profiles
}
