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

	post := types.NewPost(types.PostID(10), types.PostID(0), "Post", true, "desmos", map[string]string{}, 10, postOwner)
	likes := types.Reactions{
		types.NewReaction("like", 11, liker),
		types.NewReaction("like", 12, otherLiker),
	}
	children := types.PostIDs{types.PostID(98), types.PostID(100)}

	response := types.NewPostResponse(post, likes, children)
	jsonData, err := json.Marshal(&response)
	assert.NoError(t, err)
	assert.Equal(t,
		`{"id":"10","parent_id":"0","message":"Post","created":"10","last_edited":"0","allows_comments":true,"subspace":"desmos","owner":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","reactions":[{"created":"11","owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4","value":"like"},{"created":"12","owner":"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae","value":"like"}],"children":["98","100"]}`,
		string(jsonData),
	)
}
