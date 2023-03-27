package v3

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v4/x/reactions/types"
)

// MigrateStore performs in-place store migrations from v2 to v3
// It removes all the duplicated reactions.
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
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
