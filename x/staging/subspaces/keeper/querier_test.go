package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/desmos-labs/desmos/x/staging/subspaces/keeper"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"time"
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
			name:           "Subspace not found returns error",
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
				suite.k.SaveSubspace(suite.ctx, *test.storedSubspace)
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
