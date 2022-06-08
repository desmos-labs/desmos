package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	poststypes "github.com/desmos-labs/desmos/v3/x/posts/types"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"

	"github.com/desmos-labs/desmos/v3/x/reactions/types"
)

// HasSubspace tells whether the subspace with the given id exists or not
func (k Keeper) HasSubspace(ctx sdk.Context, subspaceID uint64) bool {
	return k.sk.HasSubspace(ctx, subspaceID)
}

// HasPermission tells whether the given user has the provided permission inside the subspace with the specified id
func (k Keeper) HasPermission(ctx sdk.Context, subspaceID uint64, user string, permission subspacestypes.Permission) bool {
	// Report-related permissions are checked only against the root section
	return k.sk.HasPermission(ctx, subspaceID, subspacestypes.RootSectionID, user, permission)
}

// HasUserBlocked tells whether the given blocker has blocked the user inside the provided subspace
func (k Keeper) HasUserBlocked(ctx sdk.Context, blocker, user string, subspaceID uint64) bool {
	return k.rk.HasUserBlocked(ctx, blocker, user, subspaceID)
}

// HasPost tells whether the given post exists or not
func (k Keeper) HasPost(ctx sdk.Context, subspaceID uint64, postID uint64) bool {
	return k.pk.HasPost(ctx, subspaceID, postID)
}

// GetPost returns the post associated with the given id
func (k Keeper) GetPost(ctx sdk.Context, subspaceID uint64, postID uint64) (poststypes.Post, bool) {
	return k.pk.GetPost(ctx, subspaceID, postID)
}

// --------------------------------------------------------------------------------------------------------------------

// IterateRegisteredReactions iterates over all the registered reactions and performs the provided function
func (k Keeper) IterateRegisteredReactions(ctx sdk.Context, fn func(reaction types.RegisteredReaction) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RegisteredReactionPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var registeredReaction types.RegisteredReaction
		k.cdc.MustUnmarshal(iterator.Value(), &registeredReaction)

		stop := fn(registeredReaction)
		if stop {
			break
		}
	}
}

// GetRegisteredReactions returns all the stored registered reactions
func (k Keeper) GetRegisteredReactions(ctx sdk.Context) []types.RegisteredReaction {
	var reactions []types.RegisteredReaction
	k.IterateRegisteredReactions(ctx, func(reaction types.RegisteredReaction) (stop bool) {
		reactions = append(reactions, reaction)
		return false
	})
	return reactions
}

// IterateSubspaceRegisteredReactions iterates over all the given subspace registered reactions and performs the provided function
func (k Keeper) IterateSubspaceRegisteredReactions(ctx sdk.Context, subspaceID uint64, fn func(reaction types.RegisteredReaction) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SubspaceRegisteredReactionsPrefix(subspaceID))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var registeredReaction types.RegisteredReaction
		k.cdc.MustUnmarshal(iterator.Value(), &registeredReaction)

		stop := fn(registeredReaction)
		if stop {
			break
		}
	}
}

// --------------------------------------------------------------------------------------------------------------------

// IterateReactions iterates over all the reactions and performs the provided function
func (k Keeper) IterateReactions(ctx sdk.Context, fn func(reaction types.Reaction) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ReactionPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var reaction types.Reaction
		k.cdc.MustUnmarshal(iterator.Value(), &reaction)

		stop := fn(reaction)
		if stop {
			break
		}
	}
}

// GetReactions returns all the stored reactions
func (k Keeper) GetReactions(ctx sdk.Context) []types.Reaction {
	var reactions []types.Reaction
	k.IterateReactions(ctx, func(reaction types.Reaction) (stop bool) {
		reactions = append(reactions, reaction)
		return false
	})
	return reactions
}

// IterateSubspaceReactions iterates over all the given subspace reactions and performs the provided function
func (k Keeper) IterateSubspaceReactions(ctx sdk.Context, subspaceID uint64, fn func(reaction types.Reaction) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SubspaceReactionsPrefix(subspaceID))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var reaction types.Reaction
		k.cdc.MustUnmarshal(iterator.Value(), &reaction)

		stop := fn(reaction)
		if stop {
			break
		}
	}
}

// IteratePostReactions iterates over all the given post reactions and performs the provided function
func (k Keeper) IteratePostReactions(ctx sdk.Context, subspaceID uint64, postID uint64, fn func(reaction types.Reaction) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PostReactionsPrefix(subspaceID, postID))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var reaction types.Reaction
		k.cdc.MustUnmarshal(iterator.Value(), &reaction)

		stop := fn(reaction)
		if stop {
			break
		}
	}
}

// --------------------------------------------------------------------------------------------------------------------

// IterateReactionsParams iterates over all the stored subspace reactions params and performs the provided function
func (k Keeper) IterateReactionsParams(ctx sdk.Context, fn func(params types.SubspaceReactionsParams) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ReactionsParamsPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var params types.SubspaceReactionsParams
		k.cdc.MustUnmarshal(iterator.Value(), &params)

		stop := fn(params)
		if stop {
			break
		}
	}
}

// GetReactionsParams returns all the stored reactions parameters
func (k Keeper) GetReactionsParams(ctx sdk.Context) []types.SubspaceReactionsParams {
	var parameters []types.SubspaceReactionsParams
	k.IterateReactionsParams(ctx, func(params types.SubspaceReactionsParams) (stop bool) {
		parameters = append(parameters, params)
		return false
	})
	return parameters
}
