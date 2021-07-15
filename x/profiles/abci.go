package profiles

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profiles/keeper"
)

// handleEndBLock clean up the state during end block
func handleEndBlock(ctx sdk.Context, k keeper.Keeper) {
	k.DeleteUnregisteredUserRelationshipsAndBlocks(ctx)
}
