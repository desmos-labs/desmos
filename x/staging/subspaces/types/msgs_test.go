package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
)

func TestMsgCreateSubspace_Route(t *testing.T) {
	msg := types.NewMsgCreateSubspace(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"mooncake",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		types.SubspaceTypeOpen,
	)
	require.Equal(t, "subspaces", msg.Route())
}

func TestMsgCreateSubspace_Type(t *testing.T) {
	msg := types.NewMsgCreateSubspace(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"mooncake",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		types.SubspaceTypeOpen,
	)
	require.Equal(t, "create_subspace", msg.Type())
}

func TestMsgCreateSubspace_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    *types.MsgCreateSubspace
		expErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgCreateSubspace(
				"",
				"mooncake",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				types.SubspaceTypeOpen,
			),
			expErr: true,
		},
		{
			name: "invalid subspace creator address returns error",
			msg: types.NewMsgCreateSubspace(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"mooncake",
				"",
				types.SubspaceTypeOpen,
			),
			expErr: true,
		},
		{
			name: "invalid subspace name returns error",
			msg: types.NewMsgCreateSubspace(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				types.SubspaceTypeOpen,
			),
			expErr: true,
		},
		{
			name: "valid message returns no error",
			msg: types.NewMsgCreateSubspace(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"mooncake",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				types.SubspaceTypeOpen,
			),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.msg.ValidateBasic()
			if test.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgCreateSubspace_GetSignBytes(t *testing.T) {
	msg := types.NewMsgCreateSubspace(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"mooncake",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		types.SubspaceTypeOpen,
	)
	expected := `{"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","name":"mooncake","subspace_id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af","type":1}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

func TestMsgCreateSubspace_GetSigners(t *testing.T) {
	msg := types.NewMsgCreateSubspace(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"mooncake",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		types.SubspaceTypeOpen,
	)
	addr, _ := sdk.AccAddressFromBech32(msg.Creator)
	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())
}

func TestMsgEditSubspace_Route(t *testing.T) {
	msg := types.NewMsgEditSubspace(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"star",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		types.SubspaceTypeOpen,
	)
	require.Equal(t, "subspaces", msg.Route())
}

func TestMsgEditSubspace_Type(t *testing.T) {
	msg := types.NewMsgEditSubspace(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"star",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		types.SubspaceTypeOpen,
	)
	require.Equal(t, "edit_subspace", msg.Type())
}

func TestMsgEditSubspace_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    *types.MsgEditSubspace
		expErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgEditSubspace(
				"",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"star",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				types.SubspaceTypeOpen,
			),
			expErr: true,
		},
		{
			name: "invalid subspace owner address returns error",
			msg: types.NewMsgEditSubspace(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"star",
				"",
				types.SubspaceTypeOpen,
			),
			expErr: true,
		},
		{
			name: "equal subspace owner and new owner addresses returns error",
			msg: types.NewMsgEditSubspace(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"star",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				types.SubspaceTypeOpen,
			),
			expErr: true,
		},
		{
			name: "valid message returns no error",
			msg: types.NewMsgEditSubspace(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"star",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				types.SubspaceTypeOpen,
			),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.msg.ValidateBasic()
			if test.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgEditSubspace_GetSignBytes(t *testing.T) {
	msg := types.NewMsgEditSubspace(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"star",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		types.SubspaceTypeOpen,
	)
	expected := `{"editor":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","name":"star","owner":"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h","subspace_id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af","type":1}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

func TestMsgEditSubspace_GetSigners(t *testing.T) {
	msg := types.NewMsgEditSubspace(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"star",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		types.SubspaceTypeOpen,
	)
	addr, _ := sdk.AccAddressFromBech32(msg.Editor)
	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())
}

func TestMsgAddAdmin_Route(t *testing.T) {
	msg := types.NewMsgAddAdmin(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, "subspaces", msg.Route())
}

func TestMsgAddAdmin_Type(t *testing.T) {
	msg := types.NewMsgAddAdmin(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, "add_admin", msg.Type())
}

func TestMsgAddAdmin_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    *types.MsgAddAdmin
		expErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgAddAdmin(
				"",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "equals owner and admin addresses returns error",
			msg: types.NewMsgAddAdmin(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			),
			expErr: true,
		},
		{
			name: "invalid subspace owner address returns error",
			msg: types.NewMsgAddAdmin(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"",
			),
			expErr: true,
		},
		{
			name: "invalid subspace new admin address returns error",
			msg: types.NewMsgAddAdmin(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "valid message returns no error",
			msg: types.NewMsgAddAdmin(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.msg.ValidateBasic()
			if test.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgAddAdmin_GetSignBytes(t *testing.T) {
	msg := types.NewMsgAddAdmin(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	expected := `{"admin":"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h","owner":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace_id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

func TestMsgAddAdmin_GetSigners(t *testing.T) {
	msg := types.NewMsgAddAdmin(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	addr, _ := sdk.AccAddressFromBech32(msg.Owner)
	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())
}

func TestMsgRemoveAdmin_Route(t *testing.T) {
	msg := types.NewMsgRemoveAdmin(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, "subspaces", msg.Route())
}

func TestMsgRemoveAdmin_Type(t *testing.T) {
	msg := types.NewMsgRemoveAdmin(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, "remove_admin", msg.Type())
}

func TestMsgRemoveAdmin_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    *types.MsgRemoveAdmin
		expErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgRemoveAdmin(
				"",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "equals owner and admin addresses returns error",
			msg: types.NewMsgRemoveAdmin(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			),
			expErr: true,
		},
		{
			name: "invalid subspace owner address returns error",
			msg: types.NewMsgRemoveAdmin(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"",
			),
			expErr: true,
		},
		{
			name: "invalid subspace admin address returns error",
			msg: types.NewMsgRemoveAdmin(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "valid message returns no error",
			msg: types.NewMsgRemoveAdmin(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.msg.ValidateBasic()
			if test.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgRemoveAdmin_GetSignBytes(t *testing.T) {
	msg := types.NewMsgRemoveAdmin(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	expected := `{"admin":"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h","owner":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace_id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

func TestMsgRemoveAdmin_GetSigners(t *testing.T) {
	msg := types.NewMsgRemoveAdmin(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	addr, _ := sdk.AccAddressFromBech32(msg.Owner)
	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())
}

func TestMsgRegisterUser_Route(t *testing.T) {
	msg := types.NewMsgRegisterUser(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, "subspaces", msg.Route())
}

func TestMsgRegisterUser_Type(t *testing.T) {
	msg := types.NewMsgRegisterUser(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, "register_user", msg.Type())
}

func TestMsgRegisterUser_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    *types.MsgRegisterUser
		expErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgRegisterUser(
				"",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "invalid subspace admin address returns error",
			msg: types.NewMsgRegisterUser(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"",
			),
			expErr: true,
		},
		{
			name: "invalid subspace user address returns error",
			msg: types.NewMsgRegisterUser(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "valid message returns no error",
			msg: types.NewMsgRegisterUser(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.msg.ValidateBasic()
			if test.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgRegisterUser_GetSignBytes(t *testing.T) {
	msg := types.NewMsgRegisterUser(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	expected := `{"admin":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace_id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af","user":"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h"}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

func TestMsgRegisterUser_GetSigners(t *testing.T) {
	msg := types.NewMsgRegisterUser(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	addr, _ := sdk.AccAddressFromBech32(msg.Admin)
	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())
}

func TestMsgUnregisterUser_Route(t *testing.T) {
	msg := types.NewMsgUnregisterUser(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, "subspaces", msg.Route())
}

func TestMsgUnregisterUser_Type(t *testing.T) {
	msg := types.NewMsgUnregisterUser(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, "unregister_user", msg.Type())
}

func TestMsgUnregisterUser_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    *types.MsgUnregisterUser
		expErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgUnregisterUser(
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "invalid subspace admin address returns error",
			msg: types.NewMsgUnregisterUser(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"",
			),
			expErr: true,
		},
		{
			name: "invalid subspace user address returns error",
			msg: types.NewMsgUnregisterUser(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "valid message returns no error",
			msg: types.NewMsgUnregisterUser(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.msg.ValidateBasic()
			if test.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgUnregisterUser_GetSignBytes(t *testing.T) {
	msg := types.NewMsgUnregisterUser(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	expected := `{"admin":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace_id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af","user":"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h"}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

func TestMsgUnregisterUser_GetSigners(t *testing.T) {
	msg := types.NewMsgUnregisterUser(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	addr, _ := sdk.AccAddressFromBech32(msg.Admin)
	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())
}

func TestMsgBanUser_Route(t *testing.T) {
	msg := types.NewMsgBanUser(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, "subspaces", msg.Route())
}

func TestMsgBanUser_Type(t *testing.T) {
	msg := types.NewMsgBanUser(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, "block_user", msg.Type())
}

func TestMsgBanUser_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    *types.MsgBanUser
		expErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgBanUser(
				"",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "invalid subspace admin address returns error",
			msg: types.NewMsgBanUser(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"",
			),
			expErr: true,
		},
		{
			name: "invalid subspace user address returns error",
			msg: types.NewMsgBanUser(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "valid message returns no error",
			msg: types.NewMsgBanUser(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.msg.ValidateBasic()
			if test.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgBanUser_GetSignBytes(t *testing.T) {
	msg := types.NewMsgBanUser(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	expected := `{"admin":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace_id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af","user":"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h"}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

func TestMsgBanUser_GetSigners(t *testing.T) {
	msg := types.NewMsgBanUser(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	addr, _ := sdk.AccAddressFromBech32(msg.Admin)
	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())
}

func TestTestMsgUnbanUser_Route(t *testing.T) {
	msg := types.NewMsgUnbanUser(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, "subspaces", msg.Route())
}

func TestTestMsgUnbanUser_Type(t *testing.T) {
	msg := types.NewMsgUnbanUser(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, "unblock_user", msg.Type())
}

func TestTestMsgUnbanUser_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    *types.MsgUnbanUser
		expErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgUnbanUser(
				"",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "invalid subspace admin address returns error",
			msg: types.NewMsgUnbanUser(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"",
			),
			expErr: true,
		},
		{
			name: "invalid subspace user address returns error",
			msg: types.NewMsgUnbanUser(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "valid message returns no error",
			msg: types.NewMsgUnbanUser(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.msg.ValidateBasic()
			if test.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestTestMsgUnbanUser_GetSignBytes(t *testing.T) {
	msg := types.NewMsgUnbanUser(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	expected := `{"admin":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace_id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af","user":"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h"}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

func TestTestMsgUnbanUser_GetSigners(t *testing.T) {
	msg := types.NewMsgUnbanUser(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	addr, _ := sdk.AccAddressFromBech32(msg.Admin)
	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())
}
