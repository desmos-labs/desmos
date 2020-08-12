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
	return store.Has(types.RelationshipsStoreKey(id))
}

// StoreRelationship allows to store the given relationship returning an error if something goes wrong.
func (k Keeper) StoreRelationship(ctx sdk.Context, relationship types.Relationship) {
	store := ctx.KVStore(k.StoreKey)
	key := types.RelationshipsStoreKey(relationship.RelationshipID())
	store.Set(key, k.Cdc.MustMarshalBinaryBare(&relationship))
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

// deleteRelationshipFromArray remove the relationship with the given relationshipID from the given array
func deleteRelationshipFromArray(keeper Keeper, store sdk.KVStore, storeKey []byte,
	relationshipID types.RelationshipID) (relationshipsIDs []types.RelationshipID) {
	keeper.Cdc.MustUnmarshalBinaryBare(store.Get(storeKey), &relationshipsIDs)
	for index, id := range relationshipsIDs {
		if id.Equals(relationshipID) {
			relationshipsIDs = append(relationshipsIDs[:index], relationshipsIDs[index+1:]...)
			return relationshipsIDs
		}
	}
	// never reached
	return relationshipsIDs
}

// DeleteRelationship allows to delete the relationship between the given user and his counterparty
func (k Keeper) DeleteRelationship(ctx sdk.Context, relationshipID types.RelationshipID, user sdk.AccAddress) error {
	store := ctx.KVStore(k.StoreKey)
	key := types.RelationshipsStoreKey(relationshipID)
	userKey := types.UserRelationshipsStoreKey(user)

	// Get the stored relationship associated to the given relationshipID
	var relationship types.Relationship
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &relationship)

	// Check the relationship type to know how many user -> relationship couple to delete
	switch relationship.(type) {
	case types.MonodirectionalRelationship:
		if !relationship.Creator().Equals(user) {
			return fmt.Errorf("user with address %s isn't the relationship's creator", user)
		}
	case types.BidirectionalRelationship:
		if !relationship.Creator().Equals(user) && !relationship.Recipient().Equals(user) {
			return fmt.Errorf("user with address %s is neither the creator nor the recipient of the relationship", user)
		}

		if !relationship.Creator().Equals(user) {
			// delete creator -> relationship association
			creatorKey := types.UserRelationshipsStoreKey(relationship.Creator())
			relIds := deleteRelationshipFromArray(k, store, creatorKey, relationshipID)
			store.Set(creatorKey, k.Cdc.MustMarshalBinaryBare(&relIds))
		}

		if !relationship.Recipient().Equals(user) {
			// delete receiver -> relationship association
			recipientKey := types.UserRelationshipsStoreKey(relationship.Recipient())
			relIds := deleteRelationshipFromArray(k, store, recipientKey, relationshipID)
			store.Set(recipientKey, k.Cdc.MustMarshalBinaryBare(&relIds))
		}
	}

	// delete user -> relationship association
	relIds := deleteRelationshipFromArray(k, store, userKey, relationshipID)
	store.Set(userKey, k.Cdc.MustMarshalBinaryBare(&relIds))

	// delete stored relationship
	store.Delete(key)

	return nil
}
