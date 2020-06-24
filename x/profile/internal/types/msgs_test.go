package types_test

import (
	"fmt"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

// ----------------------
// --- MsgSaveProfile
// ----------------------

var user, _ = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
var testProfile = types.Profile{
	DTag:    "dtag",
	Moniker: newStrPtr("moniker"),
	Bio:     newStrPtr("biography"),
	Pictures: types.NewPictures(
		newStrPtr("https://shorturl.at/adnX3"),
		newStrPtr("https://shorturl.at/cgpyF"),
	),
	Creator: user,
}

var msgEditProfile = types.NewMsgSaveProfile(
	"monk",
	testProfile.Moniker,
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
	tests := []struct {
		name  string
		msg   types.MsgSaveProfile
		error error
	}{
		{
			name: "Empty owner returns error",
			msg: types.NewMsgSaveProfile(
				testProfile.DTag,
				testProfile.Moniker,
				testProfile.Bio,
				testProfile.Pictures.Profile,
				testProfile.Pictures.Cover,
				nil,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid creator address: "),
		},
		{
			name: "Max bio length exceeded",
			msg: types.NewMsgSaveProfile(
				testProfile.DTag,
				nil,
				newStrPtr(strings.Repeat("a", 2000)),
				nil,
				nil,
				testProfile.Creator,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Profile biography cannot exceed 1000 characters"),
		},
		{
			name:  "Min dtag length not reached",
			msg:   types.NewMsgSaveProfile("l", nil, nil, nil, nil, testProfile.Creator),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid profile dtag provided: 'l'"),
		},
		{
			name: "Max dtag length exceeded",
			msg: types.NewMsgSaveProfile(
				strings.Repeat("_", 100),
				nil,
				nil,
				nil,
				nil,
				testProfile.Creator,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("Invalid profile dtag provided: '%s'", strings.Repeat("_", 100))),
		},
		{
			name:  "Invalid dtag characters",
			msg:   types.NewMsgSaveProfile("d.tag", nil, nil, nil, nil, testProfile.Creator),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid profile dtag provided: 'd.tag'"),
		},
		{
			name: "No error message",
			msg: types.NewMsgSaveProfile(
				"_crazy_papa_21",
				newStrPtr("custom-moniker"),
				newStrPtr("custom-bio"),
				newStrPtr("https://test.com/my-custom-profile-pic"),
				newStrPtr("https://test.com/my-custom-cover-pic"),
				user,
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
	expected := `{"type":"desmos/MsgSaveProfile","value":{"bio":"biography","cover_picture":"https://shorturl.at/cgpyF","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","dtag":"monk","moniker":"moniker","profile_picture":"https://shorturl.at/adnX3"}}`
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
