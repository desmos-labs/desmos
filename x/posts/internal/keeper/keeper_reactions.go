package keeper

import (
	"bytes"
	"fmt"

	"github.com/desmos-labs/desmos/x/posts/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// -------------
// --- PostReactions
// -------------

// SavePostReaction allows to save the given postReaction inside the store.
// It assumes that the given postReaction is valid.
// If another postReaction from the same user for the same post and with the same value exists, returns an expError.
// nolint: interfacer
func (k Keeper) SavePostReaction(ctx sdk.Context, postID types.PostID, reaction types.PostReaction) error {
	store := ctx.KVStore(k.StoreKey)
	key := types.PostReactionsStoreKey(postID)

	// Get the existent reactions
	var reactions types.PostReactions
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &reactions)

	// Check for double reactions
	if reactions.ContainsReactionFrom(reaction.Owner, reaction.Value) {
		return fmt.Errorf("%s has already reacted with %s to the post with id %s",
			reaction.Owner, reaction.Value, postID)
	}

	// Check if the postReaction is a registered one
	post, _ := k.GetPost(ctx, postID)
	if _, exist := k.GetRegisteredReaction(ctx, reaction.Value, post.Subspace); !exist {
		return fmt.Errorf("postReaction with short code %s isn't registered yet and can't be used to react to the post "+
			"with ID %s and subspace %s, please register it before use", reaction.Value, postID, post.Subspace)
	}

	// Save the new postReaction
	reactions = append(reactions, reaction)
	store.Set(key, k.Cdc.MustMarshalBinaryBare(&reactions))

	return nil
}

// RemovePostReaction removes the postReaction from the given user from the post having the
// given postID. If no postReaction with the same value was previously added from the given user, an expError
// is returned.
// nolint: interfacer
func (k Keeper) RemovePostReaction(ctx sdk.Context, postID types.PostID, reaction types.PostReaction) error {
	store := ctx.KVStore(k.StoreKey)
	key := types.PostReactionsStoreKey(postID)

	// Get the existing reactions
	var reactions types.PostReactions
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &reactions)

	// Check if the user exists
	if !reactions.ContainsReactionFrom(reaction.Owner, reaction.Value) {
		return fmt.Errorf("cannot remove the postReaction with value %s from user %s as it does not exist", reaction.Value, reaction.Owner)
	}

	// Remove and save the reactions list
	if newLikes, edited := reactions.RemoveReaction(reaction.Owner, reaction.Value); edited {
		if len(newLikes) == 0 {
			store.Delete(key)
		} else {
			store.Set(key, k.Cdc.MustMarshalBinaryBare(&newLikes))
		}
	}

	return nil
}

// GetPostReactions returns the list of reactions that has been associated to the post having the given id
// nolint: interfacer
func (k Keeper) GetPostReactions(ctx sdk.Context, postID types.PostID) types.PostReactions {
	store := ctx.KVStore(k.StoreKey)

	var reactions types.PostReactions
	k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(postID)), &reactions)

	return reactions
}

// GetReactions allows to returns the list of reactions that have been stored inside the given context
func (k Keeper) GetReactions(ctx sdk.Context) map[string]types.PostReactions {
	store := ctx.KVStore(k.StoreKey)

	iterator := sdk.KVStorePrefixIterator(store, types.PostReactionsStorePrefix)
	defer iterator.Close()

	reactionsData := map[string]types.PostReactions{}
	for ; iterator.Valid(); iterator.Next() {
		var postLikes types.PostReactions
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &postLikes)
		idBytes := bytes.TrimPrefix(iterator.Key(), types.PostReactionsStorePrefix)
		reactionsData[string(idBytes)] = postLikes
	}

	return reactionsData
}

// -------------
// --- Reactions
// -------------

// RegisterReaction allows to register a new postReaction for later reference
func (k Keeper) RegisterReaction(ctx sdk.Context, reaction types.Reaction) {
	store := ctx.KVStore(k.StoreKey)
	store.Set(types.ReactionsStoreKey(reaction.ShortCode, reaction.Subspace), k.Cdc.MustMarshalBinaryBare(reaction))
}

// GetRegisteredReaction returns the registered postReaction which has the given shortcode
// and is registered to be used inside the given subspace.
// If no postReaction could be found, returns false instead.
func (k Keeper) GetRegisteredReaction(
	ctx sdk.Context, shortcode string, subspace string,
) (reaction types.Reaction, exist bool) {
	store := ctx.KVStore(k.StoreKey)
	key := types.ReactionsStoreKey(shortcode, subspace)

	if !store.Has(key) {
		return types.Reaction{}, false
	}

	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &reaction)
	return reaction, true
}

// GetRegisteredReactions returns all the registered reactions
func (k Keeper) GetRegisteredReactions(ctx sdk.Context) (reactions types.Reactions) {
	store := ctx.KVStore(k.StoreKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ReactionsStorePrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var reaction types.Reaction
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &reaction)
		reactions = append(reactions, reaction)
	}

	return reactions
}
