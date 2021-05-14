package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/ibc/profiles/types"
	"github.com/stretchr/testify/require"
)

var validConnectionMsg = types.NewMsgCreateIBCAccountConnection(
	"ibc-profiles",
	"desmos-0",
	types.NewIBCAccountConnectionPacketData(
		"cosmos",
		"cosmos1u33w3u4ler4654phrpt6xqvh92ch0v6mcjrj97",
		"02bcd0738e3b7e0f6650c8e6eb10bd4266fd6818c92a0283b4cb0884f046051c3e",
		"cosmos1qdasq0mzpajknaaj32kf9lk5nmcy8g65mddd4p",
		"143bbd3131d76232f973f84a9ea7be751243044315056dbecb968942da97e474401caa1c8f2c4ce5e48052cf44066717f2166c21a7277de9911d75c57eca598d",
		"1a18d5f012ce0e8258fd3455c01b48249bb019231e416c4323ab2bb170b4ad0951b370138d2ea69a376feb942d3c619c9152d63a6d2e0232aaff77162df66636",
	),
	1000,
)

var validLinkMsg = types.NewMsgCreateIBCAccountLink(
	"ibc-profiles",
	"desmos-0",
	types.NewIBCAccountLinkPacketData(
		"cosmos",
		"cosmos1fc6dg2f85hsd7nc7c4ng89uad7hzh6thpwa6xf",
		"039d2b5c2e5be3f0f3e1665f622734aa21bbfcb1bc0099fbdb127c1a20d6570682",
		"8bf29f79659fb52ba3f8305b4610c195bef0eb80debf4018b68f43be01066a441e243a5634f4f327d7b7d20848490a04c3f7355d2bbe512c5f1035c775dfad3d",
	),
	1000,
)

func TestMsgIBCAccountConnection_Route(t *testing.T) {
	require.Equal(t, "ibcprofiles", validConnectionMsg.Route())
}

func TestMsgIBCAccountConnection_Type(t *testing.T) {
	require.Equal(t, "ibc_account_connection", validConnectionMsg.Type())
}

func TestMsgIBCAccountConnection_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgCreateIBCAccountConnection","value":{"channel_id":"desmos-0","packet":{"destination_address":"cosmos1qdasq0mzpajknaaj32kf9lk5nmcy8g65mddd4p","destination_signature":"1a18d5f012ce0e8258fd3455c01b48249bb019231e416c4323ab2bb170b4ad0951b370138d2ea69a376feb942d3c619c9152d63a6d2e0232aaff77162df66636","source_address":"cosmos1u33w3u4ler4654phrpt6xqvh92ch0v6mcjrj97","source_chain_prefix":"cosmos","source_pub_key":"02bcd0738e3b7e0f6650c8e6eb10bd4266fd6818c92a0283b4cb0884f046051c3e","source_signature":"143bbd3131d76232f973f84a9ea7be751243044315056dbecb968942da97e474401caa1c8f2c4ce5e48052cf44066717f2166c21a7277de9911d75c57eca598d"},"port":"ibc-profiles","timeout_timestamp":1000}}`
	require.Equal(t, expected, string(validConnectionMsg.GetSignBytes()))
}

func TestMsgIBCAccountConnection_GetSigner(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(validConnectionMsg.Packet.SourceAddress)
	require.Equal(t, []sdk.AccAddress{addr}, validConnectionMsg.GetSigners())
}

