package types_test

import (
	"encoding/json"
	"testing"

	"github.com/desmos-labs/desmos/x/posts/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

func TestPostQueryResponse_MarshalJSON(t *testing.T) {
	postOwner, _ := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	liker, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	otherLiker, _ := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")

	post := types.NewTextPost(types.PostID(10), types.PostID(0), "TextPost", true, "desmos", map[string]string{}, 10, postOwner)
	likes := types.Reactions{
		types.NewReaction("like", 11, liker),
		types.NewReaction("like", 12, otherLiker),
	}
	children := types.PostIDs{types.PostID(98), types.PostID(100)}

	textPostResponse := types.NewPostResponse(post, likes, children)

	media := types.PostMedia{
		Provider: "provider",
		URI:      "uri",
		MimeType: "text/plain",
	}

	mediaPost := types.NewMediaPost(post, []types.PostMedia{media})

	mediaPostResponse := types.NewPostResponse(mediaPost, likes, children)

	tests := []struct {
		name        string
		response    types.PostQueryResponse
		expResponse string
	}{
		{
			name:        "Post Query Response with TextPost",
			response:    textPostResponse,
			expResponse: `{"type":"TextPost","post":{"id":"10","parent_id":"0","message":"TextPost","created":"10","last_edited":"0","allows_comments":true,"subspace":"desmos","creator":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},"reactions":[{"created":"11","owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4","value":"like"},{"created":"12","owner":"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae","value":"like"}],"children":["98","100"]}`,
		},
		{
			name:        "Post Query Response with MediaPost",
			response:    mediaPostResponse,
			expResponse: `{"type":"MediaPost","post":{"id":"2","parent_id":"0","message":"media Post","created":"0","last_edited":"0","allows_comments":false,"subspace":"desmos","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},"medias":[{"provider":"ipfs","uri":"uri","mime_Type":"text/plain"}],"reactions":[{"created":"11","owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4","value":"like"},{"created":"12","owner":"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae","value":"like"}],"children":["98","100"]}`,
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
