package simulation

// DONTCOVER

import (
	"fmt"

	profilestypes "github.com/desmos-labs/desmos/v2/x/profiles/types"

	subspacessim "github.com/desmos-labs/desmos/v2/x/subspaces/simulation"
	subspacestypes "github.com/desmos-labs/desmos/v2/x/subspaces/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/cosmos/cosmos-sdk/types/module"

	profilessim "github.com/desmos-labs/desmos/v2/x/profiles/simulation"
	"github.com/desmos-labs/desmos/v2/x/relationships/types"
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
	profiles := getProfilesAccounts(accounts)

	// Create and set the subspaces state
	var subspacesState subspacestypes.GenesisState
	err = simsState.Cdc.UnmarshalJSON(simsState.GenState[subspacestypes.ModuleName], &subspacesState)
	if err != nil {
		panic(err)
	}

	profileGenesis := types.NewGenesisState(
		randomRelationships(profiles, subspacesState.Subspaces, simsState, simsState.Rand.Intn(profilesNumber)),
		randomUsersBlocks(profiles, subspacesState.Subspaces, simsState, simsState.Rand.Intn(profilesNumber)),
	)

	bz, err := simsState.Cdc.MarshalJSON(profileGenesis)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated profile parameters:\n%s\n", bz)

	simsState.GenState[types.ModuleName] = simsState.Cdc.MustMarshalJSON(profileGenesis)
}

// getProfilesAccounts filters the given accounts slice and returns only the profiles
func getProfilesAccounts(accounts []authtypes.GenesisAccount) []*profilestypes.Profile {
	var profiles []*profilestypes.Profile
	for _, account := range accounts {
		if profile, ok := account.(*profilestypes.Profile); ok {
			profiles = append(profiles, profile)
		}
	}
	return profiles
}

// -------------------------------------------------------------------------------------------------------------------

// randomRelationships returns randomly generated genesis relationships and their associated users - IDs map
func randomRelationships(
	profiles []*profilestypes.Profile, subspaces []subspacestypes.Subspace, simState *module.SimulationState, number int,
) []types.Relationship {
	relationships := make([]types.Relationship, number)
	for index := 0; index < number; {
		profile1 := profilessim.RandomProfile(simState.Rand, profiles)
		profile2 := profilessim.RandomProfile(simState.Rand, profiles)

		// Skip same profiles
		if profile1.GetAddress().Equals(profile2.GetAddress()) {
			continue
		}

		subspace, _ := subspacessim.RandomSubspace(simState.Rand, subspaces)
		relationship := types.NewRelationship(
			profile1.GetAddress().String(),
			profile2.GetAddress().String(),
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
	profiles []*profilestypes.Profile, subspaces []subspacestypes.Subspace, simState *module.SimulationState, number int,
) []types.UserBlock {
	usersBlocks := make([]types.UserBlock, number)
	for index := 0; index < number; {
		profile1 := profilessim.RandomProfile(simState.Rand, profiles)
		profile2 := profilessim.RandomProfile(simState.Rand, profiles)

		// Skip same profiles
		if profile1.GetAddress().Equals(profile2.GetAddress()) {
			continue
		}

		subspace, _ := subspacessim.RandomSubspace(simState.Rand, subspaces)
		block := types.NewUserBlock(
			profile1.GetAddress().String(),
			profile2.GetAddress().String(),
			"reason",
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
