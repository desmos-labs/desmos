package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// SaveRelationship allows to store the given relationship returning an error if he's already present.
func (k Keeper) SaveRelationship(ctx sdk.Context, relationship types.Relationship) error {
	store := ctx.KVStore(k.storeKey)
	key := types.RelationshipsStoreKey(relationship.Creator)

	relationships := types.MustUnmarshalRelationships(k.cdc, store.Get(key))
	for _, rel := range relationships {
		if rel.Equal(relationship) {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"relationship already exists with %s", relationship.Recipient)
		}
	}

	relationships = append(relationships, relationship)
	store.Set(key, types.MustMarshalRelationships(k.cdc, relationships))
	return nil
}

// GetUserRelationships allows to list all the stored relationships that involve the given user.
func (k Keeper) GetUserRelationships(ctx sdk.Context, user string) []types.Relationship {
	var relationships []types.Relationship
	k.IterateRelationships(ctx, func(index int64, relationship types.Relationship) (stop bool) {
		if relationship.Creator == user || relationship.Recipient == user {
			relationships = append(relationships, relationship)
		}
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
	key := types.RelationshipsStoreKey(relationship.Creator)

	relationships := types.MustUnmarshalRelationships(k.cdc, store.Get(key))
	relationships, found := types.RemoveRelationship(relationships, relationship)

	// The relationship didn't exist, so return an error
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"relationship between %s and %s for subspace %s not found",
			relationship.Creator, relationship.Recipient, relationship.Subspace)
	}

	// Delete the key if no relationships are left.
	// This cleans up the store avoiding export/import tests to fail due to a different number of keys present.
	if len(relationships) == 0 {
		store.Delete(key)
	} else {
		store.Set(key, types.MustMarshalRelationships(k.cdc, relationships))
	}
	return nil
}
