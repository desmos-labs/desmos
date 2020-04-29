package types_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/desmos-labs/desmos/x/posts/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
		"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
		"",
		"Post",
		true,
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		map[string]string{},
		time.Date(2020, 2, 2, 15, 0, 0, 0, timeZone),
		postOwner,
	)

	medias := types.NewPostMedias(
		types.NewPostMedia("https://uri.com", "text/plain"),
	)

	pollData := types.NewPollData(
		"poll?",
		time.Date(2050, 1, 1, 15, 15, 00, 000, timeZone),
		types.PollAnswers{answer, answer2},
		true,
		false,
		true,
	)

	answersDetails := types.NewUserAnswers(
		types.NewUserAnswer([]types.AnswerID{types.AnswerID(1)}, liker),
	)

	children := types.PostIDs{id, id}

	reactionsResponses := []types.ReactionQueryResponse{
		types.NewReactionQueryResponse("https://example.com/like", ":like:", liker),
		types.NewReactionQueryResponse("👍", ":+1:", otherLiker),
	}

	tests := []struct {
		name        string
		response    types.PostQueryResponse
		expResponse string
	}{
		{
			name: "Post Query Response with Post that contains media and poll",
			response: types.NewPostResponse(
				post.WithMedias(medias).WithPollData(pollData),
				answersDetails,
				reactionsResponses,
				children,
			),
			expResponse: `{"id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","parent_id":"","message":"Post","created":"2020-02-02T15:00:00Z","last_edited":"0001-01-01T00:00:00Z","allows_comments":true,"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e","creator":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","medias":[{"uri":"https://uri.com","mime_type":"text/plain"}],"poll_data":{"question":"poll?","provided_answers":[{"id":"1","text":"Yes"},{"id":"2","text":"No"}],"end_date":"2050-01-01T15:15:00Z","is_open":true,"allows_multiple_answers":false,"allows_answer_edits":true},"poll_answers":[{"answers":["1"],"user":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"}],"reactions":[{"value":"https://example.com/like","shortcode":":like:","owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},{"value":"👍","shortcode":":+1:","owner":"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae"}],"children":["dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1"]}`,
		},
		{
			name: "Post Query with Post that not contains poll",
			response: types.NewPostResponse(
				post.WithMedias(medias),
				nil,
				reactionsResponses,
				children,
			),
			expResponse: `{"id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","parent_id":"","message":"Post","created":"2020-02-02T15:00:00Z","last_edited":"0001-01-01T00:00:00Z","allows_comments":true,"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e","creator":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","medias":[{"uri":"https://uri.com","mime_type":"text/plain"}],"reactions":[{"value":"https://example.com/like","shortcode":":like:","owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},{"value":"👍","shortcode":":+1:","owner":"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae"}],"children":["dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1"]}`,
		},
		{
			name: "Post Query Response with Post that not contains media",
			response: types.NewPostResponse(
				post.WithPollData(pollData),
				answersDetails,
				reactionsResponses,
				children,
			),
			expResponse: `{"id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","parent_id":"","message":"Post","created":"2020-02-02T15:00:00Z","last_edited":"0001-01-01T00:00:00Z","allows_comments":true,"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e","creator":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","poll_data":{"question":"poll?","provided_answers":[{"id":"1","text":"Yes"},{"id":"2","text":"No"}],"end_date":"2050-01-01T15:15:00Z","is_open":true,"allows_multiple_answers":false,"allows_answer_edits":true},"poll_answers":[{"answers":["1"],"user":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"}],"reactions":[{"value":"https://example.com/like","shortcode":":like:","owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},{"value":"👍","shortcode":":+1:","owner":"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae"}],"children":["dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1"]}`,
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

	medias := types.NewPostMedias(types.NewPostMedia("https://uri.com", "text/plain"))

	pollData := types.NewPollData(
		"poll?",
		time.Now().UTC().Add(time.Hour),
		types.NewPollAnswers(
			types.NewPollAnswer(types.AnswerID(1), "Yes"),
			types.NewPollAnswer(types.AnswerID(2), "No"),
		),
		true,
		false,
		true,
	)

	postResponse := types.NewPostResponse(
		types.NewPost(
			"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
			"",
			"Post",
			true,
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			map[string]string{},
			time.Date(2020, 2, 2, 15, 0, 0, 0, timeZone),
			postOwner,
		).WithMedias(medias).WithPollData(pollData),
		types.NewUserAnswers(
			types.NewUserAnswer([]types.AnswerID{types.AnswerID(1)}, liker),
		),
		[]types.ReactionQueryResponse{
			types.NewReactionQueryResponse("https://example.com/like", ":like:", liker),
			types.NewReactionQueryResponse("👍", ":+1:", otherLiker),
		},
		types.PostIDs{id, id},
	)

	expected := `
ID: dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1
Reactions: [{https://example.com/like :like: cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4} {👍 :+1: cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae}]
Children: [dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1, dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1]
`
	stringResponse := postResponse.String()
	require.Equal(t, strings.TrimSpace(expected), stringResponse)
}
