package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/tendermint/tendermint/libs/kv"
)

// DecodeStore unmarshals the KVPair's Value to the corresponding posts type
func DecodeStore(cdc *codec.Codec, kvA, kvB kv.Pair) string {
	switch {
	case bytes.Equal(kvA.Key, types.LastPostIDStoreKey):
		var idA, idB types.PostID
		cdc.MustUnmarshalBinaryBare(kvA.Value, &idA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &idB)
		return fmt.Sprintf("LastPostIDA: %s\nLastPostIDB: %s\n", idA, idB)
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
		var reactionsA, reactionB types.Reactions
		cdc.MustUnmarshalBinaryBare(kvA.Value, &reactionsA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &reactionB)
		return fmt.Sprintf("ReactionsA: %s\nReactionsB: %s\n", reactionsA, reactionB)
	default:
		panic(fmt.Sprintf("invalid account key %X", kvA.Key))
	}
}
