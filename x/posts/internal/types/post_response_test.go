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
	postOwner, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	assert.NoError(t, err)

	liker, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	assert.NoError(t, err)

	otherLiker, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	assert.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	assert.NoError(t, err)

	date := time.Date(2020, 2, 2, 15, 0, 0, 0, timeZone)
	medias := types.PostMedias{
		types.PostMedia{
			URI:      "https://uri.com",
			MimeType: "text/plain",
		},
	}

	post := types.NewPost(
		types.PostID(10),
		types.PostID(0),
		"Post",
		true,
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		map[string]string{},
		date,
		postOwner,
		medias,
	)
	postNoMedia := types.NewPost(
		types.PostID(10),
		types.PostID(0),
		"Post",
		true,
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		map[string]string{},
		date,
		postOwner,
		types.PostMedias{},
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
			expResponse: `{"id":"10","parent_id":"0","message":"Post","created":"2020-02-02T15:00:00Z","last_edited":"0001-01-01T00:00:00Z","allows_comments":true,"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e","creator":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","medias":[{"uri":"https://uri.com","mime_Type":"text/plain"}],"reactions":[{"owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4","value":"like"},{"owner":"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae","value":"like"}],"children":["98","100"]}`,
		},
		{
			name:        "Post Query Response with Post that not contains media",
			response:    types.NewPostResponse(postNoMedia, likes, children),
			expResponse: `{"id":"10","parent_id":"0","message":"Post","created":"2020-02-02T15:00:00Z","last_edited":"0001-01-01T00:00:00Z","allows_comments":true,"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e","creator":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","reactions":[{"owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4","value":"like"},{"owner":"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae","value":"like"}],"children":["98","100"]}`,
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
	postOwner, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	assert.NoError(t, err)

	liker, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	assert.NoError(t, err)

	otherLiker, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	assert.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	assert.NoError(t, err)

	date := time.Date(2020, 2, 2, 15, 0, 0, 0, timeZone)
	medias := types.PostMedias{
		types.PostMedia{
			URI:      "https://uri.com",
			MimeType: "text/plain",
		},
	}

	post := types.NewPost(
		types.PostID(10),
		types.PostID(0),
		"Post",
		true,
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		map[string]string{},
		date,
		postOwner,
		medias,
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
