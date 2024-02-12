package wasm_test

import (
	"encoding/json"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v7/app"
	profilestypes "github.com/desmos-labs/desmos/v7/x/profiles/types"
	"github.com/desmos-labs/desmos/v7/x/subspaces/types"
	"github.com/desmos-labs/desmos/v7/x/subspaces/wasm"
)

func TestMsgsParser_ParseCustomMsgs(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	parser := wasm.NewWasmMsgParser(cdc)
	contractAddr, err := sdk.AccAddressFromBech32("cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr")
	require.NoError(t, err)

	wrongMsgBz, err := json.Marshal(profilestypes.ProfilesMsg{DeleteProfile: nil})
	require.NoError(t, err)

	expiration := time.Date(2023, 1, 11, 1, 1, 1, 1, time.UTC)

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
			name: "create subspace message is parsed correctly",
			msg: buildCreateSubspaceRequest(cdc,
				types.NewMsgCreateSubspace(
					"test",
					"test",
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
			)},
		},
		{
			name: "edit subspace message is parsed correctly",
			msg: buildEditSubspaceRequest(cdc,
				types.NewMsgEditSubspace(
					1,
					"test",
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
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)},
		},
		{
			name: "delete subspace message is parsed correctly",
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
			name: "create user group message is parsed correctly",
			msg: buildCreateUserGroupRequest(cdc,
				types.NewMsgCreateUserGroup(
					1,
					0,
					"test",
					"",
					types.NewPermissions(types.PermissionEverything),
					[]string{"cosmos16yhs7fgqnf6fjm4tftv66g2smtmee62wyg780l"},
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				),
			),
			shouldErr: false,
			expMsgs: []sdk.Msg{types.NewMsgCreateUserGroup(
				1,
				0,
				"test",
				"",
				types.NewPermissions(types.PermissionEverything),
				[]string{"cosmos16yhs7fgqnf6fjm4tftv66g2smtmee62wyg780l"},
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)},
		},
		{
			name: "edit user group message is parsed correctly",
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
			name: "set user group message is parsed correctly",
			msg: buildSetUserGroupPermissionsRequest(cdc,
				types.NewMsgSetUserGroupPermissions(
					1,
					1,
					types.NewPermissions(types.PermissionEverything),
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				),
			),
			shouldErr: false,
			expMsgs: []sdk.Msg{
				types.NewMsgSetUserGroupPermissions(
					1,
					1,
					types.NewPermissions(types.PermissionEverything),
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				),
			},
		},
		{
			name: "delete user group message is parsed correctly",
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
			name: "add user to user group message is parsed correctly",
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
			name: "remove user from user group message is parsed correctly",
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
			name: "set user permissions message is parsed correctly",
			msg: buildSetUserPermissionsRequest(cdc,
				types.NewMsgSetUserPermissions(
					1,
					0,
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					types.NewPermissions(types.PermissionEverything),
					"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
				),
			),
			shouldErr: false,
			expMsgs: []sdk.Msg{
				types.NewMsgSetUserPermissions(
					1,
					0,
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					types.NewPermissions(types.PermissionEverything),
					"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
				),
			},
		},
		{
			name: "grant treasury authorization message is parsed correctly",
			msg: buildGrantTreasuryAuthorizationRequest(cdc,
				types.NewMsgGrantTreasuryAuthorization(
					1,
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
					authz.NewGenericAuthorization("/cosmos.bank.v1beta1.MsgSend"),
					&expiration,
				),
			),
			shouldErr: false,
			expMsgs: []sdk.Msg{
				types.NewMsgGrantTreasuryAuthorization(
					1,
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
					authz.NewGenericAuthorization("/cosmos.bank.v1beta1.MsgSend"),
					&expiration,
				),
			},
		},
		{
			name: "revoke treasury authorization message is parsed correctly",
			msg: buildRevokeTreasuryAuthorizationRequest(cdc,
				types.NewMsgRevokeTreasuryAuthorization(
					1,
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
					"/cosmos.bank.v1beta1.MsgSend",
				),
			),
			shouldErr: false,
			expMsgs: []sdk.Msg{
				types.NewMsgRevokeTreasuryAuthorization(
					1,
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
					"/cosmos.bank.v1beta1.MsgSend",
				),
			},
		},
		{
			name: "grant user allowance message is parsed correctly",
			msg: buildGrantAllowanceRequest(cdc, types.NewMsgGrantAllowance(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				types.NewUserGrantee("cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0"),
				&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10)))},
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{
				types.NewMsgGrantAllowance(
					1,
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					types.NewUserGrantee("cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10)))},
				),
			},
		},
		{
			name: "revoke user allowance message is parsed correctly",
			msg: buildRevokeAllowanceRequest(cdc, types.NewMsgRevokeAllowance(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				types.NewUserGrantee("cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0"),
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{
				types.NewMsgRevokeAllowance(
					1,
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					types.NewUserGrantee("cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0"),
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
