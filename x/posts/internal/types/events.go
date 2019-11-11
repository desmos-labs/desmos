package types

// Magpie module event types
const (
	EventTypeCreatePost = "create_post"
	EventTypeEditPost   = "edit_post"
	EventTypeLikePost   = "like_post"
	EventTypeUnlikePost = "unlike_post"

	// Post attributes
	AttributeKeyPostID       = "post_id"
	AttributeKeyPostParentID = "post_parent_id"
	AttributeKeyPostOwner    = "post_owner"
	AttributeKeyPostEditTime = "post_edit_time"

	// Like attributes
	AttributeKeyLikeOwner = "liker"

	// Generic attributes
	AttributeKeyCreationTime  = "creation_time"
	AttributeKeyNamespace     = "namespace"
	AttributeKeyExternalOwner = "external_owner"
)
