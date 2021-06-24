package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

// NewDecodeStore returns a new decoder that unmarshals the KVPair's Value
// to the corresponding posts type
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
			var postReactionsA, postReactionsB types.PostReaction
			cdc.MustUnmarshalBinaryBare(kvA.Value, &postReactionsA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &postReactionsB)
			return fmt.Sprintf("PostReactionsA: %s\nPostReactionsB: %s\n", postReactionsA, postReactionsB)

		case bytes.HasPrefix(kvA.Key, types.RegisteredReactionsStorePrefix):
			var reactionA, reactionB types.RegisteredReaction
			cdc.MustUnmarshalBinaryBare(kvA.Value, &reactionA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &reactionB)
			return fmt.Sprintf("ReactionA: %s\nReactionB: %s\n", reactionA, reactionB)

		case bytes.HasPrefix(kvA.Key, types.ReportsStorePrefix):
			var reportsA, reportsB types.Reports
			cdc.MustUnmarshalBinaryBare(kvA.Value, &reportsA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &reportsB)
			return fmt.Sprintf("ReportsA: %s\nReportsB: %s\n", reportsA.Reports, reportsB.Reports)
		default:
			panic(fmt.Sprintf("unexpected %s key %X (%s)", types.ModuleName, kvA.Key, kvA.Key))
		}
	}
}
