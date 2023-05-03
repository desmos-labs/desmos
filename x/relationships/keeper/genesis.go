package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v5/x/relationships/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		k.GetAllRelationships(ctx),
		k.GetAllUserBlocks(ctx),
	)
}

// GetAllRelationships returns the list of all stored relationships
func (k Keeper) GetAllRelationships(ctx sdk.Context) []types.Relationship {
	var relationships []types.Relationship
	k.IterateRelationships(ctx, func(index int64, relationship types.Relationship) (stop bool) {
		relationships = append(relationships, relationship)
		return false
	})
	return relationships
}

// GetAllUserBlocks returns the list of all stored blocks
func (k Keeper) GetAllUserBlocks(ctx sdk.Context) []types.UserBlock {
	var blocks []types.UserBlock
	k.IterateUsersBlocks(ctx, func(_ int64, block types.UserBlock) (stop bool) {
		blocks = append(blocks, block)
		return false
	})
	return blocks
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
