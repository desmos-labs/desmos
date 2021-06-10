package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// SaveRelationship allows to store the given relationship returning an error if he's already present.
func (k Keeper) SaveRelationship(ctx sdk.Context, relationship types.Relationship) error {
	store := ctx.KVStore(k.storeKey)
	key := types.RelationshipsStoreKey(relationship.Creator, relationship.Subspace, relationship.Recipient)

	if store.Has(key) {
		
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "relationship already exists with %s", relationship.Recipient)
	}

	store.Set(key, types.MustMarshalRelationship(k.cdc, relationship))
	return nil
}

func (k Keeper) GetRelationship(ctx sdk.Context, creator, subspace, recipient string) (types.Relationship, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.RelationshipsStoreKey(creator, subspace, recipient)

	if !store.Has(key) {
		return types.Relationship{}, false
	}

	bz := store.Get(key)

	var rel types.Relationship
	k.cdc.MustUnmarshalBinaryBare(bz, &rel)
	return rel, true
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
func (k Keeper) RemoveRelationship(ctx sdk.Context, relationship types.Relationship) error {
	store := ctx.KVStore(k.storeKey)
	key := types.RelationshipsStoreKey(relationship.Creator, relationship.Subspace, relationship.Recipient)
	if !store.Has(key) {
		return fmt.Errorf("relationship does not exist")
	}
	store.Delete(key)
	return nil
}
