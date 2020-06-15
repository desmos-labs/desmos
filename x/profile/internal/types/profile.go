package types

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Profile represents a generic account on Desmos, containing the information of a single user
type Profile struct {
	DTag     string         `json:"dtag" yaml:"dtag"`
	Moniker  *string        `json:"moniker,omitempty" yaml:"moniker,omitempty"`
	Bio      *string        `json:"bio,omitempty" yaml:"bio,omitempty"`
	Pictures *Pictures      `json:"pictures,omitempty" yaml:"pictures,omitempty"`
	Creator  sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewProfile(dtag string, creator sdk.AccAddress) Profile {
	return Profile{DTag: dtag, Creator: creator}
}

// WithDTag updates profile's DTag with the given one
func (profile Profile) WithDTag(dtag string) Profile {
	profile.DTag = dtag
	return profile
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
	if profile.Pictures == nil {
		arePicturesEquals = other.Pictures == nil
	} else {
		arePicturesEquals = profile.Pictures.Equals(other.Pictures)
	}

	return profile.DTag == other.DTag &&
		profile.Bio == other.Bio &&
		arePicturesEquals &&
		profile.Creator.Equals(other.Creator)
}

// Validate check the validity of the Profile
func (profile Profile) Validate() error {
	if profile.Creator.Empty() {
		return fmt.Errorf("profile creator cannot be empty or blank")
	}

	if !DTagRegEx.MatchString(profile.DTag) {
		return fmt.Errorf("invalid profile dtag")
	}

	if profile.Moniker != nil && (len(*profile.Moniker) < MinMonikerLength || len(*profile.Moniker) > MaxMonikerLength) {
		return fmt.Errorf("invalid profile moniker. Length should be between %d and %d", MinMonikerLength, MaxMonikerLength)
	}

	if profile.Pictures != nil {
		if err := profile.Pictures.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type Profiles []Profile
