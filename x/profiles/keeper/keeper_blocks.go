package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// SaveUserBlock allows to store the given block inside the store, returning an error if
// something goes wrong.
// It requires the blocker to have a registered profile.
func (k Keeper) SaveUserBlock(ctx sdk.Context, userBlock types.UserBlock) error {
	// Check the blocker to make sure they have a profile
	if !k.HasProfile(ctx, userBlock.Blocker) {
		return sdkerrors.Wrapf(types.ErrProfileNotFound, "blocker does not have a profile")
	}

	// Check to make sure the blocker and blocked users are not the same
	if userBlock.Blocker == userBlock.Blocked {
		return sdkerrors.Wrap(types.ErrInvalidBlock, "blocker and blocked cannot be the same user")
	}

	store := ctx.KVStore(k.storeKey)
	key := types.UserBlockStoreKey(userBlock.Blocker, userBlock.Subspace, userBlock.Blocked)
	if store.Has(key) {
		return sdkerrors.Wrapf(types.ErrBlockAlreadyCreated,
			"the user with address %s has already been blocked", userBlock.Blocked)
	}

	store.Set(key, k.cdc.MustMarshalBinaryBare(&userBlock))
	return nil
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

// IsUserBlocked tells if the given blocker has blocked the given blocked user
func (k Keeper) IsUserBlocked(ctx sdk.Context, blocker, blocked string) bool {
	return k.HasUserBlocked(ctx, blocker, blocked, "")
}

// HasUserBlocked returns true if the provided blocker has blocked the given user for the given subspace.
// If the provided subspace is empty, all subspaces will be checked
func (k Keeper) HasUserBlocked(ctx sdk.Context, blocker, blocked, subspace string) bool {
	if subspace != "" {
		store := ctx.KVStore(k.storeKey)
		key := types.UserBlockStoreKey(blocker, subspace, blocked)

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

// DeleteUserBlock allows to the specified blocker to unblock the given blocked user.
func (k Keeper) DeleteUserBlock(ctx sdk.Context, blocker, blocked string, subspace string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.UserBlockStoreKey(blocker, subspace, blocked)
	if !store.Has(key) {
		return sdkerrors.Wrapf(types.ErrBlockNotFound,
			"block from %s towards %s for subspace %s not found", blocker, blocked, subspace)
	}
	store.Delete(key)
	return nil
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
		store.Delete(types.UserBlockStoreKey(block.Blocker, block.Subspace, block.Blocked))
	}
}
