package keeper_test

import (
	"context"
	"time"

	relationshipstypes "github.com/desmos-labs/desmos/x/relationships/types"

	"github.com/desmos-labs/desmos/x/profiles/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) Test_handleMsgSaveProfile() {
	tests := []struct {
		name             string
		existentProfiles []types.Profile
		msg              *types.MsgSaveProfile
		expProfiles      []types.Profile
		expErr           error
		expEvent         sdk.Event
	}{
		{
			name: "Profile saved (with no previous profile created)",
			msg: types.NewMsgSaveProfile(
				"custom_dtag",
				"my-moniker",
				"my-bio",
				"https://test.com/profile-picture",
				"https://test.com/cover-pic",
				suite.testData.profile.Creator,
			),
			expProfiles: []types.Profile{
				types.NewProfile(
					"custom_dtag",
					"my-moniker",
					"my-bio",
					types.NewPictures(
						"https://test.com/profile-picture",
						"https://test.com/cover-pic",
					),
					suite.testData.profile.CreationDate,
					suite.testData.profile.Creator,
				),
			},
			expEvent: sdk.NewEvent(
				types.EventTypeProfileSaved,
				sdk.NewAttribute(types.AttributeProfileDtag, "custom_dtag"),
				sdk.NewAttribute(types.AttributeProfileCreator, suite.testData.profile.Creator),
				sdk.NewAttribute(types.AttributeProfileCreationTime, suite.testData.profile.CreationDate.Format(time.RFC3339)),
			),
		},
		{
			name: "Profile saved (with previous profile created)",
			existentProfiles: []types.Profile{
				types.NewProfile(
					"test_dtag",
					"old-moniker",
					"old-biography",
					types.NewPictures(
						"https://test.com/old-profile-pic",
						"https://test.com/old-cover-pic",
					),
					suite.testData.profile.CreationDate,
					suite.testData.profile.Creator,
				),
			},
			msg: types.NewMsgSaveProfile(
				"other_dtag",
				"moniker",
				"biography",
				"https://test.com/profile-pic",
				"https://test.com/cover-pic",
				suite.testData.profile.Creator,
			),
			expProfiles: []types.Profile{
				types.NewProfile(
					"other_dtag",
					"moniker",
					"biography",
					types.NewPictures(
						"https://test.com/profile-pic",
						"https://test.com/cover-pic",
					),
					suite.testData.profile.CreationDate,
					suite.testData.profile.Creator,
				),
			},
			expEvent: sdk.NewEvent(
				types.EventTypeProfileSaved,
				sdk.NewAttribute(types.AttributeProfileDtag, "other_dtag"),
				sdk.NewAttribute(types.AttributeProfileCreator, suite.testData.profile.Creator),
				sdk.NewAttribute(types.AttributeProfileCreationTime, suite.testData.profile.CreationDate.Format(time.RFC3339)),
			),
		},
		{
			name: "Profile not edited because of the invalid profile picture",
			existentProfiles: []types.Profile{
				suite.testData.profile,
				types.NewProfile(
					"custom_dtag",
					"biography",
					"",
					types.NewPictures("", ""),
					suite.testData.profile.CreationDate,
					suite.testData.profile.Creator,
				),
			},
			msg: types.NewMsgSaveProfile(
				"custom_dtag",
				"",
				"",
				"invalid-pic",
				"",
				suite.testData.profile.Creator,
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid profile picture uri provided"),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			suite.ctx = suite.ctx.WithBlockTime(suite.testData.profile.CreationDate)
			suite.keeper.SetParams(suite.ctx, types.DefaultParams())

			for _, acc := range test.existentProfiles {
				err := suite.keeper.StoreProfile(suite.ctx, acc)
				suite.Require().NoError(err)
			}

			server := keeper.NewMsgServerImpl(suite.keeper)
			_, err := server.SaveProfile(context.Background(), test.msg)

			if test.expErr != nil {
				suite.Error(err)
				suite.Require().Equal(test.expErr.Error(), err.Error())
			}

			if test.expErr == nil {
				suite.Require().NoError(err)

				profs := suite.keeper.GetProfiles(suite.ctx)
				suite.Len(profs, len(test.expProfiles))
				for index, profile := range profs {
					suite.True(profile.Equal(test.expProfiles[index]))
				}

				// Check the events
				suite.Len(suite.ctx.EventManager().Events(), 1)
				suite.Contains(suite.ctx.EventManager().Events(), test.expEvent)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgDeleteProfile() {
	tests := []struct {
		name            string
		existentAccount *types.Profile
		msg             *types.MsgDeleteProfile
		expErr          error
	}{
		{
			name:            "Profile doesn't exists",
			existentAccount: nil,
			msg:             types.NewMsgDeleteProfile(suite.testData.profile.Creator),
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				"no profile associated with this address: cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
		},
		{
			name:            "Profile deleted successfully",
			existentAccount: &suite.testData.profile,
			msg:             types.NewMsgDeleteProfile(suite.testData.profile.Creator),
			expErr:          nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			if test.existentAccount != nil {
				err := suite.keeper.StoreProfile(suite.ctx, *test.existentAccount)
				suite.Require().NoError(err)
			}

			server := keeper.NewMsgServerImpl(suite.keeper)
			res, err := server.DeleteProfile(context.Background(), test.msg)

			if res == nil {
				suite.NotNil(err)
				suite.Require().Equal(test.expErr.Error(), err.Error())
			}
			if res != nil {
				// Check the events
				createAccountEv := sdk.NewEvent(
					types.EventTypeProfileDeleted,
					sdk.NewAttribute(types.AttributeProfileDtag, "dtag"),
					sdk.NewAttribute(types.AttributeProfileCreator, test.msg.Creator),
				)

				suite.Len(suite.ctx.EventManager().Events(), 1)
				suite.Contains(suite.ctx.EventManager().Events(), createAccountEv)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgRequestDTagTransfer() {
	tests := []struct {
		name           string
		msg            *types.MsgRequestDTagTransfer
		isBlocked      bool
		hasProfile     bool
		storedDTagReqs []types.DTagTransferRequest
		expErr         error
		expEvent       sdk.Event
	}{
		{
			name:      "Blocked receiver making request returns error",
			msg:       types.NewMsgRequestDTagTransfer(suite.testData.user, suite.testData.otherUser),
			isBlocked: true,
			expErr: sdkerrors.Wrapf(
				sdkerrors.ErrInvalidRequest,
				"The user with address %s has blocked you",
				suite.testData.user,
			),
		},
		{
			name:           "No DTag to transfer returns error",
			msg:            types.NewMsgRequestDTagTransfer(suite.testData.otherUser, suite.testData.user),
			storedDTagReqs: nil,
			hasProfile:     false,
			expErr: sdkerrors.Wrapf(
				sdkerrors.ErrInvalidRequest,
				"The user with address %s doesn't have a profile yet so their DTag cannot be transferred",
				suite.testData.otherUser,
			),
		},
		{
			name: "Already present request returns error",
			msg:  types.NewMsgRequestDTagTransfer(suite.testData.user, suite.testData.otherUser),
			storedDTagReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			hasProfile: true,
			expErr: sdkerrors.Wrapf(
				sdkerrors.ErrInvalidRequest,
				"the transfer request from %s to %s has already been made",
				suite.testData.otherUser, suite.testData.user,
			),
		},
		{
			name:           "Not already present request saved correctly",
			msg:            types.NewMsgRequestDTagTransfer(suite.testData.user, suite.testData.otherUser),
			storedDTagReqs: nil,
			hasProfile:     true,
			expErr:         nil,
			expEvent: sdk.NewEvent(
				types.EventTypeDTagTransferRequest,
				sdk.NewAttribute(types.AttributeDTagToTrade, "dtag"),
				sdk.NewAttribute(types.AttributeRequestReceiver, suite.testData.user),
				sdk.NewAttribute(types.AttributeRequestSender, suite.testData.otherUser),
			),
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, req := range test.storedDTagReqs {
				err := suite.keeper.SaveDTagTransferRequest(suite.ctx, req)
				suite.Require().NoError(err)
			}

			if test.hasProfile {
				suite.keeper.AssociateDtagWithAddress(suite.ctx, "dtag", suite.testData.user)
			}

			if test.isBlocked {
				err := suite.relKeeper.SaveUserBlock(suite.ctx, relationshipstypes.NewUserBlock(
					suite.testData.user,
					suite.testData.otherUser,
					"test",
					"",
				))
				suite.Require().NoError(err)
			}

			server := keeper.NewMsgServerImpl(suite.keeper)
			res, err := server.RequestDTagTransfer(context.Background(), test.msg)

			if res == nil {
				suite.NotNil(err)
				suite.Require().Equal(test.expErr.Error(), err.Error())
			}
			if res != nil {
				// Check the events
				createAccountEv := sdk.NewEvent(
					types.EventTypeDTagTransferRequest,
					sdk.NewAttribute(types.AttributeDTagToTrade, "dtag"),
					sdk.NewAttribute(types.AttributeRequestReceiver, suite.testData.user),
					sdk.NewAttribute(types.AttributeRequestSender, suite.testData.otherUser),
				)

				suite.Len(suite.ctx.EventManager().Events(), 1)
				suite.Contains(suite.ctx.EventManager().Events(), createAccountEv)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgAcceptDTagTransfer() {
	tests := []struct {
		name           string
		msg            *types.MsgAcceptDTagTransfer
		storedDTagReqs []types.DTagTransferRequest
		storedProfiles []types.Profile
		expErr         error
		expEvent       sdk.Event
	}{
		{
			name:   "No request made from the receiving user returns error",
			msg:    types.NewMsgAcceptDTagTransfer("newDtag", suite.testData.user, suite.testData.otherUser),
			expErr: sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "no request made from %s", suite.testData.otherUser),
		},
		{
			name: "No existent profile for the dTag owner returns error",
			storedDTagReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			msg:    types.NewMsgAcceptDTagTransfer("newDtag", suite.testData.user, suite.testData.otherUser),
			expErr: sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "profile of %s doesn't exist", suite.testData.user),
		},
		{
			name: "Already existent dTag for current owner edited profile returns error",
			msg:  types.NewMsgAcceptDTagTransfer("newDtag", suite.testData.user, suite.testData.otherUser),
			storedDTagReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			storedProfiles: []types.Profile{
				types.NewProfile(
					"dtag",
					"",
					"",
					types.NewPictures("", ""),
					suite.ctx.BlockTime(),
					suite.testData.user,
				),
				types.NewProfile(
					"newDtag",
					"",
					"",
					types.NewPictures("", ""),
					suite.ctx.BlockTime(),
					"cosmos1lkqrqrns0ekttzrs678thh5f4prcgasthqcxph",
				),
			},
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "a profile with dtag: newDtag has already been created"),
		},
		{
			name: "Current owner DTag different that the requested one",
			storedDTagReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			storedProfiles: []types.Profile{
				types.NewProfile(
					"dtag1",
					"",
					"",
					types.NewPictures("", ""),
					suite.ctx.BlockTime(),
					suite.testData.user,
				),
			},
			msg:    types.NewMsgAcceptDTagTransfer("newDtag", suite.testData.user, suite.testData.otherUser),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the owner's DTag is different from the one to be exchanged"),
		},
		{
			name:           "Dtag exchanged correctly (not existent receiver profile)",
			storedDTagReqs: []types.DTagTransferRequest{types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser)},
			storedProfiles: []types.Profile{
				types.NewProfile(
					"dtag",
					"",
					"",
					types.NewPictures("", ""),
					suite.ctx.BlockTime(),
					suite.testData.user,
				),
			},
			msg: types.NewMsgAcceptDTagTransfer("newDtag", suite.testData.user, suite.testData.otherUser),
			expEvent: sdk.NewEvent(
				types.EventTypeDTagTransferAccept,
				sdk.NewAttribute(types.AttributeDTagToTrade, "dtag"),
				sdk.NewAttribute(types.AttributeNewDTag, "newDtag"),
				sdk.NewAttribute(types.AttributeRequestReceiver, suite.testData.user),
				sdk.NewAttribute(types.AttributeRequestSender, suite.testData.otherUser),
			),
		},
		{
			name:           "Dtag exchanged correctly (already existent receiver profile)",
			storedDTagReqs: []types.DTagTransferRequest{types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser)},
			storedProfiles: []types.Profile{
				types.NewProfile(
					"dtag",
					"",
					"",
					types.NewPictures("", ""),
					suite.ctx.BlockTime(),
					suite.testData.user,
				),
				types.NewProfile(
					"previous",
					"",
					"",
					types.NewPictures("", ""),
					suite.ctx.BlockTime(),
					suite.testData.otherUser,
				),
			},
			msg: types.NewMsgAcceptDTagTransfer("newDtag", suite.testData.user, suite.testData.otherUser),
			expEvent: sdk.NewEvent(
				types.EventTypeDTagTransferAccept,
				sdk.NewAttribute(types.AttributeDTagToTrade, "dtag"),
				sdk.NewAttribute(types.AttributeNewDTag, "newDtag"),
				sdk.NewAttribute(types.AttributeRequestReceiver, suite.testData.user),
				sdk.NewAttribute(types.AttributeRequestSender, suite.testData.otherUser),
			),
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, req := range test.storedDTagReqs {
				err := suite.keeper.SaveDTagTransferRequest(suite.ctx, req)
				suite.Require().NoError(err)
			}

			for _, prof := range test.storedProfiles {
				err := suite.keeper.StoreProfile(suite.ctx, prof)
				suite.Require().NoError(err)
			}

			server := keeper.NewMsgServerImpl(suite.keeper)
			_, err := server.AcceptDTagTransfer(context.Background(), test.msg)

			if err == nil {
				// Check the events
				createAccountEv := sdk.NewEvent(
					types.EventTypeDTagTransferAccept,
					sdk.NewAttribute(types.AttributeDTagToTrade, "dtag"),
					sdk.NewAttribute(types.AttributeNewDTag, "newDtag"),
					sdk.NewAttribute(types.AttributeRequestReceiver, suite.testData.user),
					sdk.NewAttribute(types.AttributeRequestSender, suite.testData.otherUser),
				)

				suite.Len(suite.ctx.EventManager().Events(), 1)
				suite.Contains(suite.ctx.EventManager().Events(), createAccountEv)
			}

			if err != nil {
				suite.NotNil(err)
				suite.Require().Equal(test.expErr.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgRefuseDTagRequest() {
	tests := []struct {
		name           string
		msg            *types.MsgRefuseDTagTransfer
		storedDTagReqs []types.DTagTransferRequest
		expErr         error
		expEvent       sdk.Event
	}{
		{
			name:           "No requests found returns error",
			msg:            types.NewMsgRefuseDTagTransferRequest(suite.testData.user, suite.testData.otherUser),
			storedDTagReqs: nil,
			expErr:         sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "no requests to be deleted"),
		},
		{
			name: "Deletion runs correctly",
			msg:  types.NewMsgRefuseDTagTransferRequest(suite.testData.otherUser, suite.testData.user),
			storedDTagReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			expErr: nil,
			expEvent: sdk.NewEvent(
				types.EventTypeDTagTransferRefuse,
				sdk.NewAttribute(types.AttributeRequestReceiver, suite.testData.user),
				sdk.NewAttribute(types.AttributeRequestSender, suite.testData.otherUser),
			),
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, req := range test.storedDTagReqs {
				err := suite.keeper.SaveDTagTransferRequest(suite.ctx, req)
				suite.Require().NoError(err)
			}

			server := keeper.NewMsgServerImpl(suite.keeper)
			_, err := server.RefuseDTagTransfer(context.Background(), test.msg)

			if err == nil {
				// Check the data
				suite.Len(suite.ctx.EventManager().Events(), 1)
				suite.Contains(suite.ctx.EventManager().Events(), test.expEvent)
			}

			if err != nil {
				suite.NotNil(err)
				suite.Require().Equal(test.expErr.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgCancelDTagRequest() {
	tests := []struct {
		name           string
		msg            *types.MsgCancelDTagTransfer
		storedDTagReqs []types.DTagTransferRequest
		expErr         error
		expEvent       sdk.Event
	}{
		{
			name:           "No requests found returns error",
			msg:            types.NewMsgCancelDTagTransferRequest(suite.testData.user, suite.testData.otherUser),
			storedDTagReqs: nil,
			expErr:         sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "no requests to be deleted"),
		},
		{
			name: "Deletion runs correctly",
			msg:  types.NewMsgCancelDTagTransferRequest(suite.testData.otherUser, suite.testData.user),
			storedDTagReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			expErr: nil,
			expEvent: sdk.NewEvent(
				types.EventTypeDTagTransferCancel,
				sdk.NewAttribute(types.AttributeRequestReceiver, suite.testData.user),
				sdk.NewAttribute(types.AttributeRequestSender, suite.testData.otherUser),
			),
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, req := range test.storedDTagReqs {
				err := suite.keeper.SaveDTagTransferRequest(suite.ctx, req)
				suite.Require().NoError(err)
			}

			server := keeper.NewMsgServerImpl(suite.keeper)
			_, err := server.CancelDTagTransfer(context.Background(), test.msg)

			if err == nil {
				suite.Len(suite.ctx.EventManager().Events(), 1)
				suite.Contains(suite.ctx.EventManager().Events(), test.expEvent)
			}

			if err != nil {
				suite.NotNil(err)
				suite.Require().Equal(test.expErr.Error(), err.Error())
			}
		})
	}
}
