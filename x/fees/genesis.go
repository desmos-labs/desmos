package fees

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/fees/keeper"
	"github.com/desmos-labs/desmos/x/fees/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return types.GenesisState{
		Params: k.GetParams(ctx),
	}
}

// InitGenesis initializes the chain state based on the given GenesisState
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) []abci.ValidatorUpdate {
	k.SetParams(ctx, data.Params)

	return []abci.ValidatorUpdate{}
}
