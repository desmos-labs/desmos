package v3

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	poststypes "github.com/desmos-labs/desmos/v4/x/posts/types"
	"github.com/desmos-labs/desmos/v4/x/reactions/types"
)

// MigrateStore performs in-place store migrations from v2 to v3
// It removes all the duplicated reactions and fix missing reaction ids.
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, pk types.PostsKeeper, cdc codec.BinaryCodec) error {
	err := RemoveDuplicatedReactions(ctx, storeKey, cdc)
	if err != nil {
		return err
	}

	return FixMissingNextReactionIDs(ctx, storeKey, pk, cdc)
}

// RemoveDuplicatedReactions removes all the duplicated reactions
func RemoveDuplicatedReactions(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	reactionsStore := prefix.NewStore(store, types.ReactionPrefix)

	iter := reactionsStore.Iterator(nil, nil)
	defer iter.Close()

	// Get duplicated reactions
	var reactions []types.Reaction
	var duplicatedReactions []types.Reaction
	for ; iter.Valid(); iter.Next() {
		var reaction types.Reaction
		err := cdc.Unmarshal(iter.Value(), &reaction)
		if err != nil {
			return err
		}
		reactions = append(reactions, reaction)

		if contains(reactions, reaction) {
			duplicatedReactions = append(duplicatedReactions, reaction)
		}
	}

	// Remove duplicated reactions
	for _, reaction := range duplicatedReactions {
		store.Delete(types.ReactionStoreKey(reaction.SubspaceID, reaction.PostID, reaction.ID))
	}

	return nil
}

func contains(reactions []types.Reaction, reaction types.Reaction) bool {
	var count = 0
	for _, r := range reactions {
		if r.SubspaceID == reaction.SubspaceID &&
			r.PostID == reaction.PostID &&
			r.Author == reaction.Author &&
			r.Value.Equal(reaction.Value) {
			count++
		}
	}
	return count > 1
}

// --------------------------------------------------------------------------------------------------------------------

// FixMissingNextReactionIDs fixes the missing next reaction ids for existing posts
func FixMissingNextReactionIDs(ctx sdk.Context, storeKey storetypes.StoreKey, pk types.PostsKeeper, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	pk.IteratePosts(ctx, func(post poststypes.Post) bool {
		// Skip if the next reaction ID key exists
		if store.Has(types.NextReactionIDStoreKey(post.SubspaceID, post.ID)) {
			return false
		}

		// Get max reaction ID of the post
		iter := sdk.KVStorePrefixIterator(store, types.PostReactionsPrefix(post.SubspaceID, post.ID))
		maxReactionID := uint32(0)
		for ; iter.Valid(); iter.Next() {
			var reaction types.Reaction

			cdc.MustUnmarshal(iter.Value(), &reaction)
			if reaction.ID > maxReactionID {
				maxReactionID = reaction.ID
			}
		}
		iter.Close()

		// Set next reaction id
		store.Set(types.NextReactionIDStoreKey(post.SubspaceID, post.ID), types.GetReactionIDBytes(maxReactionID+1))

		return false
	})

	return nil
}
