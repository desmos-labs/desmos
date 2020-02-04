package keeper_test

import (
	"fmt"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/assert"
)

// -------------
// --- Posts
// -------------

func TestKeeper_GetLastPostId(t *testing.T) {
	tests := []struct {
		name       string
		existingID types.PostID
		expected   types.PostID
	}{
		{
			name:     "First ID returns correct value",
			expected: types.PostID(0),
		},
		{
			name:       "Existing ID returns correct value",
			existingID: types.PostID(3),
			expected:   types.PostID(3),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			if test.existingID.Valid() {
				store := ctx.KVStore(k.StoreKey)
				store.Set([]byte(types.LastPostIDStoreKey), k.Cdc.MustMarshalBinaryBare(test.existingID))
			}

			actual := k.GetLastPostID(ctx)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestKeeper_SavePost(t *testing.T) {
	tests := []struct {
		name                 string
		existingPosts        types.Posts
		lastPostID           types.PostID
		newPost              types.Post
		expParentCommentsIDs types.PostIDs
		expLastID            types.PostID
		expMedias            types.PostMedias
		expPollData          *types.PollData
	}{
		{
			name: "Post with ID already present",
			existingPosts: types.Posts{
				types.NewPost(types.PostID(1),
					types.PostID(0),
					"Post",
					false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					map[string]string{},
					testPost.Created,
					testPost.Creator,
					testPost.Medias,
					testPost.PollData,
				),
			},
			lastPostID: types.PostID(1),
			newPost: types.NewPost(types.PostID(1),
				types.PostID(0),
				"New post",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPost.Creator,
				testPost.Medias,
				testPost.PollData,
			),
			expParentCommentsIDs: []types.PostID{},
			expLastID:            types.PostID(1),
			expMedias:            testPost.Medias,
			expPollData:          testPost.PollData,
		},
		{
			name: "Post which ID is not already present",
			existingPosts: types.Posts{
				types.NewPost(types.PostID(1),
					types.PostID(0),
					"Post",
					false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					map[string]string{},
					testPost.Created,
					testPost.Creator,
					testPost.Medias,
					testPost.PollData,
				),
			},
			lastPostID: types.PostID(1),
			newPost: types.NewPost(types.PostID(15),
				types.PostID(0),
				"New post",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPost.Creator,
				testPost.Medias,
				testPost.PollData,
			),
			expParentCommentsIDs: []types.PostID{},
			expLastID:            types.PostID(15),
			expMedias:            testPost.Medias,
			expPollData:          testPost.PollData,
		},
		{
			name: "Post with valid parent ID",
			existingPosts: []types.Post{
				types.NewPost(types.PostID(1),
					types.PostID(0),
					"Parent",
					false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					map[string]string{},
					testPost.Created,
					testPost.Creator,
					testPost.Medias,
					testPost.PollData,
				),
			},
			lastPostID: types.PostID(1),
			newPost: types.NewPost(types.PostID(15),
				types.PostID(1),
				"Comment",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPost.Creator,
				testPost.Medias,
				testPost.PollData,
			),
			expParentCommentsIDs: []types.PostID{types.PostID(15)},
			expLastID:            types.PostID(15),
			expMedias:            testPost.Medias,
			expPollData:          testPost.PollData,
		},
		{
			name: "Post with ID greater ID than Last ID stored",
			existingPosts: types.Posts{
				types.NewPost(types.PostID(4),
					types.PostID(0),
					"Post lesser",
					false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					map[string]string{},
					testPost.Created,
					testPostOwner,
					testPost.Medias,
					testPost.PollData,
				),
			},
			lastPostID: types.PostID(4),
			newPost: types.NewPost(types.PostID(5),
				types.PostID(0),
				"New post greater",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{"key": "value"},
				testPost.Created,
				testPostOwner,
				testPost.Medias,
				testPost.PollData,
			),
			expParentCommentsIDs: []types.PostID{},
			expLastID:            types.PostID(5),
			expMedias:            testPost.Medias,
			expPollData:          testPost.PollData,
		},
		{
			name: "Post with ID lesser ID than Last ID stored",
			existingPosts: types.Posts{
				types.NewPost(types.PostID(4),
					types.PostID(0),
					"Post ID greater",
					false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					map[string]string{},
					testPost.Created,
					testPostOwner,
					testPost.Medias,
					testPost.PollData,
				),
			},
			lastPostID: types.PostID(4),
			newPost: types.NewPost(types.PostID(3),
				types.PostID(0),
				"New post ID lesser",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPostOwner,
				testPost.Medias,
				testPost.PollData,
			),
			expParentCommentsIDs: []types.PostID{},
			expLastID:            types.PostID(4),
			expMedias:            testPost.Medias,
			expPollData:          testPost.PollData,
		},
		{
			name:          "Post without medias is saved properly",
			existingPosts: types.Posts{},
			lastPostID:    types.PostID(0),
			newPost: types.NewPost(
				types.PostID(1),
				types.PostID(0),
				"Post without medias",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPostOwner,
				nil,
				testPost.PollData,
			),
			expParentCommentsIDs: []types.PostID{},
			expLastID:            types.PostID(1),
			expMedias:            nil,
			expPollData:          testPost.PollData,
		},
		{
			name:          "Post without poll data is saved properly",
			existingPosts: types.Posts{},
			lastPostID:    types.PostID(0),
			newPost: types.NewPost(types.PostID(1),
				types.PostID(0),
				"New post ID lesser",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPostOwner,
				testPost.Medias,
				nil,
			),
			expParentCommentsIDs: []types.PostID{},
			expLastID:            types.PostID(1),
			expMedias:            testPost.Medias,
			expPollData:          nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			store := ctx.KVStore(k.StoreKey)
			for _, p := range test.existingPosts {
				store.Set([]byte(types.PostStorePrefix+p.PostID.String()), k.Cdc.MustMarshalBinaryBare(p))
				store.Set([]byte(types.LastPostIDStoreKey), k.Cdc.MustMarshalBinaryBare(test.lastPostID))
			}

			// Save the post
			k.SavePost(ctx, test.newPost)

			// Check the stored post
			var expected types.Post
			k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostStorePrefix+test.newPost.PostID.String())), &expected)
			assert.True(t, expected.Equals(test.newPost))

			// Check the latest post id
			var lastPostID types.PostID
			k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LastPostIDStoreKey)), &lastPostID)
			assert.Equal(t, test.expLastID, lastPostID)

			// Check the parent comments
			var parentCommentsIDs []types.PostID
			k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostCommentsStorePrefix+test.newPost.ParentID.String())), &parentCommentsIDs)
			assert.True(t, test.expParentCommentsIDs.Equals(parentCommentsIDs))
		})
	}
}

