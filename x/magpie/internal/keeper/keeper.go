package keeper

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/kwunyeung/desmos/x/magpie/internal/types"
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

func (k Keeper) getPostStoreKey(postId types.PostId) []byte {
	return []byte(types.PostStorePrefix + postId.String())
}

// GetLastPostId returns the last post id that has been used
func (k Keeper) GetLastPostId(ctx sdk.Context) (id types.PostId) {
	store := ctx.KVStore(k.storeKey)
	k.cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LastPostIdStoreKey)), &id)
	return id
}

// SetLastPostId allows to set the last used post id
func (k Keeper) SetLastPostId(ctx sdk.Context, id types.PostId) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(types.LastPostIdStoreKey), k.cdc.MustMarshalBinaryBare(&id))
}

// CreatePost allows to create a new post checking for any id conflict with exiting posts
func (k Keeper) CreatePost(ctx sdk.Context, post types.Post) sdk.Error {
	if _, exists := k.GetPost(ctx, post.Id); exists {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Post with id %s already exists", post.Id))
	}

	return k.SavePost(ctx, post)
}

// SavePost allows to save the given post inside the current context
func (k Keeper) SavePost(ctx sdk.Context, post types.Post) sdk.Error {
	if post.Owner.Empty() {
		return sdk.ErrInvalidAddress("Post owner cannot be empty")
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(k.getPostStoreKey(post.Id), k.cdc.MustMarshalBinaryBare(&post))

	// Save the last post id
	k.SetLastPostId(ctx, post.Id)

	return nil
}

// GetPost returns the post having the given id inside the current context.
func (k Keeper) GetPost(ctx sdk.Context, id types.PostId) (post types.Post, found bool) {
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

// GetPosts returns the list of all the posts that are stored into the current state
func (k Keeper) GetPosts(ctx sdk.Context) []types.Post {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.PostStorePrefix))

	var posts []types.Post
	for ; iterator.Valid(); iterator.Next() {
		var post types.Post
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &post)
		posts = append(posts, post)
	}

	return posts
}

// -------------
// --- Likes
// -------------

func (k Keeper) getLikeStoreKey(id types.LikeId) []byte {
	return []byte(types.LikeStorePrefix + id.String())
}

// GetLastLikeId returns the last like id that has been used
func (k Keeper) GetLastLikeId(ctx sdk.Context) (id types.LikeId) {
	store := ctx.KVStore(k.storeKey)
	k.cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LastLikeIdStoreKey)), &id)
	return id
}

// SetLastLikeId allows to set the last used like id
func (k Keeper) SetLastLikeId(ctx sdk.Context, id types.LikeId) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(types.LastLikeIdStoreKey), k.cdc.MustMarshalBinaryBare(&id))
}

// AddLikeToPost allows to add a new like to a given post
func (k Keeper) AddLikeToPost(ctx sdk.Context, post types.Post, like types.Like) sdk.Error {
	// Set the correct post id inside the like
	like.PostId = post.Id

	// Store the like and update the last like id
	if err := k.SaveLike(ctx, like); err != nil {
		return err
	}
	k.SetLastLikeId(ctx, like.Id)

	// Update the likes counter and save the post
	post.Likes = post.Likes + 1
	return k.SavePost(ctx, post)
}

// SaveLike allows to save the given like inside the store
func (k Keeper) SaveLike(ctx sdk.Context, like types.Like) sdk.Error {
	if like.Owner.Empty() || !like.PostId.Valid() {
		return sdk.ErrUnauthorized("Liker and post id must exist.")
	}

	// Check for any pre-existing likes with the same id
	if _, found := k.GetLike(ctx, like.Id); found {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Like with id %s already existing", like.Id))
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(k.getLikeStoreKey(like.Id), k.cdc.MustMarshalBinaryBare(&like))

	return nil
}

// GetLike returns the like having the given id
func (k Keeper) GetLike(ctx sdk.Context, id types.LikeId) (like types.Like, found bool) {
	store := ctx.KVStore(k.storeKey)

	key := k.getLikeStoreKey(id)
	if !store.Has(key) {
		return types.NewLike(), false
	}

	bz := store.Get(key)
	k.cdc.MustUnmarshalBinaryBare(bz, &like)
	return like, true
}

// GetLikes allows to returns the list of likes that have been stored inside the given context
func (k Keeper) GetLikes(ctx sdk.Context) []types.Like {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.LikeStorePrefix))

	var likes []types.Like
	for ; iterator.Valid(); iterator.Next() {
		var like types.Like
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &like)
		likes = append(likes, like)
	}

	return likes
}

// -------------
// --- Sessions
// -------------

func (k Keeper) getSessionStoreKey(id types.SessionId) []byte {
	return []byte(types.SessionStorePrefix + id.String())
}

// GetLastLikeId returns the last like id that has been used
func (k Keeper) GetLastSessionId(ctx sdk.Context) (id types.SessionId) {
	store := ctx.KVStore(k.storeKey)
	k.cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LastSessionIdStoreKey)), &id)
	return id
}

// SetLastSessionId allows to set the last used like id
func (k Keeper) SetLastSessionId(ctx sdk.Context, id types.SessionId) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(types.LastSessionIdStoreKey), k.cdc.MustMarshalBinaryBare(&id))
}

// CreateSession allows to create a new session checking that no other session
// with the same id already exist
func (k Keeper) CreateSession(ctx sdk.Context, session types.Session) sdk.Error {
	// Check for any previously existing session
	if _, found := k.GetSession(ctx, session.Id); found {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Session with id %s already exists", session.Id))
	}

	return k.SaveSession(ctx, session)
}

// SaveSession allows to save a session inside the given context
func (k Keeper) SaveSession(ctx sdk.Context, session types.Session) sdk.Error {
	if session.Owner.Empty() {
		return sdk.ErrInvalidAddress("Owner address cannot be empty")
	}

	// Save the session
	store := ctx.KVStore(k.storeKey)
	store.Set(k.getSessionStoreKey(session.Id), k.cdc.MustMarshalBinaryBare(session))

	// Update the last used session id
	k.SetLastSessionId(ctx, session.Id)

	return nil
}

// GetSession returns the session having the specified id
func (k Keeper) GetSession(ctx sdk.Context, id types.SessionId) (session types.Session, found bool) {
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

// GetSessions returns the list of all the sessions present inside the current context
func (k Keeper) GetSessions(ctx sdk.Context) []types.Session {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.SessionStorePrefix))

	var sessions []types.Session
	for ; iterator.Valid(); iterator.Next() {
		var session types.Session
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &session)
		sessions = append(sessions, session)
	}

	return sessions
}
