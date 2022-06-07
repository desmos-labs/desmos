package types

// DONTCOVER

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Event Hooks
// These can be utilized to communicate between a reports keeper and another
// keeper which must take particular actions when reports/reasons change
// state. The second keeper must implement this interface, which then the
// reports keeper can call.

// ReportsHooks event hooks for posts objects (noalias)
type ReportsHooks interface {
	AfterReportSaved(ctx sdk.Context, subspaceID uint64, reportID uint64)   // Must be called when a report is saved
	AfterReportDeleted(ctx sdk.Context, subspaceID uint64, reportID uint64) // Must be called when a report is deleted

	AfterReasonSaved(ctx sdk.Context, subspaceID uint64, reasonID uint32)   // Must be called when a reason is saved
	AfterReasonDeleted(ctx sdk.Context, subspaceID uint64, reasonID uint32) // Must be called when a reason is deleted
}

// --------------------------------------------------------------------------------------------------------------------

// MultiReportsHooks combines multiple subspaces hooks, all hook functions are run in array sequence
type MultiReportsHooks []ReportsHooks

func NewMultiReportsHooks(hooks ...ReportsHooks) MultiReportsHooks {
	return hooks
}

// AfterReportSaved implements ReportsHooks
func (h MultiReportsHooks) AfterReportSaved(ctx sdk.Context, subspaceID uint64, reportID uint64) {
	for _, hook := range h {
		hook.AfterReportSaved(ctx, subspaceID, reportID)
	}
}

// AfterReportDeleted implements ReportsHooks
func (h MultiReportsHooks) AfterReportDeleted(ctx sdk.Context, subspaceID uint64, reportID uint64) {
	for _, hook := range h {
		hook.AfterReportDeleted(ctx, subspaceID, reportID)
	}
}

// AfterReasonSaved implements ReportsHooks
func (h MultiReportsHooks) AfterReasonSaved(ctx sdk.Context, subspaceID uint64, reasonID uint32) {
	for _, hook := range h {
		hook.AfterReasonSaved(ctx, subspaceID, reasonID)
	}
}

// AfterReasonDeleted implements ReportsHooks
func (h MultiReportsHooks) AfterReasonDeleted(ctx sdk.Context, subspaceID uint64, reasonID uint32) {
	for _, hook := range h {
		hook.AfterReasonDeleted(ctx, subspaceID, reasonID)
	}
}
