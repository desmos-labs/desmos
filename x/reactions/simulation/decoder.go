package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/desmos-labs/desmos/v3/x/reactions/types"
)

// NewDecodeStore returns a new decoder that unmarshals the KVPair's Value
// to the corresponding reactions type
func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.HasPrefix(kvA.Key, types.NextRegisteredReactionIDPrefix):
			var nextRegisteredReactionIDA, nextRegisteredReactionIDB uint32
			nextRegisteredReactionIDA = types.GetRegisteredReactionIDFromBytes(kvA.Value)
			nextRegisteredReactionIDB = types.GetRegisteredReactionIDFromBytes(kvB.Value)
			return fmt.Sprintf("NextRegisteredReactionIDA: %d\nNextRegisteredReactionIDB: %d\n",
				nextRegisteredReactionIDA, nextRegisteredReactionIDB)

		case bytes.HasPrefix(kvA.Key, types.RegisteredReactionPrefix):
			var reactionA, reactionB types.RegisteredReaction
			cdc.MustUnmarshal(kvA.Value, &reactionA)
			cdc.MustUnmarshal(kvB.Value, &reactionB)
			return fmt.Sprintf("RegisteredReactionA: %s\nRegisteredReactionB: %s\n",
				&reactionA, &reactionB)

		case bytes.HasPrefix(kvA.Key, types.NextReactionIDPrefix):
			var nextReactionIDA, nextReactionIDB uint32
			nextReactionIDA = types.GetReactionIDFromBytes(kvA.Value)
			nextReactionIDB = types.GetReactionIDFromBytes(kvB.Value)
			return fmt.Sprintf("NextReactionIDA: %d\nNextReactionIDB: %d\n",
				nextReactionIDA, nextReactionIDB)

		case bytes.HasPrefix(kvA.Key, types.ReactionPrefix):
			var reactionA, reactionB types.Reaction
			cdc.MustUnmarshal(kvA.Value, &reactionA)
			cdc.MustUnmarshal(kvB.Value, &reactionB)
			return fmt.Sprintf("ReactionA: %s\nReactionB: %s\n",
				&reactionA, &reactionB)

		case bytes.HasPrefix(kvA.Key, types.ReactionsParamsPrefix):
			var paramsA, paramsB types.SubspaceReactionsParams
			cdc.MustUnmarshal(kvA.Value, &paramsA)
			cdc.MustUnmarshal(kvB.Value, &paramsB)
			return fmt.Sprintf("SubspaceParamsA: %s\nSubspaceParamsB: %s\n",
				&paramsA, &paramsB)

		default:
			panic(fmt.Sprintf("unexpected %s key %X (%s)", types.ModuleName, kvA.Key, kvA.Key))
		}
	}
}
