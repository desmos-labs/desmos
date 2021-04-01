package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/links/types"
)

// InitGenesis initializes the chain state based on the given GenesisState
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {

	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	genesis := types.NewGenesisState(
		k.GetPort(ctx),
		k.GetLinks(ctx),
	)

	return genesis
}
