package types_test

import (
	"testing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
	"github.com/stretchr/testify/require"
)

var msgGrantUserAllowance = types.NewMsgGrantUserAllowance(
	1,
	"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
	"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
	&feegrant.BasicAllowance{},
)

func TestMsgGrantUserAllowance_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgGrantUserAllowance.Route())
}

func TestMsgGrantUserAllowance_Type(t *testing.T) {
	require.Equal(t, types.ActionGrantUserAllowance, msgGrantUserAllowance.Type())
}

func TestMsgGrantUserAllowance_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgGrantUserAllowance
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgGrantUserAllowance(
				0,
				"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
				"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
				&feegrant.BasicAllowance{},
			),
			shouldErr: true,
		},
		{
			name: "invalid granter returns error",
			msg: types.NewMsgGrantUserAllowance(
				1,
				"",
				"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
				&feegrant.BasicAllowance{},
			),
			shouldErr: true,
		},
		{
			name: "invalid grantee returns error",
			msg: types.NewMsgGrantUserAllowance(
				1,
				"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
				"",
				&feegrant.BasicAllowance{},
			),
			shouldErr: true,
		},
		{
			name: "granter self-grant returns error",
			msg: types.NewMsgGrantUserAllowance(
				1,
				"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
				"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
				&feegrant.BasicAllowance{},
			),
			shouldErr: true,
		},
		{
			name: "invalid allowance returns no error",
			msg: &types.MsgGrantUserAllowance{
				SubspaceID: 1,
				Granter:    "cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
				Grantee:    "cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgGrantUserAllowance,
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

func TestMsgGrantUserAllowance_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgGrantUserAllowance","value":{"allowance":{"spend_limit":[]},"grantee":"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez","granter":"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgGrantUserAllowance.GetSignBytes()))
}

func TestMsgGrantUserAllowance_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgGrantUserAllowance.Granter)
	require.Equal(t, []sdk.AccAddress{addr}, msgGrantUserAllowance.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgRevokeUserAllowance = types.NewMsgRevokeUserAllowance(
	1,
	"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
	"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
)

func TestMsgRevokeUserAllowance_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgRevokeUserAllowance.Route())
}

func TestMsgRevokeUserAllowance_Type(t *testing.T) {
	require.Equal(t, types.ActionRevokeUserAllowance, msgRevokeUserAllowance.Type())
}

func TestMsgRevokeUserAllowance_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgRevokeUserAllowance
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgRevokeUserAllowance(
				0,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
			),
			shouldErr: true,
		},
		{
			name: "invalid granter returns error",
			msg: types.NewMsgRevokeUserAllowance(
				1,
				"",
				"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
			),
			shouldErr: true,
		},
		{
			name: "invalid grantee returns error",
			msg: types.NewMsgRevokeUserAllowance(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgRevokeUserAllowance,
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

func TestMsgRevokeUserAllowance_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgRevokeUserAllowance","value":{"grantee":"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez","granter":"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgRevokeUserAllowance.GetSignBytes()))
}

func TestMsgRevokeUserAllowance_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgRevokeUserAllowance.Granter)
	require.Equal(t, []sdk.AccAddress{addr}, msgRevokeUserAllowance.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgGrantGroupAllowance = types.NewMsgGrantGroupAllowance(
	1,
	"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
	1,
	&feegrant.BasicAllowance{},
)

func TestMsgGrantGroupAllowance_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgGrantGroupAllowance.Route())
}

func TestMsgGrantGroupAllowance_Type(t *testing.T) {
	require.Equal(t, types.ActionGrantGroupAllowance, msgGrantGroupAllowance.Type())
}

func TestMsgGrantGroupAllowance_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgGrantGroupAllowance
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgGrantGroupAllowance(
				0,
				"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
				1,
				&feegrant.BasicAllowance{},
			),
			shouldErr: true,
		},
		{
			name: "invalid granter returns error",
			msg: types.NewMsgGrantGroupAllowance(
				1,
				"",
				1,
				&feegrant.BasicAllowance{},
			),
			shouldErr: true,
		},
		{
			name: "invalid group id returns error",
			msg: types.NewMsgGrantGroupAllowance(
				1,
				"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
				0,
				&feegrant.BasicAllowance{},
			),
			shouldErr: true,
		},
		{
			name: "invalid allowance returns no error",
			msg: &types.MsgGrantGroupAllowance{
				SubspaceID: 1,
				Granter:    "cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
				GroupID:    1,
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgGrantGroupAllowance,
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

func TestMsgGrantGroupAllowance_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgGrantGroupAllowance","value":{"allowance":{"spend_limit":[]},"granter":"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0","group_id":1,"subspace_id":"1"}}`
	require.Equal(t, expected, string(msgGrantGroupAllowance.GetSignBytes()))
}

func TestMsgGrantGroupAllowance_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgGrantGroupAllowance.Granter)
	require.Equal(t, []sdk.AccAddress{addr}, msgGrantGroupAllowance.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgRevokeGroupAllowance = types.NewMsgRevokeGroupAllowance(
	1,
	"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
	1,
)

func TestMsgRevokeGroupAllowance_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgRevokeGroupAllowance.Route())
}

func TestMsgRevokeGroupAllowance_Type(t *testing.T) {
	require.Equal(t, types.ActionRevokeGroupAllowance, msgRevokeGroupAllowance.Type())
}

func TestMsgRevokeGroupAllowance_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgRevokeGroupAllowance
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgRevokeGroupAllowance(
				0,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				1,
			),
			shouldErr: true,
		},
		{
			name: "invalid granter returns error",
			msg: types.NewMsgRevokeGroupAllowance(
				1,
				"",
				1,
			),
			shouldErr: true,
		},
		{
			name: "invalid group id returns error",
			msg: types.NewMsgRevokeGroupAllowance(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				0,
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgRevokeGroupAllowance,
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

func TestMsgRevokeGroupAllowance_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgRevokeGroupAllowance","value":{"granter":"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0","group_id":1,"subspace_id":"1"}}`
	require.Equal(t, expected, string(msgRevokeGroupAllowance.GetSignBytes()))
}

func TestMsgRevokeGroupAllowance_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgRevokeGroupAllowance.Granter)
	require.Equal(t, []sdk.AccAddress{addr}, msgRevokeGroupAllowance.GetSigners())
}
