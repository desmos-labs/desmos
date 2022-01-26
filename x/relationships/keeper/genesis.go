package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/desmos-labs/desmos/v2/x/relationships/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		k.GetAllRelationships(ctx),
		k.GetAllUsersBlocks(ctx),
	)
}

// InitGenesis initializes the chain state based on the given GenesisState
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) []abci.ValidatorUpdate {
	// Store the relationships
	for _, relationship := range data.Relationships {
		err := k.SaveRelationship(ctx, relationship)
		if err != nil {
			panic(err)
		}
	}

	// Store the user blocks
	for _, userBlock := range data.Blocks {
		err := k.SaveUserBlock(ctx, userBlock)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
