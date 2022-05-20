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

// GetSubspaceSections returns all the sections for the given subspace
func (k Keeper) GetSubspaceSections(ctx sdk.Context, subspaceID uint64) []types.Section {
	var sections []types.Section
	k.IterateSubspaceSections(ctx, subspaceID, func(index int64, section types.Section) (stop bool) {
		sections = append(sections, section)
		return false
	})
	return sections
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

// IterateSectionChildren iterates over all the children of the given section and performs the provided function
func (k Keeper) IterateSectionChildren(ctx sdk.Context, subspaceID uint64, sectionID uint32, fn func(index int64, section types.Section) (stop bool)) {
	index := int64(0)
	k.IterateSubspaceSections(ctx, subspaceID, func(_ int64, section types.Section) (stop bool) {
		stop = false
		if section.ID != sectionID && section.ParentID == sectionID {
			stop = fn(index, section)
			index++
		}
		return stop
	})
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

// GetSectionUserGroups returns all the user groups present inside the given section
func (k Keeper) GetSectionUserGroups(ctx sdk.Context, subspaceID uint64, sectionID uint32) []types.UserGroup {
	var groups []types.UserGroup
	k.IterateSectionUserGroups(ctx, subspaceID, sectionID, func(index int64, group types.UserGroup) (stop bool) {
		groups = append(groups, group)
		return false
	})
	return groups
}

// --------------------------------------------------------------------------------------------------------------------

// IterateUserGroupsMembers iterates over all the group member entries and performs the provided function
func (k Keeper) IterateUserGroupsMembers(ctx sdk.Context, fn func(index int64, entry types.UserGroupMemberEntry) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.GroupsMembersPrefix
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		subspaceID, groupID, user := types.SplitGroupMemberStoreKey(iterator.Key())
		stop := fn(i, types.NewUserGroupMemberEntry(subspaceID, groupID, user))
		if stop {
			break
		}
		i++
	}
}

// IterateUserGroupMembers iterates over all the members of the group with the given id present inside the given subspace
func (k Keeper) IterateUserGroupMembers(ctx sdk.Context, subspaceID uint64, groupID uint32, fn func(index int64, member sdk.AccAddress) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.GroupMembersPrefix(subspaceID, groupID)
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

// GetUserGroupMembers returns all the members of a group inside a specific subspace
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
func (k Keeper) IterateUserPermissions(ctx sdk.Context, fn func(index int64, entry types.UserPermission) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.UserPermissionsStorePrefix
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		subspaceID, sectionID, user := types.SplitUserAddressPermissionKey(iterator.Key())
		permission := types.UnmarshalPermission(iterator.Value())

		stop := fn(i, types.NewUserPermission(subspaceID, sectionID, user, permission))
		if stop {
			break
		}
		i++
	}
}

// IterateSubspaceUserPermissions iterates over all the user permissions set for the subspace with the given id
func (k Keeper) IterateSubspaceUserPermissions(ctx sdk.Context, subspaceID uint64, fn func(index int64, entry types.UserPermission) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.SubspacePermissionsPrefix(subspaceID)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		subspaceID, sectionID, user := types.SplitUserAddressPermissionKey(iterator.Key())
		permission := types.UnmarshalPermission(iterator.Value())

		stop := fn(i, types.NewUserPermission(subspaceID, sectionID, user, permission))
		if stop {
			break
		}
		i++
	}
}

// GetSubspaceUserPermissions returns all the user permissions set for the given subspace
func (k Keeper) GetSubspaceUserPermissions(ctx sdk.Context, subspaceID uint64) []types.UserPermission {
	var entries []types.UserPermission
	k.IterateSubspaceUserPermissions(ctx, subspaceID, func(index int64, entry types.UserPermission) (stop bool) {
		entries = append(entries, entry)
		return false
	})
	return entries
}

// IterateSectionUserPermissions iterates over all the permissions set for the given section and performs the provided function
func (k Keeper) IterateSectionUserPermissions(ctx sdk.Context, subspaceID uint64, sectionID uint32, fn func(index int64, entry types.UserPermission) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.SectionPermissionsPrefix(subspaceID, sectionID)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		subspaceID, sectionID, user := types.SplitUserAddressPermissionKey(iterator.Key())
		permission := types.UnmarshalPermission(iterator.Value())

		stop := fn(i, types.NewUserPermission(subspaceID, sectionID, user, permission))
		if stop {
			break
		}
		i++
	}
}

// GetSectionUserPermissions returns all the user permissions set inside the specific section
func (k Keeper) GetSectionUserPermissions(ctx sdk.Context, subspaceID uint64, sectionID uint32) []types.UserPermission {
	var entries []types.UserPermission
	k.IterateSectionUserPermissions(ctx, subspaceID, sectionID, func(index int64, entry types.UserPermission) (stop bool) {
		entries = append(entries, entry)
		return false
	})
	return entries
}
