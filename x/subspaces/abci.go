package subspaces

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v5/x/subspaces/keeper"
)

// BeginBlocker called every block, remove expired allowances
func BeginBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	// Remove allowances that have been expired by the current block time
	keeper.RemoveExpiredAllowances(ctx, ctx.BlockTime())
}
