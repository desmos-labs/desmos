package keeper

import (
	"bytes"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

// IterateSubspaces iterates through the subspaces set and performs the given function
func (k Keeper) IterateSubspaces(ctx sdk.Context, fn func(subspace types.Subspace) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SubspacePrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var subspace types.Subspace
		k.cdc.MustUnmarshal(iterator.Value(), &subspace)
		stop := fn(subspace)
		if stop {
			break
		}
	}
}

// GetAllSubspaces returns a list of all the subspaces that have been store inside the given context
func (k Keeper) GetAllSubspaces(ctx sdk.Context) []types.Subspace {
	var subspaces []types.Subspace
	k.IterateSubspaces(ctx, func(subspace types.Subspace) (stop bool) {
		subspaces = append(subspaces, subspace)
		return false
	})

	return subspaces
}

// --------------------------------------------------------------------------------------------------------------------

// IterateSections iterates over all the sections stored and performs the provided function
func (k Keeper) IterateSections(ctx sdk.Context, fn func(section types.Section) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SectionsPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var section types.Section
		k.cdc.MustUnmarshal(iterator.Value(), &section)
		stop := fn(section)
		if stop {
			break
		}
	}
}

// GetAllSections returns all the stored sections
func (k Keeper) GetAllSections(ctx sdk.Context) []types.Section {
	var sections []types.Section
	k.IterateSections(ctx, func(section types.Section) (stop bool) {
		sections = append(sections, section)
		return false
	})
	return sections
}

// IterateSubspaceSections iterates over all the sections for the given subspace and performs the provided function
func (k Keeper) IterateSubspaceSections(ctx sdk.Context, subspaceID uint64, fn func(section types.Section) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SubspaceSectionsPrefix(subspaceID))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var section types.Section
		k.cdc.MustUnmarshal(iterator.Value(), &section)
		stop := fn(section)
		if stop {
			break
		}
	}
}

// GetSubspaceSections returns all the sections for the given subspace
func (k Keeper) GetSubspaceSections(ctx sdk.Context, subspaceID uint64) []types.Section {
	var sections []types.Section
	k.IterateSubspaceSections(ctx, subspaceID, func(section types.Section) (stop bool) {
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
	if section.ID == types.RootSectionID || stop {
		// End the iteration only if the user has told us to stop, or if we reached the root
		return
	}

	// Continue to follow the path from the parent section up to the root
	k.IterateSectionPath(ctx, section.SubspaceID, section.ParentID, fn)
}

// IterateSectionChildren iterates over all the children of the given section and performs the provided function
func (k Keeper) IterateSectionChildren(ctx sdk.Context, subspaceID uint64, sectionID uint32, fn func(section types.Section) (stop bool)) {
	k.IterateSubspaceSections(ctx, subspaceID, func(section types.Section) (stop bool) {
		stop = false
		if section.ID != sectionID && section.ParentID == sectionID {
			stop = fn(section)
		}
		return stop
	})
}

// --------------------------------------------------------------------------------------------------------------------

// IterateUserGroups iterates over all the users groups stored
func (k Keeper) IterateUserGroups(ctx sdk.Context, fn func(group types.UserGroup) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GroupsPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var group types.UserGroup
		k.cdc.MustUnmarshal(iterator.Value(), &group)
		stop := fn(group)
		if stop {
			break
		}
	}
}

// GetAllUserGroups returns the information (name and members) for all the groups of all the subspaces
func (k Keeper) GetAllUserGroups(ctx sdk.Context) []types.UserGroup {
	var groups []types.UserGroup
	k.IterateUserGroups(ctx, func(group types.UserGroup) (stop bool) {
		groups = append(groups, group)
		return false
	})
	return groups
}

// IterateSubspaceUserGroups allows iterating over all the groups that are part of the subspace having the given id
func (k Keeper) IterateSubspaceUserGroups(ctx sdk.Context, subspaceID uint64, fn func(group types.UserGroup) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SubspaceGroupsPrefix(subspaceID))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var group types.UserGroup
		k.cdc.MustUnmarshal(iterator.Value(), &group)
		stop := fn(group)
		if stop {
			break
		}
	}
}

// GetSubspaceUserGroups returns the list of all groups present inside a given subspace
func (k Keeper) GetSubspaceUserGroups(ctx sdk.Context, subspaceID uint64) []types.UserGroup {
	var groups []types.UserGroup
	k.IterateSubspaceUserGroups(ctx, subspaceID, func(group types.UserGroup) (stop bool) {
		groups = append(groups, group)
		return false
	})
	return groups
}

// IterateSectionUserGroups iterates over all the user groups for the given section and performs the provided function
func (k Keeper) IterateSectionUserGroups(ctx sdk.Context, subspaceID uint64, sectionID uint32, fn func(group types.UserGroup) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SectionGroupsPrefix(subspaceID, sectionID))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var group types.UserGroup
		k.cdc.MustUnmarshal(iterator.Value(), &group)

		// Perform the function only on
		stop := fn(group)
		if stop {
			break
		}
	}
}

