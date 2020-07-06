package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
)

// -------------
// --- PostReactions
// -------------

func (suite *KeeperTestSuite) TestKeeper_SaveReaction() {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	liker, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	suite.NoError(err)

	otherLiker, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	suite.NoError(err)

	tests := []struct {
		name           string
		storedReaction types.PostReactions
		postID         types.PostID
		reaction       types.PostReaction
		storedPost     types.Post
		error          error
		expectedStored types.PostReactions
	}{
		{
			name:           "Reaction from same user already present returns expError",
			storedReaction: types.PostReactions{types.NewPostReaction(":like:", "👍", liker)},
			postID:         id,
			reaction:       types.NewPostReaction(":like:", "👍", liker),
			storedPost: types.NewPost(
				id,
				testPost.ParentID,
				testPost.Message,
				testPost.AllowsComments,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPost.Creator,
			),
			error:          fmt.Errorf("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4 has already reacted with :like: to the post with id 19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"),
			expectedStored: types.PostReactions{types.NewPostReaction(":like:", "👍", liker)},
		},
		{
			name:           "First liker is stored properly",
			storedReaction: types.PostReactions{},
			postID:         id,
			reaction:       types.NewPostReaction(":like:", "👍", liker),
			storedPost: types.NewPost(
				id,
				testPost.ParentID,
				testPost.Message,
				testPost.AllowsComments,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPost.Creator,
			),
			error:          nil,
			expectedStored: types.PostReactions{types.NewPostReaction(":like:", "👍", liker)},
		},
		{
			name:           "Second liker is stored properly",
			storedReaction: types.PostReactions{types.NewPostReaction(":like:", "👍", liker)},
			postID:         id,
			reaction:       types.NewPostReaction(":like:", "👍", otherLiker),
			storedPost: types.NewPost(
				id,
				testPost.ParentID,
				testPost.Message,
				testPost.AllowsComments,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPost.Creator,
			),
			error: nil,
			expectedStored: types.PostReactions{
				types.NewPostReaction(":like:", "👍", liker),
				types.NewPostReaction(":like:", "👍", otherLiker),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() // reset
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			if len(test.storedReaction) != 0 {
				store.Set(types.PostReactionsStoreKey(test.postID), suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedReaction))
			}

			suite.keeper.SavePost(suite.ctx, test.storedPost)

			err := suite.keeper.SavePostReaction(suite.ctx, test.postID, test.reaction)
			suite.Equal(test.error, err)

			var stored types.PostReactions
			suite.keeper.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(test.postID)), &stored)
			suite.Equal(test.expectedStored, stored)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_RemoveReaction() {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	liker, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	suite.NoError(err)

	tests := []struct {
		name           string
		storedLikes    types.PostReactions
		postID         types.PostID
		liker          sdk.AccAddress
		value          string
		shortcode      string
		error          error
		expectedStored types.PostReactions
	}{
		{
			name:           "PostReaction from same liker is removed properly",
			storedLikes:    types.PostReactions{types.NewPostReaction(":like:", "👍", liker)},
			postID:         id,
			liker:          liker,
			shortcode:      ":like:",
			value:          "👍",
			error:          nil,
			expectedStored: types.PostReactions{},
		},
		{
			name:           "Non existing reaction returns error - Creator",
			storedLikes:    types.PostReactions{},
			postID:         id,
			liker:          liker,
			shortcode:      ":like:",
			value:          "👍",
			error:          fmt.Errorf("cannot remove the reaction with value :like: from user cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4 as it does not exist"),
			expectedStored: types.PostReactions{},
		},
		{
			name:           "Non existing reaction returns error - Reaction",
			storedLikes:    types.PostReactions{types.NewPostReaction(":like:", "👍", liker)},
			postID:         id,
			liker:          liker,
			shortcode:      ":smile:",
			value:          "😊",
			error:          fmt.Errorf("cannot remove the reaction with value :smile: from user cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4 as it does not exist"),
			expectedStored: types.PostReactions{types.NewPostReaction(":like:", "👍", liker)},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {

			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			if len(test.storedLikes) != 0 {
				store.Set(types.PostReactionsStoreKey(test.postID), suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedLikes))
			}

			err := suite.keeper.RemovePostReaction(suite.ctx, test.postID, types.NewPostReaction(test.shortcode, test.value, test.liker))
			suite.Equal(test.error, err)

			var stored types.PostReactions
			suite.keeper.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(test.postID)), &stored)

			suite.Len(stored, len(test.expectedStored))
			for index, like := range test.expectedStored {
				suite.Equal(like, stored[index])
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetPostReactions() {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	liker, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	suite.NoError(err)

	otherLiker, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	suite.NoError(err)

	tests := []struct {
		name               string
		likes              types.PostReactions
		postID             types.PostID
		storedPost         types.Post
		registeredReaction types.Reaction
	}{
		{
			name:   "Empty list are returned properly",
			likes:  types.PostReactions{},
			postID: id,
		},
		{
			name: "Valid list of likes is returned properly",
			likes: types.PostReactions{
				types.NewPostReaction(":smile:", "😊", otherLiker),
				types.NewPostReaction(":smile:", "😊", liker),
			},
			postID:             id,
			storedPost:         testPost,
			registeredReaction: testRegisteredReaction,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			for _, l := range test.likes {
				suite.keeper.SavePost(suite.ctx, test.storedPost)
				suite.keeper.RegisterReaction(suite.ctx, test.registeredReaction)
				err := suite.keeper.SavePostReaction(suite.ctx, test.postID, l)
				suite.NoError(err)
			}

			stored := suite.keeper.GetPostReactions(suite.ctx, test.postID)

			suite.Len(stored, len(test.likes))
			for _, l := range test.likes {
				suite.Contains(stored, l)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetReactions() {
	id := "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"
	id2 := "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"
	liker1, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	suite.NoError(err)

	liker2, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	suite.NoError(err)

	tests := []struct {
		name  string
		likes map[string]types.PostReactions
	}{
		{
			name:  "Empty likes data are returned correctly",
			likes: map[string]types.PostReactions{},
		},
		{
			name: "Non empty likes data are returned correcly",
			likes: map[string]types.PostReactions{
				id: {
					types.NewPostReaction(":smile:", "😊", liker1),
					types.NewPostReaction(":smile:", "😊", liker2),
				},
				id2: {
					types.NewPostReaction(":smile:", "😊", liker1),
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			for postID, likes := range test.likes {
				store.Set(types.PostReactionsStoreKey(types.PostID(postID)), suite.keeper.Cdc.MustMarshalBinaryBare(likes))
			}

			likesData := suite.keeper.GetReactions(suite.ctx)
			suite.Equal(test.likes, likesData)
		})
	}
}

// -------------
// --- Reactions
// -------------

func (suite *KeeperTestSuite) TestKeeper_RegisterReaction() {
	var testOwner, _ = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	reaction := types.NewReaction(
		testOwner,
		":smile:",
		"https://smile.jpg",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)

	store := suite.ctx.KVStore(suite.keeper.StoreKey)
	key := types.ReactionsStoreKey(reaction.ShortCode, reaction.Subspace)

	suite.keeper.RegisterReaction(suite.ctx, reaction)

	var actualReaction types.Reaction

	bz := store.Get(key)
	suite.keeper.Cdc.MustUnmarshalBinaryBare(bz, &actualReaction)

	suite.Equal(reaction, actualReaction)
}

func (suite *KeeperTestSuite) TestKeeper_DoesReactionForShortcodeExist() {
	var testOwner, _ = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")

	reaction := types.NewReaction(
		testOwner,
		":smile:",
		"https://smile.jpg",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)

	tests := []struct {
		name           string
		storedReaction types.Reaction
		shortCode      string
		expBool        bool
	}{
		{
			name:           "reaction for given short code exists",
			storedReaction: reaction,
			shortCode:      ":smile:",
			expBool:        true,
		},
		{
			name:           "reaction for the given short code doesn't exist",
			storedReaction: reaction,
			shortCode:      ":test:",
			expBool:        false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			key := types.ReactionsStoreKey(reaction.ShortCode, reaction.Subspace)
			store.Set(key, suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedReaction))

			actualReaction, exist := suite.keeper.GetRegisteredReaction(suite.ctx, test.shortCode, reaction.Subspace)
			if test.shortCode == reaction.ShortCode {
				suite.True(exist)
				suite.Equal(test.storedReaction, actualReaction)
			} else {
				suite.False(exist)
				suite.Equal(types.Reaction{}, actualReaction)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_ListReactions() {
	store := suite.ctx.KVStore(suite.keeper.StoreKey)

	var testOwner, _ = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	reaction := types.NewReaction(
		testOwner,
		":smile:",
		"https://smile.jpg",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)
	reaction2 := types.NewReaction(
		testOwner,
		":thumbs_up:",
		"https://thumbs_up.jpg",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)
	reactions := types.Reactions{reaction, reaction2}

	for _, reaction := range reactions {
		key := types.ReactionsStoreKey(reaction.ShortCode, reaction.Subspace)
		store.Set(key, suite.keeper.Cdc.MustMarshalBinaryBare(reaction))
	}

	actualReactions := suite.keeper.GetRegisteredReactions(suite.ctx)

	suite.Equal(reactions, actualReactions)

}
