package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

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
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	// Store the relationships
	for _, relationship := range data.Relationships {
		k.SaveRelationship(ctx, relationship)
	}

	// Store the user blocks
	for _, userBlock := range data.Blocks {
		k.SaveUserBlock(ctx, userBlock)
	}
}
