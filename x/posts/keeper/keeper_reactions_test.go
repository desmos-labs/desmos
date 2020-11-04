package keeper_test

import (
	"fmt"

	"github.com/desmos-labs/desmos/x/posts/keeper"

	"github.com/desmos-labs/desmos/x/posts/types"
)

func (suite *KeeperTestSuite) TestKeeper_SavePostReaction() {
	tests := []struct {
		name            string
		storedReactions []types.PostReaction
		storedPost      types.Post
		postID          string
		reaction        types.PostReaction
		error           error
		expectedStored  types.PostReactions
	}{
		{
			name: "Reaction from same user already present returns error",
			storedReactions: types.PostReactions{
				types.NewPostReaction(":like:", "👍", "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			},
			postID:   "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			reaction: types.NewPostReaction(":like:", "👍", "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			storedPost: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       suite.testData.post.ParentID,
				Message:        suite.testData.post.Message,
				Created:        suite.testData.post.Created,
				LastEdited:     suite.testData.post.LastEdited,
				AllowsComments: suite.testData.post.AllowsComments,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        suite.testData.post.Creator,
			},
			error: fmt.Errorf("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4 has already reacted with :like: to the post with id 19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"),
			expectedStored: types.PostReactions{
				types.NewPostReaction(":like:", "👍", "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			},
		},
		{
			name:            "First liker is stored properly",
			storedReactions: types.PostReactions{},
			postID:          "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			reaction:        types.NewPostReaction(":like:", "👍", "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			storedPost: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       suite.testData.post.ParentID,
				Message:        suite.testData.post.Message,
				Created:        suite.testData.post.Created,
				LastEdited:     suite.testData.post.LastEdited,
				AllowsComments: suite.testData.post.AllowsComments,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        suite.testData.post.Creator,
			},
			error: nil,
			expectedStored: types.PostReactions{
				types.NewPostReaction(":like:", "👍", "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			},
		},
		{
			name: "Second liker is stored properly",
			storedReactions: types.PostReactions{
				types.NewPostReaction(":like:", "👍", "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			},
			postID:   "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			reaction: types.NewPostReaction(":like:", "👍", "cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae"),
			storedPost: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       suite.testData.post.ParentID,
				Message:        suite.testData.post.Message,
				Created:        suite.testData.post.Created,
				LastEdited:     suite.testData.post.LastEdited,
				AllowsComments: suite.testData.post.AllowsComments,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        suite.testData.post.Creator,
			},
			error: nil,
			expectedStored: types.PostReactions{
				types.NewPostReaction(":like:", "👍", "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
				types.NewPostReaction(":like:", "👍", "cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae"),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() // reset
			store := suite.ctx.KVStore(suite.storeKey)
			if len(test.storedReactions) != 0 {
				wrapped := keeper.WrappedPostReactions{Reactions: test.storedReactions}
				store.Set(types.PostReactionsStoreKey(test.postID), suite.cdc.MustMarshalBinaryBare(&wrapped))
			}

			suite.k.SavePost(suite.ctx, test.storedPost)

			err := suite.k.SavePostReaction(suite.ctx, test.postID, test.reaction)
			suite.Require().Equal(test.error, err)

			var stored keeper.WrappedPostReactions
			suite.cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(test.postID)), &stored)
			suite.Require().Equal(test.expectedStored, stored.Reactions)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeletePostReaction() {
	tests := []struct {
		name            string
		storedReactions types.PostReactions
		postID          string
		liker           string
		value           string
		shortcode       string
		error           error
		expectedStored  types.PostReactions
	}{
		{
			name: "PostReaction from same liker is removed properly",
			storedReactions: types.PostReactions{
				types.NewPostReaction(":like:", "👍", "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			},
			postID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			liker:          "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			shortcode:      ":like:",
			value:          "👍",
			error:          nil,
			expectedStored: types.PostReactions{},
		},
		{
			name:            "Non existing reaction returns error - Creator",
			storedReactions: types.PostReactions{},
			postID:          "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			liker:           "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			shortcode:       ":like:",
			value:           "👍",
			error:           fmt.Errorf("cannot remove the reaction with value :like: from user cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4 as it does not exist"),
			expectedStored:  types.PostReactions{},
		},
		{
			name: "Non existing reaction returns error - Reaction",
			storedReactions: types.PostReactions{
				types.NewPostReaction(":like:", "👍", "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			},
			postID:    "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			liker:     "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			shortcode: ":smile:",
			value:     "😊",
			error:     fmt.Errorf("cannot remove the reaction with value :smile: from user cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4 as it does not exist"),
			expectedStored: types.PostReactions{
				types.NewPostReaction(":like:", "👍", "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {

			store := suite.ctx.KVStore(suite.storeKey)
			if len(test.storedReactions) != 0 {
				wrapped := keeper.WrappedPostReactions{Reactions: test.storedReactions}
				store.Set(types.PostReactionsStoreKey(test.postID), suite.cdc.MustMarshalBinaryBare(&wrapped))
			}

			err := suite.k.DeletePostReaction(suite.ctx, test.postID, types.NewPostReaction(test.shortcode, test.value, test.liker))
			suite.Require().Equal(test.error, err)

			var stored keeper.WrappedPostReactions
			suite.cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(test.postID)), &stored)

			suite.Len(stored, len(test.expectedStored))
			for index, like := range test.expectedStored {
				suite.Require().Equal(like, stored.Reactions[index])
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetPostReactions() {

	tests := []struct {
		name               string
		storedPost         types.Post
		registeredReaction types.RegisteredReaction
		reactions          types.PostReactions
		postID             string
	}{
		{
			name:      "Empty list are returned properly",
			reactions: types.PostReactions{},
			postID:    "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		},
		{
			name: "Valid list of reactions is returned properly",
			reactions: types.PostReactions{
				types.NewPostReaction(":smile:", "😊", "cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae"),
				types.NewPostReaction(":smile:", "😊", "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			},
			postID:             "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			storedPost:         suite.testData.post,
			registeredReaction: suite.testData.registeredReaction,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			for _, l := range test.reactions {
				suite.k.SavePost(suite.ctx, test.storedPost)
				suite.k.SaveRegisteredReaction(suite.ctx, test.registeredReaction)
				err := suite.k.SavePostReaction(suite.ctx, test.postID, l)
				suite.Require().NoError(err)
			}

			stored := suite.k.GetPostReactions(suite.ctx, test.postID)

			suite.Len(stored, len(test.reactions))
			for _, l := range test.reactions {
				suite.Contains(stored, l)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetPostReactionsEntries() {
	tests := []struct {
		name    string
		entries []types.PostReactionsEntry
	}{
		{
			name:    "Empty reactions data are returned correctly",
			entries: nil,
		},
		{
			name: "Non empty reactions data are returned correcly",
			entries: []types.PostReactionsEntry{
				types.NewPostReactionsEntry(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					[]types.PostReaction{
						types.NewPostReaction(
							":smile:",
							"😊",
							"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						),
						types.NewPostReaction(
							":smile:",
							"😊",
							"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
						),
					},
				),
				types.NewPostReactionsEntry(
					"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					[]types.PostReaction{
						types.NewPostReaction(
							":smile:",
							"😊",
							"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						),
					},
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.storeKey)
			for _, entry := range test.entries {
				wrapped := keeper.WrappedPostReactions{Reactions: entry.Reactions}
				store.Set(types.PostReactionsStoreKey(entry.PostId), suite.cdc.MustMarshalBinaryBare(&wrapped))
			}

			likesData := suite.k.GetPostReactionsEntries(suite.ctx)
			suite.Require().Equal(test.entries, likesData)
		})
	}
}

// ___________________________________________________________________________________________________________________

func (suite *KeeperTestSuite) TestKeeper_SaveRegisteredReaction() {
	reaction := types.NewRegisteredReaction(
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		":smile:",
		"https://smile.jpg",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)

	store := suite.ctx.KVStore(suite.storeKey)
	key := types.ReactionsStoreKey(reaction.ShortCode, reaction.Subspace)

	suite.k.SaveRegisteredReaction(suite.ctx, reaction)

	var actualReaction types.RegisteredReaction

	bz := store.Get(key)
	suite.cdc.MustUnmarshalBinaryBare(bz, &actualReaction)

	suite.Require().Equal(reaction, actualReaction)
}

func (suite *KeeperTestSuite) TestKeeper_GetRegisteredReaction() {
	reaction := types.NewRegisteredReaction(
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		":smile:",
		"https://smile.jpg",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)

	tests := []struct {
		name           string
		storedReaction types.RegisteredReaction
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
			store := suite.ctx.KVStore(suite.storeKey)
			key := types.ReactionsStoreKey(reaction.ShortCode, reaction.Subspace)
			store.Set(key, suite.cdc.MustMarshalBinaryBare(&test.storedReaction))

			actualReaction, exist := suite.k.GetRegisteredReaction(suite.ctx, test.shortCode, reaction.Subspace)
			if test.shortCode == reaction.ShortCode {
				suite.True(exist)
				suite.Require().Equal(test.storedReaction, actualReaction)
			} else {
				suite.Require().False(exist)
				suite.Require().Equal(types.RegisteredReaction{}, actualReaction)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetRegisteredReactions() {
	reactions := types.Reactions{
		types.NewRegisteredReaction(
			"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			":smile:",
			"https://smile.jpg",
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
		types.NewRegisteredReaction(
			"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			":thumbs_up:",
			"https://thumbs_up.jpg",
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
	}

	store := suite.ctx.KVStore(suite.storeKey)
	for _, reaction := range reactions {
		key := types.ReactionsStoreKey(reaction.ShortCode, reaction.Subspace)
		store.Set(key, suite.cdc.MustMarshalBinaryBare(&reaction))
	}

	actualReactions := suite.k.GetRegisteredReactions(suite.ctx)
	suite.Require().Equal(reactions, actualReactions)
}
