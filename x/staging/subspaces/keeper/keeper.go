package keeper

import (
	"bytes"
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
	key := types.SubspaceStoreKey(subspace.ID)

	// Check if the subspace already exists inside the store
	if store.Has(key) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %s already exists", subspace.ID)
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

// GetAllSubspaces returns a list of all the subspaces that have been store inside the given context
func (k Keeper) GetAllSubspaces(ctx sdk.Context) []types.Subspace {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SubspaceStorePrefix)
	defer iterator.Close()

	var subspaces []types.Subspace
	for ; iterator.Valid(); iterator.Next() {
		var subspace types.Subspace
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &subspace)
		subspaces = append(subspaces, subspace)
	}

	return subspaces
}

// addUserToList insert the given user inside a users list of a specific susbspace identified by the given subspaceId;
// this list is stored under the given storeKey.
// It returns error when the user is already present in that list.
func (k Keeper) addUserToList(ctx sdk.Context, storeKey []byte, subspaceId, user, error string) error {
	store := ctx.KVStore(k.storeKey)

	wrapped := types.MustUnmarshalUsers(k.cdc, store.Get(storeKey))
	if wrapped.IsPresent(user) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, error, user, subspaceId)
	}

	wrapped.Users = append(wrapped.Users, user)
	store.Set(storeKey, types.MustMarshalUsers(k.cdc, wrapped))
	return nil
}

// removeUserFromList remove the given user from a users list of a specific subspace identified by the given subspaceId;
// this list is stored under the given storeKey.
// It returns error when the user is not present in that list.
func (k Keeper) removeUserFromList(ctx sdk.Context, storeKey []byte, subspaceId, user, error string) error {
	store := ctx.KVStore(k.storeKey)

	wrapped := types.MustUnmarshalUsers(k.cdc, store.Get(storeKey))
	users, found := types.RemoveUser(wrapped.Users, user)

	// The user isn't present inside the list, return error
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, error, user, subspaceId)
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

// AddAdminToSubspace insert the newAdmin inside the admins list of the given subspace if its not present.
// Returns an error if the admin is already present.
func (k Keeper) AddAdminToSubspace(ctx sdk.Context, subspaceId, user string) error {
	if err := k.addUserToList(ctx, types.AdminsStoreKey(subspaceId), subspaceId, user,
		"the user: %s is already an admin of the subspace: %s"); err != nil {
		return err
	}
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
	// If the admin doesn't exist, return error
	if err := k.removeUserFromList(ctx, types.AdminsStoreKey(subspaceId), subspaceId, admin,
		"this address: %s is not an admin of the subspace %s"); err != nil {
		return err
	}
	return nil
}

// UnblockPostsForUser give a user the possibility to post inside the given subspace.
// It returns error when the user can already post inside the subspace.
func (k Keeper) UnblockPostsForUser(ctx sdk.Context, user, subspaceId string) error {
	if err := k.removeUserFromList(ctx, types.BlockedToPostUsersKey(subspaceId), subspaceId, user,
		"the user: %s is already allowed to post inside the subspace: %s"); err != nil {
		return err
	}
	return nil
}

// BlockPostsForUser block the given user to post anything inside the given subspace.
// It returns error if the user already can't post inside the subspace.
func (k Keeper) BlockPostsForUser(ctx sdk.Context, userToBlock, subspaceId string) error {
	if err := k.addUserToList(ctx, types.BlockedToPostUsersKey(subspaceId), subspaceId, userToBlock,
		"the user: %s already can't post inside the subspace: %s"); err != nil {
		return err
	}
	return nil
}

// GetSubspaceBlockedUsers returns a list of all the blocked users unable to post inside the given subspace
func (k Keeper) GetSubspaceBlockedUsers(ctx sdk.Context, subspaceId string) types.Users {
	store := ctx.KVStore(k.storeKey)
	key := types.BlockedToPostUsersKey(subspaceId)

	return types.MustUnmarshalUsers(k.cdc, store.Get(key))
}

// GetSubspaceAdminsEntry returns a list of all the subspaces associated with their admins
func (k Keeper) GetSubspaceAdminsEntry(ctx sdk.Context) []types.SubspaceAdminsEntry {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.AdminsStorePrefix)
	defer iterator.Close()

	var entries []types.SubspaceAdminsEntry
	for ; iterator.Valid(); iterator.Next() {
		admins := types.MustUnmarshalUsers(k.cdc, iterator.Value())
		idBytes := bytes.TrimPrefix(iterator.Key(), types.AdminsStorePrefix)
		subspaceId := string(idBytes)
		entries = append(entries, types.NewAdminsEntries(subspaceId, admins))
	}

	return entries
}

// GetBlockedToPostUsers returns a list of all the subspaces associated with the users not allowed to post inside of them
func (k Keeper) GetBlockedToPostUsers(ctx sdk.Context) []types.BlockedUsersEntry {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.AdminsStorePrefix)
	defer iterator.Close()

	var entries []types.BlockedUsersEntry
	for ; iterator.Valid(); iterator.Next() {
		users := types.MustUnmarshalUsers(k.cdc, iterator.Value())
		idBytes := bytes.TrimPrefix(iterator.Key(), types.BlockedUsersPostsPrefix)
		subspaceId := string(idBytes)
		entries = append(entries, types.NewBlockedUsersEntry(subspaceId, users))
	}

	return entries
}
