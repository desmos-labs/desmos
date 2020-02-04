package types

import (
	"regexp"
)

var (
	SubspaceRegEx = regexp.MustCompile("^[a-fA-F0-9]{64}$")
)

const (
	ModuleName = "posts"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	PostStorePrefix          = "post:"
	LastPostIDStoreKey       = "last_post_id"
	PostCommentsStorePrefix  = "comments:"
	PostReactionsStorePrefix = "reactions:"
	PollAnswersStorePrefix   = "poll_answers:"

	MaxPostMessageLength            = 500
	MaxOptionalDataFieldsNumber     = 10
	MaxOptionalDataFieldValueLength = 200

	ActionCreatePost         = "create_post"
	ActionEditPost           = "edit_post"
	ActionAnswerPoll         = "answer_poll"
	ActionAddPostReaction    = "add_post_reaction"
	ActionRemovePostReaction = "remove_post_reaction"

	// Queries
	QuerierRoute     = ModuleName
	QueryPost        = "post"
	QueryPosts       = "posts"
	QueryPollAnswers = "poll-answers"
)
