package v040

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	ModuleName = "profile"
)

// GenesisState contains the data of a v0.4.0 genesis state for the profile module
type GenesisState struct {
	Profiles []Profile `json:"profiles"`
}

// Profile is a struct for a Profile
type Profile struct {
	Moniker  string         `json:"moniker"`
	Name     *string        `json:"name,omitempty"`
	Surname  *string        `json:"surname,omitempty"`
	Bio      *string        `json:"bio,omitempty"`
	Pictures *Pictures      `json:"pictures,omitempty"`
	Creator  sdk.AccAddress `json:"creator,omitempty" `
}

// Pictures is a struct for Profile Pictures
type Pictures struct {
	Profile *string `json:"profile,omitempty"`
	Cover   *string `json:"cover,omitempty"`
}
