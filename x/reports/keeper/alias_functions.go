package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/reports/types"

	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// HasSubspace tells whether the subspace with the given id exists or not
func (k Keeper) HasSubspace(ctx sdk.Context, subspaceID uint64) bool {
	return k.sk.HasSubspace(ctx, subspaceID)
}

// HasPermission tells whether the given user has the provided permission inside the subspace with the specified id
func (k Keeper) HasPermission(ctx sdk.Context, subspaceID uint64, user sdk.AccAddress, permission subspacestypes.Permission) bool {
	return k.sk.HasPermission(ctx, subspaceID, user, permission)
}

// HasUserBlocked tells whether the given blocker has blocked the user inside the provided subspace
func (k Keeper) HasUserBlocked(ctx sdk.Context, blocker, user string, subspaceID uint64) bool {
	return k.rk.HasUserBlocked(ctx, blocker, user, subspaceID)
}

// --------------------------------------------------------------------------------------------------------------------

// IterateSubspaceReasons iterates over all the given subspace reasons and performs the provided function
func (k Keeper) IterateSubspaceReasons(ctx sdk.Context, subspaceID uint64, fn func(index int64, reason types.Reason) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SubspaceReasonsPrefix(subspaceID))
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var reason types.Reason
		k.cdc.MustUnmarshal(iterator.Value(), &reason)
		stop := fn(i, reason)
		if stop {
			break
		}
		i++
	}
}

// IterateSubspaceReports iterates over all the given subspace reports and performs the provided function
func (k Keeper) IterateSubspaceReports(ctx sdk.Context, subspaceID uint64, fn func(index int64, report types.Report) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SubspaceReportsPrefix(subspaceID))
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var report types.Report
		k.cdc.MustUnmarshal(iterator.Value(), &report)
		stop := fn(i, report)
		if stop {
			break
		}
		i++
	}
}
