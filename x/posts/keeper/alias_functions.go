package keeper

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"

	subspacetypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
)

// HasSubspace checks whether the given subspace exists or not
func (k *Keeper) HasSubspace(ctx sdk.Context, subspaceID uint64) bool {
	return k.sk.HasSubspace(ctx, subspaceID)
}

// HasPermission checks whether the given user has the provided permissions or not
func (k *Keeper) HasPermission(ctx sdk.Context, subspaceID uint64, user sdk.AccAddress, permission subspacetypes.Permission) bool {
	return k.sk.HasPermission(ctx, subspaceID, user, permission)
}

// --------------------------------------------------------------------------------------------------------------------

// IteratePostIDs iterates over all the next post ids and performs the provided function
func (k Keeper) IteratePostIDs(ctx sdk.Context, fn func(index int64, subspaceID uint64, postID uint64) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PostIDPrefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		subspaceID := subspacetypes.GetSubspaceIDFromBytes(bytes.TrimPrefix(iterator.Key(), types.PostIDPrefix))
		postID := types.GetPostIDFromBytes(iterator.Value())
		stop := fn(i, subspaceID, postID)
		if stop {
			break
		}
		i++
	}
}

// --------------------------------------------------------------------------------------------------------------------

// IteratePosts iterates over all the posts stored inside the context and performs the provided function
func (k Keeper) IteratePosts(ctx sdk.Context, fn func(index int64, post types.Post) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PostPrefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var post types.Post
		k.cdc.MustUnmarshal(iterator.Value(), &post)
		stop := fn(i, post)
		if stop {
			break
		}
		i++
	}
}

// IterateSubspacePosts iterates over all the posts stored inside the given subspace and performs the provided function
func (k Keeper) IterateSubspacePosts(ctx sdk.Context, subspaceID uint64, fn func(index int64, post types.Post) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SubspacePostsPrefix(subspaceID))
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var post types.Post
		k.cdc.MustUnmarshal(iterator.Value(), &post)
		stop := fn(i, post)
		if stop {
			break
		}
		i++
	}
}

// GetSubspacePosts returns all the posts associated to the given subspace
func (k Keeper) GetSubspacePosts(ctx sdk.Context, subspaceID uint64) []types.Post {
	var posts []types.Post
	k.IterateSubspacePosts(ctx, subspaceID, func(index int64, post types.Post) (stop bool) {
		posts = append(posts, post)
		return false
	})
	return posts
}

// --------------------------------------------------------------------------------------------------------------------

// IterateAttachments iterates over all the attachments in the given context and performs the provided function
func (k Keeper) IterateAttachments(ctx sdk.Context, fn func(index int64, attachment types.Attachment) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.AttachmentPrefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var attachment types.Attachment
		k.cdc.MustUnmarshal(iterator.Value(), &attachment)
		stop := fn(i, attachment)
		if stop {
			break
		}
		i++
	}
}

// IteratePostAttachments iterates through the attachments associated with the provided post and performs the given function
func (k Keeper) IteratePostAttachments(ctx sdk.Context, subspaceID uint64, postID uint64, fn func(index int64, attachment types.Attachment) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PostAttachmentsPrefix(subspaceID, postID))
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var attachment types.Attachment
		k.cdc.MustUnmarshal(iterator.Value(), &attachment)
		stop := fn(i, attachment)
		if stop {
			break
		}
		i++
	}
}

// GetPostAttachments returns all the attachments associated to the given post
func (k Keeper) GetPostAttachments(ctx sdk.Context, subspaceID uint64, postID uint64) []types.Attachment {
	var attachments []types.Attachment
	k.IteratePostAttachments(ctx, subspaceID, postID, func(index int64, attachment types.Attachment) (stop bool) {
		attachments = append(attachments, attachment)
		return false
	})
	return attachments
}

// --------------------------------------------------------------------------------------------------------------------

// IterateUserAnswers iterates over all the poll user answers and performs the provided function
func (k Keeper) IterateUserAnswers(ctx sdk.Context, fn func(index int64, answer types.UserAnswer) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PollAnswerPrefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var answer types.UserAnswer
		k.cdc.MustUnmarshal(iterator.Value(), &answer)
		stop := fn(i, answer)
		if stop {
			break
		}
		i++
	}
}

// IteratePollUserAnswers iterates through the answers to the given poll and performs the provided function
func (k Keeper) IteratePollUserAnswers(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32, fn func(index int64, answer types.UserAnswer) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PollAnswersPrefix(subspaceID, postID, pollID))
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var answer types.UserAnswer
		k.cdc.MustUnmarshal(iterator.Value(), &answer)
		stop := fn(i, answer)
		if stop {
			break
		}
		i++
	}
}

// GetPollUserAnswers returns all the user answers for the given poll
func (k Keeper) GetPollUserAnswers(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32) []types.UserAnswer {
	var answers []types.UserAnswer
	k.IteratePollUserAnswers(ctx, subspaceID, postID, pollID, func(index int64, answer types.UserAnswer) (stop bool) {
		answers = append(answers, answer)
		return false
	})
	return answers
}

// --------------------------------------------------------------------------------------------------------------------

// IteratePollsTallyResults iterates over all the polls tally results and performs the provided function
func (k Keeper) IteratePollsTallyResults(ctx sdk.Context, fn func(index int64, results types.PollTallyResults) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PollTallyResultPrefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var results types.PollTallyResults
		k.cdc.MustUnmarshal(iterator.Value(), &results)
		stop := fn(i, results)
		if stop {
			break
		}
		i++
	}
}
