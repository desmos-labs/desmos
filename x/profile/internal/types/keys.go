package types

import "regexp"

const (
	ModuleName = "profile"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	MaxNameSurnameLength = 500
	MaxMonikerLength     = 20
	MaxBioLength         = 1000

	ActionCreateProfile = "create_profile"
	ActionEditProfile   = "edit_profile"
	ActionDeleteProfile = "delete_profile"

	//Queries
	QuerierRoute  = ModuleName
	QueryProfile  = "profile"
	QueryProfiles = "profiles"
)

var (
	TxHashRegEx = regexp.MustCompile("^[a-fA-F0-9]{64}$")

	ProfileStorePrefix = []byte("profile")
)

// ProfileStoreKey turns a moniker to a key used to store a profile into the profiles store
func ProfileStoreKey(moniker string) []byte {
	return append(ProfileStorePrefix, []byte(moniker)...)
}
