package keeper

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

// WithWasmKeeper decorates profiles keeper with the cosmwasm keeper
func (k Keeper) WithWasmKeeper(wasmKeeper wasmkeeper.Keeper) Keeper {
	k.wasmKeeper = wasmKeeper
	return k
}

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

func (k Keeper) GetPermissionedContract(ctx sdk.Context, admin, contractAddress string) types.PermissionedContract {
	store := ctx.KVStore(k.storeKey)
	key := types.PermissionedContractKey(admin, contractAddress)

	var permissionedContract types.PermissionedContract
	cBz := store.Get(key)

	k.cdc.MustUnmarshal(cBz, &permissionedContract)

	return permissionedContract
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

func (k Keeper) UpdateDtagAuctionStatus(ctx sdk.Context, contractAddress string, msg types.SudoMsg) error {
	auctionStatus := types.NewUpdateDTagAuctionStatusMsg(msg.UpdateDtagAuctionStatus.User, msg.UpdateDtagAuctionStatus.TransferStatus)
	message, err := auctionStatus.Marshal()
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
