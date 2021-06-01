package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// StoreChainLink stores the given chain link inside the current context.
// It assumes that the given chain link has already been validated.
func (k Keeper) StoreChainLink(ctx sdk.Context, user string, link types.ChainLink) error {
	if _, found := k.GetChainLink(ctx, link.ChainConfig.Name, link.Address); found {
		return fmt.Errorf("chain link already exists")
	}

	// check target address has a profile or not
	if link.ChainConfig.Name == "desmos" {
		if _, err := sdk.AccAddressFromBech32(link.Address); err != nil {
			return err
		}
		_, found, err := k.GetProfile(ctx, link.Address)
		if err != nil {
			return err
		}
		if found {
			return fmt.Errorf("the target address has profiles already")
		}
	}

	// Check if address has the profile
	profile, found, err := k.GetProfile(ctx, user)
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

// DeleteLink allows to delete a link associated with the given address and chain name inside the current context.
// It assumes that the related link exists.
func (k Keeper) DeleteChainLink(ctx sdk.Context, owner, chainName, target string) error {
	// Check if address has the profile and get the profile
	profile, found, err := k.GetProfile(ctx, owner)
	if err != nil {
		return err
	}
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, ("non existent profile on owner address"))
	}

	isTargetExist := false
	newChainsLinks := []types.ChainLink{}
	// Try to find the target link
	for _, link := range profile.ChainsLinks {
		currChainName := link.ChainConfig.Name
		currAddr := link.Address
		if currChainName == chainName && currAddr == target {
			isTargetExist = true
			continue
		}
		newChainsLinks = append(newChainsLinks, link)
	}

	if !isTargetExist {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, ("non existent target chain link in the profile"))
	}
	// Update profile status
	profile.ChainsLinks = newChainsLinks
	err = k.StoreProfile(ctx, profile)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	key := types.ChainsLinksStoreKey(chainName, target)
	store.Delete(key)
	return nil
}
