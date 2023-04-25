package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/posts/types"
)

func (suite *KeeperTestSuite) TestKeeper_AfterSubspaceSaved() {
	testCases := []struct {
		name       string
		setup      func()
		store      func(ctx sdk.Context)
		subspaceID uint64
		check      func(ctx sdk.Context)
	}{
		{
			name:       "post id is set properly",
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				stored, err := suite.k.GetNextPostID(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint64(1), stored)
			},
		},
		{
			name: "post id is not overwritten",
			store: func(ctx sdk.Context) {
				suite.k.SetNextPostID(ctx, 1, 2)
			},
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				stored, err := suite.k.GetNextPostID(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint64(2), stored)
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
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
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
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
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

func (suite *KeeperTestSuite) TestKeeper_AfterSectionDeleted() {
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
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
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
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
				))
			},
			subspaceID: 1,
			sectionID:  1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasPost(ctx, 1, 1))
				suite.Require().False(suite.k.HasPost(ctx, 1, 2))
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

			// Call the method that should call the hook
			suite.k.Hooks().AfterSubspaceSectionDeleted(ctx, tc.subspaceID, tc.sectionID)

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
