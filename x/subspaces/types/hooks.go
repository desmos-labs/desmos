package types

// DONTCOVER

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Event Hooks
// These can be utilized to communicate between a subspaces keeper and another
// keeper which must take particular actions when subspaces/groups/permissions change
// state. The second keeper must implement this interface, which then the
// subspaces keeper can call.

// SubspacesHooks event hooks for subspaces objects (noalias)
type SubspacesHooks interface {
	AfterSubspaceSaved(ctx sdk.Context, subspaceID uint64)   // Must be called when a subspace is saved
	AfterSubspaceDeleted(ctx sdk.Context, subspaceID uint64) // Must be called when a subspace is deleted

	AfterSubspaceGroupSaved(ctx sdk.Context, subspaceID uint64, groupID uint32)                              // Must be called when a subspace group is created
	AfterSubspaceGroupMemberAdded(ctx sdk.Context, subspaceID uint64, groupID uint32, user sdk.AccAddress)   // Must be called when a user is added to a group
	AfterSubspaceGroupMemberRemoved(ctx sdk.Context, subspaceID uint64, groupID uint32, user sdk.AccAddress) // Must be called when a user is removed from a group
	AfterSubspaceGroupDeleted(ctx sdk.Context, subspaceID uint64, groupID uint32)                            // Must be called when a subspace group is deleted

	AfterUserPermissionSet(ctx sdk.Context, subspaceID uint64, user sdk.AccAddress, permissions Permissions) // Must be called when a permissions is set for a user
	AfterUserPermissionRemoved(ctx sdk.Context, subspaceID uint64, user sdk.AccAddress)                      // Must be called when a permissions is removed for a user
}

// --------------------------------------------------------------------------------------------------------------------

// MultiSubspacesHooks combines multiple subspaces hooks, all hook functions are run in array sequence
type MultiSubspacesHooks []SubspacesHooks

func NewMultiSubspacesHooks(hooks ...SubspacesHooks) MultiSubspacesHooks {
	return hooks
}

// AfterSubspaceSaved implements SubspacesHook
func (h MultiSubspacesHooks) AfterSubspaceSaved(ctx sdk.Context, subspaceID uint64) {
	for _, hook := range h {
		hook.AfterSubspaceSaved(ctx, subspaceID)
	}
}

// AfterSubspaceDeleted implements SubspacesHook
func (h MultiSubspacesHooks) AfterSubspaceDeleted(ctx sdk.Context, subspaceID uint64) {
	for _, hook := range h {
		hook.AfterSubspaceDeleted(ctx, subspaceID)
	}
}

// AfterSubspaceGroupSaved implements SubspacesHook
func (h MultiSubspacesHooks) AfterSubspaceGroupSaved(ctx sdk.Context, subspaceID uint64, groupID uint32) {
	for _, hook := range h {
		hook.AfterSubspaceGroupSaved(ctx, subspaceID, groupID)
	}
}

// AfterSubspaceGroupMemberAdded implements SubspacesHook
func (h MultiSubspacesHooks) AfterSubspaceGroupMemberAdded(ctx sdk.Context, subspaceID uint64, groupID uint32, user sdk.AccAddress) {
	for _, hook := range h {
		hook.AfterSubspaceGroupMemberAdded(ctx, subspaceID, groupID, user)
	}
}

// AfterSubspaceGroupMemberRemoved implements SubspacesHook
func (h MultiSubspacesHooks) AfterSubspaceGroupMemberRemoved(ctx sdk.Context, subspaceID uint64, groupID uint32, user sdk.AccAddress) {
	for _, hook := range h {
		hook.AfterSubspaceGroupMemberRemoved(ctx, subspaceID, groupID, user)
	}
}

// AfterSubspaceGroupDeleted implements SubspacesHook
func (h MultiSubspacesHooks) AfterSubspaceGroupDeleted(ctx sdk.Context, subspaceID uint64, groupID uint32) {
	for _, hook := range h {
		hook.AfterSubspaceGroupDeleted(ctx, subspaceID, groupID)
	}
}

// AfterUserPermissionSet implements SubspacesHook
func (h MultiSubspacesHooks) AfterUserPermissionSet(ctx sdk.Context, subspaceID uint64, user sdk.AccAddress, permissions Permissions) {
	for _, hook := range h {
		hook.AfterUserPermissionSet(ctx, subspaceID, user, permissions)
	}
}

// AfterUserPermissionRemoved implements SubspacesHook
func (h MultiSubspacesHooks) AfterUserPermissionRemoved(ctx sdk.Context, subspaceID uint64, user sdk.AccAddress) {
	for _, hook := range h {
		hook.AfterUserPermissionRemoved(ctx, subspaceID, user)
	}
}
