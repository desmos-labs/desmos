package keeper_test

import (
	"fmt"
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
	"github.com/desmos-labs/desmos/x/profile/internal/types/msgs"
	"testing"

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
		profile models.Profile
		expErr  error
	}{
		{
			name: "Max name length exceeded",
			profile: models.Profile{
				Moniker:  testProfile.Moniker,
				Name:     &invalidMaxLenField,
				Surname:  testProfile.Surname,
				Bio:      testProfile.Bio,
				Pictures: models.NewPictures(testProfile.Pictures.Profile, testProfile.Pictures.Cover),
				Creator:  testProfile.Creator,
			},
			expErr: fmt.Errorf("Profile name cannot exceed 1000 characters"),
		},
		{
			name: "Min name length not reached",
			profile: models.Profile{
				Moniker:  testProfile.Moniker,
				Name:     &invalidMinLenField,
				Surname:  testProfile.Surname,
				Bio:      testProfile.Bio,
				Pictures: models.NewPictures(testProfile.Pictures.Profile, testProfile.Pictures.Cover),
				Creator:  testProfile.Creator,
			},
			expErr: fmt.Errorf("Profile name cannot be less than 2 characters"),
		},
		{
			name: "Max surname length exceeded",
			profile: models.Profile{
				Moniker:  testProfile.Moniker,
				Name:     testProfile.Name,
				Surname:  &invalidMaxLenField,
				Bio:      testProfile.Bio,
				Pictures: models.NewPictures(testProfile.Pictures.Profile, testProfile.Pictures.Cover),
				Creator:  testProfile.Creator,
			},
			expErr: fmt.Errorf("Profile surname cannot exceed 1000 characters"),
		},
		{
			name: "Min surname length not reached",
			profile: models.Profile{
				Moniker:  testProfile.Moniker,
				Name:     testProfile.Name,
				Surname:  &invalidMinLenField,
				Bio:      testProfile.Bio,
				Pictures: models.NewPictures(testProfile.Pictures.Profile, testProfile.Pictures.Cover),
				Creator:  testProfile.Creator,
			},
			expErr: fmt.Errorf("Profile surname cannot be less than 2 characters"),
		},
		{
			name: "Max bio length exceeded",
			profile: models.Profile{
				Moniker:  testProfile.Moniker,
				Name:     testProfile.Name,
				Surname:  testProfile.Surname,
				Bio:      &invalidBio,
				Pictures: models.NewPictures(testProfile.Pictures.Profile, testProfile.Pictures.Cover),
				Creator:  testProfile.Creator,
			},
			expErr: fmt.Errorf("Profile biography cannot exceed 1000 characters"),
		},
		{
			name: "Min moniker length not reached",
			profile: models.Profile{
				Moniker:  "l",
				Name:     testProfile.Name,
				Surname:  testProfile.Surname,
				Bio:      testProfile.Bio,
				Pictures: models.NewPictures(testProfile.Pictures.Profile, testProfile.Pictures.Cover),
				Creator:  testProfile.Creator,
			},
			expErr: fmt.Errorf("Profile moniker cannot be less than 2 characters"),
		},
		{
			name: "Max moniker length exceeded",
			profile: models.Profile{
				Moniker:  invalidMonikerLen,
				Name:     testProfile.Name,
				Surname:  testProfile.Surname,
				Bio:      testProfile.Bio,
				Pictures: models.NewPictures(testProfile.Pictures.Profile, testProfile.Pictures.Cover),
				Creator:  testProfile.Creator,
			},
			expErr: fmt.Errorf("Profile moniker cannot exceed 30 characters"),
		},
		{
			name: "Invalid profile pictures returns error",
			profile: models.Profile{
				Name:     &name,
				Surname:  &surname,
				Moniker:  "moniker",
				Bio:      &bio,
				Pictures: models.NewPictures(&invalidPic, testProfile.Pictures.Cover),
				Creator:  testPostOwner,
			},
			expErr: fmt.Errorf("invalid profile picture uri provided"),
		},
		{
			name: "Valid profile returns no error",
			profile: models.Profile{
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
			k.SetParams(ctx, models.DefaultParams())
			actual := keeper.ValidateProfile(ctx, k, test.profile)
			require.Equal(t, test.expErr, actual)
		})
	}
}

