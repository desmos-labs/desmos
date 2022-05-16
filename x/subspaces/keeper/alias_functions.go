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

// IterateSections iterates over all the sections stored and performs the provided function
func (k Keeper) IterateSections(ctx sdk.Context, fn func(index int64, section types.Section) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SectionsPrefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var section types.Section
		k.cdc.MustUnmarshal(iterator.Value(), &section)
		stop := fn(i, section)
		if stop {
			break
		}
		i++
	}
}

// GetAllSections returns all the stored sections
func (k Keeper) GetAllSections(ctx sdk.Context) []types.Section {
	var sections []types.Section
	k.IterateSections(ctx, func(index int64, section types.Section) (stop bool) {
		sections = append(sections, section)
		return false
	})
	return sections
}

// IterateSubspaceSections iterates over all the sections for the given subspace and performs the provided function
func (k Keeper) IterateSubspaceSections(ctx sdk.Context, subspaceID uint64, fn func(index int64, section types.Section) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SubspaceSectionsPrefix(subspaceID))
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var section types.Section
		k.cdc.MustUnmarshal(iterator.Value(), &section)
		stop := fn(i, section)
		if stop {
			break
		}
		i++
	}
}

// IterateSectionPath iterates the path that leads from the section having the given id up towards the root section
// and performs the provided function on all the sections that are encountered over the path (including
// the initial section having the specified id).
func (k Keeper) IterateSectionPath(ctx sdk.Context, subspaceID uint64, sectionID uint32, fn func(section types.Section) (stop bool)) {
	section, found := k.GetSection(ctx, subspaceID, sectionID)
	if !found {
		return
	}

	stop := fn(section)
	if section.ID == 0 || stop {
		// End the iteration only if the user has told us to stop, or if we reached the root (section id 0)
		return
	}

	// Continue to follow the path from the parent section up to the root
	k.IterateSectionPath(ctx, section.SubspaceID, section.ParentID, fn)
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

// GetAllUserGroups returns the information (name and members) for all the groups of all the subspaces
func (k Keeper) GetAllUserGroups(ctx sdk.Context) []types.UserGroup {
	var groups []types.UserGroup
	k.IterateUserGroups(ctx, func(index int64, group types.UserGroup) (stop bool) {
		groups = append(groups, group)
		return false
	})
	return groups
}

// IterateSubspaceUserGroups allows iterating over all the groups that are part of the subspace having the given id
func (k Keeper) IterateSubspaceUserGroups(ctx sdk.Context, subspaceID uint64, fn func(index int64, group types.UserGroup) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SubspaceGroupsPrefix(subspaceID))
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

// GetSubspaceUserGroups returns the list of all groups present inside a given subspace
func (k Keeper) GetSubspaceUserGroups(ctx sdk.Context, subspaceID uint64) []types.UserGroup {
	var groups []types.UserGroup
	k.IterateSubspaceUserGroups(ctx, subspaceID, func(index int64, group types.UserGroup) (stop bool) {
		groups = append(groups, group)
		return false
	})
	return groups
}

// IterateSectionUserGroups iterates over all the user groups for the given section and performs the provided function
func (k Keeper) IterateSectionUserGroups(ctx sdk.Context, subspaceID uint64, sectionID uint32, fn func(index int64, group types.UserGroup) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SectionGroupsPrefix(subspaceID, sectionID))
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var group types.UserGroup
		k.cdc.MustUnmarshal(iterator.Value(), &group)

		// Perform the function only on
		stop := fn(i, group)
		if stop {
			break
		}
		i++
	}
}

// IterateUserGroupMembers iterates over all the members of the group with the given name present inside the given subspace
func (k Keeper) IterateUserGroupMembers(ctx sdk.Context, subspaceID uint64, groupID uint32, fn func(index int64, member sdk.AccAddress) (stop bool)) {
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

// GetUserGroupMembers iterates returns all the members of a group inside a specific subspace
func (k Keeper) GetUserGroupMembers(ctx sdk.Context, subspaceID uint64, groupID uint32) []sdk.AccAddress {
	var members []sdk.AccAddress
	k.IterateUserGroupMembers(ctx, subspaceID, groupID, func(index int64, member sdk.AccAddress) (stop bool) {
		members = append(members, member)
		return false
	})
	return members
}

// --------------------------------------------------------------------------------------------------------------------

// IterateUserPermissions iterates over all the stored user permissions
func (k Keeper) IterateUserPermissions(ctx sdk.Context, fn func(index int64, subspaceID uint64, sectionID uint32, user sdk.AccAddress, permission types.Permission) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.UserPermissionsStorePrefix
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		subspaceID, sectionID, user := types.SplitUserAddressPermissionKey(append(prefix, iterator.Key()...))
		permission := types.UnmarshalPermission(iterator.Value())

		stop := fn(i, subspaceID, sectionID, user, permission)
		if stop {
			break
		}
		i++
	}
}

// IterateSubspaceUserPermissions iterates over all the user permissions set for the subspace with the given id
func (k Keeper) IterateSubspaceUserPermissions(ctx sdk.Context, subspaceID uint64, fn func(index int64, sectionID uint32, user sdk.AccAddress, permission types.Permission) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.SubspacePermissionsPrefix(subspaceID)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		_, sectionID, user := types.SplitUserAddressPermissionKey(append(prefix, iterator.Key()...))
		permission := types.UnmarshalPermission(iterator.Value())

		stop := fn(i, sectionID, user, permission)
		if stop {
			break
		}
		i++
	}
}

// IterateSectionUserPermissions iterates over all the permissions set for the given section and performs the provided function
func (k Keeper) IterateSectionUserPermissions(ctx sdk.Context, subspaceID uint64, sectionID uint32, fn func(index int64, user sdk.AccAddress, permission types.Permission) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.SectionPermissionsPrefix(subspaceID, sectionID)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		_, _, user := types.SplitUserAddressPermissionKey(append(prefix, iterator.Key()...))
		permission := types.UnmarshalPermission(iterator.Value())

		stop := fn(i, user, permission)
		if stop {
			break
		}
		i++
	}
}
