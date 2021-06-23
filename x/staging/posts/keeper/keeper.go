package keeper

import (
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.BinaryMarshaler

	paramSubspace paramstypes.Subspace // Reference to the ParamsStore to get and set posts specific params
	rk            RelationshipsKeeper  // Relationships k to keep track of blocked users
	sk            SubspacesKeeper      // Subspaces k to make checks on posts based on their subspace
}

// NewKeeper creates new instances of the posts Keeper
func NewKeeper(
	cdc codec.BinaryMarshaler, storeKey sdk.StoreKey,
	paramSpace paramstypes.Subspace, rk RelationshipsKeeper, sk SubspacesKeeper,
) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:      storeKey,
		cdc:           cdc,
		paramSubspace: paramSpace,
		rk:            rk,
		sk:            sk,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
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

	// Save the query key if the key does not exist
	subspaceKey := types.SubspacePostKey(post.Subspace, post.PostID)
	if !store.Has(subspaceKey) {
		store.Set(subspaceKey, []byte(post.PostID))
	}

	// Save the comments to the parent post, if it is valid
	if types.IsValidPostID(post.ParentID) {
		commentKey := types.CommentsStoreKey(post.ParentID, post.PostID)
		if !store.Has(commentKey) {
			store.Set(commentKey, []byte(post.PostID))
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

// GetPostCommentsIDs returns the IDs of all the comments associated to the post
// having the given postID
func (k Keeper) GetPostCommentsIDs(ctx sdk.Context, postID string) []string {
	ids := []string{}

	k.IterateCommentIDsByPost(ctx, postID, func(_ int64, commentID string) bool {
		ids = append(ids, commentID)
		return false
	})

	return ids
}

// GetPosts returns the list of all the posts that are stored into the current state
func (k Keeper) GetPosts(ctx sdk.Context) []types.Post {
	var posts []types.Post
	k.IteratePosts(ctx, func(_ int64, post types.Post) (stop bool) {
		posts = append(posts, post)
		return false
	})

	return posts
}

// -------------
// --- Subspaces
// -------------

func (k Keeper) CheckUserPermissionOnSubspace(ctx sdk.Context, subspaceID string, user string) error {
	return k.sk.CheckSubspaceUserPermission(ctx, subspaceID, user)
}
