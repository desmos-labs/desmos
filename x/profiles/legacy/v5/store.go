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
// - remove all chain links that are not valid anymore due to the new rules
// - add missing chain links owner keys
//
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec, legacyAmino *codec.LegacyAmino) error {
	store := ctx.KVStore(storeKey)

	// Fix the application links
	err := fixApplicationLinks(store, cdc)
	if err != nil {
		return err
	}

	// Fix the chain links
	err = fixChainLinks(store, cdc, legacyAmino)
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
func fixChainLinks(store sdk.KVStore, cdc codec.BinaryCodec, legacyAmino *codec.LegacyAmino) error {
	chainLinkStore := prefix.NewStore(store, types.ChainLinksPrefix)
	chainLinksIterator := chainLinkStore.Iterator(nil, nil)

	var validChainLinks []types.ChainLink
	var invalidChainLinks []types.ChainLink
	for ; chainLinksIterator.Valid(); chainLinksIterator.Next() {
		var chainLink types.ChainLink
		err := cdc.Unmarshal(chainLinksIterator.Value(), &chainLink)
		if err != nil {
			return err
		}

		// Make sure the signed value is valid, if it's a transaction
		err = chainLink.Proof.Verify(cdc, legacyAmino, chainLink.User, chainLink.GetAddressData())
		if err == nil {
			validChainLinks = append(validChainLinks, chainLink)
		} else {
			invalidChainLinks = append(invalidChainLinks, chainLink)
		}
	}

	chainLinksIterator.Close()

	// Delete invalid chain links
	for _, link := range invalidChainLinks {
		store.Delete(types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, link.GetAddressData().GetValue()))
	}

	// Store the owners of valid chain links
	for _, link := range validChainLinks {
		store.Set(types.ChainLinkOwnerKey(link.ChainConfig.Name, link.GetAddressData().GetValue(), link.User), []byte{0x01})
	}

	return nil
}
