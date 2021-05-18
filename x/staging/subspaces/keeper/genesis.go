package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		k.GetAllSubspaces(ctx),
	)
}

// InitGenesis initializes the chain state based on the given GenesisState
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {

	for _, subspace := range data.Subspaces {
		if err := subspace.Validate(); err != nil {
			panic(err)
		}
		if k.DoesSubspaceExists(ctx, subspace.ID) {
			panic(fmt.Sprintf("The subspace %s already exist", subspace.ID))
		}
		k.SaveSubspace(ctx, subspace)
	}
}
