package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
)

// IterateProfile iterates through the profiles set and performs the provided function
func (k Keeper) IterateProfile(ctx sdk.Context, fn func(index int64, profile types.Profile) (stop bool)) {
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ProfileStorePrefix)
	defer iterator.Close()
	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var profile types.Profile
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &profile)
		stop := fn(i, profile)
		if stop {
			break
		}
		i++
	}
}