func TestKeeper_GetPost(t *testing.T) {
	tests := []struct {
		name       string
		postExists bool
		ID         types.PostID
		expected   types.Post
	}{
		{
			name:     "Non existent post is not found",
			ID:       types.PostID(123),
			expected: types.Post{},
		},
		{
			name:       "Existing post is found properly",
			ID:         types.PostID(45),
			postExists: true,
			expected: types.NewPost(
				types.PostID(45),
				types.PostID(0),
				"Post",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPostOwner,
				testPost.Medias,
				testPost.PollData,
			),
		},
		{
			name:       "Existing post without medias is found properly",
			ID:         types.PostID(45),
			postExists: true,
			expected: types.NewPost(
				types.PostID(45),
				types.PostID(0),
				"Post",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPostOwner,
				nil,
				testPost.PollData,
			),
		},
		{
			name:       "Existing post without poll is found properly",
			ID:         types.PostID(45),
			postExists: true,
			expected: types.NewPost(
				types.PostID(45),
				types.PostID(0),
				"Post",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPostOwner,
				testPost.Medias,
				nil,
			),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)

			if test.postExists {
				store.Set([]byte(types.PostStorePrefix+test.expected.PostID.String()), k.Cdc.MustMarshalBinaryBare(&test.expected))
			}

			expected, found := k.GetPost(ctx, test.ID)
			assert.Equal(t, test.postExists, found)
			if test.postExists {
				assert.True(t, expected.Equals(test.expected))
			}
		})
	}
}

