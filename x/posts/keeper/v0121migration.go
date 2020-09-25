package keeper

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"

	v0100 "github.com/desmos-labs/desmos/x/posts/keeper/legacy/v0.10.0"
	"github.com/desmos-labs/desmos/x/posts/types"
)

// PerformSeptemberFixMigration fixes all the errors made inside the september-upgrade
// migration process. To do this, it performs the following operations:
// 1. Remove all the keys that should not be there and were added during the posts migration.
// 2. Truly migrate all the posts.
func (k Keeper) PerformSeptemberFixMigration(ctx sdk.Context) error {
	store := ctx.KVStore(k.StoreKey)

	// Delete unwanted keys
	if err := k.deleteUnwatedKeys(store); err != nil {
		return err
	}

	// Migrate the posts
	if err := k.migratePosts(store); err != nil {
		return err
	}

	return nil
}

// deleteUnwatedKeys removes all the keys from the posts store that should not be there.
// These are all the keys that were created wrongly due to a bug inside the september upgrade migration.
func (k Keeper) deleteUnwatedKeys(store sdk.KVStore) error {
	iterator := store.Iterator(nil, nil)

	// Get all the keys that should be deleted
	var keysToDelete [][]byte
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		if bytes.HasPrefix(key, types.PostStorePrefix) {
			continue
		}

		if bytes.HasPrefix(key, types.PostIndexedIDStorePrefix) {
			continue
		}

		if bytes.HasPrefix(key, types.PostTotalNumberPrefix) {
			continue
		}

		if bytes.HasPrefix(key, types.PostCommentsStorePrefix) {
			continue
		}

		if bytes.HasPrefix(key, types.PostReactionsStorePrefix) {
			continue
		}

		if bytes.HasPrefix(key, types.ReactionsStorePrefix) {
			continue
		}

		if bytes.HasPrefix(key, types.PollAnswersStorePrefix) {
			continue
		}

		keysToDelete = append(keysToDelete, key)
	}

	// Close the iterator
	iterator.Close()

	// Check iteration errors
	if err := iterator.Error(); err != nil {
		return err
	}
	return nil
}

// migratePosts performs the migration of all the posts from version v0.10.0 to v0.12.0.
// To do this it reads all the keys inside the store having the proper prefix, it migrates the values,
// and then it writes all the new values associating them with the already existing keys.
func (k Keeper) migratePosts(store sdk.KVStore) error {
	iterator := sdk.KVStorePrefixIterator(store, types.PostStorePrefix)

	// Get all the keys
	var keys [][]byte
	for ; iterator.Valid(); iterator.Next() {
		keys = append(keys, iterator.Key())
	}

	// Close the iterator
	iterator.Close()

	// Check iteration errors
	if err := iterator.Error(); err != nil {
		return err
	}

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
