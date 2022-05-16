package keeper

// DONTCOVER

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// Implements SubspacesHooks interface
var _ types.SubspacesHooks = Keeper{}

// AfterSubspaceSaved - call if hook is registered
func (k Keeper) AfterSubspaceSaved(ctx sdk.Context, subspaceID uint64) {
	if k.hooks != nil {
		k.hooks.AfterSubspaceSaved(ctx, subspaceID)
	}
}

// AfterSubspaceDeleted - call if hook is registered
func (k Keeper) AfterSubspaceDeleted(ctx sdk.Context, subspaceID uint64) {
	if k.hooks != nil {
		k.hooks.AfterSubspaceDeleted(ctx, subspaceID)
	}
}

// AfterSubspaceSectionSaved - call if hook is registered
func (k Keeper) AfterSubspaceSectionSaved(ctx sdk.Context, subspaceID uint64, sectionID uint32) {
	if k.hooks != nil {
		k.hooks.AfterSubspaceSectionSaved(ctx, subspaceID, sectionID)
	}
}

// AfterSubspaceSectionDeleted - call if hook is registered
func (k Keeper) AfterSubspaceSectionDeleted(ctx sdk.Context, subspaceID uint64, sectionID uint32) {
	if k.hooks != nil {
		k.hooks.AfterSubspaceSectionDeleted(ctx, subspaceID, sectionID)
	}
}

// AfterSubspaceGroupSaved - call if hook is registered
func (k Keeper) AfterSubspaceGroupSaved(ctx sdk.Context, subspaceID uint64, groupID uint32) {
	if k.hooks != nil {
		k.hooks.AfterSubspaceGroupSaved(ctx, subspaceID, groupID)
	}
}

// AfterSubspaceGroupMemberAdded - call if hook is registered
func (k Keeper) AfterSubspaceGroupMemberAdded(ctx sdk.Context, subspaceID uint64, groupID uint32, user sdk.AccAddress) {
	if k.hooks != nil {
		k.hooks.AfterSubspaceGroupMemberAdded(ctx, subspaceID, groupID, user)
	}
}

// AfterSubspaceGroupMemberRemoved - call if hook is registered
func (k Keeper) AfterSubspaceGroupMemberRemoved(ctx sdk.Context, subspaceID uint64, groupID uint32, user sdk.AccAddress) {
	if k.hooks != nil {
		k.hooks.AfterSubspaceGroupMemberRemoved(ctx, subspaceID, groupID, user)
	}
}

// AfterSubspaceGroupDeleted - call if hook is registered
func (k Keeper) AfterSubspaceGroupDeleted(ctx sdk.Context, subspaceID uint64, groupID uint32) {
	if k.hooks != nil {
		k.hooks.AfterSubspaceGroupDeleted(ctx, subspaceID, groupID)
	}
}

// AfterUserPermissionSet - call if hook is registered
func (k Keeper) AfterUserPermissionSet(ctx sdk.Context, subspaceID uint64, user sdk.AccAddress, permissions types.Permission) {
	if k.hooks != nil {
		k.hooks.AfterUserPermissionSet(ctx, subspaceID, user, permissions)
	}
}

// AfterUserPermissionRemoved - call if hook is registered
func (k Keeper) AfterUserPermissionRemoved(ctx sdk.Context, subspaceID uint64, user sdk.AccAddress) {
	if k.hooks != nil {
		k.hooks.AfterUserPermissionRemoved(ctx, subspaceID, user)
	}
}
