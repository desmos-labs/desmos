package keeper

import (
	"bytes"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts"
	"github.com/desmos-labs/desmos/x/report/internal/types"
	"github.com/desmos-labs/desmos/x/report/internal/types/models"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	StoreKey sdk.StoreKey // Unexposed key to access store from sdk.Context
	Cdc      *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the reports Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		StoreKey: storeKey,
		Cdc:      cdc,
	}
}

// SaveReport allows to save the given report inside the current context.
// It assumes that the given report has already been validated.
// If the same report has already been inserted, nothing will be changed.
func (k Keeper) SaveReport(ctx sdk.Context, postID posts.PostID, report types.Report) bool {
	store := ctx.KVStore(k.StoreKey)
	key := models.ReportStoreKey(postID)
	// Get the list of reports related to the given postID
	var reports models.Reports
	bz := store.Get(key)
	k.Cdc.MustUnmarshalBinaryBare(bz, &reports)

	// try to append the given report
	reports, appended := reports.AppendIfMissing(report)
	if appended {
		store.Set(key, k.Cdc.MustMarshalBinaryBare(&reports))
	}

	return appended
}

// GetPostReports returns the list of reports associated with the given postID.
// If no report is associated with the given postID the function will returns an empty list.
func (k Keeper) GetPostReports(ctx sdk.Context, postID posts.PostID) (reports types.Reports) {
	store := ctx.KVStore(k.StoreKey)

	// Get the list of reports related to the given postID
	bz := store.Get(models.ReportStoreKey(postID))
	k.Cdc.MustUnmarshalBinaryBare(bz, reports)

	return reports
}

// GetReportsMap allows to returns the list of reports that have been stored inside the given context
func (k Keeper) GetReportsMap(ctx sdk.Context) map[string]types.Reports {
	store := ctx.KVStore(k.StoreKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ReportsStorePrefix)
	defer iterator.Close()

	reportsData := map[string]types.Reports{}
	for ; iterator.Valid(); iterator.Next() {
		var reports types.Reports
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &reports)
		idBytes := bytes.TrimPrefix(iterator.Key(), types.ReportsStorePrefix)
		reportsData[string(idBytes)] = reports
	}

	return reportsData
}
