package types

// Magpie module event types
const (
	EventTypeCreatePost = "create_post"
	EventTypeEditPost   = "edit_post"
	EventTypeLikePost   = "like"
	EventTypeUnlikePost = "unlike"

	AttributeKeyPostOnwer = "power_owner"
	AttributeKeyLiker     = "liker"
	AttributeKeyTime      = "time"
	AttributeKeyPostID    = "post_id"
	AttributeKeyLikeID    = "like_id"

	AttributeValueCategory = ModuleName
)
