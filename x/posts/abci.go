package posts

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
)

// EndBlocker called upon each block end to close expired polls
// TODO look how to iterate only over open poll
func EndBlocker(ctx sdk.Context, keeper Keeper) {

	store := ctx.KVStore(keeper.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.PostStorePrefix))

	for ; iterator.Valid(); iterator.Next() {
		var post types.Post
		keeper.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &post)
		if ctx.BlockTime().After(post.PollData.EndDate) || ctx.BlockTime().Equal(post.PollData.EndDate) {
			post.PollData.Open = false
			post.LastEdited = ctx.BlockTime()
			keeper.SavePost(ctx, post)

			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeClosePoll,
					sdk.NewAttribute(types.AttributeKeyPostID, post.PostID.String()),
					sdk.NewAttribute(types.AttributeKeyPostOwner, post.Creator.String()),
				),
			)
		}

	}
}