// GetSectionUserGroups returns all the user groups present inside the given section
func (k Keeper) GetSectionUserGroups(ctx sdk.Context, subspaceID uint64, sectionID uint32) []types.UserGroup {
	var groups []types.UserGroup
	k.IterateSectionUserGroups(ctx, subspaceID, sectionID, func(group types.UserGroup) (stop bool) {
		groups = append(groups, group)
		return false
	})
	return groups
}

// --------------------------------------------------------------------------------------------------------------------

// IterateUserGroupsMembers iterates over all the group member entries and performs the provided function
func (k Keeper) IterateUserGroupsMembers(ctx sdk.Context, fn func(entry types.UserGroupMemberEntry) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.GroupsMembersPrefix
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		subspaceID, groupID, user := types.SplitGroupMemberStoreKey(iterator.Key())
		stop := fn(types.NewUserGroupMemberEntry(subspaceID, groupID, user))
		if stop {
			break
		}
	}
}

// IterateUserGroupMembers iterates over all the members of the group with the given id present inside the given subspace
func (k Keeper) IterateUserGroupMembers(ctx sdk.Context, subspaceID uint64, groupID uint32, fn func(member string) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.GroupMembersPrefix(subspaceID, groupID)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		member := types.GetAddressFromBytes(bytes.TrimPrefix(iterator.Key(), prefix))

		stop := fn(member)
		if stop {
			break
		}
	}
}

// GetUserGroupMembers returns all the members of a group inside a specific subspace
func (k Keeper) GetUserGroupMembers(ctx sdk.Context, subspaceID uint64, groupID uint32) []string {
	var members []string
	k.IterateUserGroupMembers(ctx, subspaceID, groupID, func(member string) (stop bool) {
		members = append(members, member)
		return false
	})
	return members
}

// --------------------------------------------------------------------------------------------------------------------

// IterateUserPermissions iterates over all the stored user permissions
func (k Keeper) IterateUserPermissions(ctx sdk.Context, fn func(entry types.UserPermission) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.UserPermissionsStorePrefix
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var permission types.UserPermission
		k.cdc.MustUnmarshal(iterator.Value(), &permission)

		stop := fn(permission)
		if stop {
			break
		}
	}
}

// IterateSubspaceUserPermissions iterates over all the user permissions set for the subspace with the given id
func (k Keeper) IterateSubspaceUserPermissions(ctx sdk.Context, subspaceID uint64, fn func(entry types.UserPermission) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.SubspacePermissionsPrefix(subspaceID)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var permission types.UserPermission
		k.cdc.MustUnmarshal(iterator.Value(), &permission)

		stop := fn(permission)
		if stop {
			break
		}
	}
}

// GetSubspaceUserPermissions returns all the user permissions set for the given subspace
func (k Keeper) GetSubspaceUserPermissions(ctx sdk.Context, subspaceID uint64) []types.UserPermission {
	var entries []types.UserPermission
	k.IterateSubspaceUserPermissions(ctx, subspaceID, func(entry types.UserPermission) (stop bool) {
		entries = append(entries, entry)
		return false
	})
	return entries
}

// IterateSectionUserPermissions iterates over all the permissions set for the given section and performs the provided function
func (k Keeper) IterateSectionUserPermissions(ctx sdk.Context, subspaceID uint64, sectionID uint32, fn func(entry types.UserPermission) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.SectionPermissionsPrefix(subspaceID, sectionID)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var permission types.UserPermission
		k.cdc.MustUnmarshal(iterator.Value(), &permission)

		stop := fn(permission)
		if stop {
			break
		}
	}
}

// GetSectionUserPermissions returns all the user permissions set inside the specific section
func (k Keeper) GetSectionUserPermissions(ctx sdk.Context, subspaceID uint64, sectionID uint32) []types.UserPermission {
	var entries []types.UserPermission
	k.IterateSectionUserPermissions(ctx, subspaceID, sectionID, func(entry types.UserPermission) (stop bool) {
		entries = append(entries, entry)
		return false
	})
	return entries
}

// --------------------------------------------------------------------------------------------------------------------

func (k Keeper) GetAllGrants(ctx sdk.Context) []types.Grant {
	return append(k.GetAllUserGrants(ctx), k.GetAllUserGroupsGrants(ctx)...)
}

// IterateUserGrants iterates through the user grants and performs the given function
func (k Keeper) IterateUserGrants(ctx sdk.Context, fn func(grant types.Grant) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.UserAllowancePrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var grant types.Grant
		k.cdc.MustUnmarshal(iterator.Value(), &grant)
		stop := fn(grant)
		if stop {
			break
		}
	}
}

