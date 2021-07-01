package keeper

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

// IteratePosts iterates through the posts set and performs the provided function
func (k Keeper) IteratePosts(ctx sdk.Context, fn func(index int64, post types.Post) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PostStorePrefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var post types.Post
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &post)
		stop := fn(i, post)
		if stop {
			break
		}
		i++
	}
}

// ValidatePost checks if the given post is valid according to the current posts' module params
func (k Keeper) ValidatePost(ctx sdk.Context, post types.Post) error {
	params := k.GetParams(ctx)
	maxMsgLen := params.MaxPostMessageLength.Int64()
	maxOpFieldNum := params.MaxAdditionalAttributesFieldsNumber.Int64()
	maxOpFieldValLen := params.MaxAdditionalAttributesFieldValueLength.Int64()

	if int64(len(post.Message)) > maxMsgLen {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("post with id %s has more than %d characters", post.PostID, maxMsgLen))
	}

	if int64(len(post.AdditionalAttributes)) > maxOpFieldNum {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("post with id %s contains additional attributes with more than %d key-value pairs",
				post.PostID, maxOpFieldNum))
	}

	for _, additionalAttribute := range post.AdditionalAttributes {
		if int64(len(strings.TrimSpace(additionalAttribute.Value))) > maxOpFieldValLen {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("post with id %s has additional attributes with key %s which value exceeds %d characters.",
					post.PostID, additionalAttribute.Key, maxOpFieldValLen))
		}
	}

	return post.Validate()
}

// IsCreatorBlockedBySomeTags checks if some of the post's tags have blocked the post's creator
func (k Keeper) IsCreatorBlockedBySomeTags(ctx sdk.Context, attachments types.Attachments, creator, subspace string) error {
	for _, attachment := range attachments {
		for _, tag := range attachment.Tags {
			// check if the request's receiver has blocked the sender before
			if k.IsUserBlocked(ctx, tag, creator, subspace) {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
					fmt.Sprintf("The user with address %s has blocked you", tag))
			}
		}
	}
	return nil
}

// ExtractReactionValueAndShortcode parse the given registeredReactions returning its correct value and shortcode
func (k Keeper) ExtractReactionValueAndShortcode(ctx sdk.Context, reaction string, subspace string) (string, string, error) {
	var reactionShortcode, reactionValue string

	// Parse registeredReactions adding the variation selector-16 to let the emoji being readable
	parsedReaction := strings.ReplaceAll(reaction, "️", "")

	if emojiReact, found := types.GetEmojiByShortCodeOrValue(reaction); found {
		reactionShortcode = emojiReact.Shortcodes[0]
		reactionValue = emojiReact.Value
	} else {
		// The registeredReactions is a shortcode that should be registered
		regReaction, registered := k.GetRegisteredReaction(ctx, reaction, subspace)
		if !registered {
			return "", "", sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("short code %s must be registered before using it", parsedReaction))
		}

		reactionShortcode = regReaction.ShortCode
		reactionValue = regReaction.Value
	}

	return reactionShortcode, reactionValue, nil
}

// IterateUserAnswers iterates through the user answers and perform the provided function
func (k Keeper) IterateUserAnswers(ctx sdk.Context, fn func(index int64, answer types.UserAnswer) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.UserAnswersStorePrefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		answer := types.MustUnmarshalUserAnswer(k.cdc, iterator.Value())
		stop := fn(i, answer)
		if stop {
			break
		}
		i++
	}
}

// IterateUserAnswersByPost iterates through the user answers with the given post id and performs the provided function
func (k Keeper) IterateUserAnswersByPost(ctx sdk.Context, postID string, fn func(index int64, answer types.UserAnswer) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.UserAnswersByPostPrefix(postID))
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		answer := types.MustUnmarshalUserAnswer(k.cdc, iterator.Value())
		stop := fn(i, answer)
		if stop {
			break
		}
		i++
	}
}

// IterateRegisteredReactions iterates through the registered reactions and performs the provided function
func (k Keeper) IterateRegisteredReactions(ctx sdk.Context, fn func(index int64, reaction types.RegisteredReaction) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RegisteredReactionsStorePrefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		reaction := types.MustUnmarshalRegisteredReaction(k.cdc, iterator.Value())
		stop := fn(i, reaction)
		if stop {
			break
		}
		i++
	}
}

// IteratePostReactions iterates through the post reactions and performs the provided function
func (k Keeper) IteratePostReactions(ctx sdk.Context, fn func(index int64, reaction types.PostReaction) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PostReactionsStorePrefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		reaction := types.MustUnmarshalPostReaction(k.cdc, iterator.Value())

		stop := fn(i, reaction)
		if stop {
			break
		}
		i++
	}
}

// IteratePostReactionsByPost iterates through the post reactions added to the post with the given id and performs the provided function
func (k Keeper) IteratePostReactionsByPost(ctx sdk.Context, postID string, fn func(index int64, reaction types.PostReaction) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PostReactionsPrefix(postID))
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		reaction := types.MustUnmarshalPostReaction(k.cdc, iterator.Value())
		stop := fn(i, reaction)
		if stop {
			break
		}
		i++
	}
}

// IteratePostReports iterates through the post's reports with the given id and performs the provided function
func (k Keeper) IteratePostReports(ctx sdk.Context, postID string, fn func(index int64, report types.Report) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ReportsByPostIDPrefix(postID))
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		report := types.MustUnmarshalReport(k.cdc, iterator.Value())
		stop := fn(i, report)
		if stop {
			break
		}
		i++
	}
}
