package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// SaveRelationship allows to store the given relationship returning an error if he's already present.
// It requires the creator to have a profile.
func (k Keeper) SaveRelationship(ctx sdk.Context, relationship types.Relationship) error {
	// Check the creator to make sure they have a profile
	if !k.HasProfile(ctx, relationship.Creator) {
		return sdkerrors.Wrap(types.ErrProfileNotFound, "relationship creator does not have a profile")
	}

	// Check to make sure the creator and recipient are not the same
	if relationship.Creator == relationship.Recipient {
		return sdkerrors.Wrap(types.ErrInvalidRelationship,
			"relationship creator and recipient cannot be the same user")
	}

	store := ctx.KVStore(k.storeKey)
	key := types.RelationshipsStoreKey(relationship.Creator, relationship.Subspace, relationship.Recipient)

	if store.Has(key) {
		return sdkerrors.Wrapf(types.ErrDuplicatedRelationship, "recipient: %s", relationship.Recipient)
	}

	store.Set(key, types.MustMarshalRelationship(k.cdc, relationship))
	return nil
}

// GetRelationship returns the relationship existing between the provided creator and recipient inside the given subspace
func (k Keeper) GetRelationship(ctx sdk.Context, creator, subspace, recipient string) (types.Relationship, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.RelationshipsStoreKey(creator, subspace, recipient)

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
func (k Keeper) RemoveRelationship(ctx sdk.Context, relationship types.Relationship) error {
	store := ctx.KVStore(k.storeKey)
	key := types.RelationshipsStoreKey(relationship.Creator, relationship.Subspace, relationship.Recipient)
	if !store.Has(key) {
		return sdkerrors.Wrapf(types.ErrRelationshipNotFound,
			"relationship between %s and %s for subspace %s not found",
			relationship.Creator, relationship.Recipient, relationship.Subspace)
	}
	store.Delete(key)
	return nil
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
		store.Delete(types.RelationshipsStoreKey(relationship.Creator, relationship.Subspace, relationship.Recipient))
	}
}
