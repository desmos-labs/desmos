package keeper_test

import (
	"fmt"

	"github.com/desmos-labs/desmos/testutil"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/keeper"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestMsgServer_RequestDTagTransfer() {
	testCases := []struct {
		name         string
		store        func(ctx sdk.Context)
		storedBlocks []types.UserBlock
		msg          *types.MsgRequestDTagTransfer
		shouldErr    bool
		expEvents    sdk.Events
	}{
		{
			name: "blocked receiver making request returns error",
			store: func(ctx sdk.Context) {
				block := types.NewUserBlock(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"This user has been blocked",
					"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocker)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocked)))
				suite.Require().NoError(suite.k.SaveUserBlock(ctx, block))

			},
			msg: types.NewMsgRequestDTagTransfer(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "no DTag to transfer returns error",
			msg: types.NewMsgRequestDTagTransfer(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "Already present request returns error",
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"dtag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(request.Receiver)))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			msg: types.NewMsgRequestDTagTransfer(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "not already present request is saved correctly",
			store: func(ctx sdk.Context) {
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")))
			},
			msg: types.NewMsgRequestDTagTransfer(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeDTagTransferRequest,
					sdk.NewAttribute(types.AttributeDTagToTrade, fmt.Sprintf("%s-dtag", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")),
					sdk.NewAttribute(types.AttributeRequestSender, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeRequestReceiver, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
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

			server := keeper.NewMsgServerImpl(suite.k)
			_, err := server.RequestDTagTransfer(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServer_CancelDTagTransfer() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgCancelDTagTransferRequest
		shouldErr bool
		expEvents sdk.Events
	}{
		{
			name:      "request not found returns error",
			msg:       types.NewMsgCancelDTagTransferRequest("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			expEvents: sdk.EmptyEvents(),
			shouldErr: true,
		},
		{
			name: "found request is cancelled correctly",
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"dtag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				)

				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(request.Receiver)))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			msg: types.NewMsgCancelDTagTransferRequest(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeDTagTransferCancel,
					sdk.NewAttribute(types.AttributeRequestSender, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeRequestReceiver, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
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

			server := keeper.NewMsgServerImpl(suite.k)
			_, err := server.CancelDTagTransferRequest(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServer_AcceptDTagTransfer() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgAcceptDTagTransferRequest
		shouldErr bool
		expEvents sdk.Events
	}{
		{
			name: "returns an error if there are no request made from the receiving user",
			msg: types.NewMsgAcceptDTagTransferRequest(
				"newDtag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "returns an error if the request receiver does not have a profile",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				request := types.NewDTagTransferRequest(
					"dtag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				)
				store.Set(
					types.DTagTransferRequestStoreKey(request.Sender, request.Receiver),
					suite.cdc.MustMarshal(&request),
				)
			},
			msg: types.NewMsgAcceptDTagTransferRequest(
				"newDtag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "returns an error if the new DTag has already been chosen by another user",
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"dtag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(request.Sender)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(request.Receiver)))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))

				profile := testutil.ProfileFromAddr("cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn")
				profile.DTag = "NewDTag"
				suite.Require().NoError(suite.k.StoreProfile(ctx, profile))
			},
			msg: types.NewMsgAcceptDTagTransferRequest(
				"newDtag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "returns an error when current owner DTag is different from the requested one",
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"dtag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(request.Sender)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(request.Receiver)))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))

				profile := testutil.ProfileFromAddr("cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn")
				profile.DTag = "dtag"
				suite.Require().NoError(suite.k.StoreProfile(ctx, profile))
			},
			msg: types.NewMsgAcceptDTagTransferRequest(
				"NewDTag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expEvents: sdk.EmptyEvents(),
			shouldErr: true,
		},
		{
			name: "DTag is exchanged correctly (not existent sender profile)",
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"DTag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				)

				receiverProfile := testutil.ProfileFromAddr(request.Receiver)
				receiverProfile.DTag = "DTag"
				suite.Require().NoError(suite.k.StoreProfile(ctx, receiverProfile))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			msg: types.NewMsgAcceptDTagTransferRequest(
				"NewDtag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeDTagTransferAccept,
					sdk.NewAttribute(types.AttributeDTagToTrade, "DTag"),
					sdk.NewAttribute(types.AttributeNewDTag, "NewDtag"),
					sdk.NewAttribute(types.AttributeRequestSender, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeRequestReceiver, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
				),
			},
		},
		{
			name: "DTag exchanged correctly (already existent sender profile)",
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"DTag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				)
				receiverProfile := testutil.ProfileFromAddr(request.Receiver)
				receiverProfile.DTag = "DTag"
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(request.Sender)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, receiverProfile))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			msg: types.NewMsgAcceptDTagTransferRequest(
				"NewDtag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeDTagTransferAccept,
					sdk.NewAttribute(types.AttributeDTagToTrade, "DTag"),
					sdk.NewAttribute(types.AttributeNewDTag, "NewDtag"),
					sdk.NewAttribute(types.AttributeRequestSender, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeRequestReceiver, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
				),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			suite.k.SetParams(ctx, types.DefaultParams())
			if tc.store != nil {
				tc.store(ctx)
			}

			server := keeper.NewMsgServerImpl(suite.k)
			_, err := server.AcceptDTagTransferRequest(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServer_RefuseDTagTransfer() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgRefuseDTagTransferRequest
		shouldErr bool
		expEvents sdk.Events
	}{
		{
			name: "not found request returns error",
			msg: types.NewMsgRefuseDTagTransferRequest(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "found request is refused correctly",
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"dtag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(request.Receiver)))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			msg: types.NewMsgRefuseDTagTransferRequest(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeDTagTransferRefuse,
					sdk.NewAttribute(types.AttributeRequestSender, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeRequestReceiver, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
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

			server := keeper.NewMsgServerImpl(suite.k)
			_, err := server.RefuseDTagTransferRequest(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())
			}
		})
	}
}
