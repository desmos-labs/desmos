package types_test

//
//import (
//	"testing"
//
//	types2 "github.com/desmos-labs/desmos/v2/x/subspaces/types"
//
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	"github.com/stretchr/testify/require"
//)
//
//func TestMsgCreateSubspace_Route(t *testing.T) {
//	msg := types2.NewMsgCreateSubspace(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"mooncake",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//		types2.SubspaceTypeOpen,
//	)
//	require.Equal(t, "subspaces", msg.Route())
//}
//
//func TestMsgCreateSubspace_Type(t *testing.T) {
//	msg := types2.NewMsgCreateSubspace(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"mooncake",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//		types2.SubspaceTypeOpen,
//	)
//	require.Equal(t, "create_subspace", msg.Type())
//}
//
//func TestMsgCreateSubspace_ValidateBasic(t *testing.T) {
//	tests := []struct {
//		name   string
//		msg    *types2.MsgCreateSubspace
//		expErr bool
//	}{
//		{
//			name: "invalid subspace id returns error",
//			msg: types2.NewMsgCreateSubspace(
//				"",
//				"mooncake",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//				types2.SubspaceTypeOpen,
//			),
//			expErr: true,
//		},
//		{
//			name: "invalid subspace creator address returns error",
//			msg: types2.NewMsgCreateSubspace(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"mooncake",
//				"",
//				types2.SubspaceTypeOpen,
//			),
//			expErr: true,
//		},
//		{
//			name: "invalid subspace name returns error",
//			msg: types2.NewMsgCreateSubspace(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//				types2.SubspaceTypeOpen,
//			),
//			expErr: true,
//		},
//		{
//			name: "valid message returns no error",
//			msg: types2.NewMsgCreateSubspace(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"mooncake",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//				types2.SubspaceTypeOpen,
//			),
//		},
//	}
//
//	for _, test := range tests {
//		test := test
//		t.Run(test.name, func(t *testing.T) {
//			err := test.msg.ValidateBasic()
//			if test.expErr {
//				require.Error(t, err)
//			} else {
//				require.NoError(t, err)
//			}
//		})
//	}
//}
//
//func TestMsgCreateSubspace_GetSignBytes(t *testing.T) {
//	msg := types2.NewMsgCreateSubspace(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"mooncake",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//		types2.SubspaceTypeOpen,
//	)
//	expected := `{"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af","name":"mooncake","subspace_type":1}`
//	require.Equal(t, expected, string(msg.GetSignBytes()))
//}
//
//func TestMsgCreateSubspace_GetSigners(t *testing.T) {
//	msg := types2.NewMsgCreateSubspace(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"mooncake",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//		types2.SubspaceTypeOpen,
//	)
//	addr, _ := sdk.AccAddressFromBech32(msg.Creator)
//	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())
//}
//
//func TestMsgEditSubspace_Route(t *testing.T) {
//	msg := types2.NewMsgEditSubspace(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"star",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//		types2.SubspaceTypeOpen,
//	)
//	require.Equal(t, "subspaces", msg.Route())
//}
//
//func TestMsgEditSubspace_Type(t *testing.T) {
//	msg := types2.NewMsgEditSubspace(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"star",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//		types2.SubspaceTypeOpen,
//	)
//	require.Equal(t, "edit_subspace", msg.Type())
//}
//
//func TestMsgEditSubspace_ValidateBasic(t *testing.T) {
//	tests := []struct {
//		name   string
//		msg    *types2.MsgEditSubspace
//		expErr bool
//	}{
//		{
//			name: "invalid subspace id returns error",
//			msg: types2.NewMsgEditSubspace(
//				"",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//				"star",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//				types2.SubspaceTypeOpen,
//			),
//			expErr: true,
//		},
//		{
//			name: "invalid subspace owner address returns error",
//			msg: types2.NewMsgEditSubspace(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//				"star",
//				"",
//				types2.SubspaceTypeOpen,
//			),
//			expErr: true,
//		},
//		{
//			name: "equal subspace owner and new owner addresses returns error",
//			msg: types2.NewMsgEditSubspace(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//				"star",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//				types2.SubspaceTypeOpen,
//			),
//			expErr: true,
//		},
//		{
//			name: "valid message returns no error",
//			msg: types2.NewMsgEditSubspace(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//				"star",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//				types2.SubspaceTypeOpen,
//			),
//		},
//	}
//
//	for _, test := range tests {
//		test := test
//		t.Run(test.name, func(t *testing.T) {
//			err := test.msg.ValidateBasic()
//			if test.expErr {
//				require.Error(t, err)
//			} else {
//				require.NoError(t, err)
//			}
//		})
//	}
//}
//
//func TestMsgEditSubspace_GetSignBytes(t *testing.T) {
//	msg := types2.NewMsgEditSubspace(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"star",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//		types2.SubspaceTypeOpen,
//	)
//	expected := `{"editor":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af","name":"star","owner":"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h","subspace_type":1}`
//	require.Equal(t, expected, string(msg.GetSignBytes()))
//}
//
//func TestMsgEditSubspace_GetSigners(t *testing.T) {
//	msg := types2.NewMsgEditSubspace(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"star",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//		types2.SubspaceTypeOpen,
//	)
//	addr, _ := sdk.AccAddressFromBech32(msg.Editor)
//	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())
//}
//
//func TestMsgAddAdmin_Route(t *testing.T) {
//	msg := types2.NewMsgAddAdmin(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	require.Equal(t, "subspaces", msg.Route())
//}
//
//func TestMsgAddAdmin_Type(t *testing.T) {
//	msg := types2.NewMsgAddAdmin(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	require.Equal(t, "add_admin", msg.Type())
//}
//
//func TestMsgAddAdmin_ValidateBasic(t *testing.T) {
//	tests := []struct {
//		name   string
//		msg    *types2.MsgAddAdmin
//		expErr bool
//	}{
//		{
//			name: "invalid subspace id returns error",
//			msg: types2.NewMsgAddAdmin(
//				"",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//			),
//			expErr: true,
//		},
//		{
//			name: "equals owner and admin addresses returns error",
//			msg: types2.NewMsgAddAdmin(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//			),
//			expErr: true,
//		},
//		{
//			name: "invalid subspace owner address returns error",
//			msg: types2.NewMsgAddAdmin(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//				"",
//			),
//			expErr: true,
//		},
//		{
//			name: "invalid subspace new admin address returns error",
//			msg: types2.NewMsgAddAdmin(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//			),
//			expErr: true,
//		},
//		{
//			name: "valid message returns no error",
//			msg: types2.NewMsgAddAdmin(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//			),
//			expErr: false,
//		},
//	}
//
//	for _, test := range tests {
//		test := test
//		t.Run(test.name, func(t *testing.T) {
//			err := test.msg.ValidateBasic()
//			if test.expErr {
//				require.Error(t, err)
//			} else {
//				require.NoError(t, err)
//			}
//		})
//	}
//}
//
//func TestMsgAddAdmin_GetSignBytes(t *testing.T) {
//	msg := types2.NewMsgAddAdmin(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	expected := `{"admin":"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h","owner":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace_id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"}`
//	require.Equal(t, expected, string(msg.GetSignBytes()))
//}
//
//func TestMsgAddAdmin_GetSigners(t *testing.T) {
//	msg := types2.NewMsgAddAdmin(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	addr, _ := sdk.AccAddressFromBech32(msg.Owner)
//	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())
//}
//
//func TestMsgRemoveAdmin_Route(t *testing.T) {
//	msg := types2.NewMsgRemoveAdmin(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	require.Equal(t, "subspaces", msg.Route())
//}
//
//func TestMsgRemoveAdmin_Type(t *testing.T) {
//	msg := types2.NewMsgRemoveAdmin(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	require.Equal(t, "remove_admin", msg.Type())
//}
//
//func TestMsgRemoveAdmin_ValidateBasic(t *testing.T) {
//	tests := []struct {
//		name   string
//		msg    *types2.MsgRemoveAdmin
//		expErr bool
//	}{
//		{
//			name: "invalid subspace id returns error",
//			msg: types2.NewMsgRemoveAdmin(
//				"",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//			),
//			expErr: true,
//		},
//		{
//			name: "equals owner and admin addresses returns error",
//			msg: types2.NewMsgRemoveAdmin(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//			),
//			expErr: true,
//		},
//		{
//			name: "invalid subspace owner address returns error",
//			msg: types2.NewMsgRemoveAdmin(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//				"",
//			),
//			expErr: true,
//		},
//		{
//			name: "invalid subspace admin address returns error",
//			msg: types2.NewMsgRemoveAdmin(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//			),
//			expErr: true,
//		},
//		{
//			name: "valid message returns no error",
//			msg: types2.NewMsgRemoveAdmin(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//			),
//			expErr: false,
//		},
//	}
//
//	for _, test := range tests {
//		test := test
//		t.Run(test.name, func(t *testing.T) {
//			err := test.msg.ValidateBasic()
//			if test.expErr {
//				require.Error(t, err)
//			} else {
//				require.NoError(t, err)
//			}
//		})
//	}
//}
//
//func TestMsgRemoveAdmin_GetSignBytes(t *testing.T) {
//	msg := types2.NewMsgRemoveAdmin(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	expected := `{"admin":"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h","owner":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace_id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"}`
//	require.Equal(t, expected, string(msg.GetSignBytes()))
//}
//
//func TestMsgRemoveAdmin_GetSigners(t *testing.T) {
//	msg := types2.NewMsgRemoveAdmin(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	addr, _ := sdk.AccAddressFromBech32(msg.Owner)
//	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())
//}
//
//func TestMsgRegisterUser_Route(t *testing.T) {
//	msg := types2.NewMsgRegisterUser(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	require.Equal(t, "subspaces", msg.Route())
//}
//
//func TestMsgRegisterUser_Type(t *testing.T) {
//	msg := types2.NewMsgRegisterUser(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	require.Equal(t, "register_user", msg.Type())
//}
//
//func TestMsgRegisterUser_ValidateBasic(t *testing.T) {
//	tests := []struct {
//		name   string
//		msg    *types2.MsgRegisterUser
//		expErr bool
//	}{
//		{
//			name: "invalid subspace id returns error",
//			msg: types2.NewMsgRegisterUser(
//				"",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//			),
//			expErr: true,
//		},
//		{
//			name: "invalid subspace admin address returns error",
//			msg: types2.NewMsgRegisterUser(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//				"",
//			),
//			expErr: true,
//		},
//		{
//			name: "invalid subspace user address returns error",
//			msg: types2.NewMsgRegisterUser(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//			),
//			expErr: true,
//		},
//		{
//			name: "valid message returns no error",
//			msg: types2.NewMsgRegisterUser(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//			),
//			expErr: false,
//		},
//	}
//
//	for _, test := range tests {
//		test := test
//		t.Run(test.name, func(t *testing.T) {
//			err := test.msg.ValidateBasic()
//			if test.expErr {
//				require.Error(t, err)
//			} else {
//				require.NoError(t, err)
//			}
//		})
//	}
//}
//
//func TestMsgRegisterUser_GetSignBytes(t *testing.T) {
//	msg := types2.NewMsgRegisterUser(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	expected := `{"admin":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace_id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af","user":"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h"}`
//	require.Equal(t, expected, string(msg.GetSignBytes()))
//}
//
//func TestMsgRegisterUser_GetSigners(t *testing.T) {
//	msg := types2.NewMsgRegisterUser(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	addr, _ := sdk.AccAddressFromBech32(msg.Admin)
//	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())
//}
//
//func TestMsgUnregisterUser_Route(t *testing.T) {
//	msg := types2.NewMsgUnregisterUser(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	require.Equal(t, "subspaces", msg.Route())
//}
//
//func TestMsgUnregisterUser_Type(t *testing.T) {
//	msg := types2.NewMsgUnregisterUser(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	require.Equal(t, "unregister_user", msg.Type())
//}
//
//func TestMsgUnregisterUser_ValidateBasic(t *testing.T) {
//	tests := []struct {
//		name   string
//		msg    *types2.MsgUnregisterUser
//		expErr bool
//	}{
//		{
//			name: "invalid subspace id returns error",
//			msg: types2.NewMsgUnregisterUser(
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//				"",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//			),
//			expErr: true,
//		},
//		{
//			name: "invalid subspace admin address returns error",
//			msg: types2.NewMsgUnregisterUser(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//				"",
//			),
//			expErr: true,
//		},
//		{
//			name: "invalid subspace user address returns error",
//			msg: types2.NewMsgUnregisterUser(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//			),
//			expErr: true,
//		},
//		{
//			name: "valid message returns no error",
//			msg: types2.NewMsgUnregisterUser(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//			),
//			expErr: false,
//		},
//	}
//
//	for _, test := range tests {
//		test := test
//		t.Run(test.name, func(t *testing.T) {
//			err := test.msg.ValidateBasic()
//			if test.expErr {
//				require.Error(t, err)
//			} else {
//				require.NoError(t, err)
//			}
//		})
//	}
//}
//
//func TestMsgUnregisterUser_GetSignBytes(t *testing.T) {
//	msg := types2.NewMsgUnregisterUser(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	expected := `{"admin":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace_id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af","user":"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h"}`
//	require.Equal(t, expected, string(msg.GetSignBytes()))
//}
//
//func TestMsgUnregisterUser_GetSigners(t *testing.T) {
//	msg := types2.NewMsgUnregisterUser(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	addr, _ := sdk.AccAddressFromBech32(msg.Admin)
//	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())
//}
//
//func TestMsgBanUser_Route(t *testing.T) {
//	msg := types2.NewMsgBanUser(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	require.Equal(t, "subspaces", msg.Route())
//}
//
//func TestMsgBanUser_Type(t *testing.T) {
//	msg := types2.NewMsgBanUser(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	require.Equal(t, "block_user", msg.Type())
//}
//
//func TestMsgBanUser_ValidateBasic(t *testing.T) {
//	tests := []struct {
//		name   string
//		msg    *types2.MsgBanUser
//		expErr bool
//	}{
//		{
//			name: "invalid subspace id returns error",
//			msg: types2.NewMsgBanUser(
//				"",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//			),
//			expErr: true,
//		},
//		{
//			name: "invalid subspace admin address returns error",
//			msg: types2.NewMsgBanUser(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//				"",
//			),
//			expErr: true,
//		},
//		{
//			name: "invalid subspace user address returns error",
//			msg: types2.NewMsgBanUser(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//			),
//			expErr: true,
//		},
//		{
//			name: "valid message returns no error",
//			msg: types2.NewMsgBanUser(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//			),
//			expErr: false,
//		},
//	}
//
//	for _, test := range tests {
//		test := test
//		t.Run(test.name, func(t *testing.T) {
//			err := test.msg.ValidateBasic()
//			if test.expErr {
//				require.Error(t, err)
//			} else {
//				require.NoError(t, err)
//			}
//		})
//	}
//}
//
//func TestMsgBanUser_GetSignBytes(t *testing.T) {
//	msg := types2.NewMsgBanUser(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	expected := `{"admin":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace_id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af","user":"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h"}`
//	require.Equal(t, expected, string(msg.GetSignBytes()))
//}
//
//func TestMsgBanUser_GetSigners(t *testing.T) {
//	msg := types2.NewMsgBanUser(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	addr, _ := sdk.AccAddressFromBech32(msg.Admin)
//	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())
//}
//
//func TestTestMsgUnbanUser_Route(t *testing.T) {
//	msg := types2.NewMsgUnbanUser(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	require.Equal(t, "subspaces", msg.Route())
//}
//
//func TestTestMsgUnbanUser_Type(t *testing.T) {
//	msg := types2.NewMsgUnbanUser(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	require.Equal(t, "unblock_user", msg.Type())
//}
//
//func TestTestMsgUnbanUser_ValidateBasic(t *testing.T) {
//	tests := []struct {
//		name   string
//		msg    *types2.MsgUnbanUser
//		expErr bool
//	}{
//		{
//			name: "invalid subspace id returns error",
//			msg: types2.NewMsgUnbanUser(
//				"",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//			),
//			expErr: true,
//		},
//		{
//			name: "invalid subspace admin address returns error",
//			msg: types2.NewMsgUnbanUser(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//				"",
//			),
//			expErr: true,
//		},
//		{
//			name: "invalid subspace user address returns error",
//			msg: types2.NewMsgUnbanUser(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//			),
//			expErr: true,
//		},
//		{
//			name: "valid message returns no error",
//			msg: types2.NewMsgUnbanUser(
//				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//			),
//			expErr: false,
//		},
//	}
//
//	for _, test := range tests {
//		test := test
//		t.Run(test.name, func(t *testing.T) {
//			err := test.msg.ValidateBasic()
//			if test.expErr {
//				require.Error(t, err)
//			} else {
//				require.NoError(t, err)
//			}
//		})
//	}
//}
//
//func TestTestMsgUnbanUser_GetSignBytes(t *testing.T) {
//	msg := types2.NewMsgUnbanUser(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	expected := `{"admin":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace_id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af","user":"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h"}`
//	require.Equal(t, expected, string(msg.GetSignBytes()))
//}
//
//func TestTestMsgUnbanUser_GetSigners(t *testing.T) {
//	msg := types2.NewMsgUnbanUser(
//		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
//		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
//		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
//	)
//	addr, _ := sdk.AccAddressFromBech32(msg.Admin)
//	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())
//}
