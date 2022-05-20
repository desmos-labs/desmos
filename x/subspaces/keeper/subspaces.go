package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// SetSubspaceID sets the new subspace id to the store
func (k Keeper) SetSubspaceID(ctx sdk.Context, subspaceID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.SubspaceIDKey, types.GetSubspaceIDBytes(subspaceID))
}

// GetSubspaceID gets the highest subspace id
func (k Keeper) GetSubspaceID(ctx sdk.Context) (subspaceID uint64, err error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.SubspaceIDKey)
	if bz == nil {
		return 0, sdkerrors.Wrap(types.ErrInvalidGenesis, "initial subspace ID hasn't been set")
	}

	subspaceID = types.GetSubspaceIDFromBytes(bz)
	return subspaceID, nil
}

// --------------------------------------------------------------------------------------------------------------------

// SaveSubspace saves the given subspace inside the current context.
func (k Keeper) SaveSubspace(ctx sdk.Context, subspace types.Subspace) {
	store := ctx.KVStore(k.storeKey)

	// Store the subspace
	store.Set(types.SubspaceKey(subspace.ID), k.cdc.MustMarshal(&subspace))

	// If the initial group id does not exist, create it now
	groupIDKey := types.GroupIDStoreKey(subspace.ID)
	if !store.Has(groupIDKey) {
		store.Set(groupIDKey, types.GetGroupIDBytes(1))
	}

	// If the subspace does not have the default group, create it now
	if !k.HasUserGroup(ctx, subspace.ID, 0) {
		k.SaveUserGroup(ctx, types.DefaultUserGroup(subspace.ID))
	}

	k.Logger(ctx).Info("subspace saved", "id", subspace.ID)
	k.AfterSubspaceSaved(ctx, subspace.ID)
}

// HasSubspace tells whether the given subspace exists or not
func (k Keeper) HasSubspace(ctx sdk.Context, subspaceID uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.SubspaceKey(subspaceID))
}

// GetSubspace returns the subspace associated with the given id.
// If there is no subspace associated with the given id the function will return an empty subspace and false.
func (k Keeper) GetSubspace(ctx sdk.Context, subspaceID uint64) (subspace types.Subspace, found bool) {
	store := ctx.KVStore(k.storeKey)

	key := types.SubspaceKey(subspaceID)
	if !store.Has(key) {
		return types.Subspace{}, false
	}

	k.cdc.MustUnmarshal(store.Get(key), &subspace)
	return subspace, true
}

// DeleteSubspace allows to delete the subspace with the given id
func (k Keeper) DeleteSubspace(ctx sdk.Context, subspaceID uint64) {
	// Delete the subspace
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.SubspaceKey(subspaceID))

	// Delete the group id
	store.Delete(types.GroupIDStoreKey(subspaceID))

	// Delete all user groups
	for _, group := range k.GetSubspaceGroups(ctx, subspaceID) {
		k.DeleteUserGroup(ctx, subspaceID, group.ID)
	}

	// Delete all the permissions for this subspace
	var usersWithPermissions []sdk.AccAddress
	k.IterateSubspacePermissions(ctx, subspaceID, func(_ int64, entry types.UserPermission) (stop bool) {
		user, err := sdk.AccAddressFromBech32(entry.User)
		if err != nil {
			panic(err)
		}

		usersWithPermissions = append(usersWithPermissions, user)
		return false
	})

	for _, member := range usersWithPermissions {
		k.RemoveUserPermissions(ctx, subspaceID, member)
	}

	// Log the subspace deletion
	k.Logger(ctx).Info("subspace deleted", "id", subspaceID)
	k.AfterSubspaceDeleted(ctx, subspaceID)
}
