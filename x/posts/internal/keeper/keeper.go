package keeper

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"

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
func (k Keeper) SavePost(ctx sdk.Context, post types.Post) {
	store := ctx.KVStore(k.StoreKey)

	// Save the post
	store.Set([]byte(types.PostStorePrefix+post.PostID.String()), k.Cdc.MustMarshalBinaryBare(&post))

	// Set the last post id only if the current post has a greater one than the last one stored
	if post.PostID > k.GetLastPostID(ctx) {
		store.Set([]byte(types.LastPostIDStoreKey), k.Cdc.MustMarshalBinaryBare(&post.PostID))
	}

	// Save the comments to the parent post, if it is valid
	if post.ParentID.Valid() {
		parentCommentsKey := []byte(types.PostCommentsStorePrefix + post.ParentID.String())

		var commentsIDs types.PostIDs
		k.Cdc.MustUnmarshalBinaryBare(store.Get(parentCommentsKey), &commentsIDs)

		if editedIDs, appended := commentsIDs.AppendIfMissing(post.PostID); appended {
			store.Set(parentCommentsKey, k.Cdc.MustMarshalBinaryBare(&editedIDs))
		}
	}
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

// GetPostChildrenIDs returns the IDs of all the children posts associated to the post
// having the given postID
// nolint: interfacer
func (k Keeper) GetPostChildrenIDs(ctx sdk.Context, postID types.PostID) types.PostIDs {
	store := ctx.KVStore(k.StoreKey)

	var postIDs types.PostIDs
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostCommentsStorePrefix+postID.String())), &postIDs)
	return postIDs
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

// GetPostsFiltered retrieves posts filtered by a given set of params which
// include pagination parameters along with the creator address, the parent id and the creation time.
//
// NOTE: If no filters are provided, all posts will be returned in paginated
// form.
func (k Keeper) GetPostsFiltered(ctx sdk.Context, params types.QueryPostsParams) types.Posts {
	posts := k.GetPosts(ctx)
	filteredPosts := make(types.Posts, 0, len(posts))

	for _, p := range posts {
		matchParentID, matchCreationTime, matchAllowsComments, matchSubspace, matchCreator := true, true, true, true, true

		// match parent id if valid
		if params.ParentID != nil {
			matchParentID = params.ParentID.Equals(p.ParentID)
		}

		// match creation time if valid height
		if params.CreationTime.GTE(sdk.ZeroInt()) {
			matchCreationTime = params.CreationTime.Equal(p.Created)
		}

		// match allows comments
		if params.AllowsComments != nil {
			matchAllowsComments = *params.AllowsComments == p.AllowsComments
		}

		// match subspace if provided
		if len(params.Subspace) > 0 {
			matchSubspace = params.Subspace == p.Subspace
		}

		// match creator address (if supplied)
		if len(params.Creator) > 0 {
			matchCreator = params.Creator.Equals(p.Creator)
		}

		if matchParentID && matchCreationTime && matchAllowsComments && matchSubspace && matchCreator {
			filteredPosts = append(filteredPosts, p)
		}
	}

	// Default page
	page := params.Page
	if page == 0 {
		page = 1
	}

	start, end := client.Paginate(len(filteredPosts), page, params.Limit, 100)
	if start < 0 || end < 0 {
		filteredPosts = types.Posts{}
	} else {
		filteredPosts = filteredPosts[start:end]
	}

	return filteredPosts
}

// -------------
// --- MediaPost
// -------------

// SaveMediaProvider add the given provider to the supported providers' list.
func (k Keeper) SaveMediaProvider(ctx sdk.Context, provider string) {
	store := ctx.KVStore(k.StoreKey)

	mediaProviders := k.GetMediaProviders(ctx)

	if updatedProviders, appended := mediaProviders.AppendIfMissing(provider); appended {
		store.Set([]byte(types.MediaProvidersStoreKey), k.Cdc.MustMarshalBinaryBare(&updatedProviders))
	}
}

// GetMediaProviders returns the lists of all supported media providers.
func (k Keeper) GetMediaProviders(ctx sdk.Context) (providers types.Strings) {
	store := ctx.KVStore(k.StoreKey)

	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.MediaProvidersStoreKey)), &providers)

	if len(providers) == 0 {
		return types.Strings{}
	}

	return providers
}

// IsSupportedProvider returns true if all the providers of medias are supported ones, false otherwise.
func (k Keeper) IsSupportedProvider(supportedProviders types.Strings, medias types.PostMedias) bool {
	for _, media := range medias {
		if !supportedProviders.Contains(media.Provider) {
			return false
		}
	}

	return true
}

// SaveMediaMimeType add the given mimeType to the supported mimeTypes' list.
func (k Keeper) SaveMediaMimeType(ctx sdk.Context, mimeType string) {
	store := ctx.KVStore(k.StoreKey)

	mediaMimeTypes := k.GetMediaMimeTypes(ctx)

	if updatedMimeTypes, appended := mediaMimeTypes.AppendIfMissing(mimeType); appended {
		store.Set([]byte(types.MediaMimeTypeStoreKey), k.Cdc.MustMarshalBinaryBare(&updatedMimeTypes))
	}
}

// GetMediaMimeTypes returns the lists of all supported media mime types.
func (k Keeper) GetMediaMimeTypes(ctx sdk.Context) (mimeTypes types.Strings) {
	store := ctx.KVStore(k.StoreKey)

	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.MediaMimeTypeStoreKey)), &mimeTypes)

	if len(mimeTypes) == 0 {
		return types.Strings{}
	}

	return mimeTypes
}

