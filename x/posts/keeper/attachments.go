package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
)

// SetAttachmentID sets the new attachment id for the given post to the store
func (k Keeper) SetAttachmentID(ctx sdk.Context, subspaceID uint64, postID uint64, attachmentID uint32) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.AttachmentIDStoreKey(subspaceID, postID), types.GetAttachmentIDBytes(attachmentID))
}

// HasAttachmentID checks whether the given post already has an attachment id
func (k Keeper) HasAttachmentID(ctx sdk.Context, subspaceID uint64, postID uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.AttachmentIDStoreKey(subspaceID, postID))
}

// GetAttachmentID gets the highest attachment id for the given post
func (k Keeper) GetAttachmentID(ctx sdk.Context, subspaceID uint64, postID uint64) (attachmentID uint32, err error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.AttachmentIDStoreKey(subspaceID, postID))
	if bz == nil {
		return 0, sdkerrors.Wrapf(types.ErrInvalidGenesis, "initial attachment ID hasn't been set for post %d within subspace %d", postID, subspaceID)
	}

	attachmentID = types.GetAttachmentIDFromBytes(bz)
	return attachmentID, nil
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

	// Delete the attachment
	store.Delete(types.AttachmentStoreKey(subspaceID, postID, attachmentID))

	// Delete the poll user answers
	for _, answer := range k.GetUserAnswers(ctx, subspaceID, postID, attachmentID) {
		k.DeleteUserAnswer(ctx, subspaceID, postID, attachmentID, answer.User)
	}

	// Delete the poll tally results
	k.DeletePollTallyResults(ctx, subspaceID, postID, attachmentID)
}
