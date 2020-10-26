package common

import (
	"regexp"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

const (
	ModuleName = "posts"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionCreatePost         = "create_post"
	ActionEditPost           = "edit_post"
	ActionAnswerPoll         = "answer_poll"
	ActionAddPostReaction    = "add_post_reaction"
	ActionRemovePostReaction = "remove_post_reaction"
	ActionRegisterReaction   = "register_reaction"

	// Queries
	QuerierRoute             = ModuleName
	QueryPost                = "post"
	QueryPosts               = "posts"
	QueryPollAnswers         = "poll-answers"
	QueryRegisteredReactions = "registered-reactions"
	QueryParams              = "params"

	// Sorting
	PostSortByCreationDate  = "created"
	PostSortByID            = "id"
	PostSortOrderAscending  = "ascending"
	PostSortOrderDescending = "descending"
)

var (
	postIDRegEx    = regexp.MustCompile(`^[a-fA-F0-9]{64}$`)
	shortCodeRegEx = regexp.MustCompile(`:[a-z0-9+-]([a-z0-9\d_-])*:`)

	ModuleAddress = authtypes.NewModuleAddress(ModuleName)

	PostStorePrefix          = []byte("post")
	PostIndexedIDStorePrefix = []byte("p_index")
	PostTotalNumberPrefix    = []byte("number_of_posts")
	PostCommentsStorePrefix  = []byte("comments")
	PostReactionsStorePrefix = []byte("p_reactions")
	ReactionsStorePrefix     = []byte("reactions")
	PollAnswersStorePrefix   = []byte("poll_answers")
)

// IsValidPostID tells whether the given value represents a valid post id or not
func IsValidPostID(value string) bool {
	return postIDRegEx.MatchString(value)
}

// IsValidReactionCode tells whether the given value is a valid emoji shortcode or not
func IsValidReactionCode(value string) bool {
	return shortCodeRegEx.MatchString(value)
}
