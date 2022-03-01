package simulation

// DONTCOVER

import (
	"math/rand"
	"time"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

// RandomGenesisSubspace picks a random genesis subspace from the given slice
func RandomGenesisSubspace(r *rand.Rand, subspaces []types.GenesisSubspace) types.GenesisSubspace {
	return subspaces[r.Intn(len(subspaces))]
}

// RandomSubspace picks a random subspace from an array and returns its position as well as value.
func RandomSubspace(r *rand.Rand, subspaces []types.Subspace) (types.Subspace, int) {
	idx := r.Intn(len(subspaces))
	return subspaces[idx], idx
}

// GenerateRandomSubspace generates a new subspace containing random data
func GenerateRandomSubspace(r *rand.Rand, accs []simtypes.Account) types.Subspace {
	simAccount, _ := simtypes.RandomAcc(r, accs)
	creator := simAccount.Address.String()

	return types.NewSubspace(
		RandomID(r),
		RandomName(r),
		RandomDescription(r),
		creator,
		creator,
		creator,
		RandomDate(r),
	)
}

// RandomID returns a new random ID
func RandomID(r *rand.Rand) uint64 {
	return r.Uint64()
}

// RandomName returns a random subspace name
func RandomName(r *rand.Rand) string {
	return simtypes.RandStringOfLength(r, 10)
}

// RandomDescription returns a random subspace description
func RandomDescription(r *rand.Rand) string {
	return simtypes.RandStringOfLength(r, 30)
}

// RandomDate returns a random post creation date
func RandomDate(r *rand.Rand) time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2010, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := r.Int63n(delta) + min
	return time.Unix(sec, 0).Truncate(time.Millisecond)
}

// RandomString returns a random string from the given slice
func RandomString(r *rand.Rand, strings []string) string {
	return strings[r.Intn(len(strings))]
}

// RandomGroup returns a random group selecting it from the list of groups given
func RandomGroup(r *rand.Rand, groups []types.UserGroup) types.UserGroup {
	return groups[r.Intn(len(groups))]
}

// RandomPermission returns a random permission from the given slice
func RandomPermission(r *rand.Rand, permissions []types.Permission) types.Permission {
	return permissions[r.Intn(len(permissions))]
}

// RandomAddress returns a random address from the slice given
func RandomAddress(r *rand.Rand, addresses []sdk.AccAddress) sdk.AccAddress {
	return addresses[r.Intn(len(addresses))]
}

// RandomAuthAccount returns a random account from the slice given
func RandomAuthAccount(r *rand.Rand, accounts []authtypes.AccountI) authtypes.AccountI {
	return accounts[r.Intn(len(accounts))]
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
