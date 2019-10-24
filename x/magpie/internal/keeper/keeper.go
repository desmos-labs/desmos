package keeper

import (
	"fmt"
	"strings"
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
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey, coinKeeper bank.Keeper) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

// -------------
// --- Posts
// -------------

func (k Keeper) getPostStoreKey(postId string) []byte {
	return []byte(types.PostStorePrefix + postId)
}

// CreatePost allows to create a new post checking for any id conflict with exiting posts
func (k Keeper) CreatePost(ctx sdk.Context, post types.Post) sdk.Error {
	if _, exists := k.GetPost(ctx, post.ID); exists {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Post with id %s already exists", post.ID))
	}

	return k.SavePost(ctx, post)
}

// SavePost allows to save the given post inside the current context
func (k Keeper) SavePost(ctx sdk.Context, post types.Post) sdk.Error {
	if post.Owner.Empty() {
		return sdk.ErrInvalidAddress("Post owner cannot be empty")
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(k.getPostStoreKey(post.ID), k.cdc.MustMarshalBinaryBare(&post))
	return nil
}

// GetPost returns the post having the given id inside the current context.
func (k Keeper) GetPost(ctx sdk.Context, id string) (post types.Post, found bool) {
	store := ctx.KVStore(k.storeKey)

	key := k.getPostStoreKey(id)
	if !store.Has(key) {
		return types.NewPost(), false
	}

	k.cdc.MustUnmarshalBinaryBare(store.Get(key), &post)
	return post, true
}

// EditPosts allows to edit the message associated with the given post
func (k Keeper) EditPostMessage(ctx sdk.Context, post types.Post, message string) sdk.Error {
	post.Message = message
	return k.SavePost(ctx, post)
}

// GetPostsIterator returns an iterator over the whole set of posts
func (k Keeper) GetPostsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.PostStorePrefix))
}

// -------------
// --- Likes
// -------------

func (k Keeper) getLikeStoreKey(id string) []byte {
	return []byte(types.LikeStorePrefix + id)
}

// SavePostLike allows to save a new like to the given post
func (k Keeper) SavePostLike(ctx sdk.Context, post types.Post, like types.Like) sdk.Error {
	if like.Owner.Empty() || len(strings.TrimSpace(like.PostID)) == 0 {
		return sdk.ErrUnauthorized("Liker and post id must exist.")
	}

	store := ctx.KVStore(k.storeKey)

	// Check for any pre-existing likes with the same id
	if store.Has(k.getLikeStoreKey(like.ID)) {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Like with id %s already existing", like.ID))
	}

	// Set the like data
	like.PostID = post.ID

	// store the like
	store.Set(k.getLikeStoreKey(like.ID), k.cdc.MustMarshalBinaryBare(like))

	// update the likes counter
	post.Likes = post.Likes + 1
	return k.SavePost(ctx, post)
}

// GetLike returns the like having the given id
func (k Keeper) GetLike(ctx sdk.Context, id string) (like types.Like, found bool) {
	store := ctx.KVStore(k.storeKey)

	key := k.getLikeStoreKey(id)
	if !store.Has(key) {
		return types.NewLike(), false
	}

	bz := store.Get(key)
	k.cdc.MustUnmarshalBinaryBare(bz, &like)
	return like, true
}

// -------------
// --- Sessions
// -------------

func (k Keeper) getSessionStoreKey(id string) []byte {
	return []byte(types.SessionStorePrefix + id)
}

// CreateSession allows to create a new session checking that no other session
// with the same id already exist
func (k Keeper) CreateSession(ctx sdk.Context, session types.Session) sdk.Error {
	// Check for any previously existing session
	if _, found := k.GetSession(ctx, session.ID); found {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Session with id %s already exists", session.ID))
	}

	return k.SaveSession(ctx, session)
}

// SaveSession allows to save a session inside the given context
func (k Keeper) SaveSession(ctx sdk.Context, session types.Session) sdk.Error {
	if session.Owner.Empty() {
		return sdk.ErrInvalidAddress("Owner address cannot be empty")
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(k.getSessionStoreKey(session.ID), k.cdc.MustMarshalBinaryBare(session))
	return nil
}

// GetSession returns the session having the specified id
func (k Keeper) GetSession(ctx sdk.Context, id string) (session types.Session, found bool) {
	store := ctx.KVStore(k.storeKey)

	key := k.getSessionStoreKey(id)
	if !store.Has(key) {
		return types.NewSession(), false
	}

	bz := store.Get(key)
	k.cdc.MustUnmarshalBinaryBare(bz, &session)
	return session, true
}

// EditSessionExpiration allows to edit the expiration time of the given session
func (k Keeper) EditSessionExpiration(ctx sdk.Context, session types.Session, expiry time.Time) sdk.Error {
	session.Expiry = expiry
	return k.SaveSession(ctx, session)
}
