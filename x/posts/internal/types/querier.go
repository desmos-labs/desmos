package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// QueryPostsParams Params for query 'custom/posts/posts'
type QueryPostsParams struct {
	Page  int
	Limit int

	SortBy    string // Field that should determine the sorting
	SortOrder string // Either ascending or descending

	ParentID       PostID
	CreationTime   *time.Time
	AllowsComments *bool
	Subspace       string
	Creator        sdk.AccAddress
	Hashtags       []string
}

func DefaultQueryPostsParams(page, limit int) QueryPostsParams {
	return QueryPostsParams{
		Page:  page,
		Limit: limit,

		SortBy:    PostSortByCreationDate,
		SortOrder: PostSortOrderAscending,

		ParentID:       "",
		CreationTime:   nil,
		AllowsComments: nil,
		Subspace:       "",
		Creator:        nil,
		Hashtags:       nil,
	}
}
