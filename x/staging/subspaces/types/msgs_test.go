package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMsgCreateSubspace_Route(t *testing.T) {
	msg := types.NewMsgCreateSubspace(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"mooncake",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		true,
	)
	require.Equal(t, "subspaces", msg.Route())
}

func TestMsgCreateSubspace_Type(t *testing.T) {
	msg := types.NewMsgCreateSubspace(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"mooncake",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		true,
	)
	require.Equal(t, "create_subspace", msg.Type())
}

func TestMsgCreateSubspace_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types.MsgCreateSubspace
		error error
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgCreateSubspace(
				"",
				"mooncake",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				true,
			),
			error: sdkerrors.Wrap(types.ErrInvalidSubspace, "subspace id must be a valid SHA-256 hash"),
		},
		{
			name: "invalid subspace creator address returns error",
			msg: types.NewMsgCreateSubspace(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"mooncake",
				"",
				true,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator address"),
		},
		{
			name: "invalid subspace name returns error",
			msg: types.NewMsgCreateSubspace(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				true,
			),
			error: sdkerrors.Wrap(types.ErrInvalidSubspaceName, "subspace name cannot be empty or blank"),
		},
		{
			name: "valid message returns no error",
			msg: types.NewMsgCreateSubspace(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"mooncake",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				true,
			),
			error: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			returnedError := test.msg.ValidateBasic()
			if test.error == nil {
				require.Nil(t, returnedError)
			} else {
				require.NotNil(t, returnedError)
				require.Equal(t, test.error.Error(), returnedError.Error())
			}
		})
	}
}

