package reports

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	posts "github.com/desmos-labs/desmos/x/posts/types"
	reportsKeeper "github.com/desmos-labs/desmos/x/reports/keeper"
	reportsTypes "github.com/desmos-labs/desmos/x/reports/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func ExportGenesis(ctx sdk.Context, k reportsKeeper.Keeper) reportsTypes.GenesisState {
	return reportsTypes.GenesisState{
		Reports: k.GetReportsMap(ctx),
	}
}

// InitGenesis initializes the chain state based on the given GenesisState
func InitGenesis(ctx sdk.Context, keeper reportsKeeper.Keeper, data reportsTypes.GenesisState) []abci.ValidatorUpdate {
	for postID, reports := range data.Reports {
		postID, err := posts.ParsePostID(postID)
		if err != nil {
			panic(err)
		}
		for _, report := range reports {
			keeper.SaveReport(ctx, postID, report)
		}
	}

	return []abci.ValidatorUpdate{}
}
