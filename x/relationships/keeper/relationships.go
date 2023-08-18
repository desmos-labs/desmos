package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/x/relationships/types"
)

// SaveRelationship allows to store the given relationship returning an error if something goes wrong
func (k Keeper) SaveRelationship(ctx sdk.Context, relationship types.Relationship) {
	store := ctx.KVStore(k.storeKey)
	key := types.RelationshipsStoreKey(relationship.Creator, relationship.Counterparty, relationship.SubspaceID)
	store.Set(key, types.MustMarshalRelationship(k.cdc, relationship))
}

// HasRelationship tells whether the relationship between the creator and counterparty
// already exists for the given subspace
func (k Keeper) HasRelationship(ctx sdk.Context, user, counterparty string, subspaceID uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.RelationshipsStoreKey(user, counterparty, subspaceID))
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

// DeleteRelationship deletes the given relationship
func (k Keeper) DeleteRelationship(ctx sdk.Context, user, counterparty string, subspaceID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.RelationshipsStoreKey(user, counterparty, subspaceID))
}
