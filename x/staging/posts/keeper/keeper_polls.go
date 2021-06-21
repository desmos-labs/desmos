package keeper

import (
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

// SaveUserAnswer save the user answer inside the current context
// It assumes that the post exists and has a Poll inside it.
// If answer are already present, the old ones will be overridden.
func (k Keeper) SaveUserAnswer(ctx sdk.Context, answer types.UserAnswer) {
	store := ctx.KVStore(k.storeKey)

	sort.Slice(answer.Answers, func(i, j int) bool {
		return answer.Answers[i] < answer.Answers[j]
	})
	bz := types.MustMarshalUserAnswer(k.cdc, answer)
	store.Set(types.UserAnswersStoreKey(answer.PostID, answer.User), bz)
}

// GetUserAnswer returns the user answer created by the given user address and associated to the post having the given id.
// If no user answer could be found, returns false instead.
func (k Keeper) GetUserAnswer(ctx sdk.Context, postID, user string) (types.UserAnswer, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.UserAnswersStoreKey(postID, user)
	if !store.Has(key) {
		return types.UserAnswer{}, false
	}
	return types.MustUnmarshalUserAnswer(k.cdc, store.Get(key)), true
}

// GetUserAnswersByPost returns the list of all the user answers associated with the given postID that are stored into the current state.
func (k Keeper) GetUserAnswersByPost(ctx sdk.Context, postID string) []types.UserAnswer {
	var answers []types.UserAnswer
	k.IterateUserAnswersByPost(ctx, postID, func(_ int64, answer types.UserAnswer) bool {
		answers = append(answers, answer)
		return false
	})
	return answers
}

// GetAllUserAnswers returns all the user answers.
func (k Keeper) GetAllUserAnswers(ctx sdk.Context) []types.UserAnswer {
	var answers []types.UserAnswer
	k.IterateUserAnswers(ctx, func(_ int64, answer types.UserAnswer) bool {
		answers = append(answers, answer)
		return false
	})
	return answers
}
