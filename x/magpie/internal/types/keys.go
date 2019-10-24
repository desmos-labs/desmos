package types

const (
	ModuleName = "magpie"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	PostStorePrefix       = "post:"
	LastPostIdStoreKey    = "last_post_id"
	LikeStorePrefix       = "like:"
	LastLikeIdStoreKey    = "last_like_id"
	SessionStorePrefix    = "session:"
	LastSessionIdStoreKey = "last_session_id"
)
