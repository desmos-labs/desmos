package models

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	ModuleName = "relationships"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionCreateRelationship = "create_relationship"
	ActionDeleteRelationship = "delete_relationship"

	// Queries
	QuerierRoute           = ModuleName
	QueryUserRelationships = "user_relationships"
	QueryRelationships     = "relationships"
)

var (
	RelationshipsStorePrefix = []byte("relationships")
)

// RelationshipsStoreKey turns a user address to a key used to store a Address -> []Address couple
func RelationshipsStoreKey(user sdk.AccAddress) []byte {
	return append(RelationshipsStorePrefix, []byte(user)...)
}
