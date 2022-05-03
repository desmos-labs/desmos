package posts

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/posts/keeper"
	"github.com/desmos-labs/desmos/v3/x/posts/types"
)

// EndBlocker called every block, process ended polls
func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	// Iterate over all the active polls that have been ended by the current block time
	keeper.IterateActivePollsQueue(ctx, ctx.BlockTime(), func(index int64, poll types.Attachment) (stop bool) {
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

		// When poll ends
		keeper.AfterPollVotingPeriodEnded(ctx, poll.SubspaceID, poll.PostID, poll.ID)

		return false
	})
}
