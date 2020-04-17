package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/require"
)

// -------------
// --- Posts
// -------------

func TestKeeper_SavePost(t *testing.T) {
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
				types.NewPost(id,
					"",
					"Post",
					false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					map[string]string{},
					testPost.Created,
					testPost.Creator,
				),
			},
			newPost: types.NewPost(id,
				"",
				"New post",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPost.Creator,
			),
			expParentCommentsIDs: []types.PostID{},
		},
		{
			name: "Post which ID is not already present",
			existingPosts: types.Posts{
				types.NewPost(id,
					"",
					"Post",
					false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					map[string]string{},
					testPost.Created,
					testPost.Creator,
				),
			},
			newPost: types.NewPost(id,
				"",
				"New post",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPost.Creator,
			),
			expParentCommentsIDs: []types.PostID{},
		},
		{
			name: "Post with valid parent ID",
			existingPosts: []types.Post{
				types.NewPost(id,
					"",
					"Parent",
					false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					map[string]string{},
					testPost.Created,
					testPost.Creator,
				),
			},
			newPost: types.NewPost(id2,
				id,
				"Comment",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPost.Creator,
			),
			expParentCommentsIDs: []types.PostID{id2},
		},
		{
			name: "Post with ID greater ID than Last ID stored",
			existingPosts: types.Posts{
				types.NewPost(id,
					"",
					"Post lesser",
					false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					map[string]string{},
					testPost.Created,
					testPostOwner,
				),
			},
			newPost: types.NewPost(id,
				"",
				"New post greater",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{"key": "value"},
				testPost.Created,
				testPostOwner,
			),
			expParentCommentsIDs: []types.PostID{},
		},
		{
			name: "Post with ID lesser ID than Last ID stored",
			existingPosts: types.Posts{
				types.NewPost(id,
					"",
					"Post ID greater",
					false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					map[string]string{},
					testPost.Created,
					testPostOwner,
				),
			},
			newPost: types.NewPost(id,
				"",
				"New post ID lesser",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPostOwner,
			),
			expParentCommentsIDs: []types.PostID{},
		},
		{
			name:          "Post without medias is saved properly",
			existingPosts: types.Posts{},
			newPost: types.NewPost(
				id,
				"",
				"Post without medias",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPostOwner,
			).WithMedias(testPost.Medias),
			expParentCommentsIDs: []types.PostID{},
		},
		{
			name:          "Post without poll data is saved properly",
			existingPosts: types.Posts{},
			newPost: types.NewPost(id,
				"",
				"New post ID lesser",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPostOwner,
			).WithPollData(*testPost.PollData),
			expParentCommentsIDs: []types.PostID{},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			store := ctx.KVStore(k.StoreKey)
			for _, p := range test.existingPosts {
				store.Set(types.PostStoreKey(p.PostID), k.Cdc.MustMarshalBinaryBare(p))
			}

			// Save the post
			k.SavePost(ctx, test.newPost)

			// Check the stored post
			var expected types.Post
			k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(test.newPost.PostID)), &expected)
			require.True(t, expected.Equals(test.newPost))

			// Check the parent comments
			var parentCommentsIDs []types.PostID
			k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostCommentsStoreKey(test.newPost.ParentID)), &parentCommentsIDs)
			require.True(t, test.expParentCommentsIDs.Equals(parentCommentsIDs))
		})
	}
}

func TestKeeper_GetPost(t *testing.T) {
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
			expected: types.NewPost(
				id,
				"",
				"Post",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPostOwner,
			),
		},
		{
			name:       "Existing post with medias is found properly",
			ID:         id,
			postExists: true,
			expected: types.NewPost(
				id,
				"",
				"Post",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPostOwner,
			).WithMedias(testPost.Medias),
		},
		{
			name:       "Existing post with poll is found properly",
			ID:         id,
			postExists: true,
			expected: types.NewPost(
				id,
				"",
				"Post",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPostOwner,
			).WithPollData(*testPost.PollData),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)

			if test.postExists {
				store.Set(types.PostStoreKey(test.expected.PostID), k.Cdc.MustMarshalBinaryBare(&test.expected))
			}

			expected, found := k.GetPost(ctx, test.ID)
			require.Equal(t, test.postExists, found)
			if test.postExists {
				require.True(t, expected.Equals(test.expected))
			}
		})
	}
}

func TestKeeper_GetPostChildrenIDs(t *testing.T) {
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
				types.NewPost(id, "", "Original post", false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", map[string]string{},
					testPost.Created, testPost.Creator),
				types.NewPost(id2, id, "First commit", false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", map[string]string{},
					testPost.Created, testPost.Creator),
				types.NewPost(id3, "", "Second post", false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", map[string]string{},
					testPost.Created, testPost.Creator),
				types.NewPost(id4, id3, "Comment to second post", false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", map[string]string{},
					testPost.Created, testPost.Creator),
			},
			postID:         id,
			expChildrenIDs: types.PostIDs{id2},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			for _, p := range test.storedPosts {
				k.SavePost(ctx, p)
			}

			storedChildrenIDs := k.GetPostChildrenIDs(ctx, test.postID)
			require.Len(t, storedChildrenIDs, len(test.expChildrenIDs))

			for _, id := range test.expChildrenIDs {
				require.Contains(t, storedChildrenIDs, id)
			}
		})
	}
}

func TestKeeper_GetPosts(t *testing.T) {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
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
				types.NewPost(
					id,
					"",
					"",
					false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					map[string]string{},
					testPost.Created,
					testPostOwner,
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			store := ctx.KVStore(k.StoreKey)
			for _, p := range test.posts {
				store.Set(types.PostStoreKey(p.PostID), k.Cdc.MustMarshalBinaryBare(p))
			}

			posts := k.GetPosts(ctx)
			for index, post := range test.posts {
				require.True(t, post.Equals(posts[index]))
			}
		})
	}
}

func TestKeeper_GetPostsFiltered(t *testing.T) {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	id2 := types.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd")
	id3 := types.PostID("4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")
	id4 := types.PostID("a33e173b6b96129f74acf41b5219a6bbc9f90e9e41f37115f1ce7f1f5860211c")
	id5 := types.PostID("84a5d9fc5f0acd2bb9c0a49ecaefabbe4698372e1ae88d32f9f6f80b3c0ab95e")
	boolTrue := true

	creator1, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	creator2, err := sdk.AccAddressFromBech32("cosmos1jlhazemxvu0zn9y77j6afwmpf60zveqw5480l2")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2020, 1, 1, 1, 1, 0, 0, timeZone)

	posts := types.Posts{
		types.NewPost(
			id2,
			id,
			"Post 1 #test #desmos",
			false,
			"",
			map[string]string{},
			date,
			creator1,
		),
		types.NewPost(
			id3,
			id,
			"Post 2",
			true,
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			map[string]string{},
			time.Date(2020, 2, 1, 1, 1, 0, 0, timeZone),
			creator2,
		),
		types.NewPost(
			id4,
			id5,
			"Post 3",
			false,
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			map[string]string{},
			time.Date(2020, 3, 1, 1, 1, 0, 0, timeZone),
			creator2,
		),
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
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			for _, post := range posts {
				k.SavePost(ctx, post)
			}
			result := k.GetPostsFiltered(ctx, test.filter)

			require.Len(t, result, len(test.expected))
			for index, post := range result {
				require.True(t, test.expected[index].Equals(post))
			}
		})
	}
}
