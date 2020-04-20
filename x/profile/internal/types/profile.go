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
	Name     *string        `json:"name,omitempty" yaml:"name,omitempty"`
	Surname  *string        `json:"surname,omitempty" yaml:"surname,omitempty"`
	Bio      *string        `json:"bio,omitempty" yaml:"bio,omitempty"`
	Pictures *Pictures      `json:"pictures,omitempty" yaml:"pictures,omitempty"`
	Creator  sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewProfile(moniker string, creator sdk.AccAddress) Profile {
	return Profile{
		Moniker: moniker,
		Creator: creator,
	}
}

// WithSurname updates profile's name with the given one
func (profile Profile) WithName(name *string) Profile {
	profile.Name = name
	return profile
}

// WithSurname updates profile's surname with the given one
func (profile Profile) WithSurname(surname *string) Profile {
	profile.Surname = surname
	return profile
}

// WithBio updates profile's bio with the given one
func (profile Profile) WithBio(bio *string) Profile {
	profile.Bio = bio
	return profile
}

// WithPicture updates profile's pictures with the given one
func (profile Profile) WithPictures(pictures *Pictures) Profile {
	profile.Pictures = pictures
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
	return profile.Name == other.Name &&
		profile.Surname == other.Surname &&
		profile.Moniker == other.Moniker &&
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
