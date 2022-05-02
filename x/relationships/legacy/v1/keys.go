package v1

// DONTCOVER

import (
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

var (
	RelationshipsStorePrefix = []byte("relationships")
	UsersBlocksStorePrefix   = []byte("users_blocks")
)

// SubspaceRelationshipsPrefix returns the prefix used to store all relationships for the given subspace
func SubspaceRelationshipsPrefix(subspaceID uint64) []byte {
	return append(RelationshipsStorePrefix, subspacestypes.GetSubspaceIDBytes(subspaceID)...)
}

// UserRelationshipsSubspacePrefix returns the prefix used to store all the relationships created by the user
// with the given address for the subspace having the given id
func UserRelationshipsSubspacePrefix(subspace uint64, user string) []byte {
	return append(SubspaceRelationshipsPrefix(subspace), []byte(user)...)
}

// RelationshipsStoreKey returns the store key used to store the relationships containing the given data
func RelationshipsStoreKey(user, counterparty string, subspace uint64) []byte {
	return append(UserRelationshipsSubspacePrefix(subspace, user), []byte(counterparty)...)
}

// SubspaceBlocksPrefix returns the store prefix used to store the blocks for the given subspace
func SubspaceBlocksPrefix(subspaceID uint64) []byte {
	return append(UsersBlocksStorePrefix, subspacestypes.GetSubspaceIDBytes(subspaceID)...)
}

// BlockerSubspacePrefix returns the store prefix used to store the blocks that the given blocker
// has created inside the specified subspace
func BlockerSubspacePrefix(subspaceID uint64, blocker string) []byte {
	return append(SubspaceBlocksPrefix(subspaceID), []byte(blocker)...)
}

// UserBlockStoreKey returns the store key used to save the block made by the given blocker,
// inside the specified subspace and towards the given blocked user
func UserBlockStoreKey(blocker, blockedUser string, subspace uint64) []byte {
	return append(BlockerSubspacePrefix(subspace, blocker), []byte(blockedUser)...)
}