// IsSupportedMimeType returns true if all the mime types of medias are supported ones, false otherwise.
func (k Keeper) IsSupportedMimeType(supportedMimeTypes types.Strings, medias types.PostMedias) bool {
	for _, media := range medias {
		if !supportedMimeTypes.Contains(media.Provider) {
			return false
		}
	}

	return true
}

// getMediaPostStoreKey returns the MediaPostStorePrefix with MediaPost ID's appended
func (k Keeper) getMediaPostStoreKey(postID types.PostID) []byte {
	return []byte(types.MediaPostStorePrefix + postID.String())
}

// SaveMediaPost allows to save the given media post inside the current context.
// It assumes that the media post is already been validated
func (k Keeper) SaveMediaPost(ctx sdk.Context, mediaPost types.MediaPost) {
	store := ctx.KVStore(k.StoreKey)

	//Save the post inside mediaPost
	k.SavePost(ctx, mediaPost.Post)

	//Save the medias associated with mediaPost.Post ID
	store.Set(k.getMediaPostStoreKey(mediaPost.Post.PostID), k.Cdc.MustMarshalBinaryBare(&mediaPost.Medias))
}

func (k Keeper) GetPostMedias(ctx sdk.Context, id types.PostID) (medias types.PostMedias, found bool) {
	store := ctx.KVStore(k.StoreKey)

	key := k.getMediaPostStoreKey(id)

	if !store.Has(key) {
		return types.PostMedias{}, false
	}

	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &medias)

	return medias, true
}

/*
func (k Keeper) SavePostMedia(ctx sdk.Context, mediaPost types.MediaPost) {
	store := ctx.KVStore(k.StoreKey)

	store.Set([]byte(types.MediaPostStorePrefix+mediaPost.Post.PostID.String()), k.Cdc.MustMarshalBinaryBare(&mediaPost))

	// Set the last post id only if the current post has a greater one than the last one stored
	if mediaPost.Post.PostID > k.GetLastPostID(ctx) {
		store.Set([]byte(types.LastPostIDStoreKey), k.Cdc.MustMarshalBinaryBare(&mediaPost.Post.PostID))
	}

	// Save the comments to the parent post, if it is valid
	if mediaPost.Post.ParentID.Valid() {
		parentCommentsKey := []byte(types.PostCommentsStorePrefix + mediaPost.Post.ParentID.String())

		var commentsIDs types.PostIDs
		k.Cdc.MustUnmarshalBinaryBare(store.Get(parentCommentsKey), &commentsIDs)

		if editedIDs, appended := commentsIDs.AppendIfMissing(mediaPost.Post.PostID); appended {
			store.Set(parentCommentsKey, k.Cdc.MustMarshalBinaryBare(&editedIDs))
		}
	}
}
*/

// -------------
// --- Reactions
// -------------

// SaveReaction allows to save the given reaction inside the store.
// It assumes that the given reaction is valid.
// If another reaction from the same user for the same post and with the same value exists, returns an expError.
// nolint: interfacer
func (k Keeper) SaveReaction(ctx sdk.Context, postID types.PostID, reaction types.Reaction) sdk.Error {
	store := ctx.KVStore(k.StoreKey)
	key := []byte(types.PostReactionsStorePrefix + postID.String())

	// Get the existent reactions
	var reactions types.Reactions
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &reactions)

	// Check for double reactions
	if reactions.ContainsReactionFrom(reaction.Owner, reaction.Value) {
		msg := fmt.Sprintf("%s has already reacted with %s to the post with id %s",
			reaction.Owner, reaction.Value, postID)
		return sdk.ErrUnknownRequest(msg)
	}

	// Save the new reaction
	reactions = append(reactions, reaction)
	store.Set(key, k.Cdc.MustMarshalBinaryBare(&reactions))

	return nil
}

// RemoveReaction removes the reaction from the given user from the post having the
// given postID. If no reaction with the same value was previously added from the given user, an expError
// is returned.
// nolint: interfacer
func (k Keeper) RemoveReaction(ctx sdk.Context, postID types.PostID, user sdk.AccAddress, value string) sdk.Error {
	store := ctx.KVStore(k.StoreKey)
	key := []byte(types.PostReactionsStorePrefix + postID.String())

	// Get the existing reactions
	var reactions types.Reactions
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &reactions)

	// Check if the user exists
	if !reactions.ContainsReactionFrom(user, value) {
		msg := fmt.Sprintf("Cannot remove the reaction with value %s from user %s as it does not exist",
			value, user)
		return sdk.ErrUnauthorized(msg)
	}

	// Remove and save the reactions list
	if newLikes, edited := reactions.RemoveReaction(user, value); edited {
		if len(newLikes) == 0 {
			store.Delete(key)
		} else {
			store.Set(key, k.Cdc.MustMarshalBinaryBare(&newLikes))
		}
	}

	return nil
}

// GetPostReactions returns the list of reactions that has been associated to the post having the given id
// nolint: interfacer
func (k Keeper) GetPostReactions(ctx sdk.Context, postID types.PostID) types.Reactions {
	store := ctx.KVStore(k.StoreKey)

	var reactions types.Reactions
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostReactionsStorePrefix+postID.String())), &reactions)
	return reactions
}

// GetReactions allows to returns the list of reactions that have been stored inside the given context
func (k Keeper) GetReactions(ctx sdk.Context) map[types.PostID]types.Reactions {
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.PostReactionsStorePrefix))

	reactionsData := map[types.PostID]types.Reactions{}
	for ; iterator.Valid(); iterator.Next() {
		var postLikes types.Reactions
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &postLikes)
		postID, _ := types.ParsePostID(strings.TrimPrefix(string(iterator.Key()), types.PostReactionsStorePrefix))
		reactionsData[postID] = postLikes
	}

	return reactionsData
}
