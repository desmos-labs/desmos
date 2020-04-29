package types

import (
	"regexp"

	"github.com/cosmos/cosmos-sdk/x/supply"
)

const (
	ModuleName = "posts"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	MaxPostMessageLength            = 500
	MaxOptionalDataFieldsNumber     = 10
	MaxOptionalDataFieldValueLength = 200

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
)

var (
	Sha256RegEx    = regexp.MustCompile(`^[a-fA-F0-9]{64}$`)
	HashtagRegEx   = regexp.MustCompile(`[^\S]|^#([^\s#.,!)]+)$`)
	ShortCodeRegEx = regexp.MustCompile(`:[a-z0-9+-]([a-z0-9\d_-])*:`)
	URIRegEx       = regexp.MustCompile(
		`^(?:http(s)?://)[\w.-]+(?:\.[\w.-]+)+[\w\-._~:/?#[\]@!$&'()*+,;=.]+$`)

	ModuleAddress = supply.NewModuleAddress(ModuleName)

	PostStorePrefix          = []byte("post")
	PostCommentsStorePrefix  = []byte("comments")
	PostReactionsStorePrefix = []byte("p_reactions")
	ReactionsStorePrefix     = []byte("reactions")
	PollAnswersStorePrefix   = []byte("poll_answers")
)

// PostStoreKey turns an id to a key used to store a post into the posts store
// nolint: interfacer
func PostStoreKey(id PostID) []byte {
	return append(PostStorePrefix, []byte(id)...)
}

// PostCommentsStoreKey turns an id to a key used to store a post's comments into the posts store
// nolint: interfacer
func PostCommentsStoreKey(id PostID) []byte {
	return append(PostCommentsStorePrefix, []byte(id)...)
}

// PostCommentsStoreKey turns an id to a key used to store a post's reactions into the posts store
// nolint: interfacer
func PostReactionsStoreKey(id PostID) []byte {
	return append(PostReactionsStorePrefix, []byte(id)...)
}

// ReactionsStoreKey turns the combination of shortCode and subspace to a key used to store a reaction into the reaction's store
// nolint: interfacer
func ReactionsStoreKey(shortCode, subspace string) []byte {
	return append(ReactionsStorePrefix, []byte(shortCode+subspace)...)
}

// PollAnswersStoreKey turns an id to a key used to store a post's poll answers into the posts store
// nolint: interfacer
func PollAnswersStoreKey(id PostID) []byte {
	return append(PollAnswersStorePrefix, []byte(id)...)
}
