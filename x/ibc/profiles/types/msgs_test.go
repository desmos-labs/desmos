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
		"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
		"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
		"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq",
		"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
		"42dd1f8d98c5de91a12259cf46098104132f69b61eaa24e112bf504d17e1a0b71274dad981bbb4a13dc440905a19be92eaf4497940751f431c530cc4d68e78b0",
	),
	1000,
)

var validLinkMsg = types.NewMsgCreateIBCAccountLink(
	"ibc-profiles",
	"desmos-0",
	types.NewIBCAccountLinkPacketData(
		"cosmos",
		"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
		"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
		"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
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
	expected := `{"type":"desmos/MsgCreateIBCAccountConnection","value":{"channel_id":"desmos-0","packet":{"destination_address":"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq","destination_signature":"42dd1f8d98c5de91a12259cf46098104132f69b61eaa24e112bf504d17e1a0b71274dad981bbb4a13dc440905a19be92eaf4497940751f431c530cc4d68e78b0","source_address":"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70","source_chain_prefix":"cosmos","source_pub_key":"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561","source_signature":"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c"},"port":"ibc-profiles","timeout_timestamp":1000}}`
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
					"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
					"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
					"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq",
					"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
					"42dd1f8d98c5de91a12259cf46098104132f69b61eaa24e112bf504d17e1a0b71274dad981bbb4a13dc440905a19be92eaf4497940751f431c530cc4d68e78b0",
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
					"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
					"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
					"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq",
					"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
					"42dd1f8d98c5de91a12259cf46098104132f69b61eaa24e112bf504d17e1a0b71274dad981bbb4a13dc440905a19be92eaf4497940751f431c530cc4d68e78b0",
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
					"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
					"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
					"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq",
					"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
					"42dd1f8d98c5de91a12259cf46098104132f69b61eaa24e112bf504d17e1a0b71274dad981bbb4a13dc440905a19be92eaf4497940751f431c530cc4d68e78b0",
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
					"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
					"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
					"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq",
					"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
					"42dd1f8d98c5de91a12259cf46098104132f69b61eaa24e112bf504d17e1a0b71274dad981bbb4a13dc440905a19be92eaf4497940751f431c530cc4d68e78b0",
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
	expected := `{"type":"desmos/MsgCreateIBCAccountLink","value":{"channel_id":"desmos-0","packet":{"signature":"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c","source_address":"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70","source_chain_prefix":"cosmos","source_pub_key":"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561"},"port":"ibc-profiles","timeout_timestamp":"1000"}}`
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
					"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
					"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
					"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
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
					"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
					"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
					"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
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
					"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
					"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
					"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
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
					"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
					"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
					"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
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
