package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	postskeeper "github.com/desmos-labs/desmos/x/posts/keeper"
	"github.com/desmos-labs/desmos/x/reports/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey          // Unexposed key to access store from sdk.Context
	cdc      codec.BinaryMarshaler // The wire codec for binary encoding/decoding.

	postKeeper postskeeper.Keeper // Post's keeper to perform checks on the postIDs
}

// NewKeeper creates new instances of the stored Keeper
func NewKeeper(cdc codec.BinaryMarshaler, storeKey sdk.StoreKey, pk postskeeper.Keeper) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		postKeeper: pk,
	}
}

// CheckPostExistence checks if a post with the given postID is present inside
// the current context and returns a boolean indicating that.
func (k Keeper) CheckPostExistence(ctx sdk.Context, postID string) bool {
	_, exist := k.postKeeper.GetPost(ctx, postID)
	return exist
}

// SaveReport allows to save the given stored inside the current context.
// It assumes that the given stored has already been validated.
// If the same stored has already been inserted, nothing will be changed.
func (k Keeper) SaveReport(ctx sdk.Context, report types.Report) error {
	store := ctx.KVStore(k.storeKey)
	key := types.ReportStoreKey(report.PostId)

	// Get the list of stored related to the given postID
	reports, err := k.UnmarshalReports(store.Get(key))
	if err != nil {
		return err
	}

	// Append the given stored
	reports = append(reports, report)

	bz, err := k.MarshalReports(reports)
	if err != nil {
		return err
	}

	store.Set(key, bz)
	return nil
}

// GetPostReports returns the list of stored associated with the given postID.
// If no stored is associated with the given postID the function will returns an empty list.
func (k Keeper) GetPostReports(ctx sdk.Context, postID string) ([]types.Report, error) {
	store := ctx.KVStore(k.storeKey)
	return k.UnmarshalReports(store.Get(types.ReportStoreKey(postID)))
}

// GetReports returns the list of all the stored that have been stored inside the given context
func (k Keeper) GetReports(ctx sdk.Context) ([]types.Report, error) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ReportsStorePrefix)
	defer iterator.Close()

	var reports []types.Report
	for ; iterator.Valid(); iterator.Next() {
		postReports, err := k.UnmarshalReports(iterator.Value())
		if err != nil {
			return nil, err
		}

		reports = append(reports, postReports...)
	}

	return reports, nil
}

// MarshalReports marshals a list of Report. If the given type implements
// the Marshaler interface, it is treated as a Proto-defined message and
// serialized that way. Otherwise, it falls back on the internal Amino codec.
func (k Keeper) MarshalReports(reports []types.Report) ([]byte, error) {
	return codec.MarshalAny(k.cdc, reports)
}

// UnmarshalReports returns a list of Report interface from raw encoded
// blocks bytes. An error is returned upon decoding failure.
func (k Keeper) UnmarshalReports(bz []byte) ([]types.Report, error) {
	var reports []types.Report
	if err := codec.UnmarshalAny(k.cdc, &reports, bz); err != nil {
		return nil, err
	}

	return reports, nil
}
