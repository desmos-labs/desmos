package keeper_test

import (
	"time"

	"github.com/golang/mock/gomock"

	poststypes "github.com/desmos-labs/desmos/v6/x/posts/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/v6/x/reports/keeper"
	"github.com/desmos-labs/desmos/v6/x/reports/types"
	subspacestypes "github.com/desmos-labs/desmos/v6/x/subspaces/types"
)

func (suite *KeeperTestSuite) TestMsgServer_CreateReport() {
	testCases := []struct {
		name        string
		setup       func()
		setupCtx    func(ctx sdk.Context) sdk.Context
		store       func(ctx sdk.Context)
		msg         *types.MsgCreateReport
		shouldErr   bool
		expResponse *types.MsgCreateReportResponse
		expEvents   sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "user without profile returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh").
					Return(false)
			},
			msg: types.NewMsgCreateReport(
				1,
				[]uint32{1},
				"This content is spam!",
				types.NewUserTarget("cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd"),
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "non existing subspace returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh").
					Return(true)

				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(false)
			},
			msg: types.NewMsgCreateReport(
				1,
				[]uint32{1},
				"This content is spam!",
				types.NewUserTarget("cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd"),
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "non existing reason returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh").
					Return(true)

				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextReportID(ctx, 1, 1)
			},
			msg: types.NewMsgCreateReport(
				1,
				[]uint32{1},
				"This content is spam!",
				types.NewUserTarget("cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd"),
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "no permission returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh").
					Return(true)

				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
						types.PermissionReportContent,
					).
					Return(false)
			},
			store: func(ctx sdk.Context) {
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
				[]uint32{1},
				"This content is spam!",
				types.NewUserTarget("cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd"),
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "invalid report data returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh").
					Return(true)

				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
						types.PermissionReportContent,
					).
					Return(true)
			},
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
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
				[]uint32{1},
				"This content is spam!",
				types.NewUserTarget(""),
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "duplicated report returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh").
					Return(true)

				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
						types.PermissionReportContent,
					).
					Return(true)
			},
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextReportID(ctx, 1, 1)

				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam, or the user is spamming",
				))

				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"This user is spamming",
					types.NewUserTarget("cosmos1wprgptc8ktt0eemrn2znpxv8crdxm8tdpkdr7w"),
					"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			msg: types.NewMsgCreateReport(
				1,
				[]uint32{1},
				"This user is spamming!",
				types.NewUserTarget("cosmos1wprgptc8ktt0eemrn2znpxv8crdxm8tdpkdr7w"),
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "valid request works properly - user target",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh").
					Return(true)

				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
						types.PermissionReportContent,
					).
					Return(true)

				suite.rk.
					EXPECT().
					HasUserBlocked(
						gomock.Any(),
						"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
						"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
						uint64(1),
					).
					Return(false)
			},
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
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
				[]uint32{1},
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
					types.EventTypeCreateReport,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyReportID, "1"),
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
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh").
					Return(true)

				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
						types.PermissionReportContent,
					).
					Return(true)

				suite.pk.EXPECT().
					GetPost(gomock.Any(), uint64(1), uint64(1)).
					Return(poststypes.NewPost(
						1,
						0,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						1,
						nil,
						nil,
						nil,
						poststypes.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					), true)

				suite.rk.
					EXPECT().
					HasUserBlocked(
						gomock.Any(),
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
						uint64(1),
					).
					Return(false)
			},
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
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
				[]uint32{1},
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
					types.EventTypeCreateReport,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyReportID, "1"),
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
			if tc.setup != nil {
				tc.setup()
			}
			if tc.setupCtx != nil {
				ctx = tc.setupCtx(ctx)
			}
			if tc.store != nil {
				tc.store(ctx)
			}

			msgServer := keeper.NewMsgServerImpl(suite.k)
			res, err := msgServer.CreateReport(ctx, tc.msg)
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

