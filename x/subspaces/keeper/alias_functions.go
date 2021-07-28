package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/subspaces/types"
)

// IterateSubspaces iterates through the subspaces set and performs the given function
func (k Keeper) IterateSubspaces(ctx sdk.Context, fn func(index int64, subspace types.Subspace) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SubspaceStorePrefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var subspace types.Subspace
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &subspace)
		stop := fn(i, subspace)
		if stop {
			break
		}
		i++
	}
}

// GetAllSubspaces returns a list of all the subspaces that have been store inside the given context
func (k Keeper) GetAllSubspaces(ctx sdk.Context) []types.Subspace {
	var subspaces []types.Subspace
	k.IterateSubspaces(ctx, func(_ int64, subspace types.Subspace) (stop bool) {
		subspaces = append(subspaces, subspace)
		return false
	})

	return subspaces
}

// IterateSubspaceAdmins iterates over all the admins of the subspace having the given id
func (k Keeper) IterateSubspaceAdmins(ctx sdk.Context, id string, fn func(index int64, admin string) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SubspaceAdminsPrefix(id))
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
func (k Keeper) GetAllAdmins(ctx sdk.Context) []types.UsersEntry {
	var entries []types.UsersEntry
	k.IterateSubspaces(ctx, func(_ int64, subspace types.Subspace) (stop bool) {

		var admins []string
		k.IterateSubspaceAdmins(ctx, subspace.ID, func(_ int64, admin string) (stop bool) {
			admins = append(admins, admin)
			return false
		})

		entries = append(entries, types.NewUsersEntry(subspace.ID, admins))
		return false
	})

	return entries
}

// IterateSubspaceRegisteredUsers iterates over all the registered users of the subspace having the given id
func (k Keeper) IterateSubspaceRegisteredUsers(ctx sdk.Context, id string, fn func(index int64, user string) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SubspaceRegisteredUsersPrefix(id))
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
func (k Keeper) GetAllRegisteredUsers(ctx sdk.Context) []types.UsersEntry {
	var entries []types.UsersEntry
	k.IterateSubspaces(ctx, func(_ int64, subspace types.Subspace) (stop bool) {

		var users []string
		k.IterateSubspaceRegisteredUsers(ctx, subspace.ID, func(_ int64, user string) (stop bool) {
			users = append(users, user)
			return false
		})

		entries = append(entries, types.NewUsersEntry(subspace.ID, users))
		return false
	})

	return entries
}

// IterateSubspaceBannedUsers iterates over all the banned users of the subspace having the given id
func (k Keeper) IterateSubspaceBannedUsers(ctx sdk.Context, id string, fn func(index int64, user string) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SubspaceBannedUsersPrefix(id))
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
func (k Keeper) GetAllBannedUsers(ctx sdk.Context) []types.UsersEntry {
	var entries []types.UsersEntry
	k.IterateSubspaces(ctx, func(_ int64, subspace types.Subspace) (stop bool) {

		var users []string
		k.IterateSubspaceBannedUsers(ctx, subspace.ID, func(_ int64, user string) (stop bool) {
			users = append(users, user)
			return false
		})

		entries = append(entries, types.NewUsersEntry(subspace.ID, users))
		return false
	})

	return entries
}

// checkSubspaceAdmin checks if the subspace with the given id exists and
// if the address belongs to the owner of the subspace or one of its admins.
func (k Keeper) checkSubspaceAdmin(ctx sdk.Context, id, address string) error {
	subspace, found := k.GetSubspace(ctx, id)
	if !found {
		return sdkerrors.Wrapf(types.ErrInvalidSubspaceID, "the subspace with id %s doesn't exist", id)
	}

	if subspace.Owner != address {
		store := ctx.KVStore(k.storeKey)
		if !store.Has(types.SubspaceAdminKey(subspace.ID, address)) {
			return sdkerrors.Wrapf(types.ErrInvalidSubspaceAdmin, address)
		}
	}

	return nil
}

// checkSubspaceOwner checks if the subspace with the given id exists and
// if the address belongs to the owner of the subspace.
func (k Keeper) checkSubspaceOwner(ctx sdk.Context, id, address string) error {
	subspace, found := k.GetSubspace(ctx, id)
	if !found {
		return sdkerrors.Wrapf(types.ErrInvalidSubspaceID, "the subspace with id %s doesn't exist", id)
	}

	if subspace.Owner != address {
		return sdkerrors.Wrapf(types.ErrInvalidSubspaceOwner, address)
	}

	return nil
}

// IterateTokenomicsPair iterates through the tokenomics pairs set and performs the given function
func (k Keeper) IterateTokenomicsPair(ctx sdk.Context, fn func(index int64, tokenomicsPair types.Tokenomics) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.TokenomicsPairPrefix)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var pair types.Tokenomics
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &pair)
		stop := fn(i, pair)
		if stop {
			break
		}
		i++
	}
}

// GetAllTokenomicsPairs returns a list of all the tokenomicsPairs saved inside the current contex
func (k Keeper) GetAllTokenomicsPairs(ctx sdk.Context) []types.Tokenomics {
	var pairs []types.Tokenomics
	k.IterateTokenomicsPair(ctx, func(_ int64, pair types.Tokenomics) (stop bool) {
		pairs = append(pairs, pair)
		return false
	})

	return pairs
}
