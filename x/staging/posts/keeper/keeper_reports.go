package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

// CheckReportValidity ensure that the given report is valid
// It returns error if not
func (k Keeper) CheckReportValidity(ctx sdk.Context, report types.Report) error {
	reportReasons := k.GetParams(ctx).AllowedReasons

	if !report.AreReasonsValid(reportReasons) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid report reason")
	}

	return report.Validate()
}

// SaveReport allows to save the given report inside the current context.
// It assumes that the given report has already been validated.
// If the same report has already been inserted, nothing will be changed.
func (k Keeper) SaveReport(ctx sdk.Context, report types.Report) error {
	store := ctx.KVStore(k.storeKey)
	key := types.ReportStoreKey(report.PostID, report.User)

	// Check if the report already exist
	if store.Has(key) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%s already reported post with id %s",
			report.User, report.PostID)
	}

	if err := k.CheckReportValidity(ctx, report); err != nil {
		return err
	}

	store.Set(key, types.MustMarshalReport(k.cdc, report))

	k.Logger(ctx).Info("reported post", "post-id", report.PostID, "from", report.User)
	return nil
}

// GetPostReports returns the list of reports associated with the given postID.
// If no report is associated with the given postID the function will returns an empty list.
func (k Keeper) GetPostReports(ctx sdk.Context, postID string) []types.Report {
	var reports []types.Report
	k.IteratePostReportsByPost(ctx, postID, func(_ int64, report types.Report) bool {
		reports = append(reports, report)
		return false
	})
	return reports
}

// GetAllReports returns the list of all the reports that have been stored inside the given context
func (k Keeper) GetAllReports(ctx sdk.Context) []types.Report {
	var reports []types.Report
	k.IterateReports(ctx, func(_ int64, report types.Report) bool {
		reports = append(reports, report)
		return false
	})
	return reports
}
