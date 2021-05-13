package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/ibc/profiles/types"
	"github.com/stretchr/testify/require"
)

var validConnectionMsg = types.NewMsgCreateIBCAccountConnection(
	"links",
	"desmos-0",
	types.NewIBCAccountConnectionPacketData(
		"cosmos",
		"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
		"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
		"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn",
		"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
		"fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561",
	),
	1000,
)

var validLinkMsg = types.NewMsgCreateIBCAccountLink(
	"links",
	"desmos-0",
	types.NewIBCAccountLinkPacketData(
		"cosmos",
		"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
		"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
		"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
	),
	1000,
)

func TestMsgIBCAccountConnection_Route(t *testing.T) {
	require.Equal(t, "links", validConnectionMsg.Route())
}

func TestMsgIBCAccountConnection_Type(t *testing.T) {
	require.Equal(t, "ibc_account_connection", validConnectionMsg.Type())
}

func TestMsgIBCAccountConnection_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgCreateIBCAccountConnection","value":{"channel_id":"desmos-0","packet":{"destination_address":"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn","destination_signature":"fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561","source_address":"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn","source_chain_prefix":"cosmos","source_pub_key":"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b","source_signature":"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2"},"port":"links","timeout_timestamp":1000}}`
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
				"links",
				"desmos-0",
				types.NewIBCAccountConnectionPacketData(
					"cosmos",
					"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
					"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
					"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn",
					"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
					"fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561",
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
					"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
					"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
					"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn",
					"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
					"fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561",
				),
				1000,
			),
			expPass: false,
		},
		{
			name: "Invalid channel",
			msg: types.NewMsgCreateIBCAccountConnection(
				"links",
				"(invalidchannel)",
				types.NewIBCAccountConnectionPacketData(
					"cosmos",
					"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
					"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
					"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn",
					"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
					"fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561",
				),
				1000,
			),
			expPass: false,
		},
		{
			name: "Invalid packet",
			msg: types.NewMsgCreateIBCAccountConnection(
				"links",
				"desmos-0",
				types.NewIBCAccountConnectionPacketData(
					"",
					"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
					"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
					"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn",
					"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
					"fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561",
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
	require.Equal(t, "links", validLinkMsg.Route())
}

func TestMsgIBCAccountLink_Type(t *testing.T) {
	require.Equal(t, "ibc_account_link", validLinkMsg.Type())
}

func TestMsgIBCAccountLink_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgCreateIBCAccountLink","value":{"channel_id":"desmos-0","packet":{"signature":"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2","source_address":"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn","source_chain_prefix":"cosmos","source_pub_key":"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b"},"port":"links","timeout_timestamp":"1000"}}`
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
				"links",
				"desmos-0",
				types.NewIBCAccountLinkPacketData(
					"cosmos",
					"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
					"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
					"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
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
					"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
					"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
					"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
				),
				1000,
			),
			expPass: false,
		},
		{
			name: "Invalid channel",
			msg: types.NewMsgCreateIBCAccountLink(
				"links",
				"(invalidchannel)",
				types.NewIBCAccountLinkPacketData(
					"cosmos",
					"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
					"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
					"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
				),
				1000,
			),
			expPass: false,
		},
		{
			name: "Invalid packet",
			msg: types.NewMsgCreateIBCAccountLink(
				"links",
				"desmos-0",
				types.NewIBCAccountLinkPacketData(
					"",
					"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
					"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
					"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
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