// IterateSubspaceUserGrants iterates over all the user grants inside the subspace with the given id
func (k Keeper) IterateSubspaceUserGrants(ctx sdk.Context, subspaceID uint64, fn func(grant types.Grant) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SubspaceUserAllowancePrefix(subspaceID))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var grant types.Grant
		k.cdc.MustUnmarshal(iterator.Value(), &grant)
		stop := fn(grant)
		if stop {
			break
		}
	}
}

// GetAllUserGrants returns a list of all the user grants that have been store inside the given context
func (k Keeper) GetAllUserGrants(ctx sdk.Context) []types.Grant {
	var grants []types.Grant
	k.IterateUserGrants(ctx, func(grant types.Grant) (stop bool) {
		grants = append(grants, grant)
		return false
	})

	return grants
}

// GetSubspaceUserGrants returns all the user grants inside the given subspace
func (k Keeper) GetSubspaceUserGrants(ctx sdk.Context, subspaceID uint64) []types.Grant {
	var grants []types.Grant
	k.IterateSubspaceUserGrants(ctx, subspaceID, func(grant types.Grant) (stop bool) {
		grants = append(grants, grant)
		return false
	})
	return grants
}

// --------------------------------------------------------------------------------------------------------------------

// IterateUserGroupsGrants iterates through the group grants and performs the given function
func (k Keeper) IterateUserGroupsGrants(ctx sdk.Context, fn func(grant types.Grant) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GroupAllowancePrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var grant types.Grant
		k.cdc.MustUnmarshal(iterator.Value(), &grant)
		stop := fn(grant)
		if stop {
			break
		}
	}
}

// IterateSubspaceUserGroupGrants iterates over all the group grants inside the subspace with the given id
func (k Keeper) IterateSubspaceUserGroupGrants(ctx sdk.Context, subspaceID uint64, fn func(grant types.Grant) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	prefix := types.SubspaceGroupAllowancePrefix(subspaceID)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var grant types.Grant
		k.cdc.MustUnmarshal(iterator.Value(), &grant)
		stop := fn(grant)
		if stop {
			break
		}
	}
}

// IterateUserGroupGrants iterates over all the group grants inside the group with the given id
func (k Keeper) IterateUserGroupGrants(ctx sdk.Context, subspaceID uint64, groupID uint32, fn func(grant types.Grant) (stop bool)) {
	k.IterateSubspaceUserGroupGrants(ctx, subspaceID, func(grant types.Grant) (stop bool) {
		if grant.Target.GetCachedValue().(*types.GroupTarget).GroupID == groupID {
			stop = fn(grant)
		}
		return stop
	})
}

// IterateSubspaceGranterGroupGrants iterates over all the group grants for the given granter and performs the provided function
func (k Keeper) IterateSubspaceGranterGroupGrants(ctx sdk.Context, subspaceID uint64, granter string, fn func(entry types.Grant) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	prefix := types.GranterGroupAllowancePrefix(subspaceID, granter)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var grant types.Grant
		k.cdc.MustUnmarshal(iterator.Value(), &grant)

		stop := fn(grant)
		if stop {
			break
		}
	}
}

// GetAllUserGroupsGrants returns a list of all the group grants that have been store inside the given context
func (k Keeper) GetAllUserGroupsGrants(ctx sdk.Context) []types.Grant {
	var grants []types.Grant
	k.IterateUserGroupsGrants(ctx, func(grant types.Grant) (stop bool) {
		grants = append(grants, grant)
		return false
	})

	return grants
}

// GetSubspaceUserGroupsGrants returns all the group grants inside the given subspace
func (k Keeper) GetSubspaceUserGroupsGrants(ctx sdk.Context, subspaceID uint64) []types.Grant {
	var grants []types.Grant
	k.IterateSubspaceUserGroupGrants(ctx, subspaceID, func(grant types.Grant) (stop bool) {
		grants = append(grants, grant)
		return false
	})
	return grants
}

// GetUserGroupGrants returns all the group grants inside the given group
func (k Keeper) GetUserGroupGrants(ctx sdk.Context, subspaceID uint64, groupID uint32) []types.Grant {
	var grants []types.Grant
	k.IterateUserGroupGrants(ctx, subspaceID, groupID, func(grant types.Grant) (stop bool) {
		grants = append(grants, grant)
		return false
	})
	return grants
}

// --------------------------------------------------------------------------------------------------------------------

// createAccount creates a new account for the given user who does not exist on chain
func (k Keeper) creatAccount(ctx sdk.Context, user string) {
	addr, err := sdk.AccAddressFromBech32(user)
	if err != nil {
		panic(err)
	}
	if !k.ak.HasAccount(ctx, addr) {
		defer telemetry.IncrCounter(1, "new", "account")
		k.ak.SetAccount(ctx, authtypes.NewBaseAccountWithAddress(addr))
	}
}
