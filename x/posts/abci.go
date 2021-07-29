package posts

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/keeper"
)

// EndBlocker takes care of executing the tokenomics related to posts
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	k.ExecuteTokenomics(ctx)
}
