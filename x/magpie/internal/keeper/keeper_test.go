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

func defaultPostId() types.PostId {
	return types.PostId(1)
}

func TestKeeper_GetLastPostId_FirstId(t *testing.T) {
	_, ctx, k := SetupTestInput()
	assert.Equal(t, types.PostId(0), k.GetLastPostId(ctx))
}

func TestKeeper_GetLastPostId_Existing(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	ids := []types.PostId{types.PostId(0), types.PostId(3), types.PostId(18446744073709551615)}

	store := ctx.KVStore(k.storeKey)
	for _, id := range ids {
		store.Set([]byte(types.LastPostIdStoreKey), cdc.MustMarshalBinaryBare(id))
		assert.Equal(t, id, k.GetLastPostId(ctx))
	}
}

func TestKeeper_SetLastPostId(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	ids := []types.PostId{types.PostId(0), types.PostId(3), types.PostId(18446744073709551615)}

	store := ctx.KVStore(k.storeKey)
	for _, id := range ids {
		k.SetLastPostId(ctx, id)
		var stored types.PostId
		cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LastPostIdStoreKey)), &stored)
		assert.Equal(t, id, stored)
	}
}

func TestKeeper_CreatePost_DuplicateId(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	existing := types.Post{PostID: defaultPostId(), Owner: TestPostOwner}

	store := ctx.KVStore(k.storeKey)
	store.Set(k.getPostStoreKey(existing.PostID), cdc.MustMarshalBinaryBare(&existing))

	err := k.CreatePost(ctx, existing)
	assert.Error(t, err)
	assert.Contains(t, err.Result().Log, "Post with id 1 already exists")
}

func TestKeeper_CreatePost_NewPost(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	post := types.Post{PostID: defaultPostId(), Owner: TestPostOwner}
	err := k.CreatePost(ctx, post)
	assert.NoError(t, err)

	var stored types.Post
	store := ctx.KVStore(k.storeKey)
	cdc.MustUnmarshalBinaryBare(store.Get(k.getPostStoreKey(post.PostID)), &stored)
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

	post := types.Post{PostID: defaultPostId(), Owner: TestPostOwner}
	err := k.SavePost(ctx, post)
	assert.NoError(t, err)

	var stored types.Post
	store := ctx.KVStore(k.storeKey)
	cdc.MustUnmarshalBinaryBare(store.Get(k.getPostStoreKey(post.PostID)), &stored)
	assert.Equal(t, post, stored)
}

func TestKeeper_SavePost_UpgradesLastIdProperly(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	post1 := types.Post{PostID: defaultPostId(), Owner: TestPostOwner}
	post2 := types.Post{PostID: post1.PostID.Next(), Owner: TestPostOwner}
	post3 := types.Post{PostID: post2.PostID.Next(), Owner: TestPostOwner}

	_ = k.SavePost(ctx, post1)
	_ = k.SavePost(ctx, post2)
	_ = k.SavePost(ctx, post3)

	var lastId uint64
	store := ctx.KVStore(k.storeKey)
	cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LastPostIdStoreKey)), &lastId)
	assert.Equal(t, uint64(3), lastId)

}

func TestKeeper_GetPost_NonExistent(t *testing.T) {
	_, ctx, k := SetupTestInput()

	_, found := k.GetPost(ctx, defaultPostId())
	assert.False(t, found)
}

func TestKeeper_GetPost_Existent(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	post := types.Post{PostID: defaultPostId(), Owner: TestPostOwner}
	store := ctx.KVStore(k.storeKey)
	store.Set(k.getPostStoreKey(post.PostID), cdc.MustMarshalBinaryBare(&post))

	stored, found := k.GetPost(ctx, post.PostID)
	assert.True(t, found)
	assert.Equal(t, post, stored)
}

func TestKeeper_EditPostMessage(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	post := types.Post{PostID: defaultPostId(), Owner: TestPostOwner, Message: "initial message"}

	store := ctx.KVStore(k.storeKey)
	store.Set(k.getPostStoreKey(post.PostID), cdc.MustMarshalBinaryBare(&post))

	err := k.EditPostMessage(ctx, post, "New message")
	assert.NoError(t, err)

	var updated types.Post
	cdc.MustUnmarshalBinaryBare(store.Get(k.getPostStoreKey(post.PostID)), &updated)
	assert.Equal(t, "New message", updated.Message)
}

func TestKeeper_GetPosts_EmptyList(t *testing.T) {
	_, ctx, k := SetupTestInput()
	posts := k.GetPosts(ctx)
	assert.Empty(t, posts)
}

