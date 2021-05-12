package types

// Subspaces module event types
const (
	EventTypeCreateSubspace  = "subspace_created"
	EventTypeAddAdmin        = "admin_added"
	EventTypeEnableUserPosts = "allowed_user_posts"
	EventTypeBlockUserPosts  = "blocked_user_posts"

	// Subspaces attributes
	AttributeKeySubspaceId           = "subspace_id"
	AttributeKeySubspaceCreator      = "subspace_creator"
	AttributeKeySubspaceNewAdmin     = "new_admin"
	AttributeKeySubspaceRemovedAdmin = "removed_admin"
	AttributeKeyAllowedUser          = "allowed_user"
	AttributeKeyBlockedUser          = "blocked_user"
)
