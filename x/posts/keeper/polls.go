package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
)

// HasPoll tells whether the specified post contains a poll with the provided id
func (k Keeper) HasPoll(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32) bool {
	attachment, found := k.GetAttachment(ctx, subspaceID, postID, pollID)
	if !found {
		return false
	}

	_, ok := attachment.Sum.(*types.Attachment_Poll)
	return ok
}

// --------------------------------------------------------------------------------------------------------------------

// SaveUserAnswer stores the given poll answer into the current context
func (k Keeper) SaveUserAnswer(ctx sdk.Context, answer types.UserAnswer) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PollAnswerStoreKey(answer.SubspaceID, answer.PostID, answer.PollID, answer.User), k.cdc.MustMarshal(&answer))

	k.AfterPollAnswerSaved(ctx, answer.SubspaceID, answer.PostID, answer.PollID, answer.User)
}

// HasUserAnswer tells whether a user answer to the specified poll exists
func (k Keeper) HasUserAnswer(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32, user sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.PollAnswerStoreKey(subspaceID, postID, pollID, user))
}

// GetUserAnswer returns the user answer from the given user for the specified poll.
// If there is no answer result associated with the given poll and user the function will return an empty answer and false.
func (k Keeper) GetUserAnswer(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32, user sdk.AccAddress) (answer types.UserAnswer, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.PollAnswerStoreKey(subspaceID, postID, pollID, user)
	if !store.Has(key) {
		return types.UserAnswer{}, false
	}
	k.cdc.MustUnmarshal(store.Get(key), &answer)
	return answer, true
}

// DeleteUserAnswer deletes the user answer from the provided poll
func (k Keeper) DeleteUserAnswer(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32, user sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.PollAnswerStoreKey(subspaceID, postID, pollID, user))

	k.AfterPollAnswerDeleted(ctx, subspaceID, postID, pollID, user)
}

// --------------------------------------------------------------------------------------------------------------------

// SavePollTallyResults stores the given results inside the current context
func (k Keeper) SavePollTallyResults(ctx sdk.Context, results types.PollTallyResults) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PollTallyResultsStoreKey(results.SubspaceID, results.PostID, results.PollID), k.cdc.MustMarshal(&results))

	k.AfterPollTallyResultsSaved(ctx, results.SubspaceID, results.PostID, results.PollID)
}

// HasPollTallyResults tells whether the tally results for the specified poll exist or not
func (k Keeper) HasPollTallyResults(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.PollTallyResultsStoreKey(subspaceID, postID, pollID))
}

// GetPollTallyResults returns the tally results from the given poll.
// If there is no tally result associated with the given poll the function will return an empty result and false.
func (k Keeper) GetPollTallyResults(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32) (results types.PollTallyResults, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.PollTallyResultsStoreKey(subspaceID, postID, pollID)
	if !store.Has(key) {
		return types.PollTallyResults{}, false
	}

	k.cdc.MustUnmarshal(store.Get(key), &results)
	return results, true
}

// DeletePollTallyResults deletes the tally results for the given poll
func (k Keeper) DeletePollTallyResults(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.PollTallyResultsStoreKey(subspaceID, postID, pollID))

	k.AfterPollTallyResultsDeleted(ctx, subspaceID, postID, pollID)
}
