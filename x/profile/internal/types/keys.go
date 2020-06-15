package types

import (
	"regexp"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "profiles"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	MinNameSurnameLength = 2
	MaxNameSurnameLength = 500
	MinDtagLength        = 2
	MaxDtagLength        = 30
	MaxBioLength         = 1000

	ActionSaveProfile   = "save_profile"
	ActionDeleteProfile = "delete_profile"

	//Queries
	QuerierRoute  = ModuleName
	QueryProfile  = "profile"
	QueryProfiles = "all"
)

var (
	TxHashRegEx = regexp.MustCompile("^[a-fA-F0-9]{64}$")
	URIRegEx    = regexp.MustCompile(
		`^(?:http(s)?://)[\w.-]+(?:\.[\w.-]+)+[\w\-._~:/?#[\]@!$&'()*+,;=.]+$`)

	ProfileStorePrefix = []byte("profile")
	DtagStorePrefix    = []byte("dtag")
)

// ProfileStoreKey turns an address to a key used to store a profile into the profiles store
func ProfileStoreKey(address sdk.AccAddress) []byte {
	return append(ProfileStorePrefix, address...)
}

// DtagStoreKey turns a dtag to a key used to store a dtag -> address couple
func DtagStoreKey(dtag string) []byte {
	return append(DtagStorePrefix, []byte(dtag)...)
}
