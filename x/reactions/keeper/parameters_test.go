package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/reactions/types"
)

func (suite *KeeperTestSuite) TestKeeper_SaveSubspaceParams() {
	testCases := []struct {
		name   string
		store  func(ctx sdk.Context)
		params types.SubspaceReactionsParams
		check  func(ctx sdk.Context)
	}{
		{
			name: "non existing params are stored properly",
			params: types.NewSubspaceReactionsParams(
				1,
				types.NewRegisteredReactionValueParams(true),
				types.NewFreeTextValueParams(true, 100, ""),
			),
			check: func(ctx sdk.Context) {
				params := suite.k.GetSubspaceReactionsParams(ctx, 1)
				suite.Require().Equal(types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(true, 100, ""),
				), params)
			},
		},
		{
			name: "existing params are overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspaceParams(ctx, types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(true, 10, ""),
				))
			},
			params: types.NewSubspaceReactionsParams(
				1,
				types.NewRegisteredReactionValueParams(true),
				types.NewFreeTextValueParams(true, 100, "[a-zA-Z]"),
			),
			check: func(ctx sdk.Context) {
				params := suite.k.GetSubspaceReactionsParams(ctx, 1)
				suite.Require().Equal(types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(true, 100, "[a-zA-Z]"),
				), params)
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

			suite.k.SaveSubspaceParams(ctx, tc.params)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_HasSubspaceParams() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		expResult  bool
	}{
		{
			name:       "non existing params returns false",
			subspaceID: 1,
			expResult:  false,
		},
		{
			name: "existing params returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspaceParams(ctx, types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(true, 100, ""),
				))
			},
			subspaceID: 1,
			expResult:  true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			result := suite.k.HasSubspaceParams(ctx, tc.subspaceID)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetSubspaceReactionsParams() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		expParams  types.SubspaceReactionsParams
	}{
		{
			name:       "non existing params returns default params",
			subspaceID: 1,
			expParams:  types.DefaultReactionsParams(1),
		},
		{
			name: "existing params returns correct value",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspaceParams(ctx, types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(true, 100, ""),
				))
			},
			subspaceID: 1,
			expParams: types.NewSubspaceReactionsParams(
				1,
				types.NewRegisteredReactionValueParams(true),
				types.NewFreeTextValueParams(true, 100, ""),
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			params := suite.k.GetSubspaceReactionsParams(ctx, tc.subspaceID)
			suite.Require().Equal(tc.expParams, params)
		})
	}
}
