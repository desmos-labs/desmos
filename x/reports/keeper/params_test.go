package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/reports/types"
)

func (suite *KeeperTestsuite) TestKeeper_SetParams() {
	testCases := []struct {
		name   string
		store  func(ctx sdk.Context)
		params types.Params
		check  func(ctx sdk.Context)
	}{
		{
			name:   "default params are saved correctly",
			params: types.DefaultParams(),
			check: func(ctx sdk.Context) {
				stored := suite.k.GetParams(ctx)
				suite.Require().Equal(types.DefaultParams(), stored)
			},
		},
		{
			name: "params are overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.DefaultParams())
			},
			params: types.NewParams(types.NewStandardReasons(
				types.NewStandardReason(1, "Pornography", "This content contains pornography"),
			)),
			check: func(ctx sdk.Context) {
				store := suite.k.GetParams(ctx)
				suite.Require().Equal(types.NewParams(types.NewStandardReasons(
					types.NewStandardReason(1, "Pornography", "This content contains pornography"),
				)), store)
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

			suite.k.SetParams(ctx, tc.params)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetParams() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		expParams types.Params
	}{
		{
			name: "default params are returned properly",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.DefaultParams())
			},
			expParams: types.DefaultParams(),
		},
		{
			name: "custom params are returned properly",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.NewParams(types.NewStandardReasons(
					types.NewStandardReason(1, "Pornography", "This content contains pornography"),
				)))
			},
			expParams: types.NewParams(types.NewStandardReasons(
				types.NewStandardReason(1, "Pornography", "This content contains pornography"),
			)),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			params := suite.k.GetParams(ctx)
			suite.Require().True(tc.expParams.Equal(params))
		})
	}
}
