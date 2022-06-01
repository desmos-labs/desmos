package keeper_test

import (
	"time"

	poststypes "github.com/desmos-labs/desmos/v3/x/posts/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/reports/types"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

func (suite *KeeperTestsuite) TestKeeper_AfterSubspaceSaved() {
	testCases := []struct {
		name     string
		store    func(ctx sdk.Context)
		subspace subspacestypes.Subspace
		check    func(ctx sdk.Context)
	}{
		{
			name: "saving a subspaces adds the correct keys",
			subspace: subspacestypes.NewSubspace(
				1,
				"Test subspace",
				"This is a test subspace",
				"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
				"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
				"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
				time.Date(2020, 1, 2, 12, 00, 00, 000, time.UTC),
			),
			check: func(ctx sdk.Context) {
				storedReasonID, err := suite.k.GetNextReasonID(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint32(1), storedReasonID)

				storedReportID, err := suite.k.GetNextReportID(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint64(1), storedReportID)
			},
		},
	}

	// Set the hooks
	suite.sk.SetHooks(suite.k.Hooks())

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.sk.SaveSubspace(ctx, tc.subspace)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_AfterSubspaceDeleted() {
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
					1,
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

	// Set the hooks
	suite.sk.SetHooks(suite.k.Hooks())

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.sk.DeleteSubspace(ctx, tc.subspaceID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_AfterPostDeleted() {
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
				suite.pk.SavePost(ctx, poststypes.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					poststypes.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					1,
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

	// Set the hooks
	suite.pk.SetHooks(suite.k.Hooks())

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.pk.DeletePost(ctx, tc.subspaceID, tc.postID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
