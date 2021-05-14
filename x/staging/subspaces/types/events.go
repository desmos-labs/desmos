package types

// Subspaces module event types
const (
	EventTypeCreateSubspace    = "subspace_created"
	EventTypeAddAdmin          = "admin_added"
	EventTypeEnableUserPosts   = "allowed_user_posts"
	EventTypeBlockUserPosts    = "blocked_user_posts"
	EventTypeTransferOwnership = "transfer_subspace_ownership"

	// Subspaces attributes
	AttributeKeySubspaceId           = "subspace_id"
	AttributeKeySubspaceCreator      = "subspace_creator"
	AttributeKeySubspaceNewAdmin     = "new_admin"
	AttributeKeySubspaceRemovedAdmin = "removed_admin"
	AttributeKeyEnabledToPostUser    = "enabled_user_to_post"
	AttributeKeyDisabledToPostUser   = "disabled_user_to_post"
	AttributeKeyNewOwner             = "new_owner"
)
