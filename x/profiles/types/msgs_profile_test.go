package types_test

import (
	"testing"
	"time"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

var addr, _ = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
var testProfile, _ = types.NewProfile(
	"dtag",
	"nickname",
	"biography",
	types.NewPictures(
		"https://shorturl.at/adnX3",
		"https://shorturl.at/cgpyF",
	),
	time.Unix(100, 0),
	authtypes.NewBaseAccountWithAddress(addr),
)

// ___________________________________________________________________________________________________________________

var msgEditProfile = types.NewMsgSaveProfile(
	"monk",
	testProfile.Nickname,
	testProfile.Bio,
	testProfile.Pictures.Profile,
	testProfile.Pictures.Cover,
	testProfile.GetAddress().String(),
)

func TestMsgSaveProfile_Route(t *testing.T) {
	require.Equal(t, "profiles", msgEditProfile.Route())
}

func TestMsgSaveProfile_Type(t *testing.T) {
	require.Equal(t, "save_profile", msgEditProfile.Type())
}

func TestMsgSaveProfile_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types.MsgSaveProfile
		error error
	}{
		{
			name: "Empty owner returns error",
			msg: types.NewMsgSaveProfile(
				testProfile.DTag,
				testProfile.Nickname,
				testProfile.Bio,
				testProfile.Pictures.Profile,
				testProfile.Pictures.Cover,
				"",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator: "),
		},
		{
			name:  "Invalid empty dtag returns error",
			msg:   types.NewMsgSaveProfile("", "", "", "", "", testProfile.GetAddress().String()),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "profile dtag cannot be empty or blank"),
		},
		{
			name: "No error message",
			msg: types.NewMsgSaveProfile(
				"_crazy_papa_21",
				"custom-nickname",
				"custom-bio",
				"https://test.com/my-custom-profile-pic",
				"https://test.com/my-custom-cover-pic",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
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
	expected := `{"type":"desmos/MsgSaveProfile","value":{"bio":"biography","cover_picture":"https://shorturl.at/cgpyF","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","dtag":"monk","nickname":"nickname","profile_picture":"https://shorturl.at/adnX3"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgSaveProfile_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgEditProfile.Creator)
	require.Equal(t, []sdk.AccAddress{addr}, msgEditProfile.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgDeleteProfile = types.NewMsgDeleteProfile(
	testProfile.GetAddress().String(),
)

func TestMsgDeleteProfile_Route(t *testing.T) {
	require.Equal(t, "profiles", msgDeleteProfile.Route())
}

func TestMsgDeleteProfile_Type(t *testing.T) {
	require.Equal(t, "delete_profile", msgDeleteProfile.Type())
}

func TestMsgDeleteProfile_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types.MsgDeleteProfile
		error error
	}{
		{
			name:  "Empty owner returns error",
			msg:   types.NewMsgDeleteProfile(""),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator: "),
		},
		{
			name:  "Valid message returns no error",
			msg:   types.NewMsgDeleteProfile(testProfile.GetAddress().String()),
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
	addr, _ := sdk.AccAddressFromBech32(msgDeleteProfile.Creator)
	require.Equal(t, []sdk.AccAddress{addr}, msgDeleteProfile.GetSigners())
}
