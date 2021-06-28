package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
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

// --------------------------------------------------------------------------------------------------------------------

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

// IterateRelationships iterates through the relationships and perform the provided function
func (k Keeper) IterateRelationships(ctx sdk.Context, fn func(index int64, relationship types.Relationship) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.RelationshipsStorePrefix)
	defer iterator.Close()

	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {
		relationship := types.MustUnmarshalRelationship(k.cdc, iterator.Value())

		stop := fn(i, relationship)

		if stop {
			break
		}

		i++
	}
}

// IterateUserRelationships iterates through the relationships with the given user address
// and performs the provided function
func (k Keeper) IterateUserRelationships(ctx sdk.Context, user string, fn func(index int64, relationship types.Relationship) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.UserRelationshipsPrefix(user))
	defer iterator.Close()

	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {
		relationship := types.MustUnmarshalRelationship(k.cdc, iterator.Value())

		stop := fn(i, relationship)

		if stop {
			break
		}
		i++
	}
}

// IterateBlockedUsers iterates through the list of users blocked by the specified user and performs the given function
func (k Keeper) IterateBlockedUsers(ctx sdk.Context, user string, fn func(index int64, blocks types.UserBlock) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.BlockerPrefix(user))
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		block := types.MustUnmarshalUserBlock(k.cdc, iterator.Value())
		stop := fn(i, block)
		if stop {
			break
		}
		i++
	}
}

// --------------------------------------------------------------------------------------------------------------------

// IterateUserApplicationLinks iterates through all the application links realted to the given user
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

	k.ak.IterateAccounts(ctx, func(account authtypes.AccountI) (stop bool) {
		k.IterateUserApplicationLinks(ctx, account.GetAddress().String(), func(_ int64, link types.ApplicationLink) (stop bool) {
			links = append(links, link)
			return false
		})

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
