package types_test

import (
	"github.com/desmos-labs/desmos/x/posts/types"
	"testing"
	"time"

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
var pollData = NewPollData(
	"poll?",
	time.Date(2050, 1, 1, 15, 15, 00, 000, timeZone),
	models.NewPollAnswers(
		NewPollAnswer(models.AnswerID(1), "Yes"),
		NewPollAnswer(models.AnswerID(2), "No"),
	),
	false,
	true,
)
var attachments = models.NewAttachments(NewAttachment("https://uri.com", "text/plain", nil))

var id = PostID("dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1")
var msgCreatePost = types.NewMsgCreatePost(
	"My new post",
	id,
	false,
	"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	nil,
	testOwner,
	attachments,
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

	invalidPollData := NewPollData("", msgCreatePost.PollData.EndDate,
		msgCreatePost.PollData.ProvidedAnswers, true, true)

	tests := []struct {
		name  string
		msg   types.MsgCreatePost
		error error
	}{
		{
			name: "Empty owner returns error",
			msg: types.NewMsgCreatePost(
				"Message",
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				nil,
				msgCreatePost.Attachments,
				msgCreatePost.PollData,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator address: "),
		},
		{
			name: "Empty message returns error if attachments, poll data and message are empty",
			msg: types.NewMsgCreatePost(
				"",
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				creator,
				nil,
				nil,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post message, attachments or poll are required and cannot be all blank or empty"),
		},
		{
			name: "Non-empty message returns no error if attachments are empty",
			msg: types.NewMsgCreatePost(
				"message",
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				creator,
				nil,
				msgCreatePost.PollData,
			),
			error: nil,
		},
		{
			name: "Non-empty message returns no error if attachments aren't empty",
			msg: types.NewMsgCreatePost(
				"message",
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				creator,
				msgCreatePost.Attachments,
				msgCreatePost.PollData,
			),
			error: nil,
		},
		{
			name: "Empty message returns no error if poll isn't empty",
			msg: types.NewMsgCreatePost(
				"",
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				creator,
				nil,
				msgCreatePost.PollData,
			),
			error: nil,
		},
		{
			name: "Non-empty message returns no error if poll is empty",
			msg: types.NewMsgCreatePost(
				"message",
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				creator,
				nil,
				nil,
			),
			error: nil,
		},
		{
			name: "Empty subspace returns error",
			msg: types.NewMsgCreatePost(
				"My message",
				"",
				false,
				"",
				nil,
				creator,
				msgCreatePost.Attachments,
				msgCreatePost.PollData,
			),
			error: sdkerrors.Wrap(ErrInvalidSubspace, "post subspace must be a valid sha-256 hash"),
		},
		{
			name: "Empty URI in medias returns error",
			msg: types.NewMsgCreatePost(
				"future post",
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				creator,
				models.Attachments{
					Attachment{
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
			msg: types.NewMsgCreatePost(
				"My message",
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				creator,
				models.Attachments{Attachment{
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
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				creator,
				models.Attachments{
					Attachment{
						URI:      "https://example.com",
						MimeType: "",
					},
				},
				msgCreatePost.PollData,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "mime type must be specified and cannot be empty"),
		},
		{
			name: "Message with invalid pollData returns error",
			msg: types.NewMsgCreatePost(
				"My message",
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				creator,
				nil,
				&invalidPollData,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "missing poll title"),
		},
		{
			name: "Valid message does not return any error",
			msg: types.NewMsgCreatePost(
				"Message",
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				[]OptionalDataEntry{
					{"lorem", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras in dapibus tortor, in iaculis nunc. Integer ac bibendum nisi. Curabitur faucibus vestibulum tincidunt. Donec interdum tincidunt cras amet."},
					{"date", "2020-01-01T00:00.000Z"},
					{"text", "Welcome to Desmos"},
					{"int", "0"},
					{"json", `{"key":"value"}`},
					{"double", "12.0"},
					{"array", `["first","second"]`},
				},
				creator,
				models.Attachments{
					Attachment{
						URI:      "https://uri.com",
						MimeType: "text/plain",
					},
				},
				msgCreatePost.PollData,
			),
			error: nil,
		},
		{
			name: "Message with empty attachments non-empty poll and non-empty message returns no error",
			msg: types.NewMsgCreatePost(
				"My message",
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				creator,
				nil,
				msgCreatePost.PollData,
			),
			error: nil,
		},
		{
			name: "Message with non-empty attachments, non-empty poll and non-empty message returns no error",
			msg: types.NewMsgCreatePost(
				"My message",
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				creator,
				msgCreatePost.Attachments,
				msgCreatePost.PollData,
			),
			error: nil,
		},
		{
			name: "Message with non-empty attachments, non empty poll and empty message returns no error",
			msg: types.NewMsgCreatePost(
				"",
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				creator,
				msgCreatePost.Attachments,
				msgCreatePost.PollData,
			),
			error: nil,
		},
		{
			name: "Message with empty attachments, non empty poll and empty message returns no error",
			msg: types.NewMsgCreatePost(
				"",
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				creator,
				nil,
				msgCreatePost.PollData,
			),
			error: nil,
		},
		{
			name: "Message with empty poll, non-empty attachments and non empty message returns no error",
			msg: types.NewMsgCreatePost(
				"My message",
				"",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				creator,
				models.Attachments{
					Attachment{
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
				id,
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				[]OptionalDataEntry{{"field", "value"}},
				testOwner,
				models.Attachments{
					Attachment{
						URI:      "https://uri.com",
						MimeType: "text/plain",
					},
				},
				msgCreatePost.PollData,
			),
			expSignJSON: `{"type":"desmos/MsgCreatePost","value":{"allows_comments":false,"attachments":[{"mime_type":"text/plain","uri":"https://uri.com"}],"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","message":"My new post","optional_data":[{"key":"field","value":"value"}],"parent_id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","poll_data":{"allows_answer_edits":true,"allows_multiple_answers":false,"end_date":"2050-01-01T15:15:00Z","provided_answers":[{"id":"1","text":"Yes"},{"id":"2","text":"No"}],"question":"poll?"},"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`,
		},
		{
			name: "Message with empty external reference",
			msg: types.NewMsgCreatePost(
				"My post",
				id,
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				testOwner,
				models.Attachments{
					Attachment{
						URI:      "https://uri.com",
						MimeType: "text/plain",
					},
				},
				msgCreatePost.PollData,
			),
			expSignJSON: `{"type":"desmos/MsgCreatePost","value":{"allows_comments":false,"attachments":[{"mime_type":"text/plain","uri":"https://uri.com"}],"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","message":"My post","parent_id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","poll_data":{"allows_answer_edits":true,"allows_multiple_answers":false,"end_date":"2050-01-01T15:15:00Z","provided_answers":[{"id":"1","text":"Yes"},{"id":"2","text":"No"}],"question":"poll?"},"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`,
		},
		{
			name: "Message with empty attachments",
			msg: types.NewMsgCreatePost(
				"My Post without attachments",
				id,
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				testOwner,
				models.Attachments{},
				msgCreatePost.PollData,
			),
			expSignJSON: `{"type":"desmos/MsgCreatePost","value":{"allows_comments":false,"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","message":"My Post without attachments","parent_id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","poll_data":{"allows_answer_edits":true,"allows_multiple_answers":false,"end_date":"2050-01-01T15:15:00Z","provided_answers":[{"id":"1","text":"Yes"},{"id":"2","text":"No"}],"question":"poll?"},"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`,
		},
		{
			name: "Message with empty poll data",
			msg: types.NewMsgCreatePost(
				"My Post without attachments",
				id,
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				testOwner,
				models.Attachments{
					Attachment{
						URI:      "https://uri.com",
						MimeType: "text/plain",
					},
				},
				nil,
			),
			expSignJSON: `{"type":"desmos/MsgCreatePost","value":{"allows_comments":false,"attachments":[{"mime_type":"text/plain","uri":"https://uri.com"}],"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","message":"My Post without attachments","parent_id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`,
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

func TestMsgCreatePost_ReadJSON(t *testing.T) {
	json := `{"type":"desmos/MsgCreatePost","value":{"parent_id":"","message":"","allows_comments":true,"subspace":"2bdf5932925584b9a86470bea60adce69041608a447f84a3317723aa5678ec88","optional_data":[{"key":"local_id","value":"2020-09-15T10:17:54.101972"}],"creator":"cosmos10txl52f64zmp2j7eywawlv9t4xxc4e0wnjlhq9","poll_data":{"question":"What is it better?","end_date":"2020-10-15T08:17:45.639Z","is_open":true,"allows_multiple_answers":false,"allows_answer_edits":false,"provided_answers":[{"id":"0","text":"Sushi\t"},{"id":"1","text":"Pizza"}]}}}`

	var msg types.MsgCreatePost
	err := msgs.MsgsCodec.UnmarshalJSON([]byte(json), &msg)
	require.NoError(t, err)
}

// ___________________________________________________________________________________________________________________

var msgEditPost = types.NewMsgEditPost(id, "Edited post message", attachments, &pollData, testOwner)

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
			msg:   types.NewMsgEditPost("", "Edited post message", attachments, &pollData, testOwner),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid post id: "),
		},
		{
			name:  "Invalid editor returns error",
			msg:   types.NewMsgEditPost(id, "Edited post message", attachments, &pollData, nil),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid editor address: "),
		},
		{
			name: "Non-empty message returns no error if attachments are empty",
			msg: types.NewMsgEditPost(
				id,
				"message",
				nil,
				nil,
				testOwner,
			),
			error: nil,
		},
		{
			name: "Non-empty message returns no error if attachments aren't empty",
			msg: types.NewMsgEditPost(
				id,
				"message",
				msgCreatePost.Attachments,
				nil,
				testOwner,
			),
			error: nil,
		},
		{
			name: "Empty message returns no error if poll isn't empty",
			msg: types.NewMsgEditPost(
				id,
				"",
				nil,
				msgCreatePost.PollData,
				testOwner,
			),
			error: nil,
		},
		{
			name: "Non-empty message returns no error if poll is empty",
			msg: types.NewMsgEditPost(
				id,
				"message",
				nil,
				nil,
				testOwner,
			),
			error: nil,
		},
		{
			name: "Empty message returns error if message, attachments and poll are empty",
			msg: types.NewMsgEditPost(
				id,
				"",
				nil,
				nil,
				testOwner,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"post message, attachments or poll are required and cannot be all blank or empty"),
		},
		{
			name: "Empty URI in medias returns error",
			msg: types.NewMsgEditPost(
				id,
				"future post",
				models.Attachments{
					Attachment{
						URI:      "",
						MimeType: "text/plain",
					},
				},
				msgCreatePost.PollData,
				testOwner,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid uri provided"),
		},
		{
			name: "Invalid URI in message returns error",
			msg: types.NewMsgEditPost(
				id,
				"My message",
				models.Attachments{Attachment{
					URI:      "invalid-uri",
					MimeType: "text/plain",
				}},
				msgCreatePost.PollData,
				testOwner,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid uri provided"),
		},
		{
			name: "Empty mime type in message returns error",
			msg: types.NewMsgEditPost(
				id,
				"My message",
				models.Attachments{
					Attachment{
						URI:      "https://example.com",
						MimeType: "",
					},
				},
				msgCreatePost.PollData,
				testOwner,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "mime type must be specified and cannot be empty"),
		},
		{
			name: "Invalid pollData returns error",
			msg: types.NewMsgEditPost(
				id,
				"My message",
				attachments,
				&PollData{
					Question: "",
					ProvidedAnswers: models.NewPollAnswers(
						NewPollAnswer(models.AnswerID(1), "Yes"),
						NewPollAnswer(models.AnswerID(2), "No"),
					),
					EndDate:           time.Date(2050, 1, 1, 15, 15, 00, 000, timeZone),
					AllowsAnswerEdits: true,
				},
				testOwner,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "missing poll title"),
		},
		{
			name:  "Valid message returns no error",
			msg:   types.NewMsgEditPost(id, "Edited post message", attachments, &pollData, testOwner),
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
	expected := `{"type":"desmos/MsgEditPost","value":{"attachments":[{"mime_type":"text/plain","uri":"https://uri.com"}],"editor":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","message":"Edited post message","poll_data":{"allows_answer_edits":true,"allows_multiple_answers":false,"end_date":"2050-01-01T15:15:00Z","provided_answers":[{"id":"1","text":"Yes"},{"id":"2","text":"No"}],"question":"poll?"},"post_id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgEditPost_GetSigners(t *testing.T) {
	actual := msgEditPost.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgEditPost.Editor, actual[0])
}

// ___________________________________________________________________________________________________________________

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
			msg:   types.NewMsgAddPostReaction("", ":like:", testOwner),
			error: sdkerrors.Wrap(ErrInvalidPostID, ""),
		},
		{
			name:  "Invalid user returns error",
			msg:   types.NewMsgAddPostReaction(id, ":like:", nil),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid user address: "),
		},
		{
			name:  "Invalid value returns error",
			msg:   types.NewMsgAddPostReaction(id, "like", testOwner),
			error: sdkerrors.Wrap(ErrInvalidReactionCode, "like"),
		},
		{
			name:  "Valid message returns no error (with shortcode)",
			msg:   types.NewMsgAddPostReaction(id, ":like:", testOwner),
			error: nil,
		},
		{
			name:  "Valid message returns no error (with emoji)",
			msg:   types.NewMsgAddPostReaction(id, "🤩", testOwner),
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

// ___________________________________________________________________________________________________________________

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
			msg:   types.NewMsgRemovePostReaction("", testOwner, ":+1:"),
			error: sdkerrors.Wrap(ErrInvalidPostID, ""),
		},
		{
			name:  "Invalid user address: ",
			msg:   types.NewMsgRemovePostReaction(id, nil, ":like:"),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid user address: "),
		},
		{
			name:  "Blank value returns no error",
			msg:   types.NewMsgRemovePostReaction(id, testOwner, ""),
			error: sdkerrors.Wrap(ErrInvalidReactionCode, ""),
		},
		{
			name:  "Valid message returns no error (with shortcode)",
			msg:   types.NewMsgRemovePostReaction(id, testOwner, ":+1:"),
			error: nil,
		},
		{
			name:  "Valid message returns no error (with emoji)",
			msg:   types.NewMsgRemovePostReaction(id, testOwner, "🤩"),
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

// ___________________________________________________________________________________________________________________

var msgAnswerPollPost = types.NewMsgAnswerPoll(id, []models.AnswerID{1, 2}, testOwner)

func TestMsgAnswerPollPost_Route(t *testing.T) {
	actual := msgAnswerPollPost.Route()
	require.Equal(t, "posts", actual)
}

func TestMsgAnswerPollPost_Type(t *testing.T) {
	actual := msgAnswerPollPost.Type()
	require.Equal(t, "answer_poll", actual)
}

func TestMsgAnswerPollPost_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   types.MsgAnswerPoll
		error error
	}{
		{
			name:  "Invalid post id",
			msg:   types.NewMsgAnswerPoll("", []models.AnswerID{1, 2}, msgAnswerPollPost.Answerer),
			error: sdkerrors.Wrap(ErrInvalidPostID, ""),
		},
		{
			name:  "Invalid answerer address",
			msg:   types.NewMsgAnswerPoll(id, []models.AnswerID{1, 2}, nil),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid answerer address: "),
		},
		{
			name:  "Returns error when no answer is provided",
			msg:   types.NewMsgAnswerPoll(id, []models.AnswerID{}, msgAnswerPollPost.Answerer),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "provided answers must contains at least one answer"),
		},
		{
			name: "Valid message returns no error",
			msg:  types.NewMsgAnswerPoll(id, []models.AnswerID{1, 2}, msgAnswerPollPost.Answerer),
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

func TestMsgAnswerPollPost_GetSignBytes(t *testing.T) {
	actual := msgAnswerPollPost.GetSignBytes()
	expected := `{"type":"desmos/MsgAnswerPoll","value":{"answerer":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","answers":["1","2"],"post_id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgAnswerPollPost_GetSigners(t *testing.T) {
	actual := msgAnswerPollPost.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgAnswerPollPost.Answerer, actual[0])
}

// ___________________________________________________________________________________________________________________

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
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator address: "),
		},
		{
			name: "Empty short code returns error",
			msg: types.NewMsgRegisterReaction(testOwner, "", "https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			error: sdkerrors.Wrap(ErrInvalidReactionCode, ""),
		},
		{
			name: "Invalid short code returns error",
			msg: types.NewMsgRegisterReaction(testOwner, ":smile", "https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			error: sdkerrors.Wrap(ErrInvalidReactionCode, ":smile"),
		},
		{
			name: "Empty value returns error",
			msg: types.NewMsgRegisterReaction(testOwner, ":smile:", "",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			error: sdkerrors.Wrap(commonerrors.ErrInvalidURI, "reaction value should be a valid uri"),
		},
		{
			name: "Invalid value returns error (url)",
			msg: types.NewMsgRegisterReaction(testOwner, ":smile:", "htp://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			error: sdkerrors.Wrap(commonerrors.ErrInvalidURI, "reaction value should be a valid uri"),
		},
		{
			name: "Invalid value returns error (unicode)",
			msg: types.NewMsgRegisterReaction(testOwner, ":smile:", "U+1",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			error: sdkerrors.Wrap(commonerrors.ErrInvalidURI, "reaction value should be a valid uri"),
		},
		{
			name:  "Valid emoji value returns no error",
			msg:   types.NewMsgRegisterReaction(testOwner, ":smile:", "💙", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			error: sdkerrors.Wrap(commonerrors.ErrInvalidURI, "reaction value should be a valid uri"),
		},
		{
			name: "Invalid subspace returns error",
			msg: types.NewMsgRegisterReaction(testOwner, ":smile:", "https://smile.jpg",
				"1234"),
			error: sdkerrors.Wrap(ErrInvalidSubspace, "reaction subspace must be a valid sha-256 hash"),
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
