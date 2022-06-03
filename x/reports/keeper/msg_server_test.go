package keeper_test

import (
	"time"

	poststypes "github.com/desmos-labs/desmos/v3/x/posts/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/reports/keeper"
	"github.com/desmos-labs/desmos/v3/x/reports/types"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

func (suite *KeeperTestsuite) TestMsgServer_CreateReport() {
	testCases := []struct {
		name        string
		setupCtx    func(ctx sdk.Context) sdk.Context
		store       func(ctx sdk.Context)
		msg         *types.MsgCreateReport
		shouldErr   bool
		expResponse *types.MsgCreateReportResponse
		expEvents   sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			msg: types.NewMsgCreateReport(
				1,
				1,
				"This content is spam!",
				types.NewUserTarget("cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd"),
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "non existing reason returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SetNextReportID(ctx, 1, 1)
			},
			msg: types.NewMsgCreateReport(
				1,
				1,
				"This content is spam!",
				types.NewUserTarget("cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd"),
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "no permission returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SetNextReportID(ctx, 1, 1)

				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam, or the user is spamming",
				))
			},
			msg: types.NewMsgCreateReport(
				1,
				1,
				"This content is spam!",
				types.NewUserTarget("cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd"),
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "invalid report data returns error",
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SetNextReportID(ctx, 1, 1)

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
					subspacestypes.PermissionReportContent,
				)

				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam, or the user is spamming",
				))
			},
			msg: types.NewMsgCreateReport(
				1,
				1,
				"This content is spam!",
				types.NewUserTarget(""),
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "valid request works properly - user target",
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SetNextReportID(ctx, 1, 1)

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
					subspacestypes.PermissionReportContent,
				)

				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam, or the user is spamming",
				))
			},
			msg: types.NewMsgCreateReport(
				1,
				1,
				"This content is spam!",
				types.NewUserTarget("cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd"),
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: false,
			expResponse: &types.MsgCreateReportResponse{
				ReportID:     1,
				CreationDate: time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgCreateReport{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh"),
				),
				sdk.NewEvent(
					types.EventTypeCreateReport,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyReportID, "1"),
					sdk.NewAttribute(types.AttributeKeyReasonID, "1"),
					sdk.NewAttribute(types.AttributeKeyReporter, "cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh"),
					sdk.NewAttribute(types.AttributeKeyCreationTime, time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC).Format(time.RFC3339)),
				),
				sdk.NewEvent(
					types.EventTypeReportUser,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyUser, "cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd"),
					sdk.NewAttribute(types.AttributeKeyReporter, "cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh"),
				),
			},
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				nextReportID := types.GetReportIDFromBytes(store.Get(types.NextReportIDStoreKey(1)))
				suite.Require().Equal(uint64(2), nextReportID)
			},
		},
		{
			name: "valid request works properly - post target",
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SetNextReportID(ctx, 1, 1)

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
					subspacestypes.PermissionReportContent,
				)

				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam, or the user is spamming",
				))

				suite.pk.SavePost(ctx, poststypes.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					poststypes.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			msg: types.NewMsgCreateReport(
				1,
				1,
				"This content is spam!",
				types.NewPostTarget(1),
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: false,
			expResponse: &types.MsgCreateReportResponse{
				ReportID:     1,
				CreationDate: time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgCreateReport{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh"),
				),
				sdk.NewEvent(
					types.EventTypeCreateReport,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyReportID, "1"),
					sdk.NewAttribute(types.AttributeKeyReasonID, "1"),
					sdk.NewAttribute(types.AttributeKeyReporter, "cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh"),
					sdk.NewAttribute(types.AttributeKeyCreationTime, time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC).Format(time.RFC3339)),
				),
				sdk.NewEvent(
					types.EventTypeReportPost,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyPostID, "1"),
					sdk.NewAttribute(types.AttributeKeyReporter, "cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh"),
				),
			},
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				nextReportID := types.GetReportIDFromBytes(store.Get(types.NextReportIDStoreKey(1)))
				suite.Require().Equal(uint64(2), nextReportID)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setupCtx != nil {
				ctx = tc.setupCtx(ctx)
			}
			if tc.store != nil {
				tc.store(ctx)
			}

			msgServer := keeper.NewMsgServerImpl(suite.k)
			res, err := msgServer.CreateReport(sdk.WrapSDKContext(ctx), tc.msg)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expResponse, res)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_DeleteReport() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgDeleteReport
		shouldErr bool
		expEvents sdk.Events
	}{
		{
			name: "non existing subspace returns error",
			msg: types.NewMsgDeleteReport(
				1,
				1,
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "non existing report returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			msg: types.NewMsgDeleteReport(
				1,
				1,
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "no permission returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					1,
					"This content is spam",
					types.NewUserTarget("cosmos1pjffdtweghpyxru9alssyqtdkq8mn6sepgstgm"),
					"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			msg: types.NewMsgDeleteReport(
				1,
				1,
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "invalid signer returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					1,
					"This content is spam",
					types.NewUserTarget("cosmos1pjffdtweghpyxru9alssyqtdkq8mn6sepgstgm"),
					"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			msg: types.NewMsgDeleteReport(
				1,
				1,
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "report creator can delete the report properly",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
					subspacestypes.PermissionDeleteOwnReports,
				)

				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					1,
					"This content is spam",
					types.NewUserTarget("cosmos1pjffdtweghpyxru9alssyqtdkq8mn6sepgstgm"),
					"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			msg: types.NewMsgDeleteReport(
				1,
				1,
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgDeleteReport{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh"),
				),
				sdk.NewEvent(
					types.EventTypeDeleteReport,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyReportID, "1"),
				),
			},
		},
		{
			name: "moderator can delete the report properly",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
					subspacestypes.PermissionManageReports,
				)

				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					1,
					"This content is spam",
					types.NewUserTarget("cosmos1pjffdtweghpyxru9alssyqtdkq8mn6sepgstgm"),
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			msg: types.NewMsgDeleteReport(
				1,
				1,
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgDeleteReport{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh"),
				),
				sdk.NewEvent(
					types.EventTypeDeleteReport,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyReportID, "1"),
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

			msgServer := keeper.NewMsgServerImpl(suite.k)
			_, err := msgServer.DeleteReport(sdk.WrapSDKContext(ctx), tc.msg)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_SupportStandardReason() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		msg         *types.MsgSupportStandardReason
		shouldErr   bool
		expResponse *types.MsgSupportStandardReasonResponse
		expEvents   sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			msg: types.NewMsgSupportStandardReason(
				1,
				1,
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "non existing standard reason returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SetParams(ctx, types.NewParams(nil))
			},
			msg: types.NewMsgSupportStandardReason(
				1,
				1,
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "no permission returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SetParams(ctx, types.NewParams([]types.StandardReason{
					types.NewStandardReason(1, "Spam", "This content is spam"),
				}))
			},
			msg: types.NewMsgSupportStandardReason(
				1,
				1,
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "not found next reason id returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
					subspacestypes.PermissionManageReasons,
				)

				suite.k.SetParams(ctx, types.NewParams([]types.StandardReason{
					types.NewStandardReason(1, "Spam", "This content is spam"),
				}))
			},
			msg: types.NewMsgSupportStandardReason(
				1,
				1,
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "valid request returns no error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SetNextReasonID(ctx, 1, 1)

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
					subspacestypes.PermissionManageReasons,
				)

				suite.k.SetParams(ctx, types.NewParams([]types.StandardReason{
					types.NewStandardReason(1, "Spam", "This content is spam"),
				}))
			},
			msg: types.NewMsgSupportStandardReason(
				1,
				1,
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: false,
			expResponse: &types.MsgSupportStandardReasonResponse{
				ReasonsID: 1,
			},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgSupportStandardReason{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh"),
				),
				sdk.NewEvent(
					types.EventTypeSupportStandardReason,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyStandardReasonID, "1"),
					sdk.NewAttribute(types.AttributeKeyReasonID, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				nextReasonID := types.GetReasonIDFromBytes(store.Get(types.NextReasonIDStoreKey(1)))
				suite.Require().Equal(uint32(2), nextReasonID)
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

			msgServer := keeper.NewMsgServerImpl(suite.k)
			res, err := msgServer.SupportStandardReason(sdk.WrapSDKContext(ctx), tc.msg)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expResponse, res)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_AddReason() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		msg         *types.MsgAddReason
		shouldErr   bool
		expResponse *types.MsgAddReasonResponse
		expEvents   sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			msg: types.NewMsgAddReason(
				1,
				"Spam",
				"This content is spam",
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "no permission returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			msg: types.NewMsgAddReason(
				1,
				"Spam",
				"This content is spam",
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "no next reason id returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
					subspacestypes.PermissionManageReasons,
				)
			},
			msg: types.NewMsgAddReason(
				1,
				"Spam",
				"This content is spam",
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "invalid reason returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SetNextReasonID(ctx, 1, 1)

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
					subspacestypes.PermissionManageReasons,
				)
			},
			msg: types.NewMsgAddReason(
				1,
				"",
				"This content is spam",
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "correct request returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SetNextReasonID(ctx, 1, 1)

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
					subspacestypes.PermissionManageReasons,
				)
			},
			msg: types.NewMsgAddReason(
				1,
				"Spam",
				"This content is spam",
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: false,
			expResponse: &types.MsgAddReasonResponse{
				ReasonID: 1,
			},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgAddReason{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh"),
				),
				sdk.NewEvent(
					types.EventTypeAddReason,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyReasonID, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				nextReasonID := types.GetReasonIDFromBytes(store.Get(types.NextReasonIDStoreKey(1)))
				suite.Require().Equal(uint64(2), nextReasonID)
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

			msgServer := keeper.NewMsgServerImpl(suite.k)
			res, err := msgServer.AddReason(sdk.WrapSDKContext(ctx), tc.msg)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expResponse, res)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_RemoveReason() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgRemoveReason
		shouldErr bool
		expEvents sdk.Events
	}{
		{
			name: "non existing subspace returns error",
			msg: types.NewMsgRemoveReason(
				1,
				1,
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "non existing reason returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			msg: types.NewMsgRemoveReason(
				1,
				1,
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "no permission returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				))
			},
			msg: types.NewMsgRemoveReason(
				1,
				1,
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "valid request works properly",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				))

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
					subspacestypes.PermissionManageReasons,
				)
			},
			msg: types.NewMsgRemoveReason(
				1,
				1,
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgRemoveReason{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh"),
				),
				sdk.NewEvent(
					types.EventTypeRemoveReason,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyReasonID, "1"),
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

			msgServer := keeper.NewMsgServerImpl(suite.k)
			_, err := msgServer.RemoveReason(sdk.WrapSDKContext(ctx), tc.msg)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())
			}
		})
	}
}
