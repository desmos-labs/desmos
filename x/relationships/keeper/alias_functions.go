package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v7/x/relationships/types"
)

// IterateRelationships iterates through the relationships and perform the provided function
func (k Keeper) IterateRelationships(ctx sdk.Context, fn func(index int64, relationship types.Relationship) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.RelationshipsStorePrefix)
	defer iterator.Close()

	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {
		relationship := types.MustUnmarshalRelationship(k.cdc, iterator.Value())

		stop := fn(i, relationship)

		if stop {
			break
		}

		i++
	}
}

// IterateSubspaceRelationships iterates over all the relationships present inside the given subspace
func (k Keeper) IterateSubspaceRelationships(ctx sdk.Context, subspaceID uint64, fn func(index int64, relationship types.Relationship) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.SubspaceRelationshipsPrefix(subspaceID))
	defer iterator.Close()

	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {
		relationship := types.MustUnmarshalRelationship(k.cdc, iterator.Value())
		stop := fn(i, relationship)
		if stop {
			break
		}
		i++
	}
}

// IterateUsersBlocks iterates through the list of user blocks and performs the given function
func (k Keeper) IterateUsersBlocks(ctx sdk.Context, fn func(index int64, block types.UserBlock) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.UsersBlocksStorePrefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		block := types.MustUnmarshalUserBlock(k.cdc, iterator.Value())
		stop := fn(i, block)
		if stop {
			break
		}
		i++
	}
}

// IterateSubspaceUsersBlocks iterates through the list of user blocks stored for a subspace and performs the given function
func (k Keeper) IterateSubspaceUsersBlocks(ctx sdk.Context, subspaceID uint64, fn func(index int64, block types.UserBlock) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SubspaceBlocksPrefix(subspaceID))
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		block := types.MustUnmarshalUserBlock(k.cdc, iterator.Value())
		stop := fn(i, block)
		if stop {
			break
		}
		i++
	}
}
