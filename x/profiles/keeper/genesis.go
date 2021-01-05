package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		k.GetProfiles(ctx),
		k.GetDTagTransferRequests(ctx),
		k.GetParams(ctx),
	)
}

// InitGenesis initializes the chain state based on the given GenesisState
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) []abci.ValidatorUpdate {
	k.SetParams(ctx, data.Params)

	for _, profile := range data.Profiles {
		err := k.ValidateProfile(ctx, profile)
		if err != nil {
			panic(err)
		}

		err = k.StoreProfile(ctx, profile)
		if err != nil {
			panic(err)
		}
	}

	for _, request := range data.DtagTransferRequests {
		err := k.SaveDTagTransferRequest(ctx, request)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
