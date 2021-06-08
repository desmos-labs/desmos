package keeper_test

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/desmos-labs/desmos/x/staging/subspaces/keeper"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
)

func (suite *KeeperTestsuite) Test_querySubspace() {
	tests := []struct {
		name           string
		path           []string
		storedSubspace *types.Subspace
		expErr         bool
		expResult      types.QuerySubspaceResponse
	}{
		{
			name:           "Invalid query endpoint",
			path:           []string{"invalid"},
			storedSubspace: nil,
			expErr:         true,
		},
		{
			name:           "Invalid subspace ID returns error",
			path:           []string{types.QuerySubspace, ""},
			storedSubspace: nil,
			expErr:         true,
		},
		{
			name:           "Subspace not expFound returns error",
			path:           []string{types.QuerySubspace, "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"},
			storedSubspace: nil,
			expErr:         true,
		},
		{
			name: "Subspace returned correctly",
			path: []string{types.QuerySubspace, "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"},
			storedSubspace: &types.Subspace{
				ID:           "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Name:         "test",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				CreationTime: time.Time{},
			},
			expErr: false,
			expResult: types.QuerySubspaceResponse{
				Subspace: types.Subspace{
					ID:           "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Name:         "test",
					Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					CreationTime: time.Time{},
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			if test.storedSubspace != nil {
				_ = suite.k.SaveSubspace(suite.ctx, *test.storedSubspace, test.storedSubspace.Owner)
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
			name: "Return all the subspaces correctly",
			path: []string{types.QuerySubspaces},
			storedSubspaces: []types.Subspace{
				{
					ID:           "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Name:         "test",
					Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					CreationTime: time.Time{},
				},
				{
					ID:           "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					Name:         "test",
					Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					CreationTime: time.Time{},
				},
			},
			expErr: false,
			expResult: types.QuerySubspacesResponse{
				Subspaces: []types.Subspace{
					{
						ID:           "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						Name:         "test",
						Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						CreationTime: time.Time{},
					},
					{
						ID:           "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
						Name:         "test",
						Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						CreationTime: time.Time{},
					},
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
