package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	types2 "github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

// IterateSubspaces iterates through the subspaces set and performs the given function
func (k Keeper) IterateSubspaces(ctx sdk.Context, fn func(index int64, subspace types2.Subspace) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types2.SubspaceStorePrefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var subspace types2.Subspace
		k.cdc.MustUnmarshal(iterator.Value(), &subspace)
		stop := fn(i, subspace)
		if stop {
			break
		}
		i++
	}
}

// GetAllSubspaces returns a list of all the subspaces that have been store inside the given context
func (k Keeper) GetAllSubspaces(ctx sdk.Context) []types2.Subspace {
	var subspaces []types2.Subspace
	k.IterateSubspaces(ctx, func(_ int64, subspace types2.Subspace) (stop bool) {
		subspaces = append(subspaces, subspace)
		return false
	})

	return subspaces
}

// IterateSubspaceAdmins iterates over all the admins of the subspace having the given id
func (k Keeper) IterateSubspaceAdmins(ctx sdk.Context, id string, fn func(index int64, admin string) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types2.SubspaceAdminsPrefix(id))
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		stop := fn(i, string(iterator.Value()))
		if stop {
			break
		}
		i++
	}
}

// GetAllAdmins returns the entries containing the data of all the admins of all the subspaces
func (k Keeper) GetAllAdmins(ctx sdk.Context) []types2.UsersEntry {
	var entries []types2.UsersEntry
	k.IterateSubspaces(ctx, func(_ int64, subspace types2.Subspace) (stop bool) {

		var admins []string
		k.IterateSubspaceAdmins(ctx, subspace.ID, func(_ int64, admin string) (stop bool) {
			admins = append(admins, admin)
			return false
		})

		entries = append(entries, types2.NewUsersEntry(subspace.ID, admins))
		return false
	})

	return entries
}

// IterateSubspaceRegisteredUsers iterates over all the registered users of the subspace having the given id
func (k Keeper) IterateSubspaceRegisteredUsers(ctx sdk.Context, id string, fn func(index int64, user string) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types2.SubspaceRegisteredUsersPrefix(id))
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		stop := fn(i, string(iterator.Value()))
		if stop {
			break
		}
		i++
	}
}

// GetAllRegisteredUsers returns the entries containing the data of all the registered users of all the subspaces
func (k Keeper) GetAllRegisteredUsers(ctx sdk.Context) []types2.UsersEntry {
	var entries []types2.UsersEntry
	k.IterateSubspaces(ctx, func(_ int64, subspace types2.Subspace) (stop bool) {

		var users []string
		k.IterateSubspaceRegisteredUsers(ctx, subspace.ID, func(_ int64, user string) (stop bool) {
			users = append(users, user)
			return false
		})

		entries = append(entries, types2.NewUsersEntry(subspace.ID, users))
		return false
	})

	return entries
}

// IterateSubspaceBannedUsers iterates over all the banned users of the subspace having the given id
func (k Keeper) IterateSubspaceBannedUsers(ctx sdk.Context, id string, fn func(index int64, user string) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types2.SubspaceBannedUsersPrefix(id))
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		stop := fn(i, string(iterator.Value()))
		if stop {
			break
		}
		i++
	}
}

// GetAllBannedUsers returns the entries containing the data of all the banned users of all the subspaces
func (k Keeper) GetAllBannedUsers(ctx sdk.Context) []types2.UsersEntry {
	var entries []types2.UsersEntry
	k.IterateSubspaces(ctx, func(_ int64, subspace types2.Subspace) (stop bool) {

		var users []string
		k.IterateSubspaceBannedUsers(ctx, subspace.ID, func(_ int64, user string) (stop bool) {
			users = append(users, user)
			return false
		})

		entries = append(entries, types2.NewUsersEntry(subspace.ID, users))
		return false
	})

	return entries
}

// checkSubspaceAdmin checks if the subspace with the given id exists and
// if the address belongs to the owner of the subspace or one of its admins.
func (k Keeper) checkSubspaceAdmin(ctx sdk.Context, id, address string) error {
	subspace, found := k.GetSubspace(ctx, id)
	if !found {
		return sdkerrors.Wrapf(types2.ErrInvalidSubspaceID, "the subspace with id %s doesn't exist", id)
	}

	if subspace.Owner != address {
		store := ctx.KVStore(k.storeKey)
		if !store.Has(types2.SubspaceAdminKey(subspace.ID, address)) {
			return sdkerrors.Wrapf(types2.ErrInvalidSubspaceAdmin, address)
		}
	}

	return nil
}

// checkSubspaceOwner checks if the subspace with the given id exists and
// if the address belongs to the owner of the subspace.
func (k Keeper) checkSubspaceOwner(ctx sdk.Context, id, address string) error {
	subspace, found := k.GetSubspace(ctx, id)
	if !found {
		return sdkerrors.Wrapf(types2.ErrInvalidSubspaceID, "the subspace with id %s doesn't exist", id)
	}

	if subspace.Owner != address {
		return sdkerrors.Wrapf(types2.ErrInvalidSubspaceOwner, address)
	}

	return nil
}
