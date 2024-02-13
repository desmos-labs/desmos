package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v7/x/relationships/types"
)

// SaveUserBlock allows to store the given block inside the store
func (k Keeper) SaveUserBlock(ctx sdk.Context, userBlock types.UserBlock) {
	store := ctx.KVStore(k.storeKey)
	key := types.UserBlockStoreKey(userBlock.Blocker, userBlock.Blocked, userBlock.SubspaceID)
	store.Set(key, k.cdc.MustMarshal(&userBlock))
}

// HasUserBlocked returns true if the provided blocker has blocked the given user for the given subspace.
func (k Keeper) HasUserBlocked(ctx sdk.Context, blocker, blocked string, subspaceID uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.UserBlockStoreKey(blocker, blocked, subspaceID))
}

// GetUserBlock returns the user block that has been created from the given blocker towards the provided
// blocked address for the specified subspace.
func (k Keeper) GetUserBlock(ctx sdk.Context, blocker, blocked string, subspaceID uint64) (userBlock types.UserBlock, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.UserBlockStoreKey(blocker, blocked, subspaceID)
	if !store.Has(key) {
		return types.UserBlock{}, false
	}

	return types.MustUnmarshalUserBlock(k.cdc, store.Get(key)), true
}

// DeleteUserBlock allows to the specified blocker to unblock the given blocked user
func (k Keeper) DeleteUserBlock(ctx sdk.Context, blocker, blocked string, subspaceID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.UserBlockStoreKey(blocker, blocked, subspaceID))
}
