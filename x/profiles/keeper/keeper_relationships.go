package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

// SaveUserRelationshipAssociation allows to save the user/relationship association
func (k Keeper) SaveUserRelationshipAssociation(ctx sdk.Context, user sdk.AccAddress, id types.RelationshipID) {
	store := ctx.KVStore(k.StoreKey)
	key := types.UserRelationshipsStoreKey(user)

	var ids []types.RelationshipID
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &ids)
	ids = append(ids, id)

	store.Set(key, k.Cdc.MustMarshalBinaryBare(&ids))
}

// DoesRelationshipExist checks if the given id has an associated relationship or not
func (k Keeper) DoesRelationshipExist(ctx sdk.Context, id types.RelationshipID) bool {
	store := ctx.KVStore(k.StoreKey)
	if store.Has(types.RelationshipsStoreKey(id)) {
		return true
	}
	return false
}

// StoreRelationship allows to store the given relationship returning an error if something goes wrong.
func (k Keeper) StoreRelationship(ctx sdk.Context, relationship types.Relationship) {
	store := ctx.KVStore(k.StoreKey)
	key := types.RelationshipsStoreKey(relationship.ID())
	store.Set(key, k.Cdc.MustMarshalBinaryBare(&relationship))
}

// DeleteRelationship allows to delete the relationship between the given user and his counterparty
func (k Keeper) DeleteRelationship(ctx sdk.Context, user, counterparty sdk.AccAddress) error {
	store := ctx.KVStore(k.StoreKey)
	key := types.RelationshipsStoreKey(user)

	var relationships types.Relationships
	deleted := false
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &relationships)

	for index, r := range relationships {
		if r.Creator().Equals(user) && r.Recipient().Equals(counterparty) {
			relationships = append(relationships[:index], relationships[index+1:]...)
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("no relationship found between %s and %s", user, counterparty)
	}

	store.Set(key, k.Cdc.MustMarshalBinaryBare(&relationships))
	return nil
}

// GetUserRelationships allows to list all the relationships that involve the given user.
func (k Keeper) GetUserRelationships(ctx sdk.Context, user sdk.AccAddress) types.Relationships {
	store := ctx.KVStore(k.StoreKey)
	key := types.UserRelationshipsStoreKey(user)

	// Get all the relationshipIDs related to the given user
	var relationshipsIDs []types.RelationshipID
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &relationshipsIDs)

	var relationships types.Relationships
	for _, relID := range relationshipsIDs {
		var relationship types.Relationship
		k.Cdc.MustUnmarshalBinaryBare(store.Get(types.RelationshipsStoreKey(relID)), &relationship)
		relationships = append(relationships, relationship)
	}

	return relationships
}
