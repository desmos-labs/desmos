package v5

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/profiles/types"
)

// MigrateStore performs in-place store migrations from v5 to v6
// The migration includes:
//
// - add missing application links owner keys
// - add missing chain links owner keys
//
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	// Fix the application links
	err := fixApplicationLinks(store, cdc)
	if err != nil {
		return err
	}

	// Fix the chain links
	err = fixChainLinks(store, cdc)
	if err != nil {
		return err
	}

	return nil
}

// fixApplicationLinks fixes the application links by adding the missing owner keys
func fixApplicationLinks(store sdk.KVStore, cdc codec.BinaryCodec) error {
	applicationLinksStore := prefix.NewStore(store, types.ApplicationLinkPrefix)
	applicationLinksIterator := applicationLinksStore.Iterator(nil, nil)

	var applicationLinks []types.ApplicationLink
	for ; applicationLinksIterator.Valid(); applicationLinksIterator.Next() {
		var applicationLink types.ApplicationLink
		err := cdc.Unmarshal(applicationLinksIterator.Value(), &applicationLink)
		if err != nil {
			return err
		}

		applicationLinks = append(applicationLinks, applicationLink)
	}

	applicationLinksIterator.Close()

	for _, link := range applicationLinks {
		store.Set(types.ApplicationLinkOwnerKey(link.Data.Application, link.Data.Username, link.User), []byte{0x01})
	}

	return nil
}

// fixChainLinks fixes the chain links by adding the missing owner keys
func fixChainLinks(store sdk.KVStore, cdc codec.BinaryCodec) error {
	chainLinkStore := prefix.NewStore(store, types.ChainLinksPrefix)
	chainLinksIterator := chainLinkStore.Iterator(nil, nil)

	var chainLinks []types.ChainLink
	for ; chainLinksIterator.Valid(); chainLinksIterator.Next() {
		var chainLink types.ChainLink
		err := cdc.Unmarshal(chainLinksIterator.Value(), &chainLink)
		if err != nil {
			return err
		}

		chainLinks = append(chainLinks, chainLink)
	}

	chainLinksIterator.Close()

	for _, link := range chainLinks {
		store.Set(types.ChainLinkOwnerKey(link.ChainConfig.Name, link.GetAddressData().GetValue(), link.User), []byte{0x01})
	}

	return nil
}
