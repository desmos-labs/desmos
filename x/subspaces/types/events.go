package types

// Subspaces module event types
const (
	EventTypeCreateSubspace         = "create_subspace"
	EventTypeEditSubspace           = "edit_subspace"
	EventTypeAddAdmin               = "admin_added"
	EventTypeRemoveAdmin            = "remove_admin"
	EventTypeRegisterUser           = "register_user"
	EventTypeUnregisterUser         = "unregister_user"
	EventTypeBanUser                = "ban_user"
	EventTypeUnbanUser              = "unban_user"
	EventTypeSaveTokenomicsContract = "save_tokenomics_contract"

	AttributeKeySubspaceID           = "subspace_id"
	AttributeKeySubspaceName         = "subspace_name"
	AttributeKeySubspaceCreator      = "subspace_creator"
	AttributeKeyCreationTime         = "creation_date"
	AttributeKeySubspaceNewAdmin     = "new_admin"
	AttributeKeySubspaceRemovedAdmin = "removed_admin"
	AttributeKeyRegisteredUser       = "registered_user"
	AttributeKeyUnregisteredUser     = "unregistered_user"
	AttributeKeyBanUser              = "banned_user"
	AttributeKeyUnbannedUser         = "unbanned_user"
	AttributeKeyNewOwner             = "new_owner"
	AttributeKeyContractAddress      = "contract_address"
)
