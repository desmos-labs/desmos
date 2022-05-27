package keeper

import (
	"bytes"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	subspacetypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
)

// HasSubspace checks whether the given subspace exists or not
func (k Keeper) HasSubspace(ctx sdk.Context, subspaceID uint64) bool {
	return k.sk.HasSubspace(ctx, subspaceID)
}

// HasPermission checks whether the given user has the provided permissions or not
func (k Keeper) HasPermission(ctx sdk.Context, subspaceID uint64, user sdk.AccAddress, permission subspacetypes.Permission) bool {
	return k.sk.HasPermission(ctx, subspaceID, user, permission)
}

// HasUserBlocked tells whether the given blocker has blocked the user inside the provided subspace
func (k Keeper) HasUserBlocked(ctx sdk.Context, blocker, user string, subspaceID uint64) bool {
	return k.rk.HasUserBlocked(ctx, blocker, user, subspaceID)
}

// HasRelationship tells whether the relationship between the user and counterparty exists for the given subspace
func (k Keeper) HasRelationship(ctx sdk.Context, user, counterparty string, subspaceID uint64) bool {
	return k.rk.HasRelationship(ctx, user, counterparty, subspaceID)
}

// --------------------------------------------------------------------------------------------------------------------

// IteratePostIDs iterates over all the next post ids and performs the provided function
func (k Keeper) IteratePostIDs(ctx sdk.Context, fn func(index int64, subspaceID uint64, postID uint64) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.NextPostIDPrefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		subspaceID := subspacetypes.GetSubspaceIDFromBytes(bytes.TrimPrefix(iterator.Key(), types.NextPostIDPrefix))
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

// GetPosts returns all the posts stored inside the given context
func (k Keeper) GetPosts(ctx sdk.Context) []types.Post {
	var posts []types.Post
	k.IteratePosts(ctx, func(index int64, post types.Post) (stop bool) {
		posts = append(posts, post)
		return false
	})
	return posts
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

// --------------------------------------------------------------------------------------------------------------------

// IterateActivePolls iterates over the polls in the active polls queue and performs the provided function
func (k Keeper) IterateActivePolls(ctx sdk.Context, fn func(index int64, poll types.Attachment) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ActivePollQueuePrefix)
	defer iterator.Close()

	index := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		subspaceID, postID, pollID, _ := types.SplitActivePollQueueKey(iterator.Key())
		attachment, found := k.GetAttachment(ctx, subspaceID, postID, pollID)
		if !found || !types.IsPoll(attachment) {
			panic(fmt.Sprintf("poll %d %d %d does not exist", subspaceID, postID, pollID))
		}

		stop := fn(index, attachment)
		if stop {
			break
		}
		index++
	}
}

// IterateActivePollsQueue iterates over the polls that are still active by the time given performs the provided function
func (k Keeper) IterateActivePollsQueue(ctx sdk.Context, endTime time.Time, fn func(index int64, poll types.Attachment) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := store.Iterator(types.ActivePollQueuePrefix, sdk.PrefixEndBytes(types.ActivePollByTimeKey(endTime)))
	defer iterator.Close()

	index := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		subspaceID, postID, pollID, _ := types.SplitActivePollQueueKey(iterator.Key())
		attachment, found := k.GetAttachment(ctx, subspaceID, postID, pollID)
		if !found || !types.IsPoll(attachment) {
			panic(fmt.Sprintf("poll %d %d %d does not exist", subspaceID, postID, pollID))
		}

		stop := fn(index, attachment)
		if stop {
			break
		}
		index++
	}
}

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

// IterateUserAnswers iterates over all the polls user answers and performs the provided function
func (k Keeper) IterateUserAnswers(ctx sdk.Context, fn func(index int64, answer types.UserAnswer) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.UserAnswerPrefix)
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
