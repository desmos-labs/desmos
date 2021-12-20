package keeper

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
	"github.com/desmos-labs/desmos/v2/x/profiles/wasm"
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

func (k Keeper) IteratePermissionedContracts(ctx sdk.Context, fn func(index int64, contract types.PermissionedContract) bool) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PermissionedContractsPrefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var contract types.PermissionedContract
		k.cdc.MustUnmarshal(iterator.Value(), &contract)
		stop := fn(i, contract)
		if stop {
			break
		}
		i++
	}
}

func (k Keeper) composeAuctionMessage(user string) (json.RawMessage, error) {
	auctionStatus := wasm.NewUpdateDtagAuctionStatus(user)
	bz, err := json.Marshal(&auctionStatus)
	if err != nil {
		return nil, err
	}
	return bz, nil
}

func (k Keeper) UpdateDtagAuctionStatus(ctx sdk.Context, contractAddress, userAddress string) error {
	message, err := k.composeAuctionMessage(userAddress)
	if err != nil {
		return err
	}
	address, err := sdk.AccAddressFromBech32(contractAddress)
	if err != nil {
		return err
	}
	_, err = k.wasmKeeper.Sudo(ctx, address, message)
	if err != nil {
		return err
	}

	return nil
}
