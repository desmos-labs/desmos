package simulation

// DONTCOVER

import (
	"fmt"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// RandomizedGenState generates a random GenesisState for profile
func RandomizedGenState(simsState *module.SimulationState) {

	profileGenesis := types.NewGenesisState(
		randomDTagTransferRequests(simsState),
		randomRelationships(simsState),
		randomUsersBlocks(simsState),
		types.NewParams(
			RandomUsernameParams(simsState.Rand),
			RandomDTagParams(simsState.Rand),
			RandomBioParams(simsState.Rand),
		),
	)

	bz, err := simsState.Cdc.MarshalJSON(profileGenesis)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated profile parameters:\n%s\n", bz)

	simsState.GenState[types.ModuleName] = simsState.Cdc.MustMarshalJSON(profileGenesis)
}

// randomDTagTransferRequests returns randomly generated genesis dTag transfer requests
func randomDTagTransferRequests(simState *module.SimulationState) []types.DTagTransferRequest {
	dTagTransferRequestsNumber := simState.Rand.Intn(20)

	dTagTransferRequests := make([]types.DTagTransferRequest, dTagTransferRequestsNumber)
	for i := 0; i < dTagTransferRequestsNumber; i++ {
		simAccount, _ := simtypes.RandomAcc(simState.Rand, simState.Accounts)
		simAccount2, _ := simtypes.RandomAcc(simState.Rand, simState.Accounts)
		dTagTransferRequests[i] = types.NewDTagTransferRequest(
			RandomDTag(simState.Rand),
			simAccount.Address.String(),
			simAccount2.Address.String(),
		)
	}

	return dTagTransferRequests
}

// randomRelationships returns randomly generated genesis relationships and their associated users - IDs map
func randomRelationships(simState *module.SimulationState) []types.Relationship {
	relationshipsNumber := simState.Rand.Intn(simtypes.RandIntBetween(simState.Rand, 1, 30))

	relationships := make([]types.Relationship, relationshipsNumber)
	for index := 0; index < relationshipsNumber; {
		sender, _ := simtypes.RandomAcc(simState.Rand, simState.Accounts)
		receiver, _ := simtypes.RandomAcc(simState.Rand, simState.Accounts)
		if !sender.Equals(receiver) {
			newRelationship := types.NewRelationship(
				sender.Address.String(),
				receiver.Address.String(),
				RandomSubspace(simState.Rand),
			)
			if !containsRelationship(relationships, newRelationship) {
				relationships[index] = newRelationship
				index++
			}
		}
	}

	return relationships
}

// containsRelationship returns true iff the given slice contains the given relationship
func containsRelationship(slice []types.Relationship, relationship types.Relationship) bool {
	for _, rel := range slice {
		if rel.Equal(relationship) {
			return true
		}
	}
	return false
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
