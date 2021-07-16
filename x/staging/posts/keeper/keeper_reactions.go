package keeper

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/staging/posts/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SavePostReaction allows to save the given reaction inside the store.
// It assumes that the given reaction is valid and already registered.
// If another reaction from the same user for the same post and with the same value exists, returns an expError.
func (k Keeper) SavePostReaction(ctx sdk.Context, reaction types.PostReaction) error {
	store := ctx.KVStore(k.storeKey)

	key := types.PostReactionsStoreKey(reaction.PostID, reaction.Owner, reaction.ShortCode)
	// Check for double reactions
	if store.Has(key) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"%s has already reacted with %s to the post with id %s",
			reaction.Owner, reaction.ShortCode, reaction.PostID)
	}

	// Save the new reaction
	store.Set(key, types.MustMarshalPostReaction(k.cdc, reaction))
	return nil
}

// DeletePostReaction removes the reaction from the given data.
// If no reaction with the same value was previously added from the given user, an expError
// is returned.
func (k Keeper) DeletePostReaction(ctx sdk.Context, reaction types.PostReaction) error {
	store := ctx.KVStore(k.storeKey)

	key := types.PostReactionsStoreKey(reaction.PostID, reaction.Owner, reaction.ShortCode)
	if !store.Has(key) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"cannot remove the reaction with value %s from user %s as it does not exist",
			reaction.ShortCode, reaction.Owner)
	}

	store.Delete(key)
	return nil
}

// GetAllRegisteredReactions returns all the post reactions
func (k Keeper) GetAllPostReactions(ctx sdk.Context) []types.PostReaction {
	var reactions []types.PostReaction
	k.IteratePostReactions(ctx, func(_ int64, reaction types.PostReaction) bool {
		reactions = append(reactions, reaction)
		return false
	})
	return reactions
}

// GetPostReaction returns the post reaction for the given postID, owner and short code.
// If such reaction does not exist, returns false instead.
func (k Keeper) GetPostReaction(ctx sdk.Context, postID, owner, shortCode string) (types.PostReaction, bool) {
	store := ctx.KVStore(k.storeKey)

	key := types.PostReactionsStoreKey(postID, owner, shortCode)
	if !store.Has(key) {
		return types.PostReaction{}, false
	}
	return types.MustUnmarshalPostReaction(k.cdc, store.Get(key)), true
}

// ___________________________________________________________________________________________________________________

// SaveRegisteredReaction allows to register a new reaction for later reference
func (k Keeper) SaveRegisteredReaction(ctx sdk.Context, reaction types.RegisteredReaction) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.RegisteredReactionsStoreKey(reaction.Subspace, reaction.ShortCode), k.cdc.MustMarshalBinaryBare(&reaction))

	k.Logger(ctx).Info("registered reaction", "shortcode", reaction.ShortCode, "subspace", reaction.Subspace)
}

// GetRegisteredReaction returns the registered reactions which has the given shortcode
// and is registered to be used inside the given subspace.
// If no reaction could be found, returns false instead.
func (k Keeper) GetRegisteredReaction(
	ctx sdk.Context, shortcode string, subspace string,
) (reaction types.RegisteredReaction, exist bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.RegisteredReactionsStoreKey(subspace, shortcode)

	if !store.Has(key) {
		return types.RegisteredReaction{}, false
	}

	k.cdc.MustUnmarshalBinaryBare(store.Get(key), &reaction)
	return reaction, true
}

// GetRegisteredReactions returns all the registered reactions
func (k Keeper) GetRegisteredReactions(ctx sdk.Context) []types.RegisteredReaction {
	var reactions []types.RegisteredReaction

	k.IterateRegisteredReactions(ctx, func(_ int64, reaction types.RegisteredReaction) bool {
		reactions = append(reactions, reaction)
		return false
	})

	return reactions
}
