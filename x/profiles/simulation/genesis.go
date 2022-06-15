package simulation

// DONTCOVER

import (
	"fmt"
	"time"

	"github.com/desmos-labs/desmos/v3/testutil/profilestesting"

	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/desmos-labs/desmos/v3/x/profiles/types"
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

	chainLinks := randomChainLinks(profiles, simsState)

	profileGenesis := types.NewGenesisState(
		randomDTagTransferRequests(profiles, simsState, simsState.Rand.Intn(profilesNumber)),
		types.NewParams(
			RandomNicknameParams(simsState.Rand),
			RandomDTagParams(simsState.Rand),
			RandomBioParams(simsState.Rand),
			RandomOracleParams(simsState.Rand),
		),
		types.IBCPortID,
		chainLinks,
		getDefaultExternalAddressEntries(chainLinks, simsState),
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

// -------------------------------------------------------------------------------------------------------------------

// randomChainLinks returns randomly generated genesis chain links
func randomChainLinks(
	profiles []*types.Profile, simsState *module.SimulationState,
) []types.ChainLink {
	linksNumber := simsState.Rand.Intn(100)
	links := make([]types.ChainLink, linksNumber)
	for i := 0; i < linksNumber; {
		// Get a random profile
		profile := RandomProfile(simsState.Rand, profiles)
		chainAccount := profilestesting.GetChainLinkAccount("cosmos", "cosmos")
		link := chainAccount.GetBech32ChainLink(profile.GetAddress().String(), time.Date(2022, 6, 9, 0, 0, 0, 0, time.UTC))

		if !containsChainLink(links, link) {
			links[i] = link
			i++
		}
	}
	return links
}

func containsChainLink(slice []types.ChainLink, link types.ChainLink) bool {
	for _, l := range slice {
		if l.User == link.User && l.Address.Equal(link.Address) && l.ChainConfig.Name == link.ChainConfig.Name {
			return true
		}
	}
	return false
}

// -------------------------------------------------------------------------------------------------------------------

// getDefaultExternalAddressEntries returns randomly generated genesis default external address entries
func getDefaultExternalAddressEntries(
	links []types.ChainLink, simsState *module.SimulationState,
) []types.DefaultExternalAddressEntry {
	entries := make([]types.DefaultExternalAddressEntry, 0, len(links))
	for _, link := range links {
		entry := types.NewDefaultExternalAddressEntry(link.User, link.ChainConfig.Name, link.GetAddressData().GetValue())
		if !containsDefaultExternalAddressEntry(entries, entry) {
			entries = append(entries, entry)
		}
	}
	return entries
}

func containsDefaultExternalAddressEntry(slice []types.DefaultExternalAddressEntry, entry types.DefaultExternalAddressEntry) bool {
	for _, e := range slice {
		if e.Owner == entry.Owner && e.ChainName == entry.ChainName {
			return true
		}
	}
	return false
}
