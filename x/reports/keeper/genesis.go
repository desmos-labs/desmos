package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/reports/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	reports, err := k.GetReports(ctx)
	if err != nil {
		panic(err)
	}

	return types.NewGenesisState(
		reports,
	)
}

// InitGenesis initializes the chain state based on the given GenesisState
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) []abci.ValidatorUpdate {
	for _, report := range data.Reports {
		err := k.SaveReport(ctx, report)
		if err != nil {
			panic(err)
		}
	}

	return []abci.ValidatorUpdate{}
}
