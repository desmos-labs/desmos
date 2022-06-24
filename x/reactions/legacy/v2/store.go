package v2

import (
	"github.com/cosmos/cosmos-sdk/codec"
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
// The things done here are the following:
// 1. setting next registered reaction id
// 2. setting next reaction id for existing posts
// 3. setting the reaction params for existing subspaces
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, sk SubspacesKeeper, pk PostsKeeper, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	// Set the next registered reaction id for all the subspaces
	sk.IterateSubspaces(ctx, func(subspace subspacestypes.Subspace) (stop bool) {
		store.Set(types.NextRegisteredReactionIDStoreKey(subspace.ID), types.GetRegisteredReactionIDBytes(1))

		params := types.DefaultReactionsParams(subspace.ID)
		store.Set(types.SubspaceReactionsParamsStoreKey(subspace.ID), cdc.MustMarshal(&params))
		return false
	})

	// Set the next reaction id for all the posts
	pk.IteratePosts(ctx, func(post poststypes.Post) (stop bool) {
		store.Set(types.NextReactionIDStoreKey(post.SubspaceID, post.ID), types.GetReactionIDBytes(1))
		return false
	})

	return nil
}
