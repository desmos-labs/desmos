package types

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

	AfterSubspaceGroupSaved(ctx sdk.Context, subspaceID uint64, groupName string)                              // Must be called when a subspace group is created
	AfterSubspaceGroupMemberAdded(ctx sdk.Context, subspaceID uint64, groupName string, user sdk.AccAddress)   // Must be called when a user is added to a group
	AfterSubspaceGroupMemberRemoved(ctx sdk.Context, subspaceID uint64, groupName string, user sdk.AccAddress) // Must be called when a user is removed from a group
	AfterSubspaceGroupDeleted(ctx sdk.Context, subspaceID uint64, groupName string)                            // Must be called when a subspace group is deleted

	AfterPermissionSet(ctx sdk.Context, subspaceID uint64, target string, permissions Permission) // Must be called when a permission is set
	AfterPermissionRemoved(ctx sdk.Context, subspaceID uint64, target string)                     // Must be called when a permission is removed
}

// --------------------------------------------------------------------------------------------------------------------

// MultiSubspacesHooks combines multiple subspaces hooks, all hook functions are run in array sequence
type MultiSubspacesHooks []SubspacesHooks

func NewMultiStakingHooks(hooks ...SubspacesHooks) MultiSubspacesHooks {
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
func (h MultiSubspacesHooks) AfterSubspaceGroupSaved(ctx sdk.Context, subspaceID uint64, groupName string) {
	for _, hook := range h {
		hook.AfterSubspaceGroupSaved(ctx, subspaceID, groupName)
	}
}

// AfterSubspaceGroupMemberAdded implements SubspacesHook
func (h MultiSubspacesHooks) AfterSubspaceGroupMemberAdded(ctx sdk.Context, subspaceID uint64, groupName string, user sdk.AccAddress) {
	for _, hook := range h {
		hook.AfterSubspaceGroupMemberAdded(ctx, subspaceID, groupName, user)
	}
}

// AfterSubspaceGroupMemberRemoved implements SubspacesHook
func (h MultiSubspacesHooks) AfterSubspaceGroupMemberRemoved(ctx sdk.Context, subspaceID uint64, groupName string, user sdk.AccAddress) {
	for _, hook := range h {
		hook.AfterSubspaceGroupMemberRemoved(ctx, subspaceID, groupName, user)
	}
}

// AfterSubspaceGroupDeleted implements SubspacesHook
func (h MultiSubspacesHooks) AfterSubspaceGroupDeleted(ctx sdk.Context, subspaceID uint64, groupName string) {
	for _, hook := range h {
		hook.AfterSubspaceGroupDeleted(ctx, subspaceID, groupName)
	}
}

// AfterPermissionSet implements SubspacesHook
func (h MultiSubspacesHooks) AfterPermissionSet(ctx sdk.Context, subspaceID uint64, target string, permissions Permission) {
	for _, hook := range h {
		hook.AfterPermissionSet(ctx, subspaceID, target, permissions)
	}
}

// AfterPermissionRemoved implements SubspacesHook
func (h MultiSubspacesHooks) AfterPermissionRemoved(ctx sdk.Context, subspaceID uint64, target string) {
	for _, hook := range h {
		hook.AfterPermissionRemoved(ctx, subspaceID, target)
	}
}
