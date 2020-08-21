package keeper_test

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/profiles/keeper"
	"github.com/desmos-labs/desmos/x/profiles/types"
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
				"test_dtag",
				newStrPtr("moniker"),
				newStrPtr("biography"),
				newStrPtr("https://test.com/profile-pic"),
				newStrPtr("https://test.com/cover-pic"),
				suite.testData.profile.Creator,
			),
			expProfiles: types.NewProfiles(
				types.NewProfile("test_dtag", suite.testData.profile.Creator, suite.testData.profile.CreationDate).
					WithMoniker(newStrPtr("moniker")).
					WithBio(newStrPtr("biography")).
					WithPictures(
						newStrPtr("https://test.com/profile-pic"),
						newStrPtr("https://test.com/cover-pic"),
					),
			),
			expEvent: sdk.NewEvent(
				types.EventTypeProfileSaved,
				sdk.NewAttribute(types.AttributeProfileDtag, "test_dtag"),
				sdk.NewAttribute(types.AttributeProfileCreator, suite.testData.profile.Creator.String()),
				sdk.NewAttribute(types.AttributeProfileCreationTime, suite.testData.profile.CreationDate.Format(time.RFC3339)),
			),
		},
		{
			name: "Profile saving fails due to wrong tag",
			existentProfiles: types.NewProfiles(
				suite.testData.profile,
				types.NewProfile("editor_dtag", suite.testData.profile.Creator, suite.testData.profile.CreationDate).
					WithBio(newStrPtr("biography")),
			),
			msg: types.NewMsgSaveProfile(
				"editor_tag",
				newStrPtr("new-moniker"),
				newStrPtr("new-bio"),
				nil,
				nil,
				suite.testData.profile.Creator, // Use the same user
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "wrong dtag provided. Make sure to use the current one"),
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

func (suite *KeeperTestSuite) Test_handleMsgCreateMonoDirectionalRelationship() {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)
	receiver, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)
	rel := types.NewMonodirectionalRelationship(sender, receiver)

	tests := []struct {
		name               string
		msg                types.MsgCreateMonoDirectionalRelationship
		storedRelationship *types.MonodirectionalRelationship
		expErr             error
		expEvent           sdk.Event
	}{
		{
			name:               "Relationship already created returns error",
			msg:                types.NewMsgCreateMonoDirectionalRelationship(sender, receiver),
			storedRelationship: &rel,
			expErr:             sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("relationship with %s has already been made", receiver)),
		},
		{
			name:               "Relationship has been saved correctly",
			msg:                types.NewMsgCreateMonoDirectionalRelationship(sender, receiver),
			storedRelationship: nil,
			expErr:             nil,
			expEvent: sdk.NewEvent(
				types.EventTypeMonodirectionalRelationshipCreated,
				sdk.NewAttribute(types.AttributeRelationshipSender, rel.Sender.String()),
				sdk.NewAttribute(types.AttributeRelationshipReceiver, rel.Receiver.String()),
			),
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			if test.storedRelationship != nil {
				suite.keeper.StoreRelationship(suite.ctx, *test.storedRelationship)
			}

			handler := keeper.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			if test.expErr != nil {
				suite.Error(err)
				suite.Equal(test.expErr.Error(), err.Error())
			}

			if test.expErr == nil {
				suite.NoError(err)

				suite.True(suite.keeper.DoesRelationshipExist(suite.ctx, rel.ID))

				// Check the events
				suite.Len(res.Events, 1)
				suite.Contains(res.Events, test.expEvent)
			}

		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgRequestBiDirectionalRelationship() {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)
	receiver, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)
	rel := types.NewBiDirectionalRelationship(sender, receiver, types.Sent)

	tests := []struct {
		name               string
		msg                types.MsgRequestBidirectionalRelationship
		storedRelationship *types.BidirectionalRelationship
		expErr             error
		expEvent           sdk.Event
	}{
		{
			name:               "Relationship already created returns error",
			msg:                types.NewMsgRequestBidirectionalRelationship(sender, receiver, "hello"),
			storedRelationship: &rel,
			expErr:             sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("relationship request to %s has already been made", receiver)),
		},
		{
			name:               "Relationship has been saved correctly",
			msg:                types.NewMsgRequestBidirectionalRelationship(sender, receiver, "hello"),
			storedRelationship: nil,
			expErr:             nil,
			expEvent: sdk.NewEvent(
				types.EventTypeBidirectionalRelationshipRequested,
				sdk.NewAttribute(types.AttributeRelationshipMessage, "hello"),
				sdk.NewAttribute(types.AttributeRelationshipSender, rel.Sender.String()),
				sdk.NewAttribute(types.AttributeRelationshipReceiver, rel.Receiver.String()),
				sdk.NewAttribute(types.AttributeRelationshipStatus, rel.Status.String()),
			),
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			if test.storedRelationship != nil {
				suite.keeper.StoreRelationship(suite.ctx, *test.storedRelationship)
			}

			handler := keeper.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			if test.expErr != nil {
				suite.Error(err)
				suite.Equal(test.expErr.Error(), err.Error())
			}

			if test.expErr == nil {
				suite.NoError(err)

				suite.True(suite.keeper.DoesRelationshipExist(suite.ctx, rel.ID))

				// Check the events
				suite.Len(res.Events, 1)
				suite.Contains(res.Events, test.expEvent)
			}

		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgAcceptBiDirectionalRelationship() {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)
	receiver, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)
	var abstractRelBiAccepted types.Relationship
	abstractRelBiAccepted = types.NewBiDirectionalRelationship(sender, receiver, types.Accepted)

	var abstractRelBiSent types.Relationship
	abstractRelBiSent = types.NewBiDirectionalRelationship(sender, receiver, types.Sent)

	var abstractRelMono types.Relationship
	abstractRelMono = types.NewMonodirectionalRelationship(sender, receiver)

	tests := []struct {
		name               string
		msg                types.MsgAcceptBidirectionalRelationship
		storedRelationship *types.Relationship
		expErr             error
		expEvent           sdk.Event
	}{
		{
			name:               "Relationship doesn't exist and returns error",
			msg:                types.NewMsgAcceptBidirectionalRelationship(abstractRelBiAccepted.RelationshipID(), receiver),
			storedRelationship: nil,
			expErr:             sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("relationship with id %s doesn't exist", abstractRelBiAccepted.RelationshipID())),
		},
		{
			name:               "Relationship already accepted returns error",
			msg:                types.NewMsgAcceptBidirectionalRelationship(abstractRelBiAccepted.RelationshipID(), receiver),
			storedRelationship: &abstractRelBiAccepted,
			expErr:             sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("the relationship with id: %s has already been accepted", abstractRelBiAccepted.RelationshipID())),
		},
		{
			name:               "Relationship with wrong receiver returns error",
			msg:                types.NewMsgAcceptBidirectionalRelationship(abstractRelBiSent.RelationshipID(), sender),
			storedRelationship: &abstractRelBiSent,
			expErr:             sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("%s isn't the recipient of the relationship with ID: %s", sender, abstractRelBiSent.RelationshipID())),
		},
		{
			name:               "Relationship with wrong type returns error",
			msg:                types.NewMsgAcceptBidirectionalRelationship(abstractRelMono.RelationshipID(), receiver),
			storedRelationship: &abstractRelMono,
			expErr:             sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("the relationship with id: %s is not a bidirectional relationship and cannot be accepted", abstractRelMono.RelationshipID())),
		},
		{
			name:               "Relationship has been saved correctly",
			msg:                types.NewMsgAcceptBidirectionalRelationship(abstractRelBiSent.RelationshipID(), receiver),
			storedRelationship: &abstractRelBiSent,
			expErr:             nil,
			expEvent: sdk.NewEvent(
				types.EventTypeBidirectionalRelationshipAccepted,
				sdk.NewAttribute(types.AttributeRelationshipID, abstractRelBiSent.RelationshipID().String()),
				sdk.NewAttribute(types.AttributeRelationshipReceiver, abstractRelBiSent.Recipient().String()),
				sdk.NewAttribute(types.AttributeRelationshipStatus, types.Accepted.String()),
			),
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			if test.storedRelationship != nil {
				rel := *test.storedRelationship
				if _, ok := rel.(types.BidirectionalRelationship); ok {
					suite.keeper.SaveUserRelationshipAssociation(suite.ctx, []sdk.AccAddress{rel.Creator()}, rel.RelationshipID())
				}
				suite.keeper.SaveUserRelationshipAssociation(suite.ctx, []sdk.AccAddress{rel.Recipient()}, rel.RelationshipID())
				suite.keeper.StoreRelationship(suite.ctx, *test.storedRelationship)
			}

			handler := keeper.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			if test.expErr != nil {
				suite.Error(err)
				suite.Equal(test.expErr.Error(), err.Error())
			}

			if test.expErr == nil {
				suite.NoError(err)

				suite.True(suite.keeper.DoesRelationshipExist(suite.ctx, abstractRelBiSent.RelationshipID()))

				// Check the events
				suite.Len(res.Events, 1)
				suite.Contains(res.Events, test.expEvent)
			}

		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgDenyBidirectionalRelationship() {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)
	receiver, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)
	var abstractRelBiAccepted types.Relationship
	abstractRelBiAccepted = types.NewBiDirectionalRelationship(sender, receiver, types.Accepted)

	var abstractRelBiSent types.Relationship
	abstractRelBiSent = types.NewBiDirectionalRelationship(sender, receiver, types.Sent)

	var abstractRelMono types.Relationship
	abstractRelMono = types.NewMonodirectionalRelationship(sender, receiver)

	tests := []struct {
		name               string
		msg                types.MsgDenyBidirectionalRelationship
		storedRelationship *types.Relationship
		expErr             error
		expEvent           sdk.Event
	}{
		{
			name:               "Relationship doesn't exist and returns error",
			msg:                types.NewMsgDenyBidirectionalRelationship(abstractRelBiAccepted.RelationshipID(), receiver),
			storedRelationship: nil,
			expErr:             sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("relationship with id %s doesn't exist", abstractRelBiAccepted.RelationshipID())),
		},
		{
			name:               "Relationship already accepted returns error",
			msg:                types.NewMsgDenyBidirectionalRelationship(abstractRelBiAccepted.RelationshipID(), receiver),
			storedRelationship: &abstractRelBiAccepted,
			expErr:             sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("the relationship with id: %s has already been accepted", abstractRelBiAccepted.RelationshipID())),
		},
		{
			name:               "Relationship with wrong receiver returns error",
			msg:                types.NewMsgDenyBidirectionalRelationship(abstractRelBiSent.RelationshipID(), sender),
			storedRelationship: &abstractRelBiSent,
			expErr:             sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("%s isn't the recipient of the relationship with ID: %s", sender, abstractRelBiSent.RelationshipID())),
		},
		{
			name:               "Relationship with wrong type returns error",
			msg:                types.NewMsgDenyBidirectionalRelationship(abstractRelMono.RelationshipID(), receiver),
			storedRelationship: &abstractRelMono,
			expErr:             sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("the relationship with id: %s is not a bidirectional relationship and cannot be denied", abstractRelMono.RelationshipID())),
		},
		{
			name:               "Relationship has been saved correctly",
			msg:                types.NewMsgDenyBidirectionalRelationship(abstractRelBiSent.RelationshipID(), receiver),
			storedRelationship: &abstractRelBiSent,
			expErr:             nil,
			expEvent: sdk.NewEvent(
				types.EventTypeBidirectionalRelationshipDenied,
				sdk.NewAttribute(types.AttributeRelationshipID, abstractRelBiSent.RelationshipID().String()),
				sdk.NewAttribute(types.AttributeRelationshipReceiver, abstractRelBiSent.Recipient().String()),
				sdk.NewAttribute(types.AttributeRelationshipStatus, types.Denied.String()),
			),
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			if test.storedRelationship != nil {
				rel := *test.storedRelationship
				if _, ok := rel.(types.BidirectionalRelationship); ok {
					suite.keeper.SaveUserRelationshipAssociation(suite.ctx, []sdk.AccAddress{rel.Creator()}, rel.RelationshipID())
				}
				suite.keeper.SaveUserRelationshipAssociation(suite.ctx, []sdk.AccAddress{rel.Recipient()}, rel.RelationshipID())
				suite.keeper.StoreRelationship(suite.ctx, *test.storedRelationship)
			}

			handler := keeper.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			if test.expErr != nil {
				suite.Error(err)
				suite.Equal(test.expErr.Error(), err.Error())
			}

			if test.expErr == nil {
				suite.NoError(err)

				suite.True(suite.keeper.DoesRelationshipExist(suite.ctx, abstractRelBiSent.RelationshipID()))

				// Check the events
				suite.Len(res.Events, 1)
				suite.Contains(res.Events, test.expEvent)
			}

		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgDeleteRelationship() {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)
	receiver, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)
	var abstractRelBiAccepted types.Relationship
	abstractRelBiAccepted = types.NewBiDirectionalRelationship(sender, receiver, types.Accepted)

	var abstractRelMono types.Relationship
	abstractRelMono = types.NewMonodirectionalRelationship(sender, receiver)

	tests := []struct {
		name               string
		msg                types.MsgDeleteRelationship
		storedRelationship *types.Relationship
		expErr             error
		expEvent           sdk.Event
	}{
		{
			name:               "Relationship doesn't exist and returns error",
			msg:                types.NewMsgDeleteRelationship(abstractRelBiAccepted.RelationshipID(), receiver),
			storedRelationship: nil,
			expErr:             sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("relationship with id %s doesn't exist", abstractRelBiAccepted.RelationshipID())),
		},
		{
			name:               "Unauthorized user returns error",
			msg:                types.NewMsgDeleteRelationship(abstractRelMono.RelationshipID(), receiver),
			storedRelationship: &abstractRelMono,
			expErr:             sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("user with address %s isn't the relationship's creator", receiver)),
		},
		{
			name:               "User delete relationship successfully",
			msg:                types.NewMsgDeleteRelationship(abstractRelMono.RelationshipID(), sender),
			storedRelationship: &abstractRelMono,
			expErr:             nil,
			expEvent: sdk.NewEvent(
				types.EventTypeRelationshipsDeleted,
				sdk.NewAttribute(types.AttributeRelationshipID, abstractRelMono.RelationshipID().String()),
				sdk.NewAttribute(types.AttributeRelationshipSender, abstractRelMono.Creator().String()),
			),
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			if test.storedRelationship != nil {
				rel := *test.storedRelationship
				suite.keeper.SaveUserRelationshipAssociation(suite.ctx, []sdk.AccAddress{rel.Creator()}, rel.RelationshipID())
				if _, ok := rel.(types.BidirectionalRelationship); ok {
					suite.keeper.SaveUserRelationshipAssociation(suite.ctx, []sdk.AccAddress{rel.Recipient()}, rel.RelationshipID())
				}
				suite.keeper.StoreRelationship(suite.ctx, *test.storedRelationship)
			}

			handler := keeper.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			if test.expErr != nil {
				suite.Error(err)
				suite.Equal(test.expErr.Error(), err.Error())
			}

			if test.expErr == nil {
				suite.NoError(err)

				suite.False(suite.keeper.DoesRelationshipExist(suite.ctx, abstractRelMono.RelationshipID()))

				// Check the events
				suite.Len(res.Events, 1)
				suite.Contains(res.Events, test.expEvent)
			}

		})
	}

}
