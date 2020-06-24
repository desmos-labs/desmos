package keeper_test

import (
	"testing"
	"time"

	"github.com/desmos-labs/desmos/x/profile/internal/keeper"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func Test_validateProfile(t *testing.T) {
	invalidPic := "pic"
	invalidMonikerLen := "asdserhrtyjeqrgdfhnr1asdserhrtyjeqrgdfhnr1"
	invalidMaxLenField := "9YfrVVi3UEI1ymN7n6isScyHNSt30xG6Jn1EDxEXxWOn0voSMIKqLhHsBfnZoXEXeFlAO5qMwjNGvgoiNBtoMfR78J2SNhBz" +
		"wNxlTky9DCJ2F2luh9cTc7umcHl2BDwSepE1Iijn4htrP7vcKWgIgHYh73oNmF7PTiU1gmL2G8W4XB06bpDLFb0eLzPbSGLe51" +
		"25k9tljhFBdgSPtoKuLQUQPGC3IqyyTIqQEpLeNpmbiJUDmbqQ1tyyS8mDC7WQEYv8uuYU90pjBSkGJQs2FI2Q7hIHL202O1SF" +
		"sTkJ5H9v30Jry3HqmjxYv1yG1PWah2Gkg7xP0toSdEXObDE9YWo6LMDO29yyTrohCwG9RHo04l8jfJOUbuer7BrXmWodFuGhIcd" +
		"C43T4R4l5a5P6zWlUkWuhYZCtX1dpfENb4wlDNHd2r1TFCblNs7COKSUINVd8swxR2lEzRO2mwE39mvUEBEHi0S06QtU1m8Chv" +
		"6ou0LSnJMCTq9YfrVVi3UEI1ymN7n6isScyHNSt30xG6Jn1EDxEXxWOn0voSMIKqLhHsBfnZoXEXeFlAO5qMwjNGvgoiNBtoMfR78J2SNhBz" +
		"wNxlTky9DCJ2F2luh9cTc7umcHl2BDwSepE1Iijn4htrP7vcKWgIgHYh73oNmF7PTiU1gmL2G8W4XB06bpDLFb0eLzPbSGLe51" +
		"25k9tljhFBdgSPtoKuLQUQPGC3IqyyTIqQEpLeNpmbiJUDmbqQ1tyyS8mDC7WQEYv8uuYU90pjBSkGJQs2FI2Q7hIHL202O1SF" +
		"sTkJ5H9v30Jry3HqmjxYv1yG1PWah2Gkg7xP0toSdEXObDE9YWo6LMDO29yyTrohCwG9RHo04l8jfJOUbuer7BrXmWodFuGhIcd" +
		"C43T4R4l5a5P6zWlUkWuhYZCtX1dpfENb4wlDNHd2r1TFCblNs7COKSUINVd8swxR2lEzRO2mwE39mvUEBEHi0S06QtU1m8Chv" +
		"6ou0LSnJMCTq"
	invalidBio := "9YfrVVi3UEI1ymN7n6isScyHNSt30xG6Jn1EDxEXxWOn0voSMIKqLhHsBfnZoXEXeFlAO5qMwjNGvgoiNBtoMfR78J2SNhBz" +
		"wNxlTky9DCJ2F2luh9cTc7umcHl2BDwSepE1Iijn4htrP7vcKWgIgHYh73oNmF7PTiU1gmL2G8W4XB06bpDLFb0eLzPbSGLe51" +
		"25k9tljhFBdgSPtoKuLQUQPGC3IqyyTIqQEpLeNpmbiJUDmbqQ1tyyS8mDC7WQEYv8uuYU90pjBSkGJQs2FI2Q7hIHL202O1SF" +
		"sTkJ5H9v30Jry3HqmjxYv1yG1PWah2Gkg7xP0toSdEXObDE9YWo6LMDO29yyTrohCwG9RHo04l8jfJOUbuer7BrXmWodFuGhIcd" +
		"C43T4R4l5a5P6zWlUkWuhYZCtX1dpfENb4wlDNHd2r1TFCblNs7COKSUINVd8swxR2lEzRO2mwE39mvUEBEHi0S06QtU1m8Chv" +
		"6ou0LSnJMCTq9YfrVVi3UEI1ymN7n6isScyHNSt30xG6Jn1EDxEXxWOn0voSMIKqLhHsBfnZoXEXeFlAO5qMwjNGvgoiNBtoMfR78J2SNhBz" +
		"wNxlTky9DCJ2F2luh9cTc7umcHl2BDwSepE1Iijn4htrP7vcKWgIgHYh73oNmF7PTiU1gmL2G8W4XB06bpDLFb0eLzPbSGLe51" +
		"25k9tljhFBdgSPtoKuLQUQPGC3IqyyTIqQEpLeNpmbiJUDmbqQ1tyyS8mDC7WQEYv8uuYU90pjBSkGJQs2FI2Q7hIHL202O1SF" +
		"sTkJ5H9v30Jry3HqmjxYv1yG1PWah2Gkg7xP0toSdEXObDE9YWo6LMDO29yyTrohCwG9RHo04l8jfJOUbuer7BrXmWodFuGhIcd" +
		"C43T4R4l5a5P6zWlUkWuhYZCtX1dpfENb4wlDNHd2r1TFCblNs7COKSUINVd8swxR2lEzRO2mwE39mvUEBEHi0S06QtU1m8Chv" +
		"6ou0LSnJMCTq"
	invalidMinLenField := "l"

	tests := []struct {
		name    string
		profile types.Profile
		expErr  error
	}{
		{
			name: "Max name length exceeded",
			profile: types.Profile{
				Moniker:  testProfile.Moniker,
				Name:     &invalidMaxLenField,
				Surname:  testProfile.Surname,
				Bio:      testProfile.Bio,
				Pictures: types.NewPictures(testProfile.Pictures.Profile, testProfile.Pictures.Cover),
				Creator:  testProfile.Creator,
			},
			expErr: fmt.Errorf("Profile name cannot exceed 1000 characters"),
		},
		{
			name: "Min name length not reached",
			profile: types.Profile{
				Moniker:  testProfile.Moniker,
				Name:     &invalidMinLenField,
				Surname:  testProfile.Surname,
				Bio:      testProfile.Bio,
				Pictures: types.NewPictures(testProfile.Pictures.Profile, testProfile.Pictures.Cover),
				Creator:  testProfile.Creator,
			},
			expErr: fmt.Errorf("Profile name cannot be less than 2 characters"),
		},
		{
			name: "Max surname length exceeded",
			profile: types.Profile{
				Moniker:  testProfile.Moniker,
				Name:     testProfile.Name,
				Surname:  &invalidMaxLenField,
				Bio:      testProfile.Bio,
				Pictures: types.NewPictures(testProfile.Pictures.Profile, testProfile.Pictures.Cover),
				Creator:  testProfile.Creator,
			},
			expErr: fmt.Errorf("Profile surname cannot exceed 1000 characters"),
		},
		{
			name: "Min surname length not reached",
			profile: types.Profile{
				Moniker:  testProfile.Moniker,
				Name:     testProfile.Name,
				Surname:  &invalidMinLenField,
				Bio:      testProfile.Bio,
				Pictures: types.NewPictures(testProfile.Pictures.Profile, testProfile.Pictures.Cover),
				Creator:  testProfile.Creator,
			},
			expErr: fmt.Errorf("Profile surname cannot be less than 2 characters"),
		},
		{
			name: "Max bio length exceeded",
			profile: types.Profile{
				Moniker:  testProfile.Moniker,
				Name:     testProfile.Name,
				Surname:  testProfile.Surname,
				Bio:      &invalidBio,
				Pictures: types.NewPictures(testProfile.Pictures.Profile, testProfile.Pictures.Cover),
				Creator:  testProfile.Creator,
			},
			expErr: fmt.Errorf("Profile biography cannot exceed 1000 characters"),
		},
		{
			name: "Min moniker length not reached",
			profile: types.Profile{
				Moniker:  "l",
				Name:     testProfile.Name,
				Surname:  testProfile.Surname,
				Bio:      testProfile.Bio,
				Pictures: types.NewPictures(testProfile.Pictures.Profile, testProfile.Pictures.Cover),
				Creator:  testProfile.Creator,
			},
			expErr: fmt.Errorf("Profile moniker cannot be less than 2 characters"),
		},
		{
			name: "Max moniker length exceeded",
			profile: types.Profile{
				Moniker:  invalidMonikerLen,
				Name:     testProfile.Name,
				Surname:  testProfile.Surname,
				Bio:      testProfile.Bio,
				Pictures: types.NewPictures(testProfile.Pictures.Profile, testProfile.Pictures.Cover),
				Creator:  testProfile.Creator,
			},
			expErr: fmt.Errorf("Profile moniker cannot exceed 30 characters"),
		},
		{
			name: "Invalid profile pictures returns error",
			profile: types.Profile{
				Name:     &name,
				Surname:  &surname,
				Moniker:  "moniker",
				Bio:      &bio,
				Pictures: types.NewPictures(&invalidPic, testProfile.Pictures.Cover),
				Creator:  testPostOwner,
			},
			expErr: fmt.Errorf("invalid profile picture uri provided"),
		},
		{
			name: "Valid profile returns no error",
			profile: types.Profile{
				Name:     &name,
				Surname:  &surname,
				Moniker:  "moniker",
				Bio:      &bio,
				Pictures: testPictures,
				Creator:  testPostOwner,
			},
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
			existentAccounts: types.NewProfiles(
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
			existentAccounts: types.NewProfiles(
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
			existentAccounts: types.Profiles{testProfile},
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
