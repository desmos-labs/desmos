package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

type SubspacesKeeper interface {
	IterateSubspaces(ctx sdk.Context, fn func(index int64, subspaces subspacestypes.Subspace) (stop bool))
}

// MigrateStore performs in-place store migrations from v1 to v2
// The only thing that is done here is setting up the next post id key for existing subspaces.
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, sk SubspacesKeeper) error {
	store := ctx.KVStore(storeKey)

	// Set the next post id for all the subspaces
	sk.IterateSubspaces(ctx, func(index int64, subspaces subspacestypes.Subspace) (stop bool) {
		store.Set(types.NextPostIDStoreKey(subspaces.ID), types.GetPostIDBytes(1))
		return false
	})

	return nil
}
