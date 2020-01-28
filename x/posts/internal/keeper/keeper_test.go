package keeper_test

import (
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
					"desmos",
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
				"desmos",
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
					"desmos",
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
				"desmos",
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
					"desmos",
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
				"desmos",
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
					"desmos",
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
				"desmos",
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
					"desmos",
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
				"desmos",
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
				"desmos",
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
				"desmos",
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
				"desmos",
				map[string]string{},
				testPost.Created,
				testPostOwner,
				testPost.Medias,
				testPost.PollData,
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
				types.NewPost(types.PostID(10), types.PostID(0), "Original post", false, "desmos", map[string]string{}, testPost.Created, testPost.Creator, testPost.Medias, testPost.PollData),
				types.NewPost(types.PostID(55), types.PostID(10), "First commit", false, "desmos", map[string]string{}, testPost.Created, testPost.Creator, testPost.Medias, testPost.PollData),
				types.NewPost(types.PostID(11), types.PostID(0), "Second post", false, "desmos", map[string]string{}, testPost.Created, testPost.Creator, testPost.Medias, testPost.PollData),
				types.NewPost(types.PostID(104), types.PostID(11), "Comment to second post", false, "desmos", map[string]string{}, testPost.Created, testPost.Creator, testPost.Medias, testPost.PollData),
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
					"desmos",
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
			"desmos",
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
			"desmos",
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
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, Subspace: "desmos"},
			expected: types.Posts{posts[1], posts[2]},
		},
		{
			name:     "Creator mather works properly",
			filter:   types.QueryPostsParams{Page: 1, Limit: 5, Creator: creator2},
			expected: types.Posts{posts[1], posts[2]},
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
			assert.Len(t, result, len(test.expected))
			for index, post := range result {
				assert.True(t, test.expected[index].Equals(post))
			}
		})
	}
}

func TestKeeper_SavePollPostAnswers(t *testing.T) {
	creator1, _ := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")

	tests := []struct {
		name            string
		postID          types.PostID
		answerer        sdk.AccAddress
		answers         []uint64
		previousAnswers []uint64
		expAnswers      []uint64
	}{
		{
			name:            "Save answers with no previous answers in this context",
			postID:          types.PostID(1),
			answers:         []uint64{1, 2},
			previousAnswers: nil,
			expAnswers:      []uint64{1, 2},
			answerer:        creator1,
		},
		{
			name:            "Save answers and overridden the previous ones",
			postID:          types.PostID(1),
			answers:         []uint64{1},
			previousAnswers: []uint64{2},
			expAnswers:      []uint64{1},
			answerer:        creator1,
		},
	}

	for _, test := range tests {

		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)

			if test.previousAnswers != nil {
				store.Set([]byte(types.PollAnswersStorePrefix+test.postID.String()+test.answerer.String()),
					k.Cdc.MustMarshalBinaryBare(test.previousAnswers))
			}

			k.SavePollPostAnswers(ctx, test.postID, test.answers, test.answerer)

			var actualAnswers []uint64
			answersBz := store.Get([]byte(types.PollAnswersStorePrefix + test.postID.String() + test.answerer.String()))
			k.Cdc.MustUnmarshalBinaryBare(answersBz, &actualAnswers)
			assert.Equal(t, test.expAnswers, actualAnswers)
		})
	}
}

func TestKeeper_GetPollPostUserAnswers(t *testing.T) {
	creator1, _ := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")

	tests := []struct {
		name       string
		postID     types.PostID
		user       sdk.AccAddress
		answers    []uint64
		expAnswers []uint64
	}{
		{
			name:       "User hadn't post any answer",
			postID:     types.PostID(1),
			user:       creator1,
			answers:    nil,
			expAnswers: []uint64{},
		},
		{
			name:       "User had post answers",
			postID:     types.PostID(1),
			user:       creator1,
			answers:    []uint64{1, 2},
			expAnswers: []uint64{1, 2},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			if test.answers != nil {
				k.SavePollPostAnswers(ctx, test.postID, test.answers, test.user)
			}

			actualAnswers := k.GetPollPostUserAnswers(ctx, test.postID, test.user)

			assert.Equal(t, test.expAnswers, actualAnswers)
		})
	}
}

func TestKeeper_GetPollTotalAnswersAmount(t *testing.T) {
	creator, _ := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")

	tests := []struct {
		name     string
		postID   types.PostID
		answers  []uint64
		expTotal sdk.Int
	}{
		{
			name:     "Get the total number of poll answers",
			postID:   types.PostID(1),
			answers:  []uint64{1, 2, 3, 4, 5, 6, 7, 8},
			expTotal: sdk.NewInt(8),
		},
		{
			name:     "Zero answers to poll",
			postID:   types.PostID(1),
			answers:  nil,
			expTotal: sdk.NewInt(0),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			if test.answers != nil {
				k.SavePollPostAnswers(ctx, test.postID, test.answers, creator)
			}

			total := k.GetPollTotalAnswersAmount(ctx, test.postID)

			assert.Equal(t, total, test.expTotal)
		})
	}
}

func TestKeeper_GetAnswerTotalVotes(t *testing.T) {
	creator, _ := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	creator2, _ := sdk.AccAddressFromBech32("cosmos1jlhazemxvu0zn9y77j6afwmpf60zveqw5480l2")

	tests := []struct {
		name     string
		postID   types.PostID
		answerID uint64
		answers  [][]uint64
		users    []sdk.AccAddress
		expTotal sdk.Int
	}{
		{
			name:     "Get the total votes for an answers",
			postID:   types.PostID(1),
			answerID: uint64(1),
			answers:  [][]uint64{{1, 3}, {1, 5}},
			users:    []sdk.AccAddress{creator, creator2},
			expTotal: sdk.NewInt(2),
		},
		{
			name:     "Answer with 0 votes",
			postID:   types.PostID(1),
			answerID: uint64(1),
			answers:  [][]uint64{{2, 3}, {2, 3}},
			users:    []sdk.AccAddress{creator, creator2},
			expTotal: sdk.NewInt(0),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			for index, user := range test.users {
				k.SavePollPostAnswers(ctx, test.postID, test.answers[index], user)
			}

			totalVotes := k.GetAnswerTotalVotes(ctx, test.postID, test.answerID)

			assert.Equal(t, test.expTotal, totalVotes)
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
		error          sdk.Error
		expectedStored types.Reactions
	}{
		{
			name:           "Reaction from same user already present returns expError",
			storedLikes:    types.Reactions{types.NewReaction("like", liker)},
			postID:         types.PostID(10),
			like:           types.NewReaction("like", liker),
			error:          sdk.ErrUnknownRequest("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4 has already reacted with like to the post with id 10"),
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
		error          sdk.Error
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
			error:          sdk.ErrUnauthorized("Cannot remove the reaction with value like from user cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4 as it does not exist"),
			expectedStored: types.Reactions{},
		},
		{
			name:           "Non existing reaction returns error - Value",
			storedLikes:    types.Reactions{types.NewReaction("like", liker)},
			postID:         types.PostID(15),
			liker:          liker,
			value:          "reaction",
			error:          sdk.ErrUnauthorized("Cannot remove the reaction with value reaction from user cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4 as it does not exist"),
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
