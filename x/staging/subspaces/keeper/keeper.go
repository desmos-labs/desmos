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

// SaveSubspace saves the given subspaces inside the current context.
// It assumes that the subspaces has been validated already.
func (k Keeper) SaveSubspace(ctx sdk.Context, subspace types.Subspace) error {
	store := ctx.KVStore(k.storeKey)
	key := types.SubspaceStoreKey(subspace.ID)

	// Check if the subspaces already exists inside the store
	if store.Has(key) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "the subspaces with id %s already exists", subspace.ID)
	}

	store.Set(key, k.cdc.MustMarshalBinaryBare(&subspace))
	return nil
}

// DoesSubspaceExists returns true if the subspaces with the given id exists inside the store.
func (k Keeper) DoesSubspaceExists(ctx sdk.Context, subspaceID string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.SubspaceStoreKey(subspaceID))
}

// GetSubspace returns the subspaces associated with the given ID.
// If there is no subspaces associated with the given ID the function will return an error.
func (k Keeper) GetSubspace(ctx sdk.Context, subspaceID string) (subspace types.Subspace, found bool) {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(types.SubspaceStoreKey(subspaceID)) {
		return types.Subspace{}, false
	}

	k.cdc.MustUnmarshalBinaryBare(store.Get(types.SubspaceStoreKey(subspaceID)), &subspace)
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

// TransferSubspaceOwnership transfer the ownership of the subspaces with the given subspaceID to the newOwner.
// It returns error if the subspaces doesnt exist.
func (k Keeper) TransferSubspaceOwnership(ctx sdk.Context, subspaceID, newOwner string) {
	store := ctx.KVStore(k.storeKey)
	key := types.SubspaceStoreKey(subspaceID)

	var subspace types.Subspace
	k.cdc.MustUnmarshalBinaryBare(store.Get(key), &subspace)

	// set new owner
	subspace.Owner = newOwner
	store.Set(key, k.cdc.MustMarshalBinaryBare(&subspace))
}

// AddAdminToSubspace insert the newAdmin inside the admins list of the given subspaces if its not present.
// Returns an error if the admin is already present.
func (k Keeper) AddAdminToSubspace(ctx sdk.Context, subspaceID, user string) error {
	if err := k.AddUserToList(ctx, types.AdminsStoreKey(subspaceID), subspaceID, user,
		"the user: %s is already an admin of the subspaces: %s"); err != nil {
		return err
	}
	return nil
}

// IsAdmin returns true if the given address is an admin inside the given subspaces id, false otherwise.
func (k Keeper) IsAdmin(ctx sdk.Context, address, subspaceID string) bool {
	store := ctx.KVStore(k.storeKey)
	key := types.AdminsStoreKey(subspaceID)

	admins := types.MustUnmarshalUsers(k.cdc, store.Get(key))
	return admins.IsPresent(address)
}

// GetSubspaceAdmins returns a list of all the subspaces admins
func (k Keeper) GetSubspaceAdmins(ctx sdk.Context, subspaceID string) types.Users {
	store := ctx.KVStore(k.storeKey)
	key := types.AdminsStoreKey(subspaceID)

	return types.MustUnmarshalUsers(k.cdc, store.Get(key))
}

// RemoveAdminFromSubspace remove the given admin from the given subspaces.
// It returns error when the admin is not present inside the subspaces.
func (k Keeper) RemoveAdminFromSubspace(ctx sdk.Context, subspaceID, admin string) error {
	// If the admin doesn't exist, return error
	if err := k.RemoveUserFromList(ctx, types.AdminsStoreKey(subspaceID), subspaceID, admin,
		"this address: %s is not an admin of the subspaces %s"); err != nil {
		return err
	}
	return nil
}

// UnblockPostsForUser give a user the possibility to post inside the given subspaces.
// It returns error when the user can already post inside the subspaces.
func (k Keeper) UnblockPostsForUser(ctx sdk.Context, user, subspaceID string) error {
	if err := k.RemoveUserFromList(ctx, types.BlockedToPostUsersKey(subspaceID), subspaceID, user,
		"the user: %s is already allowed to post inside the subspaces: %s"); err != nil {
		return err
	}
	return nil
}

// BlockPostsForUser block the given user to post anything inside the given subspaces.
// It returns error if the user already can't post inside the subspaces.
func (k Keeper) BlockPostsForUser(ctx sdk.Context, userToBlock, subspaceID string) error {
	if err := k.AddUserToList(ctx, types.BlockedToPostUsersKey(subspaceID), subspaceID, userToBlock,
		"the user: %s already can't post inside the subspaces: %s"); err != nil {
		return err
	}
	return nil
}

// GetSubspaceBlockedUsers returns a list of all the blocked users unable to post inside the given subspaces
func (k Keeper) GetSubspaceBlockedUsers(ctx sdk.Context, subspaceID string) types.Users {
	store := ctx.KVStore(k.storeKey)
	key := types.BlockedToPostUsersKey(subspaceID)

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
		subspaceID := string(idBytes)
		entries = append(entries, types.NewAdminsEntries(subspaceID, admins))
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
		subspaceID := string(idBytes)
		entries = append(entries, types.NewBlockedUsersEntry(subspaceID, users))
	}

	return entries
}
