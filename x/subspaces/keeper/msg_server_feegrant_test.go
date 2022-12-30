package keeper_test

import (
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

func (suite *KeeperTestSuite) TestMsgServer_GrantUserAllowance() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgGrantUserAllowance
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "subspace not found returns error",
			msg: types.NewMsgGrantUserAllowance(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
				&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			),
			shouldErr: true,
		},
		{
			name: "granter has no grant allowances permission returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			msg: types.NewMsgGrantUserAllowance(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
				&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			),
			shouldErr: true,
		},
		{
			name: "duplicated grant returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SetUserPermissions(ctx, 1, types.RootSectionID, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", types.NewPermissions(types.PermissionGrantAllowances))
				grant, err := types.NewUserGrant(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5", &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))})
				suite.Require().NoError(err)
				suite.k.SaveUserGrant(ctx, grant)
			},
			msg: types.NewMsgGrantUserAllowance(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
				&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			),
			shouldErr: true,
		},
		{
			name: "unpack allowance returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SetUserPermissions(ctx, 1, types.RootSectionID, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", types.NewPermissions(types.PermissionGrantAllowances))
			},
			msg: &types.MsgGrantUserAllowance{
				SubspaceID: 1,
				Granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				Grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name: "user allowance set properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SetUserPermissions(ctx, 1, types.RootSectionID, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", types.NewPermissions(types.PermissionGrantAllowances))
			},
			msg:       types.NewMsgGrantUserAllowance(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5", &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))}),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgGrantUserAllowance{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
				),
				sdk.NewEvent(
					types.EventTypeGrantUserAllowance,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyGranter, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
					sdk.NewAttribute(types.AttributeKeyGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				),
			},
			check: func(ctx sdk.Context) {
				grant, found := suite.k.GetUserGrant(ctx, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.Require().True(found)

				expected, err := types.NewUserGrant(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5", &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))})
				suite.Require().NoError(err)
				suite.Require().Equal(expected, grant)
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

			// Run the message
			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.GrantUserAllowance(sdk.WrapSDKContext(ctx), tc.msg)

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

func (suite *KeeperTestSuite) TestMsgServer_RevokeUserAllowance() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgRevokeUserAllowance
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "allowance does not exist returns error",
			msg: types.NewMsgRevokeUserAllowance(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "user allowance revoked properly",
			store: func(ctx sdk.Context) {
				grant, err := types.NewUserGrant(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5", &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))})
				suite.Require().NoError(err)
				suite.k.SaveUserGrant(ctx, grant)
			},
			msg:       types.NewMsgRevokeUserAllowance(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgRevokeUserAllowance{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
				),
				sdk.NewEvent(
					types.EventTypeRevokeUserAllowance,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyGranter, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
					sdk.NewAttribute(types.AttributeKeyGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				),
			},
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasUserGrant(ctx, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"))
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

			// Run the message
			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.RevokeUserAllowance(sdk.WrapSDKContext(ctx), tc.msg)

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

func (suite *KeeperTestSuite) TestMsgServer_GranGroupAllowance() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgGrantGroupAllowance
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "subspace not found returns error",
			msg: types.NewMsgGrantGroupAllowance(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				1,
				&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			),
			shouldErr: true,
		},
		{
			name: "group not found returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			msg: types.NewMsgGrantGroupAllowance(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				1,
				&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			),
			shouldErr: true,
		},
		{
			name: "granter has no grant allowances permission returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(1, 0, 1, "test", "test", nil))
			},
			msg: types.NewMsgGrantGroupAllowance(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				1,
				&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			),
			shouldErr: true,
		},
		{
			name: "duplicated grant returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(1, 0, 1, "test", "test", nil))
				suite.k.SetUserPermissions(ctx, 1, types.RootSectionID, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", types.NewPermissions(types.PermissionGrantAllowances))
				grant, err := types.NewGroupGrant(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", 1, &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))})
				suite.Require().NoError(err)
				suite.k.SaveGroupGrant(ctx, grant)
			},
			msg: types.NewMsgGrantGroupAllowance(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				1,
				&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			),
			shouldErr: true,
		},
		{
			name: "unpack allowance returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(1, 0, 1, "test", "test", nil))
				suite.k.SetUserPermissions(ctx, 1, types.RootSectionID, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", types.NewPermissions(types.PermissionGrantAllowances))
			},
			msg: &types.MsgGrantGroupAllowance{
				SubspaceID: 1,
				Granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				GroupID:    1,
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name: "user allowance set properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(1, 0, 1, "test", "test", nil))
				suite.k.SetUserPermissions(ctx, 1, types.RootSectionID, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", types.NewPermissions(types.PermissionGrantAllowances))
			},
			msg:       types.NewMsgGrantGroupAllowance(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", 1, &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))}),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgGrantGroupAllowance{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
				),
				sdk.NewEvent(
					types.EventTypeGrantGroupAllowance,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyGranter, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
					sdk.NewAttribute(types.AttributeKeyUserGroupID, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				grant, found := suite.k.GetGroupGrant(ctx, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", 1)
				suite.Require().True(found)

				expected, err := types.NewGroupGrant(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", 1, &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))})
				suite.Require().NoError(err)
				suite.Require().Equal(expected, grant)
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

			// Run the message
			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.GrantGroupAllowance(sdk.WrapSDKContext(ctx), tc.msg)

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

func (suite *KeeperTestSuite) TestMsgServer_RevokeGroupAllowance() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgRevokeGroupAllowance
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "allowance does not exist returns error",
			msg: types.NewMsgRevokeGroupAllowance(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				1,
			),
			shouldErr: true,
		},
		{
			name: "group allowance revoked properly",
			store: func(ctx sdk.Context) {
				grant, err := types.NewGroupGrant(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", 1, &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))})
				suite.Require().NoError(err)
				suite.k.SaveGroupGrant(ctx, grant)
			},
			msg:       types.NewMsgRevokeGroupAllowance(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", 1),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgRevokeGroupAllowance{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
				),
				sdk.NewEvent(
					types.EventTypeRevokeGroupAllowance,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyGranter, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
					sdk.NewAttribute(types.AttributeKeyUserGroupID, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasGroupGrant(ctx, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", 1))
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

			// Run the message
			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.RevokeGroupAllowance(sdk.WrapSDKContext(ctx), tc.msg)

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
