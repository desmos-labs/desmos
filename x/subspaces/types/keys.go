package types

// DONTCOVER

const (
	ModuleName = "subspaces"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionCreateSubspace   = "create_subspace"
	ActionEditSubspace     = "edit_subspace"
	ActionAddAdmin         = "add_admin"
	ActionRemoveAdmin      = "remove_admin"
	ActionRegisterUser     = "register_user"
	ActionUnregisterUser   = "unregister_user"
	ActionBlockUser        = "block_user"
	ActionUnblockUser      = "unblock_user"
	ActionInsertTokenomics = "insert_tokenomics"

	QuerierRoute = ModuleName

	DoNotModify = "[do-not-modify]"
)

var (
	adminPrefix          = []byte("prefix")
	registeredUserPrefix = []byte("user")
	bannedUserPrefix     = []byte("banned")
	SubspaceStorePrefix  = []byte("subspace")
	TokenomicsPairPrefix = []byte("tokenomics")
)

// SubspaceStoreKey turns an id to a key used to store a subspace into the subspaces store
func SubspaceStoreKey(id string) []byte {
	return append(SubspaceStorePrefix, []byte(id)...)
}

// SubspaceAdminsPrefix returns the prefix used to store each admin
// of the subspace with the given id
func SubspaceAdminsPrefix(id string) []byte {
	return append([]byte(id), adminPrefix...)
}

// SubspaceAdminKey returns the key used to store the given admin as an admin
// of the subspace with the given id
func SubspaceAdminKey(id string, admin string) []byte {
	return append(SubspaceAdminsPrefix(id), []byte(admin)...)
}

// SubspaceRegisteredUsersPrefix returns the prefix used to store each registered user
// of the subspace with the given id
func SubspaceRegisteredUsersPrefix(id string) []byte {
	return append([]byte(id), registeredUserPrefix...)
}

// SubspaceRegisteredUserKey returns the key used to store the given user as a registered user
// of the subspace with the given id
func SubspaceRegisteredUserKey(id string, user string) []byte {
	return append(SubspaceRegisteredUsersPrefix(id), []byte(user)...)
}

// SubspaceBannedUsersPrefix returns the prefix used to store each banned user
// of the subspace with the given id
func SubspaceBannedUsersPrefix(id string) []byte {
	return append([]byte(id), bannedUserPrefix...)
}

// SubspaceBannedUserKey returns the key used to store the given user as a banned user
// of the subspace with the given id
func SubspaceBannedUserKey(id string, user string) []byte {
	return append(SubspaceBannedUsersPrefix(id), []byte(user)...)
}

// TokenomicsPairKey turns an id into a key used to store a tokenomics pair inside the store
func TokenomicsPairKey(id string) []byte {
	return append(TokenomicsPairPrefix, []byte(id)...)
}
