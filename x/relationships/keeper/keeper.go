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

// GetUserRelationships allows to list all the storedRelationships that involve the given user.
func (k Keeper) GetUserRelationships(ctx sdk.Context, user sdk.AccAddress) []sdk.AccAddress {
	store := ctx.KVStore(k.StoreKey)
	key := types.RelationshipsStoreKey(user)

	var relationships []sdk.AccAddress
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &relationships)

	return relationships
}

// GetUsersRelationships allows to returns the map of all the users and their associated storedRelationships
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
