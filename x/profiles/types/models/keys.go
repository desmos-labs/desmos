package models

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "profiles"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionSaveProfile                       = "save_profile"
	ActionDeleteProfile                     = "delete_profile"
	ActionCreateMonoDirectionalRelationship = "create_mono_directional_relationship"
	ActionRequestBiDirectionalRelationship  = "request_bi_directional_relationship"
	ActionAcceptBiDirectionalRelationship   = "accept_bi_directional_relationship"
	ActionDenyBiDirectionalRelationship     = "deny_bi_directional_relationship"
	ActionDeleteRelationship                = "delete_relationship"

	//Queries
	QuerierRoute       = ModuleName
	QueryProfile       = "profile"
	QueryProfiles      = "all"
	QueryRelationships = "relationships"
	QueryParams        = "params"
)

var (
	ProfileStorePrefix      = []byte("profile")
	DtagStorePrefix         = []byte("dtag")
	RelationshipsPrefix     = []byte("relationships")
	UserRelationshipsPrefix = []byte("user")
)

// ProfileStoreKey turns an address to a key used to store a profile into the profiles store
func ProfileStoreKey(address sdk.AccAddress) []byte {
	return append(ProfileStorePrefix, address...)
}

// DtagStoreKey turns a dtag to a key used to store a dtag -> address couple
func DtagStoreKey(dtag string) []byte {
	return append(DtagStorePrefix, []byte(dtag)...)
}

// RelationshipsStoreKey turns a relID to a key used to store a RelationshipID -> Relationship couple
func RelationshipsStoreKey(relID RelationshipID) []byte {
	return append(RelationshipsPrefix, []byte(relID)...)
}

func UserRelationshipsStoreKey(address sdk.AccAddress) []byte {
	return append(UserRelationshipsPrefix, address...)
}
