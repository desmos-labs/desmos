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
func (k Keeper) GetGroupsInheritedPermissions(ctx sdk.Context, subspaceID uint64, target string) types.Permission {
	// TODO
	return 0
}

// HasPermission checks whether the specific target has the given permission inside a specific subspace
func (k Keeper) HasPermission(ctx sdk.Context, subspaceID uint64, target string, permission types.Permission) (bool, error) {
	store := ctx.KVStore(k.storeKey)

	// TODO: Check if the target is the current owner
	subspace, found := k.GetSubspace(ctx, subspaceID)
	if !found {
		return false, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d does not exist", subspaceID)
	}

	var defaultPermission = types.PermissionNothing

	// If the target is the owner of the subspace, they should have all the permission
	if subspace.Owner == target {
		defaultPermission = types.PermissionEverything
	}

	// TODO: Check if the target is a user and is part of a group
	return (k.GetPermissions(ctx, subspaceID, target) & permission) == permission, nil
}

// SetPermissions sets the given permission for the specific target inside a single subspace
func (k Keeper) SetPermissions(ctx sdk.Context, subspaceID uint64, target string, permissions uint32) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PermissionStoreKey(subspaceID, target), types.MarshalPermission(permissions))
}
