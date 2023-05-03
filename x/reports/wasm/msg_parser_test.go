package wasm_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v5/app"
	profilestypes "github.com/desmos-labs/desmos/v5/x/profiles/types"
	"github.com/desmos-labs/desmos/v5/x/reports/types"
	"github.com/desmos-labs/desmos/v5/x/reports/wasm"
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
			name: "create report json message is parsed correctly",
			msg: buildCreateReportRequest(cdc,
				types.NewMsgCreateReport(
					1,
					[]uint32{1},
					"test",
					types.NewPostTarget(1),
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				),
			),
			shouldErr: false,
			expMsgs: []sdk.Msg{types.NewMsgCreateReport(
				1,
				[]uint32{1},
				"test",
				types.NewPostTarget(1),
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)},
		},
		{
			name: "delete report json message is parsed correctly",
			msg: buildDeleteReportRequest(cdc, types.NewMsgDeleteReport(
				1,
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{types.NewMsgDeleteReport(
				1,
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69")},
		},
		{
			name: "support standard reason message is parsed correctly",
			msg: buildSupportStandardReasonRequest(cdc, types.NewMsgSupportStandardReason(
				1,
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{types.NewMsgSupportStandardReason(
				1,
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)},
		},
		{
			name: "add reason message is parsed correctly",
			msg: buildAddReasonRequest(cdc, types.NewMsgAddReason(
				1,
				"test",
				"test",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{types.NewMsgAddReason(
				1,
				"test",
				"test",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)},
		},
		{
			name: "remove reason message is parsed correctly",
			msg: buildRemoveReasonRequest(cdc, types.NewMsgRemoveReason(
				1,
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{types.NewMsgRemoveReason(
				1,
				1,
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
