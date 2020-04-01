package types

import (
	"encoding/json"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Profile represents a generic account on Desmos, containing the information of a single user
type Profile struct {
	Name             string         `json:"name,omitempty"`
	Surname          string         `json:"surname,omitempty"`
	Moniker          string         `json:"moniker"`
	Bio              string         `json:"bio,omitempty"`
	Pictures         *Pictures      `json:"pictures,omitempty"`
	VerifiedServices []ServiceLink  `json:"verified_services,omitempty"` // List of all the trusted services linked to this profile
	ChainLinks       []ChainLink    `json:"chain_links,omitempty"`       // List of all the other chain accounts linked to this profile
	Creator          sdk.AccAddress `json:"creator,omitempty" `
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

	if len(profile.VerifiedServices) != len(other.VerifiedServices) {
		return false
	}

	if len(profile.ChainLinks) != len(other.ChainLinks) {
		return false
	}

	for index, service := range profile.VerifiedServices {
		if !service.Equals(other.VerifiedServices[index]) {
			return false
		}
	}

	for index, chainLink := range profile.ChainLinks {
		if !chainLink.Equals(other.ChainLinks[index]) {
			return false
		}
	}

	return profile.Name == other.Name &&
		profile.Surname == other.Surname &&
		profile.Moniker == other.Moniker &&
		profile.Bio == other.Bio &&
		profile.Pictures.Equals(other.Pictures) &&
		profile.Creator.Equals(other.Creator)
}

func (profile Profile) Validate() error {
	if profile.Creator.Empty() {
		return fmt.Errorf("profile creator cannot be empty or blank")
	}

	if len(strings.TrimSpace(profile.Moniker)) == 0 {
		return fmt.Errorf("profile moniker cannot be empty or blank")
	}

	if len(profile.VerifiedServices) != 0 {
		for _, verifiedService := range profile.VerifiedServices {
			if err := verifiedService.Validate(); err != nil {
				return err
			}
		}
	}

	if len(profile.ChainLinks) != 0 {
		for _, chainLink := range profile.ChainLinks {
			if err := chainLink.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}

type Profiles []Profile
