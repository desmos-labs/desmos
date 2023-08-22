package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/x/reports/types"
)

func (suite *KeeperTestSuite) TestKeeper_AfterSubspaceSaved() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		check      func(ctx sdk.Context)
	}{
		{
			name:       "saving a subspaces adds the correct keys",
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				storedReasonID, err := suite.k.GetNextReasonID(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint32(1), storedReasonID)

				storedReportID, err := suite.k.GetNextReportID(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint64(1), storedReportID)
			},
		},
		{
			name: "reason and report ids are not overwritten",
			store: func(ctx sdk.Context) {
				suite.k.SetNextReportID(ctx, 1, 2)
				suite.k.SetNextReasonID(ctx, 1, 2)
			},
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				storedReasonID, err := suite.k.GetNextReasonID(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint32(2), storedReasonID)

				storedReportID, err := suite.k.GetNextReportID(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint64(2), storedReportID)
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

			suite.k.Hooks().AfterSubspaceSaved(ctx, tc.subspaceID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_AfterSubspaceDeleted() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		check      func(ctx sdk.Context)
	}{
		{
			name: "deleting a subspace removes all the reasons and reports",
			store: func(ctx sdk.Context) {
				suite.k.SetNextReasonID(ctx, 1, 2)
				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				))

				suite.k.SetNextReportID(ctx, 1, 2)
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"",
					types.NewPostTarget(1),
					"cosmos1atdl3cpms89md5qa3rxtql0drtgftch2zgkr7v",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)

				// Check the reasons data
				suite.Require().False(store.Has(types.NextReasonIDStoreKey(1)))
				suite.Require().False(store.Has(types.ReasonStoreKey(1, 1)))

				// Check the reports data
				suite.Require().False(store.Has(types.NextReportIDStoreKey(1)))
				suite.Require().False(store.Has(types.ReportStoreKey(1, 1)))
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

			suite.k.Hooks().AfterSubspaceDeleted(ctx, tc.subspaceID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_AfterPostDeleted() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		check      func(ctx sdk.Context)
	}{
		{
			name: "deleting a post removes all the associated reports",
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"",
					types.NewPostTarget(1),
					"cosmos1atdl3cpms89md5qa3rxtql0drtgftch2zgkr7v",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			subspaceID: 1,
			postID:     1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)

				// Make sure the report does not exist
				suite.Require().False(store.Has(types.ReportStoreKey(1, 1)))
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

			suite.k.Hooks().AfterPostDeleted(ctx, tc.subspaceID, tc.postID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
