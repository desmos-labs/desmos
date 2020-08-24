package profiles

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profiles/keeper"
	"github.com/desmos-labs/desmos/x/profiles/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return types.GenesisState{
		Profiles:           k.GetProfiles(ctx),
		Params:             k.GetParams(ctx),
		UsersRelationships: k.GetUsersRelationshipsMap(ctx),
	}
}

// InitGenesis initializes the chain state based on the given GenesisState
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) []abci.ValidatorUpdate {
	k.SetParams(ctx, data.Params)

	for _, profile := range data.Profiles {
		if err := keeper.ValidateProfile(ctx, k, profile); err != nil {
			panic(err)
		}
		if err := k.SaveProfile(ctx, profile); err != nil {
			panic(err)
		}
	}

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

	return nil
}
