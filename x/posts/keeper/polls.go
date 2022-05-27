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
	return types.IsPoll(attachment)
}

// GetPoll returns the poll having the given id.
// If not poll with the given id is found, the function returns nil and false.
func (k Keeper) GetPoll(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32) (poll *types.Poll, found bool) {
	attachment, found := k.GetAttachment(ctx, subspaceID, postID, pollID)
	if !found {
		return nil, false
	}

	poll, ok := attachment.Content.GetCachedValue().(*types.Poll)
	return poll, ok
}

// Tally iterates over the votes and returns the tally results of a poll
func (k Keeper) Tally(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32) *types.PollTallyResults {
	poll, found := k.GetPoll(ctx, subspaceID, postID, pollID)
	if !found {
		return nil
	}

	// Create the map index -> count(votes)
	results := make(map[uint32]uint64, len(poll.ProvidedAnswers))
	for i := range poll.ProvidedAnswers {
		results[uint32(i)] = 0
	}

	k.IteratePollUserAnswers(ctx, subspaceID, postID, pollID, func(_ int64, answer types.UserAnswer) (stop bool) {
		// Update the results
		for _, answerIndex := range answer.AnswersIndexes {
			results[answerIndex]++
		}

		// Delete the user answer
		k.DeleteUserAnswer(ctx, answer.SubspaceID, answer.PostID, answer.PollID, answer.User)

		return false
	})

	tallyResults := make([]types.PollTallyResults_AnswerResult, len(results))
	for index, count := range results {
		tallyResults[int(index)] = types.NewAnswerResult(index, count)
	}

	return types.NewPollTallyResults(tallyResults)
}

// --------------------------------------------------------------------------------------------------------------------

// InsertActivePollQueue inserts a poll into the active poll queue
func (k Keeper) InsertActivePollQueue(ctx sdk.Context, poll types.Attachment) {
	store := ctx.KVStore(k.storeKey)
	bz := types.GetPollIDBytes(poll.SubspaceID, poll.PostID, poll.ID)
	content := poll.Content.GetCachedValue().(*types.Poll)
	store.Set(types.ActivePollQueueKey(poll.SubspaceID, poll.PostID, poll.ID, content.EndDate), bz)
}

// RemoveFromActivePollQueue removes a poll from the active poll queue
func (k Keeper) RemoveFromActivePollQueue(ctx sdk.Context, poll types.Attachment) {
	store := ctx.KVStore(k.storeKey)
	content := poll.Content.GetCachedValue().(*types.Poll)
	store.Delete(types.ActivePollQueueKey(poll.SubspaceID, poll.PostID, poll.ID, content.EndDate))
}

// --------------------------------------------------------------------------------------------------------------------

// SaveUserAnswer stores the given poll answer into the current context
func (k Keeper) SaveUserAnswer(ctx sdk.Context, answer types.UserAnswer) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PollAnswerStoreKey(answer.SubspaceID, answer.PostID, answer.PollID, answer.User), k.cdc.MustMarshal(&answer))

	k.AfterPollAnswerSaved(ctx, answer.SubspaceID, answer.PostID, answer.PollID, answer.User)
}

// HasUserAnswer tells whether a user answer to the specified poll exists
func (k Keeper) HasUserAnswer(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32, user string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.PollAnswerStoreKey(subspaceID, postID, pollID, user))
}

// GetUserAnswer returns the user answer from the given user for the specified poll.
// If there is no answer result associated with the given poll and user the function will return an empty answer and false.
func (k Keeper) GetUserAnswer(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32, user string) (answer types.UserAnswer, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.PollAnswerStoreKey(subspaceID, postID, pollID, user)
	if !store.Has(key) {
		return types.UserAnswer{}, false
	}
	k.cdc.MustUnmarshal(store.Get(key), &answer)
	return answer, true
}

// DeleteUserAnswer deletes the user answer from the provided poll
func (k Keeper) DeleteUserAnswer(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32, user string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.PollAnswerStoreKey(subspaceID, postID, pollID, user))

	k.AfterPollAnswerDeleted(ctx, subspaceID, postID, pollID, user)
}
