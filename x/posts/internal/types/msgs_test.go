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
var date = time.Date(2020, 1, 1, 12, 0, 0, 0, timeZone)
var msgCreatePost = types.NewMsgCreatePost(
	"My new post",
	types.PostID(53),
	false,
	"desmos",
	map[string]string{},
	testOwner,
	date,
	types.PostMedias{
		types.PostMedia{
			URI:      "https://uri.com",
			MimeType: "text/plain",
		},
	},
)

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
			name: "Empty owner returns error",
			msg: types.NewMsgCreatePost(
				"Message",
				types.PostID(0),
				false,
				"desmos",
				map[string]string{},
				nil,
				date,
				msgCreatePost.Medias,
			),
			error: sdk.ErrInvalidAddress("Invalid creator address: "),
		},
		{
			name: "Empty message returns error",
			msg: types.NewMsgCreatePost(
				"",
				types.PostID(0),
				false,
				"desmos",
				map[string]string{},
				creator,
				date,
				msgCreatePost.Medias,
			),
			error: sdk.ErrUnknownRequest("Post message cannot be empty nor blank"),
		},
		{
			name: "Very long message returns error",
			msg: types.NewMsgCreatePost(
				`
				Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque massa felis, aliquam sed ipsum at, 
				mollis pharetra quam. Vestibulum nec nulla ante. Praesent sed dignissim turpis. Curabitur aliquam nunc 
				eu nisi porta, eu gravida purus faucibus. Duis commodo sagittis lacus, vitae luctus enim vulputate a. 
				Nulla tempor eget nunc vitae vulputate. Nulla facilities. Donec sollicitudin odio in arcu efficitur, 
				sit amet vestibulum diam ullamcorper. Ut ac dolor in velit gravida efficitur et et erat volutpat.
				`,
				types.PostID(0),
				false,
				"desmos",
				map[string]string{},
				creator,
				date,
				msgCreatePost.Medias,
			),
			error: sdk.ErrUnknownRequest("Post message cannot exceed 500 characters"),
		},
		{
			name: "Empty subspace returns error",
			msg: types.NewMsgCreatePost(
				"My message",
				types.PostID(0),
				false,
				"",
				map[string]string{},
				creator,
				date,
				msgCreatePost.Medias,
			),
			error: sdk.ErrUnknownRequest("Post subspace cannot be empty nor blank"),
		},
		{
			name: "More than 10 optional data returns error",
			msg: types.NewMsgCreatePost(
				"My message",
				types.PostID(0),
				false,
				"desmos",
				map[string]string{
					"key1":  "value1",
					"key2":  "value2",
					"key3":  "value3",
					"key4":  "value4",
					"key5":  "value5",
					"key6":  "value6",
					"key7":  "value7",
					"key8":  "value8",
					"key9":  "value9",
					"key10": "value10",
					"key11": "value11",
				},
				creator,
				date,
				msgCreatePost.Medias,
			),
			error: sdk.ErrUnknownRequest("Post optional data cannot be longer than 10 fields"),
		},
		{
			name: "Optional data longer than 200 characters returns error",
			msg: types.NewMsgCreatePost(
				"My message",
				types.PostID(0),
				false,
				"desmos",
				map[string]string{
					"key1": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi ac ullamcorper dui, a mattis sapien. Vivamus sed massa eget felis hendrerit ultrices. Morbi pretium hendrerit nisi quis faucibus volutpat.",
				},
				creator,
				date,
				msgCreatePost.Medias,
			),
			error: sdk.ErrUnknownRequest("Post optional data value lengths cannot be longer than 200. key1 exceeds the limit"),
		},
		{
			name: "Future creation date returns error",
			msg: types.NewMsgCreatePost(
				"future post",
				types.PostID(0),
				false,
				"desmos",
				map[string]string{},
				creator,
				time.Now().UTC().Add(time.Hour),
				msgCreatePost.Medias,
			),
			error: sdk.ErrUnknownRequest("Creation date cannot be in the future"),
		},
		{
			name: "Empty URI in medias returns error",
			msg: types.NewMsgCreatePost(
				"future post",
				types.PostID(0),
				false,
				"desmos",
				map[string]string{},
				creator,
				date,
				types.PostMedias{
					types.PostMedia{
						URI:      "",
						MimeType: "text/plain",
					},
				},
			),
			error: sdk.ErrUnknownRequest("uri must be specified and cannot be empty"),
		},
		{
			name: "Invalid URI in message returns error",
			msg: types.NewMsgCreatePost(
				"My message",
				types.PostID(0),
				false,
				"desmos",
				map[string]string{},
				creator,
				date,
				types.PostMedias{types.PostMedia{
					URI:      "invalid-uri",
					MimeType: "text/plain",
				}}),
			error: sdk.ErrUnknownRequest("invalid uri provided"),
		},
		{
			name: "Empty mime type in message returns error",
			msg: types.NewMsgCreatePost(
				"My message",
				types.PostID(0),
				false,
				"desmos",
				map[string]string{},
				creator,
				date,
				types.PostMedias{types.PostMedia{
					URI:      "https://example.com",
					MimeType: "",
				}}),
			error: sdk.ErrUnknownRequest("mime type must be specified and cannot be empty"),
		},
		{
			name: "Valid message does not return any error",
			msg: types.NewMsgCreatePost(
				"Message",
				types.PostID(0),
				false,
				"desmos",
				map[string]string{
					"lorem":  "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras in dapibus tortor, in iaculis nunc. Integer ac bibendum nisi. Curabitur faucibus vestibulum tincidunt. Donec interdum tincidunt cras amet.",
					"date":   "2020-01-01T00:00.000Z",
					"text":   "Welcome to Desmos",
					"int":    "0",
					"json":   `{"key":"value"}`,
					"double": "12.0",
					"array":  `["first","second"]`,
				},
				creator,
				date,
				types.PostMedias{
					types.PostMedia{
						URI:      "https://uri.com",
						MimeType: "text/plain",
					},
				},
			),
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
	tests := []struct {
		name        string
		msg         types.MsgCreatePost
		expSignJSON string
	}{
		{
			name: "Message with non-empty external reference",
			msg: types.NewMsgCreatePost(
				"My new post",
				types.PostID(53),
				false,
				"desmos",
				map[string]string{"field": "value"},
				testOwner,
				date,
				types.PostMedias{
					types.PostMedia{
						URI:      "https://uri.com",
						MimeType: "text/plain",
					},
				},
			),
			expSignJSON: `{"type":"desmos/MsgCreatePost","value":{"allows_comments":false,"creation_date":"2020-01-01T12:00:00Z","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","message":"My new post","optional_data":{"field":"value"},"parent_id":"53","post_medias":[{"mime_Type":"text/plain","uri":"https://uri.com"}],"subspace":"desmos"}}`,
		},
		{
			name: "Message with non-empty external reference",
			msg: types.NewMsgCreatePost(
				"My post",
				types.PostID(15),
				false,
				"desmos",
				map[string]string{},
				testOwner,
				date,
				types.PostMedias{
					types.PostMedia{
						URI:      "https://uri.com",
						MimeType: "text/plain",
					},
				},
			),
			expSignJSON: `{"type":"desmos/MsgCreatePost","value":{"allows_comments":false,"creation_date":"2020-01-01T12:00:00Z","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","message":"My post","parent_id":"15","post_medias":[{"mime_Type":"text/plain","uri":"https://uri.com"}],"subspace":"desmos"}}`,
		},
		{
			name: "Message with empty medias",
			msg: types.NewMsgCreatePost(
				"My Post without medias",
				types.PostID(10),
				false,
				"desmos",
				map[string]string{},
				testOwner,
				date,
				types.PostMedias{},
			),
			expSignJSON: `{"type":"desmos/MsgCreatePost","value":{"allows_comments":false,"creation_date":"2020-01-01T12:00:00Z","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","message":"My Post without medias","parent_id":"10","subspace":"desmos"}}`,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expSignJSON, string(test.msg.GetSignBytes()))
		})
	}
}

