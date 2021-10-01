package simulation

// DONTCOVER

import (
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/desmos-labs/desmos/v2/x/staging/subspaces/types"
)

var (
	randomNames = []string{"facebook", "mooncake", "hiddenguru", "twitter", "linkedin", "snapchat"}
	ids         = []string{
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		"2bdf5932925584b9a86470bea60adce69041608a447f84a3317723aa5678ec88",
		"3d59f7548e1af2151b64135003ce63c0a484c26b9b8b166a7b1c1805ec34b00a",
		"ec8202b6f9fb16f9e26b66367afa4e037752f3c09a18cefab426165e06a424b1",
		"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
		"3f40462915a3e6026a4d790127b95ded4d870f6ab18d9af2fcbc454168255237",
	}
)

type SubspaceData struct {
	Subspace       types.Subspace
	CreatorAccount simtypes.Account
}

// RandomSubspace picks and returns a random subspace from an array and returns its
// position in the array.
func RandomSubspace(r *rand.Rand, subspaces []types.Subspace) (types.Subspace, int) {
	idx := r.Intn(len(subspaces))
	return subspaces[idx], idx
}

func RandomSubspaceData(r *rand.Rand, accs []simtypes.Account) SubspaceData {
	simAccount, _ := simtypes.RandomAcc(r, accs)
	owner := simAccount.Address.String()

	// Create a random subspace
	subspace := types.NewSubspace(
		RandomSubspaceID(r),
		RandomName(r),
		owner,
		owner,
		RandomSubspaceType(r),
		RandomDate(r),
	)

	return SubspaceData{
		Subspace:       subspace,
		CreatorAccount: simAccount,
	}
}

// RandomSubspaceID returns a random id from the above random ids array
func RandomSubspaceID(r *rand.Rand) string {
	index := r.Intn(len(ids))
	return ids[index]
}

// RandomName returns a random subspace name from the above random names
func RandomName(r *rand.Rand) string {
	idx := r.Intn(len(randomNames))
	return randomNames[idx]
}

// RandomDate returns a random post creation date
func RandomDate(r *rand.Rand) time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2010, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := r.Int63n(delta) + min
	return time.Unix(sec, 0).Truncate(time.Millisecond)
}

// RandomSubspaceType returns a random subspace type
func RandomSubspaceType(r *rand.Rand) types.SubspaceType {
	if r.Intn(101) <= 50 {
		return types.SubspaceTypeClosed
	}
	return types.SubspaceTypeOpen
}

// GetAccount gets the account having the given address from the accs list
func GetAccount(address sdk.Address, accs []simtypes.Account) *simtypes.Account {
	for _, acc := range accs {
		if acc.Address.Equals(address) {
			return &acc
		}
	}
	return nil
}
