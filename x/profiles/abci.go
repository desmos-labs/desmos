package profiles

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/profiles/keeper"
	"github.com/desmos-labs/desmos/v4/x/profiles/types"
)

// BeginBlocker is called every block, processes expired application links
func BeginBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	// Iterate over all the expiring application links
	keeper.IterateExpiringApplicationLinks(ctx, func(_ int64, link types.ApplicationLink) (stop bool) {
		// Delete the application link
		keeper.DeleteApplicationLink(ctx, link)
		return false
	})
}
