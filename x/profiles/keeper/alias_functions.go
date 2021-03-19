package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// IterateProfiles iterates through the profiles set and performs the provided function
func (k Keeper) IterateProfiles(ctx sdk.Context, fn func(index int64, profile *types.Profile) (stop bool)) {
	i := int64(0)
	k.ak.IterateAccounts(ctx, func(account authtypes.AccountI) (stop bool) {
		profile, ok := account.(*types.Profile)

		stop = false
		if ok {
			stop = fn(i, profile)
			i++
		}

		return stop
	})
}

func (k Keeper) GetProfiles(ctx sdk.Context) []*types.Profile {
	var profiles []*types.Profile
	k.IterateProfiles(ctx, func(_ int64, profile *types.Profile) (stop bool) {
		profiles = append(profiles, profile)
		return false
	})
	return profiles
}
