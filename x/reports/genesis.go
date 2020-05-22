package reports

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts"
	abci "github.com/tendermint/tendermint/abci/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	return GenesisState{
		Reports: k.GetReportsMap(ctx),
	}
}

// InitGenesis initializes the chain state based on the given GenesisState
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for postID, reports := range data.Reports {
		postID := posts.PostID(postID)
		if !postID.Valid() {
			panic(fmt.Errorf("invalid postID: %s", postID))
		}
		for _, report := range reports {
			keeper.SaveReport(ctx, postID, report)
		}
	}

	return []abci.ValidatorUpdate{}
}
