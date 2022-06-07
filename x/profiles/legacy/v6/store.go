package v6

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/profiles/types"
)

// MigrateStore performs in-place store migrations from v6 to v7
// The migration includes:
//
// - set the default external address to oldest chain link of each chain for each owner
//
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	// Set the default external addresses
	err := setDefaultExternalAddresses(store, cdc)
	if err != nil {
		return err
	}

	return nil
}

// setDefaultExternalAddresses set the default external address of each chain for each owner
func setDefaultExternalAddresses(store sdk.KVStore, cdc codec.BinaryCodec) error {
	chainLinkStore := prefix.NewStore(store, types.ChainLinksPrefix)
	chainLinksIterator := chainLinkStore.Iterator(nil, nil)

	for ; chainLinksIterator.Valid(); chainLinksIterator.Next() {
		var link types.ChainLink
		err := cdc.Unmarshal(chainLinksIterator.Value(), &link)
		if err != nil {
			return err
		}

		// Validate the source address
		srcAddrData, err := types.UnpackAddressData(cdc, link.Address)
		if err != nil {
			return err
		}

		// Update default external address if the key exists
		if store.Has(types.DefaultExternalAddressKey(link.User, link.ChainConfig.Name)) {
			addrBz := store.Get(types.DefaultExternalAddressKey(link.User, link.ChainConfig.Name))
			var defaultLink types.ChainLink
			err := cdc.Unmarshal(store.Get(types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, string(addrBz))), &defaultLink)
			if err != nil {
				return err
			}

			// Skip if the new link is after the default one
			if link.CreationTime.After(defaultLink.CreationTime) {
				continue
			}
		}

		store.Set(types.DefaultExternalAddressKey(link.User, link.ChainConfig.Name), []byte(srcAddrData.GetValue()))
	}

	chainLinksIterator.Close()
	return nil
}
