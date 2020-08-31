package models

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	ModuleName = "relationships"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionCreateRelationship = "create_relationship"
	ActionDeleteRelationship = "delete_relationship"
	ActionBlockUser          = "block_user"
	ActionUnblockUser        = "unblock_user"

	// Queries
	QuerierRoute           = ModuleName
	QueryUserRelationships = "user_relationships"
	QueryRelationships     = "relationships"
	QueryBlockedUsers      = "blocked_users"
)

var (
	RelationshipsStorePrefix = []byte("relationships")
	BlockedUsersStorePrefix  = []byte("blocked_users_of")
)

// RelationshipsStoreKey turns a user address to a key used to store a Address -> []Address couple
func RelationshipsStoreKey(user sdk.AccAddress) []byte {
	return append(RelationshipsStorePrefix, []byte(user)...)
}

// BlockedUsersStoreKey turns a user address to a key used to store a Address -> []Address couple
func BlockedUsersStoreKey(user sdk.AccAddress) []byte {
	return append(BlockedUsersStorePrefix, []byte(user)...)
}