func TestKeeper_GetPostChildrenIDs(t *testing.T) {
	tests := []struct {
		name           string
		storedPosts    types.Posts
		postID         types.PostID
		expChildrenIDs types.PostIDs
	}{
		{
			name:           "Empty children list is returned properly",
			postID:         types.PostID(76),
			expChildrenIDs: types.PostIDs{},
		},
		{
			name: "Non empty children list is returned properly",
			storedPosts: types.Posts{
				types.NewPost(types.PostID(10), types.PostID(0), "Original post", false, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", map[string]string{}, testPost.Created, testPost.Creator, testPost.Medias, testPost.PollData),
				types.NewPost(types.PostID(55), types.PostID(10), "First commit", false, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", map[string]string{}, testPost.Created, testPost.Creator, testPost.Medias, testPost.PollData),
				types.NewPost(types.PostID(11), types.PostID(0), "Second post", false, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", map[string]string{}, testPost.Created, testPost.Creator, testPost.Medias, testPost.PollData),
				types.NewPost(types.PostID(104), types.PostID(11), "Comment to second post", false, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", map[string]string{}, testPost.Created, testPost.Creator, testPost.Medias, testPost.PollData),
			},
			postID:         types.PostID(10),
			expChildrenIDs: types.PostIDs{types.PostID(55)},
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
			assert.Len(t, storedChildrenIDs, len(test.expChildrenIDs))

			for _, id := range test.expChildrenIDs {
				assert.Contains(t, storedChildrenIDs, id)
			}
		})
	}
}

func TestKeeper_GetPosts(t *testing.T) {
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
					types.PostID(13),
					types.PostID(0),
					"",
					false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					map[string]string{},
					testPost.Created,
					testPostOwner,
					testPost.Medias,
					testPost.PollData,
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
				store.Set([]byte(types.PostStorePrefix+p.PostID.String()), k.Cdc.MustMarshalBinaryBare(p))
			}

			posts := k.GetPosts(ctx)
			for index, post := range test.posts {
				assert.True(t, post.Equals(posts[index]))
			}
		})
	}
}

func TestKeeper_GetPostsFiltered(t *testing.T) {
	boolTrue := true

	creator1, _ := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	creator2, _ := sdk.AccAddressFromBech32("cosmos1jlhazemxvu0zn9y77j6afwmpf60zveqw5480l2")

	timeZone, _ := time.LoadLocation("UTC")
	date := time.Date(2020, 1, 1, 1, 1, 0, 0, timeZone)

	posts := types.Posts{
		types.NewPost(
			types.PostID(10),
			types.PostID(1),
			"Post 1",
			false,
			"",
			map[string]string{},
			date,
			creator1,
			testPost.Medias,
			testPost.PollData,
		),
		types.NewPost(
			types.PostID(11),
			types.PostID(1),
			"Post 2",
			true,
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			map[string]string{},
			time.Date(2020, 2, 1, 1, 1, 0, 0, timeZone),
			creator2,
			testPost.Medias,
			testPost.PollData,
		),
		types.NewPost(
			types.PostID(12),
			types.PostID(2),
			"Post 3",
			false,
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			map[string]string{},
			date,
			creator2,
			testPost.Medias,
			testPost.PollData,
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
			expected: types.Posts{posts[0], posts[1]},
		},
		{
			name:     "Creation time matcher works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, CreationTime: &date},
			expected: types.Posts{posts[0], posts[2]},
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
			expected: types.Posts{posts[0], posts[2], posts[1]},
		},
		{
			name:     "Sorting by date descending works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, SortBy: types.PostSortByCreationDate, SortOrder: types.PostSortOrderDescending},
			expected: types.Posts{posts[1], posts[0], posts[2]},
		},
		{
			name:     "Sorting by ID ascending works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, SortBy: types.PostSortByID, SortOrder: types.PostSortOrderAscending},
			expected: types.Posts{posts[0], posts[1], posts[2]},
		},
		{
			name:     "Sorting by ID descending works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, SortBy: types.PostSortByID, SortOrder: types.PostSortOrderDescending},
			expected: types.Posts{posts[2], posts[1], posts[0]},
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
			assert.True(t, test.expected.Equals(result), "Expected\n%s\nbut got\n%s", test.expected, result)
		})
	}
}

