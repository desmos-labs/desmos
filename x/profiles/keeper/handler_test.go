package keeper_test

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/profiles/keeper"
	"github.com/desmos-labs/desmos/x/profiles/types"
	"github.com/desmos-labs/desmos/x/relationships"
)

func (suite *KeeperTestSuite) Test_validateProfile() {
	tests := []struct {
		name    string
		profile types.Profile
		expErr  error
	}{
		{
			name: "Max moniker length exceeded",
			profile: types.NewProfile("custom_dtag", suite.testData.profile.Creator, suite.testData.profile.CreationDate).
				WithMoniker(newStrPtr(strings.Repeat("A", 1005))).
				WithBio(newStrPtr("my-bio")).
				WithPictures(
					newStrPtr("https://test.com/profile-picture"),
					newStrPtr("https://test.com/cover-pic"),
				),
			expErr: fmt.Errorf("profile moniker cannot exceed 1000 characters"),
		},
		{
			name: "Min moniker length not reached",
			profile: types.NewProfile("custom_dtag", suite.testData.profile.Creator, suite.testData.profile.CreationDate).
				WithMoniker(newStrPtr("m")).
				WithBio(newStrPtr("my-bio")).
				WithPictures(
					newStrPtr("https://test.com/profile-picture"),
					newStrPtr("https://test.com/cover-pic"),
				),
			expErr: fmt.Errorf("profile moniker cannot be less than 2 characters"),
		},
		{
			name: "Max bio length exceeded",
			profile: types.NewProfile("custom_dtag", suite.testData.profile.Creator, suite.testData.profile.CreationDate).
				WithMoniker(newStrPtr("moniker")).
				WithBio(newStrPtr(strings.Repeat("A", 1005))).
				WithPictures(
					newStrPtr("https://test.com/profile-picture"),
					newStrPtr("https://test.com/cover-pic"),
				),
			expErr: fmt.Errorf("profile biography cannot exceed 1000 characters"),
		},
		{
			name: "Invalid dtag doesn't match regEx",
			profile: types.NewProfile("custom.", suite.testData.profile.Creator, suite.testData.profile.CreationDate).
				WithMoniker(newStrPtr("moniker")).
				WithBio(newStrPtr(strings.Repeat("A", 1000))).
				WithPictures(
					newStrPtr("https://test.com/profile-picture"),
					newStrPtr("https://test.com/cover-pic"),
				),
			expErr: fmt.Errorf("invalid profile dtag, it should match the following regEx ^[A-Za-z0-9_]+$"),
		},
		{
			name: "Min dtag length not reached",
			profile: types.NewProfile("d", suite.testData.profile.Creator, suite.testData.profile.CreationDate).
				WithMoniker(newStrPtr("moniker")).
				WithBio(newStrPtr("my-bio")).
				WithPictures(
					newStrPtr("https://test.com/profile-picture"),
					newStrPtr("https://test.com/cover-pic"),
				),
			expErr: fmt.Errorf("profile dtag cannot be less than 3 characters"),
		},
		{
			name: "Max dtag length exceeded",
			profile: types.NewProfile("9YfrVVi3UEI1ymN7n6isScyHNSt30xG6Jn1EDxEXxWOn0voSMIKqLhHsBfnZoXE", suite.testData.profile.Creator, suite.testData.profile.CreationDate).
				WithMoniker(newStrPtr("moniker")).
				WithBio(newStrPtr("my-bio")).
				WithPictures(
					newStrPtr("https://test.com/profile-picture"),
					newStrPtr("https://test.com/cover-pic"),
				),
			expErr: fmt.Errorf("profile dtag cannot exceed 30 characters"),
		},
		{
			name: "Invalid profile pictures returns error",
			profile: types.NewProfile("dtag", suite.testData.profile.Creator, suite.testData.profile.CreationDate).
				WithMoniker(newStrPtr("moniker")).
				WithBio(newStrPtr("my-bio")).
				WithPictures(
					newStrPtr("pic"),
					newStrPtr("https://test.com/cover-pic"),
				),
			expErr: fmt.Errorf("invalid profile picture uri provided"),
		},
		{
			name: "Valid profile returns no error",
			profile: types.NewProfile("dtag", suite.testData.profile.Creator, suite.testData.profile.CreationDate).
				WithMoniker(newStrPtr("moniker")).
				WithBio(newStrPtr("my-bio")).
				WithPictures(
					newStrPtr("https://test.com/profile-picture"),
					newStrPtr("https://test.com/cover-pic"),
				),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() // reset
			suite.keeper.SetParams(suite.ctx, types.DefaultParams())
			actual := keeper.ValidateProfile(suite.ctx, suite.keeper, test.profile)
			suite.Equal(test.expErr, actual)
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgSaveProfile() {
	tests := []struct {
		name             string
		existentProfiles types.Profiles
		expAccount       *types.Profile
		msg              types.MsgSaveProfile
		expErr           error
		expProfiles      types.Profiles
		expEvent         sdk.Event
	}{
		{
			name: "Profile saved (with no previous profile created)",
			msg: types.NewMsgSaveProfile(
				"custom_dtag",
				newStrPtr("my-moniker"),
				newStrPtr("my-bio"),
				newStrPtr("https://test.com/profile-picture"),
				newStrPtr("https://test.com/cover-pic"),
				suite.testData.profile.Creator,
			),
			expProfiles: types.NewProfiles(
				types.NewProfile("custom_dtag", suite.testData.profile.Creator, suite.testData.profile.CreationDate).
					WithMoniker(newStrPtr("my-moniker")).
					WithBio(newStrPtr("my-bio")).
					WithPictures(
						newStrPtr("https://test.com/profile-picture"),
						newStrPtr("https://test.com/cover-pic"),
					),
			),
			expEvent: sdk.NewEvent(
				types.EventTypeProfileSaved,
				sdk.NewAttribute(types.AttributeProfileDtag, "custom_dtag"),
				sdk.NewAttribute(types.AttributeProfileCreator, suite.testData.profile.Creator.String()),
				sdk.NewAttribute(types.AttributeProfileCreationTime, suite.testData.profile.CreationDate.Format(time.RFC3339)),
			),
		},
		{
			name: "Profile saved (with previous profile created)",
			existentProfiles: types.NewProfiles(
				types.NewProfile("test_dtag", suite.testData.profile.Creator, suite.testData.profile.CreationDate).
					WithMoniker(newStrPtr("old-moniker")).
					WithBio(newStrPtr("old-biography")).
					WithPictures(
						newStrPtr("https://test.com/old-profile-pic"),
						newStrPtr("https://test.com/old-cover-pic"),
					),
			),
			msg: types.NewMsgSaveProfile(
				"other_dtag",
				newStrPtr("moniker"),
				newStrPtr("biography"),
				newStrPtr("https://test.com/profile-pic"),
				newStrPtr("https://test.com/cover-pic"),
				suite.testData.profile.Creator,
			),
			expProfiles: types.NewProfiles(
				types.NewProfile("other_dtag", suite.testData.profile.Creator, suite.testData.profile.CreationDate).
					WithMoniker(newStrPtr("moniker")).
					WithBio(newStrPtr("biography")).
					WithPictures(
						newStrPtr("https://test.com/profile-pic"),
						newStrPtr("https://test.com/cover-pic"),
					),
			),
			expEvent: sdk.NewEvent(
				types.EventTypeProfileSaved,
				sdk.NewAttribute(types.AttributeProfileDtag, "other_dtag"),
				sdk.NewAttribute(types.AttributeProfileCreator, suite.testData.profile.Creator.String()),
				sdk.NewAttribute(types.AttributeProfileCreationTime, suite.testData.profile.CreationDate.Format(time.RFC3339)),
			),
		},
		{
			name: "Profile not edited because of the invalid profile picture",
			existentProfiles: types.NewProfiles(
				suite.testData.profile,
				types.NewProfile("custom_dtag", suite.testData.profile.Creator, suite.testData.profile.CreationDate).
					WithBio(newStrPtr("biography")),
			),
			msg: types.NewMsgSaveProfile(
				"custom_dtag",
				nil,
				nil,
				newStrPtr("invalid-pic"),
				nil,
				suite.testData.profile.Creator,
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid profile picture uri provided"),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() //reset
			suite.ctx = suite.ctx.WithBlockTime(suite.testData.profile.CreationDate)

			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			suite.keeper.SetParams(suite.ctx, types.DefaultParams())
			if test.existentProfiles != nil {
				for _, acc := range test.existentProfiles {
					key := types.ProfileStoreKey(acc.Creator)
					store.Set(key, suite.keeper.Cdc.MustMarshalBinaryBare(acc))
					suite.keeper.AssociateDtagWithAddress(suite.ctx, acc.DTag, acc.Creator)
				}
			}

			handler := keeper.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			if test.expErr != nil {
				suite.Error(err)
				suite.Equal(test.expErr.Error(), err.Error())
			}

			if test.expErr == nil {
				suite.NoError(err)

				profiles := suite.keeper.GetProfiles(suite.ctx)
				suite.Len(profiles, len(test.expProfiles))
				for index, profile := range profiles {
					suite.True(profile.Equals(test.expProfiles[index]))
				}

				// Check the events
				suite.Len(res.Events, 1)
				suite.Contains(res.Events, test.expEvent)
			}

		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgDeleteProfile() {
	tests := []struct {
		name            string
		existentAccount *types.Profile
		msg             types.MsgDeleteProfile
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
			suite.SetupTest() // reset
			store := suite.ctx.KVStore(suite.keeper.StoreKey)

			if test.existentAccount != nil {
				key := types.ProfileStoreKey(test.existentAccount.Creator)
				store.Set(key, suite.keeper.Cdc.MustMarshalBinaryBare(&test.existentAccount))
				suite.keeper.AssociateDtagWithAddress(suite.ctx, test.existentAccount.DTag, test.existentAccount.Creator)
			}

			handler := keeper.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			if res == nil {
				suite.NotNil(err)
				suite.Equal(test.expErr.Error(), err.Error())
			}
			if res != nil {
				// Check the data
				suite.Equal(suite.keeper.Cdc.MustMarshalBinaryLengthPrefixed("dtag"), res.Data)

				// Check the events
				createAccountEv := sdk.NewEvent(
					types.EventTypeProfileDeleted,
					sdk.NewAttribute(types.AttributeProfileDtag, "dtag"),
					sdk.NewAttribute(types.AttributeProfileCreator, test.msg.Creator.String()),
				)

				suite.Len(res.Events, 1)
				suite.Contains(res.Events, createAccountEv)
			}

		})
	}
}

func (suite *KeeperTestSuite) TestCheckForBlockedUser() {
	user, _ := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	otherUser, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	userBlock := relationships.NewUserBlock(user, otherUser, "test", "")
	userBlock1 := relationships.NewUserBlock(otherUser, user, "test", "")

	tests := []struct {
		name       string
		user       sdk.AccAddress
		userBlocks []relationships.UserBlock
		expBool    bool
	}{
		{
			name:       "blocked user found returns true",
			user:       otherUser,
			userBlocks: []relationships.UserBlock{userBlock, userBlock1},
			expBool:    true,
		},
		{
			name:       "non blocked user not found returns false",
			user:       user,
			userBlocks: []relationships.UserBlock{userBlock},
			expBool:    false,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			res := keeper.CheckForBlockedUser(test.userBlocks, test.user)
			suite.Equal(test.expBool, res)
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgRequestDTagTransfer() {
	tests := []struct {
		name           string
		msg            types.MsgRequestDTagTransfer
		hasProfile     bool
		isBlocked      bool
		storedDTagReqs []types.DTagTransferRequest
		expErr         error
		expEvent       sdk.Event
	}{
		{
			name:      "Blocked receiver making request returns error",
			msg:       types.NewMsgRequestDTagTransfer(suite.testData.user, suite.testData.otherUser),
			isBlocked: true,
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("The user with address %s has been blocked from %s",
					suite.testData.otherUser, suite.testData.user),
			),
		},
		{
			name:           "No DTag to transfer returns error",
			msg:            types.NewMsgRequestDTagTransfer(suite.testData.otherUser, suite.testData.user),
			storedDTagReqs: nil,
			hasProfile:     false,
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("The user with address %s doesn't have a profile yet so their DTag cannot be transferred",
					suite.testData.otherUser)),
		},
		{
			name: "Already present request returns error",
			msg:  types.NewMsgRequestDTagTransfer(suite.testData.user, suite.testData.otherUser),
			storedDTagReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			hasProfile: true,
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("the transfer request from %s to %s has already been made",
					suite.testData.otherUser, suite.testData.user)),
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
				sdk.NewAttribute(types.AttributeRequestReceiver, suite.testData.user.String()),
				sdk.NewAttribute(types.AttributeRequestSender, suite.testData.otherUser.String()),
			),
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {

			if test.isBlocked {
				userBlock := relationships.NewUserBlock(suite.testData.user, suite.testData.otherUser, "test",
					"")
				_ = suite.keeper.RelKeeper.SaveUserBlock(suite.ctx, userBlock)
			}

			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			if test.storedDTagReqs != nil {
				store.Set(types.DtagTransferRequestStoreKey(suite.testData.user),
					suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedDTagReqs),
				)
			}

			if test.hasProfile {
				suite.keeper.AssociateDtagWithAddress(suite.ctx, "dtag", suite.testData.user)
			}

			handler := keeper.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			if res == nil {
				suite.NotNil(err)
				suite.Equal(test.expErr.Error(), err.Error())
			}
			if res != nil {
				// Check the data
				suite.Equal(suite.keeper.Cdc.MustMarshalBinaryLengthPrefixed(
					types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser)), res.Data,
				)

				// Check the events
				createAccountEv := sdk.NewEvent(
					types.EventTypeDTagTransferRequest,
					sdk.NewAttribute(types.AttributeDTagToTrade, "dtag"),
					sdk.NewAttribute(types.AttributeRequestReceiver, suite.testData.user.String()),
					sdk.NewAttribute(types.AttributeRequestSender, suite.testData.otherUser.String()),
				)

				suite.Len(res.Events, 1)
				suite.Contains(res.Events, createAccountEv)
			}

		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgAcceptDTagTransfer() {
	user, err := sdk.AccAddressFromBech32("cosmos1lkqrqrns0ekttzrs678thh5f4prcgasthqcxph")
	suite.NoError(err)

	tests := []struct {
		name                       string
		msg                        types.MsgAcceptDTagTransferRequest
		storedDTagReqs             []types.DTagTransferRequest
		storedOwnerProfile         *types.Profile
		storedReceivingUserProfile *types.Profile
		storedDtag                 *string
		expErr                     error
		expEvent                   sdk.Event
	}{
		{
			name:           "No request made from the receiving user returns error",
			msg:            types.NewMsgAcceptDTagTransfer("newDtag", suite.testData.user, suite.testData.otherUser),
			storedDTagReqs: nil,
			expErr:         sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("no request made from %s", suite.testData.otherUser)),
		},
		{
			name:           "No existent profile for the dTag owner returns error",
			msg:            types.NewMsgAcceptDTagTransfer("newDtag", suite.testData.user, suite.testData.otherUser),
			storedDTagReqs: []types.DTagTransferRequest{types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser)},
			expErr:         sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("profile of %s doesn't exist", suite.testData.user)),
		},
		{
			name:                       "Already existent dTag for current owner edited profile returns error",
			msg:                        types.NewMsgAcceptDTagTransfer("newDtag", suite.testData.user, suite.testData.otherUser),
			storedDTagReqs:             []types.DTagTransferRequest{types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser)},
			storedOwnerProfile:         newProfilePtr(types.NewProfile("dtag", suite.testData.user, suite.ctx.BlockTime())),
			storedReceivingUserProfile: nil,
			storedDtag:                 newStrPtr("newDtag"),
			expErr:                     sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "a profile with dtag: newDtag has already been created"),
		},
		{
			name:                       "Current owner DTag different that the requested one",
			msg:                        types.NewMsgAcceptDTagTransfer("newDtag", suite.testData.user, suite.testData.otherUser),
			storedDTagReqs:             []types.DTagTransferRequest{types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser)},
			storedOwnerProfile:         newProfilePtr(types.NewProfile("dtag1", suite.testData.user, suite.ctx.BlockTime())),
			storedReceivingUserProfile: nil,
			expErr:                     sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the owner's DTag is different from the one to be exchanged"),
		},
		{
			name:                       "Dtag exchanged correctly (not existent receiver profile)",
			msg:                        types.NewMsgAcceptDTagTransfer("newDtag", suite.testData.user, suite.testData.otherUser),
			storedDTagReqs:             []types.DTagTransferRequest{types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser)},
			storedOwnerProfile:         newProfilePtr(types.NewProfile("dtag", suite.testData.user, suite.ctx.BlockTime())),
			storedReceivingUserProfile: nil,
			expErr:                     nil,
			expEvent: sdk.NewEvent(
				types.EventTypeDTagTransferAccept,
				sdk.NewAttribute(types.AttributeDTagToTrade, "dtag"),
				sdk.NewAttribute(types.AttributeNewDTag, "newDtag"),
				sdk.NewAttribute(types.AttributeRequestReceiver, suite.testData.user.String()),
				sdk.NewAttribute(types.AttributeRequestSender, suite.testData.otherUser.String()),
			),
		},
		{
			name:                       "Dtag exchanged correctly (already existent receiver profile)",
			msg:                        types.NewMsgAcceptDTagTransfer("newDtag", suite.testData.user, suite.testData.otherUser),
			storedDTagReqs:             []types.DTagTransferRequest{types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser)},
			storedOwnerProfile:         newProfilePtr(types.NewProfile("dtag", suite.testData.user, suite.ctx.BlockTime())),
			storedReceivingUserProfile: newProfilePtr(types.NewProfile("previous", suite.testData.otherUser, suite.ctx.BlockTime())),
			expErr:                     nil,
			expEvent: sdk.NewEvent(
				types.EventTypeDTagTransferAccept,
				sdk.NewAttribute(types.AttributeDTagToTrade, "dtag"),
				sdk.NewAttribute(types.AttributeNewDTag, "newDtag"),
				sdk.NewAttribute(types.AttributeRequestReceiver, suite.testData.user.String()),
				sdk.NewAttribute(types.AttributeRequestSender, suite.testData.otherUser.String()),
			),
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			if test.storedDTagReqs != nil {
				store.Set(types.DtagTransferRequestStoreKey(suite.testData.user),
					suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedDTagReqs),
				)
			}

			if test.storedOwnerProfile != nil {
				suite.keeper.SaveProfile(suite.ctx, *test.storedOwnerProfile)
			}

			if test.storedReceivingUserProfile != nil {
				suite.keeper.SaveProfile(suite.ctx, *test.storedReceivingUserProfile)
			}

			if test.storedDtag != nil {
				suite.keeper.SaveProfile(suite.ctx, types.NewProfile(*test.storedDtag, user, suite.ctx.BlockTime()))
			}

			handler := keeper.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			if res == nil {
				suite.NotNil(err)
				suite.Equal(test.expErr.Error(), err.Error())
			}
			if res != nil {
				// Check the data
				suite.Equal(suite.keeper.Cdc.MustMarshalBinaryLengthPrefixed(test.storedOwnerProfile.DTag), res.Data)

				// Check the events
				createAccountEv := sdk.NewEvent(
					types.EventTypeDTagTransferAccept,
					sdk.NewAttribute(types.AttributeDTagToTrade, "dtag"),
					sdk.NewAttribute(types.AttributeNewDTag, "newDtag"),
					sdk.NewAttribute(types.AttributeRequestReceiver, suite.testData.user.String()),
					sdk.NewAttribute(types.AttributeRequestSender, suite.testData.otherUser.String()),
				)

				suite.Len(res.Events, 1)
				suite.Contains(res.Events, createAccountEv)
			}

		})
	}
}

