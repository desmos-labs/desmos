package profile

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	return GenesisState{
		Accounts: k.GetAccounts(ctx),
	}
}

// InitGenesis initializes the chain state based on the given GenesisState
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {

	for _, account := range data.Accounts {
		if err := keeper.SaveAccount(ctx, account); err != nil {
			panic(err)
		}
	}

	return nil
}
