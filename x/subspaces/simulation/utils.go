package simulation

// DONTCOVER

import (
	"math/rand"
	"time"

	poststypes "github.com/desmos-labs/desmos/v5/x/posts/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

var (
	validPermissions = []types.Permissions{
		types.CombinePermissions(poststypes.PermissionWrite, poststypes.PermissionEditOwnContent, poststypes.PermissionInteractWithContent),
		types.CombinePermissions(poststypes.PermissionWrite, poststypes.PermissionInteractWithContent),
		types.CombinePermissions(poststypes.PermissionWrite, poststypes.PermissionEditOwnContent),
		types.CombinePermissions(poststypes.PermissionWrite, poststypes.PermissionEditOwnContent, poststypes.PermissionInteractWithContent, types.PermissionDeleteSubspace),
		types.CombinePermissions(types.PermissionEverything),
	}
)

// RandomSubspace picks a random subspace from an array and returns its position as well as value.
func RandomSubspace(r *rand.Rand, subspaces []types.Subspace) types.Subspace {
	return subspaces[r.Intn(len(subspaces))]
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
		GenerateRandomFeeTokens(r),
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

// RandomSection returns a randomly selected section from the slice given
func RandomSection(r *rand.Rand, sections []types.Section) types.Section {
	return sections[r.Intn(len(sections))]
}

// RandomSectionName returns a random section name
func RandomSectionName(r *rand.Rand) string {
	return simtypes.RandStringOfLength(r, 10)
}

// RandomSectionDescription returns a random section description
func RandomSectionDescription(r *rand.Rand) string {
	return simtypes.RandStringOfLength(r, 20)
}

// RandomGroup returns a random group selecting it from the list of groups given
func RandomGroup(r *rand.Rand, groups []types.UserGroup) types.UserGroup {
	return groups[r.Intn(len(groups))]
}

// RandomPermission returns a random permission from the given slice
func RandomPermission(r *rand.Rand, permissions []types.Permissions) types.Permissions {
	return permissions[r.Intn(len(permissions))]
}

// RandomAddress returns a random address from the slice given
func RandomAddress(r *rand.Rand, addresses []string) string {
	return addresses[r.Intn(len(addresses))]
}

// RandomAuthAccount returns a random account from the slice given
func RandomAuthAccount(r *rand.Rand, accounts []authtypes.AccountI) authtypes.AccountI {
	return accounts[r.Intn(len(accounts))]
}

// GetAccount gets the account having the given address from the accs list
func GetAccount(address string, accs []simtypes.Account) *simtypes.Account {
	for _, acc := range accs {
		if acc.Address.String() == address {
			return &acc
		}
	}
	return nil
}

// RandomGrant returns a random user grant from the slice given
func RandomGrant(r *rand.Rand, grants []types.Grant) types.Grant {
	return grants[r.Intn(len(grants))]
}

// GenerateRandomFeeTokens generates a list of fee tokens
func GenerateRandomFeeTokens(r *rand.Rand) sdk.Coins {
	coins := make(sdk.Coins, r.Intn(10))

	for i := range coins {
		coins[i] = sdk.NewCoin(simtypes.RandStringOfLength(r, 10), sdk.NewInt(r.Int63n(1000000)))
	}

	return coins.Sort()
}
