package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/links/types"
	"github.com/stretchr/testify/require"
)

var validConnectionMsg = types.NewMsgCreateIBCAccountConnection(
	"links",
	"desmos-0",
	1000,
	"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
	"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
	"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn",
	"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
	"fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561",
)

var validLinkMsg = types.NewMsgCreateIBCAccountLink(
	"links",
	"desmos-0",
	1000,
	"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
	"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
	"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
)

func TestMsgIBCAccountConnection_Route(t *testing.T) {
	require.Equal(t, "links", validConnectionMsg.Route())
}

func TestMsgIBCAccountConnection_Type(t *testing.T) {
	require.Equal(t, "ibc_account_connection", validConnectionMsg.Type())
}

func TestMsgIBCAccountConnection_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgCreateIBCAccountConnection","value":{"channel_id":"desmos-0","destination_address":"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn","destination_signature":"fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561","port":"links","source_address":"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn","source_pub_key":"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b","source_signature":"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2","timeout_timestamp":1000}}`
	require.Equal(t, expected, string(validConnectionMsg.GetSignBytes()))
}

func TestMsgIBCAccountConnection_GetSigner(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(validConnectionMsg.SourceAddress)
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
				1000,
				"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
				"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
				"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
				"fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561",
			),
			expPass: true,
		},
		{
			name: "Invalid port",
			msg: types.NewMsgCreateIBCAccountConnection(
				"(invalidport)",
				"desmos-0",
				1000,
				"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
				"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
				"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
				"fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561",
			),
			expPass: false,
		},
		{
			name: "Invalid channel",
			msg: types.NewMsgCreateIBCAccountConnection(
				"links",
				"(invalidchannel)",
				1000,
				"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
				"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
				"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
				"fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561",
			),
			expPass: false,
		},
		{
			name: "Invalid src address",
			msg: types.NewMsgCreateIBCAccountConnection(
				"links",
				"desmos-0",
				1000,
				"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lx",
				"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
				"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
				"fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561",
			),
			expPass: false,
		},
		{
			name: "Invalid src pubkey",
			msg: types.NewMsgCreateIBCAccountConnection(
				"links",
				"desmos-0",
				1000,
				"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
				"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c2",
				"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
				"fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561",
			),
			expPass: false,
			expErr:  sdkerrors.Wrap(nil, "invalid source pubkey"),
		},
		{
			name: "Invalid src signature",
			msg: types.NewMsgCreateIBCAccountConnection(
				"links",
				"desmos-0",
				1000,
				"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
				"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
				"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e7",
				"fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561",
			),
			expPass: false,
			expErr:  sdkerrors.Wrap(nil, "invalid source signature"),
		},
		{
			name: "Invalid dest signature",
			msg: types.NewMsgCreateIBCAccountConnection(
				"links",
				"desmos-0",
				1000,
				"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
				"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
				"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
				"fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd",
			),
			expPass: false,
			expErr:  sdkerrors.Wrap(nil, "invalid destination signature"),
		},
		{
			name: "Mismatch src pubkey with address",
			msg: types.NewMsgCreateIBCAccountConnection(
				"links",
				"desmos-0",
				1000,
				"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
				"02466b245623786131225676fbcf4eb5a32c835a8acc733a989af45b0cbbcc0e8",
				"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
				"fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561",
			),
			expPass: false,
			expErr:  sdkerrors.Wrap(nil, "source pubkey and source address are mismatched"),
		},
		{
			name: "Invalid pubkey for signature",
			msg: types.NewMsgCreateIBCAccountConnection(
				"links",
				"desmos-0",
				1000,
				"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
				"02466b245623786131225676fbcf4eb5a32c835a8acc733a989af45b0cbbcc0e8",
				"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
			),
			expPass: false,
			expErr:  sdkerrors.Wrap(nil, "failed to verify source signature"),
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
	expected := `{"type":"desmos/MsgCreateIBCAccountLink","value":{"channel_id":"desmos-0","port":"links","signature":"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2","source_address":"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn","source_pub_key":"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b","timeout_timestamp":"1000"}}`
	require.Equal(t, expected, string(validLinkMsg.GetSignBytes()))
}

func TestMsgIBCAccountLink_GetSigner(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(validLinkMsg.SourceAddress)
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
				1000,
				"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
				"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
			),
			expPass: true,
		},
		{
			name: "Invalid port",
			msg: types.NewMsgCreateIBCAccountLink(
				"(invalidport)",
				"desmos-0",
				1000,
				"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
				"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
			),
			expPass: false,
		},
		{
			name: "Invalid channel",
			msg: types.NewMsgCreateIBCAccountLink(
				"links",
				"(invalidchannel)",
				1000,
				"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
				"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
			),
			expPass: false,
		},
		{
			name: "Invalid src address",
			msg: types.NewMsgCreateIBCAccountLink(
				"links",
				"desmos-0",
				1000,
				"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0",
				"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
			),
			expPass: false,
		},
		{
			name: "Invalid src pubkey",
			msg: types.NewMsgCreateIBCAccountLink(
				"links",
				"desmos-0",
				1000,
				"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
				"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
			),
			expPass: false,
			expErr:  sdkerrors.Wrap(nil, "invalid source pubkey"),
		},
		{
			name: "Invalid src signature",
			msg: types.NewMsgCreateIBCAccountLink(
				"links",
				"desmos-0",
				1000,
				"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
				"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e7",
			),
			expPass: false,
			expErr:  sdkerrors.Wrap(nil, "invalid source signature"),
		},
		{
			name: "Mismatch src pubkey with address",
			msg: types.NewMsgCreateIBCAccountLink(
				"links",
				"desmos-0",
				1000,
				"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
				"02466b245623786131225676fbcf4eb5a32c835a8acc733a989af45b0cbbcc0e8",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
			),
			expPass: false,
			expErr:  sdkerrors.Wrap(nil, "source pubkey and source address are mismatched"),
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
