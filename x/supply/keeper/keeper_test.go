package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestKeeper_CalculateCirculatingSupply() {
	testCases := []struct {
		name                      string
		multiplierFactor          sdk.Int
		expectedCirculatingSupply sdk.Int
		store                     func(ctx sdk.Context)
	}{
		{
			name:                      "circulating supply calculated correctly",
			multiplierFactor:          sdk.NewInt(1_000_000),
			expectedCirculatingSupply: sdk.NewInt(500_000),
			store: func(ctx sdk.Context) {
				suite.SupplySetup(1_000_000_000_000, 200_000_000_000, 300_000_000_000)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			circulatingSupply := suite.k.CalculateCirculatingSupply(ctx, suite.denom, tc.multiplierFactor)
			suite.Require().Equal(tc.expectedCirculatingSupply, circulatingSupply)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetConvertedTotalSupply() {
	testCases := []struct {
		name                string
		multiplierFactor    sdk.Int
		expectedTotalSupply sdk.Int
		store               func(ctx sdk.Context)
	}{
		{
			name:                "total converted supply returned correctly",
			multiplierFactor:    sdk.NewInt(1_000_000),
			expectedTotalSupply: sdk.NewInt(1_000_000),
			store: func(ctx sdk.Context) {
				suite.SupplySetup(1_000_000_000_000, 0, 0)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			totalConvertedSupply := suite.k.GetConvertedTotalSupply(ctx, suite.denom, tc.multiplierFactor)
			suite.Require().Equal(tc.expectedTotalSupply, totalConvertedSupply)
		})
	}
}
