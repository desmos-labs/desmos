package keeper_test

import (
	"time"

	"github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (suite *KeeperTestsuite) TestMsgServer_GrantTreasuryAuthorization() {
	blockTime := time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC)
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		msg         *types.MsgGrantTreasuryAuthorization
		shouldErr   bool
		expResponse *types.MsgGrantTreasuryAuthorizationResponse
		expEvents   []sdk.Event
		check       func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			msg: types.NewMsgGrantTreasuryAuthorization(
				1,
				"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				&banktypes.SendAuthorization{SpendLimit: sdk.NewCoins(sdk.NewInt64Coin("steak", 100))},
				time.Date(2023, 1, 11, 0, 0, 0, 0, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "granter without permissions returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC)))
			},
			msg: types.NewMsgGrantTreasuryAuthorization(
				1,
				"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				&banktypes.SendAuthorization{SpendLimit: sdk.NewCoins(sdk.NewInt64Coin("steak", 100))},
				time.Date(2023, 1, 11, 0, 0, 0, 0, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "invalid treasury address returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(1,
					"Test subspace",
					"This is a test subspace",
					"",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC)))
				suite.k.SetUserPermissions(ctx,
					1,
					types.RootSectionID,
					"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
					types.NewPermissions(types.PermissionManageTreasuryAuthorization))
			},
			msg: types.NewMsgGrantTreasuryAuthorization(
				1,
				"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				&banktypes.SendAuthorization{SpendLimit: sdk.NewCoins(sdk.NewInt64Coin("steak", 100))},
				time.Date(2023, 1, 11, 0, 0, 0, 0, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "invalid grantee address returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC)))
				suite.k.SetUserPermissions(ctx,
					1,
					types.RootSectionID,
					"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
					types.NewPermissions(types.PermissionManageTreasuryAuthorization))
			},
			msg: types.NewMsgGrantTreasuryAuthorization(
				1,
				"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
				"",
				&banktypes.SendAuthorization{SpendLimit: sdk.NewCoins(sdk.NewInt64Coin("steak", 100))},
				time.Date(2023, 1, 11, 0, 0, 0, 0, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "valid request creates authorization properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC)))
				suite.k.SetUserPermissions(ctx,
					1,
					types.RootSectionID,
					"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
					types.NewPermissions(types.PermissionManageTreasuryAuthorization))
			},
			msg: types.NewMsgGrantTreasuryAuthorization(
				1,
				"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				&banktypes.SendAuthorization{SpendLimit: sdk.NewCoins(sdk.NewInt64Coin("steak", 100))},
				time.Date(2023, 1, 11, 0, 0, 0, 0, time.UTC),
			),
			shouldErr:   false,
			expResponse: &types.MsgGrantTreasuryAuthorizationResponse{},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgGrantTreasuryAuthorization{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez"),
				),
				sdk.NewEvent(
					types.EventTypeGrantTreasuryAuthorization,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyGranter, "cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez"),
					sdk.NewAttribute(types.AttributeKeyGrantee, "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69"),
				),
			},
			check: func(ctx sdk.Context) {
				treasury, err := sdk.AccAddressFromBech32("cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn")
				suite.Require().NoError(err)
				grantee, err := sdk.AccAddressFromBech32("cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69")
				suite.Require().NoError(err)
				authorizations := suite.authzKeeper.GetAuthorizations(ctx, grantee, treasury)
				suite.Require().Equal(1, len(authorizations))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			ctx = ctx.WithBlockTime(blockTime)

			if tc.store != nil {
				tc.store(ctx)
			}

			// Reset any event that might have been emitted during the setup
			ctx = ctx.WithEventManager(sdk.NewEventManager())

			// Run the message
			service := keeper.NewMsgServerImpl(suite.k)
			res, err := service.GrantTreasuryAuthorization(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expResponse, res)

				for _, event := range tc.expEvents {
					suite.Require().Contains(ctx.EventManager().Events(), event)
				}

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_RevokeTreasuryAuthorization() {
	blockTime := time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC)
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		msg         *types.MsgRevokeTreasuryAuthorization
		shouldErr   bool
		expResponse *types.MsgRevokeTreasuryAuthorizationResponse
		expEvents   []sdk.Event
		check       func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			msg: types.NewMsgRevokeTreasuryAuthorization(
				1,
				"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"/cosmos.v1beta1.MsgSend",
			),
			shouldErr: true,
		},
		{
			name: "Revokeer has no permission returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC)))
			},
			msg: types.NewMsgRevokeTreasuryAuthorization(
				1,
				"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"/cosmos.v1beta1.MsgSend",
			),
			shouldErr: true,
		},
		{
			name: "invalid treasury address returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(1,
					"Test subspace",
					"This is a test subspace",
					"",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC)))
				suite.k.SetUserPermissions(ctx,
					1,
					types.RootSectionID,
					"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
					types.NewPermissions(types.PermissionManageTreasuryAuthorization))
			},
			msg: types.NewMsgRevokeTreasuryAuthorization(
				1,
				"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"/cosmos.v1beta1.MsgSend",
			),
			shouldErr: true,
		},
		{
			name: "invalid Revokeee address returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC)))
				suite.k.SetUserPermissions(ctx,
					1,
					types.RootSectionID,
					"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
					types.NewPermissions(types.PermissionManageTreasuryAuthorization))
			},
			msg: types.NewMsgRevokeTreasuryAuthorization(
				1,
				"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
				"",
				"/cosmos.v1beta1.MsgSend",
			),
			shouldErr: true,
		},
		{
			name: "non exiting authorization returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC)))
				suite.k.SetUserPermissions(ctx,
					1,
					types.RootSectionID,
					"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
					types.NewPermissions(types.PermissionManageTreasuryAuthorization))
			},
			msg: types.NewMsgRevokeTreasuryAuthorization(
				1,
				"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"/cosmos.v1beta1.MsgSend",
			),
			shouldErr: true,
		},
		{
			name: "valid request deletes authorization properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC)))
				suite.k.SetUserPermissions(ctx,
					1,
					types.RootSectionID,
					"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
					types.NewPermissions(types.PermissionManageTreasuryAuthorization))

				treasury, err := sdk.AccAddressFromBech32("cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn")
				suite.Require().NoError(err)
				grantee, err := sdk.AccAddressFromBech32("cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69")
				suite.Require().NoError(err)

				err = suite.authzKeeper.SaveGrant(ctx,
					grantee,
					treasury,
					&banktypes.SendAuthorization{SpendLimit: sdk.NewCoins(sdk.NewInt64Coin("steak", 100))},
					time.Date(2024, 1, 11, 1, 1, 1, 1, time.UTC))
				suite.Require().NoError(err)
			},
			msg: types.NewMsgRevokeTreasuryAuthorization(
				1,
				"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"/cosmos.bank.v1beta1.MsgSend",
			),
			shouldErr:   false,
			expResponse: &types.MsgRevokeTreasuryAuthorizationResponse{},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgRevokeTreasuryAuthorization{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez"),
				),
				sdk.NewEvent(
					types.EventTypeRevokeTreasuryAuthorization,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyGranter, "cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez"),
					sdk.NewAttribute(types.AttributeKeyGrantee, "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69"),
				),
			},
			check: func(ctx sdk.Context) {
				treasury, err := sdk.AccAddressFromBech32("cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn")
				suite.Require().NoError(err)
				grantee, err := sdk.AccAddressFromBech32("cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69")
				suite.Require().NoError(err)
				authorizations := suite.authzKeeper.GetAuthorizations(ctx, grantee, treasury)
				suite.Require().Equal(0, len(authorizations))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			ctx = ctx.WithBlockTime(blockTime)

			if tc.store != nil {
				tc.store(ctx)
			}

			// Reset any event that might have been emitted during the setup
			ctx = ctx.WithEventManager(sdk.NewEventManager())

			// Run the message
			service := keeper.NewMsgServerImpl(suite.k)
			res, err := service.RevokeTreasuryAuthorization(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expResponse, res)

				for _, event := range tc.expEvents {
					suite.Require().Contains(ctx.EventManager().Events(), event)
				}

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}
