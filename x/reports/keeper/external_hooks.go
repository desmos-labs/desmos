package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	poststypes "github.com/desmos-labs/desmos/v3/x/posts/types"

	"github.com/desmos-labs/desmos/v3/x/reports/types"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

type Hooks struct {
	k Keeper
}

var (
	_ subspacestypes.SubspacesHooks = Hooks{}
	_ poststypes.PostsHooks         = Hooks{}
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
	h.k.IterateSubspaceReasons(ctx, subspaceID, func(reason types.Reason) (stop bool) {
		h.k.DeleteReason(ctx, reason.SubspaceID, reason.ID)
		return false
	})

	// Delete the report id key
	h.k.DeleteNextReportID(ctx, subspaceID)

	// Delete all the reports related to this subspace
	h.k.IterateSubspaceReports(ctx, subspaceID, func(report types.Report) (stop bool) {
		h.k.DeleteReport(ctx, report.SubspaceID, report.ID)
		return false
	})
}

// AfterSubspaceGroupSaved implements subspacestypes.Hooks
func (h Hooks) AfterSubspaceGroupSaved(sdk.Context, uint64, uint32) {}

// AfterSubspaceGroupDeleted implements subspacestypes.Hooks
func (h Hooks) AfterSubspaceGroupDeleted(sdk.Context, uint64, uint32) {}

// AfterSubspaceSectionSaved implements subspacestypes.Hooks
func (h Hooks) AfterSubspaceSectionSaved(sdk.Context, uint64, uint32) {}

// AfterSubspaceSectionDeleted implements subspacestypes.Hooks
func (h Hooks) AfterSubspaceSectionDeleted(sdk.Context, uint64, uint32) {}

// AfterSubspaceGroupMemberAdded implements subspacestypes.Hooks
func (h Hooks) AfterSubspaceGroupMemberAdded(sdk.Context, uint64, uint32, string) {
}

// AfterSubspaceGroupMemberRemoved implements subspacestypes.Hooks
func (h Hooks) AfterSubspaceGroupMemberRemoved(sdk.Context, uint64, uint32, string) {
}

// AfterUserPermissionSet implements subspacestypes.Hooks
func (h Hooks) AfterUserPermissionSet(sdk.Context, uint64, uint32, string, subspacestypes.Permission) {
}

// AfterUserPermissionRemoved implements subspacestypes.Hooks
func (h Hooks) AfterUserPermissionRemoved(sdk.Context, uint64, uint32, string) {
}

// AfterPostSaved implements poststypes.PostsHooks
func (h Hooks) AfterPostSaved(sdk.Context, uint64, uint64) {}

// AfterPostDeleted implements poststypes.PostsHooks
func (h Hooks) AfterPostDeleted(ctx sdk.Context, subspaceID uint64, postID uint64) {
	// Delete all the reports related to this post
	h.k.IteratePostReports(ctx, subspaceID, postID, func(report types.Report) (stop bool) {
		h.k.DeleteReport(ctx, report.SubspaceID, report.ID)
		return false
	})
}

// AfterAttachmentSaved implements poststypes.PostsHooks
func (h Hooks) AfterAttachmentSaved(sdk.Context, uint64, uint64, uint32) {}

// AfterAttachmentDeleted implements poststypes.PostsHooks
func (h Hooks) AfterAttachmentDeleted(sdk.Context, uint64, uint64, uint32) {}

// AfterPollAnswerSaved implements poststypes.PostsHooks
func (h Hooks) AfterPollAnswerSaved(sdk.Context, uint64, uint64, uint32, string) {}

// AfterPollAnswerDeleted implements poststypes.PostsHooks
func (h Hooks) AfterPollAnswerDeleted(sdk.Context, uint64, uint64, uint32, string) {}

// AfterPollVotingPeriodEnded implements poststypes.PostsHooks
func (h Hooks) AfterPollVotingPeriodEnded(sdk.Context, uint64, uint64, uint32) {}
