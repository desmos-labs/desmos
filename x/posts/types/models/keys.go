package models

import "regexp"

var (
	hashtagRegEx = regexp.MustCompile(`[^\S]|^#([^\s#.,!)]+)$`)
)

// PostStoreKey turns an id to a key used to store a post into the posts store
//nolint: interfacer
func PostStoreKey(id PostID) []byte {
	return append(PostStorePrefix, []byte(id)...)
}

// PostIndexedIDStoreKey turns an id to a key used to store an incremental ID into the posts store
func PostIndexedIDStoreKey(id PostID) []byte {
	return append(PostIDStorePrefix, []byte(id)...)
}

// PostCommentsStoreKey turns an id to a key used to store a post's comments into the posts store
//nolint: interfacer
func PostCommentsStoreKey(id PostID) []byte {
	return append(PostCommentsStorePrefix, []byte(id)...)
}

// PostCommentsStoreKey turns an id to a key used to store a post's reactions into the posts store
//nolint: interfacer
func PostReactionsStoreKey(id PostID) []byte {
	return append(PostReactionsStorePrefix, []byte(id)...)
}

// ReactionsStoreKey turns the combination of shortCode and subspace to a key used to store a reaction into the reaction's store
//nolint: interfacer
func ReactionsStoreKey(shortCode, subspace string) []byte {
	return append(ReactionsStorePrefix, []byte(shortCode+subspace)...)
}

// PollAnswersStoreKey turns an id to a key used to store a post's poll answers into the posts store
//nolint: interfacer
func PollAnswersStoreKey(id PostID) []byte {
	return append(PollAnswersStorePrefix, []byte(id)...)
}
