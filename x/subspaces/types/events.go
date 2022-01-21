package types

// Subspaces module event types
const (
	EventTypeCreateSubspace      = "create_subspace"
	EventTypeEditSubspace        = "edit_subspace"
	EventTypeCreateUserGroup     = "create_user_group"
	EventTypeDeleteUserGroup     = "delete_user_group"
	EventTypeAddUserToGroup      = "add_group_member"
	EventTypeRemoveUserFromGroup = "delete_group_member"
	EventTypeSetPermissions      = "set_permissions"

	AttributeValueCategory      = ModuleName
	AttributeKeySubspaceID      = "subspace_id"
	AttributeKeySubspaceName    = "subspace_name"
	AttributeKeySubspaceCreator = "subspace_creator"
	AttributeKeyCreationTime    = "creation_date"
	AttributeKeyUserGroupName   = "user_group_name"
	AttributeKeyUser            = "user"
	AttributeKeyTarget          = "target"
)
