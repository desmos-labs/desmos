package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

// SaveReport allows to save the given report inside the current context.
// It assumes that the given report has already been validated.
// If the same report has already been inserted, nothing will be changed.
func (k Keeper) SaveReport(ctx sdk.Context, report types.Report) error {
	store := ctx.KVStore(k.storeKey)
	key := types.ReportStoreKey(report.PostID)

	// Get the list of reports related to the given postID
	reports := types.MustUnmarshalReports(store.Get(key), k.cdc)

	// Append the given report
	newSlice, appended := types.AppendIfMissing(reports, report)
	if !appended {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "report already exists")
	}

	store.Set(key, types.MustMarshalReports(newSlice, k.cdc))
	return nil
}

// GetPostReports returns the list of reports associated with the given postID.
// If no report is associated with the given postID the function will returns an empty list.
func (k Keeper) GetPostReports(ctx sdk.Context, postID string) []types.Report {
	store := ctx.KVStore(k.storeKey)
	return types.MustUnmarshalReports(store.Get(types.ReportStoreKey(postID)), k.cdc)
}

// GetAllReports returns the list of all the reports that have been stored inside the given context
func (k Keeper) GetAllReports(ctx sdk.Context) []types.Report {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ReportsStorePrefix)
	defer iterator.Close()

	var reports []types.Report
	for ; iterator.Valid(); iterator.Next() {
		postReports := types.MustUnmarshalReports(iterator.Value(), k.cdc)
		reports = append(reports, postReports...)
	}

	return reports
}
