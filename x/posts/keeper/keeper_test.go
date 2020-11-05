package keeper_test

import (
	"time"

	"github.com/desmos-labs/desmos/x/posts/types"
)

func (suite *KeeperTestSuite) TestKeeper_SavePost() {
	tests := []struct {
		name                 string
		existingPosts        []types.Post
		newPost              types.Post
		expParentCommentsIDs []string
		expLastID            string
	}{
		{
			name: "Post with ID already present",
			existingPosts: []types.Post{
				{
					PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:      "Post",
					Created:      suite.testData.post.Created,
					LastEdited:   suite.testData.post.LastEdited,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: nil,
					Creator:      suite.testData.post.Creator,
				},
			},
			newPost: types.Post{
				PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Message:      "New post",
				Created:      suite.testData.post.Created,
				LastEdited:   suite.testData.post.LastEdited,
				Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: nil,
				Creator:      suite.testData.post.Creator,
			},
		},
		{
			name: "Post which ID is not already present",
			existingPosts: []types.Post{
				{
					PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:      "Post",
					Created:      suite.testData.post.Created,
					LastEdited:   suite.testData.post.LastEdited,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: nil,
					Creator:      suite.testData.post.Creator,
				},
			},
			newPost: types.Post{
				PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Message:      "New post",
				Created:      suite.testData.post.Created,
				LastEdited:   suite.testData.post.LastEdited,
				Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: nil,
				Creator:      suite.testData.post.Creator,
			},
		},
		{
			name: "Post with valid parent ID",
			existingPosts: []types.Post{
				{
					PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:      "Parent",
					Created:      suite.testData.post.Created,
					LastEdited:   suite.testData.post.LastEdited,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: nil,
					Creator:      suite.testData.post.Creator,
				},
			},
			newPost: types.Post{
				PostID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				ParentID:     "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Message:      "Comment",
				Created:      suite.testData.post.Created,
				LastEdited:   suite.testData.post.LastEdited,
				Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: nil,
				Creator:      suite.testData.post.Creator,
			},
			expParentCommentsIDs: []string{
				"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
			},
		},
		{
			name: "Post with ID greater ID than Last ID stored",
			existingPosts: []types.Post{
				{
					PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:      "Post lesser",
					Created:      suite.testData.post.Created,
					LastEdited:   suite.testData.post.LastEdited,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: nil,
					Creator:      suite.testData.postOwner,
				},
			},
			newPost: types.Post{
				PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Message:      "New post greater",
				Created:      suite.testData.post.Created,
				LastEdited:   suite.testData.post.LastEdited,
				Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: []types.OptionalDataEntry{{"key", "value"}},
				Creator:      suite.testData.postOwner,
			},
		},
		{
			name: "Post with ID lesser ID than Last ID stored",
			existingPosts: []types.Post{
				{
					PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:      "Post ID greater",
					Created:      suite.testData.post.Created,
					LastEdited:   suite.testData.post.LastEdited,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: nil,
					Creator:      suite.testData.postOwner,
				},
			},
			newPost: types.Post{
				PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Message:      "New post ID lesser",
				Created:      suite.testData.post.Created,
				LastEdited:   suite.testData.post.LastEdited,
				Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: nil,
				Creator:      suite.testData.postOwner,
			},
		},
		{
			name:          "Post with medias is saved properly",
			existingPosts: []types.Post{},
			newPost: types.Post{
				PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Message:      "Post with medias",
				Created:      suite.testData.post.Created,
				LastEdited:   suite.testData.post.LastEdited,
				Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: nil,
				Creator:      suite.testData.postOwner,
				Attachments:  suite.testData.post.Attachments,
			},
		},
		{
			name:          "Post with poll data is saved properly",
			existingPosts: []types.Post{},
			newPost: types.Post{
				PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Message:      "Post with poll data",
				Created:      suite.testData.post.Created,
				LastEdited:   suite.testData.post.LastEdited,
				Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: nil,
				Creator:      suite.testData.postOwner,
				PollData:     suite.testData.post.PollData,
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.storeKey)
			for _, post := range test.existingPosts {
				store.Set(types.PostStoreKey(post.PostID), suite.cdc.MustMarshalBinaryBare(&post))
			}

			// Save the post
			suite.keeper.SavePost(suite.ctx, test.newPost)

			// Check the stored post
			var expected types.Post
			suite.cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(test.newPost.PostID)), &expected)
			suite.True(expected.Equal(test.newPost))

			// Check the parent comments
			var wrapped types.CommentIDs
			suite.cdc.MustUnmarshalBinaryBare(store.Get(types.PostCommentsStoreKey(test.newPost.ParentID)), &wrapped)
			suite.Equal(test.expParentCommentsIDs, wrapped.Ids)
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
				PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Message:      "Post",
				Created:      suite.testData.post.Created,
				LastEdited:   suite.testData.post.LastEdited,
				Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: nil,
				Creator:      suite.testData.postOwner,
			},
		},
		{
			name:       "Existing post with medias is found properly",
			ID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			postExists: true,
			expected: types.Post{
				PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Message:      "Post",
				Created:      suite.testData.post.Created,
				LastEdited:   suite.testData.post.LastEdited,
				Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: nil,
				Creator:      suite.testData.postOwner,
				Attachments:  suite.testData.post.Attachments,
			},
		},
		{
			name:       "Existing post with poll is found properly",
			ID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			postExists: true,
			expected: types.Post{
				PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Message:      "Post",
				Created:      suite.testData.post.Created,
				LastEdited:   suite.testData.post.LastEdited,
				Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: nil,
				Creator:      suite.testData.postOwner,
				PollData:     suite.testData.post.PollData,
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

			expected, found := suite.keeper.GetPost(suite.ctx, test.ID)
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
					PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:      "Original post",
					Created:      suite.testData.post.Created,
					LastEdited:   suite.testData.post.LastEdited,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: nil,
					Creator:      suite.testData.post.Creator,
				},
				{
					PostID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					ParentID:     "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:      "First commit",
					Created:      suite.testData.post.Created,
					LastEdited:   suite.testData.post.LastEdited,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: nil,
					Creator:      suite.testData.post.Creator,
				},
				{
					PostID:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					Message:      "Second post",
					Created:      suite.testData.post.Created,
					LastEdited:   suite.testData.post.LastEdited,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: nil,
					Creator:      suite.testData.post.Creator,
				},
				{
					PostID:       "a33e173b6b96129f74acf41b5219a6bbc9f90e9e41f37115f1ce7f1f5860211c",
					ParentID:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					Message:      "Comment to second post",
					Created:      suite.testData.post.Created,
					LastEdited:   suite.testData.post.LastEdited,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: nil,
					Creator:      suite.testData.post.Creator},
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
				suite.keeper.SavePost(suite.ctx, p)
			}

			storedChildrenIDs := suite.keeper.GetPostChildrenIDs(suite.ctx, test.postID)
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
					PostID:       "63b173547f1079e46885aa3ad4e36d0fe4beea8b7e2ec9c1d71ba3bff1abd909",
					Created:      suite.testData.post.Created,
					LastEdited:   suite.testData.post.LastEdited,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: nil,
					Creator:      suite.testData.postOwner,
				},
				{
					PostID:       "aad15654d10acd67b942ca39afd7a2aa071aed7c3f0b946edd2b666a037026f7",
					Created:      suite.testData.post.Created,
					LastEdited:   suite.testData.post.LastEdited,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: nil,
					Creator:      suite.testData.postOwner,
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			for _, p := range test.posts {
				suite.keeper.SavePost(suite.ctx, p)
			}

			posts := suite.keeper.GetPosts(suite.ctx)
			for index, post := range test.posts {
				suite.True(post.Equal(posts[index]))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetPostsFiltered() {
	date := time.Date(2020, 1, 1, 1, 1, 0, 0, time.UTC)
	posts := []types.Post{
		{
			PostID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
			ParentID:     "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			Message:      "Post 1 #test #desmos",
			Created:      date,
			OptionalData: nil,
			Creator:      "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
		},
		{
			PostID:         "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			ParentID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			Message:        "Post 2",
			Created:        time.Date(2020, 2, 1, 1, 1, 0, 0, time.UTC),
			AllowsComments: true,
			Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			OptionalData:   nil,
			Creator:        "cosmos1jlhazemxvu0zn9y77j6afwmpf60zveqw5480l2",
		},
		{
			PostID:       "a33e173b6b96129f74acf41b5219a6bbc9f90e9e41f37115f1ce7f1f5860211c",
			ParentID:     "84a5d9fc5f0acd2bb9c0a49ecaefabbe4698372e1ae88d32f9f6f80b3c0ab95e",
			Message:      "Post 3",
			Created:      time.Date(2020, 3, 1, 1, 1, 0, 0, time.UTC),
			Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			OptionalData: nil,
			Creator:      "cosmos1jlhazemxvu0zn9y77j6afwmpf60zveqw5480l2",
		},
	}

	tests := []struct {
		name     string
		filter   types.QueryPostsParams
		expected []types.Post
	}{
		{
			name:     "Valid pagination works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 2},
			expected: []types.Post{posts[1], posts[2]},
		},
		{
			name:     "Non existing page returns empty list",
			filter:   types.QueryPostsParams{Page: 10, Limit: 1},
			expected: []types.Post{},
		},
		{
			name:     "Invalid pagination returns all data",
			filter:   types.QueryPostsParams{Page: 1, Limit: 0},
			expected: []types.Post{posts[1], posts[2], posts[0]},
		},
		{
			name:     "Parent ID matcher works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, ParentID: posts[0].ParentID},
			expected: []types.Post{posts[1], posts[0]},
		},
		{
			name:     "Creation time matcher works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, CreationTime: &date},
			expected: []types.Post{posts[0]},
		},
		{
			name:     "Subspace mather works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, Subspace: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"},
			expected: []types.Post{posts[1], posts[2]},
		},
		{
			name:     "Creator mather works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, Creator: "cosmos1jlhazemxvu0zn9y77j6afwmpf60zveqw5480l2"},
			expected: []types.Post{posts[1], posts[2]},
		},
		{
			name:     "Sorting by date ascending works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, SortBy: types.PostSortByCreationDate, SortOrder: types.PostSortOrderAscending},
			expected: []types.Post{posts[0], posts[1], posts[2]},
		},
		{
			name:     "Sorting by date descending works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, SortBy: types.PostSortByCreationDate, SortOrder: types.PostSortOrderDescending},
			expected: []types.Post{posts[2], posts[1], posts[0]},
		},
		{
			name:     "Sorting by ID ascending works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, SortBy: types.PostSortByID, SortOrder: types.PostSortOrderAscending},
			expected: []types.Post{posts[1], posts[2], posts[0]},
		},
		{
			name:     "Sorting by ID descending works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, SortBy: types.PostSortByID, SortOrder: types.PostSortOrderDescending},
			expected: []types.Post{posts[0], posts[2], posts[1]},
		},
		{
			name:     "Filtering by hashtags works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, Hashtags: []string{"desmos", "test"}},
			expected: []types.Post{posts[0]},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			for _, post := range posts {
				suite.keeper.SavePost(suite.ctx, post)
			}
			result := suite.keeper.GetPostsFiltered(suite.ctx, test.filter)

			suite.Len(result, len(test.expected))
			for index, post := range result {
				suite.True(test.expected[index].Equal(post))
			}
		})
	}
}
