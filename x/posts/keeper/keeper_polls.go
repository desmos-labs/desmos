package keeper

import (
	"bytes"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/posts/types"
)

// SavePollAnswers save the poll's answers associated with the given postID inside the current context
// It assumes that the post exists and has a Poll inside it.
// If userAnswersDetails are already present, the old ones will be overridden.
func (k Keeper) SavePollAnswers(ctx sdk.Context, postID types.PostID, userPollAnswers types.UserAnswer) {
	store := ctx.KVStore(k.StoreKey)

	sort.Slice(
		userPollAnswers.Answers,
		func(i, j int) bool { return userPollAnswers.Answers[i] < userPollAnswers.Answers[j] },
	)

	usersAnswersDetails := k.GetPollAnswers(ctx, postID)

	if usersAnswersDetails, appended := usersAnswersDetails.AppendIfMissingOrIfUsersEquals(userPollAnswers); appended {
		store.Set(types.PollAnswersStoreKey(postID), k.Cdc.MustMarshalBinaryBare(&usersAnswersDetails))
	}

}

// GetPollAnswers returns the list of all the post polls answers associated with the given postID that are stored into the current state.
func (k Keeper) GetPollAnswers(ctx sdk.Context, postID types.PostID) types.UserAnswers {
	store := ctx.KVStore(k.StoreKey)

	var usersAnswersDetails types.UserAnswers
	answersBz := store.Get(types.PollAnswersStoreKey(postID))

	k.Cdc.MustUnmarshalBinaryBare(answersBz, &usersAnswersDetails)

	return usersAnswersDetails
}

// GetPollAnswersMap allows to returns the list of answers that have been stored inside the given context
func (k Keeper) GetPollAnswersMap(ctx sdk.Context) map[string]types.UserAnswers {
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PollAnswersStorePrefix)

	usersAnswersData := map[string]types.UserAnswers{}
	for ; iterator.Valid(); iterator.Next() {
		var userAnswers types.UserAnswers
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &userAnswers)
		idBytes := bytes.TrimPrefix(iterator.Key(), types.PollAnswersStorePrefix)
		usersAnswersData[string(idBytes)] = userAnswers
	}

	return usersAnswersData
}

// GetPollAnswersByUser retrieves post poll answers associated to the given ID and filtered by user
func (k Keeper) GetPollAnswersByUser(ctx sdk.Context, postID types.PostID, user sdk.AccAddress) []types.AnswerID {
	postPollAnswers := k.GetPollAnswers(ctx, postID)

	for _, postPollAnswers := range postPollAnswers {
		if user.Equals(postPollAnswers.User) {
			return postPollAnswers.Answers
		}
	}
	return nil
}