func TestKeeper_GetPosts_ExistingList(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	post1 := types.Post{PostID: defaultPostId()}
	post2 := types.Post{PostID: post1.PostID.Next()}

	store := ctx.KVStore(k.storeKey)
	store.Set(k.getPostStoreKey(post1.PostID), cdc.MustMarshalBinaryBare(&post1))
	store.Set(k.getPostStoreKey(post2.PostID), cdc.MustMarshalBinaryBare(&post2))

	posts := k.GetPosts(ctx)
	assert.Len(t, posts, 2)
	assert.Contains(t, posts, post1)
	assert.Contains(t, posts, post2)
}

// -------------
// --- Likes
// -------------

func defaultLikeId() types.LikeId {
	return types.LikeId(1)
}

func TestKeeper_GetLastLikeId_FirstId(t *testing.T) {
	_, ctx, k := SetupTestInput()
	assert.Equal(t, types.LikeId(0), k.GetLastLikeId(ctx))
}

func TestKeeper_GetLastLikeId_Existing(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	ids := []types.LikeId{types.LikeId(0), types.LikeId(3), types.LikeId(18446744073709551615)}

	store := ctx.KVStore(k.storeKey)
	for _, id := range ids {
		store.Set([]byte(types.LastLikeIdStoreKey), cdc.MustMarshalBinaryBare(id))
		assert.Equal(t, id, k.GetLastLikeId(ctx))
	}
}

func TestKeeper_SetLastLikeId(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	ids := []types.LikeId{types.LikeId(0), types.LikeId(3), types.LikeId(18446744073709551615)}

	store := ctx.KVStore(k.storeKey)
	for _, id := range ids {
		k.SetLastLikeId(ctx, id)
		var stored types.LikeId
		cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LastLikeIdStoreKey)), &stored)
		assert.Equal(t, id, stored)
	}
}

func TestKeeper_AddLikeToPost_EmptyOwner(t *testing.T) {
	_, ctx, k := SetupTestInput()

	like := types.Like{LikeID: defaultLikeId()}
	err := k.AddLikeToPost(ctx, types.Post{}, like)
	assert.Error(t, err)
	assert.Contains(t, err.Result().Log, "Liker and post id must exist")
}

func TestKeeper_AddLikeToPost_EmptyPostId(t *testing.T) {
	_, ctx, k := SetupTestInput()

	like := types.Like{Owner: TestPostOwner, LikeID: types.LikeId(0)}
	err := k.AddLikeToPost(ctx, types.Post{}, like)
	assert.Error(t, err)
	assert.Contains(t, err.Result().Log, "Liker and post id must exist")
}

func TestKeeper_AddLikeToPost_ExistingId(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	post := types.Post{PostID: defaultPostId()}
	like := types.Like{Owner: TestPostOwner, PostID: post.PostID, LikeID: defaultLikeId()}

	store := ctx.KVStore(k.storeKey)
	store.Set(k.getLikeStoreKey(like.LikeID), cdc.MustMarshalBinaryBare(&like))

	err := k.AddLikeToPost(ctx, post, like)
	assert.Error(t, err)
	assert.Contains(t, err.Result().Log, "Like with id 1 already existing")
}

func TestKeeper_AddLikeToPost_ValidLike(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	like := types.Like{Owner: TestPostOwner, LikeID: defaultLikeId(), PostID: defaultPostId()}
	post := types.Post{Owner: TestPostOwner, PostID: defaultPostId().Next()}

	err := k.AddLikeToPost(ctx, post, like)
	assert.NoError(t, err)

	var storedLike types.Like
	var storedPost types.Post
	store := ctx.KVStore(k.storeKey)
	cdc.MustUnmarshalBinaryBare(store.Get(k.getLikeStoreKey(like.LikeID)), &storedLike)
	cdc.MustUnmarshalBinaryBare(store.Get(k.getPostStoreKey(post.PostID)), &storedPost)

	assert.Equal(t, post.PostID, storedLike.PostID)
	assert.Equal(t, uint(1), storedPost.Likes)
}

func TestKeeper_AddLikeToPost_UpdatesLastLikeId(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	post := types.Post{PostID: defaultPostId()}
	like1 := types.Like{Owner: TestPostOwner, LikeID: defaultLikeId(), PostID: post.PostID}
	like2 := types.Like{Owner: TestPostOwner, LikeID: like1.LikeID.Next(), PostID: post.PostID}
	like3 := types.Like{Owner: TestPostOwner, LikeID: like2.LikeID.Next(), PostID: post.PostID}

	_ = k.AddLikeToPost(ctx, post, like1)
	_ = k.AddLikeToPost(ctx, post, like2)
	_ = k.AddLikeToPost(ctx, post, like3)

	var lastId uint64
	store := ctx.KVStore(k.storeKey)
	cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LastLikeIdStoreKey)), &lastId)
	assert.Equal(t, uint64(3), lastId)
}

func TestKeeper_GetLike_NonExistent(t *testing.T) {
	_, ctx, k := SetupTestInput()

	_, found := k.GetLike(ctx, defaultLikeId())
	assert.False(t, found)
}

