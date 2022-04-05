package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v3/x/supply/types"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (suite *KeeperTestSuite) TestQueryServer_TotalSupply() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		req         *types.QueryTotalSupplyRequest
		expResponse *wrapperspb.StringValue
	}{
		{
			name: "valid query returns properly",
			store: func(ctx sdk.Context) {
				suite.SupplySetup(ctx, 1_000_000_000_000, 200_000, 300_000)
			},
			req:         types.NewQueryTotalSupplyRequest(suite.denom, 6),
			expResponse: wrapperspb.String(sdk.NewInt(1_000_000).String()),
		},
		{
			name: "valid query returns properly without divider_exponent set",
			store: func(ctx sdk.Context) {
				suite.SupplySetup(ctx, 1_000_000_000_000, 200_000, 300_000)
			},
			req:         types.NewQueryTotalSupplyRequest(suite.denom, 0),
			expResponse: wrapperspb.String(sdk.NewInt(1_000_000_000_000).String()),
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
		expResponse *wrapperspb.StringValue
	}{
		{
			name: "valid query returns properly",
			store: func(ctx sdk.Context) {
				suite.SupplySetup(ctx, 1_000_000, 200_000, 300_000)
			},
			req:         types.NewQueryCirculatingSupplyRequest(suite.denom, 3),
			expResponse: wrapperspb.String(sdk.NewInt(500).String()),
		},
		{
			name: "valid query returns properly without divider_exponent",
			store: func(ctx sdk.Context) {
				suite.SupplySetup(ctx, 1_000_000, 200_000, 300_000)
			},
			req:         types.NewQueryCirculatingSupplyRequest(suite.denom, 0),
			expResponse: wrapperspb.String(sdk.NewInt(500_000).String()),
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
