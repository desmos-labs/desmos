package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/reports/types"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

type SubspacesKeeper interface {
	IterateSubspaces(ctx sdk.Context, fn func(subspaces subspacestypes.Subspace) (stop bool))
}

// MigrateStore performs in-place store migrations from v1 to v2
// The only thing that is done here is setting up the next reason id and report id keys for existing subspaces.
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, sk SubspacesKeeper) error {
	store := ctx.KVStore(storeKey)

	// Set the next reason id and report id for all the subspaces
	sk.IterateSubspaces(ctx, func(subspace subspacestypes.Subspace) (stop bool) {
		store.Set(types.NextReasonIDStoreKey(subspace.ID), types.GetReasonIDBytes(1))
		store.Set(types.NextReportIDStoreKey(subspace.ID), types.GetReportIDBytes(1))
		return false
	})

	return nil
}
