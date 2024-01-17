package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/x/profiles/types"

	"github.com/stretchr/testify/require"
)

var msgRequestTransferDTag = types.NewMsgRequestDTagTransfer(
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
)

func TestMsgRequestDTagTransfer_Route(t *testing.T) {
	require.Equal(t, "profiles", msgRequestTransferDTag.Route())
}

func TestMsgRequestDTagTransfer_Type(t *testing.T) {
	require.Equal(t, "request_dtag_transfer", msgRequestTransferDTag.Type())
}

func TestMsgRequestDTagTransfer_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgRequestDTagTransfer
		shouldErr bool
	}{
		{
			name:      "empty current owner returns error",
			msg:       types.NewMsgRequestDTagTransfer("", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			shouldErr: true,
		},
		{
			name:      "empty receiving user returns error",
			msg:       types.NewMsgRequestDTagTransfer("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", ""),
			shouldErr: true,
		},
		{
			name: "equals current owner and receiving user returns error",
			msg: types.NewMsgRequestDTagTransfer(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name:      "valid message returns no error",
			msg:       msgRequestTransferDTag,
			shouldErr: false,
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

func TestMsgRequestDTagTransfer_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgRequestTransferDTag.Sender)
	require.Equal(t, []sdk.AccAddress{addr}, msgRequestTransferDTag.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgAcceptDTagTransfer = types.NewMsgAcceptDTagTransferRequest(
	"dtag",
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
)

func TestMsgAcceptDTagTransferRequest_Route(t *testing.T) {
	require.Equal(t, "profiles", msgAcceptDTagTransfer.Route())
}

func TestMsgAcceptDTagTransferRequest_Type(t *testing.T) {
	require.Equal(t, "accept_dtag_transfer_request", msgAcceptDTagTransfer.Type())
}

func TestMsgAcceptDTagTransferRequest_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgAcceptDTagTransferRequest
		shouldErr bool
	}{
		{
			name: "empty sender user returns error",
			msg: types.NewMsgAcceptDTagTransferRequest(
				"dtag",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "empty receiver user returns error",
			msg: types.NewMsgAcceptDTagTransferRequest(
				"dtag",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
			),
			shouldErr: true,
		},
		{
			name: "equals current owner and receiving user returns error",
			msg: types.NewMsgAcceptDTagTransferRequest(
				"dtag",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "empty new DTag returns error",
			msg: types.NewMsgAcceptDTagTransferRequest(
				"",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name:      "no errors message",
			msg:       msgAcceptDTagTransfer,
			shouldErr: false,
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

func TestMsgAcceptDTagTransferRequest_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgAcceptDTagTransfer.Receiver)
	require.Equal(t, []sdk.AccAddress{addr}, msgAcceptDTagTransfer.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgRejectDTagTransfer = types.NewMsgRefuseDTagTransferRequest(
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
)

func TestMsgRejectDTagTransferRequest_Route(t *testing.T) {
	require.Equal(t, "profiles", msgRejectDTagTransfer.Route())
}

func TestMsgRejectDTagTransferRequest_Type(t *testing.T) {
	require.Equal(t, "refuse_dtag_transfer_request", msgRejectDTagTransfer.Type())
}

func TestMsgRejectDTagTransferRequest_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgRefuseDTagTransferRequest
		shouldErr bool
	}{
		{
			name: "empty sender returns error",
			msg: types.NewMsgRefuseDTagTransferRequest(
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "empty receiver returns error",
			msg: types.NewMsgRefuseDTagTransferRequest(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
			),
			shouldErr: true,
		},
		{
			name: "equals sender and receiver returns error",
			msg: types.NewMsgRefuseDTagTransferRequest(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg: types.NewMsgRefuseDTagTransferRequest(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: false,
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

func TestMsgRejectDTagTransferRequest_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgRejectDTagTransfer.Receiver)
	require.Equal(t, []sdk.AccAddress{addr}, msgRejectDTagTransfer.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgCancelDTagTransferReq = types.NewMsgCancelDTagTransferRequest(
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
)

func TestMsgCancelDTagTransferRequest_Route(t *testing.T) {
	require.Equal(t, "profiles", msgCancelDTagTransferReq.Route())
}

func TestMsgCancelDTagTransferRequest_Type(t *testing.T) {
	require.Equal(t, "cancel_dtag_transfer_request", msgCancelDTagTransferReq.Type())
}

func TestMsgCancelDTagTransferRequest_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgCancelDTagTransferRequest
		shouldErr bool
	}{
		{
			name: "Empty receiver returns error",
			msg: types.NewMsgCancelDTagTransferRequest(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
			),
			shouldErr: true,
		},
		{
			name: "Empty sender returns error",
			msg: types.NewMsgCancelDTagTransferRequest(
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "Equals sender and receiver returns error",
			msg: types.NewMsgCancelDTagTransferRequest(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name:      "No error message",
			msg:       msgCancelDTagTransferReq,
			shouldErr: false,
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

func TestMsgCancelDTagTransferRequest_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgCancelDTagTransferReq.Sender)
	require.Equal(t, []sdk.AccAddress{addr}, msgCancelDTagTransferReq.GetSigners())
}
