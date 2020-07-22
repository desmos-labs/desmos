package msgs_test

import (
	"testing"

	commonerrors "github.com/desmos-labs/desmos/x/commons/types/errors"
	postserrors "github.com/desmos-labs/desmos/x/posts/types/errors"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/posts/types/msgs"
)

var msgRegisterReaction = msgs.NewMsgRegisterReaction(testOwner, ":smile:", "https://smile.jpg",
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
		msg   msgs.MsgRegisterReaction
		error error
	}{
		{
			name: "Invalid creator returns error",
			msg: msgs.NewMsgRegisterReaction(nil, ":smile:", "https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator address: "),
		},
		{
			name: "Empty short code returns error",
			msg: msgs.NewMsgRegisterReaction(testOwner, "", "https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			error: sdkerrors.Wrap(postserrors.ErrInvalidReactionCode, ""),
		},
		{
			name: "Invalid short code returns error",
			msg: msgs.NewMsgRegisterReaction(testOwner, ":smile", "https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			error: sdkerrors.Wrap(postserrors.ErrInvalidReactionCode, ":smile"),
		},
		{
			name: "Empty value returns error",
			msg: msgs.NewMsgRegisterReaction(testOwner, ":smile:", "",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			error: sdkerrors.Wrap(commonerrors.ErrInvalidURI, "reaction value should be a valid uri"),
		},
		{
			name: "Invalid value returns error (url)",
			msg: msgs.NewMsgRegisterReaction(testOwner, ":smile:", "htp://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			error: sdkerrors.Wrap(commonerrors.ErrInvalidURI, "reaction value should be a valid uri"),
		},
		{
			name: "Invalid value returns error (unicode)",
			msg: msgs.NewMsgRegisterReaction(testOwner, ":smile:", "U+1",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			error: sdkerrors.Wrap(commonerrors.ErrInvalidURI, "reaction value should be a valid uri"),
		},
		{
			name:  "Valid emoji value returns no error",
			msg:   msgs.NewMsgRegisterReaction(testOwner, ":smile:", "💙", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			error: sdkerrors.Wrap(commonerrors.ErrInvalidURI, "reaction value should be a valid uri"),
		},
		{
			name: "Invalid subspace returns error",
			msg: msgs.NewMsgRegisterReaction(testOwner, ":smile:", "https://smile.jpg",
				"1234"),
			error: sdkerrors.Wrap(postserrors.ErrInvalidSubspace, "reaction subspace must be a valid sha-256 hash"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if test.error == nil {
				require.Nil(t, test.msg.ValidateBasic())
			} else {
				require.NotNil(t, test.msg.ValidateBasic())
				require.Equal(t, test.error.Error(), test.msg.ValidateBasic().Error())
			}
		})
	}
}

func TestMsgRegisterReaction_GetSignBytes(t *testing.T) {
	actual := msgRegisterReaction.GetSignBytes()
	expected := `{"type":"desmos/MsgRegisterReaction","value":{"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","shortcode":":smile:","subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e","value":"https://smile.jpg"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgRegisterReaction_GetSigners(t *testing.T) {
	actual := msgRegisterReaction.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgRegisterReaction.Creator, actual[0])
}
