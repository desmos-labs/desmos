package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	types2 "github.com/desmos-labs/desmos/x/posts/types"
)

func (suite *KeeperTestSuite) TestKeeper_SaveUserAnswer() {
	tests := []struct {
		name            string
		postID          string
		storedAnswers   []types2.UserAnswer
		answer          types2.UserAnswer
		expectedAnswers []types2.UserAnswer
	}{
		{
			name:   "Save answers with no previous answers in this context",
			postID: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			answer: types2.NewUserAnswer(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				[]string{"1", "2"},
			),
			storedAnswers: nil,
			expectedAnswers: []types2.UserAnswer{
				types2.NewUserAnswer("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", []string{"1", "2"}),
			},
		},
		{
			name:   "Save new answers",
			postID: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			answer: types2.NewUserAnswer(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				[]string{"1"},
			),
			storedAnswers: []types2.UserAnswer{
				types2.NewUserAnswer("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", []string{"1", "2"}),
			},
			expectedAnswers: []types2.UserAnswer{
				types2.NewUserAnswer("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", []string{"1", "2"}),
				types2.NewUserAnswer("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af", "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", []string{"1"}),
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
			suite.Require().Equal(test.expectedAnswers, suite.k.GetAllUserAnswers(suite.ctx))
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetUserAnswersByPost() {
	tests := []struct {
		name      string
		store     func(sdk.Context)
		postID    string
		expStored []types2.UserAnswer
	}{
		{
			name:      "No answers returns empty list",
			postID:    "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			expStored: nil,
		},
		{
			name: "Answers returned correctly",
			store: func(ctx sdk.Context) {
				answers := []types2.UserAnswer{
					types2.NewUserAnswer(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						[]string{"1", "2"},
					),
					types2.NewUserAnswer(
						"29de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						[]string{"1", "2"},
					),
				}
				for _, answer := range answers {
					suite.k.SaveUserAnswer(ctx, answer)
				}
			},
			postID: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			expStored: []types2.UserAnswer{
				types2.NewUserAnswer(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					[]string{"1", "2"},
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store(suite.ctx)
			}
			actual := suite.k.GetUserAnswersByPost(suite.ctx, test.postID)
			suite.Require().Equal(test.expStored, actual)
		})
	}
}

func (suite *KeeperTestSuite) Test_GetUserAnswer() {
	tests := []struct {
		name        string
		store       func(ctx sdk.Context)
		postID      string
		user        string
		shouldFound bool
	}{
		{
			name:        "non existent answer return false",
			postID:      "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			user:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			shouldFound: false,
		},
		{
			name: "existing answer return true",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserAnswer(
					ctx,
					types2.NewUserAnswer(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						[]string{"1", "2"},
					),
				)
			},
			postID:      "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			user:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			shouldFound: true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store(suite.ctx)
			}
			_, actual := suite.k.GetUserAnswer(suite.ctx, test.postID, test.user)
			suite.Require().Equal(test.shouldFound, actual)
		})
	}
}
