package simulation

// DONTCOVER

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/desmos-labs/desmos/x/fees/types"
)

// RandomizedGenState generates a random GenesisState for fees
func RandomizedGenState(simState *module.SimulationState) {
	feesGenesis := types.NewGenesisState(types.DefaultParams())
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(feesGenesis)

	fmt.Printf("Selected randomly generated fees parameters:\n%s\n%s\n",
		codec.MustMarshalJSONIndent(simState.Cdc, feesGenesis.Params.FeeDenom),
		codec.MustMarshalJSONIndent(simState.Cdc, feesGenesis.Params.MinFees),
	)
}
