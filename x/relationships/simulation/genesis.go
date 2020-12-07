package simulation

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/desmos-labs/desmos/x/relationships/types"
)

// RandomizedGenState generates a random GenesisState for profile
func RandomizedGenState(simsState *module.SimulationState) {
	profileGenesis := types.NewGenesisState(
		randomRelationships(simsState),
		randomUsersBlocks(simsState),
	)

	simsState.GenState[types.ModuleName] = simsState.Cdc.MustMarshalJSON(profileGenesis)
}

// randomRelationships returns randomly generated genesis relationships and their associated users - IDs map
func randomRelationships(simState *module.SimulationState) []types.Relationship {
	relationshipsNumber := simState.Rand.Intn(simtypes.RandIntBetween(simState.Rand, 1, 30))

	relationships := make([]types.Relationship, relationshipsNumber)
	for index := 0; index < relationshipsNumber; index++ {
		sender, _ := simtypes.RandomAcc(simState.Rand, simState.Accounts)
		receiver, _ := simtypes.RandomAcc(simState.Rand, simState.Accounts)
		if !sender.Equals(receiver) {
			relationships[index] = types.NewRelationship(
				sender.Address.String(),
				receiver.Address.String(),
				RandomSubspace(simState.Rand),
			)
		}
	}

	return relationships
}

// randomUsersBlocks
func randomUsersBlocks(simState *module.SimulationState) []types.UserBlock {
	usersBlocksNumber := simState.Rand.Intn(simtypes.RandIntBetween(simState.Rand, 1, 30))

	usersBlocks := make([]types.UserBlock, usersBlocksNumber)
	for index := 0; index < usersBlocksNumber; index++ {
		blocker, _ := simtypes.RandomAcc(simState.Rand, simState.Accounts)
		blocked, _ := simtypes.RandomAcc(simState.Rand, simState.Accounts)
		if !blocker.Equals(blocked) {
			usersBlocks[index] = types.NewUserBlock(
				blocker.Address.String(),
				blocked.Address.String(),
				"reason",
				RandomSubspace(simState.Rand),
			)
		}
	}

	return usersBlocks
}
