package msgs_test

import (
	"testing"

	postserrors "github.com/desmos-labs/desmos/x/posts/types/errors"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	emoji2 "github.com/desmos-labs/Go-Emoji-Utils"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/posts/types/models"
	"github.com/desmos-labs/desmos/x/posts/types/msgs"
)

// ----------------------
// --- MsgAddPostReaction
// ----------------------

func TestShortCodeRegEx(t *testing.T) {
	for _, emoji := range emoji2.Emojis {
		for _, shortcode := range emoji.Shortcodes {
			res := models.IsValidReactionCode(shortcode)
			if !res {
				println(shortcode)
			}
			require.True(t, res)
		}
	}
}

var msgPostReaction = msgs.NewMsgAddPostReaction(id, "like", testOwner)

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
		msg   msgs.MsgAddPostReaction
		error error
	}{
		{
			name:  "Invalid post id returns error",
			msg:   msgs.NewMsgAddPostReaction("", ":like:", testOwner),
			error: sdkerrors.Wrap(postserrors.ErrInvalidPostID, ""),
		},
		{
			name:  "Invalid user returns error",
			msg:   msgs.NewMsgAddPostReaction(id, ":like:", nil),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid user address: "),
		},
		{
			name:  "Invalid value returns error",
			msg:   msgs.NewMsgAddPostReaction(id, "like", testOwner),
			error: sdkerrors.Wrap(postserrors.ErrInvalidReactionCode, "like"),
		},
		{
			name:  "Valid message returns no error (with shortcode)",
			msg:   msgs.NewMsgAddPostReaction(id, ":like:", testOwner),
			error: nil,
		},
		{
			name:  "Valid message returns no error (with emoji)",
			msg:   msgs.NewMsgAddPostReaction(id, "🤩", testOwner),
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
	expected := `{"type":"desmos/MsgAddPostReaction","value":{"post_id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","reaction":"like","user":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
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

var msgUnlikePost = msgs.NewMsgRemovePostReaction(id, testOwner, "like")

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
		msg   msgs.MsgRemovePostReaction
		error error
	}{
		{
			name:  "Invalid post id returns error",
			msg:   msgs.NewMsgRemovePostReaction("", testOwner, ":+1:"),
			error: sdkerrors.Wrap(postserrors.ErrInvalidPostID, ""),
		},
		{
			name:  "Invalid user address: ",
			msg:   msgs.NewMsgRemovePostReaction(id, nil, ":like:"),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid user address: "),
		},
		{
			name:  "Blank value returns no error",
			msg:   msgs.NewMsgRemovePostReaction(id, testOwner, ""),
			error: sdkerrors.Wrap(postserrors.ErrInvalidReactionCode, ""),
		},
		{
			name:  "Valid message returns no error (with shortcode)",
			msg:   msgs.NewMsgRemovePostReaction(id, testOwner, ":+1:"),
			error: nil,
		},
		{
			name:  "Valid message returns no error (with emoji)",
			msg:   msgs.NewMsgRemovePostReaction(id, testOwner, "🤩"),
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
	expected := `{"type":"desmos/MsgRemovePostReaction","value":{"post_id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","reaction":"like","user":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgRemovePostReaction_GetSigners(t *testing.T) {
	actual := msgUnlikePost.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgUnlikePost.User, actual[0])
}
