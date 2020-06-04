package profile

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	return GenesisState{
		Profiles:             k.GetProfiles(ctx),
		NameSurnameLenParams: k.GetNameSurnameLenParams(ctx),
		MonikerLenParams:     k.GetMonikerLenParams(ctx),
		BioLenParams:         k.GetBioLenParams(ctx),
	}
}

// InitGenesis initializes the chain state based on the given GenesisState
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, profile := range data.Profiles {
		if err := keeper.SaveProfile(ctx, profile); err != nil {
			panic(err)
		}
	}

	keeper.SetNameSurnameLenParams(ctx, data.NameSurnameLenParams)
	keeper.SetMonikerLenParams(ctx, data.MonikerLenParams)
	keeper.SetBioLenParams(ctx, data.BioLenParams)

	return nil
}
