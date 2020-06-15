package types

import (
	"encoding/json"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Profile represents a generic account on Desmos, containing the information of a single user
type Profile struct {
	Moniker  string         `json:"moniker" yaml:"moniker"`
	Bio      *string        `json:"bio,omitempty" yaml:"bio,omitempty"`
	Pictures *Pictures      `json:"pictures,omitempty" yaml:"pictures,omitempty"`
	Creator  sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewProfile(creator sdk.AccAddress) Profile {
	return Profile{
		Creator: creator,
	}
}

//WithMoniker updates profile's moniker with the given one
func (profile Profile) WithMoniker(moniker string) Profile {
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
	return profile.Moniker == other.Moniker &&
		profile.Bio == other.Bio &&
		profile.Pictures.Equals(other.Pictures) &&
		profile.Creator.Equals(other.Creator)
}

// Validate check the validity of the Profile
func (profile Profile) Validate() error {
	if profile.Creator.Empty() {
		return fmt.Errorf("profile creator cannot be empty or blank")
	}

	if strings.TrimSpace(profile.Moniker) == "" {
		return fmt.Errorf("profile moniker cannot be empty or blank")
	}

	if profile.Pictures != nil {
		if err := profile.Pictures.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type Profiles []Profile
