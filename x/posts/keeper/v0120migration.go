package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v0100 "github.com/desmos-labs/desmos/x/posts/keeper/legacy/v0.10.0"
	"github.com/desmos-labs/desmos/x/posts/types"
)

// MigratePostsFrom0100To0120 migrates all the posts from v0.10.0 to v.12.0.
// To do this it executes the following operations one post at a time:
// 1. It reads the old post
// 2. It converts the post removing the Open field from the PollData, if any
// 3. It saves the post inside the store again
func (k Keeper) MigratePostsFrom0100To0120(ctx sdk.Context) error {
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PostStorePrefix)

	// Get all the keys
	var keys [][]byte
	for ; iterator.Valid(); iterator.Next() {
		keys = append(keys, iterator.Key())
	}

	// Check iteration errors
	err := iterator.Error()
	if err != nil {
		return err
	}

	// Close the iterator
	iterator.Close()

	// Iterate over all the keys and migrate the data
	for _, key := range keys {
		// Get the v0.10.0 post
		var v0100Post v0100.Post
		err := k.Cdc.UnmarshalBinaryBare(store.Get(key), &v0100Post)
		if err != nil {
			return err
		}

		// Convert the post
		v0120Post := types.Post{
			PostID:         types.PostID(v0100Post.PostID),
			ParentID:       types.PostID(v0100Post.ParentID),
			Message:        v0100Post.Message,
			Created:        v0100Post.Created,
			LastEdited:     v0100Post.LastEdited,
			AllowsComments: v0100Post.AllowsComments,
			Subspace:       v0100Post.Subspace,
			OptionalData:   types.OptionalData(v0100Post.OptionalData),
			Creator:        v0100Post.Creator,
			Attachments:    migrateAttachments(v0100Post.Attachments),
			PollData:       migratePollData(v0100Post.PollData),
		}

		bz, err := k.Cdc.MarshalBinaryBare(&v0120Post)
		if err != nil {
			return err
		}

		// Store the post
		store.Set(key, bz)
	}

	return nil
}

// migrateAttachments migrates the given attachments from v0.10.0 to v0.12.0
func migrateAttachments(attachments []v0100.Attachment) types.Attachments {
	var v1200Attachments = make([]types.Attachment, len(attachments))
	for index, attachment := range attachments {
		v1200Attachments[index] = types.Attachment(attachment)
	}
	return v1200Attachments
}

// migratePollData migrates the given pollData from v0.10.0 to v0.12.0
func migratePollData(pollData *v0100.PollData) *types.PollData {
	if pollData == nil {
		return nil
	}

	return &types.PollData{
		Question:              pollData.Question,
		ProvidedAnswers:       migrateProvidedAnswers(pollData.ProvidedAnswers),
		EndDate:               pollData.EndDate,
		AllowsMultipleAnswers: pollData.AllowsMultipleAnswers,
		AllowsAnswerEdits:     pollData.AllowsAnswerEdits,
	}
}

// migrateProvidedAnswers migrates the providedAnswers from v0.10.0 to v0.12.0
func migrateProvidedAnswers(providedAnswers []v0100.PollAnswer) types.PollAnswers {
	var v0120Answers = make([]types.PollAnswer, len(providedAnswers))
	for index, answer := range providedAnswers {
		v0120Answers[index] = types.PollAnswer{
			ID:   types.AnswerID(answer.ID),
			Text: answer.Text,
		}
	}
	return v0120Answers
}
