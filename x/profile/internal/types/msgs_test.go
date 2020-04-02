package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"
)

// ----------------------
// --- MsgCreateProfile
// ----------------------

var testProfileOwner, _ = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
var testPictures = types.NewPictures("profile", "cover")

var testProfile = types.Profile{
	Name:     "name",
	Surname:  "surname",
	Moniker:  "moniker",
	Bio:      "biography",
	Pictures: &testPictures,
	Creator:  testProfileOwner,
}

var msgCreateProfile = types.NewMsgCreateProfile(
	testProfile.Name,
	testProfile.Surname,
	testProfile.Moniker,
	testProfile.Bio,
	testProfile.Pictures,
	testProfile.Creator,
)

var msgEditProfile = types.NewMsgEditProfile(
	testProfile.Moniker,
	"monk",
	testProfile.Name,
	testProfile.Surname,
	testProfile.Bio,
	testProfile.Pictures,
	testProfile.Creator,
)

var msgDeleteProfile = types.NewMsgDeleteProfile(
	testProfile.Moniker,
	testProfile.Creator,
)

func TestMsgCreateProfile_Route(t *testing.T) {
	actual := msgCreateProfile.Route()
	require.Equal(t, "profile", actual)
}

func TestMsgCreateProfile_Type(t *testing.T) {
	actual := msgCreateProfile.Type()
	require.Equal(t, "create_profile", actual)
}

