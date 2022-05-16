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

	AfterSubspaceSectionSaved(ctx sdk.Context, subspaceID uint64, sectionID uint32)   // Must be called when a subspace section is saved
	AfterSubspaceSectionDeleted(ctx sdk.Context, subspaceID uint64, sectionID uint32) // Must be called when a subspace section is deleted

	AfterSubspaceGroupSaved(ctx sdk.Context, subspaceID uint64, groupID uint32)                              // Must be called when a subspace group is created
	AfterSubspaceGroupMemberAdded(ctx sdk.Context, subspaceID uint64, groupID uint32, user sdk.AccAddress)   // Must be called when a user is added to a group
	AfterSubspaceGroupMemberRemoved(ctx sdk.Context, subspaceID uint64, groupID uint32, user sdk.AccAddress) // Must be called when a user is removed from a group
	AfterSubspaceGroupDeleted(ctx sdk.Context, subspaceID uint64, groupID uint32)                            // Must be called when a subspace group is deleted

	AfterUserPermissionSet(ctx sdk.Context, subspaceID uint64, user sdk.AccAddress, permissions Permission) // Must be called when a permission is set for a user
	AfterUserPermissionRemoved(ctx sdk.Context, subspaceID uint64, user sdk.AccAddress)                     // Must be called when a permission is removed for a user
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

// AfterSubspaceSectionSaved implements SubspacesHooks
func (h MultiSubspacesHooks) AfterSubspaceSectionSaved(ctx sdk.Context, subspaceID uint64, sectionID uint32) {
	for _, hook := range h {
		hook.AfterSubspaceSectionSaved(ctx, subspaceID, sectionID)
	}
}

// AfterSubspaceSectionDeleted implements SubspacesHooks
func (h MultiSubspacesHooks) AfterSubspaceSectionDeleted(ctx sdk.Context, subspaceID uint64, sectionID uint32) {
	for _, hook := range h {
		hook.AfterSubspaceSectionDeleted(ctx, subspaceID, sectionID)
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
func (h MultiSubspacesHooks) AfterUserPermissionSet(ctx sdk.Context, subspaceID uint64, user sdk.AccAddress, permissions Permission) {
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
