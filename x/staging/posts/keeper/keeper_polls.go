package keeper

import (
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

// SaveUserAnswer save the poll's answers associated with the given postID inside the current context
// It assumes that the post exists and has a Poll inside it.
// If answer are already present, the old ones will be overridden.
func (k Keeper) SaveUserAnswer(ctx sdk.Context, answer types.UserAnswer) {
	store := ctx.KVStore(k.storeKey)

	sort.Slice(answer.Answers, func(i, j int) bool {
		return answer.Answers[i] < answer.Answers[j]
	})
	bz := types.MustMarshalUserAnswer(k.cdc, answer)
	store.Set(types.PollAnswersStoreKey(answer.PostID, answer.User), bz)
}

// GetUserAnswer returns the list of all the post polls answers associated with the given postID that are stored into the current state.
//
func (k Keeper) GetUserAnswer(ctx sdk.Context, postID, user string) (types.UserAnswer, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.PollAnswersStoreKey(postID, user)
	if !store.Has(key) {
		return types.UserAnswer{}, false
	}
	return types.MustUnmarshalUserAnswer(k.cdc, key), true
}

// GetAllUserAnswers returns the list of all the post polls answers associated with the given postID that are stored into the current state.
func (k Keeper) GetUserAnswersByPost(ctx sdk.Context, postID string) []types.UserAnswer {
	var answers []types.UserAnswer
	k.IteratePollAnswersByID(ctx, postID, func(_ int64, answer types.UserAnswer) bool {
		answers = append(answers, answer)
		return false
	})
	return answers
}

// GetAllUserAnswers returns the list of all the post polls answers associated with the given postID that are stored into the current state.
func (k Keeper) GetAllUserAnswers(ctx sdk.Context) []types.UserAnswer {
	var answers []types.UserAnswer
	k.IteratePollAnswers(ctx, func(_ int64, answer types.UserAnswer) bool {
		answers = append(answers, answer)
		return false
	})
	return answers
}
