package types

const (
	ModuleName = "subspaces"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionCreateSubspace    = "create_subspace"
	ActionAddAdmin          = "add_admin"
	ActionRemoveAdmin       = "remove_admin"
	ActionEnableUserPosts   = "enable_user_posts"
	ActionDisableUserPosts  = "disable_user_posts"
	ActionTransferOwnership = "transfer_ownership"

	// Queries
	QuerierRoute  = ModuleName
	QuerySubspace = "subspace"
)

var (
	SubspaceStorePrefix     = []byte("subspace")
	AdminsStorePrefix       = []byte("admins")
	BlockedUsersPostsPrefix = []byte("blocked")
)

// SubspaceStoreKey turns an id to a key used to store a subspace into the subspaces store
func SubspaceStoreKey(id string) []byte {
	return append(SubspaceStorePrefix, []byte(id)...)
}

// AdminsStoreKey turn an id in to a key used to store admins into the admins store
func AdminsStoreKey(subspaceId string) []byte {
	return append(AdminsStorePrefix, []byte(subspaceId)...)
}

// BlockedToPostUsersKey turn an id to a key used to store users that are not allowed to post inside a subspace
func BlockedToPostUsersKey(subspaceId string) []byte {
	return append(BlockedUsersPostsPrefix, []byte(subspaceId)...)
}
