package types

// Subspaces module event types
const (
	EventTypeCreatedSubspace              = "create_subspace"
	EventTypeEditedSubspace               = "edit_subspace"
	EventTypeDeletedSubspace              = "delete_subspace"
	EventTypeCreatedSection               = "create_section"
	EventTypeEditedSection                = "edit_section"
	EventTypeMovedSection                 = "move_section"
	EventTypeDeletedSection               = "delete_section"
	EventTypeCreatedUserGroup             = "create_user_group"
	EventTypeEditedUserGroup              = "edit_user_group"
	EvenTypeMovedUserGroup                = "move_user_group"
	EventTypeSetUserGroupPermissions      = "set_user_group_permissions"
	EventTypeDeletedUserGroup             = "delete_user_group"
	EventTypeAddedUserToGroup             = "add_group_member"
	EventTypeRemovedUserFromGroup         = "remove_group_member"
	EventTypeSetUserPermissions           = "set_user_permissions"
	EventTypeGrantedTreasuryAuthorization = "grant_treasury_authorization"
	EventTypeRevokedTreasuryAuthorization = "revoke_treasury_authorization"
	EventTypeGrantedAllowance             = "grant_allowance"
	EventTypeRevokedAllowance             = "revoke_allowance"
	EventTypeUpdatedSubspaceFeeToken      = "update_subspace_fee_token"

	AttributeKeySubspaceID      = "subspace_id"
	AttributeKeySubspaceName    = "subspace_name"
	AttributeKeySubspaceCreator = "subspace_creator"
	AttributeKeyCreationTime    = "creation_date"
	AttributeKeySectionID       = "section_id"
	AttributeKeyUserGroupID     = "user_group_id"
	AttributeKeyPermissions     = "permissions"
	AttributeKeyUser            = "user"
	AttributeKeyGranter         = "granter"
	AttributeKeyGrantee         = "grantee"
	AttributeKeyUserGrantee     = "user_grantee"
	AttributeKeyGroupGrantee    = "group_grantee"
)
