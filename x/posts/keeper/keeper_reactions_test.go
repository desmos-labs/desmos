package keeper_test

import (
	types2 "github.com/desmos-labs/desmos/x/posts/types"
)

func (suite *KeeperTestSuite) TestKeeper_SavePostReaction() {
	tests := []struct {
		name            string
		storedPosts     []types2.Post
		storedReactions []types2.PostReactionsEntry
		postID          string
		reaction        types2.PostReaction
		expError        bool
		expectedStored  []types2.PostReactionsEntry
	}{
		{
			name: "Reaction from same user already present returns error",
			storedPosts: []types2.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					ParentID:             suite.testData.post.ParentID,
					Message:              suite.testData.post.Message,
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					CommentsState:        suite.testData.post.CommentsState,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
				},
			},
			storedReactions: []types2.PostReactionsEntry{
				types2.NewPostReactionsEntry(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					[]types2.PostReaction{
						types2.NewPostReaction(
							":like:",
							"👍",
							"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						),
					},
				),
			},
			postID:   "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			reaction: types2.NewPostReaction(":like:", "👍", "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			expError: true,
			expectedStored: []types2.PostReactionsEntry{
				types2.NewPostReactionsEntry(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					[]types2.PostReaction{
						types2.NewPostReaction(
							":like:",
							"👍",
							"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						),
					}),
			},
		},
		{
			name: "First liker is stored properly",
			storedPosts: []types2.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					ParentID:             suite.testData.post.ParentID,
					Message:              suite.testData.post.Message,
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					CommentsState:        suite.testData.post.CommentsState,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
				},
			},
			postID:   "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			reaction: types2.NewPostReaction(":like:", "👍", "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			expError: false,
			expectedStored: []types2.PostReactionsEntry{
				types2.NewPostReactionsEntry(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					[]types2.PostReaction{
						types2.NewPostReaction(
							":like:",
							"👍",
							"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						),
					}),
			},
		},
		{
			name: "Second liker is stored properly",
			storedPosts: []types2.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					ParentID:             suite.testData.post.ParentID,
					Message:              suite.testData.post.Message,
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					CommentsState:        suite.testData.post.CommentsState,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
				},
			},
			storedReactions: []types2.PostReactionsEntry{
				types2.NewPostReactionsEntry(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					[]types2.PostReaction{
						types2.NewPostReaction(
							":like:",
							"👍",
							"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						),
					},
				),
			},
			postID:   "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			reaction: types2.NewPostReaction(":like:", "👍", "cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae"),
			expError: false,
			expectedStored: []types2.PostReactionsEntry{
				types2.NewPostReactionsEntry(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					[]types2.PostReaction{
						types2.NewPostReaction(
							":like:",
							"👍",
							"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						),
						types2.NewPostReaction(
							":like:",
							"👍",
							"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
						),
					}),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, post := range test.storedPosts {
				suite.k.SavePost(suite.ctx, post)
			}

			for _, entry := range test.storedReactions {
				for _, reaction := range entry.Reactions {
					err := suite.k.SavePostReaction(suite.ctx, entry.PostID, reaction)
					suite.Require().NoError(err)
				}
			}

			err := suite.k.SavePostReaction(suite.ctx, test.postID, test.reaction)

			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expectedStored, suite.k.GetPostReactionsEntries(suite.ctx))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeletePostReaction() {
	tests := []struct {
		name            string
		storedReactions []types2.PostReactionsEntry
		data            struct {
			postID   string
			reaction types2.PostReaction
		}
		expError     bool
		expReactions []types2.PostReactionsEntry
	}{
		{
			name: "Exiting reaction is removed properly",
			storedReactions: []types2.PostReactionsEntry{
				types2.NewPostReactionsEntry(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					[]types2.PostReaction{
						types2.NewPostReaction(
							":like:",
							"👍",
							"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						),
					},
				),
			},
			data: struct {
				postID   string
				reaction types2.PostReaction
			}{
				postID: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				reaction: types2.NewPostReaction(
					":like:",
					"👍",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
			expError:     false,
			expReactions: nil,
		},
		{
			name: "Non existing reaction returns error (different creator)",
			data: struct {
				postID   string
				reaction types2.PostReaction
			}{
				postID: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				reaction: types2.NewPostReaction(
					":like:",
					"👍",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
			expError: true,
		},
		{
			name: "Non existing reaction returns error (different reaction)",
			storedReactions: []types2.PostReactionsEntry{
				types2.NewPostReactionsEntry(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					[]types2.PostReaction{
						types2.NewPostReaction(
							":like:",
							"👍",
							"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						),
					}),
			},
			data: struct {
				postID   string
				reaction types2.PostReaction
			}{
				postID: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				reaction: types2.NewPostReaction(
					":smile:",
					"😊",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
			expError: true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, entry := range test.storedReactions {
				for _, reaction := range entry.Reactions {
					err := suite.k.SavePostReaction(suite.ctx, entry.PostID, reaction)
					suite.Require().NoError(err)
				}
			}

			err := suite.k.DeletePostReaction(suite.ctx, test.data.postID, test.data.reaction)

			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expReactions, suite.k.GetPostReactionsEntries(suite.ctx))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetPostReactions() {

	tests := []struct {
		name               string
		storedPost         types2.Post
		registeredReaction types2.RegisteredReaction
		reactions          []types2.PostReaction
		postID             string
	}{
		{
			name:      "Empty list are returned properly",
			reactions: []types2.PostReaction{},
			postID:    "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		},
		{
			name: "Valid list of reactions is returned properly",
			reactions: []types2.PostReaction{
				types2.NewPostReaction(":smile:", "😊", "cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae"),
				types2.NewPostReaction(":smile:", "😊", "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
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
		entries []types2.PostReactionsEntry
	}{
		{
			name:    "Empty reactions data are returned correctly",
			entries: nil,
		},
		{
			name: "Non empty reactions data are returned correcly",
			entries: []types2.PostReactionsEntry{
				types2.NewPostReactionsEntry(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					[]types2.PostReaction{
						types2.NewPostReaction(
							":smile:",
							"😊",
							"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						),
						types2.NewPostReaction(
							":smile:",
							"😊",
							"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
						),
					},
				),
				types2.NewPostReactionsEntry(
					"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					[]types2.PostReaction{
						types2.NewPostReaction(
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
				wrapped := types2.PostReactions{Reactions: entry.Reactions}
				store.Set(types2.PostReactionsStoreKey(entry.PostID), suite.cdc.MustMarshalBinaryBare(&wrapped))
			}

			likesData := suite.k.GetPostReactionsEntries(suite.ctx)
			suite.Require().Equal(test.entries, likesData)
		})
	}
}

// ___________________________________________________________________________________________________________________

func (suite *KeeperTestSuite) TestKeeper_SaveRegisteredReaction() {
	tests := []struct {
		name            string
		storedReactions []types2.RegisteredReaction
		toSave          types2.RegisteredReaction
		expStored       []types2.RegisteredReaction
	}{
		{
			name: "Already present reaction is overloaded",
			storedReactions: []types2.RegisteredReaction{
				types2.NewRegisteredReaction(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					":smile:",
					"https://smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			toSave: types2.NewRegisteredReaction(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				":smile:",
				"SMILE!",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expStored: []types2.RegisteredReaction{
				types2.NewRegisteredReaction(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					":smile:",
					"SMILE!",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
		},
		{
			name: "Not present reaction is stored properly",
			toSave: types2.NewRegisteredReaction(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				":smile:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expStored: []types2.RegisteredReaction{
				types2.NewRegisteredReaction(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					":smile:",
					"https://smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, reaction := range test.storedReactions {
				suite.k.SaveRegisteredReaction(suite.ctx, reaction)
			}

			suite.k.SaveRegisteredReaction(suite.ctx, test.toSave)

			suite.Require().Equal(test.expStored, suite.k.GetRegisteredReactions(suite.ctx))
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetRegisteredReaction() {
	tests := []struct {
		name            string
		storedReactions []types2.RegisteredReaction
		data            struct {
			shortCode string
			subspace  string
		}
		expActual types2.RegisteredReaction
		expExist  bool
	}{
		{
			name: "registeredReactions for given short code exists",
			storedReactions: []types2.RegisteredReaction{
				types2.NewRegisteredReaction(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					":smile:",
					"https://smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			data: struct {
				shortCode string
				subspace  string
			}{
				shortCode: ":smile:",
				subspace:  "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			},
			expExist: true,
			expActual: types2.NewRegisteredReaction(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				":smile:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
		},
		{
			name: "registeredReactions for the given short code doesn't exist",
			storedReactions: []types2.RegisteredReaction{
				types2.NewRegisteredReaction(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					":smile:",
					"https://smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			data: struct {
				shortCode string
				subspace  string
			}{
				shortCode: ":test:",
				subspace:  "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			},
			expExist: false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, reaction := range test.storedReactions {
				suite.k.SaveRegisteredReaction(suite.ctx, reaction)
			}

			actual, exists := suite.k.GetRegisteredReaction(suite.ctx, test.data.shortCode, test.data.subspace)

			if test.expExist {
				suite.Require().True(exists)
				suite.Require().Equal(test.expActual, actual)
			} else {
				suite.Require().False(exists)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetRegisteredReactions() {
	tests := []struct {
		name            string
		storedReactions []types2.RegisteredReaction
		expected        []types2.RegisteredReaction
	}{
		{
			name:            "Empty slice is returned properly",
			storedReactions: nil,
			expected:        nil,
		},
		{
			name: "Non empty slice is returned properly",
			storedReactions: []types2.RegisteredReaction{
				types2.NewRegisteredReaction(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					":smile:",
					"https://smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types2.NewRegisteredReaction(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					":thumbs_up:",
					"https://thumbs_up.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			expected: []types2.RegisteredReaction{
				types2.NewRegisteredReaction(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					":smile:",
					"https://smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types2.NewRegisteredReaction(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					":thumbs_up:",
					"https://thumbs_up.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, reaction := range test.storedReactions {
				suite.k.SaveRegisteredReaction(suite.ctx, reaction)
			}

			stored := suite.k.GetRegisteredReactions(suite.ctx)
			suite.Require().Equal(test.expected, stored)
		})
	}
}
