package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

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

// HasPermission checks whether the specific target has the given permission inside a specific subspace
func (k Keeper) HasPermission(ctx sdk.Context, subspaceID uint64, target string, permission types.Permission) (bool, error) {
	// Get the subspace to make sure the request is valid
	subspace, found := k.GetSubspace(ctx, subspaceID)
	if !found {
		return false, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d does not exist", subspaceID)
	}

	specificPermissions := k.GetPermissions(ctx, subspaceID, target)

	userAddr, err := sdk.AccAddressFromBech32(target)
	if err != nil {
		return types.CheckPermission(specificPermissions, permission), nil
	}

	// The owner of the subspaces has all the permissions by default
	if subspace.Owner == userAddr.String() {
		return true, nil
	}

	// Get the group permissions
	groupPermissions := k.GetGroupsInheritedPermissions(ctx, subspaceID, userAddr)

	// Check the combination of the permissions
	return types.CheckPermission(types.CombinePermissions(specificPermissions, groupPermissions), permission), nil
}

// SetPermissions sets the given permission for the specific target inside a single subspace
func (k Keeper) SetPermissions(ctx sdk.Context, subspaceID uint64, target string, permissions uint32) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PermissionStoreKey(subspaceID, target), types.MarshalPermission(permissions))
}
