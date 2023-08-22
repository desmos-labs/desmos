package keeper_test

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/x/supply/types"
)

func (suite *KeeperTestSuite) TestQueryServer_Total() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryTotalRequest
		expSupply math.Int
	}{
		{
			name: "valid query returns properly",
			store: func(ctx sdk.Context) {
				suite.setupSupply(ctx,
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1_000_000_000_000))),
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(200_000))),
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(300_000))),
				)
			},
			req:       types.NewQueryTotalRequest(sdk.DefaultBondDenom, 6),
			expSupply: sdk.NewInt(1_000_000),
		},
		{
			name: "valid query returns properly - divider equals to 0",
			store: func(ctx sdk.Context) {
				suite.setupSupply(ctx,
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1_000_000_000_000))),
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(200_000))),
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(300_000))),
				)
			},
			req:       types.NewQueryTotalRequest(sdk.DefaultBondDenom, 0),
			expSupply: sdk.NewInt(1_000_000_000_000),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			res, err := suite.k.Total(sdk.WrapSDKContext(ctx), tc.req)
			suite.Require().NoError(err)
			suite.Require().Equal(tc.expSupply, res.TotalSupply)
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_Circulating() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryCirculatingRequest
		expSupply math.Int
	}{
		{
			name: "valid query returns properly",
			store: func(ctx sdk.Context) {
				suite.setupSupply(ctx,
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1_000_000))),
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(200_000))),
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(300_000))),
				)
			},
			req:       types.NewQueryCirculatingRequest(sdk.DefaultBondDenom, 3),
			expSupply: sdk.NewInt(500),
		},
		{
			name: "valid query returns properly - divider equals to 0",
			store: func(ctx sdk.Context) {
				suite.setupSupply(ctx,
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1_000_000))),
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(200_000))),
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(300_000))),
				)
			},
			req:       types.NewQueryCirculatingRequest(sdk.DefaultBondDenom, 0),
			expSupply: sdk.NewInt(500_000),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			res, err := suite.k.Circulating(sdk.WrapSDKContext(ctx), tc.req)
			suite.Require().NoError(err)
			suite.Require().Equal(tc.expSupply, res.CirculatingSupply)
		})
	}
}
