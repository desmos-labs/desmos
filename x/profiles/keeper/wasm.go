package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

func (k Keeper) DoesPermissionedContractExist(ctx sdk.Context, admin, contractAddress string) bool {
	store := ctx.KVStore(k.storeKey)
	key := types.PermissionedContractKey(admin, contractAddress)

	return store.Has(key)
}

func (k Keeper) SavePermissionedContract(ctx sdk.Context, contract types.PermissionedContract) {
	store := ctx.KVStore(k.storeKey)
	key := types.PermissionedContractKey(contract.Admin, contract.Address)

	store.Set(key, k.cdc.MustMarshal(&contract))
}
