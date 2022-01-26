package types

import (
	subspacestypes "github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

// DONTCOVER

const (
	ModuleName = "relationships"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionCreateRelationship = "create_relationship"
	ActionDeleteRelationship = "delete_relationship"
	ActionBlockUser          = "block_user"
	ActionUnblockUser        = "unblock_user"
)

var (
	RelationshipsStorePrefix = []byte("relationships")
	UsersBlocksStorePrefix   = []byte("users_blocks")
)

// UserRelationshipsPrefix returns the prefix used to store all relationships created
// by the user with the given address
func UserRelationshipsPrefix(user string) []byte {
	return append(RelationshipsStorePrefix, []byte(user)...)
}

// UserRelationshipsSubspacePrefix returns the prefix used to store all the relationships created by the user
// with the given address for the subspace having the given id
func UserRelationshipsSubspacePrefix(user string, subspace uint64) []byte {
	return append(UserRelationshipsPrefix(user), subspacestypes.GetSubspaceIDBytes(subspace)...)
}

// RelationshipsStoreKey returns the store key used to store the relationships containing the given data
func RelationshipsStoreKey(user, counterparty string, subspace uint64) []byte {
	return append(UserRelationshipsSubspacePrefix(user, subspace), []byte(counterparty)...)
}

// BlockerPrefix returns the store prefix used to store the blocks created by the given blocker
func BlockerPrefix(blocker string) []byte {
	return append(UsersBlocksStorePrefix, []byte(blocker)...)
}

// BlockerSubspacePrefix returns the store prefix used to store the blocks that the given blocker
// has created inside the specified subspace
func BlockerSubspacePrefix(blocker string, subspace uint64) []byte {
	return append(BlockerPrefix(blocker), subspacestypes.GetSubspaceIDBytes(subspace)...)
}

// UserBlockStoreKey returns the store key used to save the block made by the given blocker,
// inside the specified subspace and towards the given blocked user
func UserBlockStoreKey(blocker string, subspace uint64, blockedUser string) []byte {
	return append(BlockerSubspacePrefix(blocker, subspace), []byte(blockedUser)...)
}