func TestMsgIBCAccountConnection_ValidateBasic(t *testing.T) {
	tests := []struct {
		name    string
		msg     *types.MsgCreateIBCAccountConnection
		expPass bool
		expErr  error
	}{
		{
			name: "Valid msg",
			msg: types.NewMsgCreateIBCAccountConnection(
				"ibc-profiles",
				"desmos-0",
				types.NewIBCAccountConnectionPacketData(
					"cosmos",
					"cosmos1u33w3u4ler4654phrpt6xqvh92ch0v6mcjrj97",
					"02bcd0738e3b7e0f6650c8e6eb10bd4266fd6818c92a0283b4cb0884f046051c3e",
					"cosmos1qdasq0mzpajknaaj32kf9lk5nmcy8g65mddd4p",
					"143bbd3131d76232f973f84a9ea7be751243044315056dbecb968942da97e474401caa1c8f2c4ce5e48052cf44066717f2166c21a7277de9911d75c57eca598d",
					"1a18d5f012ce0e8258fd3455c01b48249bb019231e416c4323ab2bb170b4ad0951b370138d2ea69a376feb942d3c619c9152d63a6d2e0232aaff77162df66636",
				),
				1000,
			),
			expPass: true,
		},
		{
			name: "Invalid port",
			msg: types.NewMsgCreateIBCAccountConnection(
				"(invalidport)",
				"desmos-0",
				types.NewIBCAccountConnectionPacketData(
					"cosmos",
					"cosmos1u33w3u4ler4654phrpt6xqvh92ch0v6mcjrj97",
					"02bcd0738e3b7e0f6650c8e6eb10bd4266fd6818c92a0283b4cb0884f046051c3e",
					"cosmos1qdasq0mzpajknaaj32kf9lk5nmcy8g65mddd4p",
					"143bbd3131d76232f973f84a9ea7be751243044315056dbecb968942da97e474401caa1c8f2c4ce5e48052cf44066717f2166c21a7277de9911d75c57eca598d",
					"1a18d5f012ce0e8258fd3455c01b48249bb019231e416c4323ab2bb170b4ad0951b370138d2ea69a376feb942d3c619c9152d63a6d2e0232aaff77162df66636",
				),
				1000,
			),
			expPass: false,
		},
		{
			name: "Invalid channel",
			msg: types.NewMsgCreateIBCAccountConnection(
				"ibc-profiles",
				"(invalidchannel)",
				types.NewIBCAccountConnectionPacketData(
					"cosmos",
					"cosmos1u33w3u4ler4654phrpt6xqvh92ch0v6mcjrj97",
					"02bcd0738e3b7e0f6650c8e6eb10bd4266fd6818c92a0283b4cb0884f046051c3e",
					"cosmos1qdasq0mzpajknaaj32kf9lk5nmcy8g65mddd4p",
					"143bbd3131d76232f973f84a9ea7be751243044315056dbecb968942da97e474401caa1c8f2c4ce5e48052cf44066717f2166c21a7277de9911d75c57eca598d",
					"1a18d5f012ce0e8258fd3455c01b48249bb019231e416c4323ab2bb170b4ad0951b370138d2ea69a376feb942d3c619c9152d63a6d2e0232aaff77162df66636",
				),
				1000,
			),
			expPass: false,
		},
		{
			name: "Invalid packet",
			msg: types.NewMsgCreateIBCAccountConnection(
				"ibc-profiles",
				"desmos-0",
				types.NewIBCAccountConnectionPacketData(
					"",
					"cosmos1u33w3u4ler4654phrpt6xqvh92ch0v6mcjrj97",
					"02bcd0738e3b7e0f6650c8e6eb10bd4266fd6818c92a0283b4cb0884f046051c3e",
					"cosmos1qdasq0mzpajknaaj32kf9lk5nmcy8g65mddd4p",
					"143bbd3131d76232f973f84a9ea7be751243044315056dbecb968942da97e474401caa1c8f2c4ce5e48052cf44066717f2166c21a7277de9911d75c57eca598d",
					"1a18d5f012ce0e8258fd3455c01b48249bb019231e416c4323ab2bb170b4ad0951b370138d2ea69a376feb942d3c619c9152d63a6d2e0232aaff77162df66636",
				),
				1000,
			),
			expPass: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if !test.expPass {
				require.Error(t, test.msg.ValidateBasic())
				if test.expErr != nil {
					require.Equal(t, test.expErr, test.msg.ValidateBasic())
				}

			} else {
				require.NoError(t, test.msg.ValidateBasic())
			}
		})
	}
}

// ___________________________________________________________________________________________________________________

func TestMsgIBCAccountLink_Route(t *testing.T) {
	require.Equal(t, "ibcprofiles", validLinkMsg.Route())
}

func TestMsgIBCAccountLink_Type(t *testing.T) {
	require.Equal(t, "ibc_account_link", validLinkMsg.Type())
}

