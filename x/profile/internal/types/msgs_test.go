package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"
)

// ----------------------
// --- MsgSaveProfile
// ----------------------

var testProfileOwner, _ = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
var testProfilePic = "https://shorturl.at/adnX3"
var testCoverPic = "https://shorturl.at/cgpyF"
var testPictures = types.NewPictures(&testProfilePic, &testCoverPic)
var name = "name"
var surname = "surname"
var bio = "biography"

var invalidMaxLenField = "9YfrVVi3UEI1ymN7n6isScyHNSt30xG6Jn1EDxEXxWOn0voSMIKqLhHsBfnZoXEXeFlAO5qMwjNGvgoiNBtoMfR78J2SNhBz" +
	"wNxlTky9DCJ2F2luh9cTc7umcHl2BDwSepE1Iijn4htrP7vcKWgIgHYh73oNmF7PTiU1gmL2G8W4XB06bpDLFb0eLzPbSGLe51" +
	"25k9tljhFBdgSPtoKuLQUQPGC3IqyyTIqQEpLeNpmbiJUDmbqQ1tyyS8mDC7WQEYv8uuYU90pjBSkGJQs2FI2Q7hIHL202O1SF" +
	"sTkJ5H9v30Jry3HqmjxYv1yG1PWah2Gkg7xP0toSdEXObDE9YWo6LMDO29yyTrohCwG9RHo04l8jfJOUbuer7BrXmWodFuGhIcd" +
	"C43T4R4l5a5P6zWlUkWuhYZCtX1dpfENb4wlDNHd2r1TFCblNs7COKSUINVd8swxR2lEzRO2mwE39mvUEBEHi0S06QtU1m8Chv" +
	"6ou0LSnJMCTq"
var invalidBio = "9YfrVVi3UEI1ymN7n6isScyHNSt30xG6Jn1EDxEXxWOn0voSMIKqLhHsBfnZoXEXeFlAO5qMwjNGvgoiNBtoMfR78J2SNhBz" +
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
var invalidMinLenField = "l"

var testProfile = types.Profile{
	Name:     &name,
	Surname:  &surname,
	Moniker:  "moniker",
	Bio:      &bio,
	Pictures: testPictures,
	Creator:  testProfileOwner,
}

var newMoniker = "monk"
var msgEditProfile = types.NewMsgSaveProfile(
	newMoniker,
	testProfile.Name,
	testProfile.Surname,
	testProfile.Bio,
	testProfile.Pictures.Profile,
	testProfile.Pictures.Cover,
	testProfile.Creator,
)

var msgDeleteProfile = types.NewMsgDeleteProfile(
	testProfile.Creator,
)

func TestMsgSaveProfile_Route(t *testing.T) {
	actual := msgEditProfile.Route()
	require.Equal(t, "profiles", actual)
}

func TestMsgSaveProfile_Type(t *testing.T) {
	actual := msgEditProfile.Type()
	require.Equal(t, "save_profile", actual)
}

