package keeper

import (
	"testing"

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
