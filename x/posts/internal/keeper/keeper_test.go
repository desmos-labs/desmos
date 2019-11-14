package keeper_test

import (
	"testing"

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
		existingId types.PostID
		expected   types.PostID
	}{
		{
			name:     "First ID returns correct value",
			expected: types.PostID(0),
		},
		{
			name:       "Existing ID returns correct value",
			existingId: types.PostID(3),
			expected:   types.PostID(3),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			if test.existingId.Valid() {
				store := ctx.KVStore(k.StoreKey)
				store.Set([]byte(types.LastPostIDStoreKey), k.Cdc.MustMarshalBinaryBare(test.existingId))
			}

			actual := k.GetLastPostID(ctx)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestKeeper_SavePost(t *testing.T) {
	tests := []struct {
		name         string
		existingPost types.Post
		newPost      types.Post
		error        sdk.Error
	}{
		{
			name:         "Duplicate ID is overridden",
			existingPost: types.NewPost(types.PostID(0), types.PostID(0), "Post", 0, testPostOwner),
			newPost:      types.NewPost(types.PostID(0), types.PostID(10), "New post", 0, testPostOwner),
			error:        nil,
		},
		{
			name:         "Not duplicate ID saved correctly",
			existingPost: types.NewPost(types.PostID(0), types.PostID(0), "Post", 0, testPostOwner),
			newPost:      types.NewPost(types.PostID(15), types.PostID(10), "New post", 0, testPostOwner),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			store := ctx.KVStore(k.StoreKey)
			store.Set(
				[]byte(types.PostStorePrefix+test.existingPost.PostID.String()),
				k.Cdc.MustMarshalBinaryBare(&test.existingPost),
			)

			err := k.SavePost(ctx, test.newPost)
			assert.Equal(t, test.error, err)

			if test.error == nil {
				var expected types.Post
				k.Cdc.MustUnmarshalBinaryBare(
					store.Get([]byte(types.PostStorePrefix+test.newPost.PostID.String())),
					&expected,
				)
				assert.Equal(t, test.newPost, expected)

				var lastPostID types.PostID
				k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LastPostIDStoreKey)), &lastPostID)
				assert.Equal(t, test.newPost.PostID, lastPostID)
			}
		})
	}
}

func TestKeeper_GetPost(t *testing.T) {
	tests := []struct {
		name          string
		existingPost  types.Post
		ID            types.PostID
		expectedFound bool
		expected      types.Post
	}{
		{
			name:          "Non existent post is not found",
			ID:            types.PostID(123),
			expectedFound: false,
			expected:      types.Post{},
		},
		{
			name:          "Existing post is found properly",
			existingPost:  types.NewPost(types.PostID(45), types.PostID(0), "Post", 0, testPostOwner),
			ID:            types.PostID(45),
			expectedFound: true,
			expected:      types.NewPost(types.PostID(45), types.PostID(0), "Post", 0, testPostOwner),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)

			if !(types.Post{}).Equals(test.existingPost) {
				store.Set(
					[]byte(types.PostStorePrefix+test.existingPost.PostID.String()),
					k.Cdc.MustMarshalBinaryBare(&test.existingPost),
				)
			}

			expected, found := k.GetPost(ctx, test.ID)
			assert.Equal(t, test.expected, expected)
			assert.Equal(t, test.expectedFound, found)
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
				types.Post{PostID: types.PostID(13)},
				types.Post{PostID: types.PostID(76)},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			store := ctx.KVStore(k.StoreKey)
			for _, p := range test.posts {
				store.Set([]byte(types.PostStorePrefix+p.PostID.String()), k.Cdc.MustMarshalBinaryBare(&p))
			}

			posts := k.GetPosts(ctx)
			assert.True(t, test.posts.Equals(posts))
		})
	}
}

// -------------
// --- Likes
// -------------

func TestKeeper_SaveLike(t *testing.T) {
	liker, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	otherLiker, _ := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")

	tests := []struct {
		name           string
		storedLikes    types.Likes
		postID         types.PostID
		like           types.Like
		error          sdk.Error
		expectedStored types.Likes
	}{
		{
			name:           "Like from same liker already present returns error",
			storedLikes:    types.Likes{types.NewLike(10, liker)},
			postID:         types.PostID(10),
			like:           types.NewLike(50, liker),
			error:          sdk.ErrUnknownRequest("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4 has already liked the post with id 10"),
			expectedStored: types.Likes{types.NewLike(10, liker)},
		},
		{
			name:           "First like is stored properly",
			storedLikes:    types.Likes{},
			postID:         types.PostID(15),
			like:           types.NewLike(15, liker),
			error:          nil,
			expectedStored: types.Likes{types.NewLike(15, liker)},
		},
		{
			name:        "Second like is stored properly",
			storedLikes: types.Likes{types.NewLike(10, liker)},
			postID:      types.PostID(87),
			like:        types.NewLike(1, otherLiker),
			error:       nil,
			expectedStored: types.Likes{
				types.NewLike(10, liker),
				types.NewLike(1, otherLiker),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			store := ctx.KVStore(k.StoreKey)
			if len(test.storedLikes) != 0 {
				store.Set([]byte(types.LikesStorePrefix+test.postID.String()), k.Cdc.MustMarshalBinaryBare(&test.storedLikes))
			}

			err := k.SaveLike(ctx, test.postID, test.like)
			assert.Equal(t, test.error, err)

			var stored types.Likes
			k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LikesStorePrefix+test.postID.String())), &stored)
			assert.Equal(t, test.expectedStored, stored)
		})
	}
}

func TestKeeper_GetLikes(t *testing.T) {
	liker1, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	liker2, _ := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")

	tests := []struct {
		name  string
		likes map[types.PostID]types.Likes
	}{
		{
			name:  "Empty likes data are returned correctly",
			likes: map[types.PostID]types.Likes{},
		},
		{
			name: "Non empty likes data are returned correcly",
			likes: map[types.PostID]types.Likes{
				types.PostID(5): {
					types.NewLike(10, liker1),
					types.NewLike(50, liker2),
				},
				types.PostID(10): {
					types.NewLike(5, liker1),
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
				store.Set([]byte(types.LikesStorePrefix+postID.String()), k.Cdc.MustMarshalBinaryBare(&likes))
			}

			likesData := k.GetLikes(ctx)
			assert.Equal(t, test.likes, likesData)
		})
	}
}
