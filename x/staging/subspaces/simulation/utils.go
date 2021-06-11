package simulation

// DONTCOVER

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
)

var (
	randomNames = []string{"facebook", "mooncake", "hiddenguru", "twitter", "linkedin", "snapchat"}
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

// RandomSubspaceID returns a randomly generated subspaceID
func RandomSubspaceID(r *rand.Rand) string {
	randBytes := make([]byte, 4)
	_, err := r.Read(randBytes)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(randBytes)
	return hex.EncodeToString(hash[:])
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
