package keeper

import (
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

func (k Keeper) SaveReport(ctx sdk.Context, postID posts.PostID, report types.Report) {
	store := ctx.KVStore(k.StoreKey)

	// Save the report
	var reports models.Reports
	store.Get(types.ReportStoreKey(postID))

}

func (k Keeper) GetPostReports(ctx sdk.Context, postID posts.PostID) []types.Report {}
