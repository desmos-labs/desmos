package keeper_test

import (
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	relationshipstypes "github.com/desmos-labs/desmos/x/staging/relationships/types"

	"github.com/desmos-labs/desmos/x/profiles/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) Test_handleMsgSaveProfile() {
	tests := []struct {
		name              string
		existentProfiles  []*types.Profile
		blockTime         time.Time
		msg               *types.MsgSaveProfile
		shouldErr         bool
		expEvents         sdk.Events
		expStoredProfiles []*types.Profile
	}{
		{
			name: "Profile saved (with no previous profile created)",
			msg: types.NewMsgSaveProfile(
				"custom_dtag",
				"my-moniker",
				"my-bio",
				"https://test.com/profile-picture",
				"https://test.com/cover-pic",
				suite.testData.profile.GetAddress().String(),
			),
			expStoredProfiles: []*types.Profile{
				suite.CheckProfileNoError(types.NewProfile(
					"custom_dtag",
					"my-moniker",
					"my-bio",
					types.NewPictures(
						"https://test.com/profile-picture",
						"https://test.com/cover-pic",
					),
					suite.testData.profile.CreationDate,
					suite.testData.profile.GetAccount(),
				)),
			},
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeProfileSaved,
					sdk.NewAttribute(types.AttributeProfileDTag, "custom_dtag"),
					sdk.NewAttribute(types.AttributeProfileCreator, suite.testData.profile.GetAddress().String()),
					sdk.NewAttribute(types.AttributeProfileCreationTime, suite.testData.profile.CreationDate.Format(time.RFC3339)),
				),
			},
		},
		{
			name:      "Profile saved (with previous profile created)",
			blockTime: suite.testData.profile.CreationDate,
			existentProfiles: []*types.Profile{
				suite.CheckProfileNoError(types.NewProfile(
					"test_dtag",
					"old-moniker",
					"old-biography",
					types.NewPictures(
						"https://test.com/old-profile-pic",
						"https://test.com/old-cover-pic",
					),
					suite.testData.profile.CreationDate,
					suite.testData.profile.GetAccount(),
				)),
			},
			msg: types.NewMsgSaveProfile(
				"other_dtag",
				"moniker",
				"biography",
				"https://test.com/profile-pic",
				"https://test.com/cover-pic",
				suite.testData.profile.GetAddress().String(),
			),
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeProfileSaved,
					sdk.NewAttribute(types.AttributeProfileDTag, "other_dtag"),
					sdk.NewAttribute(types.AttributeProfileCreator, suite.testData.profile.GetAddress().String()),
					sdk.NewAttribute(types.AttributeProfileCreationTime, suite.testData.profile.CreationDate.Format(time.RFC3339)),
				),
			},
			expStoredProfiles: []*types.Profile{
				suite.CheckProfileNoError(types.NewProfile(
					"other_dtag",
					"moniker",
					"biography",
					types.NewPictures(
						"https://test.com/profile-pic",
						"https://test.com/cover-pic",
					),
					suite.testData.profile.CreationDate,
					suite.testData.profile.GetAccount(),
				)),
			},
		},
		{
			name:      "Profile not edited because of the invalid profile picture",
			blockTime: suite.testData.profile.CreationDate,
			existentProfiles: []*types.Profile{
				suite.CheckProfileNoError(types.NewProfile(
					"custom_dtag",
					"biography",
					"",
					types.NewPictures("", ""),
					suite.testData.profile.CreationDate,
					suite.testData.profile.GetAccount(),
				)),
			},
			msg: types.NewMsgSaveProfile(
				"custom_dtag",
				"",
				"",
				"invalid-pic",
				"",
				suite.testData.profile.GetAddress().String(),
			),
			expEvents: sdk.EmptyEvents(),
			shouldErr: true,
			expStoredProfiles: []*types.Profile{
				suite.CheckProfileNoError(types.NewProfile(
					"custom_dtag",
					"biography",
					"",
					types.NewPictures("", ""),
					suite.testData.profile.CreationDate,
					suite.testData.profile.GetAccount(),
				)),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			suite.k.SetParams(suite.ctx, types.DefaultParams())

			suite.ctx = suite.ctx.WithBlockTime(test.blockTime)
			for _, acc := range test.existentProfiles {
				err := suite.k.StoreProfile(suite.ctx, acc)
				suite.Require().NoError(err)
			}

			server := keeper.NewMsgServerImpl(suite.k)
			_, err := server.SaveProfile(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())
			}

			stored := suite.k.GetProfiles(suite.ctx)
			suite.Require().Len(stored, len(test.expStoredProfiles))
			for _, profile := range stored {
				suite.Require().Contains(test.expStoredProfiles, profile)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgDeleteProfile() {
	tests := []struct {
		name           string
		storedProfiles []*types.Profile
		msg            *types.MsgDeleteProfile
		expErr         error
		expEvents      sdk.Events
	}{
		{
			name:           "Profile doesn't exists",
			storedProfiles: nil,
			msg:            types.NewMsgDeleteProfile(suite.testData.profile.GetAddress().String()),
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				"no profile associated with the following address found: cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			expEvents: sdk.EmptyEvents(),
		},
		{
			name: "Profile deleted successfully",
			storedProfiles: []*types.Profile{
				suite.testData.profile,
			},
			msg: types.NewMsgDeleteProfile(suite.testData.profile.GetAddress().String()),
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeProfileDeleted,
					sdk.NewAttribute(types.AttributeProfileCreator, suite.testData.profile.GetAddress().String()),
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, profile := range test.storedProfiles {
				err := suite.k.StoreProfile(suite.ctx, profile)
				suite.Require().NoError(err)
			}

			server := keeper.NewMsgServerImpl(suite.k)
			_, err := server.DeleteProfile(sdk.WrapSDKContext(suite.ctx), test.msg)

			suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())

			if test.expErr != nil {
				suite.Require().Error(err)
				suite.Require().Equal(test.expErr.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgRequestDTagTransfer() {
	otherAddr, err := sdk.AccAddressFromBech32(suite.testData.otherUser)
	suite.Require().NoError(err)

	otherAccAny, err := codectypes.NewAnyWithValue(authtypes.NewBaseAccountWithAddress(otherAddr))
	suite.Require().NoError(err)

	tests := []struct {
		name           string
		storedProfiles []*types.Profile
		storedDTagReqs []types.DTagTransferRequest
		storedBlocks   []relationshipstypes.UserBlock
		msg            *types.MsgRequestDTagTransfer
		shouldErr      bool
		expEvents      sdk.Events
	}{
		{
			name: "Blocked receiver making request returns error",
			storedBlocks: []relationshipstypes.UserBlock{
				relationshipstypes.NewUserBlock(
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
				err := suite.rk.SaveUserBlock(suite.ctx, block)
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
