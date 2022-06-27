package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/reactions/types"
)

// Implement ReactionsHooks interface
var _ types.ReactionsHooks = Keeper{}

// AfterReactionSaved implements types.ReactionsHooks
func (k Keeper) AfterReactionSaved(ctx sdk.Context, subspaceID uint64, reactionID uint32) {
	if k.hooks != nil {
		k.hooks.AfterReactionSaved(ctx, subspaceID, reactionID)
	}
}

// AfterReactionDeleted implements types.ReactionsHooks
func (k Keeper) AfterReactionDeleted(ctx sdk.Context, subspaceID uint64, reactionID uint32) {
	if k.hooks != nil {
		k.hooks.AfterReactionDeleted(ctx, subspaceID, reactionID)
	}
}

// AfterRegisteredReactionSaved implements types.ReactionsHooks
func (k Keeper) AfterRegisteredReactionSaved(ctx sdk.Context, subspaceID uint64, registeredReactionID uint32) {
	if k.hooks != nil {
		k.hooks.AfterRegisteredReactionSaved(ctx, subspaceID, registeredReactionID)
	}
}

// AfterRegisteredReactionDeleted implements types.ReactionsHooks
func (k Keeper) AfterRegisteredReactionDeleted(ctx sdk.Context, subspaceID uint64, registeredReactionID uint32) {
	if k.hooks != nil {
		k.hooks.AfterRegisteredReactionDeleted(ctx, subspaceID, registeredReactionID)
	}
}

// AfterReactionParamsSaved implements types.ReactionsHooks
func (k Keeper) AfterReactionParamsSaved(ctx sdk.Context, subspaceID uint64) {
	if k.hooks != nil {
		k.hooks.AfterReactionParamsSaved(ctx, subspaceID)
	}
}