func TestMsgCreateSubspace_GetSignBytes(t *testing.T) {
	msg := types.NewMsgCreateSubspace(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"mooncake",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		true,
	)
	expected := `{"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","name":"mooncake","open":true,"subspace_id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

func TestMsgCreateSubspace_GetSigners(t *testing.T) {
	msg := types.NewMsgCreateSubspace(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"mooncake",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		true,
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
	)
	require.Equal(t, "subspaces", msg.Route())
}

func TestMsgEditSubspace_Type(t *testing.T) {
	msg := types.NewMsgEditSubspace(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"star",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, "edit_subspace", msg.Type())
}

func TestMsgEditSubspace_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types.MsgEditSubspace
		error error
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgEditSubspace(
				"",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"star",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(types.ErrInvalidSubspace, "subspace id must be a valid SHA-256 hash"),
		},
		{
			name: "invalid subspace owner address returns error",
			msg: types.NewMsgEditSubspace(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"star",
				"",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid editor address"),
		},
		{
			name: "equal subspace owner and new owner addresses returns error",
			msg: types.NewMsgEditSubspace(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"star",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the owner address is equal to the editor address"),
		},
		{
			name: "valid message returns no error",
			msg: types.NewMsgEditSubspace(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"star",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			returnedError := test.msg.ValidateBasic()
			if test.error == nil {
				require.Nil(t, returnedError)
			} else {
				require.NotNil(t, returnedError)
				require.Equal(t, test.error.Error(), returnedError.Error())
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
	)
	expected := `{"editor":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","name":"star","owner":"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h","subspace_id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

func TestMsgEditSubspace_GetSigners(t *testing.T) {
	msg := types.NewMsgEditSubspace(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"star",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
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
		name  string
		msg   *types.MsgAddAdmin
		error error
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgAddAdmin(
				"",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(types.ErrInvalidSubspace, "subspace id must be a valid SHA-256 hash"),
		},
		{
			name: "equals owner and admin addresses returns error",
			msg: types.NewMsgAddAdmin(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "owner address can't be equal to admin address"),
		},
		{
			name: "invalid subspace owner address returns error",
			msg: types.NewMsgAddAdmin(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid owner address"),
		},
		{
			name: "invalid subspace new admin address returns error",
			msg: types.NewMsgAddAdmin(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid new admin address"),
		},
		{
			name: "valid message returns no error",
			msg: types.NewMsgAddAdmin(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			returnedError := test.msg.ValidateBasic()
			if test.error == nil {
				require.Nil(t, returnedError)
			} else {
				require.NotNil(t, returnedError)
				require.Equal(t, test.error.Error(), returnedError.Error())
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
		name  string
		msg   *types.MsgRemoveAdmin
		error error
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgRemoveAdmin(
				"",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(types.ErrInvalidSubspace, "subspace id must be a valid SHA-256 hash"),
		},
		{
			name: "equals owner and admin addresses returns error",
			msg: types.NewMsgRemoveAdmin(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "owner address can't be equal to admin address"),
		},
		{
			name: "invalid subspace owner address returns error",
			msg: types.NewMsgRemoveAdmin(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid owner address"),
		},
		{
			name: "invalid subspace admin address returns error",
			msg: types.NewMsgRemoveAdmin(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid admin address"),
		},
		{
			name: "valid message returns no error",
			msg: types.NewMsgRemoveAdmin(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			returnedError := test.msg.ValidateBasic()
			if test.error == nil {
				require.Nil(t, returnedError)
			} else {
				require.NotNil(t, returnedError)
				require.Equal(t, test.error.Error(), returnedError.Error())
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
		name  string
		msg   *types.MsgRegisterUser
		error error
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgRegisterUser(
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(types.ErrInvalidSubspace, "subspace id must be a valid SHA-256 hash"),
		},
		{
			name: "invalid subspace admin address returns error",
			msg: types.NewMsgRegisterUser(
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid admin address"),
		},
		{
			name: "invalid subspace user address returns error",
			msg: types.NewMsgRegisterUser(
				"",
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid user address"),
		},
		{
			name: "valid message returns no error",
			msg: types.NewMsgRegisterUser(
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			returnedError := test.msg.ValidateBasic()
			if test.error == nil {
				require.Nil(t, returnedError)
			} else {
				require.NotNil(t, returnedError)
				require.Equal(t, test.error.Error(), returnedError.Error())
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
	expected := `{"admin":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace_id":"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h","user":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"}`
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
		name  string
		msg   *types.MsgUnregisterUser
		error error
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgUnregisterUser(
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(types.ErrInvalidSubspace, "subspace id must be a valid SHA-256 hash"),
		},
		{
			name: "invalid subspace admin address returns error",
			msg: types.NewMsgUnregisterUser(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid admin address"),
		},
		{
			name: "invalid subspace user address returns error",
			msg: types.NewMsgUnregisterUser(
				"",
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid user address"),
		},
		{
			name: "valid message returns no error",
			msg: types.NewMsgUnregisterUser(
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			returnedError := test.msg.ValidateBasic()
			if test.error == nil {
				require.Nil(t, returnedError)
			} else {
				require.NotNil(t, returnedError)
				require.Equal(t, test.error.Error(), returnedError.Error())
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
	expected := `{"admin":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace_id":"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h","user":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"}`
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
		name  string
		msg   *types.MsgBanUser
		error error
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgBanUser(
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(types.ErrInvalidSubspace, "subspace id must be a valid SHA-256 hash"),
		},
		{
			name: "invalid subspace admin address returns error",
			msg: types.NewMsgBanUser(
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid admin address"),
		},
		{
			name: "invalid subspace user address returns error",
			msg: types.NewMsgBanUser(
				"",
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid user address"),
		},
		{
			name: "valid message returns no error",
			msg: types.NewMsgBanUser(
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			returnedError := test.msg.ValidateBasic()
			if test.error == nil {
				require.Nil(t, returnedError)
			} else {
				require.NotNil(t, returnedError)
				require.Equal(t, test.error.Error(), returnedError.Error())
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
	expected := `{"admin":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace_id":"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h","user":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"}`
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
		name  string
		msg   *types.MsgUnbanUser
		error error
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgUnbanUser(
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(types.ErrInvalidSubspace, "subspace id must be a valid SHA-256 hash"),
		},
		{
			name: "invalid subspace admin address returns error",
			msg: types.NewMsgUnbanUser(
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid admin address"),
		},
		{
			name: "invalid subspace user address returns error",
			msg: types.NewMsgUnbanUser(
				"",
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid user address"),
		},
		{
			name: "valid message returns no error",
			msg: types.NewMsgUnbanUser(
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			returnedError := test.msg.ValidateBasic()
			if test.error == nil {
				require.Nil(t, returnedError)
			} else {
				require.NotNil(t, returnedError)
				require.Equal(t, test.error.Error(), returnedError.Error())
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
	expected := `{"admin":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace_id":"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h","user":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"}`
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
