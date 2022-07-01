package posts

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/posts/keeper"
	"github.com/desmos-labs/desmos/v4/x/posts/types"
)

// EndBlocker called every block, process ended polls
func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	// Iterate over all the active polls that have been ended by the current block time
	keeper.IterateActivePollsQueue(ctx, ctx.BlockTime(), func(poll types.Attachment) (stop bool) {
		// Compute the poll results
		results := keeper.Tally(ctx, poll.SubspaceID, poll.PostID, poll.ID)

		// Update the content with the results
		content := poll.Content.GetCachedValue().(*types.Poll)
		content.FinalTallyResults = results

		contentAny, err := codectypes.NewAnyWithValue(content)
		if err != nil {
			panic(err)
		}
		poll.Content = contentAny

		keeper.SaveAttachment(ctx, poll)
		keeper.RemoveFromActivePollQueue(ctx, poll)

		// Emit an event
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeTallyPoll,
				sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", poll.SubspaceID)),
				sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", poll.PostID)),
				sdk.NewAttribute(types.AttributeKeyPollID, fmt.Sprintf("%d", poll.ID)),
			),
		)

		// When poll ends
		keeper.AfterPollVotingPeriodEnded(ctx, poll.SubspaceID, poll.PostID, poll.ID)

		return false
	})
}
