package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/reports/types"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		k.getSubspaceDataEntries(ctx),
		k.getAllReasons(ctx),
		k.getAllReports(ctx),
		k.GetParams(ctx),
	)
}

// getSubspaceDataEntries returns the subspaces data entries stored in the given context
func (k Keeper) getSubspaceDataEntries(ctx sdk.Context) []types.SubspaceDataEntry {
	var entries []types.SubspaceDataEntry
	k.sk.IterateSubspaces(ctx, func(subspace subspacestypes.Subspace) (stop bool) {
		nextReasonID, err := k.GetNextReasonID(ctx, subspace.ID)
		if err != nil {
			panic(err)
		}

		nextReportID, err := k.GetNextReportID(ctx, subspace.ID)
		if err != nil {
			panic(err)
		}

		entries = append(entries, types.NewSubspacesDataEntry(
			subspace.ID,
			nextReasonID,
			nextReportID,
		))

		return false
	})
	return entries
}

// getAllReasons returns all the reasons stored inside the given context
func (k Keeper) getAllReasons(ctx sdk.Context) []types.Reason {
	var reasons []types.Reason
	k.IterateReasons(ctx, func(index int64, reason types.Reason) (stop bool) {
		reasons = append(reasons, reason)
		return false
	})
	return reasons
}

// getAllReports returns all the reports stored inside the given context
func (k Keeper) getAllReports(ctx sdk.Context) []types.Report {
	var reports []types.Report
	k.IterateReports(ctx, func(index int64, report types.Report) (stop bool) {
		reports = append(reports, report)
		return false
	})
	return reports
}

// --------------------------------------------------------------------------------------------------------------------

// InitGenesis initializes the chain state based on the given GenesisState
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	// Initialize the subspaces data
	for _, entry := range data.SubspacesData {
		k.SetNextReasonID(ctx, entry.SubspaceID, entry.ReasonID)
		k.SetNextReportID(ctx, entry.SubspaceID, entry.ReportID)
	}

	// Initialize all the reasons
	for _, reason := range data.Reasons {
		k.SaveReason(ctx, reason)
	}

	// Initialize all the reports
	for _, report := range data.Reports {
		k.SaveReport(ctx, report)
	}

	// Initialize the params
	k.SetParams(ctx, data.Params)
}