func TestKeeper_GetLike_Existent(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	like := types.Like{Owner: TestPostOwner, LikeID: defaultLikeId()}
	store := ctx.KVStore(k.storeKey)
	store.Set(k.getLikeStoreKey(like.LikeID), cdc.MustMarshalBinaryBare(&like))

	stored, found := k.GetLike(ctx, like.LikeID)
	assert.True(t, found)
	assert.Equal(t, like, stored)
}

// --------------
// --- Sessions
// --------------

func defaultSessionId() types.SessionId {
	return types.SessionId(1)
}

func TestKeeper_GetLastSessionId_FirstId(t *testing.T) {
	_, ctx, k := SetupTestInput()
	assert.Equal(t, types.SessionId(0), k.GetLastSessionId(ctx))
}

func TestKeeper_GetLastSessionId_Existing(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	ids := []types.SessionId{types.SessionId(0), types.SessionId(3), types.SessionId(18446744073709551615)}

	store := ctx.KVStore(k.storeKey)
	for _, id := range ids {
		store.Set([]byte(types.LastSessionIdStoreKey), cdc.MustMarshalBinaryBare(id))
		assert.Equal(t, id, k.GetLastSessionId(ctx))
	}
}

func TestKeeper_SetLastSessionId(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	ids := []types.SessionId{types.SessionId(0), types.SessionId(3), types.SessionId(18446744073709551615)}

	store := ctx.KVStore(k.storeKey)
	for _, id := range ids {
		k.SetLastSessionId(ctx, id)
		var stored types.SessionId
		cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LastSessionIdStoreKey)), &stored)
		assert.Equal(t, id, stored)
	}
}

func TestKeeper_CreateSession_ExistingId(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	session := types.Session{SessionID: defaultSessionId()}

	store := ctx.KVStore(k.storeKey)
	store.Set(k.getSessionStoreKey(session.SessionID), cdc.MustMarshalBinaryBare(&session))

	err := k.CreateSession(ctx, session)
	assert.Error(t, err)
	assert.Contains(t, err.Result().Log, "Session with id 1 already exists")
}

func TestKeeper_CreatePost_ValidSession(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	session := types.Session{SessionID: defaultSessionId(), Owner: TestPostOwner}
	err := k.CreateSession(ctx, session)
	assert.NoError(t, err)

	var stored types.Session
	store := ctx.KVStore(k.storeKey)
	cdc.MustUnmarshalBinaryBare(store.Get(k.getSessionStoreKey(session.SessionID)), &stored)
	assert.Equal(t, session, stored)
}

func TestKeeper_SaveSession_EmptyOwner(t *testing.T) {
	_, ctx, k := SetupTestInput()

	session := types.Session{SessionID: defaultSessionId()}
	err := k.SaveSession(ctx, session)
	assert.Error(t, err)
	assert.Contains(t, err.Result().Log, "Owner address cannot be empty")
}

func TestKeeper_SaveSession_ValidSession(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	session := types.Session{Owner: TestPostOwner, SessionID: defaultSessionId()}

	err := k.SaveSession(ctx, session)
	assert.NoError(t, err)

	var stored types.Session
	store := ctx.KVStore(k.storeKey)
	cdc.MustUnmarshalBinaryBare(store.Get(k.getSessionStoreKey(session.SessionID)), &stored)
	assert.Equal(t, session, stored)
}

func TestKeeper_GetSession_NonExistent(t *testing.T) {
	_, ctx, k := SetupTestInput()

	_, found := k.GetSession(ctx, defaultSessionId())
	assert.False(t, found)
}

func TestKeeper_GetSession_Existent(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	session := types.Session{Owner: TestPostOwner, SessionID: defaultSessionId()}

	store := ctx.KVStore(k.storeKey)
	store.Set(k.getSessionStoreKey(session.SessionID), cdc.MustMarshalBinaryBare(&session))

	stored, found := k.GetSession(ctx, session.SessionID)
	assert.True(t, found)
	assert.Equal(t, session, stored)
}

func TestKeeper_EditSessionExpiration(t *testing.T) {
	cdc, ctx, k := SetupTestInput()

	session := types.Session{Owner: TestPostOwner, SessionID: defaultSessionId()}

	store := ctx.KVStore(k.storeKey)
	store.Set(k.getSessionStoreKey(session.SessionID), cdc.MustMarshalBinaryBare(&session))

	location, _ := time.LoadLocation("UTC")
	expiration := time.Date(2017, 10, 20, 11, 45, 32, 0, location)
	err := k.EditSessionExpiration(ctx, session, expiration)
	assert.NoError(t, err)

	var stored types.Session
	cdc.MustUnmarshalBinaryBare(store.Get(k.getSessionStoreKey(session.SessionID)), &stored)
	assert.Equal(t, expiration, stored.Expiry)
}
