package keeper

import (
	errors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/reactions/types"
)

// SaveSubspaceReactionsParams stores the given reactions params inside the store
func (k Keeper) SaveSubspaceReactionsParams(ctx sdk.Context, params types.SubspaceReactionsParams) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.SubspaceReactionsParamsStoreKey(params.SubspaceID), k.cdc.MustMarshal(&params))
}

// HasSubspaceReactionsParams tells whether the params for the given subspace exist or not
func (k Keeper) HasSubspaceReactionsParams(ctx sdk.Context, subspaceID uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.SubspaceReactionsParamsStoreKey(subspaceID))
}

// GetSubspaceReactionsParams returns the reactions params for the given subspace
func (k Keeper) GetSubspaceReactionsParams(ctx sdk.Context, subspaceID uint64) (params types.SubspaceReactionsParams, err error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.SubspaceReactionsParamsStoreKey(subspaceID))
	if bz == nil {
		return types.SubspaceReactionsParams{}, errors.Wrapf(types.ErrInvalidGenesis, "reactions params are not set for subspace %d", subspaceID)
	}

	k.cdc.MustUnmarshal(bz, &params)
	return params, nil
}

// DeleteSubspaceReactionsParams deletes the reactions params for the given subspace
func (k Keeper) DeleteSubspaceReactionsParams(ctx sdk.Context, subspaceID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.SubspaceReactionsParamsStoreKey(subspaceID))
}
