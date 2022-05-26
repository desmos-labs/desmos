package keeper

// DONTCOVER

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
)

// Implement PostsHooks interface
var _ types.PostsHooks = Keeper{}

// AfterPostSaved implements types.PostsHooks
func (k Keeper) AfterPostSaved(ctx sdk.Context, subspaceID uint64, postID uint64) {
	if k.hooks != nil {
		k.hooks.AfterPostSaved(ctx, subspaceID, postID)
	}
}

// AfterPostDeleted implements types.PostsHooks
func (k Keeper) AfterPostDeleted(ctx sdk.Context, subspaceID uint64, postID uint64) {
	if k.hooks != nil {
		k.hooks.AfterPostDeleted(ctx, subspaceID, postID)
	}
}

// AfterAttachmentSaved implements types.PostsHooks
func (k Keeper) AfterAttachmentSaved(ctx sdk.Context, subspaceID uint64, postID uint64, attachmentID uint32) {
	if k.hooks != nil {
		k.hooks.AfterAttachmentSaved(ctx, subspaceID, postID, attachmentID)
	}
}

// AfterAttachmentDeleted implements types.PostsHooks
func (k Keeper) AfterAttachmentDeleted(ctx sdk.Context, subspaceID uint64, postID uint64, attachmentID uint32) {
	if k.hooks != nil {
		k.hooks.AfterAttachmentDeleted(ctx, subspaceID, postID, attachmentID)
	}
}

// AfterPollAnswerSaved implements types.PostsHooks
func (k Keeper) AfterPollAnswerSaved(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32, user string) {
	if k.hooks != nil {
		k.hooks.AfterPollAnswerSaved(ctx, subspaceID, postID, pollID, user)
	}
}

// AfterPollAnswerDeleted implements types.PostsHooks
func (k Keeper) AfterPollAnswerDeleted(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32, user string) {
	if k.hooks != nil {
		k.hooks.AfterPollAnswerDeleted(ctx, subspaceID, postID, pollID, user)
	}
}

// AfterPollVotingPeriodEnded implements types.PostsHooks
func (k Keeper) AfterPollVotingPeriodEnded(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32) {
	if k.hooks != nil {
		k.hooks.AfterPollVotingPeriodEnded(ctx, subspaceID, postID, pollID)
	}
}
