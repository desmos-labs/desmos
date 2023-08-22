package wasm_test

import (
	"encoding/json"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v6/app"
	profilestypes "github.com/desmos-labs/desmos/v6/x/profiles/types"
	"github.com/desmos-labs/desmos/v6/x/relationships/types"
	"github.com/desmos-labs/desmos/v6/x/relationships/wasm"
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
			name: "create relationship json message parsed correctly",
			msg: buildCreateRelationshipRequest(cdc, types.NewMsgCreateRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				1,
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{
				types.NewMsgCreateRelationship(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					1,
				)},
		},
		{
			name: "delete relationship json message parsed correctly",
			msg: buildDeleteRelationshipRequest(cdc, types.NewMsgDeleteRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				1,
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{
				types.NewMsgDeleteRelationship(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					1,
				)},
		},
		{
			name: "block user json message parsed correctly",
			msg: buildBlockUserRequest(cdc, types.NewMsgBlockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"reason",
				0,
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{
				types.NewMsgBlockUser(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"reason",
					0,
				),
			},
		},
		{
			name: "unblock user json message parsed correctly",
			msg: buildUnblockUserRequest(cdc, types.NewMsgUnblockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				0,
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{
				types.NewMsgUnblockUser(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
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
