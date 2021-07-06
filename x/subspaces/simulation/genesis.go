package simulation

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/desmos-labs/desmos/x/subspaces/types"
)

// RandomizeGenState generates a random GenesisState for subspaces
func RandomizeGenState(simState *module.SimulationState) {
	subspaces := randomSubspaces(simState)
	subspacesGenesis := types.NewGenesisState(subspaces, nil, nil, nil)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(subspacesGenesis)
}

// randomSubspaces returns randomly generated genesis subspaces
func randomSubspaces(simState *module.SimulationState) (subspaces []types.Subspace) {
	subspaces = make([]types.Subspace, len(subspacesIds))
	for index := 0; index < len(subspaces); index++ {
		subspaceData := RandomSubspaceData(simState.Rand, simState.Accounts)
		// this is to ensure to not repeat the ID and get a mismatch on the owner while validating it
		subspaceData.Subspace.ID = subspacesIds[index]
		subspaces[index] = subspaceData.Subspace
	}

	return subspaces
}
