package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v3/x/reactions/types"
)

// SetNextRegisteredReactionID sets the next registered reaction id for the given subspace
func (k Keeper) SetNextRegisteredReactionID(ctx sdk.Context, subspaceID uint64, registeredReactionID uint32) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.NextRegisteredReactionIDStoreKey(subspaceID), types.GetRegisteredReactionIDBytes(registeredReactionID))
}

// HasNextRegisteredReactionID tells whether the next registered reaction id exists for the given subspace
func (k Keeper) HasNextRegisteredReactionID(ctx sdk.Context, subspaceID uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.NextRegisteredReactionIDStoreKey(subspaceID))
}

// GetNextRegisteredReactionID gets the next registered reaction id for the given subspace
func (k Keeper) GetNextRegisteredReactionID(ctx sdk.Context, subspaceID uint64) (registeredReactionID uint32, err error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NextRegisteredReactionIDStoreKey(subspaceID))
	if bz == nil {
		return 0, sdkerrors.Wrapf(types.ErrInvalidGenesis, "initial registered reaction id not set for subspace %d", subspaceID)
	}

	registeredReactionID = types.GetRegisteredReactionIDFromBytes(bz)
	return registeredReactionID, nil
}

// DeleteNextRegisteredReactionID removes the next registered reaction id for the given subspace
func (k Keeper) DeleteNextRegisteredReactionID(ctx sdk.Context, subspaceID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.NextRegisteredReactionIDStoreKey(subspaceID))
}

// --------------------------------------------------------------------------------------------------------------------

// SaveRegisteredReaction saves the given registered reaction inside the current context
func (k Keeper) SaveRegisteredReaction(ctx sdk.Context, reaction types.RegisteredReaction) {
	store := ctx.KVStore(k.storeKey)

	// Store the reaction
	store.Set(types.RegisteredReactionStoreKey(reaction.SubspaceID, reaction.ID), k.cdc.MustMarshal(&reaction))

	k.Logger(ctx).Debug("registered reaction saved", "subspace id", reaction.SubspaceID, "id", reaction.ID)
	k.AfterRegisteredReactionSaved(ctx, reaction.SubspaceID, reaction.ID)
}

// HasRegisteredReaction tells whether the given registered reaction exists or not
func (k Keeper) HasRegisteredReaction(ctx sdk.Context, subspaceID uint64, registeredReactionID uint32) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.RegisteredReactionStoreKey(subspaceID, registeredReactionID))
}

// GetRegisteredReaction returns the registered reaction associated with the given id.
// If there is no registered reaction with the given id the function will return an empty registered reaction and false.
func (k Keeper) GetRegisteredReaction(ctx sdk.Context, subspaceID uint64, registeredReactionID uint32) (registeredReaction types.RegisteredReaction, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.RegisteredReactionStoreKey(subspaceID, registeredReactionID))
	if bz == nil {
		return types.RegisteredReaction{}, false
	}

	k.cdc.MustUnmarshal(bz, &registeredReaction)
	return registeredReaction, true
}

// DeleteRegisteredReaction deletes the registered reaction having the given id from the store
func (k Keeper) DeleteRegisteredReaction(ctx sdk.Context, subspaceID uint64, registeredReactionID uint32) {
	registeredReaction, found := k.GetRegisteredReaction(ctx, subspaceID, registeredReactionID)
	if !found {
		return
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete(types.RegisteredReactionStoreKey(subspaceID, registeredReactionID))

	// Delete all the reactions associated to this registered reaction
	k.IterateSubspaceReactions(ctx, subspaceID, func(reaction types.Reaction) (stop bool) {
		if registeredReactionValue, ok := reaction.Value.GetCachedValue().(*types.RegisteredReactionValue); ok {
			if registeredReactionValue.RegisteredReactionID == registeredReaction.ID {
				k.DeleteReaction(ctx, reaction.SubspaceID, reaction.ID)
			}
		}
		return false
	})

	k.Logger(ctx).Debug("registered reaction deleted", "subspace id", subspaceID, "id", registeredReactionID)
	k.AfterRegisteredReactionDeleted(ctx, subspaceID, registeredReactionID)
}
