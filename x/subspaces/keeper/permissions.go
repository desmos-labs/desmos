package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

// GetPermissions returns the permissions that are currently set inside
// the subspace with the given id for the given target
func (k Keeper) GetPermissions(ctx sdk.Context, subspaceID uint64, target string) types.Permission {
	store := ctx.KVStore(k.storeKey)
	return types.UnmarshalPermission(store.Get(types.PermissionStoreKey(subspaceID, target)))
}

// GetGroupsInheritedPermissions returns the permissions that the specified target
// has inherited from all the groups that they are part of.
func (k Keeper) GetGroupsInheritedPermissions(ctx sdk.Context, subspaceID uint64, user sdk.AccAddress) types.Permission {
	var permissions []types.Permission
	k.IterateSubspaceGroups(ctx, subspaceID, func(index int64, groupName string) (stop bool) {
		if k.IsMemberOfGroup(ctx, subspaceID, groupName, user) {
			permissions = append(permissions, k.GetPermissions(ctx, subspaceID, groupName))
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
	k.IterateSubspaceGroups(ctx, subspaceID, func(index int64, groupName string) (stop bool) {
		if !k.HasPermission(ctx, subspaceID, groupName, permission) {
			// Return early if the group does not have the permission. We will check other groups anyway
			return false
		}

		// If the group has the permission, get all the members
		k.IterateGroupMembers(ctx, subspaceID, groupName, func(index int64, member sdk.AccAddress) (stop bool) {
			users = append(users, member)
			return false
		})

		return false
	})

	return users, nil
}

// HasPermission checks whether the specific target has the given permission inside a specific subspace
func (k Keeper) HasPermission(ctx sdk.Context, subspaceID uint64, target string, permission types.Permission) bool {
	// Get the subspace to make sure the request is valid
	subspace, found := k.GetSubspace(ctx, subspaceID)
	if !found {
		return false
	}

	specificPermissions := k.GetPermissions(ctx, subspaceID, target)

	userAddr, err := sdk.AccAddressFromBech32(target)
	if err != nil {
		return types.CheckPermission(specificPermissions, permission)
	}

	// The owner of the subspaces has all the permissions by default
	if subspace.Owner == userAddr.String() {
		return true
	}

	// Get the group permissions
	groupPermissions := k.GetGroupsInheritedPermissions(ctx, subspaceID, userAddr)

	// Check the combination of the permissions
	return types.CheckPermission(types.CombinePermissions(specificPermissions, groupPermissions), permission)
}

// SetPermissions sets the given permission for the specific target inside a single subspace
func (k Keeper) SetPermissions(ctx sdk.Context, subspaceID uint64, target string, permissions uint32) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PermissionStoreKey(subspaceID, target), types.MarshalPermission(permissions))
}

// RemovePermissions removes the permission for the given target inside the provided subspace
func (k Keeper) RemovePermissions(ctx sdk.Context, subspaceID uint64, target string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.PermissionStoreKey(subspaceID, target))
}
