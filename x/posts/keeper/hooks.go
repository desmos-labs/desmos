package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
)

// Implement PostsHooks interface
var _ types.PostsHooks = Keeper{}

func (k Keeper) AfterPostSaved(ctx sdk.Context, subspaceID uint64, postID uint64) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) AfterPostDeleted(ctx sdk.Context, subspaceID uint64, postID uint64) {
	//TODO implement me
	panic("implement me")
}
