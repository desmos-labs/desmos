package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v3/x/reports/types"
)

// SetReportID sets the new report id for the given subspace to the store
func (k Keeper) SetReportID(ctx sdk.Context, subspaceID uint64, reportID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ReportIDStoreKey(subspaceID), types.GetReportIDBytes(reportID))
}

// GetReportID gets the highest report id for the given subspace
func (k Keeper) GetReportID(ctx sdk.Context, subspaceID uint64) (reportID uint64, err error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ReportIDStoreKey(subspaceID))
	if bz == nil {
		return 0, sdkerrors.Wrapf(types.ErrInvalidGenesis, "initial report id hasn't been set for subspace %d", subspaceID)
	}

	reportID = types.GetReportIDFromBytes(bz)
	return reportID, nil
}

// DeleteReportID removes the report id key for the given subspace
func (k Keeper) DeleteReportID(ctx sdk.Context, subspaceID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.ReportIDStoreKey(subspaceID))
}

// --------------------------------------------------------------------------------------------------------------------

// SaveReport saves the given report inside the current context
func (k Keeper) SaveReport(ctx sdk.Context, report types.Report) {
	store := ctx.KVStore(k.storeKey)

	// Store the report
	store.Set(types.ReportStoreKey(report.SubspaceID, report.ID), k.cdc.MustMarshal(&report))

	k.Logger(ctx).Debug("report saved", "subspace id", report.SubspaceID, "id", report.ID)
	k.AfterReportSaved(ctx, report.SubspaceID, report.ID)
}

// HasReport tells whether the given report exists or not
func (k Keeper) HasReport(ctx sdk.Context, subspaceID uint64, reportID uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.ReportStoreKey(subspaceID, reportID))
}

// GetReport returns the report associated with the given id.
// If there is no report associated with the given id the function will return an empty report and false.
func (k Keeper) GetReport(ctx sdk.Context, subspaceID uint64, reportID uint64) (report types.Report, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ReportStoreKey(subspaceID, reportID))
	if bz == nil {
		return types.Report{}, false
	}

	k.cdc.MustUnmarshal(bz, &report)
	return report, true
}

// DeleteReport deletes the report having the given id from the store
func (k Keeper) DeleteReport(ctx sdk.Context, subspaceID uint64, reportID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.ReportStoreKey(subspaceID, reportID))
}
