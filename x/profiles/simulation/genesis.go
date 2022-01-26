package simulation

// DONTCOVER

import (
	"fmt"

	subspacestypes "github.com/desmos-labs/desmos/v2/x/subspaces/types"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

// RandomizedGenState generates a random GenesisState for profile
func RandomizedGenState(simsState *module.SimulationState) {
	profilesNumber := len(simsState.Accounts)
	profiles := NewRandomProfiles(simsState.Rand, simsState.Accounts, profilesNumber)

	// Update the auth state with the profiles
	var authState authtypes.GenesisState
	err := simsState.Cdc.UnmarshalJSON(simsState.GenState[authtypes.ModuleName], &authState)
	if err != nil {
		panic(err)
	}

	genAccounts, err := mergeAccountsWithProfiles(authState.Accounts, profiles)
	if err != nil {
		panic(err)
	}
	authState.Accounts = genAccounts

	bz, err := simsState.Cdc.MarshalJSON(&authState)
	if err != nil {
		panic(err)
	}
	simsState.GenState[authtypes.ModuleName] = bz

	// Create and set profiles state
	var subspacesState subspacestypes.GenesisState
	err = simsState.Cdc.UnmarshalJSON(simsState.GenState[subspacestypes.ModuleName], &subspacesState)
	if err != nil {
		panic(err)
	}

	profileGenesis := types.NewGenesisState(
		randomDTagTransferRequests(profiles, simsState, simsState.Rand.Intn(profilesNumber)),
		types.NewParams(
			RandomNicknameParams(simsState.Rand),
			RandomDTagParams(simsState.Rand),
			RandomBioParams(simsState.Rand),
			RandomOracleParams(simsState.Rand),
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

// mergeAccountsWithProfiles merges the provided x/auth genesis accounts with the given profiles, replacing
// any existing account with the associated profile (if existing).
func mergeAccountsWithProfiles(genAccounts []*codectypes.Any, profiles []*types.Profile) ([]*codectypes.Any, error) {
	// Unpack the accounts
	accounts, err := authtypes.UnpackAccounts(genAccounts)
	if err != nil {
		return nil, err
	}

	for index, account := range accounts {
		// See if the account also has a profile
		var profile *types.Profile
		for _, p := range profiles {
			if p.GetAddress().Equals(account.GetAddress()) {
				profile = p
				break
			}
		}

		// If found replace the account with the profile
		if profile != nil {
			accounts[index] = profile
		}
	}

	return authtypes.PackAccounts(accounts)
}

// -------------------------------------------------------------------------------------------------------------------

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
			profile2.DTag,
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
