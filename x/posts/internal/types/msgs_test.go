package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/assert"
)

// ----------------------
// --- MsgCreatePost
// ----------------------

var testOwner, _ = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
var msgCreatePost = types.NewMsgCreatePost("My new post", types.PostID(53), false, testOwner)

func TestMsgCreatePost_Route(t *testing.T) {
	actual := msgCreatePost.Route()
	assert.Equal(t, "posts", actual)
}

func TestMsgCreatePost_Type(t *testing.T) {
	actual := msgCreatePost.Type()
	assert.Equal(t, "create_post", actual)
}

func TestMsgCreatePost_ValidateBasic(t *testing.T) {
	creator, _ := sdk.AccAddressFromBech32("cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h")
	tests := []struct {
		name  string
		msg   types.MsgCreatePost
		error sdk.Error
	}{
		{
			name:  "Empty owner returns error",
			msg:   types.NewMsgCreatePost("Message", types.PostID(0), false, nil),
			error: sdk.ErrInvalidAddress("Invalid creator address: "),
		},
		{
			name:  "Empty post message returns error",
			msg:   types.NewMsgCreatePost("", types.PostID(0), false, creator),
			error: sdk.ErrUnknownRequest("Post message cannot be empty"),
		},
		{
			name:  "Valid message does not return any error",
			msg:   types.NewMsgCreatePost("Message", types.PostID(0), false, creator),
			error: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.error, test.msg.ValidateBasic())
		})
	}

	err := msgCreatePost.ValidateBasic()
	assert.Nil(t, err)
}

func TestMsgCreatePost_GetSignBytes(t *testing.T) {
	actual := msgCreatePost.GetSignBytes()
	expected := `{"type":"desmos/MsgCreatePost","value":{"allows_comments":false,"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","message":"My new post","parent_id":"53"}}`
	assert.Equal(t, expected, string(actual))
}

func TestMsgCreatePost_GetSigners(t *testing.T) {
	actual := msgCreatePost.GetSigners()
	assert.Equal(t, 1, len(actual))
	assert.Equal(t, msgCreatePost.Creator, actual[0])
}

// ----------------------
// --- MsgEditPost
// ----------------------

var msgEditPost = types.NewMsgEditPost(types.PostID(94), "Edited post message", testOwner)

func TestMsgEditPost_Route(t *testing.T) {
	actual := msgEditPost.Route()
	assert.Equal(t, "posts", actual)
}

func TestMsgEditPost_Type(t *testing.T) {
	actual := msgEditPost.Type()
	assert.Equal(t, "edit_post", actual)
}

func TestMsgEditPost_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   types.MsgEditPost
		error sdk.Error
	}{
		{
			name:  "Invalid post id returns error",
			msg:   types.NewMsgEditPost(types.PostID(0), "Edited post message", testOwner),
			error: sdk.ErrUnknownRequest("Invalid post id"),
		},
		{
			name:  "Invalid editor returns error",
			msg:   types.NewMsgEditPost(types.PostID(10), "Edited post message", nil),
			error: sdk.ErrInvalidAddress("Invalid editor address: "),
		},
		{
			name:  "Invalid message returns error",
			msg:   types.NewMsgEditPost(types.PostID(10), "", testOwner),
			error: sdk.ErrUnknownRequest("Post message cannot be empty"),
		},
		{
			name:  "Valid message returns no error",
			msg:   types.NewMsgEditPost(types.PostID(10), "Edited post message", testOwner),
			error: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.error, test.msg.ValidateBasic())
		})
	}
}

func TestMsgEditPost_GetSignBytes(t *testing.T) {
	actual := msgEditPost.GetSignBytes()
	expected := `{"type":"desmos/MsgEditPost","value":{"editor":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","message":"Edited post message","post_id":"94"}}`
	assert.Equal(t, expected, string(actual))
}

func TestMsgEditPost_GetSigners(t *testing.T) {
	actual := msgEditPost.GetSigners()
	assert.Equal(t, 1, len(actual))
	assert.Equal(t, msgEditPost.Editor, actual[0])
}

// ----------------------
// --- MsgLikePost
// ----------------------

var msgLike = types.NewMsgLikePost(types.PostID(94), testOwner)

func TestMsgLikePost_Route(t *testing.T) {
	actual := msgLike.Route()
	assert.Equal(t, "posts", actual)
}

func TestMsgLikePost_Type(t *testing.T) {
	actual := msgLike.Type()
	assert.Equal(t, "like_post", actual)
}

func TestMsgLikePost_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   types.MsgLikePost
		error sdk.Error
	}{
		{
			name:  "Invalid post id returns error",
			msg:   types.NewMsgLikePost(types.PostID(0), testOwner),
			error: sdk.ErrUnknownRequest("Invalid post id"),
		},
		{
			name:  "Invalid liker returns error",
			msg:   types.NewMsgLikePost(types.PostID(5), nil),
			error: sdk.ErrInvalidAddress("Invalid liker address: "),
		},
		{
			name:  "Valid message returns no error",
			msg:   types.NewMsgLikePost(types.PostID(10), testOwner),
			error: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.error, test.msg.ValidateBasic())
		})
	}
}

func TestMsgLikePost_GetSignBytes(t *testing.T) {
	actual := msgLike.GetSignBytes()
	expected := `{"type":"desmos/MsgLikePost","value":{"liker":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","post_id":"94"}}`
	assert.Equal(t, expected, string(actual))
}

func TestMsgLikePost_GetSigners(t *testing.T) {
	actual := msgLike.GetSigners()
	assert.Equal(t, 1, len(actual))
	assert.Equal(t, msgLike.Liker, actual[0])
}

// ----------------------
// --- MsgUnlikePost
// ----------------------

var msgUnlikePost = types.NewMsgUnlikePost(types.PostID(94), testOwner)

func TestMsgUnlikePost_Route(t *testing.T) {
	actual := msgUnlikePost.Route()
	assert.Equal(t, "posts", actual)
}

func TestMsgUnlikePost_Type(t *testing.T) {
	actual := msgUnlikePost.Type()
	assert.Equal(t, "unlike_post", actual)
}

func TestMsgUnlikePost_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   types.MsgUnlikePost
		error sdk.Error
	}{
		{
			name:  "Invalid post id returns error",
			msg:   types.NewMsgUnlikePost(types.PostID(0), testOwner),
			error: sdk.ErrUnknownRequest("Invalid post id"),
		},
		{
			name:  "Invalid liker address: ",
			msg:   types.NewMsgUnlikePost(types.PostID(10), nil),
			error: sdk.ErrInvalidAddress("Invalid liker address: "),
		},
		{
			name:  "Valid message returns no error",
			msg:   types.NewMsgUnlikePost(types.PostID(10), testOwner),
			error: nil,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.error, test.msg.ValidateBasic())
	}
}

func TestMsgUnlikePost_GetSignBytes(t *testing.T) {
	actual := msgUnlikePost.GetSignBytes()
	expected := `{"type":"desmos/MsgUnlikePost","value":{"liker":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","post_id":"94"}}`
	assert.Equal(t, expected, string(actual))
}

func TestMsgUnlikePost_GetSigners(t *testing.T) {
	actual := msgUnlikePost.GetSigners()
	assert.Equal(t, 1, len(actual))
	assert.Equal(t, msgUnlikePost.Liker, actual[0])
}
