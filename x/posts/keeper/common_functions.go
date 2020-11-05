package keeper

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/posts/types"
)

// IteratePosts iterates through the posts set and performs the provided function
// It makes a copy of the posts array which is done only for sorting purposes.
func (k Keeper) IteratePosts(ctx sdk.Context, fn func(index int64, post types.Post) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PostStorePrefix)
	defer iterator.Close()

	var posts []types.Post
	for ; iterator.Valid(); iterator.Next() {
		var post types.Post
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &post)
		posts = append(posts, post)
	}

	i := int64(0)
	postsSorted := make([]types.Post, len(posts))
	for _, post := range posts {
		var index types.PostIndex
		k.cdc.MustUnmarshalBinaryBare(store.Get(types.PostIndexedIDStoreKey(post.PostID)), &index)
		postsSorted[index.Value-1] = post
	}

	//freeing up memory
	//nolint
	posts = nil

	for _, post := range postsSorted {
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
	maxOpFieldNum := params.MaxOptionalDataFieldsNumber.Int64()
	maxOpFieldValLen := params.MaxOptionalDataFieldValueLength.Int64()

	if int64(len(post.Message)) > maxMsgLen {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("post with id %s has more than %d characters", post.PostID, maxMsgLen))
	}

	if int64(len(post.OptionalData)) > maxOpFieldNum {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("post with id %s contains optional data with more than %d key-value pairs",
				post.PostID, maxOpFieldNum))
	}

	for _, optionalData := range post.OptionalData {
		if int64(len(strings.TrimSpace(optionalData.Value))) > maxOpFieldValLen {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("post with id %s has optional data with key %s which value exceeds %d characters.",
					post.PostID, optionalData.Key, maxOpFieldValLen))
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
