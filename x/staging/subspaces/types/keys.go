package types

const (
	ModuleName = "subspaces"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionCreateSubspace = "create_subspace"
	ActionEditSubspace   = "edit_subspace"
	ActionAddAdmin       = "add_admin"
	ActionRemoveAdmin    = "remove_admin"
	ActionRegisterUser   = "register_user"
	ActionBlockUser      = "block_user"

	// Queries
	QuerierRoute  = ModuleName
	QuerySubspace = "subspace"
)

var (
	SubspaceStorePrefix   = []byte("subspace")
	AdminsStorePrefix     = []byte("admins")
	BlockedUsersPrefix    = []byte("blocked")
	RegisteredUsersPrefix = []byte("registered")
)

// SubspaceStoreKey turns an id to a key used to store a subspace into the subspaces store
func SubspaceStoreKey(id string) []byte {
	return append(SubspaceStorePrefix, []byte(id)...)
}

// AdminsStoreKey turn an id in to a key used to store admins into the admins store
func AdminsStoreKey(subspaceID string) []byte {
	return append(AdminsStorePrefix, []byte(subspaceID)...)
}

// BlockedUsersStoreKey turn an id to a key used to store users that are blocked inside a subspace with the given id
func BlockedUsersStoreKey(subspaceID string) []byte {
	return append(BlockedUsersPrefix, []byte(subspaceID)...)
}

// RegisteredUsersStoreKey turn an id to a key used to store users that are registered under the subspace with the given id
func RegisteredUsersStoreKey(subspaceID string) []byte {
	return append(RegisteredUsersPrefix, []byte(subspaceID)...)
}
