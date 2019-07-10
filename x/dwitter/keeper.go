package dwitter

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	coinKeeper bank.Keeper

	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the dwitter Keeper
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

func (k Keeper) SetPost(ctx sdk.Context, post Post) {
	if post.Owner.Empty() {
		return
	}

	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(post.ID), k.cdc.MustMarshalBinaryBare(post))
}

func (k Keeper) GetPost(ctx sdk.Context, id string) Post {
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(id)) {
		return NewPost()
	}

	bz := store.Get([]byte(id))
	var post Post
	k.cdc.MustUnmarshalBinaryBare(bz, &post)
	return post
}

func (k Keeper) EditPost(ctx sdk.Context, id string, message string) {
	post := k.GetPost(ctx, id)
	post.Message = message
	k.SetPost(ctx, post)
}

func (k Keeper) GetPostOwner(ctx sdk.Context, id string) sdk.AccAddress {
	return k.GetPost(ctx, id).Owner
}

// func (k Keeper) GetPostsByOwner(ctx sdk.Context, owner sdk.AccAddress) []Post {
// 	matchingPosts := []Post{}
// 	return matchingPosts
// }

func (k Keeper) GetPostLikes(ctx sdk.Context, id string) uint {
	return k.GetPost(ctx, id).Likes
}

func (k Keeper) AddPostLike(ctx sdk.Context, id string) {
	post := k.GetPost(ctx, id)
	post.Likes = post.Likes + 1
	k.SetPost(ctx, post)
}

func (k Keeper) SetLike(ctx sdk.Context, id string, like Like) {
	if like.Owner.Empty() || (len(like.PostID) == 0) {
		return
	}

	post := k.GetPost(ctx, like.PostID)
	if len(post.ID) == 0 {
		return
	}

	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(id), k.cdc.MustMarshalBinaryBare(like))

	k.AddPostLike(ctx, like.PostID)
}

func (k Keeper) GetLike(ctx sdk.Context, id string) Like {
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(id)) {
		return NewLike()
	}

	bz := store.Get([]byte(id))
	var like Like
	k.cdc.MustUnmarshalBinaryBare(bz, &like)
	return like
}

func (k Keeper) GetPostsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}
