package v5

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	v4 "github.com/desmos-labs/desmos/v7/x/posts/legacy/v4"
	"github.com/desmos-labs/desmos/v7/x/posts/types"
)

// MigrateStore performs the migration from version 4 to version 5 of the store.
// To do this, it iterates over all the post attachments, and converts them to
// the new storing format (AttachmentContent instead of Attachment).
// It also removes all the Polls that have been saved as a Poll_ProvidedAnswer's attachment.
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	// Migrate the poll attachments
	return migrateAttachments(store, cdc)
}

// migrateAttachments migrates all the attachments to v5
func migrateAttachments(store sdk.KVStore, cdc codec.BinaryCodec) error {
	attachmentsStore := prefix.NewStore(store, types.AttachmentPrefix)
	attachmentsIterator := attachmentsStore.Iterator(nil, nil)
	defer attachmentsIterator.Close()

	for ; attachmentsIterator.Valid(); attachmentsIterator.Next() {
		// Get the attachment
		var attachment v4.Attachment
		err := cdc.Unmarshal(attachmentsIterator.Value(), &attachment)
		if err != nil {
			return err
		}

		var attachmentContent types.AttachmentContent
		switch content := attachment.Content.GetCachedValue().(type) {
		case *v4.Media:
			attachmentContent = migrateMedia(content)
		case *v4.Poll:
			v5Poll, err := migratePoll(content, cdc)
			if err != nil {
				return err
			}
			attachmentContent = v5Poll
		}

		// Store the new attachment - This will override the old store key
		v5Attachment := types.NewAttachment(attachment.SubspaceID, attachment.PostID, attachment.ID, attachmentContent)
		store.Set(
			types.AttachmentStoreKey(v5Attachment.SubspaceID, v5Attachment.PostID, v5Attachment.ID),
			cdc.MustMarshal(&v5Attachment),
		)
	}

	return nil
}

// migrateMedia migrates the given media to the new format
func migrateMedia(media *v4.Media) *types.Media {
	return types.NewMedia(media.Uri, media.MimeType)
}

// migratePoll migrates the given poll to the new format
func migratePoll(poll *v4.Poll, cdc codec.BinaryCodec) (*types.Poll, error) {
	// Get the new provided answers
	providedAnswers := make([]types.Poll_ProvidedAnswer, len(poll.ProvidedAnswers))
	for i, answer := range poll.ProvidedAnswers {
		attachmentContents, err := migrateProvidedAnswerAttachments(answer.Attachments, cdc)
		if err != nil {
			return nil, err
		}
		providedAnswers[i] = types.NewProvidedAnswer(answer.Text, attachmentContents)
	}

	return types.NewPoll(
		poll.Question,
		providedAnswers,
		poll.EndDate,
		poll.AllowsMultipleAnswers,
		poll.AllowsAnswerEdits,
		migratePollFinalTallyResults(poll.FinalTallyResults),
	), nil
}

// migrateProvidedAnswerAttachments migrates the given attachments slide
// converting them into AttachmentContent instances. It also filters all
// the poll attachments to exclude them
func migrateProvidedAnswerAttachments(attachments []v4.Attachment, cdc codec.BinaryCodec) ([]types.AttachmentContent, error) {
	if attachments == nil {
		return nil, nil
	}

	var attachmentContents []types.AttachmentContent
	for _, attachment := range attachments {
		var content v4.AttachmentContent
		err := cdc.UnpackAny(attachment.Content, &content)
		if err != nil {
			return nil, err
		}

		// Convert the media only
		if media, isMedia := content.(*v4.Media); isMedia {
			attachmentContents = append(attachmentContents, types.NewMedia(media.Uri, media.MimeType))
		}
	}

	return attachmentContents, nil
}

// migratePollFinalTallyResults migrates the given v4 poll tally results into v5
func migratePollFinalTallyResults(results *v4.PollTallyResults) *types.PollTallyResults {
	if results == nil {
		return nil
	}

	answersResults := make([]types.PollTallyResults_AnswerResult, len(results.Results))
	for i, result := range results.Results {
		answersResults[i] = types.NewAnswerResult(result.AnswerIndex, result.Votes)
	}
	return types.NewPollTallyResults(answersResults)
}
