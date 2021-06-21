package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

var msgLinkApplication = types.NewMsgLinkApplication(
	types.NewData("twitter", "twitterusername"),
	types.NewOracleRequestCallData("twitter", ""),
	"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
	types.IBCPortID,
	"channel-0",
	clienttypes.NewHeight(0, 1000),
	0,
)

func TestMsgLinkApplication_Route(t *testing.T) {
	require.Equal(t, "profiles", msgLinkApplication.Route())
}

func TestMsgLinkApplication_Type(t *testing.T) {
	require.Equal(t, "link_application", msgLinkApplication.Type())
}

func TestMsgLinkApplication_ValidateBasic(t *testing.T) {
	tests := []struct {
		name      string
		msg       *types.MsgLinkApplication
		shouldErr bool
	}{
		{
			name: "invalid link data returns error",
			msg: types.NewMsgLinkApplication(
				types.NewData("", "twitteruser"),
				types.NewOracleRequestCallData(
					"twitter",
					"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
				),
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				"channel-0",
				types.IBCPortID,
				clienttypes.NewHeight(0, 1000),
				0,
			),
			shouldErr: true,
		},
		{
			name: "invalid oracle request call data returns error",
			msg: types.NewMsgLinkApplication(
				types.NewData("twitter", "twitteruser"),
				types.NewOracleRequestCallData("twitter", "calldata"),
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				"channel-0",
				types.IBCPortID,
				clienttypes.NewHeight(0, 1000),
				0,
			),
			shouldErr: true,
		},
		{
			name: "invalid channel returns error",
			msg: types.NewMsgLinkApplication(
				types.NewData("twitter", "twitteruser"),
				types.NewOracleRequestCallData(
					"twitter",
					"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
				),
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				"",
				types.IBCPortID,
				clienttypes.NewHeight(0, 1000),
				0,
			),
			shouldErr: true,
		},
		{
			name: "invalid port returns error",
			msg: types.NewMsgLinkApplication(
				types.NewData("twitter", "twitteruser"),
				types.NewOracleRequestCallData(
					"twitter",
					"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
				),
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				"",
				"channel-0",
				clienttypes.NewHeight(0, 1000),
				0,
			),
			shouldErr: true,
		},
		{
			name: "invalid signer returns error",
			msg: types.NewMsgLinkApplication(
				types.NewData("twitter", "twitteruser"),
				types.NewOracleRequestCallData(
					"twitter",
					"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
				),
				"cosmos10nsdy9qka3zv0lzw8z9cnu6kanld8jh773",
				"channel-0",
				types.IBCPortID,
				clienttypes.NewHeight(0, 1000),
				0,
			),
			shouldErr: true,
		},
		{
			name: "Valid message returns no error",
			msg: types.NewMsgLinkApplication(
				types.NewData("twitter", "twitteruser"),
				types.NewOracleRequestCallData(
					"twitter",
					"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
				),
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				types.IBCPortID,
				"channel-0",
				clienttypes.NewHeight(0, 1000),
				0,
			),
			shouldErr: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.msg.ValidateBasic()
			if test.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgLinkApplication_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgLinkApplication","value":{"call_data":{"application":"twitter"},"link_data":{"application":"twitter","username":"twitterusername"},"sender":"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773","source_channel":"channel-0","source_port":"ibc-profiles","timeout_height":{"revision_height":"1000"}}}`
	require.Equal(t, expected, string(msgLinkApplication.GetSignBytes()))
}

func TestMsgLinkApplication_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgLinkApplication.Sender)
	require.Equal(t, []sdk.AccAddress{addr}, msgLinkApplication.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgUnlinkApplication = types.NewMsgUnlinkApplication(
	"twitter",
	"twitteruser",
	"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
)

func TestMsgUnlinkApplication_Route(t *testing.T) {
	require.Equal(t, "profiles", msgLinkApplication.Route())
}

func TestMsgUnlinkApplication_Type(t *testing.T) {
	require.Equal(t, "unlink_application", msgUnlinkApplication.Type())
}

func TestMsgUnlinkApplication_ValidateBasic(t *testing.T) {
	tests := []struct {
		name      string
		msg       *types.MsgUnlinkApplication
		shouldErr bool
	}{
		{
			name: "invalid application returns error",
			msg: types.NewMsgUnlinkApplication(
				"",
				"twitteruser",
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			),
			shouldErr: true,
		},
		{
			name: "invalid username",
			msg: types.NewMsgUnlinkApplication(
				"twitter",
				"",
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			),
			shouldErr: true,
		},
		{
			name: "invalid signer returns error",
			msg: types.NewMsgUnlinkApplication(
				"twitter",
				"twitteruser",
				"cosmos10nvy9qka3zv0lzw8z9cnu6kanld8jh773",
			),
			shouldErr: true,
		},
		{
			name: "Valid message returns no error",
			msg: types.NewMsgUnlinkApplication(
				"twitter",
				"twitteruser",
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			),
			shouldErr: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.msg.ValidateBasic()
			if test.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgUnlinkApplication_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgUnlinkApplication","value":{"application":"twitter","signer":"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773","username":"twitteruser"}}`
	require.Equal(t, expected, string(msgUnlinkApplication.GetSignBytes()))
}

func TestMsgUnlinkApplication_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgUnlinkApplication.Signer)
	require.Equal(t, []sdk.AccAddress{addr}, msgUnlinkApplication.GetSigners())
}
