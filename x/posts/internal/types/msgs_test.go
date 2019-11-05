package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/assert"
)

// ----------------------
// --- MsgCreatePost
// ----------------------

var testOwner, _ = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
var timeZone, _ = time.LoadLocation("UTC")
var msgCreatePost = types.MsgCreatePost{
	ParentID:      types.PostID(53),
	Message:       "My new post",
	Owner:         testOwner,
	Namespace:     "cosmos",
	ExternalOwner: "cosmos1njrqah832yfdv8yhxnrskerzxhj5zj9e563uge",
}

func TestMsgCreatePost_Route(t *testing.T) {
	actual := msgCreatePost.Route()
	assert.Equal(t, "posts", actual)
}

func TestMsgCreatePost_Type(t *testing.T) {
	actual := msgCreatePost.Type()
	assert.Equal(t, "create_post", actual)
}

func TestMsgCreatePost_ValidateBasic_Schema_valid(t *testing.T) {
	err := msgCreatePost.ValidateBasic()
	assert.Nil(t, err)
}

func TestMsgCreatePost_GetSignBytes(t *testing.T) {
	actual := msgCreatePost.GetSignBytes()
	expected := `{"type":"desmos/MsgCreatePost","value":{"external_owner":"cosmos1njrqah832yfdv8yhxnrskerzxhj5zj9e563uge","message":"My new post","namespace":"cosmos","owner":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","parent_id":"53"}}`
	assert.Equal(t, expected, string(actual))
}

func TestMsgCreatePost_GetSigners(t *testing.T) {
	actual := msgCreatePost.GetSigners()
	assert.Equal(t, 1, len(actual))
	assert.Equal(t, msgCreatePost.Owner, actual[0])
}

// ----------------------
// --- MsgEditPost
// ----------------------

var msgEditPost = types.MsgEditPost{
	PostID:  types.PostID(94),
	Message: "Edited post message",
	Editor:  testOwner,
}

func TestMsgEditPost_Route(t *testing.T) {
	actual := msgEditPost.Route()
	assert.Equal(t, "posts", actual)
}

func TestMsgEditPost_Type(t *testing.T) {
	actual := msgEditPost.Type()
	assert.Equal(t, "edit_post", actual)
}

func TestMsgEditPost_ValidateBasic_Schema_valid(t *testing.T) {
	err := msgEditPost.ValidateBasic()
	assert.Nil(t, err)
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

var msgLike = types.MsgLikePost{
	PostID:        types.PostID(94),
	Liker:         testOwner,
	Namespace:     "cosmos",
	ExternalLiker: "cosmos1njrqah832yfdv8yhxnrskerzxhj5zj9e563uge",
}

func TestMsgLikePost_Route(t *testing.T) {
	actual := msgLike.Route()
	assert.Equal(t, "posts", actual)
}

func TestMsgLikePost_Type(t *testing.T) {
	actual := msgLike.Type()
	assert.Equal(t, "like_post", actual)
}

func TestMsgLikePost_ValidateBasic_Schema_valid(t *testing.T) {
	err := msgLike.ValidateBasic()
	assert.Nil(t, err)
}

func TestMsgLikePost_GetSignBytes(t *testing.T) {
	actual := msgLike.GetSignBytes()
	expected := `{"type":"desmos/MsgLikePost","value":{"external_liker":"cosmos1njrqah832yfdv8yhxnrskerzxhj5zj9e563uge","liker":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","namespace":"cosmos","post_id":"94"}}`
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

var msgUnlikePost = types.MsgUnlikePost{
	PostID: types.PostID(94),
	Liker:  testOwner,
	Time:   time.Date(2019, 11, 2, 10, 38, 5, 96000000, timeZone),
}

func TestMsgUnlikePost_Route(t *testing.T) {
	actual := msgUnlikePost.Route()
	assert.Equal(t, "posts", actual)
}

func TestMsgUnlikePost_Type(t *testing.T) {
	actual := msgUnlikePost.Type()
	assert.Equal(t, "unlike_post", actual)
}

func TestMsgUnlikePost_ValidateBasic_Schema_valid(t *testing.T) {
	err := msgUnlikePost.ValidateBasic()
	assert.Nil(t, err)
}

func TestMsgUnlikePost_GetSignBytes(t *testing.T) {
	actual := msgUnlikePost.GetSignBytes()
	expected := `{"type":"desmos/MsgUnlikePost","value":{"liker":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","post_id":"94","time":"2019-11-02T10:38:05.096Z"}}`
	assert.Equal(t, expected, string(actual))
}

func TestMsgUnlikePost_GetSigners(t *testing.T) {
	actual := msgUnlikePost.GetSigners()
	assert.Equal(t, 1, len(actual))
	assert.Equal(t, msgUnlikePost.Liker, actual[0])
}
