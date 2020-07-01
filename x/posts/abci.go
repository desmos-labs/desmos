package posts

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EndBlocker called upon each block end to close expired polls
// TODO look how to iterate only over open poll
func EndBlocker(ctx sdk.Context, keeper Keeper) {

	store := ctx.KVStore(keeper.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, PostStorePrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var post Post
		keeper.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &post)
		if ctx.BlockTime().After(post.PollData.EndDate) || ctx.BlockTime().Equal(post.PollData.EndDate) {
			post.PollData.Open = false
			post.LastEdited = ctx.BlockTime()
			keeper.SavePost(ctx, post)

			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					EventTypeClosePoll,
					sdk.NewAttribute(AttributeKeyPostID, post.PostID.String()),
					sdk.NewAttribute(AttributeKeyPostOwner, post.Creator.String()),
				),
			)
		}

	}
}
