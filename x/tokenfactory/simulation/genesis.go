package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	subspacessim "github.com/desmos-labs/desmos/v5/x/subspaces/simulation"
	subspacestypes "github.com/desmos-labs/desmos/v5/x/subspaces/types"
	"github.com/desmos-labs/desmos/v5/x/tokenfactory/types"
	tokenfactorytypes "github.com/osmosis-labs/osmosis/v15/x/tokenfactory/types"
)

// RandomizeGenState generates a random GenesisState for x/tokenfactory
func RandomizeGenState(simState *module.SimulationState) {
	// Read the subspaces data
	subspacesGenesisBz := simState.GenState[subspacestypes.ModuleName]
	var subspacesGenesis subspacestypes.GenesisState
	simState.Cdc.MustUnmarshalJSON(subspacesGenesisBz, &subspacesGenesis)

	genesis := &tokenfactorytypes.GenesisState{
		Params:        types.ToOsmosisTokenFactoryParams(types.NewParams(nil)),
		FactoryDenoms: randomFactoryDenoms(simState.Rand, subspacesGenesis.Subspaces),
	}

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(genesis)
}

// randomFactoryDenoms generates a list of random factory denoms
func randomFactoryDenoms(r *rand.Rand, subspaces []subspacestypes.Subspace) []tokenfactorytypes.GenesisDenom {
	if len(subspaces) == 0 {
		return nil
	}

	denomsNumber := r.Intn(20)
	denoms := make([]tokenfactorytypes.GenesisDenom, denomsNumber)
	for i := 0; i < denomsNumber; i++ {
		subspace := subspacessim.RandomSubspace(r, subspaces)

		denom, err := tokenfactorytypes.GetTokenDenom(subspace.Treasury, simtypes.RandStringOfLength(r, 6))
		if err != nil {
			panic(err)
		}

		denoms[i] = tokenfactorytypes.GenesisDenom{
			Denom: denom,
			AuthorityMetadata: tokenfactorytypes.DenomAuthorityMetadata{
				Admin: subspace.Treasury,
			},
		}
	}

	return denoms
}
