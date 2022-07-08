package types

// DONTCOVER

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Event Hooks
// These can be utilized to communicate between a posts keeper and another
// keeper which must take particular actions when posts/attachments/polls change
// state. The second keeper must implement this interface, which then the
// posts keeper can call.

// PostsHooks event hooks for posts objects (noalias)
type PostsHooks interface {
	AfterPostSaved(ctx sdk.Context, subspaceID uint64, postID uint64)   // Must be called when a post is saved
	AfterPostDeleted(ctx sdk.Context, subspaceID uint64, postID uint64) // Must be called when a post is deleted

	AfterAttachmentSaved(ctx sdk.Context, subspaceID uint64, postID uint64, attachmentID uint32)   // Must be called when a post attachment is saved
	AfterAttachmentDeleted(ctx sdk.Context, subspaceID uint64, postID uint64, attachmentID uint32) // Must be called when a post attachment is deleted

	AfterPollAnswerSaved(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32, user string)   // Must be called when a poll user answer is saved
	AfterPollAnswerDeleted(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32, user string) // Must be called when a poll user answer is deleted

	AfterPollVotingPeriodEnded(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32) // Must be called when a poll tally results are saved
}

// --------------------------------------------------------------------------------------------------------------------

// MultiPostsHooks combines multiple posts hooks, all hook functions are run in array sequence
type MultiPostsHooks []PostsHooks

func NewMultiPostsHooks(hooks ...PostsHooks) MultiPostsHooks {
	return hooks
}

// AfterPostSaved implements PostsHooks
func (h MultiPostsHooks) AfterPostSaved(ctx sdk.Context, subspaceID uint64, postID uint64) {
	for _, hook := range h {
		hook.AfterPostSaved(ctx, subspaceID, postID)
	}
}

// AfterPostDeleted implements PostsHooks
func (h MultiPostsHooks) AfterPostDeleted(ctx sdk.Context, subspaceID uint64, postID uint64) {
	for _, hook := range h {
		hook.AfterPostDeleted(ctx, subspaceID, postID)
	}
}

// AfterAttachmentSaved implements PostsHooks
func (h MultiPostsHooks) AfterAttachmentSaved(ctx sdk.Context, subspaceID uint64, postID uint64, attachmentID uint32) {
	for _, hook := range h {
		hook.AfterAttachmentSaved(ctx, subspaceID, postID, attachmentID)
	}
}

// AfterAttachmentDeleted implements PostsHooks
func (h MultiPostsHooks) AfterAttachmentDeleted(ctx sdk.Context, subspaceID uint64, postID uint64, attachmentID uint32) {
	for _, hook := range h {
		hook.AfterAttachmentDeleted(ctx, subspaceID, postID, attachmentID)
	}
}

// AfterPollAnswerSaved implements PostsHooks
func (h MultiPostsHooks) AfterPollAnswerSaved(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32, user string) {
	for _, hook := range h {
		hook.AfterPollAnswerSaved(ctx, subspaceID, postID, pollID, user)
	}
}

// AfterPollAnswerDeleted implements PostsHooks
func (h MultiPostsHooks) AfterPollAnswerDeleted(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32, user string) {
	for _, hook := range h {
		hook.AfterPollAnswerDeleted(ctx, subspaceID, postID, pollID, user)
	}
}

// AfterPollVotingPeriodEnded implements PostsHooks
func (h MultiPostsHooks) AfterPollVotingPeriodEnded(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32) {
	for _, hook := range h {
		hook.AfterPollVotingPeriodEnded(ctx, subspaceID, postID, pollID)
	}
}