func TestMsgCreateProfile_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   types.MsgCreateProfile
		error error
	}{
		{
			name: "Empty owner returns error",
			msg: types.NewMsgCreateProfile(
				testProfile.Name,
				testProfile.Surname,
				testProfile.Moniker,
				testProfile.Bio,
				testProfile.Pictures,
				nil,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid creator address: "),
		},
		{
			name: "Max name length exceeded",
			msg: types.NewMsgCreateProfile(
				"9YfrVVi3UEI1ymN7n6isScyHNSt30xG6Jn1EDxEXxWOn0voSMIKqLhHsBfnZoXEXeFlAO5qMwjNGvgoiNBtoMfR78J2SNhBz"+
					"wNxlTky9DCJ2F2luh9cTc7umcHl2BDwSepE1Iijn4htrP7vcKWgIgHYh73oNmF7PTiU1gmL2G8W4XB06bpDLFb0eLzPbSGLe51"+
					"25k9tljhFBdgSPtoKuLQUQPGC3IqyyTIqQEpLeNpmbiJUDmbqQ1tyyS8mDC7WQEYv8uuYU90pjBSkGJQs2FI2Q7hIHL202O1SF"+
					"sTkJ5H9v30Jry3HqmjxYv1yG1PWah2Gkg7xP0toSdEXObDE9YWo6LMDO29yyTrohCwG9RHo04l8jfJOUbuer7BrXmWodFuGhIcd"+
					"C43T4R4l5a5P6zWlUkWuhYZCtX1dpfENb4wlDNHd2r1TFCblNs7COKSUINVd8swxR2lEzRO2mwE39mvUEBEHi0S06QtU1m8Chv"+
					"6ou0LSnJMCTq",
				testProfile.Surname,
				testProfile.Moniker,
				testProfile.Bio,
				testProfile.Pictures,
				testProfile.Creator,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Profile name cannot exceed 500 characters"),
		},
		{
			name: "Max surname length exceeded",
			msg: types.NewMsgCreateProfile(
				testProfile.Name,
				"9YfrVVi3UEI1ymN7n6isScyHNSt30xG6Jn1EDxEXxWOn0voSMIKqLhHsBfnZoXEXeFlAO5qMwjNGvgoiNBtoMfR78J2SNhBz"+
					"wNxlTky9DCJ2F2luh9cTc7umcHl2BDwSepE1Iijn4htrP7vcKWgIgHYh73oNmF7PTiU1gmL2G8W4XB06bpDLFb0eLzPbSGLe51"+
					"25k9tljhFBdgSPtoKuLQUQPGC3IqyyTIqQEpLeNpmbiJUDmbqQ1tyyS8mDC7WQEYv8uuYU90pjBSkGJQs2FI2Q7hIHL202O1SF"+
					"sTkJ5H9v30Jry3HqmjxYv1yG1PWah2Gkg7xP0toSdEXObDE9YWo6LMDO29yyTrohCwG9RHo04l8jfJOUbuer7BrXmWodFuGhIcd"+
					"C43T4R4l5a5P6zWlUkWuhYZCtX1dpfENb4wlDNHd2r1TFCblNs7COKSUINVd8swxR2lEzRO2mwE39mvUEBEHi0S06QtU1m8Chv"+
					"6ou0LSnJMCTq",
				testProfile.Moniker,
				testProfile.Bio,
				testProfile.Pictures,
				testProfile.Creator,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Profile surname cannot exceed 500 characters"),
		},
		{
			name: "Max bio length exceeded",
			msg: types.NewMsgCreateProfile(
				testProfile.Name,
				testProfile.Surname,
				testProfile.Moniker,
				"9YfrVVi3UEI1ymN7n6isScyHNSt30xG6Jn1EDxEXxWOn0voSMIKqLhHsBfnZoXEXeFlAO5qMwjNGvgoiNBtoMfR78J2SNhBz"+
					"wNxlTky9DCJ2F2luh9cTc7umcHl2BDwSepE1Iijn4htrP7vcKWgIgHYh73oNmF7PTiU1gmL2G8W4XB06bpDLFb0eLzPbSGLe51"+
					"25k9tljhFBdgSPtoKuLQUQPGC3IqyyTIqQEpLeNpmbiJUDmbqQ1tyyS8mDC7WQEYv8uuYU90pjBSkGJQs2FI2Q7hIHL202O1SF"+
					"sTkJ5H9v30Jry3HqmjxYv1yG1PWah2Gkg7xP0toSdEXObDE9YWo6LMDO29yyTrohCwG9RHo04l8jfJOUbuer7BrXmWodFuGhIcd"+
					"C43T4R4l5a5P6zWlUkWuhYZCtX1dpfENb4wlDNHd2r1TFCblNs7COKSUINVd8swxR2lEzRO2mwE39mvUEBEHi0S06QtU1m8Chv"+
					"6ou0LSnJMCTq9YfrVVi3UEI1ymN7n6isScyHNSt30xG6Jn1EDxEXxWOn0voSMIKqLhHsBfnZoXEXeFlAO5qMwjNGvgoiNBtoMfR78J2SNhBz"+
					"wNxlTky9DCJ2F2luh9cTc7umcHl2BDwSepE1Iijn4htrP7vcKWgIgHYh73oNmF7PTiU1gmL2G8W4XB06bpDLFb0eLzPbSGLe51"+
					"25k9tljhFBdgSPtoKuLQUQPGC3IqyyTIqQEpLeNpmbiJUDmbqQ1tyyS8mDC7WQEYv8uuYU90pjBSkGJQs2FI2Q7hIHL202O1SF"+
					"sTkJ5H9v30Jry3HqmjxYv1yG1PWah2Gkg7xP0toSdEXObDE9YWo6LMDO29yyTrohCwG9RHo04l8jfJOUbuer7BrXmWodFuGhIcd"+
					"C43T4R4l5a5P6zWlUkWuhYZCtX1dpfENb4wlDNHd2r1TFCblNs7COKSUINVd8swxR2lEzRO2mwE39mvUEBEHi0S06QtU1m8Chv"+
					"6ou0LSnJMCTq",
				testProfile.Pictures,
				testProfile.Creator,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Profile biography cannot exceed 1000 characters"),
		},
		{
			name: "Empty moniker error",
			msg: types.NewMsgCreateProfile(
				testProfile.Name,
				testProfile.Surname,
				"",
				testProfile.Bio,
				testProfile.Pictures,
				testProfile.Creator,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Profile moniker cannot be blank or empty"),
		},
		{
			name: "Max moniker length exceeded",
			msg: types.NewMsgCreateProfile(
				testProfile.Name,
				testProfile.Surname,
				"asdserhrtyjeqrgdfhnr1asdserhrtyjeqrgdfhnr1",
				testProfile.Bio,
				testProfile.Pictures,
				testProfile.Creator,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Profile moniker cannot exceed 30 characters"),
		},
		{
			name: "No error message",
			msg: types.NewMsgCreateProfile(
				testProfile.Name,
				testProfile.Surname,
				testProfile.Moniker,
				testProfile.Bio,
				testProfile.Pictures,
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

func TestMsgCreateProfile_GetSignBytes(t *testing.T) {
	actual := msgCreateProfile.GetSignBytes()
	expected := `{"type":"desmos/MsgCreateProfile","value":{"bio":"biography","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","moniker":"moniker","name":"name","pictures":{"cover":"cover","profile":"profile"},"surname":"surname"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgCreateProfile_GetSigners(t *testing.T) {
	actual := msgCreateProfile.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgCreateProfile.Creator, actual[0])
}

// ----------------------
// --- MsgEditProfile
// ----------------------

func TestMsgEditProfile_Route(t *testing.T) {
	actual := msgEditProfile.Route()
	require.Equal(t, "profile", actual)
}

func TestMsgEditProfile_Type(t *testing.T) {
	actual := msgEditProfile.Type()
	require.Equal(t, "edit_profile", actual)
}

func TestMsgEditProfile_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   types.MsgEditProfile
		error error
	}{
		{
			name: "Empty owner returns error",
			msg: types.NewMsgEditProfile(
				testProfile.Moniker,
				testProfile.Moniker,
				testProfile.Name,
				testProfile.Surname,
				testProfile.Bio,
				testProfile.Pictures,
				nil,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid creator address: "),
		},
		{
			name: "Max name length exceeded",
			msg: types.NewMsgEditProfile(
				testProfile.Moniker,
				testProfile.Moniker,
				"9YfrVVi3UEI1ymN7n6isScyHNSt30xG6Jn1EDxEXxWOn0voSMIKqLhHsBfnZoXEXeFlAO5qMwjNGvgoiNBtoMfR78J2SNhBz"+
					"wNxlTky9DCJ2F2luh9cTc7umcHl2BDwSepE1Iijn4htrP7vcKWgIgHYh73oNmF7PTiU1gmL2G8W4XB06bpDLFb0eLzPbSGLe51"+
					"25k9tljhFBdgSPtoKuLQUQPGC3IqyyTIqQEpLeNpmbiJUDmbqQ1tyyS8mDC7WQEYv8uuYU90pjBSkGJQs2FI2Q7hIHL202O1SF"+
					"sTkJ5H9v30Jry3HqmjxYv1yG1PWah2Gkg7xP0toSdEXObDE9YWo6LMDO29yyTrohCwG9RHo04l8jfJOUbuer7BrXmWodFuGhIcd"+
					"C43T4R4l5a5P6zWlUkWuhYZCtX1dpfENb4wlDNHd2r1TFCblNs7COKSUINVd8swxR2lEzRO2mwE39mvUEBEHi0S06QtU1m8Chv"+
					"6ou0LSnJMCTq",
				testProfile.Surname,
				testProfile.Bio,
				testProfile.Pictures,
				testProfile.Creator,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Profile name cannot exceed 500 characters"),
		},
		{
			name: "Max surname length exceeded",
			msg: types.NewMsgEditProfile(
				testProfile.Moniker,
				testProfile.Moniker,
				testProfile.Name,
				"9YfrVVi3UEI1ymN7n6isScyHNSt30xG6Jn1EDxEXxWOn0voSMIKqLhHsBfnZoXEXeFlAO5qMwjNGvgoiNBtoMfR78J2SNhBz"+
					"wNxlTky9DCJ2F2luh9cTc7umcHl2BDwSepE1Iijn4htrP7vcKWgIgHYh73oNmF7PTiU1gmL2G8W4XB06bpDLFb0eLzPbSGLe51"+
					"25k9tljhFBdgSPtoKuLQUQPGC3IqyyTIqQEpLeNpmbiJUDmbqQ1tyyS8mDC7WQEYv8uuYU90pjBSkGJQs2FI2Q7hIHL202O1SF"+
					"sTkJ5H9v30Jry3HqmjxYv1yG1PWah2Gkg7xP0toSdEXObDE9YWo6LMDO29yyTrohCwG9RHo04l8jfJOUbuer7BrXmWodFuGhIcd"+
					"C43T4R4l5a5P6zWlUkWuhYZCtX1dpfENb4wlDNHd2r1TFCblNs7COKSUINVd8swxR2lEzRO2mwE39mvUEBEHi0S06QtU1m8Chv"+
					"6ou0LSnJMCTq",
				testProfile.Bio,
				testProfile.Pictures,
				testProfile.Creator,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Profile surname cannot exceed 500 characters"),
		},
		{
			name: "Max bio length exceeded",
			msg: types.NewMsgEditProfile(
				testProfile.Moniker,
				testProfile.Moniker,
				testProfile.Name,
				testProfile.Surname,
				"9YfrVVi3UEI1ymN7n6isScyHNSt30xG6Jn1EDxEXxWOn0voSMIKqLhHsBfnZoXEXeFlAO5qMwjNGvgoiNBtoMfR78J2SNhBz"+
					"wNxlTky9DCJ2F2luh9cTc7umcHl2BDwSepE1Iijn4htrP7vcKWgIgHYh73oNmF7PTiU1gmL2G8W4XB06bpDLFb0eLzPbSGLe51"+
					"25k9tljhFBdgSPtoKuLQUQPGC3IqyyTIqQEpLeNpmbiJUDmbqQ1tyyS8mDC7WQEYv8uuYU90pjBSkGJQs2FI2Q7hIHL202O1SF"+
					"sTkJ5H9v30Jry3HqmjxYv1yG1PWah2Gkg7xP0toSdEXObDE9YWo6LMDO29yyTrohCwG9RHo04l8jfJOUbuer7BrXmWodFuGhIcd"+
					"C43T4R4l5a5P6zWlUkWuhYZCtX1dpfENb4wlDNHd2r1TFCblNs7COKSUINVd8swxR2lEzRO2mwE39mvUEBEHi0S06QtU1m8Chv"+
					"6ou0LSnJMCTq9YfrVVi3UEI1ymN7n6isScyHNSt30xG6Jn1EDxEXxWOn0voSMIKqLhHsBfnZoXEXeFlAO5qMwjNGvgoiNBtoMfR78J2SNhBz"+
					"wNxlTky9DCJ2F2luh9cTc7umcHl2BDwSepE1Iijn4htrP7vcKWgIgHYh73oNmF7PTiU1gmL2G8W4XB06bpDLFb0eLzPbSGLe51"+
					"25k9tljhFBdgSPtoKuLQUQPGC3IqyyTIqQEpLeNpmbiJUDmbqQ1tyyS8mDC7WQEYv8uuYU90pjBSkGJQs2FI2Q7hIHL202O1SF"+
					"sTkJ5H9v30Jry3HqmjxYv1yG1PWah2Gkg7xP0toSdEXObDE9YWo6LMDO29yyTrohCwG9RHo04l8jfJOUbuer7BrXmWodFuGhIcd"+
					"C43T4R4l5a5P6zWlUkWuhYZCtX1dpfENb4wlDNHd2r1TFCblNs7COKSUINVd8swxR2lEzRO2mwE39mvUEBEHi0S06QtU1m8Chv"+
					"6ou0LSnJMCTq",
				testProfile.Pictures,
				testProfile.Creator,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Profile biography cannot exceed 1000 characters"),
		},
		{
			name: "Empty moniker error",
			msg: types.NewMsgEditProfile(
				"",
				testProfile.Moniker,
				testProfile.Name,
				testProfile.Surname,
				testProfile.Bio,
				testProfile.Pictures,
				testProfile.Creator,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Profile previous moniker cannot be blank or empty"),
		},
		{
			name: "Max previous moniker length exceeded",
			msg: types.NewMsgEditProfile(
				"asdserhrtyjeqrgdfhnr1asdserhrtyjeqrgdfhnr1",
				testProfile.Moniker,
				testProfile.Name,
				testProfile.Surname,
				testProfile.Bio,
				testProfile.Pictures,
				testProfile.Creator,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Profile previous moniker cannot exceed 30 characters"),
		},
		{
			name: "Max new moniker length exceeded",
			msg: types.NewMsgEditProfile(
				testProfile.Moniker,
				"asdserhrtyjeqrgdfhnr1asdserhrtyjeqrgdfhnr1",
				testProfile.Name,
				testProfile.Surname,
				testProfile.Bio,
				testProfile.Pictures,
				testProfile.Creator,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Profile new moniker cannot exceed 30 characters"),
		},
		{
			name: "No error message",
			msg: types.NewMsgEditProfile(
				testProfile.Moniker,
				testProfile.Moniker,
				testProfile.Name,
				testProfile.Surname,
				testProfile.Bio,
				testProfile.Pictures,
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

func TestMsgEditProfile_GetSignBytes(t *testing.T) {
	actual := msgEditProfile.GetSignBytes()
	expected := `{"type":"desmos/MsgEditProfile","value":{"bio":"biography","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","name":"name","new_moniker":"monk","pictures":{"cover":"cover","profile":"profile"},"previous_moniker":"moniker","surname":"surname"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgEditProfile_GetSigners(t *testing.T) {
	actual := msgEditProfile.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgEditProfile.Creator, actual[0])
}

// ----------------------
// --- MsgDeleteProfile
// ----------------------

func TestMsgDeleteProfile_Route(t *testing.T) {
	actual := msgDeleteProfile.Route()
	require.Equal(t, "profile", actual)
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
				testProfile.Moniker,
				nil,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid creator address: "),
		},
		{
			name: "Empty moniker returns error",
			msg: types.NewMsgDeleteProfile(
				"",
				testProfile.Creator,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "NewMoniker cannot be blank or empty"),
		},
		{
			name: "Valid message returns no error",
			msg: types.NewMsgDeleteProfile(
				testProfile.Moniker,
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
	expected := `{"type":"desmos/MsgDeleteProfile","value":{"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","moniker":"moniker"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgDeleteProfile_GetSigners(t *testing.T) {
	actual := msgDeleteProfile.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgDeleteProfile.Creator, actual[0])
}
