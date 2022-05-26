package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/posts/types"

	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// Hooks represents a wrapper struct
type Hooks struct {
	k Keeper
}

var _ subspacestypes.SubspacesHooks = Hooks{}

// Hooks creates new subspaces hooks
func (k Keeper) Hooks() Hooks { return Hooks{k} }

// AfterSubspaceSaved implements subspacestypes.Hooks
func (h Hooks) AfterSubspaceSaved(ctx sdk.Context, subspaceID uint64) {
	// Create the initial post it
	h.k.SetNextPostID(ctx, subspaceID, 1)
}

// AfterSubspaceDeleted implements subspacestypes.Hooks
func (h Hooks) AfterSubspaceDeleted(ctx sdk.Context, subspaceID uint64) {
	// Delete the post id key
	h.k.DeleteNextPostID(ctx, subspaceID)

	// Delete all the posts
	h.k.IterateSubspacePosts(ctx, subspaceID, func(_ int64, post types.Post) (stop bool) {
		h.k.DeletePost(ctx, post.SubspaceID, post.ID)
		return false
	})
}

// AfterSubspaceGroupSaved implements subspacestypes.Hooks
func (h Hooks) AfterSubspaceGroupSaved(sdk.Context, uint64, uint32) {}

// AfterSubspaceGroupMemberAdded implements subspacestypes.Hooks
func (h Hooks) AfterSubspaceGroupMemberAdded(sdk.Context, uint64, uint32, sdk.AccAddress) {}

// AfterSubspaceGroupMemberRemoved implements subspacestypes.Hooks
func (h Hooks) AfterSubspaceGroupMemberRemoved(sdk.Context, uint64, uint32, sdk.AccAddress) {}

// AfterSubspaceGroupDeleted implements subspacestypes.Hooks
func (h Hooks) AfterSubspaceGroupDeleted(sdk.Context, uint64, uint32) {}

// AfterUserPermissionSet implements subspacestypes.Hooks
func (h Hooks) AfterUserPermissionSet(sdk.Context, uint64, sdk.AccAddress, subspacestypes.Permission) {
}

// AfterUserPermissionRemoved implements subspacestypes.Hooks
func (h Hooks) AfterUserPermissionRemoved(sdk.Context, uint64, sdk.AccAddress) {}
