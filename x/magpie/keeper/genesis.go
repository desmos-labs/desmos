package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/desmos-labs/desmos/x/magpie/types"
)

// InitGenesis initializes the chain state based on the given GenesisState
// noinspection GoUnhandledErrorResult
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) []abci.ValidatorUpdate {
	if err := k.SetDefaultSessionLength(ctx, data.DefaultSessionLength); err != nil {
		panic(err)
	}

	for _, session := range data.Sessions {
		k.SaveSession(ctx, session)
	}

	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the GenesisState associated with the given context
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		k.GetDefaultSessionLength(ctx),
		k.GetSessions(ctx),
	)
}
