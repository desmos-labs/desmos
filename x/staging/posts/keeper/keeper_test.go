package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

// getPostsCommentsIDs returns the comments of the post having the given id
func (suite *KeeperTestSuite) getPostsCommentsIDs(ctx sdk.Context, postID string) []string {
	store := ctx.KVStore(suite.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PostCommentsPrefix(postID))
	defer iterator.Close()

	var commentIDs []string
	for ; iterator.Valid(); iterator.Next() {
		commentID := string(iterator.Value())
		commentIDs = append(commentIDs, commentID)
	}
	return commentIDs
}

func (suite *KeeperTestSuite) TestKeeper_SavePost() {
	tests := []struct {
		name           string
		existingPosts  []types.Post
		newPost        types.Post
		expCommentsIDs []string
		expLastID      string
	}{
		{
			name: "Post with ID already present",
			existingPosts: []types.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Post",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
				},
			},
			newPost: types.Post{
				PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Message:              "New post",
				Created:              suite.testData.post.Created,
				LastEdited:           suite.testData.post.LastEdited,
				Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				AdditionalAttributes: nil,
				Creator:              suite.testData.post.Creator,
			},
			expCommentsIDs: []string{},
		},
		{
			name: "Post which ID is not already present",
			existingPosts: []types.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Post",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
				},
			},
			newPost: types.Post{
				PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Message:              "New post",
				Created:              suite.testData.post.Created,
				LastEdited:           suite.testData.post.LastEdited,
				Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				AdditionalAttributes: nil,
				Creator:              suite.testData.post.Creator,
			},
			expCommentsIDs: []string{},
		},
		{
			name: "Post with valid parent ID",
			existingPosts: []types.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Parent",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
				},
			},
			newPost: types.Post{
				PostID:               "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				ParentID:             "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Message:              "Comment",
				Created:              suite.testData.post.Created,
				LastEdited:           suite.testData.post.LastEdited,
				Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				AdditionalAttributes: nil,
				Creator:              suite.testData.post.Creator,
			},
			expCommentsIDs: []string{
				"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
			},
		},
		{
			name: "Post with ID greater ID than Last ID stored",
			existingPosts: []types.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Post lesser",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.postOwner,
				},
			},
			newPost: types.Post{
				PostID:     "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Message:    "New post greater",
				Created:    suite.testData.post.Created,
				LastEdited: suite.testData.post.LastEdited,
				Subspace:   "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				AdditionalAttributes: []types.Attribute{
					types.NewAttribute("key", "value"),
				},
				Creator: suite.testData.postOwner,
			},
			expCommentsIDs: []string{},
		},
		{
			name: "Post with ID lesser ID than Last ID stored",
			existingPosts: []types.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Post ID greater",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.postOwner,
				},
			},
			newPost: types.Post{
				PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Message:              "New post ID lesser",
				Created:              suite.testData.post.Created,
				LastEdited:           suite.testData.post.LastEdited,
				Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				AdditionalAttributes: nil,
				Creator:              suite.testData.postOwner,
			},
			expCommentsIDs: []string{},
		},
		{
			name:          "Post with medias is saved properly",
			existingPosts: []types.Post{},
			newPost: types.Post{
				PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Message:              "Post with medias",
				Created:              suite.testData.post.Created,
				LastEdited:           suite.testData.post.LastEdited,
				Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				AdditionalAttributes: nil,
				Creator:              suite.testData.postOwner,
				Attachments:          suite.testData.post.Attachments,
			},
			expCommentsIDs: []string{},
		},
		{
			name:          "Post with poll data is saved properly",
			existingPosts: []types.Post{},
			newPost: types.Post{
				PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Message:              "Post with poll data",
				Created:              suite.testData.post.Created,
				LastEdited:           suite.testData.post.LastEdited,
				Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				AdditionalAttributes: nil,
				Creator:              suite.testData.postOwner,
				PollData:             suite.testData.post.PollData,
			},
			expCommentsIDs: []string{},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			store := suite.ctx.KVStore(suite.storeKey)
			for _, post := range test.existingPosts {
				store.Set(types.PostStoreKey(post.PostID), suite.cdc.MustMarshalBinaryBare(&post))
			}

			// Save the post
			suite.k.SavePost(suite.ctx, test.newPost)

			// Check the stored post
			var expected types.Post
			suite.cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(test.newPost.PostID)), &expected)
			suite.True(expected.Equal(test.newPost))

			// Check the post comments
			ids := suite.getPostsCommentsIDs(suite.ctx, test.newPost.ParentID)
			suite.Equal(test.expCommentsIDs, ids)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetPost() {
	tests := []struct {
		name       string
		postExists bool
		ID         string
		expected   types.Post
	}{
		{
			name:     "Non existent post is not found",
			ID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			expected: types.Post{},
		},
		{
			name:       "Existing post is found properly",
			ID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			postExists: true,
			expected: types.Post{
				PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Message:              "Post",
				Created:              suite.testData.post.Created,
				LastEdited:           suite.testData.post.LastEdited,
				Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				AdditionalAttributes: nil,
				Creator:              suite.testData.postOwner,
			},
		},
		{
			name:       "Existing post with medias is found properly",
			ID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			postExists: true,
			expected: types.Post{
				PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Message:              "Post",
				Created:              suite.testData.post.Created,
				LastEdited:           suite.testData.post.LastEdited,
				Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				AdditionalAttributes: nil,
				Creator:              suite.testData.postOwner,
				Attachments:          suite.testData.post.Attachments,
			},
		},
		{
			name:       "Existing post with poll is found properly",
			ID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			postExists: true,
			expected: types.Post{
				PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Message:              "Post",
				Created:              suite.testData.post.Created,
				LastEdited:           suite.testData.post.LastEdited,
				Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				AdditionalAttributes: nil,
				Creator:              suite.testData.postOwner,
				PollData:             suite.testData.post.PollData,
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.storeKey)

			if test.postExists {
				store.Set(types.PostStoreKey(test.expected.PostID), suite.cdc.MustMarshalBinaryBare(&test.expected))
			}

			expected, found := suite.k.GetPost(suite.ctx, test.ID)
			suite.Require().Equal(test.postExists, found)
			if test.postExists {
				suite.True(expected.Equal(test.expected))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetPostChildrenIDs() {
	tests := []struct {
		name           string
		storedPosts    []types.Post
		postID         string
		expChildrenIDs []string
	}{
		{
			name:           "Empty children list is returned properly",
			postID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			expChildrenIDs: nil,
		},
		{
			name: "Non empty children list is returned properly",
			storedPosts: []types.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Original post",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
				},
				{
					PostID:               "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					ParentID:             "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "First commit",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
				},
				{
					PostID:               "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					Message:              "Second post",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
				},
				{
					PostID:               "a33e173b6b96129f74acf41b5219a6bbc9f90e9e41f37115f1ce7f1f5860211c",
					ParentID:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					Message:              "Comment to second post",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator},
			},
			postID: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			expChildrenIDs: []string{
				"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			for _, p := range test.storedPosts {
				suite.k.SavePost(suite.ctx, p)
			}

			storedChildrenIDs := suite.getPostsCommentsIDs(suite.ctx, test.postID)
			suite.Len(storedChildrenIDs, len(test.expChildrenIDs))

			for _, id := range test.expChildrenIDs {
				suite.Contains(storedChildrenIDs, id)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetPosts() {
	tests := []struct {
		name  string
		posts []types.Post
	}{
		{
			name:  "Empty list returns empty list",
			posts: []types.Post{},
		},
		{
			name: "Existing list is returned properly",
			posts: []types.Post{
				{
					PostID:               "63b173547f1079e46885aa3ad4e36d0fe4beea8b7e2ec9c1d71ba3bff1abd909",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.postOwner,
				},
				{
					PostID:               "aad15654d10acd67b942ca39afd7a2aa071aed7c3f0b946edd2b666a037026f7",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.postOwner,
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			for _, p := range test.posts {
				suite.k.SavePost(suite.ctx, p)
			}

			posts := suite.k.GetPosts(suite.ctx)
			for index, post := range test.posts {
				suite.True(post.Equal(posts[index]))
			}
		})
	}
}
