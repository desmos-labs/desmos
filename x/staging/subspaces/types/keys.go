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
	ActionUnregisterUser = "unregister_user"
	ActionBlockUser      = "block_user"
	ActionUnblockUser    = "unblock_user"

	// Queries
	QuerierRoute   = ModuleName
	QuerySubspace  = "subspace"
	QuerySubspaces = "subspaces"
)

var (
	SubspaceStorePrefix = []byte("subspace")
)

// SubspaceStoreKey turns an id to a key used to store a subspace into the subspaces store
func SubspaceStoreKey(id string) []byte {
	return append(SubspaceStorePrefix, []byte(id)...)
}
