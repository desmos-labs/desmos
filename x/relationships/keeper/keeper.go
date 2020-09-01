package keeper

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/relationships/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	StoreKey sdk.StoreKey // Unexposed key to access store from sdk.Context
	Cdc      *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the magpie Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		StoreKey: storeKey,
		Cdc:      cdc,
	}
}

// StoreRelationship allows to store the given receiver returning an error if he's already present.
func (k Keeper) StoreRelationship(ctx sdk.Context, user, receiver sdk.AccAddress) error {
	store := ctx.KVStore(k.StoreKey)
	key := types.RelationshipsStoreKey(user)
	var relationships []sdk.AccAddress
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &relationships)

	for _, addr := range relationships {
		if addr.Equals(receiver) {
			return fmt.Errorf("relationship already exists with %s", receiver)
		}
	}

	relationships = append(relationships, receiver)
	store.Set(key, k.Cdc.MustMarshalBinaryBare(&relationships))

	return nil
}

// GetUserRelationships allows to list all the storedUserBlocks that involve the given user.
func (k Keeper) GetUserRelationships(ctx sdk.Context, user sdk.AccAddress) []sdk.AccAddress {
	store := ctx.KVStore(k.StoreKey)
	key := types.RelationshipsStoreKey(user)

	var relationships []sdk.AccAddress
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &relationships)

	return relationships
}

// GetUsersRelationships allows to returns the map of all the users and their associated storedUserBlocks
func (k Keeper) GetUsersRelationships(ctx sdk.Context) map[string][]sdk.AccAddress {
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RelationshipsStorePrefix)

	usersRelationshipsMap := map[string][]sdk.AccAddress{}
	var relationships []sdk.AccAddress
	for ; iterator.Valid(); iterator.Next() {
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &relationships)
		userBytes := bytes.TrimPrefix(iterator.Key(), types.RelationshipsStorePrefix)
		userAddr := sdk.AccAddress(userBytes)
		usersRelationshipsMap[userAddr.String()] = relationships
	}

	return usersRelationshipsMap
}

// DeleteRelationship allows to delete the relationship between the given user and his counterparty
func (k Keeper) DeleteRelationship(ctx sdk.Context, user, receiver sdk.AccAddress) {
	store := ctx.KVStore(k.StoreKey)
	key := types.RelationshipsStoreKey(user)
	var relationships []sdk.AccAddress
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &relationships)

	for index, addr := range relationships {
		if addr.Equals(receiver) {
			relationships = append(relationships[:index], relationships[index+1:]...)
			if len(relationships) == 0 {
				store.Delete(key)
			} else {
				store.Set(key, k.Cdc.MustMarshalBinaryBare(&relationships))
			}
			break
		}
	}
}

/////////////////////
/////UserBlocks/////
///////////////////

// SaveUserBlock allows to store the given block inside the store, returning an error if
// something goes wrong.
func (k Keeper) SaveUserBlock(ctx sdk.Context, userBlock types.UserBlock) error {
	store := ctx.KVStore(k.StoreKey)
	key := types.UsersBlocksStoreKey(userBlock.Blocker)
	var usersBlocks []types.UserBlock
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &usersBlocks)

	for _, ub := range usersBlocks {
		if ub.Equals(userBlock) {
			return fmt.Errorf("the user with %s address has been blocked already", userBlock.Blocked)
		}
	}

	usersBlocks = append(usersBlocks, userBlock)
	store.Set(key, k.Cdc.MustMarshalBinaryBare(&usersBlocks))

	return nil
}

// UnblockUser allows to the specified blocker to unblock the given blocked user.
func (k Keeper) UnblockUser(ctx sdk.Context, blocker, blocked sdk.AccAddress) error {
	store := ctx.KVStore(k.StoreKey)
	key := types.UsersBlocksStoreKey(blocker)
	var usersBlocks []types.UserBlock
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &usersBlocks)

	for index, ub := range usersBlocks {
		if ub.Blocker.Equals(blocker) && ub.Blocked.Equals(blocked) {
			usersBlocks = append(usersBlocks[:index], usersBlocks[index+1:]...)
			if len(usersBlocks) == 0 {
				store.Delete(key)
			} else {
				store.Set(key, k.Cdc.MustMarshalBinaryBare(&usersBlocks))
			}
			return nil
		}
	}

	return fmt.Errorf("blocked user with address %s not found", blocked)
}

// GetUserBlocks returns the list of users that the specified user has blocked.
func (k Keeper) GetUserBlocks(ctx sdk.Context, user sdk.AccAddress) []types.UserBlock {
	store := ctx.KVStore(k.StoreKey)
	key := types.UsersBlocksStoreKey(user)

	var userBlocks []types.UserBlock
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &userBlocks)

	return userBlocks
}

// GetUsersBlocks returns a list of all the users blocks inside the given context.
func (k Keeper) GetUsersBlocks(ctx sdk.Context) []types.UserBlock {
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, types.UsersBlocksStorePrefix)

	var usersBlocks []types.UserBlock
	for ; iterator.Valid(); iterator.Next() {
		var userBlocks []types.UserBlock
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &userBlocks)
		usersBlocks = append(usersBlocks, userBlocks...)
	}

	return usersBlocks
}
