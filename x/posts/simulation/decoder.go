package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/types"
	"github.com/tendermint/tendermint/libs/kv"
)

// DecodeStore unmarshals the KVPair's Reaction to the corresponding posts type
func DecodeStore(cdc *codec.Codec, kvA, kvB kv.Pair) string {
	switch {
	case bytes.HasPrefix(kvA.Key, types.PostStorePrefix):
		var postA, postB types.Post
		cdc.MustUnmarshalBinaryBare(kvA.Value, &postA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &postB)
		return fmt.Sprintf("PostA: %s\nPostB: %s\n", postA, postB)
	case bytes.HasPrefix(kvA.Key, types.PostCommentsStorePrefix):
		var commentsA, commentsB types.PostIDs
		cdc.MustUnmarshalBinaryBare(kvA.Value, &commentsA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &commentsB)
		return fmt.Sprintf("CommentsA: %s\nCommentsB: %s\n", commentsA, commentsB)
	case bytes.HasPrefix(kvA.Key, types.PostReactionsStorePrefix):
		var postReactionsA, postReactionsB types.PostReactions
		cdc.MustUnmarshalBinaryBare(kvA.Value, &postReactionsA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &postReactionsB)
		return fmt.Sprintf("PostReactionsA: %s\nPostReactionsB: %s\n", postReactionsA, postReactionsB)
	case bytes.HasPrefix(kvA.Key, types.ReactionsStorePrefix):
		var reactionA, reactionB types.Reaction
		cdc.MustUnmarshalBinaryBare(kvA.Value, &reactionA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &reactionB)
		return fmt.Sprintf("ReactionA: %s\nReactionB: %s\n", reactionA, reactionB)
	case bytes.HasPrefix(kvA.Key, types.PostIndexedIDStorePrefix):
		var indexedIDA, indexedIDB types.PostID
		cdc.MustUnmarshalBinaryBare(kvA.Value, &indexedIDA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &indexedIDB)
		return fmt.Sprintf("IndexedIDA: %s\nIndexedIDB: %s\n", indexedIDA, indexedIDB)
	case bytes.HasPrefix(kvA.Key, types.PostTotalNumberPrefix):
		var totalPostsA, totalPostsB sdk.Int
		cdc.MustUnmarshalBinaryBare(kvA.Value, &totalPostsA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &totalPostsB)
		return fmt.Sprintf("TotalPostsA: %s\nTotalPostsB: %s\n", totalPostsA, totalPostsB)
	default:
		panic(fmt.Sprintf("invalid account key %X", kvA.Key))
	}
}