func (suite *KeeperTestSuite) TestMsgServer_DeleteReport() {
	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		msg       *types.MsgDeleteReport
		shouldErr bool
		expEvents sdk.Events
	}{
		{
			name: "non existing subspace returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(false)
			},
			msg: types.NewMsgDeleteReport(
				1,
				1,
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "non existing report returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)
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
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
						types.PermissionManageReports,
					).
					Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
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
			name: "no permission returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
						types.PermissionManageReports,
					).
					Return(false)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
						types.PermissionDeleteOwnReports,
					).
					Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
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
			name: "report creator can delete the report properly",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
						types.PermissionManageReports,
					).
					Return(false)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
						types.PermissionDeleteOwnReports,
					).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
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
					types.EventTypeDeleteReport,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyReportID, "1"),
				),
			},
		},
		{
			name: "moderator can delete the report properly",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
						types.PermissionManageReports,
					).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
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
			if tc.setup != nil {
				tc.setup()
			}
			if tc.store != nil {
				tc.store(ctx)
			}

			msgServer := keeper.NewMsgServerImpl(suite.k)
			_, err := msgServer.DeleteReport(ctx, tc.msg)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServer_SupportStandardReason() {
	testCases := []struct {
		name        string
		setup       func()
		store       func(ctx sdk.Context)
		msg         *types.MsgSupportStandardReason
		shouldErr   bool
		expResponse *types.MsgSupportStandardReasonResponse
		expEvents   sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(false)
			},
			msg: types.NewMsgSupportStandardReason(
				1,
				1,
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "non existing standard reason returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)
			},
			store: func(ctx sdk.Context) {
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
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
						types.PermissionManageReasons,
					).
					Return(false)
			},
			store: func(ctx sdk.Context) {
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
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
						types.PermissionManageReasons,
					).
					Return(true)
			},
			store: func(ctx sdk.Context) {
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
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
						types.PermissionManageReasons,
					).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextReasonID(ctx, 1, 1)

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
			if tc.setup != nil {
				tc.setup()
			}
			if tc.store != nil {
				tc.store(ctx)
			}

			msgServer := keeper.NewMsgServerImpl(suite.k)
			res, err := msgServer.SupportStandardReason(ctx, tc.msg)
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

func (suite *KeeperTestSuite) TestMsgServer_AddReason() {
	testCases := []struct {
		name        string
		setup       func()
		store       func(ctx sdk.Context)
		msg         *types.MsgAddReason
		shouldErr   bool
		expResponse *types.MsgAddReasonResponse
		expEvents   sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(false)
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
			name: "no permission returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
						types.PermissionManageReasons,
					).
					Return(false)
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
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
						types.PermissionManageReasons,
					).
					Return(true)
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
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
						types.PermissionManageReasons,
					).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextReasonID(ctx, 1, 1)
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
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
						types.PermissionManageReasons,
					).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextReasonID(ctx, 1, 1)
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
			if tc.setup != nil {
				tc.setup()
			}
			if tc.store != nil {
				tc.store(ctx)
			}

			msgServer := keeper.NewMsgServerImpl(suite.k)
			res, err := msgServer.AddReason(ctx, tc.msg)
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

func (suite *KeeperTestSuite) TestMsgServer_RemoveReason() {
	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		msg       *types.MsgRemoveReason
		shouldErr bool
		expEvents sdk.Events
	}{
		{
			name: "non existing subspace returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(false)
			},
			msg: types.NewMsgRemoveReason(
				1,
				1,
				"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
			),
			shouldErr: true,
		},
		{
			name: "non existing reason returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)
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
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
						types.PermissionManageReasons,
					).
					Return(false)
			},
			store: func(ctx sdk.Context) {
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
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qycmg40ju50fx2mcc82qtkzuswjs3mj3mqekeh",
						types.PermissionManageReasons,
					).
					Return(true)
			},
			store: func(ctx sdk.Context) {
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
			shouldErr: false,
			expEvents: sdk.Events{
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
			if tc.setup != nil {
				tc.setup()
			}
			if tc.store != nil {
				tc.store(ctx)
			}

			msgServer := keeper.NewMsgServerImpl(suite.k)
			_, err := msgServer.RemoveReason(ctx, tc.msg)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServer_UpdateParams() {
	testCases := []struct {
		name      string
		msg       *types.MsgUpdateParams
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "invalid authority return error",
			msg: types.NewMsgUpdateParams(
				types.DefaultParams(),
				"invalid",
			),
			shouldErr: true,
		},
		{
			name: "set params properly",
			msg: types.NewMsgUpdateParams(
				types.DefaultParams(),
				authtypes.NewModuleAddress("gov").String(),
			),
			shouldErr: false,
			expEvents: sdk.Events{},
			check: func(ctx sdk.Context) {
				params := suite.k.GetParams(ctx)
				suite.Require().Equal(types.DefaultParams(), params)
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()

			// Reset any event that might have been emitted during the setup
			ctx = ctx.WithEventManager(sdk.NewEventManager())

			// Run the message
			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.UpdateParams(ctx, tc.msg)

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
