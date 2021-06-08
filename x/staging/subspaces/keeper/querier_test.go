package keeper_test

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/desmos-labs/desmos/x/staging/subspaces/keeper"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
)

func (suite *KeeperTestsuite) Test_querySubspace() {
	tests := []struct {
		name      string
		path      []string
		store     func(ctx sdk.Context)
		expErr    bool
		expResult types.QuerySubspaceResponse
	}{
		{
			name:   "Invalid query endpoint",
			path:   []string{"invalid"},
			expErr: true,
		},
		{
			name:   "Invalid subspace ID returns error",
			path:   []string{types.QuerySubspace, ""},
			expErr: true,
		},
		{
			name:   "Subspace not expFound returns error",
			path:   []string{types.QuerySubspace, "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"},
			expErr: true,
		},
		{
			name: "Subspace returned correctly",
			path: []string{types.QuerySubspace, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"},
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)
			},
			expErr: false,
			expResult: types.QuerySubspaceResponse{
				Subspace: types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
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

			querier := keeper.NewQuerier(suite.k, suite.legacyAminoCdc)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if test.expErr {
				suite.Error(err)
			} else {
				suite.NoError(err)
				expected := codec.MustMarshalJSONIndent(suite.legacyAminoCdc, &test.expResult)
				suite.Equal(string(expected), string(result))
			}
		})
	}
}

func (suite *KeeperTestsuite) Test_querySubspaces() {
	tests := []struct {
		name            string
		path            []string
		storedSubspaces []types.Subspace
		expErr          bool
		expResult       types.QuerySubspacesResponse
	}{
		{
			name: "Returns all the subspaces correctly",
			path: []string{types.QuerySubspaces},
			storedSubspaces: []types.Subspace{
				types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				),
				types.NewSubspace(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				),
			},
			expErr: false,
			expResult: types.QuerySubspacesResponse{
				Subspaces: []types.Subspace{
					types.NewSubspace(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"test",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						types.SubspaceTypeOpen,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					),
					types.NewSubspace(
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
						"test",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						types.SubspaceTypeOpen,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, subspace := range test.storedSubspaces {
				_ = suite.k.SaveSubspace(suite.ctx, subspace, subspace.Owner)
			}

			querier := keeper.NewQuerier(suite.k, suite.legacyAminoCdc)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if test.expErr {
				suite.Error(err)
			} else {
				suite.NoError(err)
				expected := codec.MustMarshalJSONIndent(suite.legacyAminoCdc, &test.expResult)
				suite.Equal(string(expected), string(result))
			}
		})
	}
}
