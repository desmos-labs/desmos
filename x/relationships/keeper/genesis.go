package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/relationships/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		k.GetAllRelationships(ctx),
		k.GetUsersBlocks(ctx),
	)
}

// InitGenesis initializes the chain state based on the given GenesisState
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) []abci.ValidatorUpdate {
	for _, relationship := range data.Relationships {
		if err := k.StoreRelationship(ctx, relationship); err != nil {
			panic(err)
		}
	}

	for _, userBlock := range data.Blocks {
		if err := k.SaveUserBlock(ctx, userBlock); err != nil {
			panic(err)
		}
	}

	return nil
}
