package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

// Implements StakingHooks interface
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

// AfterSubspaceGroupSaved - call if hook is registered
func (k Keeper) AfterSubspaceGroupSaved(ctx sdk.Context, subspaceID uint64, groupName string) {
	if k.hooks != nil {
		k.hooks.AfterSubspaceGroupSaved(ctx, subspaceID, groupName)
	}
}

// AfterSubspaceGroupMemberAdded - call if hook is registered
func (k Keeper) AfterSubspaceGroupMemberAdded(ctx sdk.Context, subspaceID uint64, groupName string, user sdk.AccAddress) {
	if k.hooks != nil {
		k.hooks.AfterSubspaceGroupMemberAdded(ctx, subspaceID, groupName, user)
	}
}

// AfterSubspaceGroupMemberRemoved - call if hook is registered
func (k Keeper) AfterSubspaceGroupMemberRemoved(ctx sdk.Context, subspaceID uint64, groupName string, user sdk.AccAddress) {
	if k.hooks != nil {
		k.hooks.AfterSubspaceGroupMemberRemoved(ctx, subspaceID, groupName, user)
	}
}

// AfterSubspaceGroupDeleted - call if hook is registered
func (k Keeper) AfterSubspaceGroupDeleted(ctx sdk.Context, subspaceID uint64, groupName string) {
	if k.hooks != nil {
		k.hooks.AfterSubspaceGroupDeleted(ctx, subspaceID, groupName)
	}
}

// AfterPermissionSet - call if hook is registered
func (k Keeper) AfterPermissionSet(ctx sdk.Context, subspaceID uint64, target string, permissions types.Permission) {
	if k.hooks != nil {
		k.hooks.AfterPermissionSet(ctx, subspaceID, target, permissions)
	}
}

// AfterPermissionRemoved - call if hook is registered
func (k Keeper) AfterPermissionRemoved(ctx sdk.Context, subspaceID uint64, target string) {
	if k.hooks != nil {
		k.hooks.AfterPermissionRemoved(ctx, subspaceID, target)
	}
}
