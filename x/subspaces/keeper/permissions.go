package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

// SetUserPermissions sets the given permission for the specific user inside a single subspace
func (k Keeper) SetUserPermissions(ctx sdk.Context, subspaceID uint64, user sdk.AccAddress, permissions types.Permission) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.UserPermissionStoreKey(subspaceID, user), types.MarshalPermission(permissions))

	k.AfterUserPermissionSet(ctx, subspaceID, user, permissions)
}

// HasPermission checks whether the specific user has the given permission inside a specific subspace
func (k Keeper) HasPermission(ctx sdk.Context, subspaceID uint64, user sdk.AccAddress, permission types.Permission) bool {
	// Get the subspace to make sure the request is valid
	subspace, found := k.GetSubspace(ctx, subspaceID)
	if !found {
		return false
	}

	specificPermissions := k.GetUserPermissions(ctx, subspaceID, user)

	// The owner of the subspaces has all the permissions by default
	if subspace.Owner == user.String() {
		return true
	}

	// Get the group permissions
	groupPermissions := k.GetGroupsInheritedPermissions(ctx, subspaceID, user)

	// Check the combination of the permissions
	return types.CheckPermission(types.CombinePermissions(specificPermissions, groupPermissions), permission)
}

// GetUserPermissions returns the permissions that are currently set inside
// the subspace with the given id for the given user
func (k Keeper) GetUserPermissions(ctx sdk.Context, subspaceID uint64, user sdk.AccAddress) types.Permission {
	store := ctx.KVStore(k.storeKey)
	return types.UnmarshalPermission(store.Get(types.UserPermissionStoreKey(subspaceID, user)))
}

// GetGroupsInheritedPermissions returns the permissions that the specified user
// has inherited from all the groups that they are part of.
func (k Keeper) GetGroupsInheritedPermissions(ctx sdk.Context, subspaceID uint64, user sdk.AccAddress) types.Permission {
	var permissions []types.Permission
	k.IterateSubspaceGroups(ctx, subspaceID, func(index int64, group types.UserGroup) (stop bool) {
		if k.IsMemberOfGroup(ctx, subspaceID, group.ID, user) {
			permissions = append(permissions, group.Permissions)
		}
		return false
	})
	return types.CombinePermissions(permissions...)
}

// GetUsersWithPermission returns all the users that have a given permission inside the specified subspace
func (k Keeper) GetUsersWithPermission(ctx sdk.Context, subspaceID uint64, permission types.Permission) ([]sdk.AccAddress, error) {
	subspace, found := k.GetSubspace(ctx, subspaceID)
	if !found {
		return nil, nil
	}

	// The owner must always be included as they have all the permissions
	ownerAddr, err := sdk.AccAddressFromBech32(subspace.Owner)
	if err != nil {
		return nil, err
	}

	users := []sdk.AccAddress{ownerAddr}

	// Iterate over the various groups
	k.IterateSubspaceGroups(ctx, subspaceID, func(index int64, group types.UserGroup) (stop bool) {
		if !types.CheckPermission(group.Permissions, permission) {
			// Return early if the group does not have the permission. We will check other groups anyway
			return false
		}

		// If the group has the permission, get all the members
		k.IterateGroupMembers(ctx, subspaceID, group.ID, func(index int64, member sdk.AccAddress) (stop bool) {
			users = append(users, member)
			return false
		})

		return false
	})

	// Iterate over the various individually-set permissions
	k.IterateSubspacePermissions(ctx, subspaceID, func(index int64, user sdk.AccAddress, permission types.Permission) (stop bool) {
		if !types.CheckPermission(permission, permission) {
			// Return early if the user does not have the permission. We will check other users anyway
			return false
		}

		users = append(users, user)
		return false
	})

	return users, nil
}

// RemoveUserPermissions removes the permission for the given user inside the provided subspace
func (k Keeper) RemoveUserPermissions(ctx sdk.Context, subspaceID uint64, user sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.UserPermissionStoreKey(subspaceID, user))

	k.AfterUserPermissionRemoved(ctx, subspaceID, user)
}
