package simulation

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	types2 "github.com/desmos-labs/desmos/x/subspaces/types"
)

// RandomizeGenState generates a random GenesisState for subspaces
func RandomizeGenState(simState *module.SimulationState) {
	subspaces := randomSubspaces(simState)
	subspacesGenesis := types2.NewGenesisState(subspaces, nil, nil, nil)
	simState.GenState[types2.ModuleName] = simState.Cdc.MustMarshalJSON(subspacesGenesis)
}

// randomSubspaces returns randomly generated genesis account
func randomSubspaces(simState *module.SimulationState) (subspaces []types2.Subspace) {
	subspacesNumber := simState.Rand.Intn(100)

	subspaces = make([]types2.Subspace, subspacesNumber)
	for index := 0; index < subspacesNumber; index++ {
		subspaceData := RandomSubspaceData(simState.Rand, simState.Accounts)
		subspaces[index] = subspaceData.Subspace
	}

	return subspaces
}
