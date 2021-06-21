package types_test

import (
	types2 "github.com/desmos-labs/desmos/x/posts/types"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	emoji "github.com/desmos-labs/Go-Emoji-Utils"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	commonerrors "github.com/desmos-labs/desmos/x/commons/types/errors"
	"github.com/stretchr/testify/require"
)

// ----------------------
// --- MsgCreatePost
// ----------------------

var msgCreatePost = types2.NewMsgCreatePost(
	"My new post",
	"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
	types2.CommentsStateAllowed,
	"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	nil,
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	types2.NewAttachments(
		types2.NewAttachment("https://uri.com", "text/plain", nil),
	),
	types2.NewPollData(
		"poll?",
		time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
		types2.NewPollAnswers(
			types2.NewPollAnswer("1", "Yes"),
			types2.NewPollAnswer("2", "No"),
		),
		false,
		true,
	),
)

func TestMsgCreatePost_Route(t *testing.T) {
	require.Equal(t, "posts", msgCreatePost.Route())
}

func TestMsgCreatePost_Type(t *testing.T) {
	require.Equal(t, "create_post", msgCreatePost.Type())
}

func TestMsgCreatePost_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types2.MsgCreatePost
		error error
	}{
		{
			name: "Empty owner returns error",
			msg: types2.NewMsgCreatePost(
				"Message",
				"",
				types2.CommentsStateAllowed,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				"",
				msgCreatePost.Attachments,
				msgCreatePost.PollData,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator"),
		},
		{
			name: "Empty message returns error if attachments, poll and message are empty",
			msg: types2.NewMsgCreatePost(
				"",
				"",
				types2.CommentsStateAllowed,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				nil,
				nil,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post message, attachments or poll are required and cannot be all blank or empty"),
		},
		{
			name: "Non-empty message returns no error if attachments are empty",
			msg: types2.NewMsgCreatePost(
				"message",
				"",
				types2.CommentsStateAllowed,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				nil,
				msgCreatePost.PollData,
			),
			error: nil,
		},
		{
			name: "Non-empty message returns no error if attachments aren't empty",
			msg: types2.NewMsgCreatePost(
				"message",
				"",
				types2.CommentsStateAllowed,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				msgCreatePost.Attachments,
				msgCreatePost.PollData,
			),
			error: nil,
		},
		{
			name: "Empty message returns no error if poll isn't empty",
			msg: types2.NewMsgCreatePost(
				"",
				"",
				types2.CommentsStateAllowed,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				nil,
				msgCreatePost.PollData,
			),
			error: nil,
		},
		{
			name: "Non-empty message returns no error if poll is empty",
			msg: types2.NewMsgCreatePost(
				"message",
				"",
				types2.CommentsStateAllowed,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				nil,
				nil,
			),
			error: nil,
		},
		{
			name: "Empty subspace returns error",
			msg: types2.NewMsgCreatePost(
				"My message",
				"",
				types2.CommentsStateAllowed,
				"",
				nil,
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				msgCreatePost.Attachments,
				msgCreatePost.PollData,
			),
			error: sdkerrors.Wrap(types2.ErrInvalidSubspace, "post subspace must be a valid sha-256 hash"),
		},
		{
			name: "Empty URI in medias returns error",
			msg: types2.NewMsgCreatePost(
				"future post",
				"",
				types2.CommentsStateAllowed,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				types2.Attachments{
					types2.NewAttachment("", "text/plain", nil),
				},
				msgCreatePost.PollData,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid uri provided"),
		},
		{
			name: "Invalid URI in message returns error",
			msg: types2.NewMsgCreatePost(
				"My message",
				"",
				types2.CommentsStateAllowed,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				types2.Attachments{
					types2.NewAttachment("invalid-uri", "text/plain", nil),
				},
				msgCreatePost.PollData,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid uri provided"),
		},
		{
			name: "Empty mime type in message returns error",
			msg: types2.NewMsgCreatePost(
				"My message",
				"",
				types2.CommentsStateAllowed,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				types2.Attachments{
					types2.Attachment{
						URI:      "https://example.com",
						MimeType: "",
					},
				},
				msgCreatePost.PollData,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "mime type must be specified and cannot be empty"),
		},
		{
			name: "Message with invalid PollData returns error",
			msg: types2.NewMsgCreatePost(
				"My message",
				"",
				types2.CommentsStateAllowed,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				nil,
				types2.NewPollData(
					"",
					msgCreatePost.PollData.EndDate,
					msgCreatePost.PollData.ProvidedAnswers,
					true,
					true,
				),
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "missing poll question"),
		},
		{
			name: "Valid message does not return any error",
			msg: types2.NewMsgCreatePost(
				"Message",
				"",
				types2.CommentsStateAllowed,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				[]types2.Attribute{
					{"lorem", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras in dapibus tortor, in iaculis nunc. Integer ac bibendum nisi. Curabitur faucibus vestibulum tincidunt. Donec interdum tincidunt cras amet."},
					{"date", "2020-01-01T00:00.000Z"},
					{"text", "Welcome to Desmos"},
					{"int", "0"},
					{"json", `{"key":"value"}`},
					{"double", "12.0"},
					{"array", `["first","second"]`},
				},
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				types2.Attachments{
					types2.Attachment{
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
			msg: types2.NewMsgCreatePost(
				"My message",
				"",
				types2.CommentsStateAllowed,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				nil,
				msgCreatePost.PollData,
			),
			error: nil,
		},
		{
			name: "Message with non-empty attachments, non-empty poll and non-empty message returns no error",
			msg: types2.NewMsgCreatePost(
				"My message",
				"",
				types2.CommentsStateAllowed,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				msgCreatePost.Attachments,
				msgCreatePost.PollData,
			),
			error: nil,
		},
		{
			name: "Message with non-empty attachments, non empty poll and empty message returns no error",
			msg: types2.NewMsgCreatePost(
				"",
				"",
				types2.CommentsStateAllowed,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				msgCreatePost.Attachments,
				msgCreatePost.PollData,
			),
			error: nil,
		},
		{
			name: "Message with empty attachments, non empty poll and empty message returns no error",
			msg: types2.NewMsgCreatePost(
				"",
				"",
				types2.CommentsStateAllowed,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				nil,
				msgCreatePost.PollData,
			),
			error: nil,
		},
		{
			name: "Message with empty poll, non-empty attachments and non empty message returns no error",
			msg: types2.NewMsgCreatePost(
				"My message",
				"",
				types2.CommentsStateAllowed,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				types2.Attachments{
					types2.Attachment{
						URI:      "https://example.com",
						MimeType: "text/plain",
					},
				},
				nil,
			),
			error: nil,
		},
		{
			name:  "Valid message does not return error",
			msg:   msgCreatePost,
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

func TestMsgCreatePost_GetSignBytes(t *testing.T) {
	tests := []struct {
		name        string
		msg         *types2.MsgCreatePost
		expSignJSON string
	}{
		{
			name: "Message with non-empty external reference",
			msg: types2.NewMsgCreatePost(
				"My new post",
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				types2.CommentsStateAllowed,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				[]types2.Attribute{
					types2.NewAttribute("field", "value"),
				},
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				types2.Attachments{
					types2.NewAttachment("https://uri.com", "text/plain", nil),
				},
				msgCreatePost.PollData,
			),
			expSignJSON: `{"type":"desmos/MsgCreatePost","value":{"additional_attributes":[{"key":"field","value":"value"}],"attachments":[{"mime_type":"text/plain","uri":"https://uri.com"}],"comments_state":1,"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","message":"My new post","parent_id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","poll_data":{"allows_answer_edits":true,"allows_multiple_answers":false,"end_date":"2050-01-01T15:15:00Z","provided_answers":[{"answer_id":"1","text":"Yes"},{"answer_id":"2","text":"No"}],"question":"poll?"},"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`,
		},
		{
			name: "Message with empty external reference",
			msg: types2.NewMsgCreatePost(
				"My post",
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				types2.CommentsStateAllowed,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				types2.Attachments{
					types2.NewAttachment("https://uri.com", "text/plain", nil),
				},
				msgCreatePost.PollData,
			),
			expSignJSON: `{"type":"desmos/MsgCreatePost","value":{"attachments":[{"mime_type":"text/plain","uri":"https://uri.com"}],"comments_state":1,"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","message":"My post","parent_id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","poll_data":{"allows_answer_edits":true,"allows_multiple_answers":false,"end_date":"2050-01-01T15:15:00Z","provided_answers":[{"answer_id":"1","text":"Yes"},{"answer_id":"2","text":"No"}],"question":"poll?"},"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`,
		},
		{
			name: "Message with empty attachments",
			msg: types2.NewMsgCreatePost(
				"My Post without attachments",
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				types2.CommentsStateAllowed,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				types2.Attachments{},
				msgCreatePost.PollData,
			),
			expSignJSON: `{"type":"desmos/MsgCreatePost","value":{"comments_state":1,"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","message":"My Post without attachments","parent_id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","poll_data":{"allows_answer_edits":true,"allows_multiple_answers":false,"end_date":"2050-01-01T15:15:00Z","provided_answers":[{"answer_id":"1","text":"Yes"},{"answer_id":"2","text":"No"}],"question":"poll?"},"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`,
		},
		{
			name: "Message with empty poll poll",
			msg: types2.NewMsgCreatePost(
				"My Post without attachments",
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				types2.CommentsStateAllowed,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				types2.Attachments{
					types2.Attachment{
						URI:      "https://uri.com",
						MimeType: "text/plain",
					},
				},
				nil,
			),
			expSignJSON: `{"type":"desmos/MsgCreatePost","value":{"attachments":[{"mime_type":"text/plain","uri":"https://uri.com"}],"comments_state":1,"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","message":"My Post without attachments","parent_id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`,
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
	addr, _ := sdk.AccAddressFromBech32(msgCreatePost.Creator)
	require.Equal(t, []sdk.AccAddress{addr}, msgCreatePost.GetSigners())
}

func TestMsgCreatePost_ReadJSON(t *testing.T) {
	json := `{"type":"desmos/MsgCreatePost","value":{"parent_id":"","message":"","allows_comments":true,"subspace":"2bdf5932925584b9a86470bea60adce69041608a447f84a3317723aa5678ec88","additional_attributes":[{"key":"local_id","value":"2020-09-15T10:17:54.101972"}],"creator":"cosmos10txl52f64zmp2j7eywawlv9t4xxc4e0wnjlhq9","poll_data":{"question":"What is it better?","end_date":"2020-10-15T08:17:45.639Z","is_open":true,"allows_multiple_answers":false,"allows_answer_edits":false,"provided_answers":[{"id":"0","text":"Sushi\t"},{"id":"1","text":"Pizza"}]}}}`

	var msg types2.MsgCreatePost
	err := types2.ModuleCdc.UnmarshalJSON([]byte(json), &msg)
	require.NoError(t, err)
}

// ___________________________________________________________________________________________________________________

var msgEditPost = types2.NewMsgEditPost(
	"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
	"Edited post message",
	types2.CommentsStateAllowed,
	types2.NewAttachments(
		types2.NewAttachment("https://uri.com", "text/plain", nil),
	),
	types2.NewPollData(
		"poll?",
		time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
		types2.NewPollAnswers(
			types2.NewPollAnswer("1", "Yes"),
			types2.NewPollAnswer("2", "No"),
		),
		false,
		true,
	),
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
)

func TestMsgEditPost_Route(t *testing.T) {
	require.Equal(t, "posts", msgEditPost.Route())
}

func TestMsgEditPost_Type(t *testing.T) {
	require.Equal(t, "edit_post", msgEditPost.Type())
}

func TestMsgEditPost_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types2.MsgEditPost
		error error
	}{
		{
			name: "Invalid post id returns error",
			msg: types2.NewMsgEditPost(
				"",
				"Edited post message",
				types2.CommentsStateAllowed,
				types2.NewAttachments(
					types2.NewAttachment("https://uri.com", "text/plain", nil),
				),
				types2.NewPollData(
					"poll?",
					time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
					types2.NewPollAnswers(
						types2.NewPollAnswer("1", "Yes"),
						types2.NewPollAnswer("2", "No"),
					),
					false,
					true,
				),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(types2.ErrInvalidPostID, ""),
		},
		{
			name: "Invalid editor returns error",
			msg: types2.NewMsgEditPost(
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				"Edited post message",
				types2.CommentsStateAllowed,
				types2.NewAttachments(types2.NewAttachment("https://uri.com", "text/plain", nil)),
				types2.NewPollData(
					"poll?",
					time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
					types2.NewPollAnswers(
						types2.NewPollAnswer("1", "Yes"),
						types2.NewPollAnswer("2", "No"),
					),
					false,
					true,
				),
				"",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid editor"),
		},
		{
			name: "Non-empty message returns no error if attachments are empty",
			msg: types2.NewMsgEditPost(
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				"message",
				types2.CommentsStateAllowed,
				nil,
				nil,
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: nil,
		},
		{
			name: "Non-empty message returns no error if attachments aren't empty",
			msg: types2.NewMsgEditPost(
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				"message",
				types2.CommentsStateAllowed,
				msgCreatePost.Attachments,
				nil,
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: nil,
		},
		{
			name: "Empty message returns no error if poll isn't empty",
			msg: types2.NewMsgEditPost(
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				"",
				types2.CommentsStateAllowed,
				nil,
				msgCreatePost.PollData,
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: nil,
		},
		{
			name: "Empty message returns no error if attachments aren't empty",
			msg: types2.NewMsgEditPost(
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				"",
				types2.CommentsStateAllowed,
				nil,
				msgCreatePost.PollData,
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: nil,
		},
		{
			name: "Non-empty message returns no error if poll is empty",
			msg: types2.NewMsgEditPost(
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				"message",
				types2.CommentsStateAllowed,
				nil,
				nil,
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: nil,
		},
		{
			name: "Empty message returns error if message, attachments and poll are empty",
			msg: types2.NewMsgEditPost(
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				"",
				types2.CommentsStateAllowed,
				nil,
				nil,
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"post message, attachments or poll are required and cannot be all blank or empty"),
		},
		{
			name: "Empty URI in medias returns error",
			msg: types2.NewMsgEditPost(
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				"future post",
				types2.CommentsStateAllowed,
				types2.Attachments{
					types2.NewAttachment("", "text/plain", nil),
				},
				msgCreatePost.PollData,
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid uri provided"),
		},
		{
			name: "Invalid URI in message returns error",
			msg: types2.NewMsgEditPost(
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				"My message",
				types2.CommentsStateAllowed,
				types2.Attachments{
					types2.NewAttachment("invalid-uri", "text/plain", nil),
				},
				msgCreatePost.PollData,
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid uri provided"),
		},
		{
			name: "Empty mime type in message returns error",
			msg: types2.NewMsgEditPost(
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				"My message",
				types2.CommentsStateAllowed,
				types2.Attachments{
					types2.NewAttachment("https://example.com", "", nil),
				},
				msgCreatePost.PollData,
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "mime type must be specified and cannot be empty"),
		},
		{
			name: "Message with invalid PollData returns error",
			msg: types2.NewMsgEditPost(
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				"My message",
				types2.CommentsStateAllowed,
				types2.NewAttachments(types2.NewAttachment("https://uri.com", "text/plain", nil)),
				&types2.PollData{
					Question: "",
					ProvidedAnswers: types2.NewPollAnswers(
						types2.NewPollAnswer("1", "Yes"),
						types2.NewPollAnswer("2", "No"),
					),
					EndDate:           time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
					AllowsAnswerEdits: true,
				},
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "missing poll question"),
		},
		{
			name:  "Valid message returns no error",
			msg:   msgEditPost,
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
	expected := `{"type":"desmos/MsgEditPost","value":{"attachments":[{"mime_type":"text/plain","uri":"https://uri.com"}],"comments_state":1,"editor":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","message":"Edited post message","poll_data":{"allows_answer_edits":true,"allows_multiple_answers":false,"end_date":"2050-01-01T15:15:00Z","provided_answers":[{"answer_id":"1","text":"Yes"},{"answer_id":"2","text":"No"}],"question":"poll?"},"post_id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1"}}`
	require.Equal(t, expected, string(msgEditPost.GetSignBytes()))
}

func TestMsgEditPost_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgEditPost.Editor)
	require.Equal(t, []sdk.AccAddress{addr}, msgEditPost.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgPostReaction = types2.NewMsgAddPostReaction(
	"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
	"like",
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
)

func TestShortCodeRegEx(t *testing.T) {
	for _, em := range emoji.Emojis {
		for _, shortcode := range em.Shortcodes {
			res := types2.IsValidReactionCode(shortcode)
			if !res {
				println(shortcode)
			}
			require.True(t, res)
		}
	}
}

func TestMsgAddPostReaction_Route(t *testing.T) {
	require.Equal(t, "posts", msgPostReaction.Route())
}

func TestMsgAddPostReaction_Type(t *testing.T) {
	require.Equal(t, "add_post_reaction", msgPostReaction.Type())
}

func TestMsgAddPostReaction_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types2.MsgAddPostReaction
		error error
	}{
		{
			name:  "Invalid post id returns error",
			msg:   types2.NewMsgAddPostReaction("", ":like:", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			error: sdkerrors.Wrap(types2.ErrInvalidPostID, ""),
		},
		{
			name: "Invalid user returns error",
			msg: types2.NewMsgAddPostReaction(
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				":like:",
				"",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid user"),
		},
		{
			name: "Invalid value returns error",
			msg: types2.NewMsgAddPostReaction(
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				"like",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(types2.ErrInvalidReactionCode, "like"),
		},
		{
			name: "Valid message returns no error (with shortcode)",
			msg: types2.NewMsgAddPostReaction(
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				":like:",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: nil,
		},
		{
			name: "Valid message returns no error (with emoji)",
			msg: types2.NewMsgAddPostReaction(
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				"🤩",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
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
	expected := `{"type":"desmos/MsgAddPostReaction","value":{"post_id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","reaction":"like","user":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(msgPostReaction.GetSignBytes()))
}

func TestMsgAddPostReaction_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgPostReaction.User)
	require.Equal(t, []sdk.AccAddress{addr}, msgPostReaction.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgRemovePostReaction = types2.NewMsgRemovePostReaction(
	"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	"like",
)

func TestMsgRemovePostReaction_Route(t *testing.T) {
	require.Equal(t, "posts", msgRemovePostReaction.Route())
}

func TestMsgRemovePostReaction_Type(t *testing.T) {
	require.Equal(t, "remove_post_reaction", msgRemovePostReaction.Type())
}

func TestMsgRemovePostReaction_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types2.MsgRemovePostReaction
		error error
	}{
		{
			name:  "Invalid post id returns error",
			msg:   types2.NewMsgRemovePostReaction("", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", ":+1:"),
			error: sdkerrors.Wrap(types2.ErrInvalidPostID, ""),
		},
		{
			name: "Invalid user address: ",
			msg: types2.NewMsgRemovePostReaction(
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				"",
				":like:",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid user"),
		},
		{
			name: "Blank value returns no error",
			msg: types2.NewMsgRemovePostReaction(
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
			),
			error: sdkerrors.Wrap(types2.ErrInvalidReactionCode, ""),
		},
		{
			name: "Valid message returns no error (with shortcode)",
			msg: types2.NewMsgRemovePostReaction(
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				":+1:",
			),
			error: nil,
		},
		{
			name: "Valid message returns no error (with emoji)",
			msg: types2.NewMsgRemovePostReaction(
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"🤩",
			),
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
	expected := `{"type":"desmos/MsgRemovePostReaction","value":{"post_id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","reaction":"like","user":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(msgRemovePostReaction.GetSignBytes()))
}

func TestMsgRemovePostReaction_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgRemovePostReaction.User)
	require.Equal(t, []sdk.AccAddress{addr}, msgRemovePostReaction.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgAnswerPollPost = types2.NewMsgAnswerPoll(
	"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
	[]string{"1", "2"},
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
)

func TestMsgAnswerPollPost_Route(t *testing.T) {
	require.Equal(t, "posts", msgAnswerPollPost.Route())
}

func TestMsgAnswerPollPost_Type(t *testing.T) {
	require.Equal(t, "answer_poll", msgAnswerPollPost.Type())
}

func TestMsgAnswerPollPost_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types2.MsgAnswerPoll
		error error
	}{
		{
			name:  "Invalid post id",
			msg:   types2.NewMsgAnswerPoll("", []string{"1", "2"}, msgAnswerPollPost.Answerer),
			error: sdkerrors.Wrap(types2.ErrInvalidPostID, ""),
		},
		{
			name: "Invalid answerer address",
			msg: types2.NewMsgAnswerPoll(
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				[]string{"1", "2"},
				"",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid answerer"),
		},
		{
			name: "Returns error when no answer is provided",
			msg: types2.NewMsgAnswerPoll(
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				[]string{},
				msgAnswerPollPost.Answerer,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "provided answer must contains at least one answer"),
		},
		{
			name: "Valid message returns no error",
			msg: types2.NewMsgAnswerPoll(
				"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				[]string{"1", "2"},
				msgAnswerPollPost.Answerer,
			),
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
	expected := `{"type":"desmos/MsgAnswerPoll","value":{"answerer":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","answers":["1","2"],"post_id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1"}}`
	require.Equal(t, expected, string(msgAnswerPollPost.GetSignBytes()))
}

func TestMsgAnswerPollPost_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgAnswerPollPost.Answerer)
	require.Equal(t, []sdk.AccAddress{addr}, msgAnswerPollPost.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgRegisterReaction = types2.NewMsgRegisterReaction(
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	":smile:",
	"https://smile.jpg",
	"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
)

func TestMsgRegisterReaction_Route(t *testing.T) {
	require.Equal(t, "posts", msgRegisterReaction.Route())
}

func TestMsgRegisterReaction_Type(t *testing.T) {
	require.Equal(t, "register_reaction", msgRegisterReaction.Type())
}

func TestMsgRegisterReaction_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types2.MsgRegisterReaction
		error error
	}{
		{
			name: "Invalid creator returns error",
			msg: types2.NewMsgRegisterReaction(
				"",
				":smile:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator"),
		},
		{
			name: "Empty short code returns error",
			msg: types2.NewMsgRegisterReaction(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(types2.ErrInvalidReactionCode, ""),
		},
		{
			name: "Invalid short code returns error",
			msg: types2.NewMsgRegisterReaction(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				":smile",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(types2.ErrInvalidReactionCode, ":smile"),
		},
		{
			name: "Empty value returns error",
			msg: types2.NewMsgRegisterReaction(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				":smile:",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(commonerrors.ErrInvalidURI, "reaction value should be a valid uri"),
		},
		{
			name: "Invalid value returns error (url)",
			msg: types2.NewMsgRegisterReaction(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				":smile:",
				"htp://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(commonerrors.ErrInvalidURI, "reaction value should be a valid uri"),
		},
		{
			name: "Invalid value returns error (unicode)",
			msg: types2.NewMsgRegisterReaction(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				":smile:",
				"U+1",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(commonerrors.ErrInvalidURI, "reaction value should be a valid uri"),
		},
		{
			name: "Valid emoji value returns no error",
			msg: types2.NewMsgRegisterReaction(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				":smile:",
				"💙",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(commonerrors.ErrInvalidURI, "reaction value should be a valid uri"),
		},
		{
			name: "Invalid subspace returns error",
			msg: types2.NewMsgRegisterReaction(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				":smile:",
				"https://smile.jpg",
				"1234",
			),
			error: sdkerrors.Wrap(types2.ErrInvalidSubspace, "reaction subspace must be a valid sha-256 hash"),
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
	expected := `{"type":"desmos/MsgRegisterReaction","value":{"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","short_code":":smile:","subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e","value":"https://smile.jpg"}}`
	require.Equal(t, expected, string(msgRegisterReaction.GetSignBytes()))
}

func TestMsgRegisterReaction_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgRegisterReaction.Creator)
	require.Equal(t, []sdk.AccAddress{addr}, msgRegisterReaction.GetSigners())
}

// ___________________________________________________________________________________________________________________

func TestMsgReportPost_Route(t *testing.T) {
	msg := types2.NewMsgReportPost(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"type",
		"message",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, "posts", msg.Route())
}

func TestMsgReportPost_Type(t *testing.T) {
	msg := types2.NewMsgReportPost(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"type",
		"message",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, "report_post", msg.Type())
}

func TestMsgReportPost_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types2.MsgReportPost
		error error
	}{
		{
			name: "invalid post id returns error",
			msg: types2.NewMsgReportPost(
				"123",
				"type",
				"message",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(types2.ErrInvalidPostID, "123"),
		},
		{
			name: "invalid report type returns error",
			msg: types2.NewMsgReportPost(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"",
				"message",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "report type cannot be empty"),
		},
		{
			name: "invalid report message returns error",
			msg: types2.NewMsgReportPost(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"scam",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "report message cannot be empty"),
		},
		{
			name: "invalid report creator returns error",
			msg: types2.NewMsgReportPost(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"scam",
				"message",
				"address",
			),
			error: sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid report creator: address"),
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

func TestMsgReportPost_GetSignBytes(t *testing.T) {
	msg := types2.NewMsgReportPost(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"type",
		"message",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	expected := `{"type":"desmos/MsgReportPost","value":{"message":"message","post_id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af","report_type":"type","user":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

func TestNewMsgReportPost_GetSigners(t *testing.T) {
	msg := types2.NewMsgReportPost(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"type",
		"message",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	addr, _ := sdk.AccAddressFromBech32(msg.User)
	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())
}
