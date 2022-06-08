package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/reactions/types"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

type Hooks struct {
	k Keeper
}

var (
	_ subspacestypes.SubspacesHooks = Hooks{}
)

// Hooks creates a new reports hooks
func (k Keeper) Hooks() Hooks { return Hooks{k} }

// AfterSubspaceSaved implements subspacestypes.Hooks
func (h Hooks) AfterSubspaceSaved(ctx sdk.Context, subspaceID uint64) {
	// Create the initial registered reaction and reaction id
	if !h.k.HasNextRegisteredReactionID(ctx, subspaceID) {
		h.k.SetNextRegisteredReactionID(ctx, subspaceID, 1)
	}

	// Crete the initial reactions params
	if !h.k.HasSubspaceReactionsParams(ctx, subspaceID) {
		h.k.SaveSubspaceReactionsParams(ctx, types.DefaultReactionsParams(subspaceID))
	}
}

// AfterSubspaceDeleted implements subspacestypes.Hooks
func (h Hooks) AfterSubspaceDeleted(ctx sdk.Context, subspaceID uint64) {
	// Delete the next registered reaction id key
	h.k.DeleteNextRegisteredReactionID(ctx, subspaceID)

	// Delete all the registered reactions related to this subspace
	h.k.IterateSubspaceRegisteredReactions(ctx, subspaceID, func(reaction types.RegisteredReaction) (stop bool) {
		h.k.DeleteRegisteredReaction(ctx, reaction.SubspaceID, reaction.ID)
		return false
	})

	// Delete the reactions params
	h.k.DeleteSubspaceReactionsParams(ctx, subspaceID)
}

// AfterSubspaceGroupSaved implements subspacestypes.Hooks
func (h Hooks) AfterSubspaceGroupSaved(sdk.Context, uint64, uint32) {}

// AfterSubspaceGroupDeleted implements subspacestypes.Hooks
func (h Hooks) AfterSubspaceGroupDeleted(sdk.Context, uint64, uint32) {}

// AfterSubspaceSectionSaved implements subspacestypes.Hooks
func (h Hooks) AfterSubspaceSectionSaved(sdk.Context, uint64, uint32) {}

// AfterSubspaceSectionDeleted implements subspacestypes.Hooks
func (h Hooks) AfterSubspaceSectionDeleted(sdk.Context, uint64, uint32) {}

// AfterSubspaceGroupMemberAdded implements subspacestypes.Hooks
func (h Hooks) AfterSubspaceGroupMemberAdded(sdk.Context, uint64, uint32, string) {
}

// AfterSubspaceGroupMemberRemoved implements subspacestypes.Hooks
func (h Hooks) AfterSubspaceGroupMemberRemoved(sdk.Context, uint64, uint32, string) {
}

// AfterUserPermissionSet implements subspacestypes.Hooks
func (h Hooks) AfterUserPermissionSet(sdk.Context, uint64, uint32, string, subspacestypes.Permission) {
}

// AfterUserPermissionRemoved implements subspacestypes.Hooks
func (h Hooks) AfterUserPermissionRemoved(sdk.Context, uint64, uint32, string) {
}

// AfterPostSaved implements poststypes.PostsHooks
func (h Hooks) AfterPostSaved(ctx sdk.Context, subspaceID uint64, postID uint64) {
	// Set the next reaction id
	if !h.k.HasNextReactionID(ctx, subspaceID, postID) {
		h.k.SetNextReactionID(ctx, subspaceID, postID, 1)
	}
}

// AfterPostDeleted implements poststypes.PostsHooks
func (h Hooks) AfterPostDeleted(ctx sdk.Context, subspaceID uint64, postID uint64) {
	// Delete the next reaction id key
	h.k.DeleteNextReactionID(ctx, postID, subspaceID)

	// Delete all the reactions related to this post
	h.k.IteratePostReactions(ctx, subspaceID, postID, func(reaction types.Reaction) (stop bool) {
		h.k.DeleteReaction(ctx, reaction.SubspaceID, reaction.PostID, reaction.ID)
		return false
	})
}

// AfterAttachmentSaved implements poststypes.PostsHooks
func (h Hooks) AfterAttachmentSaved(sdk.Context, uint64, uint64, uint32) {}

// AfterAttachmentDeleted implements poststypes.PostsHooks
func (h Hooks) AfterAttachmentDeleted(sdk.Context, uint64, uint64, uint32) {}

// AfterPollAnswerSaved implements poststypes.PostsHooks
func (h Hooks) AfterPollAnswerSaved(sdk.Context, uint64, uint64, uint32, string) {}

// AfterPollAnswerDeleted implements poststypes.PostsHooks
func (h Hooks) AfterPollAnswerDeleted(sdk.Context, uint64, uint64, uint32, string) {}

// AfterPollVotingPeriodEnded implements poststypes.PostsHooks
func (h Hooks) AfterPollVotingPeriodEnded(sdk.Context, uint64, uint64, uint32) {}
