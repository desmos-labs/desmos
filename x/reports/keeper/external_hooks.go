package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/reports/types"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

type Hooks struct {
	k Keeper
}

var (
	_ subspacestypes.SubspacesHooks = Hooks{}

	// TODO: Add the posts hooks to delete all the reports of a post once that's deleted
)

// Hooks creates a new reports hooks
func (k Keeper) Hooks() Hooks { return Hooks{k} }

// AfterSubspaceSaved implements subspacestypes.Hooks
func (h Hooks) AfterSubspaceSaved(ctx sdk.Context, subspaceID uint64) {
	// Create the initial reason and report id
	h.k.SetNextReasonID(ctx, subspaceID, 1)
	h.k.SetNextReportID(ctx, subspaceID, 1)
}

// AfterSubspaceDeleted implements subspacestypes.Hooks
func (h Hooks) AfterSubspaceDeleted(ctx sdk.Context, subspaceID uint64) {
	// Delete the reason id key
	h.k.DeleteNextReasonID(ctx, subspaceID)

	// Delete all the reasons related to this subspace
	h.k.IterateSubspaceReasons(ctx, subspaceID, func(_ int64, reason types.Reason) (stop bool) {
		h.k.DeleteReason(ctx, reason.SubspaceID, reason.ID)
		return false
	})

	// Delete the report id key
	h.k.DeleteNextReportID(ctx, subspaceID)

	// Delete all the reports related to this subspace
	h.k.IterateSubspaceReports(ctx, subspaceID, func(_ int64, report types.Report) (stop bool) {
		h.k.DeleteReport(ctx, report.SubspaceID, report.ID)
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
