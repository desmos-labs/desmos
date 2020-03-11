package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/require"
)

// ----------------------
// --- MsgCreatePost
// ----------------------

var testOwner, _ = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
var timeZone, _ = time.LoadLocation("UTC")
var date = time.Date(2020, 1, 1, 12, 0, 0, 0, timeZone)
var pollData = types.NewPollData(
	"poll?",
	time.Date(2050, 1, 1, 15, 15, 00, 000, timeZone),
	types.NewPollAnswers(
		types.NewPollAnswer(types.AnswerID(1), "Yes"),
		types.NewPollAnswer(types.AnswerID(2), "No"),
	),
	true,
	false,
	true,
)
var msgCreatePost = types.NewMsgCreatePost(
	"My new post",
	types.PostID(53),
	false,
	"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	map[string]string{},
	testOwner,
	date,
	types.NewPostMedias(types.NewPostMedia("https://uri.com", "text/plain")),
	&pollData,
)

func TestMsgCreatePost_Route(t *testing.T) {
	actual := msgCreatePost.Route()
	require.Equal(t, "posts", actual)
}

func TestMsgCreatePost_Type(t *testing.T) {
	actual := msgCreatePost.Type()
	require.Equal(t, "create_post", actual)
}

func TestMsgCreatePost_ValidateBasic(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h")
	require.NoError(t, err)

	tests := []struct {
		name  string
		msg   types.MsgCreatePost
		error error
	}{
		{
			name: "Empty owner returns error",
			msg: types.NewMsgCreatePost(
				"Message",
				types.PostID(0),
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				nil,
				date,
				msgCreatePost.Medias,
				msgCreatePost.PollData,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid creator address: "),
		},
		{
			name: "Empty message returns error if medias and message are empty",
			msg: types.NewMsgCreatePost(
				"",
				types.PostID(0),
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				creator,
				date,
				nil,
				msgCreatePost.PollData,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Post message or medias are required and cannot be both blank or empty"),
		},
		{
			name: "Non-empty message returns no error if medias are empty",
			msg: types.NewMsgCreatePost(
				"message",
				types.PostID(0),
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				creator,
				date,
				nil,
				msgCreatePost.PollData,
			),
			error: nil,
		},
		{
			name: "Non-empty message returns no error if medias aren't empty",
			msg: types.NewMsgCreatePost(
				"message",
				types.PostID(0),
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				creator,
				date,
				msgCreatePost.Medias,
				msgCreatePost.PollData,
			),
			error: nil,
		},
		{
			name: "Empty message returns no error if medias aren't empty",
			msg: types.NewMsgCreatePost(
				"",
				types.PostID(0),
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				creator,
				date,
				msgCreatePost.Medias,
				msgCreatePost.PollData,
			),
			error: nil,
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
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				creator,
				date,
				msgCreatePost.Medias,
				msgCreatePost.PollData,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Post message cannot exceed 500 characters"),
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
				msgCreatePost.PollData,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Post subspace must be a valid sha-256 hash"),
		},
		{
			name: "More than 10 optional data returns error",
			msg: types.NewMsgCreatePost(
				"My message",
				types.PostID(0),
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
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
				msgCreatePost.PollData,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Post optional data cannot be longer than 10 fields"),
		},
		{
			name: "Optional data longer than 200 characters returns error",
			msg: types.NewMsgCreatePost(
				"My message",
				types.PostID(0),
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{
					"key1": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi ac ullamcorper dui, a mattis sapien. Vivamus sed massa eget felis hendrerit ultrices. Morbi pretium hendrerit nisi quis faucibus volutpat.",
				},
				creator,
				date,
				msgCreatePost.Medias,
				msgCreatePost.PollData,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Post optional data value lengths cannot be longer than 200. key1 exceeds the limit"),
		},
		{
			name: "Future creation date returns error",
			msg: types.NewMsgCreatePost(
				"future post",
				types.PostID(0),
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				creator,
				time.Now().UTC().Add(time.Hour),
				msgCreatePost.Medias,
				msgCreatePost.PollData,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Creation date cannot be in the future"),
		},
		{
			name: "Empty URI in medias returns error",
			msg: types.NewMsgCreatePost(
				"future post",
				types.PostID(0),
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				creator,
				date,
				types.PostMedias{
					types.PostMedia{
						URI:      "",
						MimeType: "text/plain",
					},
				},
				msgCreatePost.PollData,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "uri must be specified and cannot be empty"),
		},
		{
			name: "Invalid URI in message returns error",
			msg: types.NewMsgCreatePost(
				"My message",
				types.PostID(0),
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				creator,
				date,
				types.PostMedias{types.PostMedia{
					URI:      "invalid-uri",
					MimeType: "text/plain",
				}},
				msgCreatePost.PollData,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid uri provided"),
		},
		{
			name: "Empty mime type in message returns error",
			msg: types.NewMsgCreatePost(
				"My message",
				types.PostID(0),
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				creator,
				date,
				types.PostMedias{
					types.PostMedia{
						URI:      "https://example.com",
						MimeType: "",
					},
				},
				msgCreatePost.PollData,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "mime type must be specified and cannot be empty"),
		},
		{
			name: "Valid message does not return any error",
			msg: types.NewMsgCreatePost(
				"Message",
				types.PostID(0),
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
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
				msgCreatePost.PollData,
			),
			error: nil,
		},
		{
			name: "Message with empty medias returns no error",
			msg: types.NewMsgCreatePost(
				"My message",
				types.PostID(0),
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				creator,
				date,
				nil,
				msgCreatePost.PollData,
			),
			error: nil,
		},
		{
			name: "Message with empty poll returns no error",
			msg: types.NewMsgCreatePost(
				"My message",
				types.PostID(0),
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				creator,
				date,
				types.PostMedias{
					types.PostMedia{
						URI:      "https://example.com",
						MimeType: "text/plain",
					},
				},
				nil,
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

	err = msgCreatePost.ValidateBasic()
	require.Nil(t, err)
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
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{"field": "value"},
				testOwner,
				date,
				types.PostMedias{
					types.PostMedia{
						URI:      "https://uri.com",
						MimeType: "text/plain",
					},
				},
				msgCreatePost.PollData,
			),
			expSignJSON: `{"type":"desmos/MsgCreatePost","value":{"allows_comments":false,"creation_date":"2020-01-01T12:00:00Z","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","medias":[{"mime_type":"text/plain","uri":"https://uri.com"}],"message":"My new post","optional_data":{"field":"value"},"parent_id":"53","poll_data":{"allows_answer_edits":true,"allows_multiple_answers":false,"end_date":"2050-01-01T15:15:00Z","is_open":true,"provided_answers":[{"id":"1","text":"Yes"},{"id":"2","text":"No"}],"question":"poll?"},"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`,
		},
		{
			name: "Message with empty external reference",
			msg: types.NewMsgCreatePost(
				"My post",
				types.PostID(15),
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testOwner,
				date,
				types.PostMedias{
					types.PostMedia{
						URI:      "https://uri.com",
						MimeType: "text/plain",
					},
				},
				msgCreatePost.PollData,
			),
			expSignJSON: `{"type":"desmos/MsgCreatePost","value":{"allows_comments":false,"creation_date":"2020-01-01T12:00:00Z","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","medias":[{"mime_type":"text/plain","uri":"https://uri.com"}],"message":"My post","parent_id":"15","poll_data":{"allows_answer_edits":true,"allows_multiple_answers":false,"end_date":"2050-01-01T15:15:00Z","is_open":true,"provided_answers":[{"id":"1","text":"Yes"},{"id":"2","text":"No"}],"question":"poll?"},"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`,
		},
		{
			name: "Message with empty medias",
			msg: types.NewMsgCreatePost(
				"My Post without medias",
				types.PostID(10),
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testOwner,
				date,
				types.PostMedias{},
				msgCreatePost.PollData,
			),
			expSignJSON: `{"type":"desmos/MsgCreatePost","value":{"allows_comments":false,"creation_date":"2020-01-01T12:00:00Z","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","message":"My Post without medias","parent_id":"10","poll_data":{"allows_answer_edits":true,"allows_multiple_answers":false,"end_date":"2050-01-01T15:15:00Z","is_open":true,"provided_answers":[{"id":"1","text":"Yes"},{"id":"2","text":"No"}],"question":"poll?"},"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`,
		},
		{
			name: "Message with empty poll data",
			msg: types.NewMsgCreatePost(
				"My Post without medias",
				types.PostID(10),
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testOwner,
				date,
				types.PostMedias{
					types.PostMedia{
						URI:      "https://uri.com",
						MimeType: "text/plain",
					},
				},
				nil,
			),
			expSignJSON: `{"type":"desmos/MsgCreatePost","value":{"allows_comments":false,"creation_date":"2020-01-01T12:00:00Z","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","medias":[{"mime_type":"text/plain","uri":"https://uri.com"}],"message":"My Post without medias","parent_id":"10","subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expSignJSON, string(test.msg.GetSignBytes()))
		})
	}
}

func TestMsgCreatePost_GetSigners(t *testing.T) {
	actual := msgCreatePost.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgCreatePost.Creator, actual[0])
}

// ----------------------
// --- MsgEditPost
// ----------------------

var editDate = time.Date(2010, 1, 1, 15, 0, 0, 0, timeZone)
var msgEditPost = types.NewMsgEditPost(types.PostID(94), "Edited post message", testOwner, editDate)

func TestMsgEditPost_Route(t *testing.T) {
	actual := msgEditPost.Route()
	require.Equal(t, "posts", actual)
}

func TestMsgEditPost_Type(t *testing.T) {
	actual := msgEditPost.Type()
	require.Equal(t, "edit_post", actual)
}

func TestMsgEditPost_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   types.MsgEditPost
		error error
	}{
		{
			name:  "Invalid post id returns error",
			msg:   types.NewMsgEditPost(types.PostID(0), "Edited post message", testOwner, editDate),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid post id"),
		},
		{
			name:  "Invalid editor returns error",
			msg:   types.NewMsgEditPost(types.PostID(10), "Edited post message", nil, editDate),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid editor address: "),
		},
		{
			name:  "Blank message returns error",
			msg:   types.NewMsgEditPost(types.PostID(10), " ", testOwner, editDate),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Post message cannot be empty nor blank"),
		},
		{
			name:  "Empty message returns error",
			msg:   types.NewMsgEditPost(types.PostID(10), "", testOwner, editDate),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Post message cannot be empty nor blank"),
		},
		{
			name:  "Empty edit date returns error",
			msg:   types.NewMsgEditPost(types.PostID(10), "My new message", testOwner, time.Time{}),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid edit date"),
		},
		{
			name:  "Future edit date returns error",
			msg:   types.NewMsgEditPost(types.PostID(10), "My new message", testOwner, time.Now().Add(time.Hour)),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Edit date cannot be in the future"),
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

func TestMsgEditPost_GetSignBytes(t *testing.T) {
	actual := msgEditPost.GetSignBytes()
	expected := `{"type":"desmos/MsgEditPost","value":{"edit_date":"2010-01-01T15:00:00Z","editor":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","message":"Edited post message","post_id":"94"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgEditPost_GetSigners(t *testing.T) {
	actual := msgEditPost.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgEditPost.Editor, actual[0])
}
