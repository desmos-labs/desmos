package simulation

// DONTCOVER

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/relationships/types"
)

// RandomizedGenState generates a random GenesisState for profile
func RandomizedGenState(simsState *module.SimulationState) {
	userRelationshipsMap := randomRelationships(simsState)

	profileGenesis := types.NewGenesisState(
		userRelationshipsMap,
	)

	simsState.GenState[types.ModuleName] = simsState.Cdc.MustMarshalJSON(profileGenesis)
}

// randomRelationships returns randomly generated genesis relationships and their associated users - IDs map
func randomRelationships(simState *module.SimulationState) map[string][]sdk.AccAddress {
	relationshipsNumber := simState.Rand.Intn(sim.RandIntBetween(simState.Rand, 1, 100))
	usersRelationships := map[string][]sdk.AccAddress{}

	for index := 0; index < relationshipsNumber; index++ {
		sender, _ := sim.RandomAcc(simState.Rand, simState.Accounts)
		receiver, _ := sim.RandomAcc(simState.Rand, simState.Accounts)
		if !sender.Equals(receiver) {
			usersRelationships[sender.Address.String()] = []sdk.AccAddress{receiver.Address}
		}
	}

	return usersRelationships
}
