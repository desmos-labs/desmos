package keeper

import (
	"bytes"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/posts/types"
)

// SavePollAnswers save the poll's answers associated with the given postID inside the current context
// It assumes that the post exists and has a Poll inside it.
// If answer are already present, the old ones will be overridden.
func (k Keeper) SavePollAnswers(ctx sdk.Context, postID string, userPollAnswers types.UserAnswer) {
	store := ctx.KVStore(k.storeKey)

	sort.Slice(userPollAnswers.Answers, func(i, j int) bool {
		return userPollAnswers.Answers[i] < userPollAnswers.Answers[j]
	})

	answers := k.GetPollAnswers(ctx, postID)
	if appendedAnswers, appended := types.AppendIfMissingOrIfUsersEquals(answers, userPollAnswers); appended {
		bz := types.MustMarshalUserAnswers(k.cdc, appendedAnswers)
		store.Set(types.PollAnswersStoreKey(postID), bz)
	}
}

// GetPollAnswers returns the list of all the post polls answers associated with the given postID that are stored into the current state.
func (k Keeper) GetPollAnswers(ctx sdk.Context, postID string) []types.UserAnswer {
	store := ctx.KVStore(k.storeKey)
	return types.MustUnmarshalUserAnswers(k.cdc, store.Get(types.PollAnswersStoreKey(postID)))
}

// GetUserPollAnswersEntries allows to returns the list of answers that have been stored inside the given context
func (k Keeper) GetUserAnswersEntries(ctx sdk.Context) []types.UserAnswersEntry {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PollAnswersStorePrefix)

	var usersAnswersData []types.UserAnswersEntry
	for ; iterator.Valid(); iterator.Next() {
		userAnswers := types.MustUnmarshalUserAnswers(k.cdc, iterator.Value())
		idBytes := bytes.TrimPrefix(iterator.Key(), types.PollAnswersStorePrefix)
		usersAnswersData = append(usersAnswersData, types.NewUserAnswersEntry(string(idBytes), userAnswers))
	}

	return usersAnswersData
}

// GetPollAnswersByUser retrieves post poll answers associated to the given ID and filtered by user
func (k Keeper) GetPollAnswersByUser(ctx sdk.Context, postID string, user string) []string {
	postPollAnswers := k.GetPollAnswers(ctx, postID)

	for _, postPollAnswers := range postPollAnswers {
		if user == postPollAnswers.User {
			return postPollAnswers.Answers
		}
	}
	return nil
}
