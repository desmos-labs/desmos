package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	StoreKey sdk.StoreKey // Unexposed key to access store from sdk.Context
	Cdc      *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the magpie Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		StoreKey: storeKey,
		Cdc:      cdc,
	}
}

// -------------
// --- Posts
// -------------

func (k Keeper) getPostStoreKey(postID types.PostID) []byte {
	return []byte(types.PostStorePrefix + postID.String())
}

// GetLastPostID returns the last post id that has been used
func (k Keeper) GetLastPostID(ctx sdk.Context) types.PostID {
	store := ctx.KVStore(k.StoreKey)
	if !store.Has([]byte(types.LastPostIDStoreKey)) {
		return types.PostID(0)
	}

	var id types.PostID
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LastPostIDStoreKey)), &id)
	return id
}

// SetLastPostID allows to set the last used post id
func (k Keeper) SetLastPostID(ctx sdk.Context, id types.PostID) {
	store := ctx.KVStore(k.StoreKey)
	store.Set([]byte(types.LastPostIDStoreKey), k.Cdc.MustMarshalBinaryBare(&id))
}

// CreatePost allows to create a new post checking for any id conflict with exiting posts
func (k Keeper) CreatePost(ctx sdk.Context, post types.Post) sdk.Error {
	if _, exists := k.GetPost(ctx, post.PostID); exists {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Post with id %s already exists", post.PostID))
	}

	return k.SavePost(ctx, post)
}

// SavePost allows to save the given post inside the current context
func (k Keeper) SavePost(ctx sdk.Context, post types.Post) sdk.Error {
	if post.Owner.Empty() {
		return sdk.ErrInvalidAddress("Post owner cannot be empty")
	}

	store := ctx.KVStore(k.StoreKey)
	store.Set(k.getPostStoreKey(post.PostID), k.Cdc.MustMarshalBinaryBare(&post))

	// Save the last post id
	k.SetLastPostID(ctx, post.PostID)

	return nil
}

// GetPost returns the post having the given id inside the current context.
func (k Keeper) GetPost(ctx sdk.Context, id types.PostID) (post types.Post, found bool) {
	store := ctx.KVStore(k.StoreKey)

	key := k.getPostStoreKey(id)
	if !store.Has(key) {
		return types.NewPost(), false
	}

	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &post)
	return post, true
}

// GetPosts returns the list of all the posts that are stored into the current state
func (k Keeper) GetPosts(ctx sdk.Context) []types.Post {
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.PostStorePrefix))

	var posts []types.Post
	for ; iterator.Valid(); iterator.Next() {
		var post types.Post
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &post)
		posts = append(posts, post)
	}

	return posts
}

// -------------
// --- Likes
// -------------

func (k Keeper) getLikeStoreKey(id types.LikeID) []byte {
	return []byte(types.LikeStorePrefix + id.String())
}

// GetLastLikeID returns the last like id that has been used
func (k Keeper) GetLastLikeID(ctx sdk.Context) types.LikeID {
	store := ctx.KVStore(k.StoreKey)
	if !store.Has([]byte(types.LastLikeIDStoreKey)) {
		return types.LikeID(0)
	}

	var id types.LikeID
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LastLikeIDStoreKey)), &id)
	return id
}

// SetLastLikeID allows to set the last used like id
func (k Keeper) SetLastLikeID(ctx sdk.Context, id types.LikeID) {
	store := ctx.KVStore(k.StoreKey)
	store.Set([]byte(types.LastLikeIDStoreKey), k.Cdc.MustMarshalBinaryBare(&id))
}

// AddLikeToPost allows to add a new like to a given post
func (k Keeper) AddLikeToPost(ctx sdk.Context, post types.Post, like types.Like) sdk.Error {
	// Set the correct post id inside the like
	like.PostID = post.PostID

	// Store the like and update the last like id
	if err := k.SaveLike(ctx, like); err != nil {
		return err
	}
	k.SetLastLikeID(ctx, like.LikeID)

	return nil
}

// SaveLike allows to save the given like inside the store
func (k Keeper) SaveLike(ctx sdk.Context, like types.Like) sdk.Error {
	if like.Owner.Empty() || !like.PostID.Valid() {
		return sdk.ErrUnauthorized("Liker and post id must exist.")
	}

	// Check for any pre-existing likes with the same id
	if _, found := k.GetLike(ctx, like.LikeID); found {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Like with id %s already existing", like.LikeID))
	}

	store := ctx.KVStore(k.StoreKey)
	store.Set(k.getLikeStoreKey(like.LikeID), k.Cdc.MustMarshalBinaryBare(&like))

	return nil
}

// GetLike returns the like having the given id
func (k Keeper) GetLike(ctx sdk.Context, id types.LikeID) (like types.Like, found bool) {
	store := ctx.KVStore(k.StoreKey)

	key := k.getLikeStoreKey(id)
	if !store.Has(key) {
		return types.NewLike(), false
	}

	bz := store.Get(key)
	k.Cdc.MustUnmarshalBinaryBare(bz, &like)
	return like, true
}

// GetLikes allows to returns the list of likes that have been stored inside the given context
func (k Keeper) GetLikes(ctx sdk.Context) []types.Like {
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.LikeStorePrefix))

	var likes []types.Like
	for ; iterator.Valid(); iterator.Next() {
		var like types.Like
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &like)
		likes = append(likes, like)
	}

	return likes
}
