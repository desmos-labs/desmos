package types

// Magpie module event types
const (
	EventTypeCreatePost = "create_post"
	EventTypeEditPost   = "edit_post"
	EventTypeLikePost   = "like_post"
	EventTypeUnlikePost = "unlike_post"

	AttributeKeyPostOwner     = "post_owner"
	AttributeKeyLiker         = "liker"
	AttributeKeyCreated       = "created"
	AttributeKeyModified      = "modified"
	AttributeKeyPostID        = "post_id"
	AttributeKeyLikeID        = "like_id"
	AttributeKeyNamespace     = "namespace"
	AttributeKeyExternalOwner = "external_owner"

	AttributeValueCategory = ModuleName
)