func TestKeeper_SavePollPostAnswers(t *testing.T) {
	user, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	user2, _ := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	answers := []uint{uint(1), uint(2)}
	answers2 := []uint{uint(1)}

	tests := []struct {
		name               string
		postID             types.PostID
		userAnswersDetails types.AnswersDetails
		previousUsersAD    types.UsersAnswersDetails
		expUsersAD         types.UsersAnswersDetails
	}{
		{
			name:               "Save answers with no previous answers in this context",
			postID:             types.PostID(1),
			userAnswersDetails: types.NewAnswersDetails(answers, user),
			previousUsersAD:    nil,
			expUsersAD:         types.UsersAnswersDetails{types.NewAnswersDetails(answers, user)},
		},
		{
			name:               "Save new answers",
			postID:             types.PostID(1),
			userAnswersDetails: types.NewAnswersDetails(answers2, user2),
			previousUsersAD:    types.UsersAnswersDetails{types.NewAnswersDetails(answers, user)},
			expUsersAD: types.UsersAnswersDetails{
				types.NewAnswersDetails(answers, user),
				types.NewAnswersDetails(answers2, user2),
			},
		},
	}

	for _, test := range tests {

		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)

			if test.previousUsersAD != nil {
				store.Set([]byte(types.PollAnswersStorePrefix+test.postID.String()),
					k.Cdc.MustMarshalBinaryBare(test.previousUsersAD))
			}

			k.SavePollUserAnswers(ctx, test.postID, test.userAnswersDetails)

			var actualUsersAnswersDetails types.UsersAnswersDetails
			answersBz := store.Get([]byte(types.PollAnswersStorePrefix + test.postID.String()))
			k.Cdc.MustUnmarshalBinaryBare(answersBz, &actualUsersAnswersDetails)
			assert.Equal(t, test.expUsersAD, actualUsersAnswersDetails)
		})
	}
}

func TestKeeper_GetPostPollAnswersDetails(t *testing.T) {
	user, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	answers := []uint{uint(1), uint(2)}

	tests := []struct {
		name          string
		postID        types.PostID
		storedAnswers types.UsersAnswersDetails
	}{
		{
			name:          "No answers returns empty list",
			postID:        types.PostID(1),
			storedAnswers: nil,
		},
		{
			name:          "Answers returned correctly",
			postID:        types.PostID(1),
			storedAnswers: types.UsersAnswersDetails{types.NewAnswersDetails(answers, user)},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			if test.storedAnswers != nil {
				k.SavePollUserAnswers(ctx, test.postID, test.storedAnswers[0])
			}

			actualPostPollAnswers := k.GetPostPollAnswersDetails(ctx, test.postID)

			assert.Equal(t, test.storedAnswers, actualPostPollAnswers)
		})
	}
}

