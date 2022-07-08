package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/profiles/types"

	"github.com/stretchr/testify/require"
)

var msgEditProfile = types.NewMsgSaveProfile(
	"monk",
	"nickname",
	"biography",
	"https://shorturl.at/adnX3",
	"https://shorturl.at/cgpyF",
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
)

func TestMsgSaveProfile_Route(t *testing.T) {
	require.Equal(t, "profiles", msgEditProfile.Route())
}

func TestMsgSaveProfile_Type(t *testing.T) {
	require.Equal(t, "save_profile", msgEditProfile.Type())
}

func TestMsgSaveProfile_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgSaveProfile
		shouldErr bool
	}{
		{
			name: "empty owner returns error",
			msg: types.NewMsgSaveProfile(
				"monk",
				"nickname",
				"biography",
				"https://shorturl.at/adnX3",
				"https://shorturl.at/cgpyF",
				"",
			),
			shouldErr: true,
		},
		{
			name:      "valid message returns no error",
			msg:       msgEditProfile,
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgSaveProfile_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgSaveProfile","value":{"bio":"biography","cover_picture":"https://shorturl.at/cgpyF","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","dtag":"monk","nickname":"nickname","profile_picture":"https://shorturl.at/adnX3"}}`
	require.Equal(t, expected, string(msgEditProfile.GetSignBytes()))
}

func TestMsgSaveProfile_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgEditProfile.Creator)
	require.Equal(t, []sdk.AccAddress{addr}, msgEditProfile.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgDeleteProfile = types.NewMsgDeleteProfile(
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
)

func TestMsgDeleteProfile_Route(t *testing.T) {
	require.Equal(t, "profiles", msgDeleteProfile.Route())
}

func TestMsgDeleteProfile_Type(t *testing.T) {
	require.Equal(t, "delete_profile", msgDeleteProfile.Type())
}

func TestMsgDeleteProfile_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgDeleteProfile
		shouldErr bool
	}{
		{
			name:      "empty owner returns error",
			msg:       types.NewMsgDeleteProfile(""),
			shouldErr: true,
		},
		{
			name:      "valid message returns no error",
			msg:       msgDeleteProfile,
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgDeleteProfile_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgDeleteProfile","value":{"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(msgDeleteProfile.GetSignBytes()))
}

func TestMsgDeleteProfile_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgDeleteProfile.Creator)
	require.Equal(t, []sdk.AccAddress{addr}, msgDeleteProfile.GetSigners())
}
