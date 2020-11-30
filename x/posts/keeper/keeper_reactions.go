package keeper

import (
	"bytes"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/posts/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SavePostReaction allows to save the given reaction inside the store.
// It assumes that the given reaction is valid and already registered.
// If another reaction from the same user for the same post and with the same value exists, returns an expError.
func (k Keeper) SavePostReaction(ctx sdk.Context, postID string, reaction types.PostReaction) error {
	store := ctx.KVStore(k.storeKey)
	key := types.PostReactionsStoreKey(postID)

	// Get the existent reactions
	var reactions types.PostReactions
	k.cdc.MustUnmarshalBinaryBare(store.Get(key), &reactions)

	// Check for double reactions
	if reactions.ContainsReactionFrom(reaction.Owner, reaction.ShortCode) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"%s has already reacted with %s to the post with id %s",
			reaction.Owner, reaction.ShortCode, postID)
	}

	// Save the new reaction
	store.Set(key, k.cdc.MustMarshalBinaryBare(&types.PostReactions{
		Reactions: append(reactions.Reactions, reaction),
	}))

	return nil
}

// DeletePostReaction removes the reaction from the given user from the post having the
// given postID. If no reaction with the same value was previously added from the given user, an expError
// is returned.
func (k Keeper) DeletePostReaction(ctx sdk.Context, postID string, reaction types.PostReaction) error {
	store := ctx.KVStore(k.storeKey)
	key := types.PostReactionsStoreKey(postID)

	// Get the existing reactions
	var wrapped types.PostReactions
	k.cdc.MustUnmarshalBinaryBare(store.Get(key), &wrapped)

	// Check if the user exists
	reactions := types.NewPostReactions(wrapped.Reactions...)
	if !reactions.ContainsReactionFrom(reaction.Owner, reaction.ShortCode) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"cannot remove the reaction with value %s from user %s as it does not exist",
			reaction.ShortCode, reaction.Owner)
	}

	// Remove and save the reactions list
	if appendedReactions, edited := reactions.RemoveReaction(reaction.Owner, reaction.ShortCode); edited {
		if len(appendedReactions.Reactions) == 0 {
			store.Delete(key)
		} else {
			store.Set(key, k.cdc.MustMarshalBinaryBare(&appendedReactions))
		}
	}

	return nil
}

// GetPostReactions returns the list of reactions that has been associated to the post having the given id
func (k Keeper) GetPostReactions(ctx sdk.Context, postID string) []types.PostReaction {
	store := ctx.KVStore(k.storeKey)

	var reactions types.PostReactions
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(postID)), &reactions)

	return reactions.Reactions
}

// GetPostReactionsEntries allows to returns the list of reactions that have been stored inside the given context
func (k Keeper) GetPostReactionsEntries(ctx sdk.Context) []types.PostReactionsEntry {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.PostReactionsStorePrefix)
	defer iterator.Close()

	var reactionsData []types.PostReactionsEntry
	for ; iterator.Valid(); iterator.Next() {
		var wrapped types.PostReactions
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &wrapped)
		idBytes := bytes.TrimPrefix(iterator.Key(), types.PostReactionsStorePrefix)
		reactionsData = append(reactionsData, types.NewPostReactionsEntry(string(idBytes), wrapped.Reactions))
	}

	return reactionsData
}

// ___________________________________________________________________________________________________________________

// SaveRegisteredReaction allows to register a new reaction for later reference
func (k Keeper) SaveRegisteredReaction(ctx sdk.Context, reaction types.RegisteredReaction) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ReactionsStoreKey(reaction.ShortCode, reaction.Subspace), k.cdc.MustMarshalBinaryBare(&reaction))
}

// GetRegisteredReaction returns the registered reactions which has the given shortcode
// and is registered to be used inside the given subspace.
// If no reaction could be found, returns false instead.
func (k Keeper) GetRegisteredReaction(
	ctx sdk.Context, shortcode string, subspace string,
) (reaction types.RegisteredReaction, exist bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.ReactionsStoreKey(shortcode, subspace)

	if !store.Has(key) {
		return types.RegisteredReaction{}, false
	}

	k.cdc.MustUnmarshalBinaryBare(store.Get(key), &reaction)
	return reaction, true
}

// GetRegisteredReactions returns all the registered reactions
func (k Keeper) GetRegisteredReactions(ctx sdk.Context) (reactions []types.RegisteredReaction) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ReactionsStorePrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var reaction types.RegisteredReaction
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &reaction)
		reactions = append(reactions, reaction)
	}

	return reactions
}
