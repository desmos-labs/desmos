package magpie

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/magpie/keeper"
	"github.com/desmos-labs/desmos/x/magpie/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return types.GenesisState{
		DefaultSessionLength: k.GetDefaultSessionLength(ctx),
		Sessions:             k.GetSessions(ctx),
	}
}

// InitGenesis initializes the chain state based on the given GenesisState
// noinspection GoUnhandledErrorResult
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data types.GenesisState) []abci.ValidatorUpdate {
	if err := keeper.SetDefaultSessionLength(ctx, data.DefaultSessionLength); err != nil {
		panic(err)
	}

	for _, session := range data.Sessions {
		keeper.SaveSession(ctx, session)
	}

	return []abci.ValidatorUpdate{}
}
