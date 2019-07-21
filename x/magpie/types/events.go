package types

// Magpie module event types
const (
	EventTypeCreatePost    = "create_post"
	EventTypeEditPost      = "edit_post"
	EventTypeLikePost      = "like"
	EventTypeUnlikePost    = "unlike"
	EventTypeCreateSession = "create_session"

	AttributeKeyPostOnwer     = "power_owner"
	AttributeKeyLiker         = "liker"
	AttributeKeyCreated       = "created"
	AttributeKeyModified      = "modified"
	AttributeKeyPostID        = "post_id"
	AttributeKeyLikeID        = "like_id"
	AttributeKeySessionID     = "session_id"
	AttributeKeyNamespace     = "namespace"
	AttributeKeyExternalOwner = "external_owner"
	AttributeKeyExpiry        = "expiry"

	AttributeValueCategory = ModuleName
)
