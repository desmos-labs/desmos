package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/kwunyeung/desmos/x/magpie/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	coinKeeper bank.Keeper
	storeKey   sdk.StoreKey // Unexposed key to access store from sdk.Context
	cdc        *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the magpie Keeper
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

// SavePost allows to save the given post inside the current context
func (k Keeper) SavePost(ctx sdk.Context, post types.Post) sdk.Error {
	if post.Owner.Empty() {
		return sdk.ErrInvalidAddress("No address found.")
	}

	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(post.ID), k.cdc.MustMarshalBinaryBare(post))
	return nil
}

// GetPost returns the post having the given id inside the current context.
func (k Keeper) GetPost(ctx sdk.Context, id string) (post types.Post, found bool) {
	store := ctx.KVStore(k.storeKey)

	if !store.Has([]byte(id)) {
		return types.NewPost(), false
	}

	bz := store.Get([]byte(id))
	k.cdc.MustUnmarshalBinaryBare(bz, &post)
	return post, true
}

// EditPosts allows to edit the message associated with the given post
func (k Keeper) EditPostMessage(ctx sdk.Context, post types.Post, message string) sdk.Error {
	post.Message = message
	return k.SavePost(ctx, post)
}

// SavePostLike allows to save a new like to the given post
func (k Keeper) SavePostLike(ctx sdk.Context, post types.Post, like types.Like) sdk.Error {
	if like.Owner.Empty() || (len(like.PostID) == 0) {
		return sdk.ErrUnauthorized("Liker and post id must exist.")
	}

	// store the like
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(like.ID), k.cdc.MustMarshalBinaryBare(like))

	// update the likes counter
	post.Likes = post.Likes + 1
	return k.SavePost(ctx, post)
}

// GetLike returns the like having the given id
func (k Keeper) GetLike(ctx sdk.Context, id string) (like types.Like, found bool) {
	store := ctx.KVStore(k.storeKey)

	if !store.Has([]byte(id)) {
		return types.NewLike(), false
	}

	bz := store.Get([]byte(id))
	k.cdc.MustUnmarshalBinaryBare(bz, &like)
	return like, true
}

// GetPostsIterator returns an iterator over the whole set of posts
func (k Keeper) GetPostsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}

// SaveSession allows to save a session inside the given context
func (k Keeper) SaveSession(ctx sdk.Context, session types.Session) sdk.Error {
	if session.Owner.Empty() {
		return sdk.ErrInvalidAddress("No address found.")
	}

	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(session.ID), k.cdc.MustMarshalBinaryBare(session))
	return nil
}

// GetSession returns the session having the specified id
func (k Keeper) GetSession(ctx sdk.Context, id string) (session types.Session, found bool) {
	store := ctx.KVStore(k.storeKey)

	if !store.Has([]byte(id)) {
		return types.NewSession(), false
	}

	bz := store.Get([]byte(id))
	k.cdc.MustUnmarshalBinaryBare(bz, &session)
	return session, true
}

// EditSessionExpiration allows to edit the expiration time of the given session
func (k Keeper) EditSessionExpiration(ctx sdk.Context, session types.Session, expiry time.Time) sdk.Error {
	session.Expiry = expiry
	return k.SaveSession(ctx, session)
}
