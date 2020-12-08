package keeper

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/relationships/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey          // Unexposed key to access store from sdk.Context
	cdc      codec.BinaryMarshaler // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the relationships Keeper
func NewKeeper(cdc codec.BinaryMarshaler, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

// SaveRelationship allows to store the given relationship returning an error if he's already present.
func (k Keeper) SaveRelationship(ctx sdk.Context, relationship types.Relationship) error {
	store := ctx.KVStore(k.storeKey)
	key := types.RelationshipsStoreKey(relationship.Creator)

	relationships := types.MustUnmarshalRelationships(k.cdc, store.Get(key))
	for _, rel := range relationships {
		if rel.Equal(relationship) {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"relationship already exists with %s", relationship.Recipient)
		}
	}

	relationships = append(relationships, relationship)
	store.Set(key, types.MustMarshalRelationships(k.cdc, relationships))
	return nil
}

// GetUserRelationships allows to list all the stored relationships that involve the given user.
func (k Keeper) GetUserRelationships(ctx sdk.Context, user string) []types.Relationship {
	var relationships []types.Relationship
	k.IterateRelationships(ctx, func(index int64, relationship types.Relationship) (stop bool) {
		if relationship.Creator == user || relationship.Recipient == user {
			relationships = append(relationships, relationship)
		}
		return false
	})
	return relationships
}

// GetAllRelationships allows to returns the list of all stored relationships
func (k Keeper) GetAllRelationships(ctx sdk.Context) []types.Relationship {
	var relationships []types.Relationship
	k.IterateRelationships(ctx, func(index int64, relationship types.Relationship) (stop bool) {
		relationships = append(relationships, relationship)
		return false
	})
	return relationships
}

// RemoveRelationship allows to delete the relationship between the given user and his counterparty
func (k Keeper) RemoveRelationship(ctx sdk.Context, relationship types.Relationship) error {
	store := ctx.KVStore(k.storeKey)
	key := types.RelationshipsStoreKey(relationship.Creator)

	relationships := types.MustUnmarshalRelationships(k.cdc, store.Get(key))
	relationships, found := types.RemoveRelationship(relationships, relationship)

	// The relationship didn't exist, so return an error
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"relationship between %s and %s for subspace %s not found",
			relationship.Creator, relationship.Recipient, relationship.Subspace)
	}

	// Delete the key if no relationships are left.
	// This cleans up the store avoiding export/import tests to fail due to a different number of keys present.
	if len(relationships) == 0 {
		store.Delete(key)
	} else {
		store.Set(key, types.MustMarshalRelationships(k.cdc, relationships))
	}
	return nil
}

// ___________________________________________________________________________________________________________________

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
