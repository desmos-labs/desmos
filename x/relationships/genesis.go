package relationships

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/relationships/keeper"
	"github.com/desmos-labs/desmos/x/relationships/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return types.GenesisState{
		UsersRelationships: k.GetUsersRelationships(ctx),
		UsersBlocks:        k.GetUsersBlocks(ctx),
	}
}

// InitGenesis initializes the chain state based on the given GenesisState
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) []abci.ValidatorUpdate {
	for userAddr, relationships := range data.UsersRelationships {
		addr, err := sdk.AccAddressFromBech32(userAddr)
		if err != nil {
			panic(err)
		}
		for _, receiver := range relationships {
			err := k.StoreRelationship(ctx, addr, receiver)
			if err != nil {
				panic(err)
			}
		}
	}

	for _, userBlock := range data.UsersBlocks {
		err := k.SaveUserBlock(ctx, userBlock)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
