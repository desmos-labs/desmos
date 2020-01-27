package types_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/desmos-labs/desmos/x/posts/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

func TestPostQueryResponse_MarshalJSON(t *testing.T) {
	postOwner, _ := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	liker, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	otherLiker, _ := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")

	timeZone, _ := time.LoadLocation("UTC")
	date := time.Date(2020, 2, 2, 15, 0, 0, 0, timeZone)
	medias := types.PostMedias{
		types.PostMedia{
			URI:      "https://uri.com",
			MimeType: "text/plain",
		},
	}

	testPostEndPollDate := time.Date(2050, 1, 1, 15, 15, 00, 000, timeZone)
	pollData := types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, false, true)

	post := types.NewPost(
		types.PostID(10),
		types.PostID(0),
		"Post",
		true,
		"desmos",
		map[string]string{},
		date,
		postOwner,
		medias,
		pollData,
	)
	postNoMedia := types.NewPost(
		types.PostID(10),
		types.PostID(0),
		"Post",
		true,
		"desmos",
		map[string]string{},
		date,
		postOwner,
		types.PostMedias{},
		pollData,
	)

	likes := types.Reactions{
		types.NewReaction("like", liker),
		types.NewReaction("like", otherLiker),
	}
	children := types.PostIDs{types.PostID(98), types.PostID(100)}

	PostResponse := types.NewPostResponse(post, likes, children)

	tests := []struct {
		name        string
		response    types.PostQueryResponse
		expResponse string
	}{
		{
			name:        "Post Query Response with Post that contains media",
			response:    PostResponse,
			expResponse: `{"id":"10","parent_id":"0","message":"Post","created":"2020-02-02T15:00:00Z","last_edited":"0001-01-01T00:00:00Z","allows_comments":true,"subspace":"desmos","creator":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","medias":[{"uri":"https://uri.com","mime_Type":"text/plain"}],"poll_data":{"end_date":"2050-01-01T15:15:00Z","provided_answers":[{"id":1,"text":"Yes"},{"id":2,"text":"No"}],"title":"poll?","open":true,"allows_multiple_answers":false,"allows_answer_edits":true},"reactions":[{"owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4","value":"like"},{"owner":"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae","value":"like"}],"children":["98","100"]}`,
		},
		{
			name:        "Post Query Response with Post that not contains media",
			response:    types.NewPostResponse(postNoMedia, likes, children),
			expResponse: `{"id":"10","parent_id":"0","message":"Post","created":"2020-02-02T15:00:00Z","last_edited":"0001-01-01T00:00:00Z","allows_comments":true,"subspace":"desmos","creator":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","poll_data":{"end_date":"2050-01-01T15:15:00Z","provided_answers":[{"id":1,"text":"Yes"},{"id":2,"text":"No"}],"title":"poll?","open":true,"allows_multiple_answers":false,"allows_answer_edits":true},"reactions":[{"owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4","value":"like"},{"owner":"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae","value":"like"}],"children":["98","100"]}`,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			jsonData, err := json.Marshal(&test.response)
			assert.NoError(t, err)
			assert.Equal(t,
				test.expResponse,
				string(jsonData),
			)
		})
	}
}

func TestPostQueryResponse_String(t *testing.T) {
	postOwner, _ := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	liker, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	otherLiker, _ := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")

	timeZone, _ := time.LoadLocation("UTC")
	date := time.Date(2020, 2, 2, 15, 0, 0, 0, timeZone)
	medias := types.PostMedias{
		types.PostMedia{
			URI:      "https://uri.com",
			MimeType: "text/plain",
		},
	}
	answer := types.PollAnswer{
		ID:   uint64(1),
		Text: "Yes",
	}

	answer2 := types.PollAnswer{
		ID:   uint64(2),
		Text: "No",
	}
	pollData := types.NewPollData("poll?", time.Now().UTC().Add(time.Hour), types.PollAnswers{answer, answer2}, true, false, true)

	post := types.NewPost(
		types.PostID(10),
		types.PostID(0),
		"Post",
		true,
		"desmos",
		map[string]string{},
		date,
		postOwner,
		medias,
		pollData,
	)

	likes := types.Reactions{
		types.NewReaction("like", liker),
		types.NewReaction("like", otherLiker),
	}
	children := types.PostIDs{types.PostID(98), types.PostID(100)}

	PostResponse := types.NewPostResponse(post, likes, children)

	tests := []struct {
		name        string
		response    types.PostQueryResponse
		expResponse string
	}{
		{
			name:        "Post query response string",
			response:    PostResponse,
			expResponse: "ID - [Reactions] [Children] \n10 - [[{\"owner\":\"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4\",\"value\":\"like\"} {\"owner\":\"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae\",\"value\":\"like\"}]] [[98 100]]",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			stringResponse := test.response.String()
			assert.Equal(t, test.expResponse, stringResponse)
		})
	}
}
