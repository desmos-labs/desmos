package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// Account represents a generic account on Desmos, containing the information of a single user
type Account struct {
	Name             string         `json:"name,omitempty"`
	Surname          string         `json:"surname,omitempty"`
	Moniker          string         `json:"moniker,omitempty"`
	Bio              string         `json:"bio,omitempty"`
	Pictures         *Pictures      `json:"pictures,omitempty"`
	VerifiedServices []ServiceLink  `json:"verified_services,omitempty"` // List of all the trusted services linked to this profile
	ChainLinks       []ChainLink    `json:"chain_links,omitempty"`       // List of all the other chain accounts linked to this profile
	Creator          sdk.AccAddress `json:"creator,omitempty" `
}

// Equals allows to check whether the contents of acc are the same of other
func (acc Account) Equals(other Account) bool {

	if len(acc.VerifiedServices) != len(other.VerifiedServices) {
		return false
	}

	if len(acc.ChainLinks) != len(other.ChainLinks) {
		return false
	}

	for index, service := range acc.VerifiedServices {
		if !service.Equals(other.VerifiedServices[index]) {
			return false
		}
	}

	for index, chainLink := range acc.ChainLinks {
		if !chainLink.Equals(other.ChainLinks[index]) {
			return false
		}
	}

	return acc.Name == other.Name &&
		acc.Surname == other.Surname &&
		acc.Moniker == other.Surname &&
		acc.Bio == other.Bio &&
		acc.Pictures.Equals(other.Pictures) &&
		acc.Creator.Equals(other.Creator)
}

func (acc Account) Validate() error {
	return
}

type Accounts []Account
