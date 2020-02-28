package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/magpie/internal/types"
	"github.com/stretchr/testify/require"
)

// ----------------------
// --- MsgCreateSession
// ----------------------

var testOwner, _ = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
var msgShareDocumentSchema = types.NewMsgCreateSession(
	testOwner,
	"cosmos",
	"cosmos1njrqah832yfdv8yhxnrskerzxhj5zj9e563uge",
	"cosmospub1addwnpepqf06gxm8tf4u9af99zsuphr2jmqvr2t956me5rcx9kywmrtg6jewy8gjtcs",
	"QmZh...===",
)

func TestMsgCreateSession_Route(t *testing.T) {
	actual := msgShareDocumentSchema.Route()
	require.Equal(t, "magpie", actual)
}

func TestMsgCreateSession_Type(t *testing.T) {
	actual := msgShareDocumentSchema.Type()
	require.Equal(t, "create_session", actual)
}

func TestMsgCreateSession_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    types.MsgCreateSession
		expErr error
	}{
		{
			name: "Invalid owner",
			msg: types.NewMsgCreateSession(
				nil,
				"cosmos",
				"cosmos1njrqah832yfdv8yhxnrskerzxhj5zj9e563uge",
				"cosmospub1addwnpepqf06gxm8tf4u9af99zsuphr2jmqvr2t956me5rcx9kywmrtg6jewy8gjtcs",
				"QmZh...===",
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid session owner: "),
		},
		{
			name: "Invalid namespace",
			msg: types.NewMsgCreateSession(
				testOwner,
				"  ",
				"cosmos1njrqah832yfdv8yhxnrskerzxhj5zj9e563uge",
				"cosmospub1addwnpepqf06gxm8tf4u9af99zsuphr2jmqvr2t956me5rcx9kywmrtg6jewy8gjtcs",
				"QmZh...===",
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Session namespace cannot be empty"),
		},
		{
			name: "Invalid external owner",
			msg: types.NewMsgCreateSession(
				testOwner,
				"cosmos",
				"   ",
				"cosmospub1addwnpepqf06gxm8tf4u9af99zsuphr2jmqvr2t956me5rcx9kywmrtg6jewy8gjtcs",
				"QmZh...===",
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Session external owner cannot be empty"),
		},
		{
			name: "Invalid public key",
			msg: types.NewMsgCreateSession(
				testOwner,
				"cosmos",
				"cosmos1njrqah832yfdv8yhxnrskerzxhj5zj9e563uge",
				"   ",
				"QmZh...===",
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Signer public key cannot be empty"),
		},
		{
			name: "Invalid signature",
			msg: types.NewMsgCreateSession(
				testOwner,
				"cosmos",
				"cosmos1njrqah832yfdv8yhxnrskerzxhj5zj9e563uge",
				"cosmospub1addwnpepqf06gxm8tf4u9af99zsuphr2jmqvr2t956me5rcx9kywmrtg6jewy8gjtcs",
				"  ",
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Session signature cannot be empty"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr.Error(), test.msg.ValidateBasic().Error())
		})
	}
}

func TestMsgCreateSession_GetSignBytes(t *testing.T) {
	actual := msgShareDocumentSchema.GetSignBytes()
	expected := `{"type":"desmos/MsgCreateSession","value":{"external_owner":"cosmos1njrqah832yfdv8yhxnrskerzxhj5zj9e563uge","namespace":"cosmos","owner":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","pub_key":"cosmospub1addwnpepqf06gxm8tf4u9af99zsuphr2jmqvr2t956me5rcx9kywmrtg6jewy8gjtcs","signature":"QmZh...==="}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgCreateSession_GetSigners(t *testing.T) {
	actual := msgShareDocumentSchema.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgShareDocumentSchema.Owner, actual[0])
}
