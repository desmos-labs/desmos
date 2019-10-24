package keeper

import (
	"testing"
	"time"

	"github.com/kwunyeung/desmos/x/magpie/internal/types"
	"github.com/stretchr/testify/assert"
)

// -------------
// --- Posts
// -------------

func TestKeeper_CreatePost_DuplicateId(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	existing := types.Post{ID: "post-id", Owner: TestPostOwner}

	store := ctx.KVStore(k.storeKey)
	store.Set(k.getPostStoreKey(existing.ID), cdc.MustMarshalBinaryBare(&existing))

	err := k.CreatePost(ctx, existing)
	assert.Error(t, err)
	assert.Contains(t, err.Result().Log, "Post with id post-id already exists")
}

func TestKeeper_CreatePost_NewPost(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	post := types.Post{ID: "new-post", Owner: TestPostOwner}
	err := k.CreatePost(ctx, post)
	assert.NoError(t, err)

	var stored types.Post
	store := ctx.KVStore(k.storeKey)
	cdc.MustUnmarshalBinaryBare(store.Get(k.getPostStoreKey(post.ID)), &stored)
	assert.Equal(t, post, stored)
}

func TestKeeper_SavePost_InvalidOwner(t *testing.T) {
	_, ctx, k := SetupTestInput()

	post := types.Post{}
	err := k.SavePost(ctx, post)

	assert.Error(t, err)
	assert.Contains(t, err.Result().Log, "Post owner cannot be empty")
}

func TestKeeper_SavePost_ValidPost(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	post := types.Post{ID: "post-id", Owner: TestPostOwner}
	err := k.SavePost(ctx, post)
	assert.NoError(t, err)

	var stored types.Post
	store := ctx.KVStore(k.storeKey)
	cdc.MustUnmarshalBinaryBare(store.Get(k.getPostStoreKey(post.ID)), &stored)
	assert.Equal(t, post, stored)
}

func TestKeeper_GetPost_NonExistent(t *testing.T) {
	_, ctx, k := SetupTestInput()

	_, found := k.GetPost(ctx, "non-existent")
	assert.False(t, found)
}

func TestKeeper_GetPost_Existent(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	post := types.Post{ID: "existent", Owner: TestPostOwner}
	store := ctx.KVStore(k.storeKey)
	store.Set(k.getPostStoreKey(post.ID), cdc.MustMarshalBinaryBare(&post))

	stored, found := k.GetPost(ctx, post.ID)
	assert.True(t, found)
	assert.Equal(t, post, stored)
}

func TestKeeper_EditPostMessage(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	post := types.Post{ID: "editing-post", Owner: TestPostOwner, Message: "initial message"}

	store := ctx.KVStore(k.storeKey)
	store.Set(k.getPostStoreKey(post.ID), cdc.MustMarshalBinaryBare(&post))

	err := k.EditPostMessage(ctx, post, "New message")
	assert.NoError(t, err)

	var updated types.Post
	cdc.MustUnmarshalBinaryBare(store.Get(k.getPostStoreKey(post.ID)), &updated)
	assert.Equal(t, "New message", updated.Message)
}

// -------------
// --- Likes
// -------------

func TestKeeper_SavePostLike_EmptyOwner(t *testing.T) {
	_, ctx, k := SetupTestInput()

	like := types.Like{ID: "like"}
	err := k.SavePostLike(ctx, types.Post{}, like)
	assert.Error(t, err)
	assert.Contains(t, err.Result().Log, "Liker and post id must exist")
}

func TestKeeper_SavePostLike_EmptyPostId(t *testing.T) {
	_, ctx, k := SetupTestInput()

	like := types.Like{Owner: TestPostOwner, ID: " "}
	err := k.SavePostLike(ctx, types.Post{}, like)
	assert.Error(t, err)
	assert.Contains(t, err.Result().Log, "Liker and post id must exist")
}

func TestKeeper_SavePostLike_ExistingId(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	like := types.Like{Owner: TestPostOwner, PostID: "post-id", ID: "like-id"}

	store := ctx.KVStore(k.storeKey)
	store.Set(k.getLikeStoreKey(like.ID), cdc.MustMarshalBinaryBare(&like))

	err := k.SavePostLike(ctx, types.Post{}, like)
	assert.Error(t, err)
	assert.Contains(t, err.Result().Log, "Like with id like-id already existing")
}

