package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

func (suite *KeeperTestsuite) TestKeeper_AfterSubspaceSaved() {
	testCases := []struct {
		name       string
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
			suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
				tc.subspaceID,
				"Test",
				"Testing subspace",
				"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
				"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
				"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			))

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
