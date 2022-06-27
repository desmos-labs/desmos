package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/reactions/types"
)

func (suite *KeeperTestSuite) TestKeeper_SaveSubspaceReactionsParams() {
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
				params, err := suite.k.GetSubspaceReactionsParams(ctx, 1)
				suite.Require().NoError(err)
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
				suite.k.SaveSubspaceReactionsParams(ctx, types.NewSubspaceReactionsParams(
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
				params, err := suite.k.GetSubspaceReactionsParams(ctx, 1)
				suite.Require().NoError(err)
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

			suite.k.SaveSubspaceReactionsParams(ctx, tc.params)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_HasSubspaceReactionsParams() {
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
				suite.k.SaveSubspaceReactionsParams(ctx, types.NewSubspaceReactionsParams(
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

			result := suite.k.HasSubspaceReactionsParams(ctx, tc.subspaceID)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetSubspaceReactionsParams() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		shouldErr  bool
		expParams  types.SubspaceReactionsParams
	}{
		{
			name:       "non existing params returns error",
			subspaceID: 1,
			shouldErr:  true,
			expParams:  types.DefaultReactionsParams(1),
		},
		{
			name: "existing params returns correct value",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspaceReactionsParams(ctx, types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(true, 100, ""),
				))
			},
			subspaceID: 1,
			shouldErr:  false,
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

			params, err := suite.k.GetSubspaceReactionsParams(ctx, tc.subspaceID)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expParams, params)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteSubspaceReactionsParams() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing subspace reactions params are deleted properly",
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasSubspaceReactionsParams(ctx, 1))
			},
		},
		{
			name: "existing subspace reactions params are deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspaceReactionsParams(ctx, types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(true, 10, ""),
				))
			},
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasSubspaceReactionsParams(ctx, 1))
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

			suite.k.DeleteSubspaceReactionsParams(ctx, tc.subspaceID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
