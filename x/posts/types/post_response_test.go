package types_test

import (
	"encoding/json"
	"github.com/desmos-labs/desmos/x/posts/types"
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

	post := types.NewPost(
		"",
		"Post",
		true,
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		nil,
		time.Date(2020, 2, 2, 15, 0, 0, 0, timeZone),
		postOwner,
	)

	attachments := models.NewAttachments(
		types.NewAttachment("https://uri.com", "text/plain", []sdk.AccAddress{postOwner}),
	)

	attachmentsNoTags := models.NewAttachments(
		types.NewAttachment("https://uri.com", "text/plain", nil),
	)

	pollData := types.NewPollData(
		"poll?",
		time.Date(2050, 1, 1, 15, 15, 00, 000, timeZone),
		models.PollAnswers{
			types.NewPollAnswer(models.AnswerID(1), "Yes"),
			types.NewPollAnswer(models.AnswerID(2), "No"),
		},
		false,
		true,
	)

	answersDetails := types.NewUserAnswers(
		types.NewUserAnswer([]models.AnswerID{models.AnswerID(1)}, liker),
	)

	children := types.PostIDs{
		"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
		"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
	}

	reactionsResponses := []models.PostReaction{
		models.NewPostReaction(":like:", "https://example.com/like", liker),
		models.NewPostReaction(":+1:", "👍", otherLiker),
	}

	tests := []struct {
		name        string
		response    types.PostQueryResponse
		expResponse string
	}{
		{
			name: "Post Query Response with Post that contains attachment and poll",
			response: types.NewPostResponse(
				post.WithAttachments(attachments).WithPollData(pollData),
				answersDetails,
				reactionsResponses,
				children,
			),
			expResponse: `{"id":"230f2001f05281763b866c07badcd7b81e3708daac84db9f7bf9811934dbfa00","parent_id":"","message":"Post","created":"2020-02-02T15:00:00Z","last_edited":"0001-01-01T00:00:00Z","allows_comments":true,"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e","creator":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","attachments":[{"uri":"https://uri.com","mime_type":"text/plain","tags":["cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"]}],"poll_data":{"question":"poll?","provided_answers":[{"id":"1","text":"Yes"},{"id":"2","text":"No"}],"end_date":"2050-01-01T15:15:00Z","allows_multiple_answers":false,"allows_answer_edits":true},"poll_answers":[{"answers":["1"],"user":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"}],"reactions":[{"shortcode":":like:","value":"https://example.com/like","owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},{"shortcode":":+1:","value":"👍","owner":"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae"}],"children":["dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1"]}`,
		},
		{
			name: "Post Query Response with Post that contains attachment without tags",
			response: types.NewPostResponse(
				post.WithAttachments(attachmentsNoTags),
				answersDetails,
				reactionsResponses,
				children,
			),
			expResponse: `{"id":"3e7612dcc4d67125102101c866cdd0be54470f1bcc9ad378e03410e0e2c29705","parent_id":"","message":"Post","created":"2020-02-02T15:00:00Z","last_edited":"0001-01-01T00:00:00Z","allows_comments":true,"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e","creator":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","attachments":[{"uri":"https://uri.com","mime_type":"text/plain"}],"poll_answers":[{"answers":["1"],"user":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"}],"reactions":[{"shortcode":":like:","value":"https://example.com/like","owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},{"shortcode":":+1:","value":"👍","owner":"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae"}],"children":["dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1"]}`,
		},
		{
			name: "Post Query with Post that not contains poll",
			response: types.NewPostResponse(
				post.WithAttachments(attachments),
				nil,
				reactionsResponses,
				children,
			),
			expResponse: `{"id":"a03436d2abe887a65fa9e2b16a6a48ce39e81e09de3eb950457a64c5bc3a1237","parent_id":"","message":"Post","created":"2020-02-02T15:00:00Z","last_edited":"0001-01-01T00:00:00Z","allows_comments":true,"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e","creator":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","attachments":[{"uri":"https://uri.com","mime_type":"text/plain","tags":["cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"]}],"reactions":[{"shortcode":":like:","value":"https://example.com/like","owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},{"shortcode":":+1:","value":"👍","owner":"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae"}],"children":["dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1"]}`,
		},
		{
			name: "Post Query Response with Post that not contains attachment",
			response: types.NewPostResponse(
				post.WithPollData(pollData),
				answersDetails,
				reactionsResponses,
				children,
			),
			expResponse: `{"id":"3731f8eb41239386f42cd599cc93d06a096a2767d1926ec1915c103a2e7f5ad1","parent_id":"","message":"Post","created":"2020-02-02T15:00:00Z","last_edited":"0001-01-01T00:00:00Z","allows_comments":true,"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e","creator":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","poll_data":{"question":"poll?","provided_answers":[{"id":"1","text":"Yes"},{"id":"2","text":"No"}],"end_date":"2050-01-01T15:15:00Z","allows_multiple_answers":false,"allows_answer_edits":true},"poll_answers":[{"answers":["1"],"user":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"}],"reactions":[{"shortcode":":like:","value":"https://example.com/like","owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},{"shortcode":":+1:","value":"👍","owner":"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae"}],"children":["dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1"]}`,
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

	attachments := models.NewAttachments(types.NewAttachment("https://uri.com", "text/plain", []sdk.AccAddress{postOwner}))

	pollData := types.NewPollData(
		"poll?",
		date.Add(time.Hour),
		models.NewPollAnswers(
			types.NewPollAnswer(models.AnswerID(1), "Yes"),
			types.NewPollAnswer(models.AnswerID(2), "No"),
		),
		false,
		true,
	)

	postResponse := types.NewPostResponse(
		types.NewPost(
			"",
			"Post",
			true,
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			nil,
			time.Date(2020, 2, 2, 15, 0, 0, 0, timeZone),
			postOwner,
		).WithAttachments(attachments).WithPollData(pollData),
		types.NewUserAnswers(
			types.NewUserAnswer([]models.AnswerID{models.AnswerID(1)}, liker),
		),
		[]models.PostReaction{
			models.NewPostReaction(":like:", "https://example.com/like", liker),
			models.NewPostReaction(":+1:", "👍", otherLiker),
		},
		types.PostIDs{
			"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
			"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
		},
	)

	expected := "ID: 93ed93d2f3b3363399c8b7a4509f0530a8c863c16c755bf916901ab55ce33322\nReactions: [[Shortcode] :like: [Value] https://example.com/like [Owner] cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4 [Shortcode] :+1: [Value] 👍 [Owner] cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae]\nChildren: [dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1, dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1]"
	stringResponse := postResponse.String()
	require.Equal(t, strings.TrimSpace(expected), stringResponse)
}
