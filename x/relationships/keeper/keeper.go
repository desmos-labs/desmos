package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/relationships/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey          // Unexposed key to access store from sdk.Context
	cdc      codec.BinaryMarshaler // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the magpie Keeper
func NewKeeper(cdc codec.BinaryMarshaler, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

// StoreRelationship allows to store the given relationship returning an error if he's already present.
func (k Keeper) StoreRelationship(ctx sdk.Context, relationship types.Relationship) error {
	store := ctx.KVStore(k.storeKey)
	key := types.RelationshipsStoreKey(relationship.Creator)

	relationships, err := k.UnmarshalRelationships(store.Get(key))
	if err != nil {
		return err
	}

	for _, rel := range relationships {
		if rel.Equal(relationship) {
			return fmt.Errorf("relationship already exists with %s", relationship.Recipient)
		}
	}

	relationships = append(relationships, relationship)
	bz, err := k.MarshalRelationships(relationships)
	if err != nil {
		return err
	}

	store.Set(key, bz)
	return nil
}

// GetUserRelationships allows to list all the stored that involve the given user.
func (k Keeper) GetUserRelationships(ctx sdk.Context, user string) ([]types.Relationship, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.RelationshipsStoreKey(user))
	return k.UnmarshalRelationships(bz)
}

// GetAllRelationships allows to returns the map of all the users and their associated stored
func (k Keeper) GetAllRelationships(ctx sdk.Context) ([]types.Relationship, error) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RelationshipsStorePrefix)

	var relationships []types.Relationship
	for ; iterator.Valid(); iterator.Next() {
		userRelationships, err := k.UnmarshalRelationships(iterator.Value())
		if err != nil {
			return nil, err
		}

		relationships = append(relationships, userRelationships...)
	}

	return relationships, nil
}

// DeleteRelationship allows to delete the relationship between the given user and his counterparty
func (k Keeper) DeleteRelationship(ctx sdk.Context, relationship types.Relationship) error {
	store := ctx.KVStore(k.storeKey)
	key := types.RelationshipsStoreKey(relationship.Creator)

	relationships, err := k.UnmarshalRelationships(store.Get(key))
	if err != nil {
		return err
	}

	for index, rel := range relationships {
		if rel.Recipient == relationship.Recipient && rel.Subspace == relationship.Subspace {
			relationships = append(relationships[:index], relationships[index+1:]...)
			if len(relationships) == 0 {
				store.Delete(key)
			} else {
				bz, err := k.MarshalRelationships(relationships)
				if err != nil {
					return err
				}

				store.Set(key, bz)
			}
			break
		}
	}

	return nil
}

// MarshalRelationships marshals a list of Relationships. If the given type implements
// the Marshaler interface, it is treated as a Proto-defined message and
// serialized that way. Otherwise, it falls back on the internal Amino codec.
func (k Keeper) MarshalRelationships(relationships []types.Relationship) ([]byte, error) {
	return codec.MarshalAny(k.cdc, relationships)
}

// UnmarshalRelationships returns a list of Relationship interface from raw encoded
// relationships. An error is returned upon decoding failure.
func (k Keeper) UnmarshalRelationships(bz []byte) ([]types.Relationship, error) {
	var relationships []types.Relationship
	if err := codec.UnmarshalAny(k.cdc, &relationships, bz); err != nil {
		return nil, err
	}

	return relationships, nil
}

// -------------------
// --- Users blocks
// -------------------

// SaveUserBlock allows to store the given block inside the store, returning an error if
// something goes wrong.
func (k Keeper) SaveUserBlock(ctx sdk.Context, userBlock types.UserBlock) error {
	store := ctx.KVStore(k.storeKey)
	key := types.UsersBlocksStoreKey(userBlock.Blocker)

	usersBlocks, err := k.UnmarshalUserBlocks(store.Get(key))
	if err != nil {
		return err
	}

	for _, ub := range usersBlocks {
		if ub == userBlock {
			return fmt.Errorf("the user with %s address has already been blocked", userBlock.Blocked)
		}
	}

	usersBlocks = append(usersBlocks, userBlock)
	bz, err := k.MarshalUserBlocks(usersBlocks)
	if err != nil {
		return err
	}

	store.Set(key, bz)
	return nil
}

// DeleteUserBlock allows to the specified blocker to unblock the given blocked user.
func (k Keeper) DeleteUserBlock(ctx sdk.Context, blocker, blocked string, subspace string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.UsersBlocksStoreKey(blocker)

	usersBlocks, err := k.UnmarshalUserBlocks(store.Get(key))
	if err != nil {
		return err
	}

	for index, ub := range usersBlocks {
		if ub.Blocker == blocker && ub.Blocked == blocked && ub.Subspace == subspace {
			usersBlocks = append(usersBlocks[:index], usersBlocks[index+1:]...)
			if len(usersBlocks) == 0 {
				store.Delete(key)
			} else {
				bz, err := k.MarshalUserBlocks(usersBlocks)
				if err != nil {
					return err
				}
				store.Set(key, bz)
			}
			return nil
		}
	}

	return fmt.Errorf("blocked user with address %s not found", blocked)
}

// GetUserBlocks returns the list of users that the specified user has blocked.
func (k Keeper) GetUserBlocks(ctx sdk.Context, user string) ([]types.UserBlock, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.UsersBlocksStoreKey(user))
	return k.UnmarshalUserBlocks(bz)
}

// GetUsersBlocks returns a list of all the users blocks inside the given context.
func (k Keeper) GetUsersBlocks(ctx sdk.Context) ([]types.UserBlock, error) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.UsersBlocksStorePrefix)

	var usersBlocks []types.UserBlock
	for ; iterator.Valid(); iterator.Next() {
		userBlocks, err := k.UnmarshalUserBlocks(iterator.Value())
		if err != nil {
			return nil, err
		}

		usersBlocks = append(usersBlocks, userBlocks...)
	}

	return usersBlocks, nil
}

// MarshalUserBlocks marshals a list of UserBlock. If the given type implements
// the Marshaler interface, it is treated as a Proto-defined message and
// serialized that way. Otherwise, it falls back on the internal Amino codec.
func (k Keeper) MarshalUserBlocks(blocks []types.UserBlock) ([]byte, error) {
	return codec.MarshalAny(k.cdc, blocks)
}

// UnmarshalUserBlocks returns a list of UserBlock interface from raw encoded
// blocks bytes. An error is returned upon decoding failure.
func (k Keeper) UnmarshalUserBlocks(bz []byte) ([]types.UserBlock, error) {
	var blocks []types.UserBlock
	if err := codec.UnmarshalAny(k.cdc, &blocks, bz); err != nil {
		return nil, err
	}

	return blocks, nil
}