func TestMsgCreatePost_GetSigners(t *testing.T) {
	actual := msgCreatePost.GetSigners()
	assert.Equal(t, 1, len(actual))
	assert.Equal(t, msgCreatePost.Creator, actual[0])
}

// ----------------------
// --- MsgEditPost
// ----------------------

var editDate = time.Date(2010, 1, 1, 15, 0, 0, 0, timeZone)
var msgEditPost = types.NewMsgEditPost(types.PostID(94), "Edited post message", testOwner, editDate)

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
			msg:   types.NewMsgEditPost(types.PostID(0), "Edited post message", testOwner, editDate),
			error: sdk.ErrUnknownRequest("Invalid post id"),
		},
		{
			name:  "Invalid editor returns error",
			msg:   types.NewMsgEditPost(types.PostID(10), "Edited post message", nil, editDate),
			error: sdk.ErrInvalidAddress("Invalid editor address: "),
		},
		{
			name:  "Blank message returns error",
			msg:   types.NewMsgEditPost(types.PostID(10), " ", testOwner, editDate),
			error: sdk.ErrUnknownRequest("Post message cannot be empty nor blank"),
		},
		{
			name:  "Empty message returns error",
			msg:   types.NewMsgEditPost(types.PostID(10), "", testOwner, editDate),
			error: sdk.ErrUnknownRequest("Post message cannot be empty nor blank"),
		},
		{
			name:  "Empty edit date returns error",
			msg:   types.NewMsgEditPost(types.PostID(10), "My new message", testOwner, time.Time{}),
			error: sdk.ErrUnknownRequest("Invalid edit date"),
		},
		{
			name:  "Future edit date returns error",
			msg:   types.NewMsgEditPost(types.PostID(10), "My new message", testOwner, time.Now().Add(time.Hour)),
			error: sdk.ErrUnknownRequest("Edit date cannot be in the future"),
		},
		{
			name:  "Valid message returns no error",
			msg:   types.NewMsgEditPost(types.PostID(10), "Edited post message", testOwner, editDate),
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
	expected := `{"type":"desmos/MsgEditPost","value":{"edit_date":"2010-01-01T15:00:00Z","editor":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","message":"Edited post message","post_id":"94"}}`
	assert.Equal(t, expected, string(actual))
}

