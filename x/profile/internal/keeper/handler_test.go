package keeper_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/desmos-labs/desmos/x/profile/internal/keeper"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func Test_validateProfile(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1tg8csfcg8m8u7vu5vph9fayhfcw5hyc47mey2e")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2010, 10, 02, 12, 10, 00, 00, timeZone)

	tests := []struct {
		name    string
		profile types.Profile
		expErr  error
	}{
		{
			name: "Max moniker length exceeded",
			profile: types.NewProfile("custom_dtag", user, date).
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
			profile: types.NewProfile("custom_dtag", user, date).
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
			profile: types.NewProfile("custom_dtag", user, date).
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
			profile: types.NewProfile("custom.", user, date).
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
			profile: types.NewProfile("d", user, date).
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
			profile: types.NewProfile("9YfrVVi3UEI1ymN7n6isScyHNSt30xG6Jn1EDxEXxWOn0voSMIKqLhHsBfnZoXE", user, date).
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
			profile: types.NewProfile("dtag", user, date).
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
			profile: types.NewProfile("dtag", user, date).
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
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			k.SetParams(ctx, types.DefaultParams())
			actual := keeper.ValidateProfile(ctx, k, test.profile)
			require.Equal(t, test.expErr, actual)
		})
	}
}

func Test_handleMsgSaveProfile(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1tg8csfcg8m8u7vu5vph9fayhfcw5hyc47mey2e")
	require.NoError(t, err)

	editor, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2010, 10, 02, 12, 10, 00, 00, timeZone)

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
				user,
			),
			expProfiles: types.NewProfiles(
				types.NewProfile("custom_dtag", user, date).
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
				sdk.NewAttribute(types.AttributeProfileCreator, user.String()),
			),
		},
		{
			name: "Profile saved (with previous profile created)",
			existentProfiles: types.NewProfiles(
				types.NewProfile("test_dtag", user, date).
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
				user,
			),
			expProfiles: types.NewProfiles(
				types.NewProfile("test_dtag", user, date).
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
				sdk.NewAttribute(types.AttributeProfileCreator, user.String()),
			),
		},
		{
			name: "Profile saving fails due to wrong tag",
			existentProfiles: types.NewProfiles(
				testProfile,
				types.NewProfile("editor_dtag", editor, date).
					WithBio(newStrPtr("biography")),
			),
			msg: types.NewMsgSaveProfile(
				"editor_tag",
				newStrPtr("new-moniker"),
				newStrPtr("new-bio"),
				nil,
				nil,
				editor, // Use the same user
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "wrong dtag provided. Make sure to use the current one"),
		},
		{
			name:             "Profile not edited because of the invalid profile picture",
			existentProfiles: types.Profiles{testProfile},
			msg: types.NewMsgSaveProfile(
				"custom_dtag",
				nil,
				nil,
				newStrPtr("invalid-pic"),
				nil,
				user,
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid profile picture uri provided"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)
			k.SetParams(ctx, types.DefaultParams())
			if test.existentProfiles != nil {
				for _, acc := range test.existentProfiles {
					key := types.ProfileStoreKey(acc.Creator)
					store.Set(key, k.Cdc.MustMarshalBinaryBare(acc))
					k.AssociateDtagWithAddress(ctx, acc.DTag, acc.Creator)
				}
			}

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			if test.expErr != nil {
				require.Error(t, err)
				require.Equal(t, test.expErr.Error(), err.Error())
			}

			if test.expErr == nil {
				require.NoError(t, err)

				profiles := k.GetProfiles(ctx)
				require.Len(t, profiles, len(test.expProfiles))
				for index, profile := range profiles {
					require.True(t, profile.Equals(test.expProfiles[index]))
				}

				// Check the events
				require.Len(t, res.Events, 1)
				require.Contains(t, res.Events, test.expEvent)
			}

		})
	}
}

func Test_handleMsgDeleteProfile(t *testing.T) {
	tests := []struct {
		name            string
		existentAccount *types.Profile
		msg             types.MsgDeleteProfile
		expErr          error
	}{
		{
			name:            "Profile doesn't exists",
			existentAccount: nil,
			msg:             types.NewMsgDeleteProfile(testProfile.Creator),
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				"No profile associated with this address: cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
		},
		{
			name:            "Profile deleted successfully",
			existentAccount: &testProfile,
			msg:             types.NewMsgDeleteProfile(testProfile.Creator),
			expErr:          nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)

			if test.existentAccount != nil {
				key := types.ProfileStoreKey(test.existentAccount.Creator)
				store.Set(key, k.Cdc.MustMarshalBinaryBare(&test.existentAccount))
				k.AssociateDtagWithAddress(ctx, test.existentAccount.DTag, test.existentAccount.Creator)
			}

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			if res == nil {
				require.NotNil(t, err)
				require.Equal(t, test.expErr.Error(), err.Error())
			}
			if res != nil {
				// Check the data
				require.Equal(t, k.Cdc.MustMarshalBinaryLengthPrefixed("dtag"), res.Data)

				// Check the events
				createAccountEv := sdk.NewEvent(
					types.EventTypeProfileDeleted,
					sdk.NewAttribute(types.AttributeProfileDtag, "dtag"),
					sdk.NewAttribute(types.AttributeProfileCreator, test.msg.Creator.String()),
				)

				require.Len(t, res.Events, 1)
				require.Contains(t, res.Events, createAccountEv)
			}

		})
	}
}
