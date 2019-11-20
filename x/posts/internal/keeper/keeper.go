package keeper

import (
	"fmt"
	"strings"

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

// SavePost allows to save the given post inside the current context.
// It assumes that the given post has already been validated.
// If another post has the same ID of the given post, the old post will be overridden
func (k Keeper) SavePost(ctx sdk.Context, post types.Post) sdk.Error {
	store := ctx.KVStore(k.StoreKey)

	// Save the post and set the last post id
	store.Set([]byte(types.PostStorePrefix+post.PostID.String()), k.Cdc.MustMarshalBinaryBare(&post))
	store.Set([]byte(types.LastPostIDStoreKey), k.Cdc.MustMarshalBinaryBare(&post.PostID))
	return nil
}

// GetPost returns the post having the given id inside the current context.
// If no post having the given id can be found inside the current context, false will be returned.
func (k Keeper) GetPost(ctx sdk.Context, id types.PostID) (post types.Post, found bool) {
	store := ctx.KVStore(k.StoreKey)

	key := k.getPostStoreKey(id)
	if !store.Has(key) {
		return types.Post{}, false
	}

	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &post)
	return post, true
}

// GetPosts returns the list of all the posts that are stored into the current state.
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

// SaveLike allows to save the given like inside the store.
// It assumes that the given like is valid.
// If another like from the same owner and for the same post exists, returns an error.
// nolint: interfacer
func (k Keeper) SaveLike(ctx sdk.Context, postID types.PostID, like types.Like) sdk.Error {
	store := ctx.KVStore(k.StoreKey)
	key := []byte(types.LikesStorePrefix + postID.String())

	// Get the existent likes
	var likes types.Likes
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &likes)

	// Check for double likes
	if likes.ContainsOwnerLike(like.Owner) {
		msg := fmt.Sprintf("%s has already liked the post with id %s", like.Owner, postID.String())
		return sdk.ErrUnknownRequest(msg)
	}

	// Save the new like
	likes = append(likes, like)
	store.Set(key, k.Cdc.MustMarshalBinaryBare(&likes))

	return nil
}

// GetLikes allows to returns the list of likes that have been stored inside the given context
func (k Keeper) GetLikes(ctx sdk.Context) map[types.PostID]types.Likes {
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.LikesStorePrefix))

	likesData := map[types.PostID]types.Likes{}
	for ; iterator.Valid(); iterator.Next() {
		var postLikes types.Likes
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &postLikes)
		postID, _ := types.ParsePostID(strings.TrimPrefix(string(iterator.Key()), types.LikesStorePrefix))
		likesData[postID] = postLikes
	}

	return likesData
}
