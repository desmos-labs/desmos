package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

// SetGroupID sets the new group id for the specific subspace to the store
func (k Keeper) SetGroupID(ctx sdk.Context, subspaceID uint64, groupID uint32) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GroupIDStoreKey(subspaceID), types.GetGroupIDBytes(groupID))
}

// GetGroupID gets the highest group id for the subspace with the given id
func (k Keeper) GetGroupID(ctx sdk.Context, subspaceID uint64) (groupID uint32, err error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GroupIDStoreKey(subspaceID))
	if bz == nil {
		return 0, sdkerrors.Wrap(types.ErrInvalidGenesis, "initial group ID hasn't been set")
	}

	groupID = types.GetGroupIDFromBytes(bz)
	return groupID, nil
}

// --------------------------------------------------------------------------------------------------------------------

// SaveUserGroup saves within the subspace having the given id the provided group
func (k Keeper) SaveUserGroup(ctx sdk.Context, group types.UserGroup) {
	store := ctx.KVStore(k.storeKey)

	// Save the group
	store.Set(types.GroupStoreKey(group.SubspaceID, group.ID), k.cdc.MustMarshal(&group))

	k.Logger(ctx).Info("group saved", "subspace_id", group.SubspaceID, "group_id", group.ID)
	k.AfterSubspaceGroupSaved(ctx, group.SubspaceID, group.ID)
}

// HasUserGroup returns whether the given subspace has a group with the specified id or not
func (k Keeper) HasUserGroup(ctx sdk.Context, subspaceID uint64, groupID uint32) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GroupStoreKey(subspaceID, groupID))
}

// GetUserGroup returns the group associated with the given id inside the subspace with the provided id.
// If there is no group associated with the given id the function will return an empty group and false.
func (k Keeper) GetUserGroup(ctx sdk.Context, subspaceID uint64, groupID uint32) (group types.UserGroup, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GroupStoreKey(subspaceID, groupID)
	if !store.Has(key) {
		return group, false
	}

	k.cdc.MustUnmarshal(store.Get(key), &group)
	return group, true
}

// DeleteUserGroup deletes the group with the given id from the subspace with the provided id
func (k Keeper) DeleteUserGroup(ctx sdk.Context, subspaceID uint64, groupID uint32) {
	store := ctx.KVStore(k.storeKey)

	// Remove all the members from this group
	for _, member := range k.GetGroupMembers(ctx, subspaceID, groupID) {
		k.RemoveUserFromGroup(ctx, subspaceID, groupID, member)
	}

	// Delete the group
	store.Delete(types.GroupStoreKey(subspaceID, groupID))

	k.Logger(ctx).Info("group deleted", "subspace_id", subspaceID, "group_id", groupID)
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
