package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// SaveUserBlock allows to store the given block inside the store, returning an error if
// something goes wrong.
func (k Keeper) SaveUserBlock(ctx sdk.Context, userBlock types.UserBlock) error {
	store := ctx.KVStore(k.storeKey)
	key := types.UsersBlocksStoreKey(userBlock.Blocker)

	blocks := types.MustUnmarshalUserBlocks(k.cdc, store.Get(key))
	for _, ub := range blocks {
		if ub == userBlock {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"the user with address %s has already been blocked", userBlock.Blocked)
		}
	}

	store.Set(key, types.MustMarshalUserBlocks(k.cdc, append(blocks, userBlock)))

	k.Logger(ctx).Info("blocked user", "blocked", userBlock.Blocked, "from", userBlock.Blocker)
	return nil
}

// DeleteUserBlock allows to the specified blocker to unblock the given blocked user.
func (k Keeper) DeleteUserBlock(ctx sdk.Context, blocker, blocked string, subspace string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.UsersBlocksStoreKey(blocker)

	blocks := types.MustUnmarshalUserBlocks(k.cdc, store.Get(key))

	blocks, found := types.RemoveUserBlock(blocks, blocker, blocked, subspace)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"block from %s towards %s for subspace %s not found", blocker, blocked, subspace)
	}

	// Delete the key if no blocks are left.
	// This cleans up the store avoiding export/import tests to fail due to a different number of keys present.
	if len(blocks) == 0 {
		store.Delete(key)
	} else {
		store.Set(key, types.MustMarshalUserBlocks(k.cdc, blocks))
	}

	k.Logger(ctx).Info("unblocked user", "unblocked", blocked, "from", blocker)
	return nil
}

// GetUserBlocks returns the list of users that the specified user has blocked.
func (k Keeper) GetUserBlocks(ctx sdk.Context, user string) []types.UserBlock {
	store := ctx.KVStore(k.storeKey)
	return types.MustUnmarshalUserBlocks(k.cdc, store.Get(types.UsersBlocksStoreKey(user)))
}

// GetAllUsersBlocks returns a list of all the users blocks inside the given context.
func (k Keeper) GetAllUsersBlocks(ctx sdk.Context) []types.UserBlock {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.UsersBlocksStorePrefix)
	defer iterator.Close()

	var usersBlocks []types.UserBlock
	for ; iterator.Valid(); iterator.Next() {
		blocks := types.MustUnmarshalUserBlocks(k.cdc, iterator.Value())
		usersBlocks = append(usersBlocks, blocks...)
	}

	return usersBlocks
}

// IsUserBlocked tells if the given blocker has blocked the given blocked user
func (k Keeper) IsUserBlocked(ctx sdk.Context, blocker, blocked string) bool {
	return k.HasUserBlocked(ctx, blocker, blocked, "")
}

// HasUserBlocked returns true if the provided blocker has blocked the given user for the given subspace.
// If the provided subspace is empty, all subspaces will be checked
func (k Keeper) HasUserBlocked(ctx sdk.Context, blocker, user, subspace string) bool {
	blocks := k.GetUserBlocks(ctx, blocker)

	for _, block := range blocks {
		if block.Blocked == user {
			return subspace == "" || block.Subspace == subspace
		}
	}

	return false
}
