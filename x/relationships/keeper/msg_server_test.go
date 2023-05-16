package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"

	"github.com/desmos-labs/desmos/v5/x/relationships/keeper"
	"github.com/desmos-labs/desmos/v5/x/relationships/types"
)

func (suite *KeeperTestSuite) TestMsgServer_CreateRelationship() {
	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		msg       *types.MsgCreateRelationship
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(false)
			},
			msg: types.NewMsgCreateRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				1,
			),
			shouldErr: true,
		},
		{
			name: "blocked user returns error",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveUserBlock(ctx, types.NewUserBlock(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"",
					1,
				))
			},
			msg: types.NewMsgCreateRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				1,
			),
			shouldErr: true,
		},
		{
			name: "existing relationship returns error",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveRelationship(ctx, types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					1,
				))
			},
			msg: types.NewMsgCreateRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				1,
			),
			shouldErr: true,
		},
		{
			name: "non existing relationship is created correctly",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
			},
			msg: types.NewMsgCreateRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				1,
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgCreateRelationship{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
				),
				sdk.NewEvent(
					types.EventTypeRelationshipCreated,
					sdk.NewAttribute(types.AttributeRelationshipCreator, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeRelationshipCounterparty, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
					sdk.NewAttribute(types.AttributeKeySubspace, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				suite.Require().True(suite.k.HasRelationship(
					ctx,
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					1,
				))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setup != nil {
				tc.setup()
			}
			if tc.store != nil {
				tc.store(ctx)
			}

			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.CreateRelationship(sdk.WrapSDKContext(ctx), tc.msg)

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

func (suite *KeeperTestSuite) TestMsgServer_DeleteRelationship() {
	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		msg       *types.MsgDeleteRelationship
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(false)
			},
			msg: types.NewMsgDeleteRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				1,
			),
			shouldErr: true,
		},
		{
			name: "non existing relationship returns error",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
			},
			msg: types.NewMsgDeleteRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				1,
			),
			shouldErr: true,
		},
		{
			name: "existing relationship is deleted correctly",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveRelationship(ctx, types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					1,
				))
			},
			msg: types.NewMsgDeleteRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				1,
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgDeleteRelationship{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
				),
				sdk.NewEvent(
					types.EventTypeRelationshipDeleted,
					sdk.NewAttribute(types.AttributeRelationshipCreator, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeRelationshipCounterparty, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
					sdk.NewAttribute(types.AttributeKeySubspace, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasRelationship(ctx,
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					1,
				))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setup != nil {
				tc.setup()
			}
			if tc.store != nil {
				tc.store(ctx)
			}

			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.DeleteRelationship(sdk.WrapSDKContext(ctx), tc.msg)

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

func (suite *KeeperTestSuite) TestMsgServer_BlockUser() {
	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		msg       *types.MsgBlockUser
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(false)
			},
			msg: types.NewMsgBlockUser(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"reason",
				1,
			),
			shouldErr: true,
		},
		{
			name: "existing block returns error",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveUserBlock(ctx, types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"reason",
					1,
				))
			},
			msg: types.NewMsgBlockUser(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"reason",
				1,
			),
			shouldErr: true,
		},
		{
			name: "non existing block is stored correctly",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
			},
			msg: types.NewMsgBlockUser(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"reason",
				1,
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgBlockUser{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
				),
				sdk.NewEvent(
					types.EventTypeBlockUser,
					sdk.NewAttribute(types.AttributeKeyUserBlockBlocker, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeKeyUserBlockBlocked, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
					sdk.NewAttribute(types.AttributeKeySubspace, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				suite.Require().True(suite.k.HasUserBlocked(
					ctx,
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					1,
				))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setup != nil {
				tc.setup()
			}
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
		setup     func()
		store     func(ctx sdk.Context)
		msg       *types.MsgUnblockUser
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(false)
			},
			msg: types.NewMsgUnblockUser(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30",
				1,
			),
			shouldErr: true,
		},
		{
			name: "non existing block returns error",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
			},
			msg: types.NewMsgUnblockUser(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30",
				1,
			),
			shouldErr: true,
		},
		{
			name: "existing block is removed properly",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveUserBlock(ctx, types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"reason",
					1,
				))
			},
			msg: types.NewMsgUnblockUser(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				1,
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgUnblockUser{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
				),
				sdk.NewEvent(
					types.EventTypeUnblockUser,
					sdk.NewAttribute(types.AttributeKeyUserBlockBlocker, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeKeyUserBlockBlocked, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
					sdk.NewAttribute(types.AttributeKeySubspace, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasUserBlocked(
					ctx,
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					1,
				))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setup != nil {
				tc.setup()
			}
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
