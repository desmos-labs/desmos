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
	QueryUserBlocks        = "user_blocks"
)

var (
	RelationshipsStorePrefix = []byte("relationships")
	UsersBlocksStorePrefix   = []byte("users_blocks")
)

// RelationshipsStoreKey turns a user address to a key used to store a Address -> []Address couple
func RelationshipsStoreKey(user sdk.AccAddress) []byte {
	return append(RelationshipsStorePrefix, []byte(user)...)
}

// UsersBlocksStoreKey turns a user address to a key used to store a Address -> []Address couple
func UsersBlocksStoreKey(user sdk.AccAddress) []byte {
	return append(UsersBlocksStorePrefix, []byte(user)...)
}
