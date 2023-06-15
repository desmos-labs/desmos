package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"

	"github.com/desmos-labs/desmos/v5/x/posts/keeper"
	"github.com/desmos-labs/desmos/v5/x/posts/types"
)

func (suite *KeeperTestSuite) TestMsgServer_RequestPostOwnerTransfer() {
	testCases := []struct {
		name        string
		setup       func()
		store       func(ctx sdk.Context)
		msg         *types.MsgRequestPostOwnerTransfer
		shouldErr   bool
		expResponse *types.MsgRequestPostOwnerTransferResponse
		expEvents   sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "sender has no profile returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg").
					Return(false)
			},
			msg: types.NewMsgRequestPostOwnerTransfer(
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "receiver has no profile returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg").
					Return(true)

				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").
					Return(false)
			},
			msg: types.NewMsgRequestPostOwnerTransfer(
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "receiver has blocked sender returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg").
					Return(true)

				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").
					Return(true)

				suite.rk.EXPECT().
					HasUserBlocked(
						gomock.Any(),
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
						uint64(1),
					).
					Return(true)
			},
			msg: types.NewMsgRequestPostOwnerTransfer(
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "post not found returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg").
					Return(true)

				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").
					Return(true)

				suite.rk.EXPECT().
					HasUserBlocked(
						gomock.Any(),
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
						uint64(1),
					).
					Return(false)
			},
			msg: types.NewMsgRequestPostOwnerTransfer(
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "sender does not match the owner returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg").
					Return(true)

				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").
					Return(true)

				suite.rk.EXPECT().
					HasUserBlocked(
						gomock.Any(),
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
						uint64(1),
					).
					Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"other_owner",
				))
			},
			msg: types.NewMsgRequestPostOwnerTransfer(
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "request already exists returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg").
					Return(true)

				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").
					Return(true)

				suite.rk.EXPECT().
					HasUserBlocked(
						gomock.Any(),
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
						uint64(1),
					).
					Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))

				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(
					1,
					1,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))
			},
			msg: types.NewMsgRequestPostOwnerTransfer(
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "valid transfer is performed properly",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg").
					Return(true)

				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").
					Return(true)

				suite.rk.EXPECT().
					HasUserBlocked(
						gomock.Any(),
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
						uint64(1),
					).
					Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))
			},
			msg: types.NewMsgRequestPostOwnerTransfer(
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			expResponse: &types.MsgRequestPostOwnerTransferResponse{},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgRequestPostOwnerTransfer{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg"),
				),
				sdk.NewEvent(
					types.EventTypeRequestPostOwnerTransfer,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyPostID, "1"),
					sdk.NewAttribute(types.AttributeKeyReceiver, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd"),
					sdk.NewAttribute(types.AttributeKeySender, "cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg"),
				),
			},
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)

				var request types.PostOwnerTransferRequest
				suite.cdc.MustUnmarshal(
					store.Get(types.PostOwnerTransferRequestStoreKey(1, 1)), &request,
				)

				suite.Require().Equal(types.NewPostOwnerTransferRequest(
					1,
					1,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				), request)
			},
			shouldErr: false,
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
			res, err := msgServer.RequestPostOwnerTransfer(sdk.WrapSDKContext(ctx), tc.msg)
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

func (suite *KeeperTestSuite) TestMsgServer_CancelPostOwnerTransfer() {
	testCases := []struct {
		name        string
		setup       func()
		store       func(ctx sdk.Context)
		msg         *types.MsgCancelPostOwnerTransferRequest
		shouldErr   bool
		expResponse *types.MsgCancelPostOwnerTransferRequestResponse
		expEvents   sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "request does not exist returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg").
					Return(true)
			},
			msg: types.NewMsgCancelPostOwnerTransferRequest(
				1,
				1,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "request sender does not match the sender returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg").
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(
					1,
					1,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					"other_sender",
				))
			},
			msg: types.NewMsgCancelPostOwnerTransferRequest(
				1,
				1,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "valid request is performed properly",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg").
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(
					1,
					1,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))
			},
			msg: types.NewMsgCancelPostOwnerTransferRequest(
				1,
				1,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			expResponse: &types.MsgCancelPostOwnerTransferRequestResponse{},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgCancelPostOwnerTransferRequest{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg"),
				),
				sdk.NewEvent(
					types.EventTypeCancelPostOwnerTransfer,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyPostID, "1"),
					sdk.NewAttribute(types.AttributeKeySender, "cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg"),
				),
			},
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasPostOwnerTransferRequest(ctx, 1, 1))
			},
			shouldErr: false,
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
			res, err := msgServer.CancelPostOwnerTransferRequest(sdk.WrapSDKContext(ctx), tc.msg)
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

