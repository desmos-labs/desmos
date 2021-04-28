package types

const (
	ModuleName = "subspaces"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionCreateSubspace = "create_subspace"
	ActionAddAdmin       = "add_admin"
	ActionRemoveAdmin    = "remove_admin"
	ActionAllowUserPosts = "allow_user_posts"
	ActionBlockUserPosts = "block_user_posts"

	// Queries
	QuerierRoute = ModuleName
)

var (
	SubspaceStorePrefix = []byte("susbpace")
	AdminsStorePrefix   = []byte("admins")
)

// SubspaceStoreKey turns an id to a key used to store a subspace into the subspaces store
func SubspaceStoreKey(id string) []byte {
	return append(SubspaceStorePrefix, []byte(id)...)
}

// AdminsStoreKey turn an in to a key used to store admins into the admins store
func AdminsStoreKey(subspaceId string) []byte {
	return append(AdminsStorePrefix, []byte(subspaceId)...)
}
