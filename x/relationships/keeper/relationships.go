package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v2/x/relationships/types"
)

// SaveRelationship allows to store the given relationship returning an error if something goes wrong
func (k Keeper) SaveRelationship(ctx sdk.Context, relationship types.Relationship) error {
	// Check to make sure the creator and recipient are not the same
	if relationship.Creator == relationship.Counterparty {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "relationship creator and counterparty cannot be the same user")
	}

	store := ctx.KVStore(k.storeKey)
	key := types.RelationshipsStoreKey(relationship.Creator, relationship.Counterparty, relationship.SubspaceID)
	store.Set(key, types.MustMarshalRelationship(k.cdc, relationship))
	return nil
}

// HasRelationship tells whether the relationship between the creator and counterparty
// already exists for the given subspace
func (k Keeper) HasRelationship(ctx sdk.Context, creator, counterparty string, subspace uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.RelationshipsStoreKey(creator, counterparty, subspace))
}

// GetRelationship returns the relationship existing between the provided creator and recipient inside the given subspace
func (k Keeper) GetRelationship(ctx sdk.Context, user, counterparty string, subspaceID uint64) (types.Relationship, bool) {
	store := ctx.KVStore(k.storeKey)

	key := types.RelationshipsStoreKey(user, counterparty, subspaceID)
	if !store.Has(key) {
		return types.Relationship{}, false
	}

	return types.MustUnmarshalRelationship(k.cdc, store.Get(key)), true
}

// GetUserRelationships allows to list all the stored relationships that involve the given user.
func (k Keeper) GetUserRelationships(ctx sdk.Context, user string) []types.Relationship {
	var relationships []types.Relationship
	k.IterateUserRelationships(ctx, user, func(index int64, relationship types.Relationship) (stop bool) {
		relationships = append(relationships, relationship)
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
func (k Keeper) RemoveRelationship(ctx sdk.Context, relationship types.Relationship) {
	store := ctx.KVStore(k.storeKey)
	key := types.RelationshipsStoreKey(relationship.Creator, relationship.Counterparty, relationship.SubspaceID)
	store.Delete(key)
}

// DeleteAllUserRelationships removes all the relationships that somehow involve the given user
func (k Keeper) DeleteAllUserRelationships(ctx sdk.Context, user string) {
	var relationships []types.Relationship
	k.IterateUserRelationships(ctx, user, func(index int64, relationship types.Relationship) (stop bool) {
		relationships = append(relationships, relationship)
		return false
	})

	store := ctx.KVStore(k.storeKey)
	for _, relationship := range relationships {
		store.Delete(types.RelationshipsStoreKey(relationship.Creator, relationship.Counterparty, relationship.SubspaceID))
	}
}
