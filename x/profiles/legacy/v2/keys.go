package v2

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
func UserRelationshipsSubspacePrefix(user, subspace string) []byte {
	return append(UserRelationshipsPrefix(user), []byte(subspace)...)
}

// RelationshipsStoreKey returns the store key used to store the relationships containing the given data
func RelationshipsStoreKey(user, subspace, recipient string) []byte {
	return append(UserRelationshipsSubspacePrefix(user, subspace), []byte(recipient)...)
}

// BlockerPrefix returns the store prefix used to store the blocks created by the given blocker
func BlockerPrefix(blocker string) []byte {
	return append(UsersBlocksStorePrefix, []byte(blocker)...)
}

// BlockerSubspacePrefix returns the store prefix used to store the blocks that the given blocker
// has created inside the specified subspace
func BlockerSubspacePrefix(blocker string, subspace string) []byte {
	return append(BlockerPrefix(blocker), []byte(subspace)...)
}

// UserBlockStoreKey returns the store key used to save the block made by the given blocker,
// inside the specified subspace and towards the given blocked user
func UserBlockStoreKey(blocker string, subspace string, blockedUser string) []byte {
	return append(BlockerSubspacePrefix(blocker, subspace), []byte(blockedUser)...)
}
