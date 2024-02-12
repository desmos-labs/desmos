package keeper

import (
	"fmt"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v7/x/reports/types"
)

// SetNextReportID sets the new report id for the given subspace to the store
func (k Keeper) SetNextReportID(ctx sdk.Context, subspaceID uint64, reportID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.NextReportIDStoreKey(subspaceID), types.GetReportIDBytes(reportID))
}

// HasNextReportID tells whether a next report id exists for the given subspace
func (k Keeper) HasNextReportID(ctx sdk.Context, subspaceID uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.NextReportIDStoreKey(subspaceID))
}

// GetNextReportID gets the highest report id for the given subspace
func (k Keeper) GetNextReportID(ctx sdk.Context, subspaceID uint64) (reportID uint64, err error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NextReportIDStoreKey(subspaceID))
	if bz == nil {
		return 0, errors.Wrapf(types.ErrInvalidGenesis, "initial report id hasn't been set for subspace %d", subspaceID)
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

// validateUserReportContent validates the given target data to make sure the reported user has not blocked the reporter
func (k Keeper) validateUserReportContent(ctx sdk.Context, report types.Report, data *types.UserTarget) error {
	if k.HasUserBlocked(ctx, data.User, report.Reporter, report.SubspaceID) {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "the reported user has blocked you on subspace %d", report.SubspaceID)
	}

	return nil
}

// validatePostReportContent validates the given post report making sure that:
// - the post exists inside the given subspace
// - the post author has not blocked the reporter
func (k Keeper) validatePostReportContent(ctx sdk.Context, report types.Report, data *types.PostTarget) error {
	post, found := k.GetPost(ctx, report.SubspaceID, data.PostID)
	if !found {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "post %d does not exist inside subspace %d", data.PostID, report.SubspaceID)
	}

	if k.HasUserBlocked(ctx, post.Owner, report.Reporter, report.ID) {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "post author has blocked you on this subspace")
	}

	return nil
}

// ValidateReport validates the given report's content
func (k Keeper) ValidateReport(ctx sdk.Context, report types.Report) (err error) {
	// Validate the report
	err = report.Validate()
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Validate the content
	switch data := report.Target.GetCachedValue().(type) {
	case *types.UserTarget:
		err = k.validateUserReportContent(ctx, report, data)
	case *types.PostTarget:
		err = k.validatePostReportContent(ctx, report, data)
	}

	return err
}

// getContentKey returns the store key used to save the report reference based on its content type
func (k Keeper) getContentKey(subspaceID uint64, target types.ReportTarget, reporter string) []byte {
	var contentKey []byte
	switch data := target.(type) {
	case *types.UserTarget:
		contentKey = types.UserReportStoreKey(subspaceID, data.User, reporter)

	case *types.PostTarget:
		contentKey = types.PostReportStoreKey(subspaceID, data.PostID, reporter)
	}

	if contentKey == nil {
		panic(fmt.Errorf("unsupported content type: %T", target))
	}

	return contentKey
}

// SaveReport saves the given report inside the current context
func (k Keeper) SaveReport(ctx sdk.Context, report types.Report) {
	store := ctx.KVStore(k.storeKey)

	// Store the report
	store.Set(types.ReportStoreKey(report.SubspaceID, report.ID), k.cdc.MustMarshal(&report))

	// Set the reference for the content
	contentKey := k.getContentKey(report.SubspaceID, report.Target.GetCachedValue().(types.ReportTarget), report.Reporter)
	store.Set(contentKey, types.GetReportIDBytes(report.ID))

	k.Logger(ctx).Debug("report saved", "subspace id", report.SubspaceID, "id", report.ID)
	k.AfterReportSaved(ctx, report.SubspaceID, report.ID)
}

// HasReport tells whether the given report exists or not
func (k Keeper) HasReport(ctx sdk.Context, subspaceID uint64, reportID uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.ReportStoreKey(subspaceID, reportID))
}

// HasReported tells whether the given reporter has reported the specified target or not
func (k Keeper) HasReported(ctx sdk.Context, subspaceID uint64, reporter string, target types.ReportTarget) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(k.getContentKey(subspaceID, target, reporter))
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

	// Delete the report store key
	store.Delete(types.ReportStoreKey(report.SubspaceID, report.ID))

	// Delete the content key
	contentKey := k.getContentKey(report.SubspaceID, report.Target.GetCachedValue().(types.ReportTarget), report.Reporter)
	store.Delete(contentKey)
}
