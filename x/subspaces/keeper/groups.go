package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

// HasUserGroup returns whether the given subspace has a group with the specified name or not
func (k Keeper) HasUserGroup(ctx sdk.Context, subspaceID uint64, groupName string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GroupStoreKey(subspaceID, groupName))
}

// SaveUserGroup saves within the subspace having the given id the group with the specified name and permissions
func (k Keeper) SaveUserGroup(ctx sdk.Context, subspaceID uint64, groupName string, permissions types.Permission) {
	store := ctx.KVStore(k.storeKey)

	// Save the group
	store.Set(types.GroupStoreKey(subspaceID, groupName), []byte{0x01})

	// Save the permissions
	k.SetPermissions(ctx, subspaceID, groupName, permissions)

	k.Logger(ctx).Info("group saved", "subspace_id", subspaceID, "group_name", groupName)
	k.AfterSubspaceGroupSaved(ctx, subspaceID, groupName)
}

// DeleteUserGroup deletes the group with the given name from the subspace with the provided id
func (k Keeper) DeleteUserGroup(ctx sdk.Context, subspaceID uint64, groupName string) {
	store := ctx.KVStore(k.storeKey)

	// Remove all the members from this group
	var members []sdk.AccAddress
	k.IterateGroupMembers(ctx, subspaceID, groupName, func(index int64, member sdk.AccAddress) (stop bool) {
		members = append(members, member)
		return false
	})

	for _, member := range members {
		k.RemoveUserFromGroup(ctx, subspaceID, groupName, member)
	}

	// Remove the group permissions
	k.RemovePermissions(ctx, subspaceID, groupName)

	// Delete the group
	store.Delete(types.GroupStoreKey(subspaceID, groupName))

	k.Logger(ctx).Info("group deleted", "subspace_id", subspaceID, "group_name", groupName)
	k.AfterSubspaceGroupDeleted(ctx, subspaceID, groupName)
}

// IsMemberOfGroup returns whether the given user is part of the group with
// the specified name inside the provided subspace
func (k Keeper) IsMemberOfGroup(ctx sdk.Context, subspaceID uint64, groupName string, user sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GroupMemberStoreKey(subspaceID, groupName, user))
}

// AddUserToGroup adds the given user to the group having the provided name inside the specified subspace.
// If the group does not exist inside the subspace, it returns an error.
func (k Keeper) AddUserToGroup(ctx sdk.Context, subspaceID uint64, groupName string, user sdk.AccAddress) error {
	if !k.HasUserGroup(ctx, subspaceID, groupName) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidVersion, "group with name %s does not exist", groupName)
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(types.GroupMemberStoreKey(subspaceID, groupName, user), []byte{0x01})

	k.AfterSubspaceGroupMemberAdded(ctx, subspaceID, groupName, user)

	return nil
}

// RemoveUserFromGroup removes the specified user from the subspace group having the given name.
// If the group does not exist inside the subspace, it returns an error.
func (k Keeper) RemoveUserFromGroup(ctx sdk.Context, subspaceID uint64, groupName string, user sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GroupMemberStoreKey(subspaceID, groupName, user))

	k.AfterSubspaceGroupMemberRemoved(ctx, subspaceID, groupName, user)
}
