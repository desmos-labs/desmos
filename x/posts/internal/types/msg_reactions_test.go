package types_test

import (
	"fmt"
	"testing"

	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/require"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var msgRegisterReaction = types.NewMsgRegisterReaction(testOwner, ":smile:", "https://smile.jpg",
	"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")

func TestMsgRegisterReaction_Route(t *testing.T) {
	actual := msgRegisterReaction.Route()
	require.Equal(t, "posts", actual)
}

func TestMsgRegisterReaction_Type(t *testing.T) {
	actual := msgRegisterReaction.Type()
	require.Equal(t, "register_reaction", actual)
}

func TestMsgRegisterReaction_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   types.MsgRegisterReaction
		error error
	}{
		{
			name: "Invalid creator returns error",
			msg: types.NewMsgRegisterReaction(nil, ":smile:", "https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid creator address: %s", "")),
		},
		{
			name: "Empty short code returns error",
			msg: types.NewMsgRegisterReaction(testOwner, "", "https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "reaction short code must be an emoji short code"),
		},
		{
			name: "Invalid short code returns error",
			msg: types.NewMsgRegisterReaction(testOwner, ":smile", "https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "reaction short code must be an emoji short code"),
		},
		{
			name: "Empty value returns error",
			msg: types.NewMsgRegisterReaction(testOwner, ":smile:", "",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "reaction value should be a URL or an emoji unicode"),
		},
		{
			name: "Invalid value returns error (url)",
			msg: types.NewMsgRegisterReaction(testOwner, ":smile:", "htp://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "reaction value should be a URL or an emoji unicode"),
		},
		{
			name: "Invalid value returns error (unicode)",
			msg: types.NewMsgRegisterReaction(testOwner, ":smile:", "U+1",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "reaction value should be a URL or an emoji unicode"),
		},
		{
			name: "Invalid subspace returns error",
			msg: types.NewMsgRegisterReaction(testOwner, ":smile:", "https://smile.jpg",
				"1234"),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("reaction subspace must be a valid sha-256 hash")),
		},
	}

	for _, test := range tests {
		test := test
		returnedError := test.msg.ValidateBasic()
		if test.error == nil {
			require.Nil(t, returnedError)
		} else {
			require.NotNil(t, returnedError)
			require.Equal(t, test.error.Error(), returnedError.Error())
		}
	}
}

func TestMsgRegisterReaction_GetSignBytes(t *testing.T) {
	actual := msgRegisterReaction.GetSignBytes()
	expected := `{"type":"desmos/MsgRegisterReaction","value":{"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","short_code":":smile:","subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e","value":"https://smile.jpg"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgRegisterReaction_GetSigners(t *testing.T) {
	actual := msgRegisterReaction.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgRegisterReaction.Creator, actual[0])
}
