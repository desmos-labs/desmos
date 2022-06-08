package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
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
			name: "post id is set properly",
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
				stored, err := suite.k.GetNextPostID(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint64(1), stored)
			},
		},
		{
			name: "post id is not overwritten",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					time.Date(2020, 1, 2, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SetNextPostID(ctx, 1, 2)
			},
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
				stored, err := suite.k.GetNextPostID(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint64(2), stored)
			},
		},
	}

	// Set the subspaces hooks
	suite.sk.SetHooks(suite.k.Hooks())

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			// Call the method that should call the hook
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
			name: "subspace data are deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SetNextPostID(ctx, 1, 1)
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					1,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					2,
					"External ID",
					"This is a text",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					1,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				_, err := suite.k.GetNextPostID(ctx, 1)
				suite.Require().Error(err)

				suite.Require().False(suite.k.HasPost(ctx, 1, 1))
				suite.Require().False(suite.k.HasPost(ctx, 1, 2))
			},
		},
	}

	// Set the subspaces hooks
	suite.sk.SetHooks(suite.k.Hooks())

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			// Call the method that should call the hook
			suite.sk.DeleteSubspace(ctx, tc.subspaceID)

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_AfterSectionDeleted() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		sectionID  uint32
		check      func(ctx sdk.Context)
	}{
		{
			name: "section posts are deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SetNextPostID(ctx, 1, 1)
				suite.sk.SaveSection(ctx, subspacestypes.NewSection(1, 1, 0, "test", ""))
				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					1,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					2,
					"External ID",
					"This is a text",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					1,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			subspaceID: 1,
			sectionID:  1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasPost(ctx, 1, 1))
				suite.Require().False(suite.k.HasPost(ctx, 1, 2))
			},
		},
		{
			name: "section permissions are deleted properly",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx,
					subspacestypes.NewSubspace(
						1,
						"test",
						"test",
						"cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
						"cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
						"cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
						time.Now(),
					))
				suite.sk.SaveSection(ctx, subspacestypes.NewSection(1, 1, 0, "test", ""))
				suite.sk.SetUserPermissions(ctx, 1, 1, "cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4", subspacestypes.PermissionWrite)
				suite.Require().True(
					suite.sk.HasPermission(ctx, 1, 1, "cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4", subspacestypes.PermissionWrite),
				)
			},
			subspaceID: 1,
			sectionID:  1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.sk.HasPermission(ctx, 1, 1, "cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4", subspacestypes.PermissionWrite))
			},
		},
	}

	// Set the subspaces hooks
	suite.sk.SetHooks(suite.k.Hooks())

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			// Call the method that should call the hook
			suite.sk.DeleteSection(ctx, tc.subspaceID, tc.sectionID)

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
