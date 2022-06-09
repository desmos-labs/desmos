package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/v3/x/profiles/types"
)

// IterateProfiles iterates through the Profiles set and performs the provided function
func (k Keeper) IterateProfiles(ctx sdk.Context, fn func(index int64, profile *types.Profile) (stop bool)) {
	i := int64(0)
	k.ak.IterateAccounts(ctx, func(account authtypes.AccountI) (stop bool) {
		profile, ok := account.(*types.Profile)

		stop = false
		if ok {
			stop = fn(i, profile)
			i++
		}

		return stop
	})
}

// GetProfiles returns all the profiles that are stored inside the given context
func (k Keeper) GetProfiles(ctx sdk.Context) []*types.Profile {
	var profiles []*types.Profile
	k.IterateProfiles(ctx, func(_ int64, profile *types.Profile) (stop bool) {
		profiles = append(profiles, profile)
		return false
	})
	return profiles
}

// HasProfile returns true iff the given user has a profile, or an error if something is wrong.
func (k Keeper) HasProfile(ctx sdk.Context, user string) bool {
	_, found, err := k.GetProfile(ctx, user)
	if err != nil {
		return false
	}
	return found
}

// --------------------------------------------------------------------------------------------------------------------

// IterateDTagTransferRequests iterates over all the DTag transfer requests and performs the provided function
func (k Keeper) IterateDTagTransferRequests(
	ctx sdk.Context, fn func(index int64, dTagTransferRequest types.DTagTransferRequest) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DTagTransferRequestPrefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		request := types.MustUnmarshalDTagTransferRequest(k.cdc, iterator.Value())
		stop := fn(i, request)
		if stop {
			break
		}
		i++
	}
}

// IterateUserIncomingDTagTransferRequests iterates over all the DTag transfer request made to the given user
// and performs the provided function
func (k Keeper) IterateUserIncomingDTagTransferRequests(
	ctx sdk.Context, user string, fn func(index int64, dTagTransferRequest types.DTagTransferRequest) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.IncomingDTagTransferRequestsPrefix(user))
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		request := types.MustUnmarshalDTagTransferRequest(k.cdc, iterator.Value())
		stop := fn(i, request)
		if stop {
			break
		}
		i++
	}
}

// --------------------------------------------------------------------------------------------------------------------

// IterateApplicationLinks iterates through all the application links and performs the provided function
func (k Keeper) IterateApplicationLinks(ctx sdk.Context, fn func(index int64, link types.ApplicationLink) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ApplicationLinkPrefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		link := types.MustUnmarshalApplicationLink(k.cdc, iterator.Value())
		stop := fn(i, link)
		if stop {
			break
		}
		i++
	}
}

// IterateUserApplicationLinks iterates through all the application links related to the given user
// and performs the provided function
func (k Keeper) IterateUserApplicationLinks(ctx sdk.Context, user string, fn func(index int64, link types.ApplicationLink) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.UserApplicationLinksPrefix(user))
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		link := types.MustUnmarshalApplicationLink(k.cdc, iterator.Value())
		stop := fn(i, link)
		if stop {
			break
		}
		i++
	}
}

// GetApplicationLinks returns a slice of ApplicationLinkEntry objects containing the details of all the
// applications links entries stored inside the current context
func (k Keeper) GetApplicationLinks(ctx sdk.Context) []types.ApplicationLink {
	var links []types.ApplicationLink
	k.IterateApplicationLinks(ctx, func(index int64, link types.ApplicationLink) (stop bool) {
		links = append(links, link)
		return false
	})
	return links
}

// --------------------------------------------------------------------------------------------------------------------

// IterateChainLinks iterates through the chain links and perform the provided function
func (k Keeper) IterateChainLinks(ctx sdk.Context, fn func(index int64, link types.ChainLink) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ChainLinksPrefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		link := types.MustUnmarshalChainLink(k.cdc, iterator.Value())
		stop := fn(i, link)
		if stop {
			break
		}
		i++
	}
}

// IterateUserChainLinks iterates through all the chain links related to the given user and perform the provided function
func (k Keeper) IterateUserChainLinks(ctx sdk.Context, user string, fn func(index int64, link types.ChainLink) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.UserChainLinksPrefix(user))
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		link := types.MustUnmarshalChainLink(k.cdc, iterator.Value())

		stop := fn(i, link)
		if stop {
			break
		}
		i++
	}
}

// IterateUserChainLinksByChain iterates through the chain links related to the given user by the given chain name and perform the provided function
func (k Keeper) IterateUserChainLinksByChain(ctx sdk.Context, user string, chainName string, fn func(link types.ChainLink) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.UserChainLinksChainPrefix(user, chainName))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		link := types.MustUnmarshalChainLink(k.cdc, iterator.Value())

		stop := fn(link)
		if stop {
			break
		}
	}
}

// GetChainLinks allows to returns the list of all stored chain links
func (k Keeper) GetChainLinks(ctx sdk.Context) []types.ChainLink {
	var links []types.ChainLink
	k.IterateChainLinks(ctx, func(_ int64, link types.ChainLink) (stop bool) {
		links = append(links, link)
		return false
	})
	return links
}
