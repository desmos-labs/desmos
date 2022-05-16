package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// SetNextGroupID sets the new group id for the specific subspace to the store
func (k Keeper) SetNextGroupID(ctx sdk.Context, subspaceID uint64, groupID uint32) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.NextGroupIDStoreKey(subspaceID), types.GetGroupIDBytes(groupID))
}

// HasNextGroupID tells whether the next group id key exists for the given subspace
func (k Keeper) HasNextGroupID(ctx sdk.Context, subspaceID uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.NextGroupIDStoreKey(subspaceID))
}

// GetNextGroupID gets the highest group id for the subspace with the given id
func (k Keeper) GetNextGroupID(ctx sdk.Context, subspaceID uint64) (groupID uint32, err error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NextGroupIDStoreKey(subspaceID))
	if bz == nil {
		return 0, sdkerrors.Wrap(types.ErrInvalidGenesis, "initial group ID hasn't been set")
	}

	groupID = types.GetGroupIDFromBytes(bz)
	return groupID, nil
}

// DeleteNextGroupID deletes the next group id key for the given subspace
func (k Keeper) DeleteNextGroupID(ctx sdk.Context, subspaceID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.NextGroupIDStoreKey(subspaceID))
}

// --------------------------------------------------------------------------------------------------------------------

// SaveUserGroup saves the given user group
func (k Keeper) SaveUserGroup(ctx sdk.Context, group types.UserGroup) {
	store := ctx.KVStore(k.storeKey)

	// Save the group
	store.Set(types.GroupStoreKey(group.SubspaceID, group.SectionID, group.ID), k.cdc.MustMarshal(&group))

	k.Logger(ctx).Info("group saved", "subspace id", group.SubspaceID, "group id", group.ID)
	k.AfterSubspaceGroupSaved(ctx, group.SubspaceID, group.ID)
}

// HasUserGroup tells whether the given subspace has a group with the specified id or not
func (k Keeper) HasUserGroup(ctx sdk.Context, subspaceID uint64, groupID uint32) bool {
	found := false
	k.IterateSubspaceUserGroups(ctx, subspaceID, func(index int64, group types.UserGroup) (stop bool) {
		if group.ID == groupID {
			found = true
		}
		return found
	})
	return found
}

// GetUserGroup returns the group associated with the given id inside the subspace with the provided id.
// If there is no group associated with the given id the function will return an empty group and false.
func (k Keeper) GetUserGroup(ctx sdk.Context, subspaceID uint64, groupID uint32) (group types.UserGroup, found bool) {
	k.IterateSubspaceUserGroups(ctx, subspaceID, func(index int64, g types.UserGroup) (stop bool) {
		if g.ID == groupID {
			group = g
			found = true
		}
		return found
	})
	return group, found
}

// DeleteUserGroup deletes the group with the given id from the subspace with the provided id
func (k Keeper) DeleteUserGroup(ctx sdk.Context, subspaceID uint64, groupID uint32) {
	group, found := k.GetUserGroup(ctx, subspaceID, groupID)
	if !found {
		return
	}

	// Remove all the members from this group
	k.IterateUserGroupMembers(ctx, subspaceID, groupID, func(index int64, member sdk.AccAddress) (stop bool) {
		k.RemoveUserFromGroup(ctx, subspaceID, groupID, member)
		return false
	})

	// Delete the group
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GroupStoreKey(subspaceID, group.SectionID, group.ID))

	k.Logger(ctx).Info("group deleted", "subspace id", subspaceID, "group id", groupID)
	k.AfterSubspaceGroupDeleted(ctx, subspaceID, groupID)
}

// --------------------------------------------------------------------------------------------------------------------

// AddUserToGroup adds the given user to the group having the provided id inside the specified subspace.
// If the group does not exist inside the subspace, it returns an error.
func (k Keeper) AddUserToGroup(ctx sdk.Context, subspaceID uint64, groupID uint32, user sdk.AccAddress) error {
	if !k.HasUserGroup(ctx, subspaceID, groupID) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "group with id %d does not exist", groupID)
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(types.GroupMemberStoreKey(subspaceID, groupID, user), []byte{0x01})

	k.AfterSubspaceGroupMemberAdded(ctx, subspaceID, groupID, user)

	return nil
}

// IsMemberOfGroup returns whether the given user is part of the group with
// the specified id inside the provided subspace
func (k Keeper) IsMemberOfGroup(ctx sdk.Context, subspaceID uint64, groupID uint32, user sdk.AccAddress) bool {
	// The group with ID 0 represents the default group, so everyone is part of it
	if groupID == 0 {
		return true
	}

	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GroupMemberStoreKey(subspaceID, groupID, user))
}

// RemoveUserFromGroup removes the specified user from the subspace group having the given id.
func (k Keeper) RemoveUserFromGroup(ctx sdk.Context, subspaceID uint64, groupID uint32, user sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GroupMemberStoreKey(subspaceID, groupID, user))

	k.AfterSubspaceGroupMemberRemoved(ctx, subspaceID, groupID, user)
}
