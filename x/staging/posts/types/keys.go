package types

// DONTCOVER

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
	ActionReportPost         = "report_post"
	ActionAnswerPoll         = "answer_poll"
	ActionAddPostReaction    = "add_post_reaction"
	ActionRemovePostReaction = "remove_post_reaction"
	ActionRegisterReaction   = "register_reaction"

	// Queries
	QuerierRoute             = ModuleName
	QueryPost                = "post"
	QueryPosts               = "posts"
	QueryReports             = "reports"
	QueryPollAnswers         = "poll-answer"
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
	hashtagRegEx   = regexp.MustCompile(`[^\S]|^#([^\s#.,!)]+)$`)

	ModuleAddress = authtypes.NewModuleAddress(ModuleName)

	PostStorePrefix          = []byte("post")
	PostIndexedIDStorePrefix = []byte("p_index")
	PostTotalNumberPrefix    = []byte("number_of_posts")
	PostCommentsStorePrefix  = []byte("comments")
	PostReactionsStorePrefix = []byte("p_reactions")
	ReactionsStorePrefix     = []byte("reactions")
	PollAnswersStorePrefix   = []byte("poll_answers")
	ReportsStorePrefix       = []byte("reports")
)

// IsValidPostID tells whether the given value represents a valid post id or not
func IsValidPostID(value string) bool {
	return postIDRegEx.MatchString(value)
}

// IsValidReactionCode tells whether the given value is a valid emoji shortcode or not
func IsValidReactionCode(value string) bool {
	return shortCodeRegEx.MatchString(value)
}

// PostStoreKey turns an id to a key used to store a post into the posts store
func PostStoreKey(id string) []byte {
	return append(PostStorePrefix, []byte(id)...)
}

// PostIndexedIDStoreKey turns an id to a key used to store an incremental ID into the posts store
func PostIndexedIDStoreKey(id string) []byte {
	return append(PostIndexedIDStorePrefix, []byte(id)...)
}

// PostCommentsStoreKey turns an id to a key used to store a post's comments into the posts store
func PostCommentsStoreKey(id string) []byte {
	return append(PostCommentsStorePrefix, []byte(id)...)
}

// PostCommentsStoreKey turns an id to a key used to store a post's reactions into the posts store
func PostReactionsStoreKey(id string) []byte {
	return append(PostReactionsStorePrefix, []byte(id)...)
}

// RegisteredReactionsPrefix returns the prefix used to store all the reactions for the subspace having the given id
func RegisteredReactionsPrefix(subspace string) []byte {
	return append(ReactionsStorePrefix, []byte(subspace)...)
}

// RegisteredReactionsStoreKey returns the key used to store the registered reaction having the given short code for the given subspace
func RegisteredReactionsStoreKey(subspace, shortCode string) []byte {
	return append(RegisteredReactionsPrefix(subspace), []byte(shortCode)...)
}

// PollAnswersByIdPrefix returns the prefix used to store all the poll answers for the post having the given id
func PollAnswersByIdPrefix(id string) []byte {
	return append(PollAnswersStorePrefix, []byte(id)...)
}

// UserPollAnswersByIdPrefix returns the prefix used to store all the poll answers created by the user
// with given address for the post having the given id
func UserPollAnswersByIdPrefix(id, user string) []byte {
	return append(PollAnswersByIdPrefix(id), []byte(user)...)
}

// PollAnswersStoreKey returns the store key used to store the poll answer containing the given data
func PollAnswersStoreKey(id, user, answer string) []byte {
	return append(UserPollAnswersByIdPrefix(id, user), []byte(answer)...)
}

// ReportStoreKey turns an id into a key used to store a report inside the reports store
func ReportStoreKey(id string) []byte {
	return append(ReportsStorePrefix, []byte(id)...)
}
