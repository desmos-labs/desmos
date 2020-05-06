package common

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

	// Sorting
	PostSortByCreationDate  = "created"
	PostSortByID            = "id"
	PostSortOrderAscending  = "ascending"
	PostSortOrderDescending = "descending"
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