func TestKeeper_GetPostPollAnswersByUser(t *testing.T) {
	user, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	user2, _ := sdk.AccAddressFromBech32("cosmos1jlhazemxvu0zn9y77j6afwmpf60zveqw5480l2")
	answers := []uint{uint(1), uint(2)}

	tests := []struct {
		name          string
		storedAnswers types.AnswersDetails
		postID        types.PostID
		user          sdk.AccAddress
		expAnswers    []uint
	}{
		{
			name:          "No answers for user returns nil",
			storedAnswers: types.NewAnswersDetails(answers, user),
			postID:        types.PostID(1),
			user:          user2,
			expAnswers:    nil,
		},
		{
			name:          "Matching user returns answers made by him",
			storedAnswers: types.NewAnswersDetails(answers, user),
			postID:        types.PostID(1),
			user:          user,
			expAnswers:    answers,
		},
	}

	for _, test := range tests {
		ctx, k := SetupTestInput()

		k.SavePollUserAnswers(ctx, test.postID, test.storedAnswers)

		actualPostPollAnswers := k.GetPostPollAnswersByUser(ctx, test.postID, test.user)

		assert.Equal(t, test.expAnswers, actualPostPollAnswers)
	}
}

func TestKeeper_GetAnswersDetailsMap(t *testing.T) {
	user, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	user2, _ := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	answers := []uint{uint(1), uint(2)}

	tests := []struct {
		name    string
		usersAD map[types.PostID]types.UsersAnswersDetails
	}{
		{
			name:    "Empty users answers details data are returned correctly",
			usersAD: map[types.PostID]types.UsersAnswersDetails{},
		},
		{
			name: "Non empty users answers details data are returned correcly",
			usersAD: map[types.PostID]types.UsersAnswersDetails{
				types.PostID(1): {
					types.NewAnswersDetails(answers, user),
					types.NewAnswersDetails(answers, user2),
				},
				types.PostID(2): {
					types.NewAnswersDetails(answers, user2),
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)
			for postID, userAD := range test.usersAD {
				store.Set([]byte(types.PollAnswersStorePrefix+postID.String()), k.Cdc.MustMarshalBinaryBare(userAD))
			}

			usersADData := k.GetAnswersDetailsMap(ctx)
			assert.Equal(t, test.usersAD, usersADData)
		})
	}
}

func TestKeeper_ClosePollPost(t *testing.T) {
	postID := types.PostID(3257)

	ctx, k := SetupTestInput()

	k.SavePost(ctx, testPost)

	k.ClosePollPost(ctx, postID)

	expPost, _ := k.GetPost(ctx, postID)

	assert.Equal(t, false, expPost.PollData.Open)
}

// -------------
// --- Reactions
// -------------

func TestKeeper_SaveReaction(t *testing.T) {
	liker, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	otherLiker, _ := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")

	tests := []struct {
		name           string
		storedLikes    types.Reactions
		postID         types.PostID
		like           types.Reaction
		error          error
		expectedStored types.Reactions
	}{
		{
			name:           "Reaction from same user already present returns expError",
			storedLikes:    types.Reactions{types.NewReaction("like", liker)},
			postID:         types.PostID(10),
			like:           types.NewReaction("like", liker),
			error:          fmt.Errorf("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4 has already reacted with like to the post with id 10"),
			expectedStored: types.Reactions{types.NewReaction("like", liker)},
		},
		{
			name:           "First liker is stored properly",
			storedLikes:    types.Reactions{},
			postID:         types.PostID(15),
			like:           types.NewReaction("like", liker),
			error:          nil,
			expectedStored: types.Reactions{types.NewReaction("like", liker)},
		},
		{
			name:        "Second liker is stored properly",
			storedLikes: types.Reactions{types.NewReaction("like", liker)},
			postID:      types.PostID(87),
			like:        types.NewReaction("like", otherLiker),
			error:       nil,
			expectedStored: types.Reactions{
				types.NewReaction("like", liker),
				types.NewReaction("like", otherLiker),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			store := ctx.KVStore(k.StoreKey)
			if len(test.storedLikes) != 0 {
				store.Set([]byte(types.PostReactionsStorePrefix+test.postID.String()), k.Cdc.MustMarshalBinaryBare(&test.storedLikes))
			}

			err := k.SaveReaction(ctx, test.postID, test.like)
			assert.Equal(t, test.error, err)

			var stored types.Reactions
			k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostReactionsStorePrefix+test.postID.String())), &stored)
			assert.Equal(t, test.expectedStored, stored)
		})
	}
}

