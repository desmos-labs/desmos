package simulation

// DONTCOVER

import (
	"math/rand"
	"time"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v2/x/profiles/types"
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

	subspaces = []string{
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		"2bdf5932925584b9a86470bea60adce69041608a447f84a3317723aa5678ec88",
		"3d59f7548e1af2151b64135003ce63c0a484c26b9b8b166a7b1c1805ec34b00a",
		"ec8202b6f9fb16f9e26b66367afa4e037752f3c09a18cefab426165e06a424b1",
		"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
		"3f40462915a3e6026a4d790127b95ded4d870f6ab18d9af2fcbc454168255237",
	}

	feeCoins = sdk.NewCoins(
		sdk.NewCoin("band", sdk.NewInt(10)),
		sdk.NewCoin("atom", sdk.NewInt(10)),
		sdk.NewCoin("desmos", sdk.NewInt(10)),
		sdk.NewCoin("akt", sdk.NewInt(10)),
		sdk.NewCoin("dvpn", sdk.NewInt(10)),
		sdk.NewCoin("daric", sdk.NewInt(10)),
		sdk.NewCoin("osmo", sdk.NewInt(10)),
		sdk.NewCoin("regen", sdk.NewInt(10)),
	)
)

// NewRandomProfiles returns number random profiles
func NewRandomProfiles(r *rand.Rand, accounts []simtypes.Account, number int) []*types.Profile {
	var profiles = make([]*types.Profile, number)
	for index := range profiles {
		account := accounts[index]
		profiles[index] = NewRandomProfile(r, authtypes.NewBaseAccount(
			account.Address,
			account.PubKey,
			0,
			0,
		))
	}
	return profiles
}

// NewRandomProfile return a random ProfileData from random data and the given account
// nolint:interfacer
func NewRandomProfile(r *rand.Rand, account authtypes.AccountI) *types.Profile {
	profile, err := types.NewProfile(
		RandomDTag(r),
		RandomNickname(r),
		RandomBio(r),
		types.NewPictures(RandomProfilePic(r), RandomProfileCover(r)),
		time.Now(),
		account,
	)
	if err != nil {
		panic(err)
	}
	return profile
}

// RandomProfile picks and returns a random profile from an array
func RandomProfile(r *rand.Rand, accounts []*types.Profile) *types.Profile {
	idx := r.Intn(len(accounts))
	return accounts[idx]
}

// RandomDTagTransferRequest picks and returns a random DTag transfer request from an array of requests
func RandomDTagTransferRequest(r *rand.Rand, requests []types.DTagTransferRequest) types.DTagTransferRequest {
	idx := r.Intn(len(requests))
	return requests[idx]
}

// RandomDTag return a random DTag
func RandomDTag(r *rand.Rand) string {
	// DTag must be at least 3 characters and at most 30
	return simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 6, 30))
}

// RandomNickname return a random nickname
func RandomNickname(r *rand.Rand) string {
	return simtypes.RandStringOfLength(r, 30)
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

// GetSimAccount gets the profile having the given address from the accs list
func GetSimAccount(address sdk.Address, accs []simtypes.Account) *simtypes.Account {
	for _, acc := range accs {
		if acc.Address.Equals(address) {
			return &acc
		}
	}
	return nil
}

// RandomNicknameParams return a random set of nickname params
func RandomNicknameParams(r *rand.Rand) types.NicknameParams {
	randomMin := sdk.NewInt(int64(simtypes.RandIntBetween(r, 2, 3)))
	randomMax := sdk.NewInt(int64(simtypes.RandIntBetween(r, 30, 1000)))
	return types.NewNicknameParams(randomMin, randomMax)
}

// RandomDTagParams return a random set of nickname params
func RandomDTagParams(r *rand.Rand) types.DTagParams {
	randomMin := sdk.NewInt(int64(simtypes.RandIntBetween(r, 3, 4)))
	randomMax := sdk.NewInt(int64(simtypes.RandIntBetween(r, 30, 50)))
	return types.NewDTagParams("^[A-Za-z0-9_]+$", randomMin, randomMax)
}

// RandomBioParams return a random biography param
func RandomBioParams(r *rand.Rand) types.BioParams {
	randomMax := sdk.NewInt(int64(simtypes.RandIntBetween(r, 500, 1000)))
	return types.NewBioParams(randomMax)
}

// RandomOracleParams return a random oracle param
func RandomOracleParams(r *rand.Rand) types.OracleParams {
	randomMinCount := uint64(simtypes.RandIntBetween(r, 1, 20))
	return types.NewOracleParams(
		r.Uint64(),
		uint64(simtypes.RandIntBetween(r, int(randomMinCount), int(randomMinCount)+50)),
		randomMinCount,
		uint64(simtypes.RandIntBetween(r, 1, 10000)),
		uint64(simtypes.RandIntBetween(r, 1, 10000)),
		simtypes.RandSubsetCoins(r, feeCoins)...,
	)
}

// RandomRelationship picks and returns a random relationships from an array
func RandomRelationship(r *rand.Rand, relationships []types.Relationship) types.Relationship {
	idx := r.Intn(len(relationships))
	return relationships[idx]
}

// RandomSubspace returns a random post subspace from the above random subspaces
func RandomSubspace(r *rand.Rand) string {
	idx := r.Intn(len(subspaces))
	return subspaces[idx]
}

// RandomUserBlock picks and returns a random user block from an array
func RandomUserBlock(r *rand.Rand, userBlocks []types.UserBlock) types.UserBlock {
	idx := r.Intn(len(userBlocks))
	return userBlocks[idx]
}

// RandomDTagTransferRequests returns a new random DTag transfer request from the ones given
func RandomDTagTransferRequests(r *rand.Rand, requests []types.DTagTransferRequest) types.DTagTransferRequest {
	return requests[r.Intn(len(requests))-1]
}
