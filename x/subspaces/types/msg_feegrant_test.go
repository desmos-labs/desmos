package types_test

import (
	"testing"

	"cosmossdk.io/x/feegrant"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v6/x/subspaces/types"
)

var msgGrantAllowance = types.NewMsgGrantAllowance(
	1,
	"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
	types.NewUserGrantee("cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez"),
	&feegrant.BasicAllowance{},
)

func TestMsgGrantAllowance_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgGrantAllowance.Route())
}

func TestMsgGrantAllowance_Type(t *testing.T) {
	require.Equal(t, types.ActionGrantAllowance, msgGrantAllowance.Type())
}

func TestMsgGrantAllowance_ValidateBasic(t *testing.T) {
	granteeAny, err := codectypes.NewAnyWithValue(types.NewUserGrantee("cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez"))
	require.NoError(t, err)

	testCases := []struct {
		name      string
		msg       *types.MsgGrantAllowance
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgGrantAllowance(
				0,
				"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
				types.NewUserGrantee("cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez"),
				&feegrant.BasicAllowance{},
			),
			shouldErr: true,
		},
		{
			name: "invalid granter returns error",
			msg: types.NewMsgGrantAllowance(
				1,
				"",
				types.NewUserGrantee("cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez"),
				&feegrant.BasicAllowance{},
			),
			shouldErr: true,
		},
		{
			name: "invalid grantee returns error",
			msg: types.NewMsgGrantAllowance(
				1,
				"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
				types.NewUserGrantee(""),
				&feegrant.BasicAllowance{},
			),
			shouldErr: true,
		},
		{
			name: "granter self-grant returns error",
			msg: types.NewMsgGrantAllowance(
				1,
				"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
				types.NewUserGrantee("cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0"),
				&feegrant.BasicAllowance{},
			),
			shouldErr: true,
		},
		{
			name: "invalid allowance returns no error",
			msg: &types.MsgGrantAllowance{
				SubspaceID: 1,
				Granter:    "cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
				Grantee:    granteeAny,
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgGrantAllowance,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgGrantAllowance_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgGrantAllowance","value":{"allowance":{"spend_limit":[]},"grantee":{"type":"desmos/UserGrantee","value":{"user":"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez"}},"granter":"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgGrantAllowance.GetSignBytes()))
}

func TestMsgGrantAllowance_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgGrantAllowance.Granter)
	require.Equal(t, []sdk.AccAddress{addr}, msgGrantAllowance.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgRevokeAllowance = types.NewMsgRevokeAllowance(
	1,
	"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
	types.NewUserGrantee("cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez"),
)

func TestMsgRevokeAllowance_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgRevokeAllowance.Route())
}

func TestMsgRevokeAllowance_Type(t *testing.T) {
	require.Equal(t, types.ActionRevokeAllowance, msgRevokeAllowance.Type())
}

func TestMsgRevokeAllowance_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgRevokeAllowance
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgRevokeAllowance(
				0,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				types.NewUserGrantee("cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez"),
			),
			shouldErr: true,
		},
		{
			name: "invalid granter returns error",
			msg: types.NewMsgRevokeAllowance(
				1,
				"",
				types.NewUserGrantee("cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez"),
			),
			shouldErr: true,
		},
		{
			name: "invalid grantee returns error",
			msg: types.NewMsgRevokeAllowance(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				types.NewUserGrantee(""),
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgRevokeAllowance,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgRevokeAllowance_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgRevokeAllowance","value":{"grantee":{"type":"desmos/UserGrantee","value":{"user":"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez"}},"granter":"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgRevokeAllowance.GetSignBytes()))
}

func TestMsgRevokeAllowance_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgRevokeAllowance.Granter)
	require.Equal(t, []sdk.AccAddress{addr}, msgRevokeAllowance.GetSigners())
}
