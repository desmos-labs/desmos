package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	subspacessim "github.com/desmos-labs/desmos/v7/x/subspaces/simulation"
	subspacestypes "github.com/desmos-labs/desmos/v7/x/subspaces/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/desmos-labs/desmos/v7/x/relationships/types"
)

// RandomizedGenState generates a random GenesisState for profile
func RandomizedGenState(simsState *module.SimulationState) {
	profilesNumber := len(simsState.Accounts)

	// Get the profiles from the auth genesis state (which has already been initialized)
	var authState authtypes.GenesisState
	err := simsState.Cdc.UnmarshalJSON(simsState.GenState[authtypes.ModuleName], &authState)
	if err != nil {
		panic(err)
	}

	accounts, err := authtypes.UnpackAccounts(authState.Accounts)
	if err != nil {
		panic(err)
	}

	// Create and set the subspaces state
	var subspacesState subspacestypes.GenesisState
	err = simsState.Cdc.UnmarshalJSON(simsState.GenState[subspacestypes.ModuleName], &subspacesState)
	if err != nil {
		panic(err)
	}

	profileGenesis := types.NewGenesisState(
		randomRelationships(simsState.Rand, accounts, subspacesState.Subspaces, simsState.Rand.Intn(profilesNumber)),
		randomUsersBlocks(simsState.Rand, accounts, subspacesState.Subspaces, simsState.Rand.Intn(profilesNumber)),
	)

	bz, err := simsState.Cdc.MarshalJSON(profileGenesis)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated profile parameters:\n%s\n", bz)

	simsState.GenState[types.ModuleName] = simsState.Cdc.MustMarshalJSON(profileGenesis)
}

// -------------------------------------------------------------------------------------------------------------------

// randomRelationships returns randomly generated genesis relationships and their associated users - IDs map
func randomRelationships(
	r *rand.Rand, accounts []authtypes.GenesisAccount, subspaces []subspacestypes.Subspace, number int,
) []types.Relationship {
	if len(subspaces) == 0 {
		return nil
	}

	relationships := make([]types.Relationship, number)
	for index := 0; index < number; {
		user := RandomGenesisAccount(r, accounts)
		counterparty := RandomGenesisAccount(r, accounts)

		// Skip same profiles
		if user.GetAddress().Equals(counterparty.GetAddress()) {
			continue
		}

		subspace := subspacessim.RandomSubspace(r, subspaces)
		relationship := types.NewRelationship(
			user.GetAddress().String(),
			counterparty.GetAddress().String(),
			subspace.ID,
		)

		if !containsRelationship(relationships, relationship) {
			relationships[index] = relationship
			index++
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

// -------------------------------------------------------------------------------------------------------------------

// randomUsersBlocks
func randomUsersBlocks(
	r *rand.Rand, accounts []authtypes.GenesisAccount, subspaces []subspacestypes.Subspace, number int,
) []types.UserBlock {
	if len(subspaces) == 0 {
		return nil
	}

	usersBlocks := make([]types.UserBlock, number)
	for index := 0; index < number; {
		blocker := RandomGenesisAccount(r, accounts)
		blocked := RandomGenesisAccount(r, accounts)

		// Skip same profiles
		if blocker.GetAddress().Equals(blocked.GetAddress()) {
			continue
		}

		subspace := subspacessim.RandomSubspace(r, subspaces)
		block := types.NewUserBlock(
			blocker.GetAddress().String(),
			blocked.GetAddress().String(),
			"",
			subspace.ID,
		)

		if !containsUserBlock(usersBlocks, block) {
			usersBlocks[index] = block
			index++
		}
	}

	return usersBlocks
}

// containsUserBlock returns true iff the given slice contains the given block
func containsUserBlock(slice []types.UserBlock, block types.UserBlock) bool {
	for _, b := range slice {
		if b.Equal(block) {
			return true
		}
	}
	return false
}
