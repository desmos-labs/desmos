package keeper_test

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestKeeper_GetTotalSupply() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		denom     string
		divider   math.Int
		expSupply math.Int
	}{
		{
			name: "total supply is computed properly",
			store: func(ctx sdk.Context) {
				suite.setupSupply(ctx,
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(1_000_000_000_000))),
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(200_000))),
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(300_000))),
				)
			},
			denom:     sdk.DefaultBondDenom,
			divider:   math.NewInt(1_000_000),
			expSupply: math.NewInt(1_000_000),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			totalConvertedSupply := suite.k.GetTotalSupply(ctx, tc.denom, tc.divider)
			suite.Require().Equal(tc.expSupply, totalConvertedSupply)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetCirculatingSupply() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		denom     string
		divider   math.Int
		expSupply math.Int
	}{
		{
			name: "circulating supply is computed properly",
			store: func(ctx sdk.Context) {
				suite.setupSupply(ctx,
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(1_000_000_000_000))),
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(200_000_000_000))),
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(300_000_000_000))),
				)
			},
			denom:     sdk.DefaultBondDenom,
			divider:   math.NewInt(1_000_000),
			expSupply: math.NewInt(500_000),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			circulatingSupply := suite.k.GetCirculatingSupply(ctx, tc.denom, tc.divider)
			suite.Require().Equal(tc.expSupply, circulatingSupply)
		})
	}
}
