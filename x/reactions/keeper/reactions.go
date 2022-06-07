package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v3/x/reactions/types"
)

// SetNextReactionID sets the next reaction id for the given subspace
func (k Keeper) SetNextReactionID(ctx sdk.Context, subspaceID uint64, reactionID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.NextReactionIDStoreKey(subspaceID), types.GetReactionIDBytes(reactionID))
}

// HasNextReactionID tells whether the next reaction id exists for the given subspace
func (k Keeper) HasNextReactionID(ctx sdk.Context, subspaceID uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.NextReactionIDStoreKey(subspaceID))
}

// GetNextReactionID gets the next reaction id for the given subspace
func (k Keeper) GetNextReactionID(ctx sdk.Context, subspaceID uint64) (reactionID uint64, err error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NextReactionIDStoreKey(subspaceID))
	if bz == nil {
		return 0, sdkerrors.Wrapf(types.ErrInvalidGenesis, "initial reaction id not set for subspace %d", subspaceID)
	}

	reactionID = types.GetReactionIDFromBytes(bz)
	return reactionID, nil
}

// DeleteNextReactionID removes the next reaction id for the given subspace
func (k Keeper) DeleteNextReactionID(ctx sdk.Context, subspaceID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.NextReactionIDStoreKey(subspaceID))
}

// --------------------------------------------------------------------------------------------------------------------

// SaveReaction saves the given reaction inside the current context
func (k Keeper) SaveReaction(ctx sdk.Context, reaction types.Reaction) {
	store := ctx.KVStore(k.storeKey)

	// Store the reaction
	store.Set(types.ReactionStoreKey(reaction.SubspaceID, reaction.ID), k.cdc.MustMarshal(&reaction))

	k.AfterReactionSaved(ctx, reaction.SubspaceID, reaction.ID)
}

// HasReaction tells whether the given reaction exists or not
func (k Keeper) HasReaction(ctx sdk.Context, subspaceID uint64, reactionID uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.ReactionStoreKey(subspaceID, reactionID))
}

// GetReaction returns the reaction associated with the given id.
// If there is no reaction with the given id the function will return an empty reaction and false.
func (k Keeper) GetReaction(ctx sdk.Context, subspaceID uint64, reactionID uint64) (reaction types.Reaction, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ReactionStoreKey(subspaceID, reactionID))
	if bz == nil {
		return types.Reaction{}, false
	}

	k.cdc.MustUnmarshal(bz, &reaction)
	return reaction, true
}

// DeleteReaction deletes the reaction having the given id from the store
func (k Keeper) DeleteReaction(ctx sdk.Context, subspaceID uint64, reactionID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.ReactionStoreKey(subspaceID, reactionID))

	k.AfterReactionDeleted(ctx, subspaceID, reactionID)
}
