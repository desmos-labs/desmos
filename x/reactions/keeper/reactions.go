package keeper

import (
	"regexp"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v6/x/reactions/types"
)

// SetNextReactionID sets the next reaction id for the given subspace
func (k Keeper) SetNextReactionID(ctx sdk.Context, subspaceID uint64, postID uint64, reactionID uint32) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.NextReactionIDStoreKey(subspaceID, postID), types.GetReactionIDBytes(reactionID))
}

// HasNextReactionID tells whether the next reaction id exists for the given subspace
func (k Keeper) HasNextReactionID(ctx sdk.Context, subspaceID uint64, postID uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.NextReactionIDStoreKey(subspaceID, postID))
}

// GetNextReactionID gets the next reaction id for the given subspace
func (k Keeper) GetNextReactionID(ctx sdk.Context, subspaceID uint64, postID uint64) (reactionID uint32, err error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NextReactionIDStoreKey(subspaceID, postID))
	if bz == nil {
		return 0, errors.Wrapf(types.ErrInvalidGenesis, "initial reaction id not set for post %d inside subspace %d", postID, subspaceID)
	}

	reactionID = types.GetReactionIDFromBytes(bz)
	return reactionID, nil
}

// DeleteNextReactionID removes the next reaction id for the given subspace
func (k Keeper) DeleteNextReactionID(ctx sdk.Context, subspaceID uint64, postID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.NextReactionIDStoreKey(subspaceID, postID))
}

// --------------------------------------------------------------------------------------------------------------------

// validateRegisteredReactionValue validates the given reaction containing the provided registered reaction value
func (k Keeper) validateRegisteredReactionValue(ctx sdk.Context, reaction types.Reaction, value *types.RegisteredReactionValue) error {
	reactionsParams, err := k.GetSubspaceReactionsParams(ctx, reaction.SubspaceID)
	if err != nil {
		return err
	}
	params := reactionsParams.RegisteredReaction

	// Make sure registered reactions are enabled
	if !params.Enabled {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "registered reactions are not enabled inside this subspace")
	}

	// Make sure the registered reaction exists
	if !k.HasRegisteredReaction(ctx, reaction.SubspaceID, value.RegisteredReactionID) {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "registered reaction with id %d not found", value.RegisteredReactionID)
	}

	return nil
}

// validateFreeTextValue validates the given reaction containing the provided free text value
func (k Keeper) validateFreeTextValue(ctx sdk.Context, reaction types.Reaction, value *types.FreeTextValue) error {
	reactionsParams, err := k.GetSubspaceReactionsParams(ctx, reaction.SubspaceID)
	if err != nil {
		return err
	}
	params := reactionsParams.FreeText

	// Make sure the free text reactions are enabled
	if !params.Enabled {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "free text reactions are not enabled inside this subspace")
	}

	// Make sure the value respected the max length
	if uint32(len(value.Text)) > params.MaxLength {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "value exceed max length allowed")
	}

	// Make sure the value matches the regex
	if params.RegEx != "" {
		regEx, err := regexp.Compile(params.RegEx)
		if err != nil {
			return errors.Wrapf(sdkerrors.ErrLogic, "invalid free text regex")
		}

		if !regEx.MatchString(value.Text) {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "value does not respect required regex")
		}
	}

	return nil
}

// ValidateReaction validates the given reaction's value
func (k Keeper) ValidateReaction(ctx sdk.Context, reaction types.Reaction) (err error) {
	// Validate the reaction
	err = reaction.Validate()
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Validate the value
	switch value := reaction.Value.GetCachedValue().(type) {
	case *types.RegisteredReactionValue:
		err = k.validateRegisteredReactionValue(ctx, reaction, value)
	case *types.FreeTextValue:
		err = k.validateFreeTextValue(ctx, reaction, value)
	}

	return err
}

// SaveReaction saves the given reaction inside the current context
func (k Keeper) SaveReaction(ctx sdk.Context, reaction types.Reaction) {
	store := ctx.KVStore(k.storeKey)

	// Store the reaction
	store.Set(types.ReactionStoreKey(reaction.SubspaceID, reaction.PostID, reaction.ID), k.cdc.MustMarshal(&reaction))

	k.AfterReactionSaved(ctx, reaction.SubspaceID, reaction.ID)
}

// HasReaction tells whether the given reaction exists or not
func (k Keeper) HasReaction(ctx sdk.Context, subspaceID uint64, postID uint64, reactionID uint32) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.ReactionStoreKey(subspaceID, postID, reactionID))
}

// GetReaction returns the reaction associated with the given id.
// If there is no reaction with the given id the function will return an empty reaction and false.
func (k Keeper) GetReaction(ctx sdk.Context, subspaceID uint64, postID uint64, reactionID uint32) (reaction types.Reaction, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ReactionStoreKey(subspaceID, postID, reactionID))
	if bz == nil {
		return types.Reaction{}, false
	}

	k.cdc.MustUnmarshal(bz, &reaction)
	return reaction, true
}

// DeleteReaction deletes the reaction having the given id from the store
func (k Keeper) DeleteReaction(ctx sdk.Context, subspaceID uint64, postID uint64, reactionID uint32) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.ReactionStoreKey(subspaceID, postID, reactionID))

	k.AfterReactionDeleted(ctx, subspaceID, reactionID)
}
