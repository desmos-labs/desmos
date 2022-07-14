package v4

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/posts/types"
)

// MigrateStore performs the migration from version 3 to version 4 of the store.
// To do this, it iterates over all the polls, and removes from the store the user answers for
// polls that are already ended (and thus should have not accepted new answers).
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	return removeInvalidPollAnswers(store, cdc)
}

// removeInvalidPollAnswers iterates over all the polls and deletes the user answers for all the polls which
// final results already been tallied. This is to delete all the answers that have been added after that.
func removeInvalidPollAnswers(store sdk.KVStore, cdc codec.BinaryCodec) error {
	attachmentStore := prefix.NewStore(store, types.AttachmentPrefix)
	attachmentsIterator := attachmentStore.Iterator(nil, nil)
	defer attachmentsIterator.Close()

	for ; attachmentsIterator.Valid(); attachmentsIterator.Next() {
		// Get the attachment
		var attachment types.Attachment
		err := cdc.Unmarshal(attachmentsIterator.Value(), &attachment)
		if err != nil {
			return err
		}

		// Check if the attachment represents a poll and the final results have already been tallied
		if poll, ok := attachment.Content.GetCachedValue().(*types.Poll); ok && poll.FinalTallyResults != nil {
			// Remove all the answers that might still be there
			err = removePollAnswers(store, attachment.SubspaceID, attachment.PostID, attachment.ID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// removePollAnswers removes from the store the answers related to the given poll
func removePollAnswers(store sdk.KVStore, subspaceID uint64, postID uint64, pollID uint32) error {
	answersStore := prefix.NewStore(store, types.PollAnswersPrefix(subspaceID, postID, pollID))
	answersIterator := answersStore.Iterator(nil, nil)

	// Get the answers
	var keys [][]byte
	for ; answersIterator.Valid(); answersIterator.Next() {
		user := string(answersIterator.Key())
		keys = append(keys, types.PollAnswerStoreKey(subspaceID, postID, pollID, user))
	}

	// Close the iterator to avoid any conflict
	err := answersIterator.Close()
	if err != nil {
		return err
	}

	// Delete the various answers
	for _, key := range keys {
		store.Delete(key)
	}

	return nil
}
