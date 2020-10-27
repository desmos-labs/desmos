package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/types"
)

func (suite *KeeperTestSuite) TestKeeper_SavePollPostAnswers() {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	user, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.Require().NoError(err)

	user2, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.Require().NoError(err)

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
			postID:             id,
			userAnswersDetails: types.NewUserAnswer(answers, user),
			previousUsersAD:    nil,
			expUsersAD:         types.UserAnswers{types.NewUserAnswer(answers, user)},
		},
		{
			name:               "Save new answers",
			postID:             id,
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
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.StoreKey)

			if test.previousUsersAD != nil {
				store.Set(types.PollAnswersStoreKey(test.postID), suite.keeper.Cdc.MustMarshalBinaryBare(test.previousUsersAD))
			}

			suite.keeper.SavePollAnswers(suite.ctx, test.postID, test.userAnswersDetails)

			var actualUsersAnswersDetails types.UserAnswers
			answersBz := store.Get(types.PollAnswersStoreKey(test.postID))
			suite.keeper.Cdc.MustUnmarshalBinaryBare(answersBz, &actualUsersAnswersDetails)
			suite.Require().Equal(test.expUsersAD, actualUsersAnswersDetails)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetPostPollAnswersDetails() {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	user, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.Require().NoError(err)

	answers := []types.AnswerID{types.AnswerID(1), types.AnswerID(2)}

	tests := []struct {
		name          string
		postID        types.PostID
		storedAnswers types.UserAnswers
	}{
		{
			name:          "No answers returns empty list",
			postID:        id,
			storedAnswers: nil,
		},
		{
			name:          "Answers returned correctly",
			postID:        id,
			storedAnswers: types.UserAnswers{types.NewUserAnswer(answers, user)},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			if test.storedAnswers != nil {
				suite.keeper.SavePollAnswers(suite.ctx, test.postID, test.storedAnswers[0])
			}

			actualPostPollAnswers := suite.keeper.GetPollAnswers(suite.ctx, test.postID)

			suite.Require().Equal(test.storedAnswers, actualPostPollAnswers)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetPostPollAnswersByUser() {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	user, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.Require().NoError(err)

	user2, err := sdk.AccAddressFromBech32("cosmos1jlhazemxvu0zn9y77j6afwmpf60zveqw5480l2")
	suite.Require().NoError(err)

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
			postID:        id,
			user:          user2,
			expAnswers:    nil,
		},
		{
			name:          "Matching user returns answers made by him",
			storedAnswers: types.NewUserAnswer(answers, user),
			postID:        id,
			user:          user,
			expAnswers:    answers,
		},
	}

	for _, test := range tests {
		suite.keeper.SavePollAnswers(suite.ctx, test.postID, test.storedAnswers)

		actualPostPollAnswers := suite.keeper.GetPollAnswersByUser(suite.ctx, test.postID, test.user)
		suite.Require().Equal(test.expAnswers, actualPostPollAnswers)
	}
}
