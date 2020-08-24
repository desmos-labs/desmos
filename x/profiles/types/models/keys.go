package models

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "profiles"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionSaveProfile        = "save_profile"
	ActionDeleteProfile      = "delete_profile"
	ActionCreateRelationship = "create_relationship"
	ActionDeleteRelationship = "delete_relationship"

	//Queries
	QuerierRoute       = ModuleName
	QueryProfile       = "profile"
	QueryProfiles      = "all"
	QueryRelationships = "relationships"
	QueryParams        = "params"
)

var (
	ProfileStorePrefix       = []byte("profile")
	DtagStorePrefix          = []byte("dtag")
	RelationshipsStorePrefix = []byte("relationships")
)

// ProfileStoreKey turns an address to a key used to store a profile into the profiles store
func ProfileStoreKey(address sdk.AccAddress) []byte {
	return append(ProfileStorePrefix, address...)
}

// DtagStoreKey turns a dtag to a key used to store a dtag -> address couple
func DtagStoreKey(dtag string) []byte {
	return append(DtagStorePrefix, []byte(dtag)...)
}

// RelationshipsStoreKey turns a user address to a key used to store a Address -> []Address couple
func RelationshipsStoreKey(user sdk.AccAddress) []byte {
	return append(RelationshipsStorePrefix, []byte(user)...)
}
