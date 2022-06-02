package wasm_test

import (
	"encoding/json"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v3/app"
	profilestypes "github.com/desmos-labs/desmos/v3/x/profiles/types"
	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
	"github.com/desmos-labs/desmos/v3/x/subspaces/wasm"
	"github.com/stretchr/testify/require"
)

func TestMsgsParser_ParseCustomMsgs(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	parser := wasm.NewWasmMsgParser(cdc)
	contractAddr, err := sdk.AccAddressFromBech32("cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr")
	require.NoError(t, err)

	wrongMsgBz, err := json.Marshal(profilestypes.ProfilesMsg{DeleteProfile: cdc.MustMarshalJSON(profilestypes.NewMsgDeleteProfile(""))})
	require.NoError(t, err)

	testCases := []struct {
		name      string
		msg       json.RawMessage
		shouldErr bool
		expMsgs   []sdk.Msg
	}{
		{
			name:      "Parse wrong module message returns error",
			msg:       wrongMsgBz,
			shouldErr: true,
			expMsgs:   nil,
		},
		{
			name: "Create subspace json message is parsed correctly",
			msg: buildCreateSubspaceRequest(cdc,
				types.NewMsgCreateSubspace(
					"test",
					"test",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				),
			),
			shouldErr: false,
			expMsgs: []sdk.Msg{types.NewMsgCreateSubspace(
				"test",
				"test",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)},
		},
		{
			name: "Edit subspace json message is parsed correctly",
			msg: buildEditSubspaceRequest(cdc,
				types.NewMsgEditSubspace(
					1,
					"test",
					"",
					"",
					"",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				),
			),
			shouldErr: false,
			expMsgs: []sdk.Msg{types.NewMsgEditSubspace(
				1,
				"test",
				"",
				"",
				"",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)},
		},
		{
			name: "Delete subspace message is parsed correctly",
			msg: buildDeleteSubspaceRequest(cdc,
				types.NewMsgDeleteSubspace(
					1,
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				),
			),
			shouldErr: false,
			expMsgs: []sdk.Msg{types.NewMsgDeleteSubspace(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)},
		},
		{
			name: "Create user group message is parsed correctly",
			msg: buildCreateUserGroupRequest(cdc,
				types.NewMsgCreateUserGroup(
					1,
					"test",
					"",
					10,
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				),
			),
			shouldErr: false,
			expMsgs: []sdk.Msg{types.NewMsgCreateUserGroup(
				1,
				"test",
				"",
				10,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)},
		},
		{
			name: "Edit user group message is parsed correctly",
			msg: buildEditUserGroupRequest(cdc,
				types.NewMsgEditUserGroup(
					1,
					1,
					"test",
					"",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				),
			),
			shouldErr: false,
			expMsgs: []sdk.Msg{
				types.NewMsgEditUserGroup(
					1,
					1,
					"test",
					"",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				),
			},
		},
		{
			name: "Set user group message is parsed correctly",
			msg: buildSetUserGroupPermissionsRequest(cdc,
				types.NewMsgSetUserGroupPermissions(
					1,
					1,
					types.PermissionNothing,
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				),
			),
			shouldErr: false,
			expMsgs: []sdk.Msg{
				types.NewMsgSetUserGroupPermissions(
					1,
					1,
					types.PermissionNothing,
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				),
			},
		},
		{
			name: "Delete user group message is parsed correctly",
			msg: buildDeleteUserGroupRequest(cdc,
				types.NewMsgDeleteUserGroup(
					1,
					1,
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				),
			),
			shouldErr: false,
			expMsgs: []sdk.Msg{
				types.NewMsgDeleteUserGroup(
					1,
					1,
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				),
			},
		},
		{
			name: "Add user to user group message is parsed correctly",
			msg: buildAddUserToGroupRequest(cdc,
				types.NewMsgAddUserToUserGroup(
					1,
					1,
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				),
			),
			shouldErr: false,
			expMsgs: []sdk.Msg{
				types.NewMsgAddUserToUserGroup(
					1,
					1,
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				),
			},
		},
		{
			name: "Remove user from user group message is parsed correctly",
			msg: buildRemoveUserFromUserGroupRequest(cdc,
				types.NewMsgRemoveUserFromUserGroup(
					1,
					1,
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				),
			),
			shouldErr: false,
			expMsgs: []sdk.Msg{
				types.NewMsgRemoveUserFromUserGroup(
					1,
					1,
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				),
			},
		},
		{
			name: "Set user permissions message is parsed correctly",
			msg: buildSetUserPermissionsRequest(cdc,
				types.NewMsgSetUserPermissions(
					1,
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					types.PermissionNothing,
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				),
			),
			shouldErr: false,
			expMsgs: []sdk.Msg{
				types.NewMsgSetUserPermissions(
					1,
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					types.PermissionNothing,
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
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