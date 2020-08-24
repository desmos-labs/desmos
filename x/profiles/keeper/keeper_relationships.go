package keeper

import (
	"bytes"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

// SaveUserRelationshipAssociation allows to save the user/relationship association
func (k Keeper) SaveUserRelationshipAssociation(ctx sdk.Context, users []sdk.AccAddress, id types.RelationshipID) {
	store := ctx.KVStore(k.StoreKey)
	for _, user := range users {
		key := types.UserRelationshipsStoreKey(user)
		var ids types.RelationshipIDs
		k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &ids)
		ids = append(ids, id)
		store.Set(key, k.Cdc.MustMarshalBinaryBare(&ids))
	}
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

// GetRelationships returns all the relationships inside the current context
func (k Keeper) GetRelationships(ctx sdk.Context) types.Relationships {
	relationships := make(types.Relationships, 0)
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RelationshipsStorePrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var rel types.Relationship
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &rel)
		relationships = append(relationships, rel)
	}

	return relationships
}

// GetRelationshipFromID returns the relationship associated with the given id
func (k Keeper) GetRelationshipFromID(ctx sdk.Context, id types.RelationshipID) (rel types.Relationship, err error) {
	store := ctx.KVStore(k.StoreKey)
	bz := store.Get(types.RelationshipsStoreKey(id))
	if bz == nil {
		return rel, fmt.Errorf("relationship with id %s doesn't exist", id)
	}
	k.Cdc.MustUnmarshalBinaryBare(bz, &rel)
	return rel, nil
}

// GetUserRelationships allows to list all the relationships that involve the given user.
func (k Keeper) GetUserRelationships(ctx sdk.Context, user sdk.AccAddress) types.Relationships {
	store := ctx.KVStore(k.StoreKey)
	key := types.UserRelationshipsStoreKey(user)

	// Get all the relationshipIDs related to the given user
	var relationshipsIDs types.RelationshipIDs
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &relationshipsIDs)

	var relationships types.Relationships
	for _, relID := range relationshipsIDs {
		var relationship types.Relationship
		k.Cdc.MustUnmarshalBinaryBare(store.Get(types.RelationshipsStoreKey(relID)), &relationship)
		relationships = append(relationships, relationship)
	}

	return relationships
}

// GetUsersRelationshipsIDMap allows to returns the list of answers that have been stored inside the given context
func (k Keeper) GetUsersRelationshipsIDMap(ctx sdk.Context) map[string]types.RelationshipIDs {
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, types.UserRelationshipsStorePrefix)

	usersRelationshipsMap := map[string]types.RelationshipIDs{}
	for ; iterator.Valid(); iterator.Next() {
		var relationshipIDs types.RelationshipIDs
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &relationshipIDs)
		userBytes := bytes.TrimPrefix(iterator.Key(), types.UserRelationshipsStorePrefix)
		userAddr := sdk.AccAddress(userBytes)
		usersRelationshipsMap[userAddr.String()] = relationshipIDs
	}

	return usersRelationshipsMap
}

// deleteRelationshipFromArray remove the relationship with the given relationshipID from the given array
func deleteRelationshipFromArray(keeper Keeper, store sdk.KVStore, storeKey []byte, relationshipID types.RelationshipID) {
	var relationshipsIDs types.RelationshipIDs
	keeper.Cdc.MustUnmarshalBinaryBare(store.Get(storeKey), &relationshipsIDs)
	for index, id := range relationshipsIDs {
		if id.Equals(relationshipID) {
			relationshipsIDs = append(relationshipsIDs[:index], relationshipsIDs[index+1:]...)
			if len(relationshipsIDs) == 0 {
				store.Delete(storeKey)
			} else {
				store.Set(storeKey, keeper.Cdc.MustMarshalBinaryBare(&relationshipsIDs))
			}
			break
		}
	}
}

// DeleteRelationship allows to delete the relationship between the given user and his counterparty
// It assumes that the given relationshipID is related to an existent relationship
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

		// delete creator -> relationship association
		if !relationship.Creator().Equals(user) {
			creatorKey := types.UserRelationshipsStoreKey(relationship.Creator())
			deleteRelationshipFromArray(k, store, creatorKey, relationshipID)
		}

		// delete receiver -> relationship association
		if !relationship.Recipient().Equals(user) {
			recipientKey := types.UserRelationshipsStoreKey(relationship.Recipient())
			deleteRelationshipFromArray(k, store, recipientKey, relationshipID)
		}
	}

	// delete user -> relationship association
	deleteRelationshipFromArray(k, store, userKey, relationshipID)

	// delete stored relationship
	store.Delete(key)

	return nil
}
