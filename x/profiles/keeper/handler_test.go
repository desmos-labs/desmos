package keeper_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/desmos-labs/desmos/x/profiles/keeper"
	"github.com/desmos-labs/desmos/x/profiles/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/profiles/internal/keeper"
	"github.com/desmos-labs/desmos/x/profiles/internal/types"
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
				"No profile associated with this address: cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
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
