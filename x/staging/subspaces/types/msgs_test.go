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
	)
	require.Equal(t, "subspaces", msg.Route())
}

func TestMsgCreateSubspace_Type(t *testing.T) {
	msg := types.NewMsgCreateSubspace(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"mooncake",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
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
			name: "invalid subspace returns error",
			msg: types.NewMsgCreateSubspace(
				"",
				"mooncake",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(types.ErrInvalidSubspace, "subspace id must be a valid sha-256 hash"),
		},
		{
			name: "invalid subspace owner address returns error",
			msg: types.NewMsgCreateSubspace(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"mooncake",
				"",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid owner address"),
		},
		{
			name: "invalid subspace name returns error",
			msg: types.NewMsgCreateSubspace(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(types.ErrInvalidSubspaceName, "subspace name cannot be empty or blank"),
		},
		{
			name: "valid message returns no error",
			msg: types.NewMsgCreateSubspace(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"mooncake",
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

func TestMsgCreateSubspace_GetSignBytes(t *testing.T) {
	msg := types.NewMsgCreateSubspace(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"mooncake",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	expected := `{"name":"mooncake","owner":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace_id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

func TestMsgCreateSubspace_GetSigners(t *testing.T) {
	msg := types.NewMsgCreateSubspace(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"mooncake",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	addr, _ := sdk.AccAddressFromBech32(msg.Owner)
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
			name: "invalid subspace returns error",
			msg: types.NewMsgAddAdmin(
				"",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(types.ErrInvalidSubspace, "subspace id must be a valid sha-256 hash"),
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
	expected := `{"new_admin":"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h","owner":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace_id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"}`
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
			name: "invalid subspace returns error",
			msg: types.NewMsgRemoveAdmin(
				"",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(types.ErrInvalidSubspace, "subspace id must be a valid sha-256 hash"),
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

func TestMsgAllowUserPosts_Route(t *testing.T) {
	msg := types.NewMsgEnableUserPosts(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, "subspaces", msg.Route())
}

func TestMsgAllowUserPosts_Type(t *testing.T) {
	msg := types.NewMsgEnableUserPosts(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, "enable_user_posts", msg.Type())
}

func TestMsgAllowUserPosts_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types.MsgEnableUserPosts
		error error
	}{
		{
			name: "invalid subspace returns error",
			msg: types.NewMsgEnableUserPosts(
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(types.ErrInvalidSubspace, "subspace id must be a valid sha-256 hash"),
		},
		{
			name: "invalid subspace admin address returns error",
			msg: types.NewMsgEnableUserPosts(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid admin address"),
		},
		{
			name: "invalid subspace user address returns error",
			msg: types.NewMsgEnableUserPosts(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid user address"),
		},
		{
			name: "valid message returns no error",
			msg: types.NewMsgEnableUserPosts(
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

func TestMsgAllowUserPosts_GetSignBytes(t *testing.T) {
	msg := types.NewMsgEnableUserPosts(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	expected := `{"admin":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace_id":"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h","user":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

func TestMsgAllowUserPosts_GetSigners(t *testing.T) {
	msg := types.NewMsgEnableUserPosts(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	addr, _ := sdk.AccAddressFromBech32(msg.Admin)
	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())
}

func TestMsgBlockUserPosts_Route(t *testing.T) {
	msg := types.NewMsgDisableUserPosts(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, "subspaces", msg.Route())
}

func TestMsgBlockUserPosts_Type(t *testing.T) {
	msg := types.NewMsgDisableUserPosts(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, "disable_user_posts", msg.Type())
}

func TestMsgBlockUserPosts_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types.MsgDisableUserPosts
		error error
	}{
		{
			name: "invalid subspace returns error",
			msg: types.NewMsgDisableUserPosts(
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(types.ErrInvalidSubspace, "subspace id must be a valid sha-256 hash"),
		},
		{
			name: "invalid subspace admin address returns error",
			msg: types.NewMsgDisableUserPosts(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid admin address"),
		},
		{
			name: "invalid subspace user address returns error",
			msg: types.NewMsgDisableUserPosts(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid user address"),
		},
		{
			name: "valid message returns no error",
			msg: types.NewMsgDisableUserPosts(
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

func TestMsgBlockUserPosts_GetSignBytes(t *testing.T) {
	msg := types.NewMsgDisableUserPosts(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	expected := `{"admin":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace_id":"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h","user":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

func TestMsgBlockUserPosts_GetSigners(t *testing.T) {
	msg := types.NewMsgDisableUserPosts(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	addr, _ := sdk.AccAddressFromBech32(msg.Admin)
	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())
}

func TestMsgTransferOwnership_Route(t *testing.T) {
	msg := types.NewMsgTransferOwnership(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, "subspaces", msg.Route())
}

func TestMsgTransferOwnership_Type(t *testing.T) {
	msg := types.NewMsgTransferOwnership(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, "transfer_ownership", msg.Type())
}

func TestMsgTransferOwnership_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types.MsgTransferOwnership
		error error
	}{
		{
			name: "invalid subspace returns error",
			msg: types.NewMsgTransferOwnership(
				"",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(types.ErrInvalidSubspace, "subspace id must be a valid sha-256 hash"),
		},
		{
			name: "invalid subspace owner address returns error",
			msg: types.NewMsgTransferOwnership(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid owner address"),
		},
		{
			name: "invalid subspace new owner address returns error",
			msg: types.NewMsgTransferOwnership(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid new owner address"),
		},
		{
			name: "equal subspace owner and new owner addresses returns error",
			msg: types.NewMsgTransferOwnership(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the new owner is equal to the owner"),
		},
		{
			name: "valid message returns no error",
			msg: types.NewMsgTransferOwnership(
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

func TestMsgTransferOwnership_GetSignBytes(t *testing.T) {
	msg := types.NewMsgTransferOwnership(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	expected := `{"new_owner":"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h","owner":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace_id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

func TestMsgTransferOwnership_GetSigners(t *testing.T) {
	msg := types.NewMsgTransferOwnership(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	addr, _ := sdk.AccAddressFromBech32(msg.Owner)
	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())
}
