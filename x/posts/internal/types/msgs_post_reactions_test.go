package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/require"
)

// ----------------------
// --- MsgAddPostReaction
// ----------------------

var msgPostReaction = types.NewMsgAddPostReaction(id, "like", testOwner)

func TestMsgAddPostReaction_Route(t *testing.T) {
	actual := msgPostReaction.Route()
	require.Equal(t, "posts", actual)
}

func TestMsgAddPostReaction_Type(t *testing.T) {
	actual := msgPostReaction.Type()
	require.Equal(t, "add_post_reaction", actual)
}

func TestMsgAddPostReaction_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   types.MsgAddPostReaction
		error error
	}{
		{
			name:  "Invalid post id returns error",
			msg:   types.NewMsgAddPostReaction(nil, ":like:", testOwner),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid post id"),
		},
		{
			name:  "Invalid user returns error",
			msg:   types.NewMsgAddPostReaction(id, ":like:", nil),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid user address: "),
		},
		{
			name:  "Invalid value returns error",
			msg:   types.NewMsgAddPostReaction(id, "like", testOwner),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Reaction short code must be an emoji short code"),
		},
		{
			name:  "Valid message returns no error",
			msg:   types.NewMsgAddPostReaction(id, ":like:", testOwner),
			error: nil,
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

func TestMsgAddPostReaction_GetSignBytes(t *testing.T) {
	actual := msgPostReaction.GetSignBytes()
	expected := `{"type":"desmos/MsgAddPostReaction","value":{"post_id":"3139646530326531303563363861363065343563323839626666","user":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","value":"like"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgAddPostReaction_GetSigners(t *testing.T) {
	actual := msgPostReaction.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgPostReaction.User, actual[0])
}

// ----------------------
// --- MsgRemovePostReaction
// ----------------------

var msgUnlikePost = types.NewMsgRemovePostReaction(id, testOwner, "like")

func TestMsgRemovePostReaction_Route(t *testing.T) {
	actual := msgUnlikePost.Route()
	require.Equal(t, "posts", actual)
}

func TestMsgRemovePostReaction_Type(t *testing.T) {
	actual := msgUnlikePost.Type()
	require.Equal(t, "remove_post_reaction", actual)
}

func TestMsgRemovePostReaction_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   types.MsgRemovePostReaction
		error error
	}{
		{
			name:  "Invalid post id returns error",
			msg:   types.NewMsgRemovePostReaction(nil, testOwner, "like"),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid post id"),
		},
		{
			name:  "Invalid user address: ",
			msg:   types.NewMsgRemovePostReaction(id, nil, "like"),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid user address: "),
		},
		{
			name:  "Invalid value returns no error",
			msg:   types.NewMsgRemovePostReaction(id, testOwner, ""),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Reaction value cannot be empty nor blank"),
		},
		{
			name:  "Valid message returns no error",
			msg:   types.NewMsgRemovePostReaction(id, testOwner, "like"),
			error: nil,
		},
	}

	for _, test := range tests {
		returnedError := test.msg.ValidateBasic()
		if test.error == nil {
			require.Nil(t, returnedError)
		} else {
			require.NotNil(t, returnedError)
			require.Equal(t, test.error.Error(), returnedError.Error())
		}
	}
}

func TestMsgRemovePostReaction_GetSignBytes(t *testing.T) {
	actual := msgUnlikePost.GetSignBytes()
	expected := `{"type":"desmos/MsgRemovePostReaction","value":{"post_id":"3139646530326531303563363861363065343563323839626666","reaction":"like","user":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgRemovePostReaction_GetSigners(t *testing.T) {
	actual := msgUnlikePost.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgUnlikePost.User, actual[0])
}
