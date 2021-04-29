package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.BinaryMarshaler
}

// NewKeeper creates new instances of the subspaces keeper
func NewKeeper(storeKey sdk.StoreKey, cdc codec.BinaryMarshaler) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

// SaveSubspace saves the given subspace inside the current context.
// It assumes that the subspaces has been validated already.
func (k Keeper) SaveSubspace(ctx sdk.Context, subspace types.Subspace) error {
	store := ctx.KVStore(k.storeKey)
	key := types.SubspaceStoreKey(subspace.Id)

	// Check if the subspace already exists inside the store
	if store.Has(key) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %s already exists", subspace.Id)
	}

	store.Set(key, k.cdc.MustMarshalBinaryBare(&subspace))
	return nil
}

// DoesSubspaceExists returns true if the subspace with the given id exists inside the store.
func (k Keeper) DoesSubspaceExists(ctx sdk.Context, subspaceId string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.SubspaceStoreKey(subspaceId))
}

// GetSubspace returns the subspace associated with the given ID.
// If there is no subspace associated with the given ID the function will return an error.
func (k Keeper) GetSubspace(ctx sdk.Context, subspaceId string) (subspace types.Subspace, found bool) {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(types.SubspaceStoreKey(subspaceId)) {
		return types.Subspace{}, false
	}

	k.cdc.MustUnmarshalBinaryBare(store.Get(types.SubspaceStoreKey(subspaceId)), &subspace)
	return subspace, true
}

// AddAdminToSubspace insert the newAdmin inside the admins list of the given subspace if its not present.
// Returns an error if the admin is already present.
func (k Keeper) AddAdminToSubspace(ctx sdk.Context, subspaceId, user string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.AdminsStoreKey(subspaceId)

	admins := types.MustUnmarshalUsers(k.cdc, store.Get(key))
	if admins.IsPresent(user) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the user is already an admin of the subspace: %s", subspaceId)
	}

	admins.Users = append(admins.Users, user)
	store.Set(key, types.MustMarshalUsers(k.cdc, admins))
	return nil
}

// IsAdmin returns true if the given address is an admin inside the given subspace id, false otherwise.
func (k Keeper) IsAdmin(ctx sdk.Context, address, subspaceId string) bool {
	store := ctx.KVStore(k.storeKey)
	key := types.AdminsStoreKey(subspaceId)

	admins := types.MustUnmarshalUsers(k.cdc, store.Get(key))
	return admins.IsPresent(address)
}

// GetAllSubspaceAdmins returns a list of all the subspace admins
func (k Keeper) GetAllSubspaceAdmins(ctx sdk.Context, subspaceId string) types.Users {
	store := ctx.KVStore(k.storeKey)
	key := types.AdminsStoreKey(subspaceId)

	return types.MustUnmarshalUsers(k.cdc, store.Get(key))
}

// RemoveAdminFromSubspace remove the given admin from the given subspace.
// It returns error when the admin is not present inside the subspace.
func (k Keeper) RemoveAdminFromSubspace(ctx sdk.Context, subspaceId, admin string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.AdminsStoreKey(subspaceId)

	wrapped := types.MustUnmarshalUsers(k.cdc, store.Get(key))
	admins, found := types.RemoveUser(wrapped.Users, admin)

	// The admin doesn't exist, return error
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "this address %s is not an admin of the subspace %s",
			admin, subspaceId)
	}

	// Delete the key if no admins left inside the list.
	// This cleans up the store avoid export/import tests to fail due to a different keys number
	if len(admins) == 0 {
		store.Delete(key)
	} else {
		store.Set(key, types.MustMarshalUsers(k.cdc, types.Users{Users: admins}))
	}

	return nil
}

// AllowUserPosts give a user the possibility to post inside the given subspace.
// It returns error when the user can already post inside the subspace.
func (k Keeper) AllowUserPosts(ctx sdk.Context, newUser, subspaceId string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.AllowedToPostUsersKey(subspaceId)

	wrapped := types.MustUnmarshalUsers(k.cdc, store.Get(key))
	if wrapped.IsPresent(newUser) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the user is already allowed to post inside the subspace: %s", subspaceId)
	}

	wrapped.Users = append(wrapped.Users, newUser)
	store.Set(key, types.MustMarshalUsers(k.cdc, wrapped))
	return nil
}

func (k Keeper) BlockUserPosts(ctx sdk.Context, userToBlock, subspaceId string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.AllowedToPostUsersKey(subspaceId)

	wrapped := types.MustUnmarshalUsers(k.cdc, store.Get(key))
	users, found := types.RemoveUser(wrapped.Users, userToBlock)

	// The user doesn't exist, return error
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the user %s already can't post inside the subspace: %s", userToBlock, subspaceId)
	}

	// Delete the key if no admins left inside the list.
	// This cleans up the store avoid export/import tests to fail due to a different keys number
	if len(users) == 0 {
		store.Delete(key)
	} else {
		store.Set(key, types.MustMarshalUsers(k.cdc, types.Users{Users: users}))
	}

	return nil
}
