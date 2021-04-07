package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/links/types"
	"github.com/stretchr/testify/require"
)

var validMsg = types.NewMsgIBCAccountConnection(
	"links",
	"desmos-0",
	1000,
	"cosmos",
	"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
	"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
	"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn",
	"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
	"fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561",
)

func TestMsgIBCAccountConnection_Route(t *testing.T) {
	require.Equal(t, "links", validMsg.Route())
}

func TestMsgIBCAccountConnection_Type(t *testing.T) {
	require.Equal(t, "ibc_account_connection", validMsg.Type())
}

func TestMsgIBCAccountConnection_GetSignBytes(t *testing.T) {
	expected := `{"channel_id":"desmos-0","destination_address":"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn","destination_signature":"fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561","port":"links","source_address":"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn","source_chain_prefix":"cosmos","source_pub_key":"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b","source_signature":"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2","timeout_timestamp":1000}`
	require.Equal(t, expected, string(validMsg.GetSignBytes()))
}

func TestMsgIBCAccountConnection_GetSigner(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(validMsg.SourceAddress)
	require.Equal(t, []sdk.AccAddress{addr}, validMsg.GetSigners())
}
