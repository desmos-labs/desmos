package keeper

import (
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

// IterateUserAnswers iterates through the answers to the given poll and performs the provided function
func (k Keeper) IterateUserAnswers(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32, fn func(index int64, answer types.UserAnswer) (stop bool)) {
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

// GetUserAnswers returns all the user answers for the given poll
func (k Keeper) GetUserAnswers(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32) []types.UserAnswer {
	var answers []types.UserAnswer
	k.IterateUserAnswers(ctx, subspaceID, postID, pollID, func(index int64, answer types.UserAnswer) (stop bool) {
		answers = append(answers, answer)
		return false
	})
	return answers
}
