package types

const (
	ModuleName = "posts"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	PostStorePrefix                 = "post:"
	LastPostIDStoreKey              = "last_post_id"
	PostCommentsStorePrefix         = "comments:"
	PostReactionsStorePrefix        = "reactions:"
	MaxOptionalDataFieldsNumber     = 10
	MaxOptionalDataFieldValueLength = 200

	ActionCreatePost         = "create_post"
	ActionEditPost           = "edit_post"
	ActionAddPostReaction    = "add_post_reaction"
	ActionRemovePostReaction = "remove_post_reaction"

	// Queries
	QuerierRoute = ModuleName
	QueryPost    = "post"
	QueryPosts   = "posts"
	QueryLike    = "like"
)
