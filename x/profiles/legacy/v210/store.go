package v210

import (
	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

// MigrateStore performs in-place store migrations from v2.0 to v2.1 The
// migration includes:
//
// - Replace all app links plain text with the HEX-encoded equivalent
// - Replace all the chain links plain text with the HEX-encoded equivalent
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	err := migrateAppLinks(store, cdc)
	if err != nil {
		return err
	}

	err = migrateChainLinks(store, cdc)
	if err != nil {
		return err
	}

	return nil
}

func migrateAppLinks(store sdk.KVStore, cdc codec.BinaryCodec) error {
	iterator := sdk.KVStorePrefixIterator(store, types.UserApplicationLinkPrefix)
	defer iterator.Close()

	var keys [][]byte
	var newLinks []types.ApplicationLink
	for ; iterator.Valid(); iterator.Next() {
		var link types.ApplicationLink
		err := cdc.Unmarshal(iterator.Value(), &link)
		if err != nil {
			return err
		}

		// Change the success result value to be HEX encoded
		if result := link.Result; result != nil {
			if successResult, ok := (result.Sum).(*types.Result_Success_); ok {
				successResult.Success.Value = hex.EncodeToString([]byte(successResult.Success.Value))
			}
		}

		keys = append(keys, iterator.Key())
		newLinks = append(newLinks, link)

		store.Delete(iterator.Key())
	}

	for index, link := range newLinks {
		store.Set(keys[index], types.MustMarshalApplicationLink(cdc, link))
	}

	return nil
}

func migrateChainLinks(store sdk.KVStore, cdc codec.BinaryCodec) error {
	iterator := sdk.KVStorePrefixIterator(store, types.ChainLinksPrefix)
	defer iterator.Close()

	var keys [][]byte
	var newLinks []types.ChainLink
	for ; iterator.Valid(); iterator.Next() {
		var link types.ChainLink
		err := cdc.Unmarshal(iterator.Value(), &link)
		if err != nil {
			return err
		}

		// Change the plain text to be HEX encoded
		link.Proof.PlainText = hex.EncodeToString([]byte(link.Proof.PlainText))

		keys = append(keys, iterator.Key())
		newLinks = append(newLinks, link)

		store.Delete(iterator.Key())
	}

	for index, link := range newLinks {
		store.Set(keys[index], types.MustMarshalChainLink(cdc, link))
	}

	return nil
}
