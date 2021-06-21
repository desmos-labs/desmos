package cli

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	types2 "github.com/desmos-labs/desmos/x/posts/types"
)

// DONTCOVER

// Posts flags
const (
	flagNumLimit = "limit"
	flagPage     = "page"

	flagSortBy   = "sort-by"
	flagSorOrder = "sort-order"

	FlagParentID      = "parent-id"
	FlagAttachment    = "attachment"
	FlagPollDetails   = "poll-details"
	FlagPollAnswer    = "poll-answer"
	FlagCreationTime  = "creation-time"
	FlagCommentsState = "comments-state"
	FlagSubspace      = "subspace"
	FlagHashtag       = "hashtag"
	FlagCreator       = "creator"

	keyEndDate           = "end-date"
	keyMultipleAnswers   = "multiple-answers"
	keyAllowsAnswerEdits = "allows-answer-edits"
	keyQuestion          = "question"
)

func DefaultQueryPostsRequest(page, limit uint64) types2.QueryPostsRequest {
	return types2.QueryPostsRequest{
		Pagination: &query.PageRequest{
			Key:        nil,
			Offset:     (page - 1) * limit,
			Limit:      limit,
			CountTotal: false,
		},
		SortBy:    types2.PostSortByCreationDate,
		SortOrder: types2.PostSortOrderAscending,

		ParentId:     "",
		CreationTime: nil,
		Subspace:     "",
		Creator:      "",
		Hashtags:     nil,
	}
}
