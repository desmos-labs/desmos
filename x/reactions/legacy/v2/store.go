package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	poststypes "github.com/desmos-labs/desmos/v3/x/posts/types"
	"github.com/desmos-labs/desmos/v3/x/reactions/types"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

type SubspacesKeeper interface {
	IterateSubspaces(ctx sdk.Context, fn func(subspaces subspacestypes.Subspace) (stop bool))
}

type PostsKeeper interface {
	IteratePosts(ctx sdk.Context, fn func(post poststypes.Post) (stop bool))
}

// MigrateStore performs in-place store migrations from v1 to v2
// The only thing that is done here is setting up the next
// registered reaction id for existing subspaces and the next
// reaction id for existing posts
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, sk SubspacesKeeper, pk PostsKeeper) error {
	store := ctx.KVStore(storeKey)

	// Set the next registered reaction id for all the subspaces
	sk.IterateSubspaces(ctx, func(subspace subspacestypes.Subspace) (stop bool) {
		store.Set(types.NextRegisteredReactionIDStoreKey(subspace.ID), types.GetRegisteredReactionIDBytes(1))
		return false
	})

	// Set the next reaction id for all the posts
	pk.IteratePosts(ctx, func(post poststypes.Post) (stop bool) {
		store.Set(types.NextReactionIDStoreKey(post.SubspaceID, post.ID), types.GetReactionIDBytes(1))
		return false
	})

	return nil
}
