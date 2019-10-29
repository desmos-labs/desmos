package keeper_test

import (
	"testing"

	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/assert"
)

// -------------
// --- Posts
// -------------

func defaultPostID() types.PostID {
	return types.PostID(1)
}

func TestKeeper_GetLastPostId_FirstId(t *testing.T) {
	ctx, k := SetupTestInput()
	assert.Equal(t, types.PostID(0), k.GetLastPostID(ctx))
}

func TestKeeper_GetLastPostId_Existing(t *testing.T) {
	ctx, k := SetupTestInput()

	ids := []types.PostID{types.PostID(0), types.PostID(3), types.PostID(18446744073709551615)}

	store := ctx.KVStore(k.StoreKey)
	for _, id := range ids {
		store.Set([]byte(types.LastPostIDStoreKey), k.Cdc.MustMarshalBinaryBare(id))
		assert.Equal(t, id, k.GetLastPostID(ctx))
	}
}

func TestKeeper_SetLastPostId(t *testing.T) {
	ctx, k := SetupTestInput()

	ids := []types.PostID{types.PostID(0), types.PostID(3), types.PostID(18446744073709551615)}

	store := ctx.KVStore(k.StoreKey)
	for _, id := range ids {
		k.SetLastPostID(ctx, id)
		var stored types.PostID
		k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LastPostIDStoreKey)), &stored)
		assert.Equal(t, id, stored)
	}
}

func TestKeeper_CreatePost_DuplicateId(t *testing.T) {
	ctx, k := SetupTestInput()

	existing := types.Post{PostID: defaultPostID(), Owner: testPostOwner}

	store := ctx.KVStore(k.StoreKey)
	store.Set([]byte(types.PostStorePrefix+existing.PostID.String()), k.Cdc.MustMarshalBinaryBare(&existing))

	err := k.CreatePost(ctx, existing)
	assert.Error(t, err)
	assert.Contains(t, err.Result().Log, "Post with id 1 already exists")
}

func TestKeeper_CreatePost_NewPost(t *testing.T) {
	ctx, k := SetupTestInput()

	post := types.Post{PostID: defaultPostID(), Owner: testPostOwner}
	err := k.CreatePost(ctx, post)
	assert.NoError(t, err)

	var stored types.Post
	store := ctx.KVStore(k.StoreKey)
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostStorePrefix+post.PostID.String())), &stored)
	assert.Equal(t, post, stored)
}

func TestKeeper_SavePost_InvalidOwner(t *testing.T) {
	ctx, k := SetupTestInput()

	post := types.Post{}
	err := k.SavePost(ctx, post)

	assert.Error(t, err)
	assert.Contains(t, err.Result().Log, "Post owner cannot be empty")
}

func TestKeeper_SavePost_ValidPost(t *testing.T) {
	ctx, k := SetupTestInput()

	post := types.Post{PostID: defaultPostID(), Owner: testPostOwner}
	err := k.SavePost(ctx, post)
	assert.NoError(t, err)

	var stored types.Post
	store := ctx.KVStore(k.StoreKey)
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostStorePrefix+post.PostID.String())), &stored)
	assert.Equal(t, post, stored)
}

func TestKeeper_SavePost_UpgradesLastIdProperly(t *testing.T) {
	ctx, k := SetupTestInput()

	post1 := types.Post{PostID: defaultPostID(), Owner: testPostOwner}
	post2 := types.Post{PostID: post1.PostID.Next(), Owner: testPostOwner}
	post3 := types.Post{PostID: post2.PostID.Next(), Owner: testPostOwner}

	_ = k.SavePost(ctx, post1)
	_ = k.SavePost(ctx, post2)
	_ = k.SavePost(ctx, post3)

	var lastID uint64
	store := ctx.KVStore(k.StoreKey)
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LastPostIDStoreKey)), &lastID)
	assert.Equal(t, uint64(3), lastID)

}

func TestKeeper_GetPost_NonExistent(t *testing.T) {
	ctx, k := SetupTestInput()

	_, found := k.GetPost(ctx, defaultPostID())
	assert.False(t, found)
}

func TestKeeper_GetPost_Existent(t *testing.T) {
	ctx, k := SetupTestInput()

	post := types.Post{PostID: defaultPostID(), Owner: testPostOwner}
	store := ctx.KVStore(k.StoreKey)
	store.Set([]byte(types.PostStorePrefix+post.PostID.String()), k.Cdc.MustMarshalBinaryBare(&post))

	stored, found := k.GetPost(ctx, post.PostID)
	assert.True(t, found)
	assert.Equal(t, post, stored)
}

func TestKeeper_GetPosts_EmptyList(t *testing.T) {
	ctx, k := SetupTestInput()
	posts := k.GetPosts(ctx)
	assert.Empty(t, posts)
}

func TestKeeper_GetPosts_ExistingList(t *testing.T) {
	ctx, k := SetupTestInput()

	post1 := types.Post{PostID: defaultPostID()}
	post2 := types.Post{PostID: post1.PostID.Next()}

	store := ctx.KVStore(k.StoreKey)
	store.Set([]byte(types.PostStorePrefix+post1.PostID.String()), k.Cdc.MustMarshalBinaryBare(&post1))
	store.Set([]byte(types.PostStorePrefix+post2.PostID.String()), k.Cdc.MustMarshalBinaryBare(&post2))

	posts := k.GetPosts(ctx)
	assert.Len(t, posts, 2)
	assert.Contains(t, posts, post1)
	assert.Contains(t, posts, post2)
}

