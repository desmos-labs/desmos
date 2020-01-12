package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// QueryPostsParams Params for query 'custom/posts/posts'
type QueryPostsParams struct {
	Page  int
	Limit int

	ParentID       *PostID
	CreationTime   *time.Time
	AllowsComments *bool
	Subspace       string
	Creator        sdk.AccAddress
}

func DefaultQueryPostsParams(page, limit int) QueryPostsParams {
	return QueryPostsParams{
		Page:  page,
		Limit: limit,

		ParentID:       nil,
		CreationTime:   nil,
		AllowsComments: nil,
		Subspace:       "",
		Creator:        nil,
	}
}

// NewQueryPostsParams creates a new instance of QueryPostsParams
func NewQueryPostsParams(page, limit int,
	parentID *PostID, creationTime time.Time, allowsComments bool, subspace string, owner sdk.AccAddress,
) QueryPostsParams {
	return QueryPostsParams{
		Page:  page,
		Limit: limit,

		ParentID:       parentID,
		CreationTime:   &creationTime,
		AllowsComments: &allowsComments,
		Subspace:       subspace,
		Creator:        owner,
	}
}