func TestKeeper_SavePostLike_ValidLike(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	like := types.Like{Owner: TestPostOwner, ID: "like-id", PostID: "new post id"}
	post := types.Post{Owner: TestPostOwner, ID: "post-id"}

	err := k.SavePostLike(ctx, post, like)
	assert.NoError(t, err)

	var storedLike types.Like
	var storedPost types.Post
	store := ctx.KVStore(k.storeKey)
	cdc.MustUnmarshalBinaryBare(store.Get(k.getLikeStoreKey(like.ID)), &storedLike)
	cdc.MustUnmarshalBinaryBare(store.Get(k.getPostStoreKey(post.ID)), &storedPost)

	assert.Equal(t, post.ID, storedLike.PostID)
	assert.Equal(t, uint(1), storedPost.Likes)
}

func TestKeeper_GetLike_NonExistent(t *testing.T) {
	_, ctx, k := SetupTestInput()

	_, found := k.GetLike(ctx, "non-existent")
	assert.False(t, found)
}

func TestKeeper_GetLike_Existent(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	like := types.Like{Owner: TestPostOwner, ID: "like-id"}
	store := ctx.KVStore(k.storeKey)
	store.Set(k.getLikeStoreKey(like.ID), cdc.MustMarshalBinaryBare(&like))

	stored, found := k.GetLike(ctx, like.ID)
	assert.True(t, found)
	assert.Equal(t, like, stored)
}

// --------------
// --- Sessions
// --------------

func TestKeeper_CreateSession_ExistingId(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	session := types.Session{ID: "session-id"}

	store := ctx.KVStore(k.storeKey)
	store.Set(k.getSessionStoreKey(session.ID), cdc.MustMarshalBinaryBare(&session))

	err := k.CreateSession(ctx, session)
	assert.Error(t, err)
	assert.Contains(t, err.Result().Log, "Session with id session-id already exists")
}

func TestKeeper_CreatePost_ValidSession(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	session := types.Session{ID: "session-id", Owner: TestPostOwner}
	err := k.CreateSession(ctx, session)
	assert.NoError(t, err)

	var stored types.Session
	store := ctx.KVStore(k.storeKey)
	cdc.MustUnmarshalBinaryBare(store.Get(k.getSessionStoreKey(session.ID)), &stored)
	assert.Equal(t, session, stored)
}

func TestKeeper_SaveSession_EmptyOwner(t *testing.T) {
	_, ctx, k := SetupTestInput()

	session := types.Session{ID: "session-id"}
	err := k.SaveSession(ctx, session)
	assert.Error(t, err)
	assert.Contains(t, err.Result().Log, "Owner address cannot be empty")
}

func TestKeeper_SaveSession_ValidSession(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	session := types.Session{Owner: TestPostOwner, ID: "session-id"}

	err := k.SaveSession(ctx, session)
	assert.NoError(t, err)

	var stored types.Session
	store := ctx.KVStore(k.storeKey)
	cdc.MustUnmarshalBinaryBare(store.Get(k.getSessionStoreKey(session.ID)), &stored)
	assert.Equal(t, session, stored)
}

func TestKeeper_GetSession_NonExistent(t *testing.T) {
	_, ctx, k := SetupTestInput()

	_, found := k.GetSession(ctx, "inexistent-session")
	assert.False(t, found)
}

func TestKeeper_GetSession_Existent(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	session := types.Session{Owner: TestPostOwner, ID: "session-id"}

	store := ctx.KVStore(k.storeKey)
	store.Set(k.getSessionStoreKey(session.ID), cdc.MustMarshalBinaryBare(&session))

	stored, found := k.GetSession(ctx, session.ID)
	assert.True(t, found)
	assert.Equal(t, session, stored)
}

func TestKeeper_EditSessionExpiration(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	session := types.Session{Owner: TestPostOwner, ID: "session-id"}

	store := ctx.KVStore(k.storeKey)
	store.Set(k.getSessionStoreKey(session.ID), cdc.MustMarshalBinaryBare(&session))

	location, _ := time.LoadLocation("UTC")
	expiration := time.Date(2017, 10, 20, 11, 45, 32, 0, location)
	err := k.EditSessionExpiration(ctx, session, expiration)
	assert.NoError(t, err)

	var stored types.Session
	cdc.MustUnmarshalBinaryBare(store.Get(k.getSessionStoreKey(session.ID)), &stored)
	assert.Equal(t, expiration, stored.Expiry)
}
