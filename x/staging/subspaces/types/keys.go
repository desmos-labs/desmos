package types

const (
	ModuleName = "subspaces"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionCreateSubspace = "create_subspace"
	ActionAddAdmin       = "add_admin"
	ActionAllowUserPosts = "allow_user_posts"
	ActionBlockUserPosts = "block_user_posts"

	// Queries
	QuerierRoute = ModuleName
)
