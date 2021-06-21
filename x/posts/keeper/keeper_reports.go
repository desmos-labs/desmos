package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	types2 "github.com/desmos-labs/desmos/x/posts/types"
)

// SaveReport allows to save the given report inside the current context.
// It assumes that the given report has already been validated.
// If the same report has already been inserted, nothing will be changed.
func (k Keeper) SaveReport(ctx sdk.Context, report types2.Report) error {
	store := ctx.KVStore(k.storeKey)
	key := types2.ReportStoreKey(report.PostID)

	// Get the list of reports related to the given postID
	reports := types2.MustUnmarshalReports(store.Get(key), k.cdc)

	// Append the given report
	newSlice, appended := types2.AppendIfMissing(reports, report)
	if !appended {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "report already exists")
	}

	store.Set(key, types2.MustMarshalReports(newSlice, k.cdc))

	k.Logger(ctx).Info("reported post", "post-id", report.PostID, "from", report.User)
	return nil
}

// GetPostReports returns the list of reports associated with the given postID.
// If no report is associated with the given postID the function will returns an empty list.
func (k Keeper) GetPostReports(ctx sdk.Context, postID string) []types2.Report {
	store := ctx.KVStore(k.storeKey)
	return types2.MustUnmarshalReports(store.Get(types2.ReportStoreKey(postID)), k.cdc)
}

// GetAllReports returns the list of all the reports that have been stored inside the given context
func (k Keeper) GetAllReports(ctx sdk.Context) []types2.Report {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types2.ReportsStorePrefix)
	defer iterator.Close()

	var reports []types2.Report
	for ; iterator.Valid(); iterator.Next() {
		postReports := types2.MustUnmarshalReports(iterator.Value(), k.cdc)
		reports = append(reports, postReports...)
	}

	return reports
}
