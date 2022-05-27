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
