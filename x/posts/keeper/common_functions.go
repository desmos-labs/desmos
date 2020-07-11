package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/posts/types"
)

// IteratePosts iterates through the posts set and performs the provided function
func (k Keeper) IteratePosts(ctx sdk.Context, fn func(index int64, post types.Post) (stop bool)) {
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PostStorePrefix)
	defer iterator.Close()
	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var post types.Post
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &post)
		stop := fn(i, post)
		if stop {
			break
		}
		i++
	}
}

// ValidatePost checks if the given post is valid according to the current posts' module params
func ValidatePost(ctx sdk.Context, k Keeper, post types.Post) error {
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

	for key, value := range post.OptionalData {
		if int64(len(value)) > maxOpFieldValLen {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("post with id %s has optional data with key %s which value exceeds %d characters.",
					post.PostID, key, maxOpFieldValLen))
		}
	}

	return nil
}
