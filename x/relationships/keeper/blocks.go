package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v2/x/relationships/types"
)

// SaveUserBlock allows to store the given block inside the store
func (k Keeper) SaveUserBlock(ctx sdk.Context, userBlock types.UserBlock) {
	store := ctx.KVStore(k.storeKey)
	key := types.UserBlockStoreKey(userBlock.Blocker, userBlock.SubspaceID, userBlock.Blocked)
	store.Set(key, k.cdc.MustMarshal(&userBlock))
}

// IsUserBlocked returns true if the provided blocker has blocked the given user for the given subspace.
// If the provided subspace is empty, all subspaces will be checked
func (k Keeper) IsUserBlocked(ctx sdk.Context, blocker, blocked string, subspace uint64) bool {
	if subspace != 0 {
		store := ctx.KVStore(k.storeKey)
		key := types.UserBlockStoreKey(blocker, subspace, blocked)

		return store.Has(key)
	}

	blocks := k.GetUserBlocks(ctx, blocker)
	for _, block := range blocks {
		if block.Blocked == blocked {
			return subspace == 0 || block.SubspaceID == subspace
		}
	}

	return false
}

// GetUserBlocks returns the list of users that the specified user has blocked.
func (k Keeper) GetUserBlocks(ctx sdk.Context, blocker string) []types.UserBlock {
	var userblocks []types.UserBlock
	k.IterateUserBlocks(ctx, blocker, func(index int64, userblock types.UserBlock) (stop bool) {
		userblocks = append(userblocks, userblock)
		return false
	})
	return userblocks
}

// GetAllUsersBlocks returns a list of all the users blocks inside the given context.
func (k Keeper) GetAllUsersBlocks(ctx sdk.Context) []types.UserBlock {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.UsersBlocksStorePrefix)
	defer iterator.Close()

	var usersBlocks []types.UserBlock
	for ; iterator.Valid(); iterator.Next() {
		block := types.MustUnmarshalUserBlock(k.cdc, iterator.Value())
		usersBlocks = append(usersBlocks, block)
	}

	return usersBlocks
}

// DeleteUserBlock allows to the specified blocker to unblock the given blocked user
func (k Keeper) DeleteUserBlock(ctx sdk.Context, blocker, blocked string, subspace uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.UserBlockStoreKey(blocker, subspace, blocked))
}

// DeleteAllUserBlocks deletes all the user blocks that have been created by the given user
func (k Keeper) DeleteAllUserBlocks(ctx sdk.Context, user string) {
	var blocks []types.UserBlock
	k.IterateUserBlocks(ctx, user, func(index int64, block types.UserBlock) (stop bool) {
		blocks = append(blocks, block)
		return false
	})

	store := ctx.KVStore(k.storeKey)
	for _, block := range blocks {
		store.Delete(types.UserBlockStoreKey(block.Blocker, block.SubspaceID, block.Blocked))
	}
}
