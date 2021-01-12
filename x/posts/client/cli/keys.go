package cli

import (
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/desmos-labs/desmos/x/posts/types"
)

// Posts flags
const (
	flagNumLimit = "limit"
	flagPage     = "page"

	flagSortBy   = "sort-by"
	flagSorOrder = "sort-order"

	FlagParentID       = "parent-id"
	FlagAttachment     = "attachment"
	FlagPollDetails    = "poll-details"
	FlagPollAnswer     = "poll-answer"
	FlagCreationTime   = "creation-time"
	FlagAllowsComments = "allows-comments"
	FlagSubspace       = "subspace"
	FlagHashtag        = "hashtag"
	FlagCreator        = "creator"

	keyEndDate           = "end-date"
	keyMultipleAnswers   = "multiple-answers"
	keyAllowsAnswerEdits = "allows-answer-edits"
	keyQuestion          = "question"
)

func DefaultQueryPostsRequest(page, limit uint64) types.QueryPostsRequest {
	return types.QueryPostsRequest{
		Pagination: &query.PageRequest{
			Key:        nil,
			Offset:     (page - 1) * limit,
			Limit:      limit,
			CountTotal: false,
		},
		SortBy:    types.PostSortByCreationDate,
		SortOrder: types.PostSortOrderAscending,

		ParentID:     "",
		CreationTime: nil,
		Subspace:     "",
		Creator:      "",
		Hashtags:     nil,
	}
}
