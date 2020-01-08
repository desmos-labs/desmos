package keeper_test

import (
	"encoding/json"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/keeper"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestPostQueryResponse_MarshalJSON(t *testing.T) {
	postOwner, _ := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	liker, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	otherLiker, _ := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")

	post := types.NewPost(types.PostID(10), types.PostID(0), "Post", true, "", 10, postOwner)
	likes := types.Likes{
		types.NewLike(11, liker),
		types.NewLike(12, otherLiker),
	}
	children := types.PostIDs{types.PostID(98), types.PostID(100)}

	response := keeper.NewPostResponse(post, likes, children)
	jsonData, err := json.Marshal(&response)
	assert.NoError(t, err)
	assert.Equal(t,
		`{"post_id":"10","parent_id":"0","message":"Post","created":"10","last_edited":"0","allows_comments":true,"external_reference":"","owner":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","likes":[{"created":"11","owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},{"created":"12","owner":"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae"}],"children":["98","100"]}`,
		string(jsonData),
	)
}
