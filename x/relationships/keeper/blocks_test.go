package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/relationships/types"
)

func (suite *KeeperTestSuite) TestKeeper_SaveUserBlock() {
	testCases := []struct {
		name      string
		userBlock types.UserBlock
		check     func(ctx sdk.Context)
	}{
		{
			name: "user block is saved correctly",
			userBlock: types.NewUserBlock(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				"reason",
				0,
			),
			check: func(ctx sdk.Context) {
				suite.Require().True(suite.k.HasUserBlocked(ctx,
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					0,
				))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()

			suite.k.SaveUserBlock(ctx, tc.userBlock)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_HasUserBlocked() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		blocker    string
		blocked    string
		subspaceID uint64
		expBlocked bool
	}{
		{
			name:       "non blocked user returns false",
			blocker:    "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			blocked:    "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			expBlocked: false,
		},
		{
			name: "blocked user returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserBlock(ctx, types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"test",
					1,
				))
			},
			blocker:    "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			blocked:    "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			subspaceID: 1,
			expBlocked: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			res := suite.k.HasUserBlocked(ctx, tc.blocker, tc.blocked, tc.subspaceID)
			suite.Equal(tc.expBlocked, res)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetUserBlock() {
	testCases := []struct {
		name         string
		store        func(ctx sdk.Context)
		blocker      string
		blocked      string
		subspaceID   uint64
		expFound     bool
		expUserBlock types.UserBlock
	}{
		{
			name:       "not found user block returns false",
			blocker:    "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			blocked:    "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			subspaceID: 0,
			expFound:   false,
		},
		{
			name: "found user block returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserBlock(ctx, types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"",
					1,
				))
			},
			blocker:    "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			blocked:    "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			subspaceID: 1,
			expFound:   true,
			expUserBlock: types.NewUserBlock(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				1,
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

			userBlock, found := suite.k.GetUserBlock(ctx, tc.blocker, tc.blocked, tc.subspaceID)
			suite.Require().Equal(tc.expFound, found)
			if tc.expFound {
				suite.Require().Equal(tc.expUserBlock, userBlock)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteUserBlock() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		blocker    string
		blocked    string
		subspaceID uint64
		check      func(ctx sdk.Context)
	}{
		{
			name: "deleting user block works properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserBlock(ctx, types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"reason",
					0,
				))
				suite.k.SaveUserBlock(ctx, types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
					"reason",
					0,
				))
			},
			blocker:    "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			blocked:    "cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
			subspaceID: 0,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasRelationship(ctx,
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
					0,
				))
				suite.Require().True(suite.k.HasUserBlocked(ctx,
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					0,
				))
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

			suite.k.DeleteUserBlock(ctx, tc.blocker, tc.blocked, tc.subspaceID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