// -------------
// --- Likes
// -------------

func defaultLikeID() types.LikeID {
	return types.LikeID(1)
}

func TestKeeper_GetLastLikeId_FirstId(t *testing.T) {
	ctx, k := SetupTestInput()
	assert.Equal(t, types.LikeID(0), k.GetLastLikeID(ctx))
}

func TestKeeper_GetLastLikeId_Existing(t *testing.T) {
	ctx, k := SetupTestInput()

	ids := []types.LikeID{types.LikeID(0), types.LikeID(3), types.LikeID(18446744073709551615)}

	store := ctx.KVStore(k.StoreKey)
	for _, id := range ids {
		store.Set([]byte(types.LastLikeIDStoreKey), k.Cdc.MustMarshalBinaryBare(id))
		assert.Equal(t, id, k.GetLastLikeID(ctx))
	}
}

func TestKeeper_SetLastLikeId(t *testing.T) {
	ctx, k := SetupTestInput()

	ids := []types.LikeID{types.LikeID(0), types.LikeID(3), types.LikeID(18446744073709551615)}

	store := ctx.KVStore(k.StoreKey)
	for _, id := range ids {
		k.SetLastLikeID(ctx, id)
		var stored types.LikeID
		k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LastLikeIDStoreKey)), &stored)
		assert.Equal(t, id, stored)
	}
}

func TestKeeper_AddLikeToPost_EmptyOwner(t *testing.T) {
	ctx, k := SetupTestInput()

	like := types.Like{LikeID: defaultLikeID()}
	err := k.AddLikeToPost(ctx, types.Post{}, like)
	assert.Error(t, err)
	assert.Contains(t, err.Result().Log, "Liker and post id must exist")
}

func TestKeeper_AddLikeToPost_EmptyPostId(t *testing.T) {
	ctx, k := SetupTestInput()

	like := types.Like{Owner: testPostOwner, LikeID: types.LikeID(0)}
	err := k.AddLikeToPost(ctx, types.Post{}, like)
	assert.Error(t, err)
	assert.Contains(t, err.Result().Log, "Liker and post id must exist")
}

func TestKeeper_AddLikeToPost_ExistingId(t *testing.T) {
	ctx, k := SetupTestInput()

	post := types.Post{PostID: defaultPostID()}
	like := types.Like{Owner: testPostOwner, PostID: post.PostID, LikeID: defaultLikeID()}

	store := ctx.KVStore(k.StoreKey)
	store.Set([]byte(types.LikeStorePrefix+like.LikeID.String()), k.Cdc.MustMarshalBinaryBare(&like))

	err := k.AddLikeToPost(ctx, post, like)
	assert.Error(t, err)
	assert.Contains(t, err.Result().Log, "Like with id 1 already existing")
}

func TestKeeper_AddLikeToPost_ValidLike(t *testing.T) {
	ctx, k := SetupTestInput()

	like := types.Like{Owner: testPostOwner, LikeID: defaultLikeID(), PostID: defaultPostID()}
	post := types.Post{Owner: testPostOwner, PostID: defaultPostID().Next()}

	err := k.AddLikeToPost(ctx, post, like)
	assert.NoError(t, err)

	var storedLike types.Like
	var storedPost types.Post
	store := ctx.KVStore(k.StoreKey)
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LikeStorePrefix+like.LikeID.String())), &storedLike)
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostStorePrefix+post.PostID.String())), &storedPost)

	assert.Equal(t, post.PostID, storedLike.PostID)
	assert.Equal(t, uint(1), storedPost.Likes)
}

func TestKeeper_AddLikeToPost_UpdatesLastLikeId(t *testing.T) {
	ctx, k := SetupTestInput()

	post := types.Post{PostID: defaultPostID()}
	like1 := types.Like{Owner: testPostOwner, LikeID: defaultLikeID(), PostID: post.PostID}
	like2 := types.Like{Owner: testPostOwner, LikeID: like1.LikeID.Next(), PostID: post.PostID}
	like3 := types.Like{Owner: testPostOwner, LikeID: like2.LikeID.Next(), PostID: post.PostID}

	_ = k.AddLikeToPost(ctx, post, like1)
	_ = k.AddLikeToPost(ctx, post, like2)
	_ = k.AddLikeToPost(ctx, post, like3)

	var lastID uint64
	store := ctx.KVStore(k.StoreKey)
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LastLikeIDStoreKey)), &lastID)
	assert.Equal(t, uint64(3), lastID)
}

func TestKeeper_GetLike_NonExistent(t *testing.T) {
	ctx, k := SetupTestInput()

	_, found := k.GetLike(ctx, defaultLikeID())
	assert.False(t, found)
}

func TestKeeper_GetLike_Existent(t *testing.T) {
	ctx, k := SetupTestInput()

	like := types.Like{Owner: testPostOwner, LikeID: defaultLikeID()}
	store := ctx.KVStore(k.StoreKey)
	store.Set([]byte(types.LikeStorePrefix+like.LikeID.String()), k.Cdc.MustMarshalBinaryBare(&like))

	stored, found := k.GetLike(ctx, like.LikeID)
	assert.True(t, found)
	assert.Equal(t, like, stored)
}
