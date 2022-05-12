package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v3/x/reports/types"
)

// SetNextReportID sets the new report id for the given subspace to the store
func (k Keeper) SetNextReportID(ctx sdk.Context, subspaceID uint64, reportID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.NextReportIDStoreKey(subspaceID), types.GetReportIDBytes(reportID))
}

// GetNextReportID gets the highest report id for the given subspace
func (k Keeper) GetNextReportID(ctx sdk.Context, subspaceID uint64) (reportID uint64, err error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NextReportIDStoreKey(subspaceID))
	if bz == nil {
		return 0, sdkerrors.Wrapf(types.ErrInvalidGenesis, "initial report id hasn't been set for subspace %d", subspaceID)
	}

	reportID = types.GetReportIDFromBytes(bz)
	return reportID, nil
}

// DeleteNextReportID removes the report id key for the given subspace
func (k Keeper) DeleteNextReportID(ctx sdk.Context, subspaceID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.NextReportIDStoreKey(subspaceID))
}

// --------------------------------------------------------------------------------------------------------------------
func (k Keeper) getContentKey(report types.Report) []byte {
	var contentKey []byte
	switch data := report.Data.GetCachedValue().(type) {
	case *types.UserData:
		userAddress, err := sdk.AccAddressFromBech32(data.User)
		if err != nil {
			panic(fmt.Errorf("invalid reported user: %s", err))
		}
		contentKey = types.UserReportStoreKey(report.SubspaceID, userAddress, report.ID)

	case *types.PostData:
		contentKey = types.PostReportStoreKey(report.SubspaceID, data.PostID, report.ID)
	}

	if contentKey == nil {
		panic(fmt.Errorf("unsupported content type: %T", report.Data.GetCachedValue()))
	}

	return contentKey
}

// SaveReport saves the given report inside the current context
func (k Keeper) SaveReport(ctx sdk.Context, report types.Report) {
	store := ctx.KVStore(k.storeKey)

	// Store the report
	store.Set(types.ReportStoreKey(report.SubspaceID, report.ID), k.cdc.MustMarshal(&report))

	// Set the reference for the content
	store.Set(k.getContentKey(report), types.GetReportIDBytes(report.ID))

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
	report, found := k.GetReport(ctx, subspaceID, reportID)
	if !found {
		return
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete(types.ReportStoreKey(report.SubspaceID, report.ID))
	store.Delete(k.getContentKey(report))
}
