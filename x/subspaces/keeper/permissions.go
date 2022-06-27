package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

// SetUserPermissions sets the given permission for the specific user inside a single subspace
func (k Keeper) SetUserPermissions(ctx sdk.Context, subspaceID uint64, sectionID uint32, user string, permissions types.Permissions) {
	store := ctx.KVStore(k.storeKey)
	permission := types.NewUserPermission(subspaceID, sectionID, user, permissions)
	store.Set(types.UserPermissionStoreKey(subspaceID, sectionID, user), k.cdc.MustMarshal(&permission))

	k.AfterUserPermissionSet(ctx, subspaceID, sectionID, user, permissions)
}

// HasPermission checks whether the specific user has the given permission inside a specific subspace
func (k Keeper) HasPermission(ctx sdk.Context, subspaceID uint64, sectionID uint32, user string, permission types.Permission) bool {
	// Get the subspace to make sure the request is valid
	subspace, found := k.GetSubspace(ctx, subspaceID)
	if !found {
		return false
	}

	// The owner of the subspaces has all the permissions by default
	if subspace.Owner == user {
		return true
	}

	// Get the permissions set to the specific user
	specificPermissions := k.GetUserPermissions(ctx, subspaceID, sectionID, user)

	// Get the group permissions
	groupPermissions := k.GetGroupsInheritedPermissions(ctx, subspaceID, sectionID, user)

	// Check the combination of the permissions
	permissions := append(specificPermissions, groupPermissions...)
	return types.CheckPermission(types.CombinePermissions(permissions...), permission)
}

// getSectionPermissions gets the permissions for the given user set inside the specified section only
func (k Keeper) getSectionPermissions(ctx sdk.Context, subspaceID uint64, sectionID uint32, user string) types.Permissions {
	store := ctx.KVStore(k.storeKey)
	var permissions types.UserPermission
	k.cdc.MustUnmarshal(store.Get(types.UserPermissionStoreKey(subspaceID, sectionID, user)), &permissions)
	return permissions.Permissions
}

// GetUserPermissions returns the permissions that are currently set inside
// the subspace with the given id for the given user
func (k Keeper) GetUserPermissions(ctx sdk.Context, subspaceID uint64, sectionID uint32, user string) types.Permissions {
	if sectionID == types.RootSectionID {
		return k.getSectionPermissions(ctx, subspaceID, sectionID, user)
	}

	sectionPermissions := k.getSectionPermissions(ctx, subspaceID, sectionID, user)

	// Get the parent permissions
	var parentPermissions types.Permissions
	section, found := k.GetSection(ctx, subspaceID, sectionID)
	if found {
		parentPermissions = k.GetUserPermissions(ctx, subspaceID, section.ParentID, user)
	}

	return types.CombinePermissions(append(parentPermissions, sectionPermissions...)...)
}

// GetGroupsInheritedPermissions returns the permissions that the specified user
// has inherited from all the groups that they are part of.
func (k Keeper) GetGroupsInheritedPermissions(ctx sdk.Context, subspaceID uint64, sectionID uint32, user string) types.Permissions {
	var permissions []types.Permission

	// Iterate over the section ancestors and get all the user groups for each ancestor
	// to check if the user is part of a group
	k.IterateSectionPath(ctx, subspaceID, sectionID, func(section types.Section) (stop bool) {
		k.IterateSectionUserGroups(ctx, section.SubspaceID, section.ID, func(group types.UserGroup) (stop bool) {
			if k.IsMemberOfGroup(ctx, subspaceID, group.ID, user) {
				permissions = append(permissions, group.Permissions...)
			}
			return false
		})
		return false
	})

	return types.CombinePermissions(permissions...)
}

// GetUsersWithRootPermissions returns all the users that have a given permission inside the specified subspace
func (k Keeper) GetUsersWithRootPermissions(ctx sdk.Context, subspaceID uint64, permission types.Permissions) []string {
	subspace, found := k.GetSubspace(ctx, subspaceID)
	if !found {
		return nil
	}

	// The owner must always be included as they have all the permissions
	users := []string{subspace.Owner}

	// Iterate over the various groups
	k.IterateSectionUserGroups(ctx, subspaceID, types.RootSectionID, func(group types.UserGroup) (stop bool) {
		if !types.CheckPermissions(group.Permissions, permission) {
			// Return early if the group does not have the permission. We will check other groups anyway
			return false
		}

		// If the group has the permission, get all the members
		users = append(users, k.GetUserGroupMembers(ctx, subspaceID, group.ID)...)
		return false
	})

	// Iterate over the various individually-set permissions
	k.IterateSectionUserPermissions(ctx, subspaceID, types.RootSectionID, func(entry types.UserPermission) (stop bool) {
		if types.CheckPermissions(entry.Permissions, permission) {
			users = append(users, entry.User)
		}

		return false
	})

	return users
}

// RemoveUserPermissions removes the permission for the given user inside the provided subspace
func (k Keeper) RemoveUserPermissions(ctx sdk.Context, subspaceID uint64, sectionID uint32, user string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.UserPermissionStoreKey(subspaceID, sectionID, user))

	k.AfterUserPermissionRemoved(ctx, subspaceID, sectionID, user)
}
