package keeper_test

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/x/profiles/keeper"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) Test_handleMsgRequestDTagTransfer() {
	otherAddr, err := sdk.AccAddressFromBech32(suite.testData.otherUser)
	suite.Require().NoError(err)

	otherAccAny, err := codectypes.NewAnyWithValue(authtypes.NewBaseAccountWithAddress(otherAddr))
	suite.Require().NoError(err)

	tests := []struct {
		name           string
		storedProfiles []*types.Profile
		storedDTagReqs []types.DTagTransferRequest
		storedBlocks   []types.UserBlock
		msg            *types.MsgRequestDTagTransfer
		shouldErr      bool
		expEvents      sdk.Events
	}{
		{
			name: "Blocked receiver making request returns error",
			storedBlocks: []types.UserBlock{
				types.NewUserBlock(
					suite.testData.otherUser,
					suite.testData.user,
					"This user has been blocked",
					"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				),
			},
			msg:       types.NewMsgRequestDTagTransfer(suite.testData.user, suite.testData.otherUser),
			expEvents: sdk.EmptyEvents(),
			shouldErr: true,
		},
		{
			name:      "No DTag to transfer returns error",
			msg:       types.NewMsgRequestDTagTransfer(suite.testData.user, suite.testData.otherUser),
			expEvents: sdk.EmptyEvents(),
			shouldErr: true,
		},
		{
			name: "Already present request returns error",
			storedProfiles: []*types.Profile{
				suite.testData.profile,
				{
					DTag:    "test-dtag",
					Account: otherAccAny,
				},
			},
			storedDTagReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.profile.GetAddress().String(), suite.testData.otherUser),
			},
			msg:       types.NewMsgRequestDTagTransfer(suite.testData.profile.GetAddress().String(), suite.testData.otherUser),
			expEvents: sdk.EmptyEvents(),
			shouldErr: true,
		},
		{
			name: "Not already present request saved correctly",
			storedProfiles: []*types.Profile{
				suite.testData.profile,
			},
			msg:       types.NewMsgRequestDTagTransfer(suite.testData.user, suite.testData.profile.GetAddress().String()),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeDTagTransferRequest,
					sdk.NewAttribute(types.AttributeDTagToTrade, "dtag"),
					sdk.NewAttribute(types.AttributeRequestSender, suite.testData.user),
					sdk.NewAttribute(types.AttributeRequestReceiver, suite.testData.profile.GetAddress().String()),
				),
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, req := range test.storedDTagReqs {
				err := suite.k.SaveDTagTransferRequest(suite.ctx, req)
				suite.Require().NoError(err)
			}

			for _, profile := range test.storedProfiles {
				err := suite.k.StoreProfile(suite.ctx, profile)
				suite.Require().NoError(err)
			}

			for _, block := range test.storedBlocks {
				err := suite.k.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			server := keeper.NewMsgServerImpl(suite.k)
			_, err = server.RequestDTagTransfer(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgAcceptDTagTransfer() {
	otherAddr, err := sdk.AccAddressFromBech32(suite.testData.otherUser)
	suite.Require().NoError(err)

	newAddr, err := sdk.AccAddressFromBech32("cosmos1lkqrqrns0ekttzrs678thh5f4prcgasthqcxph")
	suite.Require().NoError(err)

	tests := []struct {
		name           string
		storedDTagReqs []types.DTagTransferRequest
		storedProfiles []*types.Profile
		msg            *types.MsgAcceptDTagTransfer
		shouldErr      bool
		expEvents      sdk.Events
	}{
		{
			name:      "No request made from the receiving user returns error",
			msg:       types.NewMsgAcceptDTagTransfer("newDtag", suite.testData.user, suite.testData.otherUser),
			expEvents: sdk.EmptyEvents(),
			shouldErr: true,
		},
		{
			name: "Return an error if the request receiver does not have a profile",
			storedDTagReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			msg:       types.NewMsgAcceptDTagTransfer("newDtag", suite.testData.user, suite.testData.otherUser),
			expEvents: sdk.EmptyEvents(),
			shouldErr: true,
		},
		{
			name: "Return an error if the new DTag has already been chosen by another user",
			storedProfiles: []*types.Profile{
				suite.CheckProfileNoError(types.NewProfile(
					"dtag",
					"",
					"",
					types.NewPictures("", ""),
					suite.ctx.BlockTime(),
					authtypes.NewBaseAccountWithAddress(otherAddr),
				)),
				suite.CheckProfileNoError(types.NewProfile(
					"newDtag",
					"",
					"",
					types.NewPictures("", ""),
					suite.ctx.BlockTime(),
					authtypes.NewBaseAccountWithAddress(newAddr),
				)),
			},
			storedDTagReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			msg:       types.NewMsgAcceptDTagTransfer("newDtag", suite.testData.user, suite.testData.otherUser),
			expEvents: sdk.EmptyEvents(),
			shouldErr: true,
		},
		{
			name: "Return an error when current owner DTag is different from the requested one",
			storedProfiles: []*types.Profile{
				suite.CheckProfileNoError(types.NewProfile(
					"dtag1",
					"",
					"",
					types.NewPictures("", ""),
					suite.ctx.BlockTime(),
					authtypes.NewBaseAccountWithAddress(otherAddr),
				)),
			},
			storedDTagReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			msg:       types.NewMsgAcceptDTagTransfer("newDtag", suite.testData.user, suite.testData.otherUser),
			expEvents: sdk.EmptyEvents(),
			shouldErr: true,
		},
		{
			name: "DTag exchanged correctly (not existent sender profile)",
			storedProfiles: []*types.Profile{
				suite.CheckProfileNoError(types.NewProfile(
					"dtag",
					"",
					"",
					types.NewPictures("", ""),
					suite.ctx.BlockTime(),
					authtypes.NewBaseAccountWithAddress(otherAddr),
				)),
			},
			storedDTagReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			msg:       types.NewMsgAcceptDTagTransfer("newDtag", suite.testData.user, suite.testData.otherUser),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeDTagTransferAccept,
					sdk.NewAttribute(types.AttributeDTagToTrade, "dtag"),
					sdk.NewAttribute(types.AttributeNewDTag, "newDtag"),
					sdk.NewAttribute(types.AttributeRequestSender, suite.testData.user),
					sdk.NewAttribute(types.AttributeRequestReceiver, suite.testData.otherUser),
				),
			},
		},
		{
			name: "DTag exchanged correctly (already existent sender profile)",
			storedProfiles: []*types.Profile{
				suite.CheckProfileNoError(types.NewProfile(
					"dtag",
					"",
					"",
					types.NewPictures("", ""),
					suite.ctx.BlockTime(),
					suite.testData.profile.GetAccount(),
				)),
				suite.CheckProfileNoError(types.NewProfile(
					"previous",
					"",
					"",
					types.NewPictures("", ""),
					suite.ctx.BlockTime(),
					authtypes.NewBaseAccountWithAddress(otherAddr),
				)),
			},
			storedDTagReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("previous", suite.testData.user, suite.testData.otherUser),
			},
			msg: types.NewMsgAcceptDTagTransfer("newDtag", suite.testData.user, suite.testData.otherUser),
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeDTagTransferAccept,
					sdk.NewAttribute(types.AttributeDTagToTrade, "previous"),
					sdk.NewAttribute(types.AttributeNewDTag, "newDtag"),
					sdk.NewAttribute(types.AttributeRequestSender, suite.testData.user),
					sdk.NewAttribute(types.AttributeRequestReceiver, suite.testData.otherUser),
				),
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			suite.k.SetParams(suite.ctx, types.DefaultParams())

			for _, req := range test.storedDTagReqs {
				err := suite.k.SaveDTagTransferRequest(suite.ctx, req)
				suite.Require().NoError(err)
			}

			for _, prof := range test.storedProfiles {
				err := suite.k.StoreProfile(suite.ctx, prof)
				suite.Require().NoError(err)
			}

			server := keeper.NewMsgServerImpl(suite.k)
			_, err = server.AcceptDTagTransfer(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgRefuseDTagRequest() {
	tests := []struct {
		name           string
		msg            *types.MsgRefuseDTagTransfer
		storedDTagReqs []types.DTagTransferRequest
		shouldErr      bool
		expEvents      sdk.Events
	}{
		{
			name:           "No requests found returns error",
			storedDTagReqs: nil,
			msg:            types.NewMsgRefuseDTagTransferRequest(suite.testData.user, suite.testData.otherUser),
			expEvents:      sdk.EmptyEvents(),
			shouldErr:      true,
		},
		{
			name: "Deletion runs correctly",
			storedDTagReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			msg:       types.NewMsgRefuseDTagTransferRequest(suite.testData.user, suite.testData.otherUser),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeDTagTransferRefuse,
					sdk.NewAttribute(types.AttributeRequestSender, suite.testData.user),
					sdk.NewAttribute(types.AttributeRequestReceiver, suite.testData.otherUser),
				),
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, req := range test.storedDTagReqs {
				err := suite.k.SaveDTagTransferRequest(suite.ctx, req)
				suite.Require().NoError(err)
			}

			server := keeper.NewMsgServerImpl(suite.k)
			_, err := server.RefuseDTagTransfer(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgCancelDTagRequest() {
	tests := []struct {
		name           string
		msg            *types.MsgCancelDTagTransfer
		storedDTagReqs []types.DTagTransferRequest
		shouldErr      bool
		expEvents      sdk.Events
	}{
		{
			name:           "No requests found returns error",
			storedDTagReqs: nil,
			msg:            types.NewMsgCancelDTagTransferRequest(suite.testData.user, suite.testData.otherUser),
			expEvents:      sdk.EmptyEvents(),
			shouldErr:      true,
		},
		{
			name: "Deletion runs correctly",
			storedDTagReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			msg:       types.NewMsgCancelDTagTransferRequest(suite.testData.user, suite.testData.otherUser),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeDTagTransferCancel,
					sdk.NewAttribute(types.AttributeRequestSender, suite.testData.user),
					sdk.NewAttribute(types.AttributeRequestReceiver, suite.testData.otherUser),
				),
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, req := range test.storedDTagReqs {
				err := suite.k.SaveDTagTransferRequest(suite.ctx, req)
				suite.Require().NoError(err)
			}

			server := keeper.NewMsgServerImpl(suite.k)
			_, err := server.CancelDTagTransfer(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())
			}
		})
	}
}
