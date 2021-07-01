package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"

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
	require.Equal(t, "request_dtag", msgRequestTransferDTag.Type())
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

func TestMsgRequestDTagTransfer_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgRequestDTagTransfer","value":{"receiver":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","sender":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(msgRequestTransferDTag.GetSignBytes()))
}

func TestMsgRequestDTagTransfer_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgRequestTransferDTag.Sender)
	require.Equal(t, []sdk.AccAddress{addr}, msgRequestTransferDTag.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgAcceptDTagTransfer = types.NewMsgAcceptDTagTransfer(
	"dtag",
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
)

func TestMsgAcceptDTagTransfer_Route(t *testing.T) {
	require.Equal(t, "profiles", msgAcceptDTagTransfer.Route())
}

func TestMsgAcceptDTagTransfer_Type(t *testing.T) {
	require.Equal(t, "accept_dtag_request", msgAcceptDTagTransfer.Type())
}

func TestMsgAcceptDTagTransfer_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgAcceptDTagTransfer
		shouldErr bool
	}{
		{
			name: "empty sender user returns error",
			msg: types.NewMsgAcceptDTagTransfer(
				"dtag",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "empty receiver user returns error",
			msg: types.NewMsgAcceptDTagTransfer(
				"dtag",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
			),
			shouldErr: true,
		},
		{
			name: "equals current owner and receiving user returns error",
			msg: types.NewMsgAcceptDTagTransfer(
				"dtag",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "empty new DTag returns error",
			msg: types.NewMsgAcceptDTagTransfer(
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

func TestMsgAcceptDTagTransfer_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgAcceptDTagTransfer","value":{"new_dtag":"dtag","receiver":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","sender":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"}}`
	require.Equal(t, expected, string(msgAcceptDTagTransfer.GetSignBytes()))
}

func TestMsgAcceptDTagTransfer_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgAcceptDTagTransfer.Receiver)
	require.Equal(t, []sdk.AccAddress{addr}, msgAcceptDTagTransfer.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgRejectDTagTransfer = types.NewMsgRefuseDTagTransferRequest(
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
)

func TestMsgRejectDTagRequest_Route(t *testing.T) {
	require.Equal(t, "profiles", msgRejectDTagTransfer.Route())
}

func TestMsgRejectDTagRequest_Type(t *testing.T) {
	require.Equal(t, "refuse_dtag_request", msgRejectDTagTransfer.Type())
}

func TestMsgRejectDTagRequest_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgRefuseDTagTransfer
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

func TestMsgRejectDTagRequest_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgRefuseDTagTransfer","value":{"receiver":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","sender":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"}}`
	require.Equal(t, expected, string(msgRejectDTagTransfer.GetSignBytes()))
}

func TestMsgRejectDTagRequest_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgRejectDTagTransfer.Receiver)
	require.Equal(t, []sdk.AccAddress{addr}, msgRejectDTagTransfer.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgCancelDTagTransferReq = types.NewMsgCancelDTagTransferRequest(
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
)

func TestMsgCancelDTagRequest_Route(t *testing.T) {
	require.Equal(t, "profiles", msgCancelDTagTransferReq.Route())
}

func TestMsgCancelDTagRequest_Type(t *testing.T) {
	require.Equal(t, "cancel_dtag_request", msgCancelDTagTransferReq.Type())
}

func TestMsgCancelDTagRequest_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgCancelDTagTransfer
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

func TestMsgCancelDTagRequest_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgCancelDTagTransfer","value":{"receiver":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","sender":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(msgCancelDTagTransferReq.GetSignBytes()))
}

func TestMsgCancelDTagRequest_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgCancelDTagTransferReq.Sender)
	require.Equal(t, []sdk.AccAddress{addr}, msgCancelDTagTransferReq.GetSigners())
}
