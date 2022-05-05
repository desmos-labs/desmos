package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
)

// SetPostID sets the new post id for the given subspace to the store
func (k Keeper) SetPostID(ctx sdk.Context, subspaceID uint64, postID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PostIDStoreKey(subspaceID), types.GetPostIDBytes(postID))
}

// GetPostID gets the highest post id for the given subspace
func (k Keeper) GetPostID(ctx sdk.Context, subspaceID uint64) (postID uint64, err error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.PostIDStoreKey(subspaceID))
	if bz == nil {
		return 0, sdkerrors.Wrapf(types.ErrInvalidGenesis, "initial post ID hasn't been set for subspace %d", subspaceID)
	}

	postID = types.GetPostIDFromBytes(bz)
	return postID, nil
}

// DeletePostID removes the post id key for the given subspace
func (k Keeper) DeletePostID(ctx sdk.Context, subspaceID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.PostIDStoreKey(subspaceID))
}

// --------------------------------------------------------------------------------------------------------------------

// ValidatePostReference checks the post reference to make sure that the referenced
// post's author has not blocked the user referencing the post
func (k Keeper) ValidatePostReference(ctx sdk.Context, postAuthor string, subspaceID uint64, referenceID uint64) error {
	// Make sure the referenced post exists
	originalPost, found := k.GetPost(ctx, subspaceID, referenceID)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %d does not exist", referenceID)
	}

	// Make sure the original author has not blocked the post author
	if k.HasUserBlocked(ctx, originalPost.Author, postAuthor, subspaceID) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "author of post %d has blocked you", referenceID)
	}

	return nil
}

// ValidatePostReply checks the original post reply settings to make sure that
// only specified users can answer to the post
func (k Keeper) ValidatePostReply(ctx sdk.Context, postAuthor string, subspaceID uint64, referenceID uint64) error {
	replyPost, found := k.GetPost(ctx, subspaceID, referenceID)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %d does not exist", referenceID)
	}

	switch replyPost.ReplySettings {
	case types.REPLY_SETTING_FOLLOWERS:
		// We need to make sure that a relationship between post author -> original author exists
		if !k.HasRelationship(ctx, postAuthor, replyPost.Author, subspaceID) {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "only followers of the author can reply to this post")
		}

	case types.REPLY_SETTING_MUTUAL:
		// We need to make sure that both relationships exist (post author -> original author and original author -> post author)
		if !k.HasRelationship(ctx, postAuthor, replyPost.Author, subspaceID) || !k.HasRelationship(ctx, replyPost.Author, postAuthor, subspaceID) {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "only mutual connections of the author can reply to this post")
		}

	case types.REPLY_SETTING_MENTIONS:
		// We need to check each mention of the original post
		if !replyPost.IsUserMentioned(postAuthor) {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "only mentioned users can reply to this post")
		}
	}

	return nil
}

// ValidatePost validates the given post based on the current params, returning an error if anything is wrong
func (k Keeper) ValidatePost(ctx sdk.Context, post types.Post) error {
	params := k.GetParams(ctx)

	// Validate the conversation reference
	if post.ConversationID != 0 {
		err := k.ValidatePostReference(ctx, post.Author, post.SubspaceID, post.ConversationID)
		if err != nil {
			return err
		}
	}

	// Validate the post references
	for _, reference := range post.ReferencedPosts {
		err := k.ValidatePostReference(ctx, post.Author, post.SubspaceID, reference.PostID)
		if err != nil {
			return err
		}

		if reference.Type == types.TYPE_REPLIED_TO {
			err = k.ValidatePostReply(ctx, post.Author, post.SubspaceID, reference.PostID)
			if err != nil {
				return err
			}
		}
	}

	// Check the post text length to make sure it's not exceeding the max length
	if uint32(len(post.Text)) > params.MaxTextLength {
		return sdkerrors.Wrapf(types.ErrInvalidPost, "text exceed max length allowed")
	}

	err := post.Validate()
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return nil
}

// SavePost saves the given post inside the current context.
func (k Keeper) SavePost(ctx sdk.Context, post types.Post) {
	store := ctx.KVStore(k.storeKey)

	// Store the post
	store.Set(types.PostStoreKey(post.SubspaceID, post.ID), k.cdc.MustMarshal(&post))

	// If the initial attachment id does not exist, create it now
	if !k.HasAttachmentID(ctx, post.SubspaceID, post.ID) {
		k.SetAttachmentID(ctx, post.SubspaceID, post.ID, 1)
	}

	k.Logger(ctx).Debug("post saved", "subpace id", post.SubspaceID, "id", post.ID)
	k.AfterPostSaved(ctx, post.SubspaceID, post.ID)
}

// HasPost tells whether the given post exists or not
func (k Keeper) HasPost(ctx sdk.Context, subspaceID uint64, postID uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.PostStoreKey(subspaceID, postID))
}

// GetPost returns the post associated with the given id.
// If there is no post associated with the given id the function will return an empty post and false.
func (k Keeper) GetPost(ctx sdk.Context, subspaceID uint64, postID uint64) (post types.Post, found bool) {
	store := ctx.KVStore(k.storeKey)

	key := types.PostStoreKey(subspaceID, postID)
	if !store.Has(key) {
		return types.Post{}, false
	}

	k.cdc.MustUnmarshal(store.Get(key), &post)
	return post, true
}

// DeletePost deletes the given post and all its attachments from the store
func (k Keeper) DeletePost(ctx sdk.Context, subspaceID uint64, postID uint64) {
	store := ctx.KVStore(k.storeKey)

	// Delete the post
	store.Delete(types.PostStoreKey(subspaceID, postID))

	// Delete all the attachments
	k.IteratePostAttachments(ctx, subspaceID, postID, func(_ int64, attachment types.Attachment) (stop bool) {
		k.DeleteAttachment(ctx, attachment.SubspaceID, attachment.PostID, attachment.ID)
		return false
	})

	k.AfterPostDeleted(ctx, subspaceID, postID)
}
