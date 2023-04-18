package keeper

import (
	errors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/posts/types"
)

// SetNextAttachmentID sets the new attachment id for the given post to the store
func (k Keeper) SetNextAttachmentID(ctx sdk.Context, subspaceID uint64, postID uint64, attachmentID uint32) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.NextAttachmentIDStoreKey(subspaceID, postID), types.GetAttachmentIDBytes(attachmentID))
}

// HasNextAttachmentID checks whether the given post already has an attachment id
func (k Keeper) HasNextAttachmentID(ctx sdk.Context, subspaceID uint64, postID uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.NextAttachmentIDStoreKey(subspaceID, postID))
}

// GetNextAttachmentID gets the highest attachment id for the given post
func (k Keeper) GetNextAttachmentID(ctx sdk.Context, subspaceID uint64, postID uint64) (attachmentID uint32, err error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NextAttachmentIDStoreKey(subspaceID, postID))
	if bz == nil {
		return 0, errors.Wrapf(types.ErrInvalidGenesis, "initial attachment ID hasn't been set for post %d within subspace %d", postID, subspaceID)
	}

	attachmentID = types.GetAttachmentIDFromBytes(bz)
	return attachmentID, nil
}

// DeleteNextAttachmentID deletes the store key used to store the next attachment id for the post having the given id
func (k Keeper) DeleteNextAttachmentID(ctx sdk.Context, subspaceID uint64, postID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.NextAttachmentIDStoreKey(subspaceID, postID))
}

// --------------------------------------------------------------------------------------------------------------------

// SaveAttachment saves the given attachment inside the current context
func (k Keeper) SaveAttachment(ctx sdk.Context, attachment types.Attachment) {
	store := ctx.KVStore(k.storeKey)

	// Store the attachment
	store.Set(types.AttachmentStoreKey(attachment.SubspaceID, attachment.PostID, attachment.ID), k.cdc.MustMarshal(&attachment))

	k.AfterAttachmentSaved(ctx, attachment.SubspaceID, attachment.PostID, attachment.ID)
}

// HasAttachment tells whether the given attachment exists or not
func (k Keeper) HasAttachment(ctx sdk.Context, subspaceID uint64, postID uint64, attachmentID uint32) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.AttachmentStoreKey(subspaceID, postID, attachmentID))
}

// GetAttachment returns the attachment associated with the given id.
// If there is no attachment associated with the given id the function will return an empty attachment and false.
func (k Keeper) GetAttachment(ctx sdk.Context, subspaceID uint64, postID uint64, attachmentID uint32) (attachment types.Attachment, found bool) {
	store := ctx.KVStore(k.storeKey)

	key := types.AttachmentStoreKey(subspaceID, postID, attachmentID)
	if !store.Has(key) {
		return types.Attachment{}, false
	}

	k.cdc.MustUnmarshal(store.Get(key), &attachment)
	return attachment, true
}

// DeleteAttachment deletes the given attachment from the current context
func (k Keeper) DeleteAttachment(ctx sdk.Context, subspaceID uint64, postID uint64, attachmentID uint32) {
	store := ctx.KVStore(k.storeKey)

	attachment, found := k.GetAttachment(ctx, subspaceID, postID, attachmentID)
	if !found {
		return
	}

	// Delete the attachment
	store.Delete(types.AttachmentStoreKey(attachment.SubspaceID, attachment.PostID, attachment.ID))

	// Remove the poll from the active queue
	if types.IsPoll(attachment) {
		// Remove the poll from the active queue (if it was there)
		k.RemoveFromActivePollQueue(ctx, attachment)

		// Delete the poll user answers
		for _, answer := range k.GetPollUserAnswers(ctx, attachment.SubspaceID, attachment.PostID, attachment.ID) {
			k.DeleteUserAnswer(ctx, attachment.SubspaceID, attachment.PostID, attachment.ID, answer.User)
		}
	}

	k.AfterAttachmentDeleted(ctx, subspaceID, postID, attachmentID)
}
