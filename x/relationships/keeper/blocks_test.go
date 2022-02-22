package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v2/x/relationships/types"
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
				blocks := suite.k.GetUserBlocks(ctx, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				suite.Require().Len(blocks, 1)
				suite.Require().Equal(types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"reason",
					0,
				), blocks[0])
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

func (suite *KeeperTestSuite) TestKeeper_IsUserBlocked() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		blocker    string
		blocked    string
		subspace   uint64
		expBlocked bool
	}{
		{
			name:       "non blocked user returns false",
			blocker:    "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			blocked:    "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			expBlocked: false,
		},
		{
			name: "blocked user returns true with generic subspace",
			store: func(ctx sdk.Context) {
				block := types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"test",
					1,
				)
				suite.k.SaveUserBlock(ctx, block)
			},
			blocker:    "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			blocked:    "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			subspace:   0,
			expBlocked: true,
		},
		{
			name: "blocked user returns true with specific subspace",
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
			subspace:   1,
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

			res := suite.k.IsUserBlocked(ctx, tc.blocker, tc.blocked, 0)
			suite.Equal(tc.expBlocked, res)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetUserBlocks() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		user      string
		expBlocks []types.UserBlock
	}{
		{
			name:      "empty slice is returned properly",
			user:      "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			expBlocks: nil,
		},
		{
			name: "non empty slice is returned properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserBlock(ctx, types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					"reason",
					1,
				))
			},
			user: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			expBlocks: []types.UserBlock{
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					"reason",
					1,
				),
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

			blocks := suite.k.GetUserBlocks(ctx, tc.user)
			suite.Require().Equal(tc.expBlocks, blocks)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetAllUsersBlocks() {
	testCases := []struct {
		name           string
		store          func(ctx sdk.Context)
		expUsersBlocks []types.UserBlock
	}{
		{
			name:           "empty users blocks list",
			expUsersBlocks: nil,
		},
		{
			name: "non-empty users blocks list",
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
			expUsersBlocks: []types.UserBlock{
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"reason",
					0,
				),
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
					"reason",
					0,
				),
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

			blocks := suite.k.GetAllUsersBlocks(ctx)
			suite.Require().Equal(tc.expUsersBlocks, blocks)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteUserBlock() {
	testCases := []struct {
		name     string
		store    func(ctx sdk.Context)
		blocker  string
		blocked  string
		subspace uint64
		check    func(ctx sdk.Context)
	}{
		{
			name: "delete user block with len(stored) > 1",
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
			blocker:  "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			blocked:  "cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
			subspace: 0,
			check: func(ctx sdk.Context) {
				blocks := suite.k.GetUserBlocks(ctx, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				suite.Require().Len(blocks, 1)
			},
		},
		{
			name: "delete user block with len(stored) == 1",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserBlock(ctx, types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"reason",
					0,
				))
			},
			blocker:  "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			blocked:  "cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
			subspace: 0,
		},
		{
			name:     "deleting a user block that does not exist returns no error",
			blocker:  "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			blocked:  "blocked",
			subspace: 0,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.k.DeleteUserBlock(ctx, tc.blocker, tc.blocked, tc.subspace)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteAllUserBlocks() {
	testCases := []struct {
		name    string
		store   func(ctx sdk.Context)
		blocker string
		check   func(ctx sdk.Context)
	}{
		{
			name: "all blocks are deleted properly",
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
			blocker: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			check: func(ctx sdk.Context) {
				blocks := suite.k.GetUserBlocks(ctx, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				suite.Require().Empty(blocks)
			},
		},
		{
			name: "delete user block with len(stored) == 1",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserBlock(ctx, types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"reason",
					0,
				))
			},
			blocker: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.k.DeleteAllUserBlocks(ctx, tc.blocker)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
