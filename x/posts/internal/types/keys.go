package types

const (
	ModuleName = "posts"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	PostStorePrefix    = "post:"
	LastPostIDStoreKey = "last_post_id"
	LikesStorePrefix   = "likes:"
	LastLikeIDStoreKey = "last_like_id"

	ActionCreatePost = "create_post"
	ActionEditPost   = "edit_post"
	ActionLikePost   = "like_post"
	ActionUnlikePost = "unlike_post"
)
