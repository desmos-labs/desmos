package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/relationships/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		k.GetAllRelationships(ctx),
		k.GetAllUsersBlocks(ctx),
	)
}

// InitGenesis initializes the chain state based on the given GenesisState
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	for _, relationship := range data.Relationships {
		err := k.SaveRelationship(ctx, relationship)
		if err != nil {
			panic(err)
		}
	}

	for _, userBlock := range data.Blocks {
		err := k.SaveUserBlock(ctx, userBlock)
		if err != nil {
			panic(err)
		}
	}
}