func TestKeeper_RemoveReaction(t *testing.T) {
	liker, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")

	tests := []struct {
		name           string
		storedLikes    types.Reactions
		postID         types.PostID
		liker          sdk.AccAddress
		value          string
		error          error
		expectedStored types.Reactions
	}{
		{
			name:           "Reaction from same liker is removed properly",
			storedLikes:    types.Reactions{types.NewReaction("like", liker)},
			postID:         types.PostID(10),
			liker:          liker,
			value:          "like",
			error:          nil,
			expectedStored: types.Reactions{},
		},
		{
			name:           "Non existing reaction returns error - Creator",
			storedLikes:    types.Reactions{},
			postID:         types.PostID(15),
			liker:          liker,
			value:          "like",
			error:          fmt.Errorf("cannot remove the reaction with value like from user cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4 as it does not exist"),
			expectedStored: types.Reactions{},
		},
		{
			name:           "Non existing reaction returns error - Value",
			storedLikes:    types.Reactions{types.NewReaction("like", liker)},
			postID:         types.PostID(15),
			liker:          liker,
			value:          "reaction",
			error:          fmt.Errorf("cannot remove the reaction with value reaction from user cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4 as it does not exist"),
			expectedStored: types.Reactions{types.NewReaction("like", liker)},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			store := ctx.KVStore(k.StoreKey)
			if len(test.storedLikes) != 0 {
				store.Set([]byte(types.PostReactionsStorePrefix+test.postID.String()), k.Cdc.MustMarshalBinaryBare(&test.storedLikes))
			}

			err := k.RemoveReaction(ctx, test.postID, test.liker, test.value)
			assert.Equal(t, test.error, err)

			var stored types.Reactions
			k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostReactionsStorePrefix+test.postID.String())), &stored)

			assert.Len(t, stored, len(test.expectedStored))
			for index, like := range test.expectedStored {
				assert.Equal(t, like, stored[index])
			}
		})
	}
}

func TestKeeper_GetPostLikes(t *testing.T) {
	liker, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	otherLiker, _ := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")

	tests := []struct {
		name   string
		likes  types.Reactions
		postID types.PostID
	}{
		{
			name:   "Empty list are returned properly",
			likes:  types.Reactions{},
			postID: types.PostID(10),
		},
		{
			name: "Valid list of likes is returned properly",
			likes: types.Reactions{
				types.NewReaction("like", otherLiker),
				types.NewReaction("like", liker),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			for _, l := range test.likes {
				_ = k.SaveReaction(ctx, test.postID, l)
			}

			stored := k.GetPostReactions(ctx, test.postID)

			assert.Len(t, stored, len(test.likes))
			for _, l := range test.likes {
				assert.Contains(t, stored, l)
			}
		})
	}
}

func TestKeeper_GetLikes(t *testing.T) {
	liker1, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	liker2, _ := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")

	tests := []struct {
		name  string
		likes map[types.PostID]types.Reactions
	}{
		{
			name:  "Empty likes data are returned correctly",
			likes: map[types.PostID]types.Reactions{},
		},
		{
			name: "Non empty likes data are returned correcly",
			likes: map[types.PostID]types.Reactions{
				types.PostID(5): {
					types.NewReaction("like", liker1),
					types.NewReaction("like", liker2),
				},
				types.PostID(10): {
					types.NewReaction("like", liker1),
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)
			for postID, likes := range test.likes {
				store.Set([]byte(types.PostReactionsStorePrefix+postID.String()), k.Cdc.MustMarshalBinaryBare(likes))
			}

			likesData := k.GetReactions(ctx)
			assert.Equal(t, test.likes, likesData)
		})
	}
}
