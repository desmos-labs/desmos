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
			name:         "Duplicate ID returns error",
			existingPost: types.NewPost(types.PostID(0), types.PostID(0), "Post", 0, testPostOwner),
			newPost:      types.NewPost(types.PostID(0), types.PostID(10), "New post", 0, testPostOwner),
			error:        sdk.ErrUnknownRequest("Post with id 0 already existing"),
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

//
//// -------------
//// --- Likes
//// -------------
//
//func defaultLikeID() types.LikeID {
//	return types.LikeID(1)
//}
//
//func TestKeeper_GetLastLikeId_FirstId(t *testing.T) {
//	ctx, k := SetupTestInput()
//	assert.Equal(t, types.LikeID(0), k.GetLastLikeID(ctx))
//}
//
//func TestKeeper_GetLastLikeId_Existing(t *testing.T) {
//	ctx, k := SetupTestInput()
//
//	ids := []types.LikeID{types.LikeID(0), types.LikeID(3), types.LikeID(18446744073709551615)}
//
//	store := ctx.KVStore(k.StoreKey)
//	for _, id := range ids {
//		store.Set([]byte(types.LastLikeIDStoreKey), k.Cdc.MustMarshalBinaryBare(id))
//		assert.Equal(t, id, k.GetLastLikeID(ctx))
//	}
//}
//
//func TestKeeper_SetLastLikeId(t *testing.T) {
//	ctx, k := SetupTestInput()
//
//	ids := []types.LikeID{types.LikeID(0), types.LikeID(3), types.LikeID(18446744073709551615)}
//
//	store := ctx.KVStore(k.StoreKey)
//	for _, id := range ids {
//		k.SetLastLikeID(ctx, id)
//		var stored types.LikeID
//		k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LastLikeIDStoreKey)), &stored)
//		assert.Equal(t, id, stored)
//	}
//}
//
//func TestKeeper_AddLikeToPost_EmptyOwner(t *testing.T) {
//	ctx, k := SetupTestInput()
//
//	like := types.Like{LikeID: defaultLikeID()}
//	err := k.AddLikeToPost(ctx, types.Post{}, like)
//	assert.Error(t, err)
//	assert.Contains(t, err.Result().Log, "Liker and post id must exist")
//}
//
//func TestKeeper_AddLikeToPost_EmptyPostId(t *testing.T) {
//	ctx, k := SetupTestInput()
//
//	like := types.Like{Owner: testPostOwner, LikeID: types.LikeID(0)}
//	err := k.AddLikeToPost(ctx, types.Post{}, like)
//	assert.Error(t, err)
//	assert.Contains(t, err.Result().Log, "Liker and post id must exist")
//}
//
//func TestKeeper_AddLikeToPost_ExistingId(t *testing.T) {
//	ctx, k := SetupTestInput()
//
//	post := types.Post{PostID: defaultPostID()}
//	like := types.Like{Owner: testPostOwner, PostID: post.PostID, LikeID: defaultLikeID()}
//
//	store := ctx.KVStore(k.StoreKey)
//	store.Set([]byte(types.LikesStorePrefix+like.LikeID.String()), k.Cdc.MustMarshalBinaryBare(&like))
//
//	err := k.AddLikeToPost(ctx, post, like)
//	assert.Error(t, err)
//	assert.Contains(t, err.Result().Log, "Like with id 1 already existing")
//}
//
//func TestKeeper_AddLikeToPost_ValidLike(t *testing.T) {
//	ctx, k := SetupTestInput()
//
//	like := types.Like{Owner: testPostOwner, LikeID: defaultLikeID(), PostID: defaultPostID()}
//	post := types.Post{Owner: testPostOwner, PostID: defaultPostID().Next()}
//
//	err := k.AddLikeToPost(ctx, post, like)
//	assert.NoError(t, err)
//
//	var storedLike types.Like
//	var storedPost types.Post
//	store := ctx.KVStore(k.StoreKey)
//	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LikesStorePrefix+like.LikeID.String())), &storedLike)
//	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostStorePrefix+post.PostID.String())), &storedPost)
//
//	assert.Equal(t, post.PostID, storedLike.PostID)
//}
//
//func TestKeeper_AddLikeToPost_UpdatesLastLikeId(t *testing.T) {
//	ctx, k := SetupTestInput()
//
//	post := types.Post{PostID: defaultPostID()}
//	like1 := types.Like{Owner: testPostOwner, LikeID: defaultLikeID(), PostID: post.PostID}
//	like2 := types.Like{Owner: testPostOwner, LikeID: like1.LikeID.Next(), PostID: post.PostID}
//	like3 := types.Like{Owner: testPostOwner, LikeID: like2.LikeID.Next(), PostID: post.PostID}
//
//	_ = k.AddLikeToPost(ctx, post, like1)
//	_ = k.AddLikeToPost(ctx, post, like2)
//	_ = k.AddLikeToPost(ctx, post, like3)
//
//	var lastID uint64
//	store := ctx.KVStore(k.StoreKey)
//	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LastLikeIDStoreKey)), &lastID)
//	assert.Equal(t, uint64(3), lastID)
//}
//
//func TestKeeper_GetLike_NonExistent(t *testing.T) {
//	ctx, k := SetupTestInput()
//
//	_, found := k.GetLike(ctx, defaultLikeID())
//	assert.False(t, found)
//}
//
//func TestKeeper_GetLike_Existent(t *testing.T) {
//	ctx, k := SetupTestInput()
//
//	like := types.Like{Owner: testPostOwner, LikeID: defaultLikeID()}
//	store := ctx.KVStore(k.StoreKey)
//	store.Set([]byte(types.LikesStorePrefix+like.LikeID.String()), k.Cdc.MustMarshalBinaryBare(&like))
//
//	stored, found := k.GetLike(ctx, like.LikeID)
//	assert.True(t, found)
//	assert.Equal(t, like, stored)
//}