func Test_handleMsgSaveProfile(t *testing.T) {
	editor, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	var name = "name"
	var surname = "surname"
	var bio = "biography"
	var newMoniker = "newMoniker"
	var invalidPic = "pic"

	testAcc2 := models.Profile{
		Name:     &name,
		Surname:  &surname,
		Moniker:  "newMoniker",
		Bio:      &bio,
		Pictures: testPictures,
		Creator:  editor,
	}

	tests := []struct {
		name             string
		existentProfiles models.Profiles
		expAccount       *models.Profile
		msg              msgs.MsgSaveProfile
		expErr           error
	}{
		{
			name:             "Profile saved (with previous profile created)",
			existentProfiles: models.Profiles{testProfile},
			msg: msgs.NewMsgSaveProfile(
				newMoniker,
				testProfile.Name,
				testProfile.Surname,
				testProfile.Bio,
				testProfile.Pictures.Profile,
				testProfile.Pictures.Cover,
				testProfile.Creator,
			),
			expErr: nil,
		},
		{
			name:             "Profile saved (with no previous profile created)",
			existentProfiles: nil,
			msg: msgs.NewMsgSaveProfile(
				newMoniker,
				testProfile.Name,
				testProfile.Surname,
				testProfile.Bio,
				testProfile.Pictures.Profile,
				testProfile.Pictures.Cover,
				testProfile.Creator,
			),
			expErr: nil,
		},
		{
			name:             "Profile not edited because the new moniker already exists",
			existentProfiles: models.Profiles{testProfile, testAcc2},
			msg: msgs.NewMsgSaveProfile(
				newMoniker,
				testProfile.Name,
				testProfile.Surname,
				testProfile.Bio,
				testProfile.Pictures.Profile,
				testProfile.Pictures.Cover,
				testPostOwner,
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "a profile with moniker: newMoniker has already been created"),
		},
		{
			name:             "Profile not edited because of the invalid pics uri",
			existentProfiles: models.Profiles{testProfile},
			msg: msgs.NewMsgSaveProfile(
				newMoniker,
				testProfile.Name,
				testProfile.Surname,
				testProfile.Bio,
				&invalidPic,
				testProfile.Pictures.Cover,
				testProfile.Creator,
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid profile picture uri provided"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)
			k.SetParams(ctx, models.DefaultParams())
			if test.existentProfiles != nil {
				for _, acc := range test.existentProfiles {
					key := models.ProfileStoreKey(acc.Creator)
					store.Set(key, k.Cdc.MustMarshalBinaryBare(acc))
					k.AssociateMonikerWithAddress(ctx, acc.Moniker, acc.Creator)
				}
			}

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			if res == nil {
				require.NotNil(t, err)
				require.Equal(t, test.expErr.Error(), err.Error())
			}
			if res != nil {

				profiles := k.GetProfiles(ctx)
				require.Len(t, profiles, 1)

				//Check the data
				require.Equal(t, k.Cdc.MustMarshalBinaryLengthPrefixed(test.msg.Moniker), res.Data)

				//Check the events
				createAccountEv := sdk.NewEvent(
					types.EventTypeProfileSaved,
					sdk.NewAttribute(types.AttributeProfileMoniker, test.msg.Moniker),
					sdk.NewAttribute(types.AttributeProfileCreator, test.msg.Creator.String()),
				)

				require.Len(t, res.Events, 1)
				require.Contains(t, res.Events, createAccountEv)
			}

		})
	}
}

func Test_handleMsgDeleteProfile(t *testing.T) {
	tests := []struct {
		name            string
		existentAccount *models.Profile
		msg             msgs.MsgDeleteProfile
		expErr          error
	}{
		{
			name:            "Profile doesnt exists",
			existentAccount: nil,
			msg:             msgs.NewMsgDeleteProfile(testProfile.Creator),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"No profile associated with this address: cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
		},
		{
			name:            "Profile deleted successfully",
			existentAccount: &testProfile,
			msg:             msgs.NewMsgDeleteProfile(testProfile.Creator),
			expErr:          nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)

			if test.existentAccount != nil {
				key := models.ProfileStoreKey(test.existentAccount.Creator)
				store.Set(key, k.Cdc.MustMarshalBinaryBare(&test.existentAccount))
				k.AssociateMonikerWithAddress(ctx, test.existentAccount.Moniker, test.existentAccount.Creator)
			}

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			if res == nil {
				require.NotNil(t, err)
				require.Equal(t, test.expErr.Error(), err.Error())
			}
			if res != nil {
				//Check the data
				require.Equal(t, k.Cdc.MustMarshalBinaryLengthPrefixed("moniker"), res.Data)

				//Check the events
				createAccountEv := sdk.NewEvent(
					types.EventTypeProfileDeleted,
					sdk.NewAttribute(types.AttributeProfileMoniker, "moniker"),
					sdk.NewAttribute(types.AttributeProfileCreator, test.msg.Creator.String()),
				)

				require.Len(t, res.Events, 1)
				require.Contains(t, res.Events, createAccountEv)
			}

		})
	}
}
