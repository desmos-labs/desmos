package keeper

import (
	"bytes"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts"
	"github.com/desmos-labs/desmos/x/reports/internal/types"
	"github.com/desmos-labs/desmos/x/reports/internal/types/models"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	postKeeper posts.Keeper // post's keeper to perform checks on the postIDs
	StoreKey   sdk.StoreKey // Unexposed key to access store from sdk.Context
	Cdc        *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the reports Keeper
func NewKeeper(pk posts.Keeper, cdc *codec.Codec, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		postKeeper: pk,
		StoreKey:   storeKey,
		Cdc:        cdc,
	}
}

// CheckPostExistence checks if a post with the given postID is present inside
// the current context and returns a boolean indicating that.
func (k Keeper) CheckPostExistence(ctx sdk.Context, postID posts.PostID) bool {
	_, exist := k.postKeeper.GetPost(ctx, postID)
	return exist
}

// SaveReport allows to save the given reports inside the current context.
// It assumes that the given reports has already been validated.
// If the same reports has already been inserted, nothing will be changed.
func (k Keeper) SaveReport(ctx sdk.Context, postID posts.PostID, report types.Report) {
	store := ctx.KVStore(k.StoreKey)
	key := models.ReportStoreKey(postID)
	// Get the list of reports related to the given postID
	var reports models.Reports
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &reports)

	// try to append the given reports
	reports = append(reports, report)
	store.Set(key, k.Cdc.MustMarshalBinaryBare(&reports))

}

// GetPostReports returns the list of reports associated with the given postID.
// If no reports is associated with the given postID the function will returns an empty list.
func (k Keeper) GetPostReports(ctx sdk.Context, postID posts.PostID) (reports types.Reports) {
	store := ctx.KVStore(k.StoreKey)

	// Get the list of reports related to the given postID
	k.Cdc.MustUnmarshalBinaryBare(store.Get(models.ReportStoreKey(postID)), &reports)
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
