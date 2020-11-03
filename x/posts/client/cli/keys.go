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

	flagParentID       = "parent-id"
	flagAttachment     = "attachment"
	flagPollDetails    = "poll-details"
	flagPollAnswer     = "poll-answer"
	flagCreationTime   = "creation-time"
	flagAllowsComments = "allows-comments"
	flagSubspace       = "subspace"
	flagHashtag        = "hashtag"
	flagCreator        = "creator"

	keyEndDate           = "end-date"
	keyMultipleAnswers   = "multiple-answers"
	keyAllowsAnswerEdits = "allows-answer-edits"
	keyQuestion          = "question"
)

func DefaultQueryPostsParams(page, limit uint64) types.QueryPostsRequest {
	return types.QueryPostsRequest{
		Pagination: &query.PageRequest{
			Key:        nil,
			Offset:     page * limit,
			Limit:      limit,
			CountTotal: false,
		},
		SortBy:    types.PostSortByCreationDate,
		SortOrder: types.PostSortOrderAscending,

		ParentID:     nil,
		CreationTime: nil,
		Subspace:     "",
		Creator:      nil,
		Hashtags:     nil,
	}
}
