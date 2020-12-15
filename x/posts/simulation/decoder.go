package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/desmos-labs/desmos/x/posts/types"
)

// NewDecodeStore returns a new decoder that unmarshals the KVPair's Value
// to the corresponding relationships type
func NewDecodeStore(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.HasPrefix(kvA.Key, types.PostStorePrefix):
			var postA, postB types.Post
			cdc.MustUnmarshalBinaryBare(kvA.Value, &postA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &postB)
			return fmt.Sprintf("PostA: %s\nPostB: %s\n", postA.String(), postB.String())

		case bytes.HasPrefix(kvA.Key, types.PostCommentsStorePrefix):
			var commentsA, commentsB types.CommentIDs
			cdc.MustUnmarshalBinaryBare(kvA.Value, &commentsA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &commentsB)
			return fmt.Sprintf("CommentsA: %s\nCommentsB: %s\n", commentsA, commentsB)

		case bytes.HasPrefix(kvA.Key, types.PostReactionsStorePrefix):
			var postReactionsA, postReactionsB types.PostReactions
			cdc.MustUnmarshalBinaryBare(kvA.Value, &postReactionsA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &postReactionsB)
			return fmt.Sprintf("PostReactionsA: %s\nPostReactionsB: %s\n", postReactionsA, postReactionsB)

		case bytes.HasPrefix(kvA.Key, types.ReactionsStorePrefix):
			var reactionA, reactionB types.RegisteredReaction
			cdc.MustUnmarshalBinaryBare(kvA.Value, &reactionA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &reactionB)
			return fmt.Sprintf("ReactionA: %s\nReactionB: %s\n", reactionA, reactionB)

		case bytes.HasPrefix(kvA.Key, types.PostIndexedIDStorePrefix):
			var indexedIDA, indexedIDB types.PostIndex
			cdc.MustUnmarshalBinaryBare(kvA.Value, &indexedIDA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &indexedIDB)
			return fmt.Sprintf("IndexedIDA: %d\nIndexedIDB: %d\n", indexedIDA.Value, indexedIDB.Value)

		case bytes.HasPrefix(kvA.Key, types.PostTotalNumberPrefix):
			var totalPostsA, totalPostsB types.PostIndex
			cdc.MustUnmarshalBinaryBare(kvA.Value, &totalPostsA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &totalPostsB)
			return fmt.Sprintf("TotalPostsA: %d\nTotalPostsB: %d\n", totalPostsA.Value, totalPostsB.Value)

		default:
			panic(fmt.Sprintf("unexpected %s key %X (%s)", types.ModuleName, kvA.Key, kvA.Key))
		}
	}
}
