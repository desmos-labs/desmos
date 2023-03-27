package keeper

import (
	errors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/reports/types"
)

// SetNextReasonID sets the new reason id for the given subspace to the store
func (k Keeper) SetNextReasonID(ctx sdk.Context, subspaceID uint64, reasonID uint32) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.NextReasonIDStoreKey(subspaceID), types.GetReasonIDBytes(reasonID))
}

// HasNextReasonID tells whether the next reason id exists for the given subspace
func (k Keeper) HasNextReasonID(ctx sdk.Context, subspaceID uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.NextReasonIDStoreKey(subspaceID))
}

// GetNextReasonID gets the highest reason id for the given subspace
func (k Keeper) GetNextReasonID(ctx sdk.Context, subspaceID uint64) (reasonID uint32, err error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NextReasonIDStoreKey(subspaceID))
	if bz == nil {
		return 0, errors.Wrapf(types.ErrInvalidGenesis, "initial reason id not set for subspace %d", subspaceID)
	}

	reasonID = types.GetReasonIDFromBytes(bz)
	return reasonID, nil
}

// DeleteNextReasonID removes the reason id key for the given subspace
func (k Keeper) DeleteNextReasonID(ctx sdk.Context, subspaceID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.NextReasonIDStoreKey(subspaceID))
}

// --------------------------------------------------------------------------------------------------------------------

// SaveReason saves the given reason inside the current context
func (k Keeper) SaveReason(ctx sdk.Context, reason types.Reason) {
	store := ctx.KVStore(k.storeKey)

	// Store the reason
	store.Set(types.ReasonStoreKey(reason.SubspaceID, reason.ID), k.cdc.MustMarshal(&reason))

	k.Logger(ctx).Debug("reason saved", "subspace id", reason.SubspaceID, "id", reason.ID)
	k.AfterReasonSaved(ctx, reason.SubspaceID, reason.ID)
}

// HasReason tells whether the given reason exists or not
func (k Keeper) HasReason(ctx sdk.Context, subspaceID uint64, reasonID uint32) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.ReasonStoreKey(subspaceID, reasonID))
}

// GetReason returns the reason associated with the given id.
// If there is no reason with the given id the function will return an empty reason and false.
func (k Keeper) GetReason(ctx sdk.Context, subspaceID uint64, reasonID uint32) (reason types.Reason, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ReasonStoreKey(subspaceID, reasonID))
	if bz == nil {
		return types.Reason{}, false
	}

	k.cdc.MustUnmarshal(bz, &reason)
	return reason, true
}

// DeleteReason deletes the reason having the given id from the store
func (k Keeper) DeleteReason(ctx sdk.Context, subspaceID uint64, reasonID uint32) {
	reason, found := k.GetReason(ctx, subspaceID, reasonID)
	if !found {
		return
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete(types.ReasonStoreKey(subspaceID, reasonID))

	// Delete all the reports associated to this reason
	k.IterateSubspaceReports(ctx, subspaceID, func(report types.Report) (stop bool) {
		if types.ContainsReason(report.ReasonsIDs, reasonID) {
			k.DeleteReport(ctx, report.SubspaceID, report.ID)
		}
		return false
	})

	k.Logger(ctx).Debug("reason deleted", "subspace id", reason.SubspaceID, "id", reason.ID)
	k.AfterReasonDeleted(ctx, reason.SubspaceID, reason.ID)
}
