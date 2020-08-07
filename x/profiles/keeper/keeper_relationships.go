package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

// SaveRelationship allows to store the given relationship returning an error if something goes wrong.
func (k Keeper) SaveRelationship(ctx sdk.Context, relationship types.Relationship) error {
	store := ctx.KVStore(k.StoreKey)
	key := types.RelationshipsStoreKey(relationship.Creator())

	var relationships []types.Relationship
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &relationships)

	// TODO study how to handle interface inside keeper
}

// DeleteRelationship allows to delete the relationship between the given user and his counterparty
func (k Keeper) DeleteRelationship(ctx sdk.Context, user, counterparty sdk.AccAddress) error {

}

// GetUserRelationships allows to list all the relationships that involve the given user.
func (k Keeper) GetUserRelationships(ctx sdk.Context, user sdk.AccAddress) []types.Relationship {

}
