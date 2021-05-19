package types

// Subspaces module event types
const (
	EventTypeCreateSubspace = "create_subspace"
	EventTypeEditSubspace   = "edit_subspace"
	EventTypeAddAdmin       = "admin_added"
	EventTypeRegisterUser   = "register_user"
	EventTypeUnregisterUser = "unregister_user"
	EventTypeBlockUser      = "block_user"
	EventTypeUnblockUser    = "unblock_user"

	// Subspaces attributes
	AttributeKeySubspaceID           = "subspace_id"
	AttributeKeySubspaceName         = "subspace_name"
	AttributeKeySubspaceCreator      = "subspace_creator"
	AttributeKeySubspaceNewAdmin     = "new_admin"
	AttributeKeySubspaceRemovedAdmin = "removed_admin"
	AttributeKeyRegisteredUser       = "registered_user"
	AttributeKeyUnregisteredUser     = "unregistered_user"
	AttributeKeyBlockedUser          = "blocked_user"
	AttributeKeyUnblockedUser        = "unblocked_user"
	AttributeKeyNewOwner             = "new_owner"
)
