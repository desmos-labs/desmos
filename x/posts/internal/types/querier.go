package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// QueryPostsParams Params for query 'custom/posts/posts'
type QueryPostsParams struct {
	Page  int
	Limit int

	ParentID       *PostID
	CreationTime   sdk.Int
	AllowsComments *bool
	Subspace       string
	Creator        sdk.AccAddress
}

func DefaultQueryPostsParams(page, limit int) QueryPostsParams {
	return QueryPostsParams{
		Page:  page,
		Limit: limit,

		ParentID:       nil,
		CreationTime:   sdk.NewInt(-1),
		AllowsComments: nil,
		Subspace:       "",
		Creator:        nil,
	}
}

// NewQueryPostsParams creates a new instance of QueryPostsParams
func NewQueryPostsParams(page, limit int,
	parentID *PostID, creationTime sdk.Int, allowsComments bool, subspace string, owner sdk.AccAddress,
) QueryPostsParams {
	return QueryPostsParams{
		Page:  page,
		Limit: limit,

		ParentID:       parentID,
		CreationTime:   creationTime,
		AllowsComments: &allowsComments,
		Subspace:       subspace,
		Creator:        owner,
	}
}
