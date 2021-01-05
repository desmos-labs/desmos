package keeper_test

import (
	"github.com/desmos-labs/desmos/x/posts/types"
)

func (suite *KeeperTestSuite) TestKeeper_SavePollAnswers() {
	tests := []struct {
		name            string
		storedAnswers   []types.UserAnswer
		postID          string
		answer          types.UserAnswer
		expectedAnswers []types.UserAnswer
	}{
		{
			name:   "Save answers with no previous answers in this context",
			postID: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			answer: types.NewUserAnswer(
				[]string{"1", "2"},
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			storedAnswers: nil,
			expectedAnswers: []types.UserAnswer{
				types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			},
		},
		{
			name:   "Save new answers",
			postID: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			answer: types.NewUserAnswer(
				[]string{"1"},
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			storedAnswers: []types.UserAnswer{
				types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			},
			expectedAnswers: []types.UserAnswer{
				types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
				types.NewUserAnswer([]string{"1"}, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
			},
		},
	}

	for _, test := range tests {

		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, answer := range test.storedAnswers {
				suite.keeper.SavePollAnswers(suite.ctx, test.postID, answer)
			}

			suite.keeper.SavePollAnswers(suite.ctx, test.postID, test.answer)
			suite.Require().Equal(test.expectedAnswers, suite.keeper.GetPollAnswers(suite.ctx, test.postID))
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetPollAnswers() {
	tests := []struct {
		name          string
		postID        string
		storedAnswers []types.UserAnswer
	}{
		{
			name:          "No answers returns empty list",
			postID:        "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			storedAnswers: nil,
		},
		{
			name:   "Answers returned correctly",
			postID: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			storedAnswers: []types.UserAnswer{
				types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			if test.storedAnswers != nil {
				suite.keeper.SavePollAnswers(suite.ctx, test.postID, test.storedAnswers[0])
			}

			actual := suite.keeper.GetPollAnswers(suite.ctx, test.postID)
			suite.Require().Equal(test.storedAnswers, actual)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetPollAnswersByUser() {
	tests := []struct {
		name          string
		storedAnswers types.UserAnswer
		postID        string
		user          string
		expAnswers    []string
	}{
		{
			name:          "No answers for user returns nil",
			storedAnswers: types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			postID:        "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			user:          "cosmos1jlhazemxvu0zn9y77j6afwmpf60zveqw5480l2",
			expAnswers:    nil,
		},
		{
			name:          "Matching user returns answers made by him",
			storedAnswers: types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			postID:        "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			user:          "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			expAnswers:    []string{"1", "2"},
		},
	}

	for _, test := range tests {
		suite.keeper.SavePollAnswers(suite.ctx, test.postID, test.storedAnswers)

		actual := suite.keeper.GetPollAnswersByUser(suite.ctx, test.postID, test.user)
		suite.Require().Equal(test.expAnswers, actual)
	}
}
