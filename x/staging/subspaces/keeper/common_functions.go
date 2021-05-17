package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
)

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

// CheckSubspaceExistenceAndAdminValidity checks if the subspaces with the given id exists and
// if the address belongs to one of its admins
func (k Keeper) CheckSubspaceExistenceAndAdminValidity(ctx sdk.Context, address, subspaceID string) error {
	if !k.DoesSubspaceExists(ctx, subspaceID) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the subspaces with id %s doesn't exist", subspaceID,
		)
	}

	if !k.IsAdmin(ctx, address, subspaceID) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the user: %s is not an admin and can't perform this operation on the subspaces: %s",
			address, subspaceID)
	}

	return nil
}

// CheckSubspaceExistenceAndOwnerValidity check if the subspaces with the given id exists and
// if the address belongs to its creator
func (k Keeper) CheckSubspaceExistenceAndOwnerValidity(ctx sdk.Context, address, subspaceID string) error {
	subspace, exist := k.GetSubspace(ctx, subspaceID)
	if !exist {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the subspaces with id %s doesn't exist", subspaceID,
		)
	}

	if subspace.Owner != address {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the user: %s is not the subspaces owner and can't perform this operation on the subspaces: %s",
			address, subspaceID,
		)
	}

	return nil
}

// AddUserToList insert the given user inside a users list of a specific susbspace identified by the given subspaceId;
// this list is stored under the given storeKey.
// It returns error when the user is already present in that list.
func (k Keeper) AddUserToList(ctx sdk.Context, storeKey []byte, subspaceID, user, error string) error {
	store := ctx.KVStore(k.storeKey)

	wrapped := types.MustUnmarshalUsers(k.cdc, store.Get(storeKey))
	if wrapped.IsPresent(user) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, error, user, subspaceID)
	}

	wrapped.Users = append(wrapped.Users, user)
	store.Set(storeKey, types.MustMarshalUsers(k.cdc, wrapped))
	return nil
}

// RemoveUserFromList remove the given user from a users list of a specific subspaces identified by the given subspaceId;
// this list is stored under the given storeKey.
// It returns error when the user is not present in that list.
func (k Keeper) RemoveUserFromList(ctx sdk.Context, storeKey []byte, subspaceID, user, error string) error {
	store := ctx.KVStore(k.storeKey)

	wrapped := types.MustUnmarshalUsers(k.cdc, store.Get(storeKey))
	users, found := types.RemoveUser(wrapped.Users, user)

	// The user isn't present inside the list, return error
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, error, user, subspaceID)
	}

	// Delete the key if no users left inside the list.
	// This cleans up the store avoid export/import tests to fail due to a different keys number
	if len(users) == 0 {
		store.Delete(storeKey)
	} else {
		store.Set(storeKey, types.MustMarshalUsers(k.cdc, types.Users{Users: users}))
	}

	return nil
}