func TestMsgIBCAccountLink_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgCreateIBCAccountLink","value":{"channel_id":"desmos-0","packet":{"signature":"8bf29f79659fb52ba3f8305b4610c195bef0eb80debf4018b68f43be01066a441e243a5634f4f327d7b7d20848490a04c3f7355d2bbe512c5f1035c775dfad3d","source_address":"cosmos1fc6dg2f85hsd7nc7c4ng89uad7hzh6thpwa6xf","source_chain_prefix":"cosmos","source_pub_key":"039d2b5c2e5be3f0f3e1665f622734aa21bbfcb1bc0099fbdb127c1a20d6570682"},"port":"ibc-profiles","timeout_timestamp":"1000"}}`
	require.Equal(t, expected, string(validLinkMsg.GetSignBytes()))
}

func TestMsgIBCAccountLink_GetSigner(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(validLinkMsg.Packet.SourceAddress)
	require.Equal(t, []sdk.AccAddress{addr}, validLinkMsg.GetSigners())
}

func TestMsgIBCAccountLink_ValidateBasic(t *testing.T) {
	tests := []struct {
		name    string
		msg     *types.MsgCreateIBCAccountLink
		expPass bool
		expErr  error
	}{
		{
			name: "Valid msg",
			msg: types.NewMsgCreateIBCAccountLink(
				"ibc-profiles",
				"desmos-0",
				types.NewIBCAccountLinkPacketData(
					"cosmos",
					"cosmos1fc6dg2f85hsd7nc7c4ng89uad7hzh6thpwa6xf",
					"039d2b5c2e5be3f0f3e1665f622734aa21bbfcb1bc0099fbdb127c1a20d6570682",
					"8bf29f79659fb52ba3f8305b4610c195bef0eb80debf4018b68f43be01066a441e243a5634f4f327d7b7d20848490a04c3f7355d2bbe512c5f1035c775dfad3d",
				),
				1000,
			),
			expPass: true,
		},
		{
			name: "Invalid port",
			msg: types.NewMsgCreateIBCAccountLink(
				"(invalidport)",
				"desmos-0",
				types.NewIBCAccountLinkPacketData(
					"cosmos",
					"cosmos1fc6dg2f85hsd7nc7c4ng89uad7hzh6thpwa6xf",
					"039d2b5c2e5be3f0f3e1665f622734aa21bbfcb1bc0099fbdb127c1a20d6570682",
					"8bf29f79659fb52ba3f8305b4610c195bef0eb80debf4018b68f43be01066a441e243a5634f4f327d7b7d20848490a04c3f7355d2bbe512c5f1035c775dfad3d",
				),
				1000,
			),
			expPass: false,
		},
		{
			name: "Invalid channel",
			msg: types.NewMsgCreateIBCAccountLink(
				"ibc-profiles",
				"(invalidchannel)",
				types.NewIBCAccountLinkPacketData(
					"cosmos",
					"cosmos1fc6dg2f85hsd7nc7c4ng89uad7hzh6thpwa6xf",
					"039d2b5c2e5be3f0f3e1665f622734aa21bbfcb1bc0099fbdb127c1a20d6570682",
					"8bf29f79659fb52ba3f8305b4610c195bef0eb80debf4018b68f43be01066a441e243a5634f4f327d7b7d20848490a04c3f7355d2bbe512c5f1035c775dfad3d",
				),
				1000,
			),
			expPass: false,
		},
		{
			name: "Invalid packet",
			msg: types.NewMsgCreateIBCAccountLink(
				"ibc-profiles",
				"desmos-0",
				types.NewIBCAccountLinkPacketData(
					"",
					"cosmos1fc6dg2f85hsd7nc7c4ng89uad7hzh6thpwa6xf",
					"039d2b5c2e5be3f0f3e1665f622734aa21bbfcb1bc0099fbdb127c1a20d6570682",
					"8bf29f79659fb52ba3f8305b4610c195bef0eb80debf4018b68f43be01066a441e243a5634f4f327d7b7d20848490a04c3f7355d2bbe512c5f1035c775dfad3d",
				),
				1000,
			),
			expPass: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if !test.expPass {
				require.Error(t, test.msg.ValidateBasic())
				if test.expErr != nil {
					require.Equal(t, test.expErr, test.msg.ValidateBasic())
				}
			} else {
				require.NoError(t, test.msg.ValidateBasic())
			}
		})
	}
}
