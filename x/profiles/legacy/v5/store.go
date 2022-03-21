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
		store.Set(types.ApplicationLinkOwnerKey(link.Data.Application, link.Data.Username, link.User), []byte(link.User))
	}

	// Fix the chain links - TODO
	//chainLinkStore := prefix.NewStore(store, types.ChainLinksPrefix)
	//chainLinksIterator := chainLinkStore.Iterator(nil, nil)
	//
	//var chainLinks []types.ChainLink
	//for ; chainLinksIterator.Valid(); chainLinksIterator.Next() {
	//	var chainLink types.ChainLink
	//	err := cdc.Unmarshal(chainLinksIterator.Value(), &chainLink)
	//	if err != nil {
	//		return err
	//	}
	//
	//	chainLinks = append(chainLinks, chainLink)
	//}
	//
	//for _, link := range chainLinks {
	//	store.Set(types.ChainLinkOwnerKey(link.ChainConfig.Name, link.GetAddressData().GetValue(), link.User), []byte(link.User))
	//}

	return nil
}
