package keeper

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// IterateSubspaces iterates through the subspaces set and performs the given function
func (k Keeper) IterateSubspaces(ctx sdk.Context, fn func(index int64, subspace types.Subspace) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SubspacePrefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var subspace types.Subspace
		k.cdc.MustUnmarshal(iterator.Value(), &subspace)
		stop := fn(i, subspace)
		if stop {
			break
		}
		i++
	}
}

// GetAllSubspaces returns a list of all the subspaces that have been store inside the given context
func (k Keeper) GetAllSubspaces(ctx sdk.Context) []types.Subspace {
	var subspaces []types.Subspace
	k.IterateSubspaces(ctx, func(_ int64, subspace types.Subspace) (stop bool) {
		subspaces = append(subspaces, subspace)
		return false
	})

	return subspaces
}

// --------------------------------------------------------------------------------------------------------------------

// IterateUserGroups iterates over all the users groups stored
func (k Keeper) IterateUserGroups(ctx sdk.Context, fn func(index int64, group types.UserGroup) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GroupsPrefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var group types.UserGroup
		k.cdc.MustUnmarshal(iterator.Value(), &group)
		stop := fn(i, group)
		if stop {
			break
		}
		i++
	}
}

// IterateSubspaceGroups allows iterating over all the groups that are part of the subspace having the given id
func (k Keeper) IterateSubspaceGroups(
	ctx sdk.Context, subspaceID uint64, fn func(index int64, group types.UserGroup) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GroupsStoreKey(subspaceID))
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var group types.UserGroup
		k.cdc.MustUnmarshal(iterator.Value(), &group)
		stop := fn(i, group)
		if stop {
			break
		}
		i++
	}
}

// GetSubspaceGroups returns the list of all groups present inside a given subspace
func (k Keeper) GetSubspaceGroups(ctx sdk.Context, subspaceID uint64) []types.UserGroup {
	var groups []types.UserGroup
	k.IterateSubspaceGroups(ctx, subspaceID, func(index int64, group types.UserGroup) (stop bool) {
		groups = append(groups, group)
		return false
	})
	return groups
}

// IterateGroupMembers iterates over all the members of the group with the given name present inside the given subspace
func (k Keeper) IterateGroupMembers(
	ctx sdk.Context, subspaceID uint64, groupID uint32, fn func(index int64, member sdk.AccAddress) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.GroupMembersStoreKey(subspaceID, groupID)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		member := types.GetAddressFromBytes(bytes.TrimPrefix(iterator.Key(), prefix))

		stop := fn(i, member)
		if stop {
			break
		}
		i++
	}
}

// GetGroupMembers iterates returns all the members of a group inside a specific subspace
func (k Keeper) GetGroupMembers(ctx sdk.Context, subspaceID uint64, groupID uint32) []sdk.AccAddress {
	var members []sdk.AccAddress
	k.IterateGroupMembers(ctx, subspaceID, groupID, func(index int64, member sdk.AccAddress) (stop bool) {
		members = append(members, member)
		return false
	})
	return members
}

// --------------------------------------------------------------------------------------------------------------------

// IterateSubspacePermissions iterates over all the permissions set for the subspace with the given id
func (k Keeper) IterateSubspacePermissions(
	ctx sdk.Context, subspaceID uint64, fn func(index int64, entry types.UserPermission) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.PermissionsStoreKey(subspaceID)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var entry types.UserPermission
		k.cdc.MustUnmarshal(iterator.Value(), &entry)

		stop := fn(i, entry)
		if stop {
			break
		}
		i++
	}
}
