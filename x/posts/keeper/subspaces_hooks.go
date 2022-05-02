package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

var _ subspacestypes.SubspacesHooks = &Keeper{}

// AfterSubspaceSaved implements subspacestypes.Hooks
func (k Keeper) AfterSubspaceSaved(ctx sdk.Context, subspaceID uint64) {
	// Create the initial post it
	k.SetPostID(ctx, subspaceID, 1)
}

// AfterSubspaceDeleted implements subspacestypes.Hooks
func (k Keeper) AfterSubspaceDeleted(ctx sdk.Context, subspaceID uint64) {
	// Delete the post id key
	k.DeletePostID(ctx, subspaceID)

	// Delete all the posts
	posts := k.GetSubspacePosts(ctx, subspaceID)
	for _, post := range posts {
		k.DeletePost(ctx, post.SubspaceID, post.ID)
	}
}

// AfterSubspaceGroupSaved implements subspacestypes.Hooks
func (k Keeper) AfterSubspaceGroupSaved(sdk.Context, uint64, uint32) {}

// AfterSubspaceGroupMemberAdded implements subspacestypes.Hooks
func (k Keeper) AfterSubspaceGroupMemberAdded(sdk.Context, uint64, uint32, sdk.AccAddress) {}

// AfterSubspaceGroupMemberRemoved implements subspacestypes.Hooks
func (k Keeper) AfterSubspaceGroupMemberRemoved(sdk.Context, uint64, uint32, sdk.AccAddress) {}

// AfterSubspaceGroupDeleted implements subspacestypes.Hooks
func (k Keeper) AfterSubspaceGroupDeleted(sdk.Context, uint64, uint32) {}

// AfterUserPermissionSet implements subspacestypes.Hooks
func (k Keeper) AfterUserPermissionSet(sdk.Context, uint64, sdk.AccAddress, subspacestypes.Permission) {
}

// AfterUserPermissionRemoved implements subspacestypes.Hooks
func (k Keeper) AfterUserPermissionRemoved(sdk.Context, uint64, sdk.AccAddress) {}
