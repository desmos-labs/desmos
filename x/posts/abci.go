package posts

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
)

// EndBlocker called upon each block end to close expired polls
func EndBlocker(ctx sdk.Context, keeper Keeper) {

	posts := keeper.GetPosts(ctx)

	for _, post := range posts {
		if ctx.BlockTime().After(post.PollData.EndDate) || ctx.BlockTime().Equal(post.PollData.EndDate) {
			keeper.ClosePollPost(ctx, post.PostID)
		}
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeClosePoll,
				sdk.NewAttribute(types.AttributeKeyPostID, post.PostID.String()),
				sdk.NewAttribute(types.AttributeKeyPostOwner, post.Creator.String()),
			),
		)
	}
}
