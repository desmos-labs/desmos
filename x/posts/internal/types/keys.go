package types

const (
	ModuleName = "posts"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	PostStorePrefix    = "post:"
	LastPostIDStoreKey = "last_post_id"
	LikeStorePrefix    = "like:"
	LastLikeIDStoreKey = "last_like_id"

	ActionCreatePost = "create_post"
	ActionEditPost   = "edit_post"
	ActionLikePost   = "like"
)
