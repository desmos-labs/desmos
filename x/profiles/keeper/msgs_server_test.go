package keeper_test

import (
	"encoding/hex"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/x/profiles/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
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
				"my-nickname",
				"my-bio",
				"https://test.com/profile-picture",
				"https://test.com/cover-pic",
				suite.testData.profile.GetAddress().String(),
			),
			expStoredProfiles: []*types.Profile{
				suite.CheckProfileNoError(types.NewProfile(
					"custom_dtag",
					"my-nickname",
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
					"old-nickname",
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
				"nickname",
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
					"nickname",
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

func (suite *KeeperTestSuite) Test_handleMsgCreateRelationship() {
	tests := []struct {
		name                string
		storedBlock         []types.UserBlock
		storedRelationships []types.Relationship
		msg                 *types.MsgCreateRelationship
		expErr              bool
		expEvents           sdk.Events
		expRelationships    []types.Relationship
	}{
		{
			name: "Relationship sender blocked by receiver returns error",
			storedBlock: []types.UserBlock{
				types.NewUserBlock(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"test",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			msg: types.NewMsgCreateRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expErr: true,
		},
		{
			name: "Existing relationship returns error",
			storedRelationships: []types.Relationship{
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			msg: types.NewMsgCreateRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expErr: true,
		},
		{
			name: "Relationship has been saved correctly",
			msg: types.NewMsgCreateRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeRelationshipCreated,
					sdk.NewAttribute(types.AttributeRelationshipSender, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeRelationshipReceiver, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
					sdk.NewAttribute(types.AttributeRelationshipSubspace, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				),
			},
			expRelationships: []types.Relationship{
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, rel := range test.storedRelationships {
				err := suite.k.SaveRelationship(suite.ctx, rel)
				suite.Require().NoError(err)
			}

			for _, block := range test.storedBlock {
				err := suite.k.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			handler := keeper.NewMsgServerImpl(suite.k)
			_, err := handler.CreateRelationship(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())

				stored := suite.k.GetAllRelationships(suite.ctx)
				suite.Require().Equal(test.expRelationships, stored)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgDeleteRelationship() {
	tests := []struct {
		name             string
		stored           []types.Relationship
		msg              *types.MsgDeleteRelationship
		expErr           bool
		expEvents        sdk.Events
		expRelationships []types.Relationship
	}{
		{
			name: "Relationship not found returns error",
			stored: []types.Relationship{
				types.NewRelationship("creator", "recipient", "subspace"),
			},
			msg:    types.NewMsgDeleteRelationship("creator", "recipient", "other_subspace"),
			expErr: true,
		},
		{
			name: "Existing relationship is removed properly and leaves empty array",
			stored: []types.Relationship{
				types.NewRelationship("creator", "recipient", "subspace"),
			},
			msg:    types.NewMsgDeleteRelationship("creator", "recipient", "subspace"),
			expErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeRelationshipsDeleted,
					sdk.NewAttribute(types.AttributeRelationshipSender, "creator"),
					sdk.NewAttribute(types.AttributeRelationshipReceiver, "recipient"),
					sdk.NewAttribute(types.AttributeRelationshipSubspace, "subspace"),
				),
			},
		},
		{
			name: "Existing relationship is removed properly and leaves not empty array",
			stored: []types.Relationship{
				types.NewRelationship("creator", "recipient", "subspace"),
				types.NewRelationship("creator", "recipient", "other_subspace"),
			},
			msg:    types.NewMsgDeleteRelationship("creator", "recipient", "subspace"),
			expErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeRelationshipsDeleted,
					sdk.NewAttribute(types.AttributeRelationshipSender, "creator"),
					sdk.NewAttribute(types.AttributeRelationshipReceiver, "recipient"),
					sdk.NewAttribute(types.AttributeRelationshipSubspace, "subspace"),
				),
			},
			expRelationships: []types.Relationship{
				types.NewRelationship("creator", "recipient", "other_subspace"),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, relationship := range test.stored {
				err := suite.k.SaveRelationship(suite.ctx, relationship)
				suite.Require().NoError(err)
			}

			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.DeleteRelationship(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())

				left := suite.k.GetAllRelationships(suite.ctx)
				suite.Require().Equal(test.expRelationships, left)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgBlockUser() {
	tests := []struct {
		name      string
		msg       *types.MsgBlockUser
		stored    []types.UserBlock
		expErr    bool
		expEvents sdk.Events
		expBlocks []types.UserBlock
	}{
		{
			name: "Existing relationship returns error",
			stored: []types.UserBlock{
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"reason",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			msg: types.NewMsgBlockUser(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"reason",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expErr: true,
		},
		{
			name:   "Block has been saved correctly",
			stored: nil,
			msg: types.NewMsgBlockUser(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"reason",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeBlockUser,
					sdk.NewAttribute(types.AttributeKeyUserBlockBlocker, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeKeyUserBlockBlocked, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
					sdk.NewAttribute(types.AttributeKeySubspace, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
					sdk.NewAttribute(types.AttributeKeyUserBlockReason, "reason"),
				),
			},
			expBlocks: []types.UserBlock{
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"reason",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, block := range test.stored {
				err := suite.k.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.BlockUser(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())

				blocks := suite.k.GetUserBlocks(suite.ctx, test.msg.Blocker)
				suite.Require().Equal(test.expBlocks, blocks)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgUnblockUser() {
	tests := []struct {
		name        string
		storedBlock []types.UserBlock
		msg         *types.MsgUnblockUser
		expErr      bool
		expEvents   sdk.Events
		expBlocks   []types.UserBlock
	}{
		{
			name:        "Invalid block returns error",
			storedBlock: []types.UserBlock{},
			msg:         types.NewMsgUnblockUser("blocker", "blocked", "subspace"),
			expErr:      true,
		},
		{
			name: "Existing block is removed and leaves empty array",
			storedBlock: []types.UserBlock{
				types.NewUserBlock("blocker", "blocked", "reason", "subspace"),
			},
			msg:    types.NewMsgUnblockUser("blocker", "blocked", "subspace"),
			expErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeUnblockUser,
					sdk.NewAttribute(types.AttributeKeyUserBlockBlocker, "blocker"),
					sdk.NewAttribute(types.AttributeKeyUserBlockBlocked, "blocked"),
					sdk.NewAttribute(types.AttributeKeySubspace, "subspace"),
				),
			},
			expBlocks: nil,
		},
		{
			name: "Existing block is removed and leaves non empty array",
			storedBlock: []types.UserBlock{
				types.NewUserBlock("blocker", "blocked", "reason", "subspace"),
				types.NewUserBlock("blocker", "blocked", "reason", "other_subspace"),
			},
			msg:    types.NewMsgUnblockUser("blocker", "blocked", "subspace"),
			expErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeUnblockUser,
					sdk.NewAttribute(types.AttributeKeyUserBlockBlocker, "blocker"),
					sdk.NewAttribute(types.AttributeKeyUserBlockBlocked, "blocked"),
					sdk.NewAttribute(types.AttributeKeySubspace, "subspace"),
				),
			},
			expBlocks: []types.UserBlock{
				types.NewUserBlock("blocker", "blocked", "reason", "other_subspace"),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, block := range test.storedBlock {
				err := suite.k.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.UnblockUser(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())

				stored := suite.k.GetAllUsersBlocks(suite.ctx)
				suite.Require().Equal(test.expBlocks, stored)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgLink() {
	var existentLinks []types.Link
	var existentProfiles []*types.Profile

	// Generate source and destination key
	srcPriv := secp256k1.GenPrivKey()
	destPriv := secp256k1.GenPrivKey()
	srcPubKey := srcPriv.PubKey()
	destPubKey := destPriv.PubKey()

	// Get bech32 encoded addresses
	srcAddr, err := bech32.ConvertAndEncode("cosmos", srcPubKey.Address().Bytes())
	suite.Require().NoError(err)
	destAddr, err := bech32.ConvertAndEncode("cosmos", destPubKey.Address().Bytes())
	suite.Require().NoError(err)

	// Get signature by signing with keys
	srcSig, err := srcPriv.Sign([]byte(srcAddr))
	suite.Require().NoError(err)
	destSig, err := destPriv.Sign(srcSig)
	suite.Require().NoError(err)

	srcSigHex := hex.EncodeToString(srcSig)
	destSigHex := hex.EncodeToString(destSig)

	blockTime := suite.testData.profile.CreationDate

	tests := []struct {
		name           string
		malleate       func()
		msg            *types.MsgLink
		shouldErr      bool
		expEvents      sdk.Events
		expStoredLinks []types.Link
	}{
		{
			name:     "Invalid msg returns error",
			malleate: func() {},
			msg: types.NewMsgLink(
				"",
				destAddr,
				srcSigHex,
				destSigHex,
			),
			shouldErr:      true,
			expEvents:      sdk.EmptyEvents(),
			expStoredLinks: nil,
		},
		{
			name:     "Invalid source address returns error (unknown source address)",
			malleate: func() {},
			msg: types.NewMsgLink(
				srcAddr,
				destAddr,
				srcSigHex,
				destSigHex,
			),
			shouldErr:      true,
			expEvents:      sdk.EmptyEvents(),
			expStoredLinks: nil,
		},
		{
			name: "Invalid source address returns error (non existent pubkey on source address)",
			malleate: func() {
				srcAccAddr, err := sdk.AccAddressFromBech32(srcAddr)
				suite.Require().NoError(err)

				srcBaseAcc := authtypes.NewBaseAccountWithAddress(srcAccAddr)
				suite.ak.SetAccount(suite.ctx, srcBaseAcc)
			},
			msg: types.NewMsgLink(
				srcAddr,
				destAddr,
				srcSigHex,
				destSigHex,
			),
			shouldErr:      true,
			expEvents:      sdk.EmptyEvents(),
			expStoredLinks: nil,
		},
		{
			name: "Invalid destination address returns error (unknown dest address)",
			malleate: func() {
				srcAccAddr, err := sdk.AccAddressFromBech32(srcAddr)
				suite.Require().NoError(err)

				srcBaseAcc := authtypes.NewBaseAccountWithAddress(srcAccAddr)
				srcBaseAcc.SetPubKey(srcPubKey)
				suite.ak.SetAccount(suite.ctx, srcBaseAcc)
			},
			msg: types.NewMsgLink(
				srcAddr,
				destAddr,
				srcSigHex,
				destSigHex,
			),
			shouldErr:      true,
			expEvents:      sdk.EmptyEvents(),
			expStoredLinks: nil,
		},
		{
			name: "Invalid destination address returns error (non existent pubkey on source address)",
			malleate: func() {
				srcAccAddr, err := sdk.AccAddressFromBech32(srcAddr)
				suite.Require().NoError(err)

				srcBaseAcc := authtypes.NewBaseAccountWithAddress(srcAccAddr)
				srcBaseAcc.SetPubKey(srcPubKey)
				suite.ak.SetAccount(suite.ctx, srcBaseAcc)

				destAccAddr, err := sdk.AccAddressFromBech32(destAddr)
				suite.Require().NoError(err)
				destBaseAcc := authtypes.NewBaseAccountWithAddress(destAccAddr)
				suite.ak.SetAccount(suite.ctx, destBaseAcc)
			},
			msg: types.NewMsgLink(
				srcAddr,
				destAddr,
				srcSigHex,
				destSigHex,
			),
			shouldErr:      true,
			expEvents:      sdk.EmptyEvents(),
			expStoredLinks: nil,
		},
		{
			name: "Invalid destination signature",
			malleate: func() {
				srcAccAddr, err := sdk.AccAddressFromBech32(srcAddr)
				suite.Require().NoError(err)

				srcBaseAcc := authtypes.NewBaseAccountWithAddress(srcAccAddr)
				srcBaseAcc.SetPubKey(srcPubKey)
				suite.ak.SetAccount(suite.ctx, srcBaseAcc)

				destAccAddr, err := sdk.AccAddressFromBech32(destAddr)
				suite.Require().NoError(err)
				destBaseAcc := authtypes.NewBaseAccountWithAddress(destAccAddr)
				destBaseAcc.SetPubKey(destPubKey)
				suite.ak.SetAccount(suite.ctx, destBaseAcc)
			},
			msg: types.NewMsgLink(
				srcAddr,
				destAddr,
				srcSigHex,
				srcSigHex,
			),
			shouldErr:      true,
			expEvents:      sdk.EmptyEvents(),
			expStoredLinks: nil,
		},
		{
			name: "Invalid source signature",
			malleate: func() {
				srcAccAddr, err := sdk.AccAddressFromBech32(srcAddr)
				suite.Require().NoError(err)

				srcBaseAcc := authtypes.NewBaseAccountWithAddress(srcAccAddr)
				srcBaseAcc.SetPubKey(destPubKey) // wrong public key
				suite.ak.SetAccount(suite.ctx, srcBaseAcc)

				destAccAddr, err := sdk.AccAddressFromBech32(destAddr)
				suite.Require().NoError(err)
				destBaseAcc := authtypes.NewBaseAccountWithAddress(destAccAddr)
				destBaseAcc.SetPubKey(destPubKey)
				suite.ak.SetAccount(suite.ctx, destBaseAcc)
			},
			msg: types.NewMsgLink(
				srcAddr,
				destAddr,
				srcSigHex,
				destSigHex,
			),
			shouldErr:      true,
			expEvents:      sdk.EmptyEvents(),
			expStoredLinks: nil,
		},
		{
			name: "existent profile on source address returns error",
			malleate: func() {
				srcAccAddr, err := sdk.AccAddressFromBech32(srcAddr)
				suite.Require().NoError(err)

				srcBaseAcc := authtypes.NewBaseAccountWithAddress(srcAccAddr)
				srcBaseAcc.SetPubKey(srcPubKey)
				suite.ak.SetAccount(suite.ctx, srcBaseAcc)

				destAccAddr, err := sdk.AccAddressFromBech32(destAddr)
				suite.Require().NoError(err)
				destBaseAcc := authtypes.NewBaseAccountWithAddress(destAccAddr)
				destBaseAcc.SetPubKey(destPubKey)
				suite.ak.SetAccount(suite.ctx, destBaseAcc)

				existentProfiles = []*types.Profile{
					suite.CheckProfileNoError(types.NewProfile(
						"custom_dtag",
						"my-nickname",
						"my-bio",
						types.NewPictures(
							"https://test.com/profile-picture",
							"https://test.com/cover-pic",
						),
						suite.testData.profile.CreationDate,
						srcBaseAcc,
					)),
				}
			},
			msg: types.NewMsgLink(
				srcAddr,
				destAddr,
				srcSigHex,
				destSigHex,
			),
			shouldErr:      true,
			expEvents:      sdk.EmptyEvents(),
			expStoredLinks: nil,
		},
		{
			name: "non existent profile on destination address returns error",
			malleate: func() {
				srcAccAddr, err := sdk.AccAddressFromBech32(srcAddr)
				suite.Require().NoError(err)

				srcBaseAcc := authtypes.NewBaseAccountWithAddress(srcAccAddr)
				srcBaseAcc.SetPubKey(srcPubKey)
				suite.ak.SetAccount(suite.ctx, srcBaseAcc)

				destAccAddr, err := sdk.AccAddressFromBech32(destAddr)
				suite.Require().NoError(err)
				destBaseAcc := authtypes.NewBaseAccountWithAddress(destAccAddr)
				destBaseAcc.SetPubKey(destPubKey)
				suite.ak.SetAccount(suite.ctx, destBaseAcc)
			},
			msg: types.NewMsgLink(
				srcAddr,
				destAddr,
				srcSigHex,
				destSigHex,
			),
			shouldErr:      true,
			expEvents:      sdk.EmptyEvents(),
			expStoredLinks: nil,
		},
		{
			name: "Link already exists returns error",
			malleate: func() {
				srcAccAddr, err := sdk.AccAddressFromBech32(srcAddr)
				suite.Require().NoError(err)

				srcBaseAcc := authtypes.NewBaseAccountWithAddress(srcAccAddr)
				srcBaseAcc.SetPubKey(srcPubKey)
				suite.ak.SetAccount(suite.ctx, srcBaseAcc)

				destAccAddr, err := sdk.AccAddressFromBech32(destAddr)
				suite.Require().NoError(err)
				destBaseAcc := authtypes.NewBaseAccountWithAddress(destAccAddr)
				destBaseAcc.SetPubKey(destPubKey)
				suite.ak.SetAccount(suite.ctx, destBaseAcc)

				existentProfiles = []*types.Profile{
					suite.CheckProfileNoError(types.NewProfile(
						"custom_dtag",
						"my-nickname",
						"my-bio",
						types.NewPictures(
							"https://test.com/profile-picture",
							"https://test.com/cover-pic",
						),
						suite.testData.profile.CreationDate,
						destBaseAcc,
					)),
				}

				existentLinks = []types.Link{
					types.NewLink(
						srcAddr,
						types.NewProof("pubKey", "signature"),
						types.NewChainConfig(suite.ctx.ChainID(), sdk.GetConfig().GetBech32AccountAddrPrefix()),
						time.Time{},
					),
				}
			},
			msg: types.NewMsgLink(
				srcAddr,
				destAddr,
				srcSigHex,
				destSigHex,
			),
			shouldErr:      true,
			expEvents:      sdk.EmptyEvents(),
			expStoredLinks: nil,
		},
		{
			name: "Create link successfully",
			malleate: func() {
				srcAccAddr, err := sdk.AccAddressFromBech32(srcAddr)
				suite.Require().NoError(err)

				srcBaseAcc := authtypes.NewBaseAccountWithAddress(srcAccAddr)
				srcBaseAcc.SetPubKey(srcPubKey)
				suite.ak.SetAccount(suite.ctx, srcBaseAcc)

				destAccAddr, err := sdk.AccAddressFromBech32(destAddr)
				suite.Require().NoError(err)
				destBaseAcc := authtypes.NewBaseAccountWithAddress(destAccAddr)
				destBaseAcc.SetPubKey(destPubKey)
				suite.ak.SetAccount(suite.ctx, destBaseAcc)

				existentProfiles = []*types.Profile{
					suite.CheckProfileNoError(types.NewProfile(
						"custom_dtag",
						"my-nickname",
						"my-bio",
						types.NewPictures(
							"https://test.com/profile-picture",
							"https://test.com/cover-pic",
						),
						suite.testData.profile.CreationDate,
						destBaseAcc,
					)),
				}

				existentLinks = nil
			},
			msg: types.NewMsgLink(
				srcAddr,
				destAddr,
				srcSigHex,
				destSigHex,
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeLink,
					sdk.NewAttribute(types.AttributeLinkSourceAddress, srcAddr),
					sdk.NewAttribute(types.AttributeLinkDestinationAddress, destAddr),
					sdk.NewAttribute(types.AttributeLinkSourceSignature, srcSigHex),
					sdk.NewAttribute(types.AttributeLinkDestinationSignature, destSigHex),
				),
			},
			expStoredLinks: []types.Link{
				types.NewLink(
					srcAddr,
					types.NewProof(hex.EncodeToString(srcPubKey.Bytes()), srcSigHex),
					types.NewChainConfig(suite.ctx.ChainID(), sdk.GetConfig().GetBech32AccountAddrPrefix()),
					blockTime,
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			test.malleate()

			suite.ctx = suite.ctx.WithBlockTime(blockTime)
			for _, acc := range existentProfiles {
				err := suite.k.StoreProfile(suite.ctx, acc)
				suite.Require().NoError(err)
			}

			for _, link := range existentLinks {
				err := suite.k.StoreLink(suite.ctx, link)
				suite.Require().NoError(err)
			}

			server := keeper.NewMsgServerImpl(suite.k)
			_, err := server.Link(sdk.WrapSDKContext(suite.ctx), test.msg)
			suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())

			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				stored := suite.k.GetAllLinks(suite.ctx)
				suite.Require().NotEmpty(stored)
				for _, link := range stored {
					suite.Require().Contains(test.expStoredLinks, link)
				}
				profile, found, err := suite.k.GetProfile(suite.ctx, destAddr)
				suite.Require().NoError(err)
				suite.Require().True(found)

				suite.Require().NotEmpty(profile.Links)
				for _, link := range profile.Links {
					suite.Require().Contains(test.expStoredLinks, link)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgUnlink() {
	link := types.NewLink(
		suite.testData.otherUser,
		types.NewProof("pubKey", "signature"),
		types.NewChainConfig("test-net", "cosmos"),
		time.Time{},
	)
	validProfile := *suite.testData.profile
	validProfile.Links = []types.Link{
		link,
	}
	tests := []struct {
		name             string
		msg              *types.MsgUnlink
		shouldErr        bool
		expEvents        sdk.Events
		existentProfiles []*types.Profile
		existentLinks    []types.Link
		expStoredLinks   []types.Link
	}{
		{
			name:             "Invalid msg returns error",
			msg:              types.NewMsgUnlink("", "", ""),
			shouldErr:        true,
			expEvents:        sdk.EmptyEvents(),
			existentProfiles: nil,
			existentLinks:    nil,
			expStoredLinks:   nil,
		},
		{
			name:             "Non existent profile on owner returns error",
			msg:              types.NewMsgUnlink(suite.testData.user, "test-net", suite.testData.otherUser),
			shouldErr:        true,
			expEvents:        sdk.EmptyEvents(),
			existentProfiles: nil,
			existentLinks:    nil,
			expStoredLinks:   nil,
		},
		{
			name:             "Non existent target link in the profile returns error",
			msg:              types.NewMsgUnlink(suite.testData.user, "test-net", suite.testData.otherUser),
			shouldErr:        true,
			expEvents:        sdk.EmptyEvents(),
			existentProfiles: []*types.Profile{suite.testData.profile},
			existentLinks:    nil,
			expStoredLinks:   nil,
		},
		{
			name:      "No error message",
			msg:       types.NewMsgUnlink(suite.testData.user, "test-net", suite.testData.otherUser),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeUnlink,
					sdk.NewAttribute(types.AttributeUnlinkOwner, suite.testData.user),
					sdk.NewAttribute(types.AttributeUnlinkChainID, "test-net"),
					sdk.NewAttribute(types.AttributeUnlinkTarget, suite.testData.otherUser),
				),
			},
			existentProfiles: []*types.Profile{&validProfile},
			existentLinks: []types.Link{
				link,
				types.NewLink(
					"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq",
					types.NewProof("pubkey", "signature"),
					types.NewChainConfig("test-net", "cosmos"),
					time.Time{},
				),
			},
			expStoredLinks: []types.Link{
				types.NewLink(
					"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq",
					types.NewProof("pubkey", "signature"),
					types.NewChainConfig("test-net", "cosmos"),
					time.Time{},
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			for _, acc := range test.existentProfiles {
				err := suite.k.StoreProfile(suite.ctx, acc)
				suite.Require().NoError(err)
			}

			for _, link := range test.existentLinks {
				err := suite.k.StoreLink(suite.ctx, link)
				suite.Require().NoError(err)
			}

			server := keeper.NewMsgServerImpl(suite.k)
			_, err := server.Unlink(sdk.WrapSDKContext(suite.ctx), test.msg)
			suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())

			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				stored := suite.k.GetAllLinks(suite.ctx)
				suite.Require().Equal(test.expStoredLinks, stored)

				profile, found, err := suite.k.GetProfile(suite.ctx, suite.testData.user)
				suite.Require().NoError(err)
				suite.Require().True(found)
				suite.Require().Empty(profile.Links)
			}
		})
	}
}
