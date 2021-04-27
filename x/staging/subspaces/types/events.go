package types

// Subspaces module event types
const (
	EventTypeCreateSubspace = "subspace_created"
	EventTypeAddAdmin       = "admin_added"
	EventTypeAllowUserPosts = "allowed_user_posts"
	EventTypeBlockUserPosts = "blocked_user_posts"

	// Subspaces attributes
	AttributeKeySubspaceId       = "subspace_id"
	AttributeKeySubspaceCreator  = "subspace_creator"
	AttributeKeySubspaceNewAdmin = "new_admin"
	AttributeKeyAllowedUser      = "allowed_user"
	AttributeKeyBlockUser        = "blocked_user"
)
