package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

func (suite *KeeperTestSuite) TestKeeper_SavePostReaction() {
	tests := []struct {
		name            string
		storedPosts     []types.Post
		storedReactions []types.PostReaction
		reaction        types.PostReaction
		expError        bool
		expectedStored  []types.PostReaction
	}{
		{
			name: "Reaction from same user already present returns error",
			storedPosts: []types.Post{
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
			storedReactions: []types.PostReaction{
				types.NewPostReaction(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					":like:",
					"👍",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
			reaction: types.NewPostReaction(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				":like:",
				"👍",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			),
			expError: true,
			expectedStored: []types.PostReaction{
				types.NewPostReaction(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					":like:",
					"👍",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
		},
		{
			name: "First liker is stored properly",
			storedPosts: []types.Post{
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
			reaction: types.NewPostReaction(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				":like:",
				"👍",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			),
			expError: false,
			expectedStored: []types.PostReaction{
				types.NewPostReaction(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					":like:",
					"👍",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
		},
		{
			name: "Second liker is stored properly",
			storedPosts: []types.Post{
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
			storedReactions: []types.PostReaction{
				types.NewPostReaction(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					":like:",
					"👍",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
			reaction: types.NewPostReaction(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				":like:",
				"👍",
				"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
			),
			expError: false,
			expectedStored: []types.PostReaction{
				types.NewPostReaction(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					":like:",
					"👍",
					"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
				),
				types.NewPostReaction(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					":like:",
					"👍",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
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

			for _, reaction := range test.storedReactions {
				err := suite.k.SavePostReaction(suite.ctx, reaction)
				suite.Require().NoError(err)
			}

			err := suite.k.SavePostReaction(suite.ctx, test.reaction)

			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expectedStored, suite.k.GetAllPostReactions(suite.ctx))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeletePostReaction() {
	tests := []struct {
		name            string
		storedReactions []types.PostReaction
		reaction        types.PostReaction
		expError        bool
		expReactions    []types.PostReaction
	}{
		{
			name: "Exiting reaction is removed properly",
			storedReactions: []types.PostReaction{
				types.NewPostReaction(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					":like:",
					"👍",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
			reaction: types.NewPostReaction(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				":like:",
				"👍",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			),
			expError:     false,
			expReactions: nil,
		},
		{
			name: "Non existing reaction returns error (different creator)",
			reaction: types.NewPostReaction(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				":like:",
				"👍",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			),
			expError: true,
		},
		{
			name: "Non existing reaction returns error (different reaction)",
			storedReactions: []types.PostReaction{
				types.NewPostReaction(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					":like:",
					"👍",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
			reaction: types.NewPostReaction(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				":smile:",
				"😊",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			),
			expError: true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, reaction := range test.storedReactions {
				err := suite.k.SavePostReaction(suite.ctx, reaction)
				suite.Require().NoError(err)
			}

			err := suite.k.DeletePostReaction(suite.ctx, test.reaction)

			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expReactions, suite.k.GetAllPostReactions(suite.ctx))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetPostReactions() {

	tests := []struct {
		name               string
		storedPost         types.Post
		registeredReaction types.RegisteredReaction
		reactions          []types.PostReaction
		postID             string
	}{
		{
			name:      "Empty list are returned properly",
			reactions: []types.PostReaction{},
			postID:    "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		},
		{
			name: "Valid list of reactions is returned properly",
			reactions: []types.PostReaction{
				types.NewPostReaction(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					":smile:",
					"😊",
					"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
				),
				types.NewPostReaction(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					":smile:",
					"😊",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
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
				err := suite.k.SavePostReaction(suite.ctx, l)
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

func (suite KeeperTestSuite) Test_GetPostReaction() {
	reactions := []types.PostReaction{
		types.NewPostReaction(
			"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			":smile:",
			"😊",
			"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
		),
	}

	tests := []struct {
		name        string
		store       func(ctx sdk.Context)
		postID      string
		owner       string
		shortCode   string
		shouldFound bool
		expResponse types.PostReaction
	}{
		{
			name:        "Empty list are returned properly",
			postID:      "",
			owner:       "",
			shortCode:   "",
			shouldFound: false,
		},
		{
			name: "Valid list of reactions is returned properly",
			store: func(ctx sdk.Context) {
				for _, r := range reactions {
					suite.k.SavePost(suite.ctx, suite.testData.post)
					err := suite.k.SavePostReaction(suite.ctx, r)
					suite.Require().NoError(err)
				}
			},
			postID:      "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			owner:       "cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
			shortCode:   ":smile:",
			shouldFound: true,
			expResponse: types.NewPostReaction(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				":smile:",
				"😊",
				"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
			),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if test.store != nil {
				test.store(ctx)
			}
			reaction, found := suite.k.GetPostReaction(ctx, test.postID, test.owner, test.shortCode)
			if !test.shouldFound {
				suite.Require().False(found)
			} else {
				suite.Require().True(found)
				suite.Require().Equal(test.expResponse, reaction)
			}
		})
	}
}

// ___________________________________________________________________________________________________________________

func (suite *KeeperTestSuite) TestKeeper_SaveRegisteredReaction() {
	tests := []struct {
		name            string
		storedReactions []types.RegisteredReaction
		toSave          types.RegisteredReaction
		expStored       []types.RegisteredReaction
	}{
		{
			name: "Already present reaction is overloaded",
			storedReactions: []types.RegisteredReaction{
				types.NewRegisteredReaction(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					":smile:",
					"https://smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			toSave: types.NewRegisteredReaction(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				":smile:",
				"SMILE!",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expStored: []types.RegisteredReaction{
				types.NewRegisteredReaction(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					":smile:",
					"SMILE!",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
		},
		{
			name: "Not present reaction is stored properly",
			toSave: types.NewRegisteredReaction(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				":smile:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expStored: []types.RegisteredReaction{
				types.NewRegisteredReaction(
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
		storedReactions []types.RegisteredReaction
		data            struct {
			shortCode string
			subspace  string
		}
		expActual types.RegisteredReaction
		expExist  bool
	}{
		{
			name: "registeredReactions for given short code exists",
			storedReactions: []types.RegisteredReaction{
				types.NewRegisteredReaction(
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
			expActual: types.NewRegisteredReaction(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				":smile:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
		},
		{
			name: "registeredReactions for the given short code doesn't exist",
			storedReactions: []types.RegisteredReaction{
				types.NewRegisteredReaction(
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
		storedReactions []types.RegisteredReaction
		expected        []types.RegisteredReaction
	}{
		{
			name:            "Empty slice is returned properly",
			storedReactions: nil,
			expected:        nil,
		},
		{
			name: "Non empty slice is returned properly",
			storedReactions: []types.RegisteredReaction{
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
			},
			expected: []types.RegisteredReaction{
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
