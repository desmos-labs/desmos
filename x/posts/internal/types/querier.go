package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// QueryPostsParams Params for query 'custom/posts/posts'
type QueryPostsParams struct {
	Page         int
	Limit        int
	Creator      sdk.AccAddress
	ParentID     *PostID
	CreationTime sdk.Int
}

// NewQueryPostsParams creates a new instance of QueryPostsParams
func NewQueryPostsParams(page, limit int, parentID *PostID, creationTime sdk.Int, creator sdk.AccAddress) QueryPostsParams {
	return QueryPostsParams{
		Page:         page,
		Limit:        limit,
		Creator:      creator,
		ParentID:     parentID,
		CreationTime: creationTime,
	}
}
