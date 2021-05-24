package simulation

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
)

// RandomizeGenState generates a random GenesisState for subspaces
func RandomizeGenState(simState *module.SimulationState) {
	subspaces := randomSubspaces(simState)
	params := types.NewParams(RandomNameParams(simState.Rand))

	subspacesGenesis := types.NewGenesisState(subspaces, params)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(subspacesGenesis)
}

// randomSubspaces returns randomly generated genesis account
func randomSubspaces(simState *module.SimulationState) (subspaces []types.Subspace) {
	subspacesNumber := simState.Rand.Intn(100)

	subspaces = make([]types.Subspace, subspacesNumber)
	for index := 0; index < subspacesNumber; index++ {
		subspaceData := RandomSubspaceData(simState.Rand, simState.Accounts)
		subspaces[index] = subspaceData.subspace
	}

	return subspaces
}
