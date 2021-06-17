package keeper_test

import (
	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

func (suite *KeeperTestSuite) TestKeeper_SavePollAnswers() {
	tests := []struct {
		name            string
		postID          string
		storedAnswers   []types.UserAnswer
		answer          types.UserAnswer
		expectedAnswers []types.UserAnswer
	}{
		{
			name:   "Save answers with no previous answers in this context",
			postID: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			answer: types.NewUserAnswer(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				[]string{"1", "2"},
			),
			storedAnswers: nil,
			expectedAnswers: []types.UserAnswer{
				types.NewUserAnswer("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", []string{"1", "2"}),
			},
		},
		{
			name:   "Save new answers",
			postID: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			answer: types.NewUserAnswer(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				[]string{"1"},
			),
			storedAnswers: []types.UserAnswer{
				types.NewUserAnswer("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", []string{"1", "2"}),
			},
			expectedAnswers: []types.UserAnswer{
				types.NewUserAnswer("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", []string{"1", "2"}),
				types.NewUserAnswer("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af", "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", []string{"1"}),
			},
		},
	}

	for _, test := range tests {

		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, answer := range test.storedAnswers {
				suite.k.SaveUserAnswer(suite.ctx, answer)
			}

			suite.k.SaveUserAnswer(suite.ctx, test.answer)
			suite.Require().Equal(test.expectedAnswers, suite.k.GetUserAnswersByID(suite.ctx, test.postID))
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
				types.NewUserAnswer("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", []string{"1", "2"}),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			if test.storedAnswers != nil {
				suite.k.SaveUserAnswer(suite.ctx, test.storedAnswers[0])
			}
			actual := suite.k.GetUserAnswersByID(suite.ctx, test.postID)
			suite.Require().Equal(test.storedAnswers, actual)
		})
	}
}
