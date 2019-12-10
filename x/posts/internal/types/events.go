package types

// Magpie module event types
const (
	EventTypePostCreated = "post_created"
	EventTypePostEdited  = "post_edited"
	EventTypePostLiked   = "post_liked"
	EventTypePostUnliked = "post_unliked"

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
