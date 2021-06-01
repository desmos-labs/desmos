package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// StoreChainLink stores the given chain link inside the current context.
// It assumes that the given chain link has already been validated.
func (k Keeper) StoreChainLink(ctx sdk.Context, link types.ChainLink, destinationAddress string) error {
	if _, found := k.GetChainLink(ctx, link.ChainConfig.Name, link.Address); found {
		return fmt.Errorf("chain link already exists")
	}

	// Check if address has the profile
	profile, found, err := k.GetProfile(ctx, destinationAddress)
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("address does not have any profile")
	}
	// Store chain link to the profile
	profile.ChainsLinks = append(profile.ChainsLinks, link)
	if err := k.StoreProfile(ctx, profile); err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	key := types.ChainsLinksStoreKey(link.ChainConfig.Name, link.Address)
	store.Set(key, k.cdc.MustMarshalBinaryBare(&link))
	return nil
}

// GetChainLink returns the chain link corresponding to the given address and the given chain name inside the current context.
func (k Keeper) GetChainLink(ctx sdk.Context, chainName string, address string) (link types.ChainLink, found bool) {
	store := ctx.KVStore((k.storeKey))

	bz := store.Get(types.ChainsLinksStoreKey(chainName, address))
	if bz == nil {
		return types.ChainLink{}, false
	}
	k.cdc.MustUnmarshalBinaryBare(bz, &link)
	return link, true
}

// GetAllChainsLinks returns a list of all the chains links inside the given context.
func (k Keeper) GetAllChainsLinks(ctx sdk.Context) []types.ChainLink {
	var links []types.ChainLink
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ChainsLinksPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var link types.ChainLink
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &link)
		links = append(links, link)
	}
	return links
}

// GetChainLinksWithPagination returns a list of the links which is paginated from all the links
func (k Keeper) GetChainsLinksWithPagination(ctx sdk.Context, page int, limit int) []types.ChainLink {
	links := k.GetAllChainsLinks(ctx)
	if page == 0 {
		page = 1
	}
	var pagedLinks []types.ChainLink
	start, end := client.Paginate(len(links), page, limit, sdkquery.DefaultLimit)
	if start < 0 || end < 0 {
		pagedLinks = []types.ChainLink{}
	} else {
		pagedLinks = links[start:end]
	}
	return pagedLinks
}

// GetUserChainsLinks returns a list of links by a given address and pagination params
func (k Keeper) GetUserChainsLinks(ctx sdk.Context, address string, page int, limit int) ([]types.ChainLink, error) {
	profile, found, err := k.GetProfile(ctx, address)
	if err != nil {
		return []types.ChainLink{}, err
	}

	if !found {
		return []types.ChainLink{}, nil
	}

	links := profile.ChainsLinks
	if page == 0 {
		page = 1
	}
	var pagedLinks []types.ChainLink
	start, end := client.Paginate(len(links), page, limit, sdkquery.DefaultLimit)
	if start < 0 || end < 0 {
		pagedLinks = []types.ChainLink{}
	} else {
		pagedLinks = links[start:end]
	}
	return pagedLinks, nil
}

// DeleteLink allows to delete a link associated with the given address and chain name inside the current context.
// It assumes that the related link exists.
func (k Keeper) DeleteChainLink(ctx sdk.Context, chainName string, address string) {
	store := ctx.KVStore(k.storeKey)
	key := types.ChainsLinksStoreKey(chainName, address)
	store.Delete(key)
}
