package keeper

// DONTCOVER

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/x/profiles/types"
)

// Implements ProfilesHooks interface
var _ types.ProfilesHooks = Keeper{}

// AfterProfileSaved - call if hook is registered
func (k Keeper) AfterProfileSaved(ctx sdk.Context, profile *types.Profile) {
	if k.hooks != nil {
		k.hooks.AfterProfileSaved(ctx, profile)
	}
}

// AfterProfileDeleted - call if hook is registered
func (k Keeper) AfterProfileDeleted(ctx sdk.Context, profile *types.Profile) {
	if k.hooks != nil {
		k.hooks.AfterProfileDeleted(ctx, profile)
	}
}

// AfterDTagTransferRequestCreated - call if hook is registered
func (k Keeper) AfterDTagTransferRequestCreated(ctx sdk.Context, request types.DTagTransferRequest) {
	if k.hooks != nil {
		k.hooks.AfterDTagTransferRequestCreated(ctx, request)
	}
}

// AfterDTagTransferRequestAccepted - call if hook is registered
func (k Keeper) AfterDTagTransferRequestAccepted(ctx sdk.Context, request types.DTagTransferRequest, newDTag string) {
	if k.hooks != nil {
		k.hooks.AfterDTagTransferRequestAccepted(ctx, request, newDTag)
	}
}

// AfterDTagTransferRequestDeleted - call if hook is registered
func (k Keeper) AfterDTagTransferRequestDeleted(ctx sdk.Context, sender, recipient string) {
	if k.hooks != nil {
		k.hooks.AfterDTagTransferRequestDeleted(ctx, sender, recipient)
	}
}

// AfterChainLinkSaved - call if hook is registered
func (k Keeper) AfterChainLinkSaved(ctx sdk.Context, link types.ChainLink) {
	if k.hooks != nil {
		k.hooks.AfterChainLinkSaved(ctx, link)
	}
}

// AfterChainLinkDeleted - call if hook is registered
func (k Keeper) AfterChainLinkDeleted(ctx sdk.Context, link types.ChainLink) {
	if k.hooks != nil {
		k.hooks.AfterChainLinkDeleted(ctx, link)
	}
}

// AfterApplicationLinkSaved - call if hook is registered
func (k Keeper) AfterApplicationLinkSaved(ctx sdk.Context, link types.ApplicationLink) {
	if k.hooks != nil {
		k.hooks.AfterApplicationLinkSaved(ctx, link)
	}
}

// AfterUserPermissionDeleted - call if hook is registered
func (k Keeper) AfterApplicationLinkDeleted(ctx sdk.Context, link types.ApplicationLink) {
	if k.hooks != nil {
		k.hooks.AfterApplicationLinkDeleted(ctx, link)
	}
}
