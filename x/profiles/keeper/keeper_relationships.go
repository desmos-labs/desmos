package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

// SaveRelationship allows to store the given relationship returning an error if something goes wrong.
func (k Keeper) SaveRelationship(ctx sdk.Context, relationship types.Relationship) error {
	store := ctx.KVStore(k.StoreKey)
	key := types.RelationshipsStoreKey(relationship.Creator())

	var relationships types.Relationships
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &relationships)

	// TODO can we improve this?
	for _, r := range relationships {
		if r.Creator().Equals(relationship.Creator()) && r.Recipient().Equals(relationship.Recipient()) {
			switch rel := r.(type) {
			case types.MonodirectionalRelationship:
				if _, isType := relationship.(types.MonodirectionalRelationship); isType {
					return fmt.Errorf("relationship between %s and %s has already been done", rel.Sender, rel.Receiver)
				}
			case types.BidirectionalRelationship:
				if _, isType := relationship.(types.BidirectionalRelationship); isType {
					return fmt.Errorf("relationship between %s and %s has already been done", rel.Sender, rel.Receiver)
				}
			}
		}
	}
	relationships = append(relationships, relationship)
	store.Set(key, k.Cdc.MustMarshalBinaryBare(&relationships))

	return nil
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
	key := types.RelationshipsStoreKey(user)

	var relationships types.Relationships
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &relationships)

	return relationships
}
