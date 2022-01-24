package keeper

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v2/x/subspaces/types"
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

// IterateSubspaceGroups allows iterating over all the groups that are part of the subspace having the given id
func (k Keeper) IterateSubspaceGroups(
	ctx sdk.Context, subspaceID uint64, fn func(index int64, groupName string) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GroupsStoreKey(subspaceID))
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		nameBz := bytes.TrimPrefix(iterator.Key(), types.GroupsStoreKey(subspaceID))
		stop := fn(i, types.GetGroupNameFromBytes(nameBz))
		if stop {
			break
		}
		i++
	}
}

// GetSubspaceGroups returns the list of all groups present inside a given subspace
func (k Keeper) GetSubspaceGroups(ctx sdk.Context, subspaceID uint64) []string {
	var groups []string
	k.IterateSubspaceGroups(ctx, subspaceID, func(index int64, groupName string) (stop bool) {
		groups = append(groups, groupName)
		return false
	})
	return groups
}

// IterateGroupMembers iterates over all the members of the group with the given name present inside the given subspace
func (k Keeper) IterateGroupMembers(
	ctx sdk.Context, subspaceID uint64, groupName string, fn func(index int64, member sdk.AccAddress) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.GroupMembersStoreKey(subspaceID, groupName)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		member := types.GetGroupMemberFromBytes(bytes.TrimPrefix(iterator.Key(), prefix))
		stop := fn(i, member)
		if stop {
			break
		}
		i++
	}
}

// GetGroupMembers iterates returns all the members of a group inside a specific subspace
func (k Keeper) GetGroupMembers(ctx sdk.Context, subspaceID uint64, groupName string) []sdk.AccAddress {
	var members []sdk.AccAddress
	k.IterateGroupMembers(ctx, subspaceID, groupName, func(index int64, member sdk.AccAddress) (stop bool) {
		members = append(members, member)
		return false
	})
	return members
}

// --------------------------------------------------------------------------------------------------------------------

// IteratePermissions iterates over all the permissions set for the subspace with the given id
func (k Keeper) IteratePermissions(
	ctx sdk.Context, subspaceID uint64, fn func(index int64, target string, permission types.Permission) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.PermissionsStoreKey(subspaceID)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		target := types.GetTargetFromBytes(bytes.TrimPrefix(iterator.Key(), prefix))
		permission := types.UnmarshalPermission(iterator.Value())

		stop := fn(i, target, permission)
		if stop {
			break
		}
		i++
	}
}
