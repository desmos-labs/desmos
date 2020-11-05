package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/magpie/types"
)

// ----------------------
// --- MsgCreateSession
// ----------------------

func TestMsgCreateSession_Route(t *testing.T) {
	msg := types.NewMsgCreateSession(
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		"cosmos",
		"cosmos1njrqah832yfdv8yhxnrskerzxhj5zj9e563uge",
		"cosmospub1addwnpepqf06gxm8tf4u9af99zsuphr2jmqvr2t956me5rcx9kywmrtg6jewy8gjtcs",
		"QmZh...===",
	)
	actual := msg.Route()
	require.Equal(t, types.RouterKey, actual)
}

func TestMsgCreateSession_Type(t *testing.T) {
	msg := types.NewMsgCreateSession(
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		"cosmos",
		"cosmos1njrqah832yfdv8yhxnrskerzxhj5zj9e563uge",
		"cosmospub1addwnpepqf06gxm8tf4u9af99zsuphr2jmqvr2t956me5rcx9kywmrtg6jewy8gjtcs",
		"QmZh...===",
	)
	actual := msg.Type()
	require.Equal(t, types.ActionCreationSession, actual)
}

func TestMsgCreateSession_ValidateBasic(t *testing.T) {
	tests := []struct {
		name   string
		msg    *types.MsgCreateSession
		expErr error
	}{
		{
			name: "Invalid owner",
			msg: types.NewMsgCreateSession(
				"",
				"cosmos",
				"cosmos1njrqah832yfdv8yhxnrskerzxhj5zj9e563uge",
				"cosmospub1addwnpepqf06gxm8tf4u9af99zsuphr2jmqvr2t956me5rcx9kywmrtg6jewy8gjtcs",
				"QmZh...===",
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, ""),
		},
		{
			name: "Invalid namespace",
			msg: types.NewMsgCreateSession(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"  ",
				"cosmos1njrqah832yfdv8yhxnrskerzxhj5zj9e563uge",
				"cosmospub1addwnpepqf06gxm8tf4u9af99zsuphr2jmqvr2t956me5rcx9kywmrtg6jewy8gjtcs",
				"QmZh...===",
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "session namespace cannot be empty"),
		},
		{
			name: "Invalid external owner",
			msg: types.NewMsgCreateSession(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos",
				"   ",
				"cosmospub1addwnpepqf06gxm8tf4u9af99zsuphr2jmqvr2t956me5rcx9kywmrtg6jewy8gjtcs",
				"QmZh...===",
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "session external owner cannot be empty"),
		},
		{
			name: "Invalid public key",
			msg: types.NewMsgCreateSession(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos",
				"cosmos1njrqah832yfdv8yhxnrskerzxhj5zj9e563uge",
				"   ",
				"QmZh...===",
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "signer's public key cannot be empty"),
		},
		{
			name: "Invalid signature",
			msg: types.NewMsgCreateSession(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos",
				"cosmos1njrqah832yfdv8yhxnrskerzxhj5zj9e563uge",
				"cosmospub1addwnpepqf06gxm8tf4u9af99zsuphr2jmqvr2t956me5rcx9kywmrtg6jewy8gjtcs",
				"  ",
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "session signature cannot be empty"),
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
	msg := types.NewMsgCreateSession(
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		"cosmos",
		"cosmos1njrqah832yfdv8yhxnrskerzxhj5zj9e563uge",
		"cosmospub1addwnpepqf06gxm8tf4u9af99zsuphr2jmqvr2t956me5rcx9kywmrtg6jewy8gjtcs",
		"QmZh...===",
	)

	actual := msg.GetSignBytes()
	expected := `{"type":"desmos/MsgCreateSession","value":{"external_owner":"cosmos1njrqah832yfdv8yhxnrskerzxhj5zj9e563uge","namespace":"cosmos","owner":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","pub_key":"cosmospub1addwnpepqf06gxm8tf4u9af99zsuphr2jmqvr2t956me5rcx9kywmrtg6jewy8gjtcs","signature":"QmZh...==="}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgCreateSession_GetSigners(t *testing.T) {
	msg := types.NewMsgCreateSession(
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		"cosmos",
		"cosmos1njrqah832yfdv8yhxnrskerzxhj5zj9e563uge",
		"cosmospub1addwnpepqf06gxm8tf4u9af99zsuphr2jmqvr2t956me5rcx9kywmrtg6jewy8gjtcs",
		"QmZh...===",
	)

	addr, _ := sdk.AccAddressFromBech32(msg.Owner)
	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())
}
