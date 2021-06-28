package simulation

// DONTCOVER

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// RandomizedGenState generates a random GenesisState for profile
func RandomizedGenState(simsState *module.SimulationState) {
	profilesNumber := simsState.Rand.Intn(len(simsState.Accounts) - 10)
	profiles := NewRandomProfiles(simsState.Rand, simsState.Accounts, profilesNumber)

	// Update the auth state with the accounts
	var authState authtypes.GenesisState
	err := simsState.Cdc.UnmarshalJSON(simsState.GenState[authtypes.ModuleName], &authState)
	if err != nil {
		panic(err)
	}

	var accountsAny = make([]*codectypes.Any, len(profiles))
	for index, profile := range profiles {
		accountsAny[index], err = codectypes.NewAnyWithValue(profile)
		if err != nil {
			panic(err)
		}
	}
	authState.Accounts = append(authState.Accounts, accountsAny...)

	bz, err := simsState.Cdc.MarshalJSON(&authState)
	if err != nil {
		panic(err)
	}
	simsState.GenState[authtypes.ModuleName] = bz

	// Create and set profiles state
	profileGenesis := types.NewGenesisState(
		randomDTagTransferRequests(profiles, simsState, simsState.Rand.Intn(profilesNumber/2)),
		randomRelationships(profiles, simsState, simsState.Rand.Intn(profilesNumber/2)),
		randomUsersBlocks(profiles, simsState, simsState.Rand.Intn(profilesNumber/2)),
		types.NewParams(
			RandomNicknameParams(simsState.Rand),
			RandomDTagParams(simsState.Rand),
			RandomBioParams(simsState.Rand),
		),
		types.IBCPortID,
		nil,
		nil,
	)

	bz, err = simsState.Cdc.MarshalJSON(profileGenesis)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated profile parameters:\n%s\n", bz)

	simsState.GenState[types.ModuleName] = simsState.Cdc.MustMarshalJSON(profileGenesis)
}

// randomDTagTransferRequests returns randomly generated genesis dTag transfer requests
func randomDTagTransferRequests(
	profiles []*types.Profile, simState *module.SimulationState, number int,
) []types.DTagTransferRequest {

	dTagTransferRequests := make([]types.DTagTransferRequest, number)
	for i := 0; i < number; {
		profile1 := RandomProfile(simState.Rand, profiles)
		profile2 := RandomProfile(simState.Rand, profiles)

		// Skip same profiles
		if profile1.GetAddress().Equals(profile2.GetAddress()) {
			continue
		}

		request := types.NewDTagTransferRequest(
			RandomDTag(simState.Rand),
			profile1.GetAddress().String(),
			profile2.GetAddress().String(),
		)

		// Skip duplicated requests
		if !containsDTagTransferRequest(dTagTransferRequests, request) {
			dTagTransferRequests[i] = request
			i++
		}
	}

	return dTagTransferRequests
}

func containsDTagTransferRequest(slice []types.DTagTransferRequest, request types.DTagTransferRequest) bool {
	for _, req := range slice {
		if req.Sender == request.Sender && req.Receiver == request.Receiver {
			return true
		}
	}
	return false
}

// -------------------------------------------------------------------------------------------------------------------

// randomRelationships returns randomly generated genesis relationships and their associated users - IDs map
func randomRelationships(
	profiles []*types.Profile, simState *module.SimulationState, number int,
) []types.Relationship {
	relationships := make([]types.Relationship, number)
	for index := 0; index < number; {
		profile1 := RandomProfile(simState.Rand, profiles)
		profile2 := RandomProfile(simState.Rand, profiles)

		// Skip same profiles
		if profile1.GetAddress().Equals(profile2.GetAddress()) {
			continue
		}

		relationship := types.NewRelationship(
			profile1.GetAddress().String(),
			profile2.GetAddress().String(),
			RandomSubspace(simState.Rand),
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
	profiles []*types.Profile, simState *module.SimulationState, number int,
) []types.UserBlock {
	usersBlocks := make([]types.UserBlock, number)
	for index := 0; index < number; {
		profile1 := RandomProfile(simState.Rand, profiles)
		profile2 := RandomProfile(simState.Rand, profiles)

		// Skip same profiles
		if profile1.GetAddress().Equals(profile2.GetAddress()) {
			continue
		}

		block := types.NewUserBlock(
			profile1.GetAddress().String(),
			profile2.GetAddress().String(),
			"reason",
			RandomSubspace(simState.Rand),
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
