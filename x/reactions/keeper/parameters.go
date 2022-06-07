package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/reactions/types"
)

// SaveSubspaceParams stores the given reactions params inside the store
func (k Keeper) SaveSubspaceParams(ctx sdk.Context, params types.SubspaceReactionsParams) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ReactionsParamsStoreKey(params.SubspaceID), k.cdc.MustMarshal(&params))
}

// HasSubspaceParams tells whether the params for the given subspace exist or not
func (k Keeper) HasSubspaceParams(ctx sdk.Context, subspaceID uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.ReactionsParamsStoreKey(subspaceID))
}

// GetSubspaceReactionsParams returns the reactions params for the given subspace.
// If the params are not set, returns types.DefaultReactionsParams instead
func (k Keeper) GetSubspaceReactionsParams(ctx sdk.Context, subspaceID uint64) types.SubspaceReactionsParams {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ReactionsParamsStoreKey(subspaceID))
	if bz == nil {
		return types.DefaultReactionsParams(subspaceID)
	}

	var params types.SubspaceReactionsParams
	k.cdc.MustUnmarshal(bz, &params)
	return params
}
