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
	MinMonikerLength     = 2
	MaxMonikerLength     = 30
	MaxBioLength         = 1000

	ActionSaveProfile   = "save_profile"
	ActionDeleteProfile = "delete_profile"

	//Queries
	QuerierRoute  = ModuleName
	QueryProfile  = "profile"
	QueryProfiles = "all"
	QueryParams   = "params"
)

var (
	TxHashRegEx = regexp.MustCompile("^[a-fA-F0-9]{64}$")
	URIRegEx    = regexp.MustCompile(
		`^(?:http(s)?://)[\w.-]+(?:\.[\w.-]+)+[\w\-._~:/?#[\]@!$&'()*+,;=.]+$`)

	ProfileStorePrefix = []byte("profile")
	MonikerStorePrefix = []byte("moniker")
)

// ProfileStoreKey turns an address to a key used to store a profile into the profiles store
func ProfileStoreKey(address sdk.AccAddress) []byte {
	return append(ProfileStorePrefix, address...)
}

// MonikerStoreKey turns a moniker to a key used to store a moniker -> address couple
func MonikerStoreKey(moniker string) []byte {
	return append(MonikerStorePrefix, []byte(moniker)...)
}