func TestMsgEditPost_GetSigners(t *testing.T) {
	actual := msgEditPost.GetSigners()
	assert.Equal(t, 1, len(actual))
	assert.Equal(t, msgEditPost.Editor, actual[0])
}

// ----------------------
// --- MsgAddPostReaction
// ----------------------

var msgLike = types.NewMsgAddPostReaction(types.PostID(94), "like", testOwner)

func TestMsgAddPostReaction_Route(t *testing.T) {
	actual := msgLike.Route()
	assert.Equal(t, "posts", actual)
}

func TestMsgAddPostReaction_Type(t *testing.T) {
	actual := msgLike.Type()
	assert.Equal(t, "add_post_reaction", actual)
}

func TestMsgAddPostReaction_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   types.MsgAddPostReaction
		error sdk.Error
	}{
		{
			name:  "Invalid post id returns error",
			msg:   types.NewMsgAddPostReaction(types.PostID(0), "like", testOwner),
			error: sdk.ErrUnknownRequest("Invalid post id"),
		},
		{
			name:  "Invalid user returns error",
			msg:   types.NewMsgAddPostReaction(types.PostID(5), "like", nil),
			error: sdk.ErrInvalidAddress("Invalid user address: "),
		},
		{
			name:  "Invalid value returns error",
			msg:   types.NewMsgAddPostReaction(types.PostID(5), "", testOwner),
			error: sdk.ErrUnknownRequest("Reaction value cannot be empty nor blank"),
		},
		{
			name:  "Valid message returns no error",
			msg:   types.NewMsgAddPostReaction(types.PostID(10), "like", testOwner),
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

func TestMsgAddPostReaction_GetSignBytes(t *testing.T) {
	actual := msgLike.GetSignBytes()
	expected := `{"type":"desmos/MsgAddPostReaction","value":{"post_id":"94","user":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","value":"like"}}`
	assert.Equal(t, expected, string(actual))
}

func TestMsgAddPostReaction_GetSigners(t *testing.T) {
	actual := msgLike.GetSigners()
	assert.Equal(t, 1, len(actual))
	assert.Equal(t, msgLike.User, actual[0])
}

// ----------------------
// --- MsgRemovePostReaction
// ----------------------

var msgUnlikePost = types.NewMsgRemovePostReaction(types.PostID(94), testOwner, "like")

func TestMsgUnlikePost_Route(t *testing.T) {
	actual := msgUnlikePost.Route()
	assert.Equal(t, "posts", actual)
}

func TestMsgUnlikePost_Type(t *testing.T) {
	actual := msgUnlikePost.Type()
	assert.Equal(t, "remove_post_reaction", actual)
}

func TestMsgUnlikePost_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   types.MsgRemovePostReaction
		error sdk.Error
	}{
		{
			name:  "Invalid post id returns error",
			msg:   types.NewMsgRemovePostReaction(types.PostID(0), testOwner, "like"),
			error: sdk.ErrUnknownRequest("Invalid post id"),
		},
		{
			name:  "Invalid user address: ",
			msg:   types.NewMsgRemovePostReaction(types.PostID(10), nil, "like"),
			error: sdk.ErrInvalidAddress("Invalid user address: "),
		},
		{
			name:  "Invalid value returns no error",
			msg:   types.NewMsgRemovePostReaction(types.PostID(10), testOwner, ""),
			error: sdk.ErrUnknownRequest("Reaction value cannot be empty nor blank"),
		},
		{
			name:  "Valid message returns no error",
			msg:   types.NewMsgRemovePostReaction(types.PostID(10), testOwner, "like"),
			error: nil,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.error, test.msg.ValidateBasic())
	}
}

func TestMsgUnlikePost_GetSignBytes(t *testing.T) {
	actual := msgUnlikePost.GetSignBytes()
	expected := `{"type":"desmos/MsgRemovePostReaction","value":{"post_id":"94","reaction":"like","user":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	assert.Equal(t, expected, string(actual))
}

func TestMsgUnlikePost_GetSigners(t *testing.T) {
	actual := msgUnlikePost.GetSigners()
	assert.Equal(t, 1, len(actual))
	assert.Equal(t, msgUnlikePost.User, actual[0])
}
