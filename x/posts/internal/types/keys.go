package types

const (
	ModuleName = "posts"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	PostStorePrefix                 = "post:"
	MediaPostStorePrefix            = "media_post:"
	MediaProvidersStoreKey          = "media_providers"
	MediaMimeTypeStoreKey           = "media_mime_type"
	LastPostIDStoreKey              = "last_post_id"
	PostCommentsStorePrefix         = "comments:"
	PostReactionsStorePrefix        = "reactions:"
	MaxOptionalDataFieldsNumber     = 10
	MaxOptionalDataFieldValueLength = 200

	ActionCreatePost         = "create_post"
	ActionCreateMediaPost    = "create_media_post"
	ActionEditPost           = "edit_post"
	ActionAddPostReaction    = "add_post_reaction"
	ActionRemovePostReaction = "remove_post_reaction"

	// Queries
	QuerierRoute = ModuleName
	QueryPost    = "post"
	QueryPosts   = "posts"
)
