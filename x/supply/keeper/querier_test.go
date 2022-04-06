package keeper_test

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v3/x/supply/keeper"
	"github.com/desmos-labs/desmos/v3/x/supply/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func (suite *KeeperTestSuite) TestQuerier_QueryTotalSupply() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		req         abci.RequestQuery
		path        []string
		expResponse []byte
	}{
		{
			name: "Query total supply returned correctly",
			store: func(ctx sdk.Context) {
				suite.SupplySetup(ctx, 1_000_000_000_000, 200_000, 400_000)
			},
			req: abci.RequestQuery{
				Path: fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryTotalSupply),
				Data: suite.legacyAminoCdc.MustMarshalJSON(types.NewQueryTotalSupplyRequest(suite.denom, 1_000_000)),
			},
			path:        []string{types.QueryTotalSupply},
			expResponse: suite.legacyAminoCdc.MustMarshalJSON(sdk.NewInt(1_000_000)),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			querier := keeper.NewQuerier(suite.k, suite.legacyAminoCdc)
			res, err := querier(ctx, tc.path, tc.req)

			suite.Require().NoError(err)
			suite.Require().Equal(tc.expResponse, res)
		})
	}
}

func (suite *KeeperTestSuite) TestQuerier_QueryCirculatingSupply() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		req         abci.RequestQuery
		path        []string
		expResponse []byte
	}{
		{
			name: "Query circulating supply returned correctly",
			store: func(ctx sdk.Context) {
				suite.SupplySetup(ctx, 1_000_000, 200_000, 300_000)
			},
			path: []string{types.QueryCirculatingSupply},
			req: abci.RequestQuery{
				Path: fmt.Sprintf(fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryCirculatingSupply)),
				Data: suite.legacyAminoCdc.MustMarshalJSON(types.NewQueryCirculatingSupplyRequest(suite.denom, 1_000)),
			},
			expResponse: codec.MustMarshalJSONIndent(suite.legacyAminoCdc, sdk.NewInt(500)),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			querier := keeper.NewQuerier(suite.k, suite.legacyAminoCdc)
			res, err := querier(ctx, tc.path, tc.req)

			suite.Require().NoError(err)
			suite.Require().Equal(tc.expResponse, res)
		})
	}
}
