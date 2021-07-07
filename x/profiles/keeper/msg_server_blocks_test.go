package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/testutil"

	"github.com/desmos-labs/desmos/x/profiles/keeper"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestMsgServer_BlockUser() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgBlockUser
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "existing block returns error",
			store: func(ctx sdk.Context) {
				block := types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"reason",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocker)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocked)))
				suite.Require().NoError(suite.k.SaveUserBlock(ctx, block))
			},
			msg: types.NewMsgBlockUser(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"reason",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			shouldErr: true,
		},
		{
			name: "non existing block is stored correctly",
			store: func(ctx sdk.Context) {
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")))
			},
			msg: types.NewMsgBlockUser(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"reason",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeBlockUser,
					sdk.NewAttribute(types.AttributeKeyUserBlockBlocker, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeKeyUserBlockBlocked, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
					sdk.NewAttribute(types.AttributeKeySubspace, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
					sdk.NewAttribute(types.AttributeKeyUserBlockReason, "reason"),
				),
			},
			check: func(ctx sdk.Context) {
				expected := []types.UserBlock{
					types.NewUserBlock(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				}
				suite.Require().Equal(expected, suite.k.GetAllUsersBlocks(ctx))
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

			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.BlockUser(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.shouldErr {
				suite.Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServer_UnblockUser() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgUnblockUser
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "invalid block returns error",
			msg: types.NewMsgUnblockUser(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"subspace",
			),
			shouldErr: true,
		},
		{
			name: "existing block is removed properly",
			store: func(ctx sdk.Context) {
				block := types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"reason",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocker)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocked)))
				suite.Require().NoError(suite.k.SaveUserBlock(ctx, block))
			},
			msg: types.NewMsgUnblockUser(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeUnblockUser,
					sdk.NewAttribute(types.AttributeKeyUserBlockBlocker, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeKeyUserBlockBlocked, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
					sdk.NewAttribute(types.AttributeKeySubspace, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				),
			},
			check: func(ctx sdk.Context) {
				suite.Require().Empty(suite.k.GetAllUsersBlocks(ctx))
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

			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.UnblockUser(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}
