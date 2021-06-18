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

// IterateUserRelationships iterates through the relationships with the given user address and performs the provided function
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
