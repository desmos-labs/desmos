package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v5/x/posts/types"
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
			name: "receiver equals to sender returns error",
			msg: types.NewMsgRequestPostOwnerTransfer(
				1,
				1,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "invalid receiver returns error",
			msg: types.NewMsgRequestPostOwnerTransfer(
				1,
				1,
				"",
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "invalid sender returns error",
			msg: types.NewMsgRequestPostOwnerTransfer(
				1,
				1,
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

var msgCancelPostOwnerTransferRequest = types.NewMsgCancelPostOwnerTransferRequest(
	1,
	1,
	"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
)

func TestMsgCancelPostOwnerTransferRequest_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgCancelPostOwnerTransferRequest.Route())
}

func TestMsgCancelPostOwnerTransferRequest_Type(t *testing.T) {
	require.Equal(t, types.ActionCancelPostOwnerTransfer, msgCancelPostOwnerTransferRequest.Type())
}

func TestMsgCancelPostOwnerTransferRequest_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgCancelPostOwnerTransferRequest
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgCancelPostOwnerTransferRequest(
				0,
				1,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			msg: types.NewMsgCancelPostOwnerTransferRequest(
				1,
				0,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "invalid sender returns error",
			msg: types.NewMsgCancelPostOwnerTransferRequest(
				1,
				1,
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgCancelPostOwnerTransferRequest,
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

func TestMsgCancelPostOwnerTransferRequest_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgCancelPostOwnerTransfer","value":{"post_id":"1","sender":"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgCancelPostOwnerTransferRequest.GetSignBytes()))
}

func TestMsgCancelPostOwnerTransferRequest_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgCancelPostOwnerTransferRequest.Sender)
	require.Equal(t, []sdk.AccAddress{addr}, msgCancelPostOwnerTransferRequest.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgAcceptPostOwnerTransferRequest = types.NewMsgAcceptPostOwnerTransferRequest(
	1,
	1,
	"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
)

func TestMsgAcceptPostOwnerTransferRequest_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgAcceptPostOwnerTransferRequest.Route())
}

func TestMsgAcceptPostOwnerTransferRequest_Type(t *testing.T) {
	require.Equal(t, types.ActionAcceptPostOwnerTransfer, msgAcceptPostOwnerTransferRequest.Type())
}

func TestMsgAcceptPostOwnerTransferRequest_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgAcceptPostOwnerTransferRequest
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgAcceptPostOwnerTransferRequest(
				0,
				1,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			msg: types.NewMsgAcceptPostOwnerTransferRequest(
				1,
				0,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "invalid receiver returns error",
			msg: types.NewMsgAcceptPostOwnerTransferRequest(
				1,
				1,
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgAcceptPostOwnerTransferRequest,
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

func TestMsgAcceptPostOwnerTransferRequest_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgAcceptPostOwnerTransfer","value":{"post_id":"1","receiver":"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgAcceptPostOwnerTransferRequest.GetSignBytes()))
}

func TestMsgAcceptPostOwnerTransferRequest_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgAcceptPostOwnerTransferRequest.Receiver)
	require.Equal(t, []sdk.AccAddress{addr}, msgAcceptPostOwnerTransferRequest.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgRefusePostOwnerTransferRequest = types.NewMsgRefusePostOwnerTransferRequest(
	1,
	1,
	"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
)

func TestMsgRefusePostOwnerTransferRequest_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgRefusePostOwnerTransferRequest.Route())
}

func TestMsgRefusePostOwnerTransferRequest_Type(t *testing.T) {
	require.Equal(t, types.ActionRefusePostOwnerTransfer, msgRefusePostOwnerTransferRequest.Type())
}

func TestMsgRefusePostOwnerTransferRequest_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgRefusePostOwnerTransferRequest
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgRefusePostOwnerTransferRequest(
				0,
				1,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			msg: types.NewMsgRefusePostOwnerTransferRequest(
				1,
				0,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "invalid receiver returns error",
			msg: types.NewMsgRefusePostOwnerTransferRequest(
				1,
				1,
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgRefusePostOwnerTransferRequest,
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

func TestMsgRefusePostOwnerTransferRequest_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgRefusePostOwnerTransfer","value":{"post_id":"1","receiver":"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgRefusePostOwnerTransferRequest.GetSignBytes()))
}

func TestMsgRefusePostOwnerTransferRequest_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgRefusePostOwnerTransferRequest.Receiver)
	require.Equal(t, []sdk.AccAddress{addr}, msgRefusePostOwnerTransferRequest.GetSigners())
}
