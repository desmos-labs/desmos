package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// SaveUserBlock allows to store the given block in specific subspace inside the store, returning an error if
// something goes wrong.
func (k Keeper) SaveUserBlock(ctx sdk.Context, userBlock types.UserBlock) error {
	store := ctx.KVStore(k.storeKey)
	key := types.UsersBlocksStoreKey(userBlock.Blocker, userBlock.Subspace, userBlock.Blocked)
	if store.Has(key) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the user with address %s has already been blocked", userBlock.Blocked)
	}

	store.Set(key, k.cdc.MustMarshalBinaryBare(&userBlock))
	return nil
}

// DeleteUserBlock allows to the specified blocker in specific subspace to unblock the given blocked user.
func (k Keeper) DeleteUserBlock(ctx sdk.Context, blocker, blocked string, subspace string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.UsersBlocksStoreKey(blocker, subspace, blocked)
	if !store.Has(key) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"block from %s towards %s for subspace %s not found", blocker, blocked, subspace)
	}
	//Delete key directly since that is 1-to-1
	store.Delete(key)
	return nil
}

// GetUserBlocks returns the list of users that the specified user has blocked.
func (k Keeper) GetUserBlocks(ctx sdk.Context, blocker string) []types.UserBlock {
	var userblocks []types.UserBlock
	k.IterateBlockedUsers(ctx, blocker, func(index int64, userblock types.UserBlock) (stop bool) {
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

// IsUserBlocked tells if the given blocker has blocked the given blocked user
func (k Keeper) IsUserBlocked(ctx sdk.Context, blocker, blocked string) bool {
	return k.HasUserBlocked(ctx, blocker, blocked, "")
}

// HasUserBlocked returns true if the provided blocker has blocked the given user for the given subspace.
// If the provided subspace is empty, all subspaces will be checked
func (k Keeper) HasUserBlocked(ctx sdk.Context, blocker, blocked, subspace string) bool {
	if subspace != "" {
		store := ctx.KVStore(k.storeKey)
		key := types.UsersBlocksStoreKey(blocker, subspace, blocked)

		return store.Has(key)
	}

	blocks := k.GetUserBlocks(ctx, blocker)
	for _, block := range blocks {
		if block.Blocked == blocked {
			return subspace == "" || block.Subspace == subspace
		}
	}

	return false
}
