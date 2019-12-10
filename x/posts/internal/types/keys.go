package types

const (
	ModuleName = "posts"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	PostStorePrefix         = "post:"
	LastPostIDStoreKey      = "last_post_id"
	PostCommentsStorePrefix = "comments:"

	LikesStorePrefix = "likes:"

	ActionCreatePost = "create_post"
	ActionEditPost   = "edit_post"
	ActionLikePost   = "like_post"
	ActionUnlikePost = "unlike_post"
)
