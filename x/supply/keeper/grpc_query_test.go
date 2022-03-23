package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v3/x/supply/types"
)

func (suite *KeeperTestSuite) TestQueryServer_TotalSupply() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		req         *types.QueryTotalSupplyRequest
		expResponse *types.QueryTotalSupplyResponse
	}{
		{
			name: "valid query returns properly",
			store: func(ctx sdk.Context) {
				suite.SupplySetup(1_000_000_000_000, 200_000, 300_000)
			},
			req:         types.NewQueryTotalSupplyRequest(suite.denom, 1_000_000),
			expResponse: &types.QueryTotalSupplyResponse{TotalSupply: sdk.NewInt(1_000_000)},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			res, err := suite.k.TotalSupply(sdk.WrapSDKContext(ctx), tc.req)
			suite.Require().NoError(err)
			suite.Require().Equal(tc.expResponse, res)
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_CirculatingSupply() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		req         *types.QueryCirculatingSupplyRequest
		expResponse *types.QueryCirculatingSupplyResponse
	}{
		{
			name: "valid query returns properly",
			store: func(ctx sdk.Context) {
				suite.SupplySetup(1_000_000, 200_000, 300_000)
			},
			req:         types.NewQueryCirculatingSupplyRequest(suite.denom, 1_000),
			expResponse: &types.QueryCirculatingSupplyResponse{CirculatingSupply: sdk.NewInt(500)},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			res, err := suite.k.CirculatingSupply(sdk.WrapSDKContext(ctx), tc.req)
			suite.Require().NoError(err)
			suite.Require().Equal(tc.expResponse, res)
		})
	}
}
