package v7

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

// MigrateStore migrates the store from version 6 to version 7.
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	// Add user allowances to expiring queue
	addAllowancesToExpiringQueue(ctx, storeKey, cdc, types.UserAllowancePrefix)

	// Add group allowances to expiring queue
	addAllowancesToExpiringQueue(ctx, storeKey, cdc, types.GroupAllowancePrefix)
	return nil
}

func addAllowancesToExpiringQueue(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, allowancePrefix []byte) {
	store := ctx.KVStore(storeKey)
	allowanceStore := prefix.NewStore(store, allowancePrefix)
	iterator := allowanceStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var grant types.Grant
		cdc.MustUnmarshal(iterator.Value(), &grant)

		expiration := grant.GetExpiration()
		if expiration != nil {
			// Re-add the prefix because the prefix store trims it out, and we need it to get the data
			allowanceKey := append([]byte(nil), allowancePrefix...)
			allowanceKey = append(allowanceKey, iterator.Key()...)
			store.Set(types.ExpiringAllowanceKey(expiration, allowanceKey), []byte{0x1})
		}
	}
}
