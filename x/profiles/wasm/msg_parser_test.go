package wasm_test

import (
	"encoding/json"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/app"
	"github.com/desmos-labs/desmos/v3/testutil/profilestesting"
	"github.com/desmos-labs/desmos/v3/x/profiles/types"
	"github.com/desmos-labs/desmos/v3/x/profiles/wasm"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

func TestMsgsParser_ParseCustomMsgs(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	parser := wasm.NewWasmMsgParser(cdc)
	contractAddr, err := sdk.AccAddressFromBech32("cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr")
	require.NoError(t, err)

	wrongMsgBz, err := json.Marshal(subspacestypes.SubspacesMsg{DeleteSubspace: nil})
	require.NoError(t, err)

	testCases := []struct {
		name      string
		msg       json.RawMessage
		shouldErr bool
		expMsgs   []sdk.Msg
	}{
		{
			name:      "wrong module message returns error",
			msg:       wrongMsgBz,
			shouldErr: true,
			expMsgs:   nil,
		},
		{
			name: "save profile json message is parsed correctly",
			msg: buildSaveProfileRequest(cdc, types.NewMsgSaveProfile(
				"test",
				"test",
				"test",
				"test",
				"test",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{
				types.NewMsgSaveProfile(
					"test",
					"test",
					"test",
					"test",
					"test",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				),
			},
		},
		{
			name: "delete profile json message is parsed correctly",
			msg: buildDeleteProfileRequest(cdc, types.NewMsgDeleteProfile(
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{
				types.NewMsgDeleteProfile(
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				),
			},
		},
		{
			name: "request dtag transfer json message is parsed correctly",
			msg: buildRequestDTagTransferRequest(cdc, types.NewMsgRequestDTagTransfer(
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{
				types.NewMsgRequestDTagTransfer(
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				),
			},
		},
		{
			name: "accept dtag transfer json message is parsed correctly",
			msg: buildAcceptDTagTransferRequest(cdc, types.NewMsgAcceptDTagTransferRequest(
				"dtag",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{
				types.NewMsgAcceptDTagTransferRequest(
					"dtag",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				),
			},
		},
		{
			name: "refuse dtag transfer json message is parsed correctly",
			msg: buildRefuseDTagTransferRequest(cdc, types.NewMsgRefuseDTagTransferRequest(
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{
				types.NewMsgRefuseDTagTransferRequest(
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				),
			},
		},
		{
			name: "cancel dtag transfer json message is parsed correctly",
			msg: buildCancelDTagTransferRequest(cdc, types.NewMsgCancelDTagTransferRequest(
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{
				types.NewMsgCancelDTagTransferRequest(
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				),
			},
		},
		{
			name: "link chain account json message is parsed correctly",
			msg: buildLinkChainAccountRequest(cdc, types.NewMsgLinkChainAccount(
				types.NewBech32Address("cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0", "cosmos"),
				types.NewProof(
					profilestesting.PubKeyFromBech32("cosmospub1addwnpepq0j8zw4t6tg3v8gh7d2d799gjhue7ewwmpg2hwr77f9kuuyzgqtrw5r6wec"),
					profilestesting.SingleSignatureProtoFromHex("ad112abb30e5240c7b9d21b4cc5421d76cfadfcd5977cca262523b5f5bc759457d4aa6d5c1eb6223db104b47aa1f222468be8eb5bb2762b971622ac5b96351b5"),
					"74657874",
				),
				types.NewChainConfig("cosmos"),
				"cosmos1u9hgsqfpe3snftr7p7fsyja3wtlmj2sgf2w9yl",
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{types.NewMsgLinkChainAccount(
				types.NewBech32Address("cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0", "cosmos"),
				types.NewProof(
					profilestesting.PubKeyFromBech32("cosmospub1addwnpepq0j8zw4t6tg3v8gh7d2d799gjhue7ewwmpg2hwr77f9kuuyzgqtrw5r6wec"),
					profilestesting.SingleSignatureProtoFromHex("ad112abb30e5240c7b9d21b4cc5421d76cfadfcd5977cca262523b5f5bc759457d4aa6d5c1eb6223db104b47aa1f222468be8eb5bb2762b971622ac5b96351b5"),
					"74657874",
				),
				types.NewChainConfig("cosmos"),
				"cosmos1u9hgsqfpe3snftr7p7fsyja3wtlmj2sgf2w9yl"),
			},
		},
		{
			name: "link application json message is parsed correctly",
			msg: buildLinkApplicationRequest(cdc, types.NewMsgLinkApplication(
				types.NewData("twitter", "twitteruser"),
				"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				types.IBCPortID,
				"channel-0",
				clienttypes.NewHeight(0, 1000),
				0,
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{
				types.NewMsgLinkApplication(
					types.NewData("twitter", "twitteruser"),
					"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					types.IBCPortID,
					"channel-0",
					clienttypes.NewHeight(0, 1000),
					0,
				),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			msgs, err := parser.ParseCustomMsgs(contractAddr, tc.msg)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tc.expMsgs, msgs)
		})
	}
}