func (suite *KeeperTestSuite) Test_deleteDTagTransferRequest() {
	tests := []struct {
		name           string
		owner          sdk.AccAddress
		sender         sdk.AccAddress
		storedDTagReqs []types.DTagTransferRequest
		expErr         error
		expEvent       sdk.Event
	}{
		{
			name:           "No requests found returns error",
			owner:          suite.testData.user,
			sender:         suite.testData.otherUser,
			storedDTagReqs: nil,
			expErr:         sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "no requests to be deleted"),
		},
		{
			name:           "Deletion runs correctly",
			owner:          suite.testData.user,
			sender:         suite.testData.otherUser,
			storedDTagReqs: []types.DTagTransferRequest{types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser)},
			expErr:         nil,
			expEvent: sdk.NewEvent(
				types.EventTypeDTagTransferRefuse,
				sdk.NewAttribute(types.AttributeRequestReceiver, suite.testData.user.String()),
				sdk.NewAttribute(types.AttributeRequestSender, suite.testData.otherUser.String())),
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			if test.storedDTagReqs != nil {
				store.Set(types.DtagTransferRequestStoreKey(suite.testData.user),
					suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedDTagReqs),
				)
			}

			res, err := keeper.DeleteDTagTransferRequest(suite.ctx, suite.keeper,
				test.owner, test.sender, types.EventTypeDTagTransferRefuse)

			if res == nil {
				suite.NotNil(err)
				suite.Equal(test.expErr.Error(), err.Error())
			}
			if res != nil {
				// Check the data
				suite.Equal(suite.keeper.Cdc.MustMarshalBinaryLengthPrefixed(test.sender), res.Data)
				suite.Len(res.Events, 1)
				suite.Contains(res.Events, test.expEvent)
			}

		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgRefuseDTagRequest() {
	tests := []struct {
		name           string
		msg            types.MsgRefuseDTagTransferRequest
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
			name:           "Deletion runs correctly",
			msg:            types.NewMsgRefuseDTagTransferRequest(suite.testData.otherUser, suite.testData.user),
			storedDTagReqs: []types.DTagTransferRequest{types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser)},
			expErr:         nil,
			expEvent: sdk.NewEvent(
				types.EventTypeDTagTransferRefuse,
				sdk.NewAttribute(types.AttributeRequestReceiver, suite.testData.user.String()),
				sdk.NewAttribute(types.AttributeRequestSender, suite.testData.otherUser.String())),
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			if test.storedDTagReqs != nil {
				store.Set(types.DtagTransferRequestStoreKey(suite.testData.user),
					suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedDTagReqs),
				)
			}

			handler := keeper.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			if res == nil {
				suite.NotNil(err)
				suite.Equal(test.expErr.Error(), err.Error())
			}
			if res != nil {
				// Check the data
				suite.Equal(suite.keeper.Cdc.MustMarshalBinaryLengthPrefixed(test.msg.Receiver), res.Data)
				suite.Len(res.Events, 1)
				suite.Contains(res.Events, test.expEvent)
			}

		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgCancelDTagRequest() {
	tests := []struct {
		name           string
		msg            types.MsgCancelDTagTransferRequest
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
			name:           "Deletion runs correctly",
			msg:            types.NewMsgCancelDTagTransferRequest(suite.testData.otherUser, suite.testData.user),
			storedDTagReqs: []types.DTagTransferRequest{types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser)},
			expErr:         nil,
			expEvent: sdk.NewEvent(
				types.EventTypeDTagTransferCancel,
				sdk.NewAttribute(types.AttributeRequestReceiver, suite.testData.user.String()),
				sdk.NewAttribute(types.AttributeRequestSender, suite.testData.otherUser.String())),
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			if test.storedDTagReqs != nil {
				store.Set(types.DtagTransferRequestStoreKey(suite.testData.user),
					suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedDTagReqs),
				)
			}

			handler := keeper.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			if res == nil {
				suite.NotNil(err)
				suite.Equal(test.expErr.Error(), err.Error())
			}
			if res != nil {
				// Check the data
				suite.Equal(suite.keeper.Cdc.MustMarshalBinaryLengthPrefixed(test.msg.Sender), res.Data)
				suite.Len(res.Events, 1)
				suite.Contains(res.Events, test.expEvent)
			}

		})
	}
}
