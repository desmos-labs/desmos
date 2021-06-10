package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
	tests := []struct {
		name  string
		msg   *types.MsgRequestDTagTransfer
		error error
	}{
		{
			name:  "Empty current owner returns error",
			msg:   types.NewMsgRequestDTagTransfer("", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address: "),
		},
		{
			name:  "Empty receiving user returns error",
			msg:   types.NewMsgRequestDTagTransfer("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", ""),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid receiver address: "),
		},
		{
			name: "Equals current owner and receiving user returns error",
			msg: types.NewMsgRequestDTagTransfer(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different"),
		},
		{
			name: "No errors message",
			msg: types.NewMsgRequestDTagTransfer(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
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

func TestMsgRequestDTagTransfer_GetSignBytes(t *testing.T) {
	actual := msgRequestTransferDTag.GetSignBytes()
	expected := `{"type":"desmos/MsgRequestDTagTransfer","value":{"receiver":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","sender":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(actual))
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
	tests := []struct {
		name  string
		msg   *types.MsgAcceptDTagTransfer
		error error
	}{
		{
			name: "Empty sender user returns error",
			msg: types.NewMsgAcceptDTagTransfer(
				"dtag",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address: "),
		},
		{
			name: "Empty receiver user returns error",
			msg: types.NewMsgAcceptDTagTransfer(
				"dtag",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid receiver address: "),
		},
		{
			name: "Equals current owner and receiving user returns error",
			msg: types.NewMsgAcceptDTagTransfer(
				"dtag",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different"),
		},
		{
			name: "Empty newDTag returns error",
			msg: types.NewMsgAcceptDTagTransfer(
				"",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "new DTag can't be empty"),
		},
		{
			name: "No errors message",
			msg: types.NewMsgAcceptDTagTransfer(
				"dtag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
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

func TestMsgAcceptDTagTransfer_GetSignBytes(t *testing.T) {
	actual := msgAcceptDTagTransfer.GetSignBytes()
	expected := `{"type":"desmos/MsgAcceptDTagTransfer","value":{"new_dtag":"dtag","receiver":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","sender":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"}}`
	require.Equal(t, expected, string(actual))
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
	tests := []struct {
		name  string
		msg   *types.MsgRefuseDTagTransfer
		error error
	}{
		{
			name: "Empty sender returns error",
			msg: types.NewMsgRefuseDTagTransferRequest(
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address: "),
		},
		{
			name: "Empty receiver returns error",
			msg: types.NewMsgRefuseDTagTransferRequest(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid receiver address: "),
		},
		{
			name: "Equals sender and receiver returns error",
			msg: types.NewMsgRefuseDTagTransferRequest(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different"),
		},
		{
			name: "No error message",
			msg: types.NewMsgRefuseDTagTransferRequest(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
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

func TestMsgRejectDTagRequest_GetSignBytes(t *testing.T) {
	actual := msgRejectDTagTransfer.GetSignBytes()
	expected := `{"type":"desmos/MsgRefuseDTagTransfer","value":{"receiver":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","sender":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"}}`
	require.Equal(t, expected, string(actual))
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
	tests := []struct {
		name  string
		msg   *types.MsgCancelDTagTransfer
		error error
	}{
		{
			name: "Empty receiver returns error",
			msg: types.NewMsgCancelDTagTransferRequest(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid receiver address: "),
		},
		{
			name: "Empty sender returns error",
			msg: types.NewMsgCancelDTagTransferRequest(
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address: "),
		},
		{
			name: "Equals sender and receiver returns error",
			msg: types.NewMsgCancelDTagTransferRequest(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different"),
		},
		{
			name: "No error message",
			msg: types.NewMsgCancelDTagTransferRequest(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
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

func TestMsgCancelDTagRequest_GetSignBytes(t *testing.T) {
	actual := msgCancelDTagTransferReq.GetSignBytes()
	expected := `{"type":"desmos/MsgCancelDTagTransfer","value":{"receiver":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","sender":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgCancelDTagRequest_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgCancelDTagTransferReq.Sender)
	require.Equal(t, []sdk.AccAddress{addr}, msgCancelDTagTransferReq.GetSigners())
}
