package msgs_test

import (
	"testing"
	"time"

	postserrors "github.com/desmos-labs/desmos/x/posts/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/posts/types/models"
	"github.com/desmos-labs/desmos/x/posts/types/msgs"
)

// ----------------------
// --- MsgCreatePost
// ----------------------

var testOwner, _ = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
var timeZone, _ = time.LoadLocation("UTC")
var date = time.Date(2020, 1, 1, 12, 0, 0, 0, timeZone)
var pollData = models.NewPollData(
	"poll?",
	time.Date(2050, 1, 1, 15, 15, 00, 000, timeZone),
	models.NewPollAnswers(
		models.NewPollAnswer(models.AnswerID(1), "Yes"),
		models.NewPollAnswer(models.AnswerID(2), "No"),
	),
	true,
	false,
	true,
)
var id = models.PostID("dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1")
var msgCreatePost = msgs.NewMsgCreatePost(
	"My new post",
	id,
	false,
	"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	map[string]string{},
	testOwner,
	date,
	models.NewPostMedias(models.NewPostMedia("https://uri.com", "text/plain", nil)),
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
		msg   msgs.MsgCreatePost
		error error
	}{
		{
			name: "Empty owner returns error",
			msg: msgs.NewMsgCreatePost(
				"Message",
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				nil,
				date,
				msgCreatePost.Medias,
				msgCreatePost.PollData,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator address: "),
		},
		{
			name: "Empty message returns error if medias, poll data and message are empty",
			msg: msgs.NewMsgCreatePost(
				"",
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				creator,
				date,
				nil,
				nil,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post message, medias or poll are required and cannot be all blank or empty"),
		},
		{
			name: "Non-empty message returns no error if medias are empty",
			msg: msgs.NewMsgCreatePost(
				"message",
				"",
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
			msg: msgs.NewMsgCreatePost(
				"message",
				"",
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
			msg: msgs.NewMsgCreatePost(
				"",
				"",
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
			name: "Empty message returns no error if poll isn't empty",
			msg: msgs.NewMsgCreatePost(
				"",
				"",
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
			name: "Non-empty message returns no error if poll is empty",
			msg: msgs.NewMsgCreatePost(
				"message",
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				creator,
				date,
				nil,
				nil,
			),
			error: nil,
		},
		{
			name: "Empty subspace returns error",
			msg: msgs.NewMsgCreatePost(
				"My message",
				"",
				false,
				"",
				map[string]string{},
				creator,
				date,
				msgCreatePost.Medias,
				msgCreatePost.PollData,
			),
			error: sdkerrors.Wrap(postserrors.ErrInvalidSubspace, "post subspace must be a valid sha-256 hash"),
		},
		{
			name: "Future creation date returns error",
			msg: msgs.NewMsgCreatePost(
				"future post",
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				creator,
				time.Now().UTC().Add(time.Hour),
				msgCreatePost.Medias,
				msgCreatePost.PollData,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "creation date cannot be in the future"),
		},
		{
			name: "Empty URI in medias returns error",
			msg: msgs.NewMsgCreatePost(
				"future post",
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				creator,
				date,
				models.PostMedias{
					models.PostMedia{
						URI:      "",
						MimeType: "text/plain",
					},
				},
				msgCreatePost.PollData,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid uri provided"),
		},
		{
			name: "Invalid URI in message returns error",
			msg: msgs.NewMsgCreatePost(
				"My message",
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				creator,
				date,
				models.PostMedias{models.PostMedia{
					URI:      "invalid-uri",
					MimeType: "text/plain",
				}},
				msgCreatePost.PollData,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid uri provided"),
		},
		{
			name: "Empty mime type in message returns error",
			msg: msgs.NewMsgCreatePost(
				"My message",
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				creator,
				date,
				models.PostMedias{
					models.PostMedia{
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
			msg: msgs.NewMsgCreatePost(
				"Message",
				"",
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
				models.PostMedias{
					models.PostMedia{
						URI:      "https://uri.com",
						MimeType: "text/plain",
					},
				},
				msgCreatePost.PollData,
			),
			error: nil,
		},
		{
			name: "Message with empty medias non-empty poll and non-empty message returns no error",
			msg: msgs.NewMsgCreatePost(
				"My message",
				"",
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
			name: "Message with non-empty medias, non-empty poll and non-empty message returns no error",
			msg: msgs.NewMsgCreatePost(
				"My message",
				"",
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
			name: "Message with non-empty medias, non empty poll and empty message returns no error",
			msg: msgs.NewMsgCreatePost(
				"",
				"",
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
			name: "Message with empty medias, non empty poll and empty message returns no error",
			msg: msgs.NewMsgCreatePost(
				"",
				"",
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
			name: "Message with empty poll, non-empty medias and non empty message returns no error",
			msg: msgs.NewMsgCreatePost(
				"My message",
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				creator,
				date,
				models.PostMedias{
					models.PostMedia{
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
		msg         msgs.MsgCreatePost
		expSignJSON string
	}{
		{
			name: "Message with non-empty external reference",
			msg: msgs.NewMsgCreatePost(
				"My new post",
				id,
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{"field": "value"},
				testOwner,
				date,
				models.PostMedias{
					models.PostMedia{
						URI:      "https://uri.com",
						MimeType: "text/plain",
					},
				},
				msgCreatePost.PollData,
			),
			expSignJSON: `{"type":"desmos/MsgCreatePost","value":{"allows_comments":false,"creation_date":"2020-01-01T12:00:00Z","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","medias":[{"mime_type":"text/plain","uri":"https://uri.com"}],"message":"My new post","optional_data":{"field":"value"},"parent_id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","poll_data":{"allows_answer_edits":true,"allows_multiple_answers":false,"end_date":"2050-01-01T15:15:00Z","is_open":true,"provided_answers":[{"id":"1","text":"Yes"},{"id":"2","text":"No"}],"question":"poll?"},"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`,
		},
		{
			name: "Message with empty external reference",
			msg: msgs.NewMsgCreatePost(
				"My post",
				id,
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testOwner,
				date,
				models.PostMedias{
					models.PostMedia{
						URI:      "https://uri.com",
						MimeType: "text/plain",
					},
				},
				msgCreatePost.PollData,
			),
			expSignJSON: `{"type":"desmos/MsgCreatePost","value":{"allows_comments":false,"creation_date":"2020-01-01T12:00:00Z","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","medias":[{"mime_type":"text/plain","uri":"https://uri.com"}],"message":"My post","parent_id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","poll_data":{"allows_answer_edits":true,"allows_multiple_answers":false,"end_date":"2050-01-01T15:15:00Z","is_open":true,"provided_answers":[{"id":"1","text":"Yes"},{"id":"2","text":"No"}],"question":"poll?"},"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`,
		},
		{
			name: "Message with empty medias",
			msg: msgs.NewMsgCreatePost(
				"My Post without medias",
				id,
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testOwner,
				date,
				models.PostMedias{},
				msgCreatePost.PollData,
			),
			expSignJSON: `{"type":"desmos/MsgCreatePost","value":{"allows_comments":false,"creation_date":"2020-01-01T12:00:00Z","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","message":"My Post without medias","parent_id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","poll_data":{"allows_answer_edits":true,"allows_multiple_answers":false,"end_date":"2050-01-01T15:15:00Z","is_open":true,"provided_answers":[{"id":"1","text":"Yes"},{"id":"2","text":"No"}],"question":"poll?"},"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`,
		},
		{
			name: "Message with empty poll data",
			msg: msgs.NewMsgCreatePost(
				"My Post without medias",
				id,
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testOwner,
				date,
				models.PostMedias{
					models.PostMedia{
						URI:      "https://uri.com",
						MimeType: "text/plain",
					},
				},
				nil,
			),
			expSignJSON: `{"type":"desmos/MsgCreatePost","value":{"allows_comments":false,"creation_date":"2020-01-01T12:00:00Z","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","medias":[{"mime_type":"text/plain","uri":"https://uri.com"}],"message":"My Post without medias","parent_id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`,
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
var msgEditPost = msgs.NewMsgEditPost(id, "Edited post message", testOwner, editDate)

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
		msg   msgs.MsgEditPost
		error error
	}{
		{
			name:  "Invalid post id returns error",
			msg:   msgs.NewMsgEditPost("", "Edited post message", testOwner, editDate),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid post id: "),
		},
		{
			name:  "Invalid editor returns error",
			msg:   msgs.NewMsgEditPost(id, "Edited post message", nil, editDate),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid editor address: "),
		},
		{
			name:  "Blank message returns error",
			msg:   msgs.NewMsgEditPost(id, " ", testOwner, editDate),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Post message cannot be empty nor blank"),
		},
		{
			name:  "Empty message returns error",
			msg:   msgs.NewMsgEditPost(id, "", testOwner, editDate),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Post message cannot be empty nor blank"),
		},
		{
			name:  "Empty edit date returns error",
			msg:   msgs.NewMsgEditPost(id, "My new message", testOwner, time.Time{}),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid edit date"),
		},
		{
			name:  "Future edit date returns error",
			msg:   msgs.NewMsgEditPost(id, "My new message", testOwner, time.Now().Add(time.Hour)),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Edit date cannot be in the future"),
		},
		{
			name:  "Valid message returns no error",
			msg:   msgs.NewMsgEditPost(id, "Edited post message", testOwner, editDate),
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
	expected := `{"type":"desmos/MsgEditPost","value":{"edit_date":"2010-01-01T15:00:00Z","editor":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","message":"Edited post message","post_id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgEditPost_GetSigners(t *testing.T) {
	actual := msgEditPost.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgEditPost.Editor, actual[0])
}
