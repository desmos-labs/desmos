package keeper

import (
	"fmt"

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

// GetStandardReason returns the standard reason with the given id.
// If no standard reason with the given id could be found, the method will return an empty standard reason and false
func (k Keeper) GetStandardReason(ctx sdk.Context, id uint32) (reason types.StandardReason, found bool) {
	for _, reason := range k.GetParams(ctx).StandardReasons {
		if reason.ID == id {
			return reason, true
		}
	}
	return types.StandardReason{}, false
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

// --------------------------------------------------------------------------------------------------------------------

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

// IteratePostReports iterates over all the reports for the given post and performs the provided function
func (k Keeper) IteratePostReports(ctx sdk.Context, subspaceID uint64, postID uint64, fn func(index int64, report types.Report) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PostReportsPrefix(subspaceID, postID))
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		reportID := types.GetReportIDFromBytes(iterator.Value())
		report, found := k.GetReport(ctx, subspaceID, reportID)
		if !found {
			panic(fmt.Errorf("report not found: subspace id %d, report id %d", subspaceID, reportID))
		}

		stop := fn(i, report)
		if stop {
			break
		}
		i++
	}
}

// IterateUserReports iterates over all the reports for the given user and performs the provided function
func (k Keeper) IterateUserReports(ctx sdk.Context, subspaceID uint64, user string, fn func(index int64, report types.Report) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.UserReportsPrefix(subspaceID, user))
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		reportID := types.GetReportIDFromBytes(iterator.Value())
		report, found := k.GetReport(ctx, subspaceID, reportID)
		if !found {
			panic(fmt.Errorf("report not found: subspace id %d, report id %d", subspaceID, reportID))
		}

		stop := fn(i, report)
		if stop {
			break
		}
		i++
	}
}