func TestMsgSaveProfile_ValidateBasic(t *testing.T) {
	invalidMonikerLen := "asdserhrtyjeqrgdfhnr1asdserhrtyjeqrgdfhnr1"
	tests := []struct {
		name  string
		msg   types.MsgSaveProfile
		error error
	}{
		{
			name: "Empty owner returns error",
			msg: types.NewMsgSaveProfile(
				testProfile.Moniker,
				testProfile.Name,
				testProfile.Surname,
				testProfile.Bio,
				testProfile.Pictures.Profile,
				testProfile.Pictures.Cover,
				nil,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid creator address: "),
		},
		{
			name: "Max name length exceeded",
			msg: types.NewMsgSaveProfile(
				testProfile.Moniker,
				&invalidMaxLenField,
				testProfile.Surname,
				testProfile.Bio,
				testProfile.Pictures.Profile,
				testProfile.Pictures.Cover,
				testProfile.Creator,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Profile name cannot exceed 500 characters"),
		},
		{
			name: "Min name length not reached",
			msg: types.NewMsgSaveProfile(
				testProfile.Moniker,
				&invalidMinLenField,
				testProfile.Surname,
				testProfile.Bio,
				testProfile.Pictures.Profile,
				testProfile.Pictures.Cover,
				testProfile.Creator,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Profile name cannot be less than 2 characters"),
		},
		{
			name: "Max surname length exceeded",
			msg: types.NewMsgSaveProfile(
				testProfile.Moniker,
				testProfile.Name,
				&invalidMaxLenField,
				testProfile.Bio,
				testProfile.Pictures.Profile,
				testProfile.Pictures.Cover,
				testProfile.Creator,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Profile surname cannot exceed 500 characters"),
		},
		{
			name: "Min surname length not reached",
			msg: types.NewMsgSaveProfile(
				testProfile.Moniker,
				testProfile.Name,
				&invalidMinLenField,
				testProfile.Bio,
				testProfile.Pictures.Profile,
				testProfile.Pictures.Cover,
				testProfile.Creator,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Profile surname cannot be less than 2 characters"),
		},
		{
			name: "Max bio length exceeded",
			msg: types.NewMsgSaveProfile(
				testProfile.Moniker,
				testProfile.Name,
				testProfile.Surname,
				&invalidBio,
				testProfile.Pictures.Profile,
				testProfile.Pictures.Cover,
				testProfile.Creator,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Profile biography cannot exceed 1000 characters"),
		},
		{
			name: "Min moniker length not reached",
			msg: types.NewMsgSaveProfile(
				"l",
				testProfile.Name,
				testProfile.Surname,
				testProfile.Bio,
				testProfile.Pictures.Profile,
				testProfile.Pictures.Cover,
				testProfile.Creator,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Profile moniker cannot be less than 2 characters"),
		},
		{
			name: "Max moniker length exceeded",
			msg: types.NewMsgSaveProfile(
				invalidMonikerLen,
				testProfile.Name,
				testProfile.Surname,
				testProfile.Bio,
				testProfile.Pictures.Profile,
				testProfile.Pictures.Cover,
				testProfile.Creator,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Profile moniker cannot exceed 30 characters"),
		},
		{
			name: "No error message",
			msg: types.NewMsgSaveProfile(
				testProfile.Moniker,
				testProfile.Name,
				testProfile.Surname,
				testProfile.Bio,
				testProfile.Pictures.Profile,
				testProfile.Pictures.Cover,
				testProfile.Creator,
			),
			error: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			returnedError := test.msg.ValidateBasic()
			if test.error == nil {
				require.Nil(t, returnedError)
			} else {
				require.NotNil(t, returnedError)
				require.Equal(t, test.error.Error(), returnedError.Error())
			}
		})
	}
}

func TestMsgSaveProfile_GetSignBytes(t *testing.T) {
	actual := msgEditProfile.GetSignBytes()
	expected := `{"type":"desmos/MsgSaveProfile","value":{"bio":"biography","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","moniker":"monk","name":"name","profile_cov":"https://shorturl.at/cgpyF","profile_pic":"https://shorturl.at/adnX3","surname":"surname"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgSaveProfile_GetSigners(t *testing.T) {
	actual := msgEditProfile.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgEditProfile.Creator, actual[0])
}

// ----------------------
// --- MsgDeleteProfile
// ----------------------

func TestMsgDeleteProfile_Route(t *testing.T) {
	actual := msgDeleteProfile.Route()
	require.Equal(t, "profiles", actual)
}

func TestMsgDeleteProfile_Type(t *testing.T) {
	actual := msgDeleteProfile.Type()
	require.Equal(t, "delete_profile", actual)
}

func TestMsgDeleteProfile_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   types.MsgDeleteProfile
		error error
	}{
		{
			name: "Empty owner returns error",
			msg: types.NewMsgDeleteProfile(
				nil,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid creator address: "),
		},
		{
			name: "Valid message returns no error",
			msg: types.NewMsgDeleteProfile(
				testProfile.Creator,
			),
			error: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			returnedError := test.msg.ValidateBasic()
			if test.error == nil {
				require.Nil(t, returnedError)
			} else {
				require.NotNil(t, returnedError)
				require.Equal(t, test.error.Error(), returnedError.Error())
			}
		})
	}
}

func TestMsgDeleteProfile_GetSignBytes(t *testing.T) {
	actual := msgDeleteProfile.GetSignBytes()
	expected := `{"type":"desmos/MsgDeleteProfile","value":{"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgDeleteProfile_GetSigners(t *testing.T) {
	actual := msgDeleteProfile.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgDeleteProfile.Creator, actual[0])
}
