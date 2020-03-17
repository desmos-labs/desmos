package keeper

import (
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
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

// GetLastPostID returns the last post id that has been used
func (k Keeper) GetLastPostID(ctx sdk.Context) types.PostID {
	store := ctx.KVStore(k.StoreKey)
	if !store.Has(types.LastPostIDStoreKey) {
		return types.PostID(0)
	}

	var id types.PostID
	k.Cdc.MustUnmarshalBinaryBare(store.Get(types.LastPostIDStoreKey), &id)
	return id
}

// SavePost allows to save the given post inside the current context.
// It assumes that the given post has already been validated.
// If another post has the same ID of the given post, the old post will be overridden
func (k Keeper) SavePost(ctx sdk.Context, post types.Post) {
	store := ctx.KVStore(k.StoreKey)

	// Save the post
	store.Set(types.PostStoreKey(post.PostID), k.Cdc.MustMarshalBinaryBare(&post))

	// Set the last post id only if the current post has a greater one than the last one stored
	if id := post.PostID; id > k.GetLastPostID(ctx) {
		store.Set(types.LastPostIDStoreKey, k.Cdc.MustMarshalBinaryBare(&id))
	}

	// Save the comments to the parent post, if it is valid
	if post.ParentID.Valid() {
		parentCommentsKey := types.PostCommentsStoreKey(post.ParentID)

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

	key := types.PostStoreKey(id)
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
	k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostCommentsStoreKey(postID)), &postIDs)
	return postIDs
}

// GetPosts returns the list of all the posts that are stored into the current state.
func (k Keeper) GetPosts(ctx sdk.Context) (posts types.Posts) {
	posts = types.Posts{}
	k.IteratePosts(ctx, func(_ int64, post types.Post) (stop bool) {
		posts = append(posts, post)
		return false
	})

	return posts
}

// GetPostsFiltered retrieves posts filtered by a given set of params which
// include pagination parameters along with the creator address, the parent id and the creation time.
//
// NOTE: If no filters are provided, all posts will be returned in paginated
// form.
func (k Keeper) GetPostsFiltered(ctx sdk.Context, params types.QueryPostsParams) types.Posts {
	filteredPosts := types.Posts{}
	k.IteratePosts(ctx, func(_ int64, post types.Post) (stop bool) {
		matchParentID, matchCreationTime, matchAllowsComments, matchSubspace, matchCreator := true, true, true, true, true

		// match parent id if valid
		if params.ParentID != nil {
			matchParentID = params.ParentID.Equals(post.ParentID)
		}

		// match creation time if valid height
		if params.CreationTime != nil {
			matchCreationTime = params.CreationTime.Equal(post.Created)
		}

		// match allows comments
		if params.AllowsComments != nil {
			matchAllowsComments = *params.AllowsComments == post.AllowsComments
		}

		// match subspace if provided
		if len(params.Subspace) > 0 {
			matchSubspace = params.Subspace == post.Subspace
		}

		// match creator address (if supplied)
		if len(params.Creator) > 0 {
			matchCreator = params.Creator.Equals(post.Creator)
		}

		if matchParentID && matchCreationTime && matchAllowsComments && matchSubspace && matchCreator {
			filteredPosts = append(filteredPosts, post)
		}

		return false
	})

	// Sort the posts
	sort.Slice(filteredPosts, func(i, j int) bool {
		var result bool
		first, second := filteredPosts[i], filteredPosts[j]

		switch params.SortBy {
		case types.PostSortByCreationDate:
			result = first.Created.Before(second.Created)
			if params.SortOrder == types.PostSortOrderDescending {
				result = first.Created.After(second.Created)
			}

		default:
			result = first.PostID < second.PostID
			if params.SortOrder == types.PostSortOrderDescending {
				result = first.PostID > second.PostID
			}
		}

		// This should never be reached
		return result
	})

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
