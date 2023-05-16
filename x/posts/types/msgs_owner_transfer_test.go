package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v4/x/posts/types"
	"github.com/stretchr/testify/require"
)

var msgRequestPostOwnerTransfer = types.NewMsgRequestPostOwnerTransfer(
	1,
	1,
	"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
	"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
)

func TestMsgRequestPostOwnerTransfer_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgRequestPostOwnerTransfer.Route())
}

func TestMsgRequestPostOwnerTransfer_Type(t *testing.T) {
	require.Equal(t, types.ActionRequestPostOwnerTransfer, msgRequestPostOwnerTransfer.Type())
}

func TestMsgRequestPostOwnerTransfer_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgRequestPostOwnerTransfer
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgRequestPostOwnerTransfer(
				0,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			msg: types.NewMsgRequestPostOwnerTransfer(
				1,
				0,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "invalid receiver returns error",
			msg: types.NewMsgRequestPostOwnerTransfer(
				1,
				0,
				"",
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "invalid sender returns error",
			msg: types.NewMsgRequestPostOwnerTransfer(
				1,
				0,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgRequestPostOwnerTransfer,
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

func TestMsgRequestPostOwnerTransfer_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgRequestPostOwnerTransfer","value":{"post_id":"1","receiver":"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd","sender":"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgRequestPostOwnerTransfer.GetSignBytes()))
}

func TestMsgRequestPostOwnerTransfer_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgRequestPostOwnerTransfer.Sender)
	require.Equal(t, []sdk.AccAddress{addr}, msgRequestPostOwnerTransfer.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgCancelPostOwnerTransfer = types.NewMsgCancelPostOwnerTransfer(
	1,
	1,
	"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
)

func TestMsgCancelPostOwnerTransfer_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgCancelPostOwnerTransfer.Route())
}

func TestMsgCancelPostOwnerTransfer_Type(t *testing.T) {
	require.Equal(t, types.ActionCancelPostOwnerTransfer, msgCancelPostOwnerTransfer.Type())
}

func TestMsgCancelPostOwnerTransfer_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgCancelPostOwnerTransfer
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgCancelPostOwnerTransfer(
				0,
				1,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			msg: types.NewMsgCancelPostOwnerTransfer(
				1,
				0,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "invalid sender returns error",
			msg: types.NewMsgCancelPostOwnerTransfer(
				1,
				0,
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgCancelPostOwnerTransfer,
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

func TestMsgCancelPostOwnerTransfer_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgCancelPostOwnerTransfer","value":{"post_id":"1","sender":"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgCancelPostOwnerTransfer.GetSignBytes()))
}

func TestMsgCancelPostOwnerTransfer_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgCancelPostOwnerTransfer.Sender)
	require.Equal(t, []sdk.AccAddress{addr}, msgCancelPostOwnerTransfer.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgAcceptPostOwnerTransfer = types.NewMsgAcceptPostOwnerTransfer(
	1,
	1,
	"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
)

func TestMsgAcceptPostOwnerTransfer_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgAcceptPostOwnerTransfer.Route())
}

func TestMsgAcceptPostOwnerTransfer_Type(t *testing.T) {
	require.Equal(t, types.ActionAcceptPostOwnerTransfer, msgAcceptPostOwnerTransfer.Type())
}

func TestMsgAcceptPostOwnerTransfer_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgAcceptPostOwnerTransfer
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgAcceptPostOwnerTransfer(
				0,
				1,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			msg: types.NewMsgAcceptPostOwnerTransfer(
				1,
				0,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "invalid receiver returns error",
			msg: types.NewMsgAcceptPostOwnerTransfer(
				1,
				0,
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgAcceptPostOwnerTransfer,
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

func TestMsgAcceptPostOwnerTransfer_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgAcceptPostOwnerTransfer","value":{"post_id":"1","receiver":"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgAcceptPostOwnerTransfer.GetSignBytes()))
}

func TestMsgAcceptPostOwnerTransfer_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgAcceptPostOwnerTransfer.Receiver)
	require.Equal(t, []sdk.AccAddress{addr}, msgAcceptPostOwnerTransfer.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgRefusePostOwnerTransfer = types.NewMsgRefusePostOwnerTransfer(
	1,
	1,
	"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
)

func TestMsgRefusePostOwnerTransfer_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgRefusePostOwnerTransfer.Route())
}

func TestMsgRefusePostOwnerTransfer_Type(t *testing.T) {
	require.Equal(t, types.ActionRefusePostOwnerTransfer, msgRefusePostOwnerTransfer.Type())
}

func TestMsgRefusePostOwnerTransfer_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgRefusePostOwnerTransfer
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgRefusePostOwnerTransfer(
				0,
				1,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			msg: types.NewMsgRefusePostOwnerTransfer(
				1,
				0,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "invalid receiver returns error",
			msg: types.NewMsgRefusePostOwnerTransfer(
				1,
				0,
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgRefusePostOwnerTransfer,
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

func TestMsgRefusePostOwnerTransfer_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgRefusePostOwnerTransfer","value":{"post_id":"1","receiver":"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgRefusePostOwnerTransfer.GetSignBytes()))
}

func TestMsgRefusePostOwnerTransfer_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgRefusePostOwnerTransfer.Receiver)
	require.Equal(t, []sdk.AccAddress{addr}, msgRefusePostOwnerTransfer.GetSigners())
}
