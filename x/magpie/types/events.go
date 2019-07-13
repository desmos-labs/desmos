package types

// Magpie module event types
const (
	EventTypeCreatePost = "create_post"
	EventTypeEditPost   = "edit_post"
	EventTypeLikePost   = "like_post"
	EventTypeUnlikePost = "unlike_post"

	AttributeKeyPostOnwer = "power_owner"
	AttributeKeyLiker     = "liker"
	AttributeKeyTime      = "time"
	AttributeKeyPostID    = "post_id"

	AttributeValueCategory = ModuleName
)
