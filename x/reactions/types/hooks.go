package types

// DONTCOVER

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Event Hooks
// These can be utilized to communicate between a reactions keeper and another
// keeper which must take particular actions when reactions change
// state. The second keeper must implement this interface, which then the
// reactions keeper can call.

// ReactionsHooks event hooks for reactions objects (noalias)
type ReactionsHooks interface {
	AfterReactionSaved(ctx sdk.Context, subspaceID uint64, reactionID uint32)   // Must be called when a reaction is saved
	AfterReactionDeleted(ctx sdk.Context, subspaceID uint64, reactionID uint32) // Must be called when a reaction is deleted

	AfterRegisteredReactionSaved(ctx sdk.Context, subspaceID uint64, registeredReactionID uint32)   // Must be called when a registered reaction is saved
	AfterRegisteredReactionDeleted(ctx sdk.Context, subspaceID uint64, registeredReactionID uint32) // Must be called when a registered reaction is deleted

	AfterReactionParamsSaved(ctx sdk.Context, subspaceID uint64) // Must be called when some reaction params are saved
}

// --------------------------------------------------------------------------------------------------------------------

// MultiReactionsHooks combines multiple posts hooks, all hook functions are run in array sequence
type MultiReactionsHooks []ReactionsHooks

func NewMultiReactionsHooks(hooks ...ReactionsHooks) MultiReactionsHooks {
	return hooks
}

// AfterReactionSaved implements ReactionsHooks
func (h MultiReactionsHooks) AfterReactionSaved(ctx sdk.Context, subspaceID uint64, reactionID uint32) {
	for _, hook := range h {
		hook.AfterReactionSaved(ctx, subspaceID, reactionID)
	}
}

// AfterReactionDeleted implements ReactionsHooks
func (h MultiReactionsHooks) AfterReactionDeleted(ctx sdk.Context, subspaceID uint64, reactionID uint32) {
	for _, hook := range h {
		hook.AfterReactionDeleted(ctx, subspaceID, reactionID)
	}
}

// AfterRegisteredReactionSaved implements ReactionsHooks
func (h MultiReactionsHooks) AfterRegisteredReactionSaved(ctx sdk.Context, subspaceID uint64, registeredReactionID uint32) {
	for _, hook := range h {
		hook.AfterRegisteredReactionSaved(ctx, subspaceID, registeredReactionID)
	}
}

// AfterRegisteredReactionDeleted implements ReactionsHooks
func (h MultiReactionsHooks) AfterRegisteredReactionDeleted(ctx sdk.Context, subspaceID uint64, registeredReactionID uint32) {
	for _, hook := range h {
		hook.AfterRegisteredReactionDeleted(ctx, subspaceID, registeredReactionID)
	}
}

// AfterReactionParamsSaved implements ReactionsHooks
func (h MultiReactionsHooks) AfterReactionParamsSaved(ctx sdk.Context, subspaceID uint64) {
	for _, hook := range h {
		hook.AfterReactionParamsSaved(ctx, subspaceID)
	}
}
