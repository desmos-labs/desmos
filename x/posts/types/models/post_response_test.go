package models_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/types/models"
	"github.com/stretchr/testify/require"
)

func TestPostQueryResponse_MarshalJSON(t *testing.T) {
	postOwner, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	liker, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	otherLiker, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	post := models.NewPost(
		"",
		"Post",
		true,
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		map[string]string{},
		time.Date(2020, 2, 2, 15, 0, 0, 0, timeZone),
		postOwner,
	)

	attachments := models.NewAttachments(
		models.NewAttachment("https://uri.com", "text/plain", []sdk.AccAddress{postOwner}),
	)

	attachmentsNoTags := models.NewAttachments(
		models.NewAttachment("https://uri.com", "text/plain", nil),
	)

	pollData := models.NewPollData(
		"poll?",
		time.Date(2050, 1, 1, 15, 15, 00, 000, timeZone),
		models.PollAnswers{
			models.NewPollAnswer(models.AnswerID(1), "Yes"),
			models.NewPollAnswer(models.AnswerID(2), "No"),
		},
		true,
		false,
		true,
	)

	answersDetails := models.NewUserAnswers(
		models.NewUserAnswer([]models.AnswerID{models.AnswerID(1)}, liker),
	)

	children := models.PostIDs{
		"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
		"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
	}

	reactionsResponses := []models.PostReaction{
		models.NewPostReaction(":like:", "https://example.com/like", liker),
		models.NewPostReaction(":+1:", "👍", otherLiker),
	}

	tests := []struct {
		name        string
		response    models.PostQueryResponse
		expResponse string
	}{
		{
			name: "Post Query Response with Post that contains attachment and poll",
			response: models.NewPostResponse(
				post.WithAttachments(attachments).WithPollData(pollData),
				answersDetails,
				reactionsResponses,
				children,
			),
			expResponse: `{"id":"d621d4996f94d9edf32487e6a6637d122ef5e5048b38214993b5eaa25420e47a","parent_id":"","message":"Post","created":"2020-02-02T15:00:00Z","last_edited":"0001-01-01T00:00:00Z","allows_comments":true,"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e","creator":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","attachments":[{"uri":"https://uri.com","mime_type":"text/plain","tags":["cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"]}],"poll_data":{"question":"poll?","provided_answers":[{"id":"1","text":"Yes"},{"id":"2","text":"No"}],"end_date":"2050-01-01T15:15:00Z","is_open":true,"allows_multiple_answers":false,"allows_answer_edits":true},"poll_answers":[{"answers":["1"],"user":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"}],"reactions":[{"shortcode":":like:","value":"https://example.com/like","owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},{"shortcode":":+1:","value":"👍","owner":"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae"}],"children":["dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1"]}`,
		},
		{
			name: "Post Query Response with Post that contains attachment without tags",
			response: models.NewPostResponse(
				post.WithAttachments(attachmentsNoTags),
				answersDetails,
				reactionsResponses,
				children,
			),
			expResponse: `{"id":"ab5a091250f9bdee2fe5b214e362aceabb14878a1c0cb0598742cbf083863d21","parent_id":"","message":"Post","created":"2020-02-02T15:00:00Z","last_edited":"0001-01-01T00:00:00Z","allows_comments":true,"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e","creator":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","attachments":[{"uri":"https://uri.com","mime_type":"text/plain"}],"poll_answers":[{"answers":["1"],"user":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"}],"reactions":[{"shortcode":":like:","value":"https://example.com/like","owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},{"shortcode":":+1:","value":"👍","owner":"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae"}],"children":["dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1"]}`,
		},
		{
			name: "Post Query with Post that not contains poll",
			response: models.NewPostResponse(
				post.WithAttachments(attachments),
				nil,
				reactionsResponses,
				children,
			),
			expResponse: `{"id":"205e9e422c4a79f67bc647cf8c49a65b878ee026d722d71ee068edad62a0169a","parent_id":"","message":"Post","created":"2020-02-02T15:00:00Z","last_edited":"0001-01-01T00:00:00Z","allows_comments":true,"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e","creator":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","attachments":[{"uri":"https://uri.com","mime_type":"text/plain","tags":["cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"]}],"reactions":[{"shortcode":":like:","value":"https://example.com/like","owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},{"shortcode":":+1:","value":"👍","owner":"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae"}],"children":["dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1"]}`,
		},
		{
			name: "Post Query Response with Post that not contains attachment",
			response: models.NewPostResponse(
				post.WithPollData(pollData),
				answersDetails,
				reactionsResponses,
				children,
			),
			expResponse: `{"id":"944d8d6da2e2423092d437ecc85ca82c3ab4d5bf16a902d82f22ff03ce5a8496","parent_id":"","message":"Post","created":"2020-02-02T15:00:00Z","last_edited":"0001-01-01T00:00:00Z","allows_comments":true,"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e","creator":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","poll_data":{"question":"poll?","provided_answers":[{"id":"1","text":"Yes"},{"id":"2","text":"No"}],"end_date":"2050-01-01T15:15:00Z","is_open":true,"allows_multiple_answers":false,"allows_answer_edits":true},"poll_answers":[{"answers":["1"],"user":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"}],"reactions":[{"shortcode":":like:","value":"https://example.com/like","owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},{"shortcode":":+1:","value":"👍","owner":"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae"}],"children":["dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1"]}`,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			jsonData, err := json.Marshal(&test.response)
			require.NoError(t, err)
			require.Equal(t, test.expResponse, string(jsonData))
		})
	}
}

func TestPostQueryResponse_String(t *testing.T) {
	postOwner, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	liker, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	otherLiker, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2020, 1, 1, 12, 0, 00, 000, timeZone)

	attachments := models.NewAttachments(models.NewAttachment("https://uri.com", "text/plain", []sdk.AccAddress{postOwner}))

	pollData := models.NewPollData(
		"poll?",
		date.Add(time.Hour),
		models.NewPollAnswers(
			models.NewPollAnswer(models.AnswerID(1), "Yes"),
			models.NewPollAnswer(models.AnswerID(2), "No"),
		),
		true,
		false,
		true,
	)

	postResponse := models.NewPostResponse(
		models.NewPost(
			"",
			"Post",
			true,
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			map[string]string{},
			time.Date(2020, 2, 2, 15, 0, 0, 0, timeZone),
			postOwner,
		).WithAttachments(attachments).WithPollData(pollData),
		models.NewUserAnswers(
			models.NewUserAnswer([]models.AnswerID{models.AnswerID(1)}, liker),
		),
		[]models.PostReaction{
			models.NewPostReaction(":like:", "https://example.com/like", liker),
			models.NewPostReaction(":+1:", "👍", otherLiker),
		},
		models.PostIDs{
			"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
			"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
		},
	)

	expected := "ID: 92cab7ec0273e74ecd6f4831e3e3800a58f3b976fc0fea2431e6fd30346e20df\nReactions: [[Shortcode] :like: [Value] https://example.com/like [Owner] cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4 [Shortcode] :+1: [Value] 👍 [Owner] cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae]\nChildren: [dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1, dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1]"
	stringResponse := postResponse.String()
	require.Equal(t, strings.TrimSpace(expected), stringResponse)
}
