package subspaces

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/x/subspaces/keeper"
)

// BeginBlocker is called every block and takes care of removing expired allowances
func BeginBlocker(ctx sdk.Context, keeper *keeper.Keeper) {
	keeper.RemoveExpiredAllowances(ctx, ctx.BlockTime())
}
