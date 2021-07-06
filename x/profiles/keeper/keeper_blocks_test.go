package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/testutil"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestKeeper_IsUserBlocked() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		blocker    string
		blocked    string
		expBlocked bool
	}{
		{
			name: "blocked user found returns true",
			store: func(ctx sdk.Context) {
				block := types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"test",
					"",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocker)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocked)))
				suite.Require().NoError(suite.k.SaveUserBlock(ctx, block))
			},
			blocker:    "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			blocked:    "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			expBlocked: true,
		},
		{
			name:       "non blocked user not found returns false",
			blocker:    "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			blocked:    "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			expBlocked: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			res := suite.k.IsUserBlocked(ctx, tc.blocker, tc.blocked)
			suite.Equal(tc.expBlocked, res)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_SaveUserBlock() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		userBlock types.UserBlock
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "already blocked user returns error",
			store: func(ctx sdk.Context) {
				block := types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					"reason",
					"subspace",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocker)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocked)))
				suite.Require().NoError(suite.k.SaveUserBlock(ctx, block))
			},
			userBlock: types.NewUserBlock(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				"reason",
				"subspace",
			),
			shouldErr: true,
		},
		{
			name: "user block added correctly",
			store: func(ctx sdk.Context) {
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")))
			},
			userBlock: types.NewUserBlock(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				"reason",
				"subspace",
			),
			shouldErr: false,
			check: func(ctx sdk.Context) {
				blocks := suite.k.GetUserBlocks(ctx, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				suite.Require().Len(blocks, 1)
				suite.Require().Equal(types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"reason",
					"subspace",
				), blocks[0])
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

			err := suite.k.SaveUserBlock(ctx, tc.userBlock)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteUserBlock() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		blocker   string
		blocked   string
		subspace  string
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "delete user block with len(stored) > 1",
			store: func(ctx sdk.Context) {
				block := types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"reason",
					"subspace",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocker)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocked)))
				suite.Require().NoError(suite.k.SaveUserBlock(ctx, block))

				block = types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
					"reason",
					"subspace",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocker)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocked)))
				suite.Require().NoError(suite.k.SaveUserBlock(ctx, block))
			},
			blocker:   "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			blocked:   "cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
			subspace:  "subspace",
			shouldErr: false,
			check: func(ctx sdk.Context) {
				blocks := suite.k.GetUserBlocks(ctx, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				suite.Require().Len(blocks, 1)
			},
		},
		{
			name: "delete user block with len(stored) == 1",
			store: func(ctx sdk.Context) {
				block := types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"reason",
					"subspace",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocker)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocked)))
				suite.Require().NoError(suite.k.SaveUserBlock(ctx, block))
			},
			blocker:   "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			blocked:   "cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
			subspace:  "subspace",
			shouldErr: false,
		},
		{
			name:      "deleting a user block that does not exist returns an error",
			blocker:   "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			blocked:   "blocked",
			subspace:  "subspace",
			shouldErr: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			err := suite.k.DeleteUserBlock(ctx, tc.blocker, tc.blocked, tc.subspace)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				if tc.check != nil {
					tc.check(ctx)
				}
			}
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
			name: "non empty slice is returned properly",
			store: func(ctx sdk.Context) {
				block := types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					"reason",
					"subspace",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocker)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocked)))
				suite.Require().NoError(suite.k.SaveUserBlock(ctx, block))
			},
			user: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			expBlocks: []types.UserBlock{
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					"reason",
					"subspace",
				),
			},
		},
		{
			name:      "empty slice is returned properly",
			user:      "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			expBlocks: nil,
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
			name: "non-empty users blocks list",
			store: func(ctx sdk.Context) {
				block := types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"reason",
					"subspace",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocker)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocked)))
				suite.Require().NoError(suite.k.SaveUserBlock(ctx, block))

				block = types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
					"reason",
					"subspace",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocker)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocked)))
				suite.Require().NoError(suite.k.SaveUserBlock(ctx, block))
			},
			expUsersBlocks: []types.UserBlock{
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"reason",
					"subspace",
				),
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
					"reason",
					"subspace",
				),
			},
		},
		{
			name:           "empty users blocks list",
			expUsersBlocks: nil,
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

func (suite *KeeperTestSuite) TestKeeper_HasUserBlocked() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		blocker    string
		blocked    string
		subspace   string
		expBlocked bool
	}{
		{
			name: "blocked user found returns true",
			store: func(ctx sdk.Context) {
				block := types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"reason",
					"subspace",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocker)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocked)))
				suite.Require().NoError(suite.k.SaveUserBlock(ctx, block))
			},
			blocker:    "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			blocked:    "cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
			subspace:   "subspace",
			expBlocked: true,
		},
		{
			name:       "blocked user not found returns false",
			blocker:    "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			blocked:    "blocked",
			subspace:   "subspace_2",
			expBlocked: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			blocked := suite.k.HasUserBlocked(ctx, tc.blocker, tc.blocked, tc.subspace)
			suite.Equal(tc.expBlocked, blocked)
		})
	}
}
