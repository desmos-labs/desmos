package types

// Subspaces module event types
const (
	EventTypeCreateSubspace          = "create_subspace"
	EventTypeEditSubspace            = "edit_subspace"
	EventTypeDeleteSubspace          = "delete_subspace"
	EventTypeCreateUserGroup         = "create_user_group"
	EventTypeEditUserGroup           = "edit_user_group"
	EventTypeSetUserGroupPermissions = "set_user_group_permissions"
	EventTypeDeleteUserGroup         = "delete_user_group"
	EventTypeAddUserToGroup          = "add_group_member"
	EventTypeRemoveUserFromGroup     = "remove_group_member"
	EventTypeSetUserPermissions      = "set_user_permissions"

	AttributeValueCategory      = ModuleName
	AttributeKeySubspaceID      = "subspace_id"
	AttributeKeySubspaceName    = "subspace_name"
	AttributeKeySubspaceCreator = "subspace_creator"
	AttributeKeyCreationTime    = "creation_date"
	AttributeKeyUserGroupID     = "user_group_id"
	AttributeKeyUser            = "user"
)
