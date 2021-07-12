package profiles

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profiles/keeper"
)

// handleEndBLock TODO: introduce
func handleEndBlock(ctx sdk.Context, k keeper.Keeper) {
	k.DeleteUnregisteredUserRelationships(ctx)
}
