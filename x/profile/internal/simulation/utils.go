package simulation

import (
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
)

var (
	dtagLetters    = "abcdefghijtuvwxyzDUVWXYZ123490_"
	monikerLetters = "abcdefghijtuvwxyzDUVWXYZ123490_ "

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
func NewRandomProfile(r *rand.Rand, account sdk.AccAddress) types.Profile {
	return types.NewProfile(RandomDTag(r), account, time.Now()).
		WithBio(RandomBio(r)).
		WithMoniker(RandomMoniker(r)).
		WithPictures(
			RandomProfilePic(r),
			RandomProfileCover(r))
}

// RandomProfile picks and returns a random profile from an array
func RandomProfile(r *rand.Rand, accounts types.Profiles) types.Profile {
	idx := r.Intn(len(accounts))
	return accounts[idx]
}

// RandomMoniker return a random dtag
func RandomDTag(r *rand.Rand) string {
	// DTag must be at least 3 characters and at most 30
	dTagLen := r.Intn(27) + 3

	b := make([]byte, dTagLen)
	for i := range b {
		b[i] = dtagLetters[r.Intn(len(dtagLetters))]
	}
	return string(b)
}

// RandomMoniker return a random moniker
func RandomMoniker(r *rand.Rand) *string {
	// Moniker must be at least 2 and at most 50 characters
	monikerLen := r.Intn(48) + 2

	b := make([]byte, monikerLen)
	for i := range b {
		b[i] = monikerLetters[r.Intn(len(monikerLetters))]
	}
	value := string(b)
	return &value
}

// RandomBio return a random bio value from the list of randomBios given
func RandomBio(r *rand.Rand) *string {
	idx := r.Intn(len(randomBios))
	return &randomBios[idx]
}

// RandomProfilePic return a random profile pic value from the list of randomProfilePics given
func RandomProfilePic(r *rand.Rand) *string {
	idx := r.Intn(len(randomProfilePics))
	return &randomProfilePics[idx]
}

// RandomProfileCover return a random profile cover from the list of randomProfileCovers
func RandomProfileCover(r *rand.Rand) *string {
	idx := r.Intn(len(randomProfileCovers))
	return &randomProfileCovers[idx]
}

// GetProfile gets the profile having the given address from the accs list
func GetSimAccount(address sdk.Address, accs []sim.Account) *sim.Account {
	for _, acc := range accs {
		if acc.Address.Equals(address) {
			return &acc
		}
	}
	return nil
}
