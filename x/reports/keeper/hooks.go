package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/reports/types"
)

// Implement ReportsHooks interface
var _ types.ReportsHooks = Keeper{}

// AfterReportSaved implements types.ReportsHooks
func (k Keeper) AfterReportSaved(ctx sdk.Context, subspaceID uint64, reportID uint64) {
	if k.hooks != nil {
		k.hooks.AfterReportSaved(ctx, subspaceID, reportID)
	}
}

// AfterReportDeleted implements types.ReportsHooks
func (k Keeper) AfterReportDeleted(ctx sdk.Context, subspaceID uint64, reportID uint64) {
	if k.hooks != nil {
		k.hooks.AfterReportDeleted(ctx, subspaceID, reportID)
	}
}

// AfterReasonSaved implements types.ReportsHooks
func (k Keeper) AfterReasonSaved(ctx sdk.Context, subspaceID uint64, reasonID uint32) {
	if k.hooks != nil {
		k.hooks.AfterReasonSaved(ctx, subspaceID, reasonID)
	}
}

// AfterReasonDeleted implements types.ReportsHooks
func (k Keeper) AfterReasonDeleted(ctx sdk.Context, subspaceID uint64, reasonID uint32) {
	if k.hooks != nil {
		k.hooks.AfterReasonDeleted(ctx, subspaceID, reasonID)
	}
}
