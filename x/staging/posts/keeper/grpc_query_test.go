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
