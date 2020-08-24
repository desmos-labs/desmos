package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/posts/types"
)

// -------------
// --- Posts
// -------------

func (suite *KeeperTestSuite) TestKeeper_SavePost() {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	id2 := types.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd")
	tests := []struct {
		name                 string
		existingPosts        types.Posts
		newPost              types.Post
		expParentCommentsIDs types.PostIDs
		expLastID            types.PostID
	}{
		{
			name: "Post with ID already present",
			existingPosts: types.Posts{
				types.Post{
					PostID:       id,
					Message:      "Post",
					Created:      suite.testData.post.Created,
					LastEdited:   suite.testData.post.LastEdited,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: map[string]string{},
					Creator:      suite.testData.post.Creator,
				},
			},
			newPost: types.Post{
				PostID:       id,
				Message:      "New post",
				Created:      suite.testData.post.Created,
				LastEdited:   suite.testData.post.LastEdited,
				Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: map[string]string{},
				Creator:      suite.testData.post.Creator,
			},
			expParentCommentsIDs: []types.PostID{},
		},
		{
			name: "Post which ID is not already present",
			existingPosts: types.Posts{
				types.Post{
					PostID:       id,
					Message:      "Post",
					Created:      suite.testData.post.Created,
					LastEdited:   suite.testData.post.LastEdited,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: map[string]string{},
					Creator:      suite.testData.post.Creator,
				},
			},
			newPost: types.Post{
				PostID:       id,
				Message:      "New post",
				Created:      suite.testData.post.Created,
				LastEdited:   suite.testData.post.LastEdited,
				Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: map[string]string{},
				Creator:      suite.testData.post.Creator,
			},
			expParentCommentsIDs: []types.PostID{},
		},
		{
			name: "Post with valid parent ID",
			existingPosts: []types.Post{
				{
					PostID:       id,
					Message:      "Parent",
					Created:      suite.testData.post.Created,
					LastEdited:   suite.testData.post.LastEdited,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: map[string]string{},
					Creator:      suite.testData.post.Creator,
				},
			},
			newPost: types.Post{
				PostID:       id2,
				ParentID:     id,
				Message:      "Comment",
				Created:      suite.testData.post.Created,
				LastEdited:   suite.testData.post.LastEdited,
				Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: map[string]string{},
				Creator:      suite.testData.post.Creator,
			},
			expParentCommentsIDs: []types.PostID{id2},
		},
		{
			name: "Post with ID greater ID than Last ID stored",
			existingPosts: types.Posts{
				types.Post{
					PostID:       id,
					Message:      "Post lesser",
					Created:      suite.testData.post.Created,
					LastEdited:   suite.testData.post.LastEdited,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: map[string]string{},
					Creator:      suite.testData.postOwner,
				},
			},
			newPost: types.Post{
				PostID:       id,
				Message:      "New post greater",
				Created:      suite.testData.post.Created,
				LastEdited:   suite.testData.post.LastEdited,
				Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: map[string]string{"key": "value"},
				Creator:      suite.testData.postOwner,
			},
			expParentCommentsIDs: []types.PostID{},
		},
		{
			name: "Post with ID lesser ID than Last ID stored",
			existingPosts: types.Posts{
				types.Post{
					PostID:       id,
					Message:      "Post ID greater",
					Created:      suite.testData.post.Created,
					LastEdited:   suite.testData.post.LastEdited,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: map[string]string{},
					Creator:      suite.testData.postOwner,
				},
			},
			newPost: types.Post{
				PostID:       id,
				Message:      "New post ID lesser",
				Created:      suite.testData.post.Created,
				LastEdited:   suite.testData.post.LastEdited,
				Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: map[string]string{},
				Creator:      suite.testData.postOwner,
			},
			expParentCommentsIDs: []types.PostID{},
		},
		{
			name:          "Post with medias is saved properly",
			existingPosts: types.Posts{},
			newPost: types.Post{
				PostID:       id,
				Message:      "Post with medias",
				Created:      suite.testData.post.Created,
				LastEdited:   suite.testData.post.LastEdited,
				Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: map[string]string{},
				Creator:      suite.testData.postOwner,
				Attachments:  suite.testData.post.Attachments,
			},
			expParentCommentsIDs: []types.PostID{},
		},
		{
			name:          "Post with poll data is saved properly",
			existingPosts: types.Posts{},
			newPost: types.Post{
				PostID:       id,
				Message:      "Post with poll data",
				Created:      suite.testData.post.Created,
				LastEdited:   suite.testData.post.LastEdited,
				Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: map[string]string{},
				Creator:      suite.testData.postOwner,
				PollData:     suite.testData.post.PollData,
			},
			expParentCommentsIDs: []types.PostID{},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			for _, p := range test.existingPosts {
				store.Set(types.PostStoreKey(p.PostID), suite.cdc.MustMarshalBinaryBare(p))
			}

			// Save the post
			suite.keeper.SavePost(suite.ctx, test.newPost)

			// Check the stored post
			var expected types.Post
			suite.keeper.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(test.newPost.PostID)), &expected)
			suite.True(expected.Equals(test.newPost))

			// Check the parent comments
			var parentCommentsIDs []types.PostID
			suite.keeper.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostCommentsStoreKey(test.newPost.ParentID)), &parentCommentsIDs)
			suite.True(test.expParentCommentsIDs.Equals(parentCommentsIDs))
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetPost() {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")

	tests := []struct {
		name       string
		postExists bool
		ID         types.PostID
		expected   types.Post
	}{
		{
			name:     "Non existent post is not found",
			ID:       id,
			expected: types.Post{},
		},
		{
			name:       "Existing post is found properly",
			ID:         id,
			postExists: true,
			expected: types.Post{
				PostID:       id,
				Message:      "Post",
				Created:      suite.testData.post.Created,
				LastEdited:   suite.testData.post.LastEdited,
				Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: map[string]string{},
				Creator:      suite.testData.postOwner,
			},
		},
		{
			name:       "Existing post with medias is found properly",
			ID:         id,
			postExists: true,
			expected: types.Post{
				PostID:       id,
				Message:      "Post",
				Created:      suite.testData.post.Created,
				LastEdited:   suite.testData.post.LastEdited,
				Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: map[string]string{},
				Creator:      suite.testData.postOwner,
				Attachments:  suite.testData.post.Attachments,
			},
		},
		{
			name:       "Existing post with poll is found properly",
			ID:         id,
			postExists: true,
			expected: types.Post{
				PostID:       id,
				Message:      "Post",
				Created:      suite.testData.post.Created,
				LastEdited:   suite.testData.post.LastEdited,
				Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: map[string]string{},
				Creator:      suite.testData.postOwner,
				PollData:     suite.testData.post.PollData,
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.StoreKey)

			if test.postExists {
				store.Set(types.PostStoreKey(test.expected.PostID), suite.keeper.Cdc.MustMarshalBinaryBare(&test.expected))
			}

			expected, found := suite.keeper.GetPost(suite.ctx, test.ID)
			suite.Equal(test.postExists, found)
			if test.postExists {
				suite.True(expected.Equals(test.expected))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetPostChildrenIDs() {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	id2 := types.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd")
	id3 := types.PostID("4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")
	id4 := types.PostID("a33e173b6b96129f74acf41b5219a6bbc9f90e9e41f37115f1ce7f1f5860211c")
	tests := []struct {
		name           string
		storedPosts    types.Posts
		postID         types.PostID
		expChildrenIDs types.PostIDs
	}{
		{
			name:           "Empty children list is returned properly",
			postID:         id,
			expChildrenIDs: types.PostIDs{},
		},
		{
			name: "Non empty children list is returned properly",
			storedPosts: types.Posts{
				types.Post{PostID: id, Message: "Original post", Created: suite.testData.post.Created,
					LastEdited: suite.testData.post.LastEdited, Subspace: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: map[string]string{}, Creator: suite.testData.post.Creator},
				types.Post{PostID: id2, ParentID: id, Message: "First commit", Created: suite.testData.post.Created,
					LastEdited: suite.testData.post.LastEdited, Subspace: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: map[string]string{}, Creator: suite.testData.post.Creator},
				types.Post{PostID: id3, Message: "Second post", Created: suite.testData.post.Created,
					LastEdited: suite.testData.post.LastEdited, Subspace: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: map[string]string{}, Creator: suite.testData.post.Creator},
				types.Post{PostID: id4, ParentID: id3, Message: "Comment to second post", Created: suite.testData.post.Created,
					LastEdited: suite.testData.post.LastEdited, Subspace: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: map[string]string{}, Creator: suite.testData.post.Creator},
			},
			postID:         id,
			expChildrenIDs: types.PostIDs{id2},
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
		posts types.Posts
	}{
		{
			name:  "Empty list returns empty list",
			posts: types.Posts{},
		},
		{
			name: "Existing list is returned properly",
			posts: types.Posts{
				types.Post{
					PostID:       "63b173547f1079e46885aa3ad4e36d0fe4beea8b7e2ec9c1d71ba3bff1abd909",
					Created:      suite.testData.post.Created,
					LastEdited:   suite.testData.post.LastEdited,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: map[string]string{},
					Creator:      suite.testData.postOwner,
				},
				types.Post{
					PostID:       "aad15654d10acd67b942ca39afd7a2aa071aed7c3f0b946edd2b666a037026f7",
					Created:      suite.testData.post.Created,
					LastEdited:   suite.testData.post.LastEdited,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: map[string]string{},
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
				suite.True(post.Equals(posts[index]))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetPostsFiltered() {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	id2 := types.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd")
	id3 := types.PostID("4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")
	id4 := types.PostID("a33e173b6b96129f74acf41b5219a6bbc9f90e9e41f37115f1ce7f1f5860211c")
	id5 := types.PostID("84a5d9fc5f0acd2bb9c0a49ecaefabbe4698372e1ae88d32f9f6f80b3c0ab95e")
	boolTrue := true

	creator1, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)

	creator2, err := sdk.AccAddressFromBech32("cosmos1jlhazemxvu0zn9y77j6afwmpf60zveqw5480l2")
	suite.NoError(err)

	timeZone, err := time.LoadLocation("UTC")
	suite.NoError(err)

	date := time.Date(2020, 1, 1, 1, 1, 0, 0, timeZone)

	posts := types.Posts{
		types.Post{
			PostID:       id2,
			ParentID:     id,
			Message:      "Post 1 #test #desmos",
			Created:      date,
			OptionalData: map[string]string{},
			Creator:      creator1,
		},
		types.Post{
			PostID:         id3,
			ParentID:       id,
			Message:        "Post 2",
			Created:        time.Date(2020, 2, 1, 1, 1, 0, 0, timeZone),
			AllowsComments: true,
			Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			OptionalData:   map[string]string{},
			Creator:        creator2,
		},
		types.Post{
			PostID:       id4,
			ParentID:     id5,
			Message:      "Post 3",
			Created:      time.Date(2020, 3, 1, 1, 1, 0, 0, timeZone),
			Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			OptionalData: map[string]string{},
			Creator:      creator2,
		},
	}

	tests := []struct {
		name     string
		filter   types.QueryPostsParams
		expected types.Posts
	}{
		{
			name:     "Valid pagination works properly",
			filter:   types.DefaultQueryPostsParams(1, 2),
			expected: types.Posts{posts[0], posts[1]},
		},
		{
			name:     "Non existing page returns empty list",
			filter:   types.DefaultQueryPostsParams(10, 1),
			expected: types.Posts{},
		},
		{
			name:     "Invalid pagination returns all data",
			filter:   types.DefaultQueryPostsParams(1, 15),
			expected: types.Posts{posts[0], posts[1], posts[2]},
		},
		{
			name:     "Parent ID matcher works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, ParentID: &posts[0].ParentID},
			expected: types.Posts{posts[1], posts[0]},
		},
		{
			name:     "Creation time matcher works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, CreationTime: &date},
			expected: types.Posts{posts[0]},
		},
		{
			name:     "Allows comments matcher works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, AllowsComments: &boolTrue},
			expected: types.Posts{posts[1]},
		},
		{
			name:     "Subspace mather works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, Subspace: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"},
			expected: types.Posts{posts[1], posts[2]},
		},
		{
			name:     "Creator mather works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, Creator: creator2},
			expected: types.Posts{posts[1], posts[2]},
		},
		{
			name:     "Sorting by date ascending works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, SortBy: types.PostSortByCreationDate, SortOrder: types.PostSortOrderAscending},
			expected: types.Posts{posts[0], posts[1], posts[2]},
		},
		{
			name:     "Sorting by date descending works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, SortBy: types.PostSortByCreationDate, SortOrder: types.PostSortOrderDescending},
			expected: types.Posts{posts[2], posts[1], posts[0]},
		},
		{
			name:     "Sorting by ID ascending works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, SortBy: types.PostSortByID, SortOrder: types.PostSortOrderAscending},
			expected: types.Posts{posts[1], posts[2], posts[0]},
		},
		{
			name:     "Sorting by ID descending works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, SortBy: types.PostSortByID, SortOrder: types.PostSortOrderDescending},
			expected: types.Posts{posts[0], posts[2], posts[1]},
		},
		{
			name:     "Filtering by hashtags works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, Hashtags: []string{"desmos", "test"}},
			expected: types.Posts{posts[0]},
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
				suite.True(test.expected[index].Equals(post))
			}
		})
	}
}
