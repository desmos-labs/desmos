package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

func (suite *KeeperTestSuite) Test_RegisteredReactions() {
	usecases := []struct {
		name            string
		storedReactions []types.RegisteredReaction
		req             *types.QueryRegisteredReactionsRequest
		expLen          int
	}{
		{
			name: "query registered reactions without subspace and pagination",
			storedReactions: []types.RegisteredReaction{
				types.NewRegisteredReaction(
					"creator",
					":smile:",
					"smile",
					"subspace1",
				),
				types.NewRegisteredReaction(
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					":fire:",
					"fire",
					"subspace2",
				),
			},
			req:    &types.QueryRegisteredReactionsRequest{},
			expLen: 2,
		},
		{
			name: "query registered reactions with a subspace",
			storedReactions: []types.RegisteredReaction{
				types.NewRegisteredReaction(
					"creator",
					":smile:",
					"smile",
					"subspace1",
				),
				types.NewRegisteredReaction(
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					":fire:",
					"fire",
					"subspace2",
				),
			},
			req:    &types.QueryRegisteredReactionsRequest{Subspace: "subspace1"},
			expLen: 1,
		},
		{
			name: "query registered reactions with pagination",
			storedReactions: []types.RegisteredReaction{
				types.NewRegisteredReaction(
					"creator",
					":smile:",
					"smile",
					"subspace1",
				),
				types.NewRegisteredReaction(
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					":fire:",
					"fire",
					"subspace2",
				),
			},
			req:    &types.QueryRegisteredReactionsRequest{Pagination: &query.PageRequest{Limit: 1}},
			expLen: 1,
		},
	}

	for _, uc := range usecases {
		suite.Run(uc.name, func() {
			suite.SetupTest()
			for _, reaction := range uc.storedReactions {
				suite.k.SaveRegisteredReaction(suite.ctx, reaction)
			}

			res, err := suite.k.RegisteredReactions(sdk.WrapSDKContext(suite.ctx), uc.req)
			suite.Require().NoError(err)
			suite.Require().NotNil(res)
			suite.Require().Equal(uc.expLen, len(res.RegisteredReactions))
		})
	}
}

func (suite *KeeperTestSuite) Test_PollAnswers() {
	usecases := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryPollAnswersRequest
		shouldErr bool
		expLen    int
	}{
		{
			name:      "invalid post id returns error",
			req:       &types.QueryPollAnswersRequest{},
			shouldErr: true,
		},
		{
			name:      "non existent post returns error",
			req:       &types.QueryPollAnswersRequest{PostId: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"},
			shouldErr: true,
		},
		{
			name: "non existent poll data in the post returns error",
			store: func(ctx sdk.Context) {
				post := types.Post{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Post with poll data",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.postOwner,
					PollData:             nil,
				}
				suite.k.SavePost(ctx, post)
			},
			req:       &types.QueryPollAnswersRequest{PostId: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"},
			shouldErr: true,
		},
		{
			name: "valid request returns properly",
			store: func(ctx sdk.Context) {
				post := types.Post{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Post with poll data",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.postOwner,
					PollData:             suite.testData.post.PollData,
				}
				suite.k.SavePost(ctx, post)

				answers := []types.UserAnswer{
					types.NewUserAnswer(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"cosmos1unacjuhyamzks5yu7qwlfuahdedd838e6fmdta",
						[]string{"1"},
					),
					types.NewUserAnswer(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						[]string{"1"},
					),
				}

				for _, answer := range answers {
					suite.k.SavePollAnswers(suite.ctx, answer)
				}

			},
			req:       &types.QueryPollAnswersRequest{PostId: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"},
			shouldErr: false,
			expLen:    2,
		},
		{
			name: "valid request with pagination returns properly",
			store: func(ctx sdk.Context) {
				post := types.Post{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Post with poll data",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.postOwner,
					PollData:             suite.testData.post.PollData,
				}
				suite.k.SavePost(ctx, post)

				answers := []types.UserAnswer{
					types.NewUserAnswer(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"cosmos1unacjuhyamzks5yu7qwlfuahdedd838e6fmdta",
						[]string{"1"},
					),
					types.NewUserAnswer(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						[]string{"1"},
					),
				}

				for _, answer := range answers {
					suite.k.SavePollAnswers(suite.ctx, answer)
				}

			},
			req: &types.QueryPollAnswersRequest{
				PostId:     "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Pagination: &query.PageRequest{Limit: 1},
			},
			shouldErr: false,
			expLen:    1,
		},
	}

	suite.SetupTest()
	for _, uc := range usecases {
		suite.Run(uc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if uc.store != nil {
				uc.store(ctx)
			}
			res, err := suite.k.PollAnswers(sdk.WrapSDKContext(ctx), uc.req)
			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
				suite.Require().Equal(uc.expLen, len(res.Answers))
			}
		})
	}
}
