package simulation

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
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
func randomRelationships(simState *module.SimulationState) map[string]types.Relationships {
	relationshipsNumber := simState.Rand.Intn(sim.RandIntBetween(simState.Rand, 1, 100))
	usersRelationships := map[string]types.Relationships{}
	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"

	for index := 0; index < relationshipsNumber; index++ {
		sender, _ := sim.RandomAcc(simState.Rand, simState.Accounts)
		receiver, _ := sim.RandomAcc(simState.Rand, simState.Accounts)
		if !sender.Equals(receiver) {
			usersRelationships[sender.Address.String()] = types.Relationships{types.NewRelationship(receiver.Address, subspace)}
		}
	}

	return usersRelationships
}

// randomUsersBlocks
func randomUsersBlocks(simState *module.SimulationState) []types.UserBlock {
	usersBlocksNumber := simState.Rand.Intn(sim.RandIntBetween(simState.Rand, 1, 100))
	var usersBlocks = make([]types.UserBlock, usersBlocksNumber)

	for index := 0; index < usersBlocksNumber; index++ {
		blocker, _ := sim.RandomAcc(simState.Rand, simState.Accounts)
		blocked, _ := sim.RandomAcc(simState.Rand, simState.Accounts)
		if !blocker.Equals(blocked) {
			usersBlocks[index] = types.NewUserBlock(blocker.Address, blocked.Address,
				"reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")
		}
	}

	return usersBlocks
}