func (suite *KeeperTestSuite) TestMsgServer_AcceptPostOwnerTransfer() {
	testCases := []struct {
		name        string
		setup       func()
		store       func(ctx sdk.Context)
		setupCtx    func(ctx sdk.Context) sdk.Context
		msg         *types.MsgAcceptPostOwnerTransferRequest
		shouldErr   bool
		expResponse *types.MsgAcceptPostOwnerTransferRequestResponse
		expEvents   sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "receiver has no profile returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").
					Return(false)
			},
			msg: types.NewMsgAcceptPostOwnerTransferRequest(
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "request does not exist returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").
					Return(true)
			},
			msg: types.NewMsgAcceptPostOwnerTransferRequest(
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "request receiver does not match the receiver returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(
					1,
					1,
					"other_receiver",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))
			},
			msg: types.NewMsgAcceptPostOwnerTransferRequest(
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "post not found returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(
					1,
					1,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))
			},
			msg: types.NewMsgAcceptPostOwnerTransferRequest(
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "invalid updated post returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "invalid_receiver").
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(
					1,
					1,
					"invalid_receiver",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))
			},
			msg: types.NewMsgAcceptPostOwnerTransferRequest(
				1,
				1,
				"invalid_receiver",
			),
			shouldErr: true,
		},
		{
			name: "request sender does not match the post sender returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(
					1,
					1,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"other_owner",
				))
			},
			msg: types.NewMsgAcceptPostOwnerTransferRequest(
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "correct request is performed properly",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").
					Return(true)
			},
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2100, 5, 17, 0, 0, 0, 0, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(
					1,
					1,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))
			},
			msg: types.NewMsgAcceptPostOwnerTransferRequest(
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			expResponse: &types.MsgAcceptPostOwnerTransferRequestResponse{},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgAcceptPostOwnerTransferRequest{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd"),
				),
				sdk.NewEvent(
					types.EventTypeAcceptPostOwnerTransfer,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyPostID, "1"),
					sdk.NewAttribute(types.AttributeKeyReceiver, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd"),
				),
			},
			shouldErr: false,
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
			res, err := msgServer.AcceptPostOwnerTransferRequest(sdk.WrapSDKContext(ctx), tc.msg)
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

func (suite *KeeperTestSuite) TestMsgServer_RefusePostOwnerTransfer() {
	testCases := []struct {
		name        string
		setup       func()
		store       func(ctx sdk.Context)
		msg         *types.MsgRefusePostOwnerTransferRequest
		shouldErr   bool
		expResponse *types.MsgRefusePostOwnerTransferRequestResponse
		expEvents   sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "request does not exist returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").
					Return(true)
			},
			msg: types.NewMsgRefusePostOwnerTransferRequest(
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "request receiver does not match the receiver returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(
					1,
					1,
					"other_receiver",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))
			},
			msg: types.NewMsgRefusePostOwnerTransferRequest(
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "correct request is performed properly",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(
					1,
					1,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))
			},
			msg: types.NewMsgRefusePostOwnerTransferRequest(
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			expResponse: &types.MsgRefusePostOwnerTransferRequestResponse{},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgRefusePostOwnerTransferRequest{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd"),
				),
				sdk.NewEvent(
					types.EventTypeRefusePostOwnerTransfer,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyPostID, "1"),
					sdk.NewAttribute(types.AttributeKeyReceiver, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd"),
				),
			},
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasPostOwnerTransferRequest(ctx, 1, 1))
			},
			shouldErr: false,
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
			res, err := msgServer.RefusePostOwnerTransferRequest(sdk.WrapSDKContext(ctx), tc.msg)
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
