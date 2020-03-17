package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/require"
)

func TestKeeper_SavePollPostAnswers(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	user2, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	answers := []types.AnswerID{types.AnswerID(1), types.AnswerID(2)}
	answers2 := []types.AnswerID{types.AnswerID(1)}

	tests := []struct {
		name               string
		postID             types.PostID
		userAnswersDetails types.UserAnswer
		previousUsersAD    types.UserAnswers
		expUsersAD         types.UserAnswers
	}{
		{
			name:               "Save answers with no previous answers in this context",
			postID:             types.PostID(1),
			userAnswersDetails: types.NewUserAnswer(answers, user),
			previousUsersAD:    nil,
			expUsersAD:         types.UserAnswers{types.NewUserAnswer(answers, user)},
		},
		{
			name:               "Save new answers",
			postID:             types.PostID(1),
			userAnswersDetails: types.NewUserAnswer(answers2, user2),
			previousUsersAD:    types.UserAnswers{types.NewUserAnswer(answers, user)},
			expUsersAD: types.UserAnswers{
				types.NewUserAnswer(answers, user),
				types.NewUserAnswer(answers2, user2),
			},
		},
	}

	for _, test := range tests {

		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)

			if test.previousUsersAD != nil {
				store.Set(types.PollAnswersStoreKey(test.postID), k.Cdc.MustMarshalBinaryBare(test.previousUsersAD))
			}

			k.SavePollAnswers(ctx, test.postID, test.userAnswersDetails)

			var actualUsersAnswersDetails types.UserAnswers
			answersBz := store.Get(types.PollAnswersStoreKey(test.postID))
			k.Cdc.MustUnmarshalBinaryBare(answersBz, &actualUsersAnswersDetails)
			require.Equal(t, test.expUsersAD, actualUsersAnswersDetails)
		})
	}
}

func TestKeeper_GetPostPollAnswersDetails(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	answers := []types.AnswerID{types.AnswerID(1), types.AnswerID(2)}

	tests := []struct {
		name          string
		postID        types.PostID
		storedAnswers types.UserAnswers
	}{
		{
			name:          "No answers returns empty list",
			postID:        types.PostID(1),
			storedAnswers: nil,
		},
		{
			name:          "Answers returned correctly",
			postID:        types.PostID(1),
			storedAnswers: types.UserAnswers{types.NewUserAnswer(answers, user)},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			if test.storedAnswers != nil {
				k.SavePollAnswers(ctx, test.postID, test.storedAnswers[0])
			}

			actualPostPollAnswers := k.GetPollAnswers(ctx, test.postID)

			require.Equal(t, test.storedAnswers, actualPostPollAnswers)
		})
	}
}

func TestKeeper_GetPostPollAnswersByUser(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	user2, err := sdk.AccAddressFromBech32("cosmos1jlhazemxvu0zn9y77j6afwmpf60zveqw5480l2")
	require.NoError(t, err)

	answers := []types.AnswerID{types.AnswerID(1), types.AnswerID(2)}

	tests := []struct {
		name          string
		storedAnswers types.UserAnswer
		postID        types.PostID
		user          sdk.AccAddress
		expAnswers    []types.AnswerID
	}{
		{
			name:          "No answers for user returns nil",
			storedAnswers: types.NewUserAnswer(answers, user),
			postID:        types.PostID(1),
			user:          user2,
			expAnswers:    nil,
		},
		{
			name:          "Matching user returns answers made by him",
			storedAnswers: types.NewUserAnswer(answers, user),
			postID:        types.PostID(1),
			user:          user,
			expAnswers:    answers,
		},
	}

	for _, test := range tests {
		ctx, k := SetupTestInput()
		k.SavePollAnswers(ctx, test.postID, test.storedAnswers)

		actualPostPollAnswers := k.GetPollAnswersByUser(ctx, test.postID, test.user)
		require.Equal(t, test.expAnswers, actualPostPollAnswers)
	}
}

func TestKeeper_GetAnswersDetailsMap(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	user2, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	require.NoError(t, err)

	answers := []types.AnswerID{types.AnswerID(1), types.AnswerID(2)}

	tests := []struct {
		name    string
		usersAD map[types.PostID]types.UserAnswers
	}{
		{
			name:    "Empty users answers details data are returned correctly",
			usersAD: map[types.PostID]types.UserAnswers{},
		},
		{
			name: "Non empty users answers details data are returned correcly",
			usersAD: map[types.PostID]types.UserAnswers{
				types.PostID(1): {
					types.NewUserAnswer(answers, user),
					types.NewUserAnswer(answers, user2),
				},
				types.PostID(2): {
					types.NewUserAnswer(answers, user2),
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)
			for postID, userAD := range test.usersAD {
				store.Set(types.PollAnswersStoreKey(postID), k.Cdc.MustMarshalBinaryBare(userAD))
			}

			usersADData := k.GetPollAnswersMap(ctx)
			require.Equal(t, test.usersAD, usersADData)
		})
	}
}
