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

// addAddressToStore insert the given address into the store identified by the given key.
// It returns true if the address is successfully inserted, false otherwise.
func (k Keeper) addAddressToStore(store sdk.KVStore, key []byte, address sdk.AccAddress) bool {
	var addresses []sdk.AccAddress
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &addresses)
	for _, addr := range addresses {
		if addr.Equals(address) {
			return false
		}
	}

	addresses = append(addresses, address)
	store.Set(key, k.Cdc.MustMarshalBinaryBare(&addresses))
	return true
}

// removeAddressFromStore remove the given address from the store identified by the given key.
func (k Keeper) removeAddressFromStore(store sdk.KVStore, key []byte, address sdk.AccAddress) {
	var addresses []sdk.AccAddress
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &addresses)

	for index, addr := range addresses {
		if addr.Equals(address) {
			addresses = append(addresses[:index], addresses[index+1:]...)
			if len(addresses) == 0 {
				store.Delete(key)
			} else {
				store.Set(key, k.Cdc.MustMarshalBinaryBare(&addresses))
			}
			break
		}
	}
}

// getUsersAssociatedAddresses returns the map of all the users related addresses
func (k Keeper) getUsersAssociatedAddresses(iterator sdk.Iterator) map[string][]sdk.AccAddress {
	usersAddressesMap := map[string][]sdk.AccAddress{}
	var relationships []sdk.AccAddress
	for ; iterator.Valid(); iterator.Next() {
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &relationships)
		userBytes := bytes.TrimPrefix(iterator.Key(), types.RelationshipsStorePrefix)
		userAddr := sdk.AccAddress(userBytes)
		usersAddressesMap[userAddr.String()] = relationships
	}

	return usersAddressesMap
}

// StoreRelationship allows to store the given receiver returning an error if he's already present.
func (k Keeper) StoreRelationship(ctx sdk.Context, user, receiver sdk.AccAddress) error {
	store := ctx.KVStore(k.StoreKey)
	key := types.RelationshipsStoreKey(user)

	if added := k.addAddressToStore(store, key, receiver); !added {
		return fmt.Errorf("relationship already exists with %s", receiver)
	}

	return nil
}

// GetUserRelationships allows to list all the storedRelationships that involve the given user.
func (k Keeper) GetUserRelationships(ctx sdk.Context, user sdk.AccAddress) []sdk.AccAddress {
	store := ctx.KVStore(k.StoreKey)
	key := types.RelationshipsStoreKey(user)

	var relationships []sdk.AccAddress
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &relationships)

	return relationships
}

// GetUsersRelationships allows to returns the map of all the users and their associated storedRelationships.
func (k Keeper) GetUsersRelationships(ctx sdk.Context) map[string][]sdk.AccAddress {
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RelationshipsStorePrefix)

	return k.getUsersAssociatedAddresses(iterator)
}

// DeleteRelationship allows to delete the relationship between the given user and his counterparty.
func (k Keeper) DeleteRelationship(ctx sdk.Context, user, receiver sdk.AccAddress) {
	store := ctx.KVStore(k.StoreKey)
	key := types.RelationshipsStoreKey(user)

	k.removeAddressFromStore(store, key, receiver)
}

// SaveUserBlock allows to store the given block inside the store, returning an error if
// something goes wrong.
func (k Keeper) SaveUserBlock(ctx sdk.Context, blocker, toBeBlocked sdk.AccAddress) error {
	store := ctx.KVStore(k.StoreKey)
	key := types.RelationshipsStoreKey(blocker)

	if added := k.addAddressToStore(store, key, toBeBlocked); !added {
		return fmt.Errorf("the user with %s address is blocked already", toBeBlocked)
	}

	return nil
}

// DeleteUserBlock allows to the specified blocker to unblock the given blocked user.
func (k Keeper) UnblockUser(ctx sdk.Context, blocker, blocked sdk.AccAddress) {
	store := ctx.KVStore(k.StoreKey)
	key := types.BlockedUsersStoreKey(blocker)
	k.removeAddressFromStore(store, key, blocked)
}

// GetUserBlockedUsers returns the list of users that the specified user has blocked.
func (k Keeper) GetUserBlockedUsers(ctx sdk.Context, user sdk.AccAddress) []sdk.AccAddress {
	store := ctx.KVStore(k.StoreKey)
	key := types.RelationshipsStoreKey(user)

	var relationships []sdk.AccAddress
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &relationships)

	return relationships
}

// GetUsersBlockedUsers returns a map of all the users and their associated blocked users.
func (k Keeper) GetUsersBlockedUsers(ctx sdk.Context) map[string][]sdk.AccAddress {
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, types.BlockedUsersStorePrefix)
	return k.getUsersAssociatedAddresses(iterator)
}
