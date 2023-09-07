package keeper_test

import (
	"fmt"

	"github.com/desmos-labs/desmos/v6/testutil/profilestesting"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/x/profiles/keeper"
	"github.com/desmos-labs/desmos/v6/x/profiles/types"
)

func (suite *KeeperTestSuite) TestMsgServer_RequestDTagTransfer() {
	testCases := []struct {
		name      string
		setup     func(ctx sdk.Context)
		store     func(ctx sdk.Context)
		msg       *types.MsgRequestDTagTransfer
		shouldErr bool
		expEvents sdk.Events
	}{
		{
			name: "blocked receiver making request returns error",
			setup: func(ctx sdk.Context) {
				suite.rk.EXPECT().HasUserBlocked(ctx, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", uint64(0)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.Require().NoError(suite.k.SaveProfile(ctx,
					profilestesting.ProfileFromAddr("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")))
				suite.Require().NoError(suite.k.SaveProfile(ctx,
					profilestesting.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")))
			},
			msg: types.NewMsgRequestDTagTransfer(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "no DTag to transfer returns error",
			setup: func(ctx sdk.Context) {
				suite.rk.EXPECT().HasUserBlocked(ctx, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", uint64(0)).Return(false)
			},
			msg: types.NewMsgRequestDTagTransfer(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "Already present request returns error",
			setup: func(ctx sdk.Context) {
				suite.rk.EXPECT().HasUserBlocked(ctx, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", uint64(0)).Return(false)
			},
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"dtag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				)
				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(request.Receiver)))
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
			setup: func(ctx sdk.Context) {
				suite.rk.EXPECT().HasUserBlocked(ctx, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", uint64(0)).Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")))
			},
			msg: types.NewMsgRequestDTagTransfer(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgRequestDTagTransfer{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
				),
				sdk.NewEvent(
					types.EventTypeDTagTransferRequest,
					sdk.NewAttribute(types.AttributeKeyDTagToTrade, fmt.Sprintf("%s-dtag", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")),
					sdk.NewAttribute(types.AttributeKeyRequestSender, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeKeyRequestReceiver, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
				),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setup != nil {
				tc.setup(ctx)
			}
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
		check     func(ctx sdk.Context)
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

				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(request.Receiver)))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			msg: types.NewMsgCancelDTagTransferRequest(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgCancelDTagTransferRequest{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
				),
				sdk.NewEvent(
					types.EventTypeDTagTransferCancel,
					sdk.NewAttribute(types.AttributeKeyRequestSender, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeKeyRequestReceiver, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
				),
			},
			check: func(ctx sdk.Context) {
				suite.Require().False(
					suite.k.HasDTagTransferRequest(ctx, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
				)
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

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServer_AcceptDTagTransfer() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		check     func(ctx sdk.Context)
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
				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(request.Sender)))
				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(request.Receiver)))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))

				profile := profilestesting.ProfileFromAddr("cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn")
				profile.DTag = "NewDTag"
				suite.Require().NoError(suite.k.SaveProfile(ctx, profile))
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
				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(request.Sender)))
				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(request.Receiver)))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))

				profile := profilestesting.ProfileFromAddr("cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn")
				profile.DTag = "dtag"
				suite.Require().NoError(suite.k.SaveProfile(ctx, profile))
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
			name: "invalid request receiver profile after exchanging returns error",
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"DTag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				)

				receiverProfile := profilestesting.ProfileFromAddr(request.Receiver)
				receiverProfile.DTag = "DTag"
				suite.Require().NoError(suite.k.SaveProfile(ctx, receiverProfile))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			msg: types.NewMsgAcceptDTagTransferRequest(
				"D",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "invalid request sender profile after exchanging returns error",
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"D",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				)

				receiverProfile := profilestesting.ProfileFromAddr(request.Receiver)
				receiverProfile.DTag = "D" // invalid DTag length
				suite.Require().NoError(suite.k.SaveProfile(ctx, receiverProfile))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			msg: types.NewMsgAcceptDTagTransferRequest(
				"NewDtag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "exchanged new DTag is registered by others returns error",
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"DTag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				)

				receiverProfile := profilestesting.ProfileFromAddr(request.Receiver)
				receiverProfile.DTag = "DTag"
				senderProfile := profilestesting.ProfileFromAddr(request.Sender)
				senderProfile.DTag = "senderDTag"
				otherProfile := profilestesting.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				otherProfile.DTag = "NewDTag"
				suite.Require().NoError(suite.k.SaveProfile(ctx, senderProfile))
				suite.Require().NoError(suite.k.SaveProfile(ctx, receiverProfile))
				suite.Require().NoError(suite.k.SaveProfile(ctx, otherProfile))

				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))

			},
			msg: types.NewMsgAcceptDTagTransferRequest(
				"NewDTag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
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

				receiverProfile := profilestesting.ProfileFromAddr(request.Receiver)
				receiverProfile.DTag = "DTag"
				suite.Require().NoError(suite.k.SaveProfile(ctx, receiverProfile))
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
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgAcceptDTagTransferRequest{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
				),
				sdk.NewEvent(
					types.EventTypeDTagTransferAccept,
					sdk.NewAttribute(types.AttributeKeyDTagToTrade, "DTag"),
					sdk.NewAttribute(types.AttributeKeyNewDTag, "NewDtag"),
					sdk.NewAttribute(types.AttributeKeyRequestSender, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeKeyRequestReceiver, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
				),
			},
		},
		{
			name: "DTag exchanged correctly (already existent sender profile))",
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"DTag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				)
				receiverProfile := profilestesting.ProfileFromAddr(request.Receiver)
				receiverProfile.DTag = "DTag"
				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(request.Sender)))
				suite.Require().NoError(suite.k.SaveProfile(ctx, receiverProfile))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			msg: types.NewMsgAcceptDTagTransferRequest(
				"NewDtag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgAcceptDTagTransferRequest{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
				),
				sdk.NewEvent(
					types.EventTypeDTagTransferAccept,
					sdk.NewAttribute(types.AttributeKeyDTagToTrade, "DTag"),
					sdk.NewAttribute(types.AttributeKeyNewDTag, "NewDtag"),
					sdk.NewAttribute(types.AttributeKeyRequestSender, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeKeyRequestReceiver, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
				),
			},
			check: func(ctx sdk.Context) {
				senderProfile, _, _ := suite.k.GetProfile(ctx, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				receiverProfile, _, _ := suite.k.GetProfile(ctx, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")

				suite.Require().NotEqual(senderProfile.DTag, "receiverDTag")
				suite.Require().Equal(receiverProfile.DTag, "NewDtag")
			},
		},
		{
			name: "DTag swapped correctly",
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"receiverDTag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				)
				receiverProfile := profilestesting.ProfileFromAddr(request.Receiver)
				receiverProfile.DTag = "receiverDTag"
				senderProfile := profilestesting.ProfileFromAddr(request.Sender)
				senderProfile.DTag = "senderDTag"
				suite.Require().NoError(suite.k.SaveProfile(ctx, senderProfile))
				suite.Require().NoError(suite.k.SaveProfile(ctx, receiverProfile))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			msg: types.NewMsgAcceptDTagTransferRequest(
				"senderDTag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgAcceptDTagTransferRequest{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
				),
				sdk.NewEvent(
					types.EventTypeDTagTransferAccept,
					sdk.NewAttribute(types.AttributeKeyDTagToTrade, "receiverDTag"),
					sdk.NewAttribute(types.AttributeKeyNewDTag, "senderDTag"),
					sdk.NewAttribute(types.AttributeKeyRequestSender, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeKeyRequestReceiver, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
				),
			},
			check: func(ctx sdk.Context) {
				senderProfile, _, _ := suite.k.GetProfile(ctx, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				receiverProfile, _, _ := suite.k.GetProfile(ctx, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")

				suite.Require().Equal(senderProfile.DTag, "receiverDTag")
				suite.Require().Equal(receiverProfile.DTag, "senderDTag")
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
				if tc.check != nil {
					tc.check(ctx)
				}
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
		check     func(ctx sdk.Context)
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
				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(request.Receiver)))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			msg: types.NewMsgRefuseDTagTransferRequest(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgRefuseDTagTransferRequest{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
				),
				sdk.NewEvent(
					types.EventTypeDTagTransferRefuse,
					sdk.NewAttribute(types.AttributeKeyRequestSender, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeKeyRequestReceiver, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
				),
			},
			check: func(ctx sdk.Context) {
				suite.Require().False(
					suite.k.HasDTagTransferRequest(ctx, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
				)
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

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
