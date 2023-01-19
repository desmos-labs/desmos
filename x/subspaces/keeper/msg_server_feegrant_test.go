package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

func (suite *KeeperTestSuite) TestMsgServer_GrantAllowance() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgGrantAllowance
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "subspace not found returns error",
			msg: types.NewMsgGrantAllowance(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			),
			shouldErr: true,
		},
		{
			name: "granter has no permission returns error",
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
			msg: types.NewMsgGrantAllowance(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			),
			shouldErr: true,
		},
		{
			name: "duplicated user grant returns error",
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
				suite.k.SetUserPermissions(ctx, 1, types.RootSectionID, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", types.NewPermissions(types.PermissionManageAllowances))
				suite.k.SaveGrant(ctx, types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))}))
			},
			msg: types.NewMsgGrantAllowance(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
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
				suite.k.SetUserPermissions(ctx, 1, types.RootSectionID, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", types.NewPermissions(types.PermissionManageAllowances))
			},
			msg: types.NewMsgGrantAllowance(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewGroupGrantee(1),
				&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			),
			shouldErr: true,
		},
		{
			name: "duplicated group grant returns error",
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
				suite.k.SetUserPermissions(ctx, 1, types.RootSectionID, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", types.NewPermissions(types.PermissionManageAllowances))
				suite.k.SaveGrant(ctx, types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))}))
			},
			msg: types.NewMsgGrantAllowance(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewGroupGrantee(1),
				&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			),
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
				suite.k.SetUserPermissions(ctx, 1, types.RootSectionID, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", types.NewPermissions(types.PermissionManageAllowances))
			},
			msg:       types.NewMsgGrantAllowance(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"), &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))}),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgGrantAllowance{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
				),
				sdk.NewEvent(
					types.EventTypeGrantAllowance,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyGranter, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
					sdk.NewAttribute(types.AttributeKeyUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				),
			},
			check: func(ctx sdk.Context) {
				grant, found := suite.k.GetUserGrant(ctx, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.Require().True(found)
				suite.Require().Equal(types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))}),
					grant)
			},
		},
		{
			name: "group allowance set properly",
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
				suite.k.SetUserPermissions(ctx, 1, types.RootSectionID, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", types.NewPermissions(types.PermissionManageAllowances))
			},
			msg:       types.NewMsgGrantAllowance(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", types.NewGroupGrantee(1), &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))}),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgGrantAllowance{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
				),
				sdk.NewEvent(
					types.EventTypeGrantAllowance,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyGranter, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
					sdk.NewAttribute(types.AttributeKeyGroupGrantee, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				grant, found := suite.k.GetGroupGrant(ctx, 1, 1)
				suite.Require().True(found)

				suite.Require().Equal(types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1), &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))}),
					grant)
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
			_, err := service.GrantAllowance(sdk.WrapSDKContext(ctx), tc.msg)

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

func (suite *KeeperTestSuite) TestMsgServer_RevokeAllowance() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgRevokeAllowance
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "granter has no permission returns error",
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
			msg: types.NewMsgRevokeAllowance(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
			),
			shouldErr: true,
		},
		{
			name: "user allowance does not exist returns error",
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
				suite.k.SetUserPermissions(ctx, 1, types.RootSectionID, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", types.NewPermissions(types.PermissionManageAllowances))
			},
			msg: types.NewMsgRevokeAllowance(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
			),
			shouldErr: true,
		},
		{
			name: "group allowance does not exist returns error",
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
				suite.k.SetUserPermissions(ctx, 1, types.RootSectionID, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", types.NewPermissions(types.PermissionManageAllowances))
			},
			msg: types.NewMsgRevokeAllowance(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewGroupGrantee(1),
			),
			shouldErr: true,
		},
		{
			name: "user allowance revoked properly",
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
				suite.k.SetUserPermissions(ctx, 1, types.RootSectionID, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", types.NewPermissions(types.PermissionManageAllowances))
				suite.k.SaveGrant(ctx, types.NewGrant(
					1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))}))
			},
			msg:       types.NewMsgRevokeAllowance(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgRevokeAllowance{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
				),
				sdk.NewEvent(
					types.EventTypeRevokeAllowance,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyGranter, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
					sdk.NewAttribute(types.AttributeKeyUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				),
			},
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasUserGrant(ctx, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"))
			},
		},
		{
			name: "group allowance revoked properly",
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
				suite.k.SetUserPermissions(ctx, 1, types.RootSectionID, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", types.NewPermissions(types.PermissionManageAllowances))
				suite.k.SaveGrant(ctx, types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))}))
			},
			msg:       types.NewMsgRevokeAllowance(1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", types.NewGroupGrantee(1)),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgRevokeAllowance{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
				),
				sdk.NewEvent(
					types.EventTypeRevokeAllowance,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyGranter, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
					sdk.NewAttribute(types.AttributeKeyGroupGrantee, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasGroupGrant(ctx, 1, 1))
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
			_, err := service.RevokeAllowance(sdk.WrapSDKContext(ctx), tc.msg)

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
