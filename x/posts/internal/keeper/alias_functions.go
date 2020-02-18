package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
)

// IteratePosts iterates through the posts set and performs the provided function
func (k Keeper) IteratePosts(ctx sdk.Context, fn func(index int64, post types.Post) (stop bool)) {
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PostStorePrefix)
	defer iterator.Close()
	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var post types.Post
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &post)
		stop := fn(i, post)
		if stop {
			break
		}
		i++
	}
}

// IsPostDuplicate returns true if another post having the same ID or content of the
// specified post already exists inside the store.
// If it exists, it is returns the reference to such post and true, otherwise it returns nil and false.
func (k Keeper) IsPostDuplicate(ctx sdk.Context, post types.Post) (*types.Post, bool) {
	var existing *types.Post

	k.IteratePosts(ctx, func(_ int64, value types.Post) (stop bool) {
		if post.IsDuplicate(value) {
			existing = &value
			return true
		}

		return false
	})

	return existing, existing != nil
}
