package simulation

//DONTCOVER

import (
	"math/rand"
	"time"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

var (
	randomBios = []string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		"Vestibulum a nulla sed purus pellentesque euismod quis ut risus.",
		"Morbi nec magna interdum, rhoncus nisl ac, posuere sapien.",
		"Duis vitae nisi efficitur, lobortis neque at, bibendum ipsum.",
		"Donec semper nisi vel mollis cursus.",
		"Curabitur quis massa id libero posuere venenatis ac ac erat.",
		"Morbi dictum elit vitae libero lobortis luctus.",
		"Nam sit amet velit venenatis est scelerisque egestas vitae nec turpis.",
		"Duis commodo sapien id ligula volutpat tincidunt in et est.",
		"Cras et magna cursus, volutpat purus at, dictum diam.",
		"Phasellus in arcu euismod, accumsan urna quis, consectetur enim.",
		"Morbi tincidunt urna sit amet vulputate bibendum.",
		"Etiam vehicula eros vel libero sollicitudin elementum.",
		"Pellentesque at nunc ac orci consequat varius.",
		"Donec aliquam libero eu purus cursus, in congue magna tempor.",
		"Vivamus a dolor scelerisque, posuere justo quis, pharetra nibh.",
	}

	randomProfilePics = []string{
		"https://shorturl.at/adnX3",
		"https://shorturl.at/adnX4",
	}

	randomProfileCovers = []string{
		"https://shorturl.at/cgpyF",
		"https://shorturl.at/cgpyG",
	}
)

// NewRandomProfile return a random ProfileData from random data and the given account
// nolint:interfacer
func NewRandomProfile(r *rand.Rand, account sdk.AccAddress) types.Profile {
	return types.NewProfile(
		RandomDTag(r),
		RandomMoniker(r),
		RandomBio(r),
		types.NewPictures(RandomProfilePic(r), RandomProfileCover(r)),
		time.Now(),
		account.String(),
	)
}

// RandomProfile picks and returns a random profile from an array
func RandomProfile(r *rand.Rand, accounts []types.Profile) types.Profile {
	idx := r.Intn(len(accounts))
	return accounts[idx]
}

// RandomDTagTransferRequest picks and returns a random dtag transfer request from an array of requests
func RandomDTagTransferRequest(r *rand.Rand, requests []types.DTagTransferRequest) types.DTagTransferRequest {
	idx := r.Intn(len(requests))
	return requests[idx]
}

// RandomDTag return a random dtag
func RandomDTag(r *rand.Rand) string {
	// DTag must be at least 3 characters and at most 30
	return simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 3, 30))
}

// RandomMoniker return a random moniker
func RandomMoniker(r *rand.Rand) string {
	randomMoniker := simtypes.RandStringOfLength(r, 30)
	return randomMoniker
}

// RandomBio return a random bio value from the list of randomBios given
func RandomBio(r *rand.Rand) string {
	idx := r.Intn(len(randomBios))
	return randomBios[idx]
}

// RandomProfilePic return a random profile pic value from the list of randomProfilePics given
func RandomProfilePic(r *rand.Rand) string {
	idx := r.Intn(len(randomProfilePics))
	return randomProfilePics[idx]
}

// RandomProfileCover return a random profile cover from the list of randomProfileCovers
func RandomProfileCover(r *rand.Rand) string {
	idx := r.Intn(len(randomProfileCovers))
	return randomProfileCovers[idx]
}

// GetProfile gets the profile having the given address from the accs list
func GetSimAccount(address sdk.Address, accs []simtypes.Account) *simtypes.Account {
	for _, acc := range accs {
		if acc.Address.Equals(address) {
			return &acc
		}
	}
	return nil
}

// RandomMonikerParams return a random set of moniker params
func RandomMonikerParams(r *rand.Rand) types.MonikerParams {
	randomMin := sdk.NewInt(int64(simtypes.RandIntBetween(r, 2, 3)))
	randomMax := sdk.NewInt(int64(simtypes.RandIntBetween(r, 30, 1000)))
	return types.NewMonikerParams(randomMin, randomMax)
}

// RandomDTagParams return a random set of moniker params
func RandomDTagParams(r *rand.Rand) types.DTagParams {
	randomMin := sdk.NewInt(int64(simtypes.RandIntBetween(r, 3, 4)))
	randomMax := sdk.NewInt(int64(simtypes.RandIntBetween(r, 30, 50)))
	return types.NewDtagParams("^[A-Za-z0-9_]+$", randomMin, randomMax)
}

// RandomBioParams return a random biography param
func RandomBioParams(r *rand.Rand) sdk.Int {
	return sdk.NewInt(int64(simtypes.RandIntBetween(r, 500, 1000)))
}
