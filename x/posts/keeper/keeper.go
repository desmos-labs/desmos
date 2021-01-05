package keeper

import (
	"sort"
	"strings"

	relationshipskeeper "github.com/desmos-labs/desmos/x/relationships/keeper"

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/posts/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.BinaryMarshaler

	paramSubspace paramstypes.Subspace       // Reference to the ParamsStore to get and set posts specific params
	rk            relationshipskeeper.Keeper // Relationships keeper to keep track of blocked users

}

// NewKeeper creates new instances of the posts Keeper
func NewKeeper(
	cdc codec.BinaryMarshaler, storeKey sdk.StoreKey,
	paramSpace paramstypes.Subspace, rk relationshipskeeper.Keeper,
) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:      storeKey,
		cdc:           cdc,
		paramSubspace: paramSpace,
		rk:            rk,
	}
}

// -------------
// --- Posts
// -------------

// IsUserBlocked tells if the given blocker has blocked the given blocked user
func (k Keeper) IsUserBlocked(ctx sdk.Context, blocker, blocked, subspace string) bool {
	return k.rk.HasUserBlocked(ctx, blocker, blocked, subspace)
}

// SavePost allows to save the given post inside the current context.
// It assumes that the given post has already been validated.
// If another post has the same ID of the given post, the old post will be overridden
func (k Keeper) SavePost(ctx sdk.Context, post types.Post) {
	store := ctx.KVStore(k.storeKey)

	// Save the post
	store.Set(types.PostStoreKey(post.PostID), k.cdc.MustMarshalBinaryBare(&post))

	// Check if the postID got an associated post, if not, increment the number of posts
	if !store.Has(types.PostIndexedIDStoreKey(post.PostID)) {
		// Retrieve the total number of posts, if null it will be equal to 0
		postIndex := types.PostIndex{Value: 0}
		if store.Has(types.PostTotalNumberPrefix) {
			k.cdc.MustUnmarshalBinaryBare(store.Get(types.PostTotalNumberPrefix), &postIndex)
		}

		postIndex = types.PostIndex{Value: postIndex.Value + 1}

		// Save the new incremental ID of the post and update the total number of posts
		store.Set(types.PostIndexedIDStoreKey(post.PostID), k.cdc.MustMarshalBinaryBare(&postIndex))
		store.Set(types.PostTotalNumberPrefix, k.cdc.MustMarshalBinaryBare(&postIndex))
	}

	// Save the comments to the parent post, if it is valid
	if types.IsValidPostID(post.ParentID) {
		parentCommentsKey := types.PostCommentsStoreKey(post.ParentID)

		var commentsIDs types.CommentIDs
		k.cdc.MustUnmarshalBinaryBare(store.Get(parentCommentsKey), &commentsIDs)
		if editedIDs, appended := commentsIDs.AppendIfMissing(post.PostID); appended {
			store.Set(parentCommentsKey, k.cdc.MustMarshalBinaryBare(&editedIDs))
		}
	}
}

// DoesPostExist returns true if the post with the given id exists inside the store.
func (k Keeper) DoesPostExist(ctx sdk.Context, id string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.PostStoreKey(id))
}

// GetPost returns the post having the given id inside the current context.
// If no post having the given id can be found inside the current context, false will be returned.
func (k Keeper) GetPost(ctx sdk.Context, id string) (post types.Post, found bool) {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(types.PostStoreKey(id)) {
		return types.Post{}, false
	}

	k.cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(id)), &post)
	return post, true
}

// GetPostChildrenIDs returns the IDs of all the children posts associated to the post
// having the given postID
// nolint: interfacer
func (k Keeper) GetPostChildrenIDs(ctx sdk.Context, postID string) []string {
	store := ctx.KVStore(k.storeKey)

	var ids types.CommentIDs
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.PostCommentsStoreKey(postID)), &ids)
	return ids.Ids
}

// GetPosts returns the list of all the posts that are stored into the current state
//sorted by their incremental ID.
func (k Keeper) GetPosts(ctx sdk.Context) []types.Post {
	var posts []types.Post
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
func (k Keeper) GetPostsFiltered(ctx sdk.Context, params types.QueryPostsParams) []types.Post {
	var filteredPosts []types.Post
	k.IteratePosts(ctx, func(_ int64, post types.Post) (stop bool) {
		matchParentID, matchCreationTime, matchSubspace, matchCreator, matchHashtags := true, true, true, true, true

		// match parent id if valid
		if types.IsValidPostID(params.ParentID) {
			matchParentID = params.ParentID == post.ParentID
		}

		// match creation time if valid height
		if params.CreationTime != nil {
			matchCreationTime = params.CreationTime.Equal(post.Created)
		}

		// match subspace if provided
		if strings.TrimSpace(params.Subspace) != "" {
			matchSubspace = params.Subspace == post.Subspace
		}

		// match creator address (if supplied)
		if strings.TrimSpace(params.Creator) != "" {
			matchCreator = params.Creator == post.Creator
		}

		// match hashtags if provided
		if params.Hashtags != nil {
			postHashtags := post.GetPostHashtags()
			matchHashtags = len(postHashtags) == len(params.Hashtags)
			sort.Strings(postHashtags)
			sort.Strings(params.Hashtags)
			for index := 0; index < len(params.Hashtags) && matchHashtags; index++ {
				matchHashtags = postHashtags[index] == params.Hashtags[index]
			}
		}

		if matchParentID && matchCreationTime && matchSubspace && matchCreator && matchHashtags {
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

	start, end := client.Paginate(len(filteredPosts), int(page), int(params.Limit), 100)
	if start < 0 || end < 0 {
		filteredPosts = []types.Post{}
	} else {
		filteredPosts = filteredPosts[start:end]
	}

	return filteredPosts
}
