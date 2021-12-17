package keeper

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
	tendermint "github.com/tendermint/tendermint/abci/types"
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

func (k Keeper) composeAuctionMessage() json.RawMessage {
	// compose the []byte message with the given attributes
	return json.RawMessage{}
}

func (k Keeper) UpdateDtagAuctionStatus(ctx sdk.Context, contractAddress string, eventAttributes []tendermint.EventAttribute) error {
	message := k.composeAuctionMessage()
	address, err := sdk.AccAddressFromBech32(contractAddress)
	if err != nil {
		return err
	}
	k.wasmKeeper.Sudo(ctx, address, message)

	return nil
}
