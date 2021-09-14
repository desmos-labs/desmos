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
func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.HasPrefix(kvA.Key, types.PostStorePrefix):
			var postA, postB types.Post
			cdc.MustUnmarshal(kvA.Value, &postA)
			cdc.MustUnmarshal(kvB.Value, &postB)
			return fmt.Sprintf("PostA: %s\nPostB: %s\n", postA.String(), postB.String())

		case bytes.HasPrefix(kvA.Key, types.PostReactionsStorePrefix):
			var reactionA, reactionB types.PostReaction
			cdc.MustUnmarshal(kvA.Value, &reactionA)
			cdc.MustUnmarshal(kvB.Value, &reactionB)
			return fmt.Sprintf("PostReactionA: %s\nPostReactionB: %s\n", reactionA, reactionB)

		case bytes.HasPrefix(kvA.Key, types.CommentsStorePrefix):
			var commentA, commentB string
			commentA = string(kvA.Value)
			commentB = string(kvA.Value)
			return fmt.Sprintf("CommentA: %s\nCommentB: %s\n", commentA, commentB)

		case bytes.HasPrefix(kvA.Key, types.RegisteredReactionsStorePrefix):
			var reactionA, reactionB types.RegisteredReaction
			cdc.MustUnmarshal(kvA.Value, &reactionA)
			cdc.MustUnmarshal(kvB.Value, &reactionB)
			return fmt.Sprintf("RegisteredReactionA: %s\nRegisteredReactionB: %s\n", reactionA, reactionB)

		case bytes.HasPrefix(kvA.Key, types.ReportsStorePrefix):
			var reportA, reportB types.Report
			cdc.MustUnmarshal(kvA.Value, &reportA)
			cdc.MustUnmarshal(kvB.Value, &reportB)
			return fmt.Sprintf("ReportsA: %s\nReportsB: %s\n", reportA, reportB)
		default:
			panic(fmt.Sprintf("unexpected %s key %X (%s)", types.ModuleName, kvA.Key, kvA.Key))
		}
	}
}
