package posts

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v5/x/posts/keeper"
	"github.com/desmos-labs/desmos/v5/x/posts/types"
)

// EndBlocker called every block, process ended polls
func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	// Iterate over all the active polls that have been ended by the current block time
	keeper.IterateActivePollsQueue(ctx, ctx.BlockTime(), func(poll types.Attachment) (stop bool) {
		keeper.EndPoll(ctx, poll)

		// Emit an event
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeTallyPoll,
				sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", poll.SubspaceID)),
				sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", poll.PostID)),
				sdk.NewAttribute(types.AttributeKeyPollID, fmt.Sprintf("%d", poll.ID)),
			),
		)

		return false
	})
}
