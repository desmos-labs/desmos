package wasm_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/app"
	profilestypes "github.com/desmos-labs/desmos/v3/x/profiles/types"
	"github.com/desmos-labs/desmos/v3/x/reactions/types"
	"github.com/desmos-labs/desmos/v3/x/reactions/wasm"
)

func TestMsgsParser_ParseCustomMsgs(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	parser := wasm.NewWasmMsgParser(cdc)
	contractAddr, err := sdk.AccAddressFromBech32("cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr")
	require.NoError(t, err)

	wrongMsgBz, err := json.Marshal(profilestypes.ProfilesMsg{DeleteProfile: nil})
	require.NoError(t, err)

	testCases := []struct {
		name      string
		msg       json.RawMessage
		shouldErr bool
		expMsgs   []sdk.Msg
	}{
		{
			name:      "parse wrong module message returns error",
			msg:       wrongMsgBz,
			shouldErr: true,
			expMsgs:   nil,
		},
		{
			name: "add reaction json message is parsed correctly",
			msg: buildAddReactionRequest(cdc, types.NewMsgAddReaction(
				1,
				1,
				types.NewRegisteredReactionValue(1),
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{types.NewMsgAddReaction(
				1,
				1,
				types.NewRegisteredReactionValue(1),
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)},
		},
		{
			name: "remove reaction json message is parsed correctly",
			msg: buildRemoveReactionRequest(cdc, types.NewMsgRemoveReaction(
				1,
				1,
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{types.NewMsgRemoveReaction(
				1,
				1,
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69")},
		},
		{
			name: "add registered reaction message is parsed correctly",
			msg: buildAddRegisteredReactionRequest(cdc, types.NewMsgAddRegisteredReaction(
				1,
				"shorthand_code",
				"display_value",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{types.NewMsgAddRegisteredReaction(
				1,
				"shorthand_code",
				"display_value",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)},
		},
		{
			name: "edit registered reaction message is parsed correctly",
			msg: buildEditRegisteredReactionRequest(cdc, types.NewMsgEditRegisteredReaction(
				1,
				1,
				"shorthand_code",
				"display_value",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{types.NewMsgEditRegisteredReaction(
				1,
				1,
				"shorthand_code",
				"display_value",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)},
		},
		{
			name: "remove registered reaction message is parsed correctly",
			msg: buildRemoveRegisteredReactionRequest(cdc, types.NewMsgRemoveRegisteredReaction(
				1,
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{types.NewMsgRemoveRegisteredReaction(
				1,
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)},
		},
		{
			name: "set reaction parameter message is parsed correctly",
			msg: buildSetReactionsParamsRequest(cdc, types.NewMsgSetReactionsParams(
				1,
				types.NewRegisteredReactionValueParams(true),
				types.NewFreeTextValueParams(true, 100, ""),
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{types.NewMsgSetReactionsParams(
				1,
				types.NewRegisteredReactionValueParams(true),
				types.NewFreeTextValueParams(true, 100, ""),
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)},
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
