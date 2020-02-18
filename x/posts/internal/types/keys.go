package types

import (
	"regexp"
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

	// Queries
	QuerierRoute     = ModuleName
	QueryPost        = "post"
	QueryPosts       = "posts"
	QueryPollAnswers = "poll-answers"
)

var (
	SubspaceRegEx = regexp.MustCompile("^[a-fA-F0-9]{64}$")

	LastPostIDStoreKey       = []byte("last_post_id")
	PostStorePrefix          = []byte("post")
	PostCommentsStorePrefix  = []byte("comments")
	PostReactionsStorePrefix = []byte("reactions")
	PollAnswersStorePrefix   = []byte("poll_answers")
	HashtagPrefix = []byte("hashtag:")
)

// AddressStoreKey turns an id to a key used to store a post into the posts store
// nolint: interfacer
func PostStoreKey(id PostID) []byte {
	return append(PostStorePrefix, []byte(id.String())...)
}

// PostCommentsStoreKey turns an id to a key used to store a post's comments into the posts store
// nolint: interfacer
func PostCommentsStoreKey(id PostID) []byte {
	return append(PostCommentsStorePrefix, []byte(id.String())...)
}

// PostCommentsStoreKey turns an id to a key used to store a post's reactions into the posts store
// nolint: interfacer
func PostReactionsStoreKey(id PostID) []byte {
	return append(PostReactionsStorePrefix, []byte(id.String())...)
}

// PollAnswersStoreKey turns an id to a key used to store a post's poll answers into the posts store
// nolint: interfacer
func PollAnswersStoreKey(id PostID) []byte {
	return append(PollAnswersStorePrefix, []byte(id.String())...)
}
