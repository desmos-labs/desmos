package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/desmos-labs/desmos/v7/x/posts/types"
)

// NewDecodeStore returns a new decoder that unmarshals the KVPair's Value
// to the corresponding subspaces type
func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.HasPrefix(kvA.Key, types.NextPostIDPrefix):
			var idA, idB uint64
			idA = types.GetPostIDFromBytes(kvA.Value)
			idB = types.GetPostIDFromBytes(kvB.Value)
			return fmt.Sprintf("PostIDA: %d\nPostIDB: %d\n", idA, idB)

		case bytes.HasPrefix(kvA.Key, types.PostPrefix):
			var postA, postB types.Post
			cdc.MustUnmarshal(kvA.Value, &postA)
			cdc.MustUnmarshal(kvB.Value, &postB)
			return fmt.Sprintf("PostA: %s\nPostB: %s\n", &postA, &postB)

		case bytes.HasPrefix(kvA.Key, types.NextAttachmentIDPrefix):
			var idA, idB uint32
			idA = types.GetAttachmentIDFromBytes(kvA.Value)
			idB = types.GetAttachmentIDFromBytes(kvB.Value)
			return fmt.Sprintf("AttachmentIDA: %d\nAttachmentIDB: %d\n", idA, idB)

		case bytes.HasPrefix(kvA.Key, types.AttachmentPrefix):
			var attachmentA, attachmentB types.Attachment
			cdc.MustUnmarshal(kvA.Value, &attachmentA)
			cdc.MustUnmarshal(kvB.Value, &attachmentB)
			return fmt.Sprintf("AttachmentA: %s\nAttachmentB: %s\n", &attachmentA, &attachmentB)

		case bytes.HasPrefix(kvA.Key, types.UserAnswerPrefix):
			var answerA, answerB types.UserAnswer
			cdc.MustUnmarshal(kvA.Value, &answerA)
			cdc.MustUnmarshal(kvB.Value, &answerB)
			return fmt.Sprintf("UserAnswerA: %s\nUserAnswerB: %s\n", &answerA, &answerB)

		case bytes.HasPrefix(kvA.Key, types.ActivePollQueuePrefix):
			subspaceIDA, postIDA, pollIDA := types.GetPollIDFromBytes(kvA.Value)
			subspaceIDB, postIDB, pollIDB := types.GetPollIDFromBytes(kvB.Value)
			return fmt.Sprintf("SubspaceIDA: %d, PostIDA: %d, PollIDA: %d\nSubspaceIDB: %d, PostIDB: %d, PollIDB: %d\n",
				subspaceIDA, postIDA, pollIDA, subspaceIDB, postIDB, pollIDB)

		default:
			panic(fmt.Sprintf("unexpected %s key %X (%s)", types.ModuleName, kvA.Key, kvA.Key))
		}
	}
}
