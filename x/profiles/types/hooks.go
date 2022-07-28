package types

// DONTCOVER

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Event Hooks
// These can be utilized to communicate between a profiles keeper and another
// keeper which must take particular actions when profiles/DTag transfer requests/app links/chain links change
// state. The second keeper must implement this interface, which then the
// profiles keeper can call.

// ProfilesHooks event hooks for profiles objects (noalias)
type ProfilesHooks interface {
	AfterProfileSaved(ctx sdk.Context, profile *Profile)   // Must be called when a profile is saved
	AfterProfileDeleted(ctx sdk.Context, profile *Profile) // Must be called when a profile is deleted

	AfterDTagTransferRequestCreated(ctx sdk.Context, request DTagTransferRequest)                  // Must be called then a DTag transfer request is created
	AfterDTagTransferRequestAccepted(ctx sdk.Context, request DTagTransferRequest, newDTag string) // Must be called when a DTag transfer request is accepted
	AfterDTagTransferRequestDeleted(ctx sdk.Context, sender, recipient string)                     // Must be called when a DTag transfer request is deleted

	AfterChainLinkSaved(ctx sdk.Context, link ChainLink)   // Must be called when a chain link is saved
	AfterChainLinkDeleted(ctx sdk.Context, link ChainLink) // Must be called when a chain link is deleted

	AfterApplicationLinkSaved(ctx sdk.Context, link ApplicationLink)   // Must be called when an application link is saved
	AfterApplicationLinkDeleted(ctx sdk.Context, link ApplicationLink) // Must be called when an application link is deleted
}

// --------------------------------------------------------------------------------------------------------------------

// MultiProfilesHooks combines multiple profiles hooks, all hook functions are run in array sequence
type MultiProfilesHooks []ProfilesHooks

func NewMultiProfilesHooks(hooks ...ProfilesHooks) MultiProfilesHooks {
	return hooks
}

// AfterProfileSaved implements ProfilesHooks
func (h MultiProfilesHooks) AfterProfileSaved(ctx sdk.Context, profile *Profile) {
	for _, hook := range h {
		hook.AfterProfileSaved(ctx, profile)
	}
}

// AfterProfileDeleted implements ProfilesHooks
func (h MultiProfilesHooks) AfterProfileDeleted(ctx sdk.Context, profile *Profile) {
	for _, hook := range h {
		hook.AfterProfileDeleted(ctx, profile)
	}
}

// AfterDTagTransferRequestCreated implements ProfilesHooks
func (h MultiProfilesHooks) AfterDTagTransferRequestCreated(ctx sdk.Context, request DTagTransferRequest) {
	for _, hook := range h {
		hook.AfterDTagTransferRequestCreated(ctx, request)
	}
}

// AfterDTagTransferRequestAccepted implements ProfilesHooks
func (h MultiProfilesHooks) AfterDTagTransferRequestAccepted(ctx sdk.Context, request DTagTransferRequest, newDTag string) {
	for _, hook := range h {
		hook.AfterDTagTransferRequestAccepted(ctx, request, newDTag)
	}
}

// AfterDTagTransferRequestDeleted implements ProfilesHooks
func (h MultiProfilesHooks) AfterDTagTransferRequestDeleted(ctx sdk.Context, sender, recipient string) {
	for _, hook := range h {
		hook.AfterDTagTransferRequestDeleted(ctx, sender, recipient)
	}
}

// AfterChainLinkSaved implements ProfilesHooks
func (h MultiProfilesHooks) AfterChainLinkSaved(ctx sdk.Context, link ChainLink) {
	for _, hook := range h {
		hook.AfterChainLinkSaved(ctx, link)
	}
}

// AfterChainLinkDeleted implements ProfilesHooks
func (h MultiProfilesHooks) AfterChainLinkDeleted(ctx sdk.Context, link ChainLink) {
	for _, hook := range h {
		hook.AfterChainLinkDeleted(ctx, link)
	}
}

// AfterApplicationLinkSaved implements ProfilesHooks
func (h MultiProfilesHooks) AfterApplicationLinkSaved(ctx sdk.Context, link ApplicationLink) {
	for _, hook := range h {
		hook.AfterApplicationLinkSaved(ctx, link)
	}
}

// AfterUserPermissionDeleted implements ProfilesHooks
func (h MultiProfilesHooks) AfterApplicationLinkDeleted(ctx sdk.Context, link ApplicationLink) {
	for _, hook := range h {
		hook.AfterApplicationLinkDeleted(ctx, link)
	}
}
