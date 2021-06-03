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
// It assumes that the subspace has been validated already.
func (k Keeper) SaveSubspace(ctx sdk.Context, subspace types.Subspace, user string) error {
	storedSubspace, found := k.GetSubspace(ctx, subspace.ID)
	// check performed when the user is trying to edit the subspace
	if found && storedSubspace.Owner != user {
		return sdkerrors.Wrapf(types.ErrInvalidSubspaceOwner,
			"%s is not the subspace owner and can't perform this operation on the subspace: %s", user, subspace.ID)
	}

	store := ctx.KVStore(k.storeKey)
	key := types.SubspaceStoreKey(subspace.ID)
	store.Set(key, k.cdc.MustMarshalBinaryBare(&subspace))
	return nil
}

// DoesSubspaceExist checks if the subspace with the given id exists.
func (k Keeper) DoesSubspaceExist(ctx sdk.Context, subspaceID string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.SubspaceStoreKey(subspaceID))
}

// GetSubspace returns the subspace associated with the given ID.
// If there is no subspace associated with the given ID the function will return an empty subspace and false.
func (k Keeper) GetSubspace(ctx sdk.Context, subspaceID string) (subspace types.Subspace, found bool) {
	if !k.DoesSubspaceExist(ctx, subspaceID) {
		return types.Subspace{}, false
	}

	store := ctx.KVStore(k.storeKey)
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.SubspaceStoreKey(subspaceID)), &subspace)
	return subspace, true
}

// AddAdminToSubspace insert the user inside the admins array of the given subspace if his not present.
// Returns an error if the user is already an admin.
func (k Keeper) AddAdminToSubspace(ctx sdk.Context, subspaceID, user, owner string) error {
	// check if the subspace exists and the admin is the actual admin of it
	subspace, err := k.checkSubspaceAndOwner(ctx, subspaceID, owner)
	if err != nil {
		return err
	}

	// check if the user we want to set as admin is already an admin
	if subspace.IsAdmin(user) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "the user with address : %s is already an admin", user)
	}

	subspace.Admins = append(subspace.Admins, user)

	store := ctx.KVStore(k.storeKey)
	key := types.SubspaceStoreKey(subspace.ID)
	store.Set(key, k.cdc.MustMarshalBinaryBare(&subspace))
	return nil
}

// RemoveAdminFromSubspace remove the given admin from the given subspace.
// It returns error when the admin is not present inside the subspace.
func (k Keeper) RemoveAdminFromSubspace(ctx sdk.Context, subspaceID, user, owner string) error {
	// check if the subspace exists and the admin is the actual admin of it
	subspace, err := k.checkSubspaceAndOwner(ctx, subspaceID, owner)
	if err != nil {
		return err
	}

	// check if the user is not an admin
	if !subspace.IsAdmin(user) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "the user with address : %s is not an admin", user)
	}

	subspace.Admins = types.RemoveUser(subspace.Admins, user)

	store := ctx.KVStore(k.storeKey)
	key := types.SubspaceStoreKey(subspace.ID)
	store.Set(key, k.cdc.MustMarshalBinaryBare(&subspace))
	return nil
}

// RegisterUserInSubspace register the user in the subspace with the given ID.
// It returns error when the user is already registered.
func (k Keeper) RegisterUserInSubspace(ctx sdk.Context, subspaceID, user, admin string) error {
	// check if the subspace exists and the admin is an actual admin
	subspace, err := k.checkSubspaceAndAdmin(ctx, subspaceID, admin)
	if err != nil {
		return err
	}

	// check if the user is already registered inside the subspace
	if subspace.IsRegistered(user) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the user with address : %s is already registered inside the subspace: %s", user, subspaceID)
	}

	subspace.RegisteredUsers = append(subspace.RegisteredUsers, user)

	store := ctx.KVStore(k.storeKey)
	key := types.SubspaceStoreKey(subspace.ID)
	store.Set(key, k.cdc.MustMarshalBinaryBare(&subspace))
	return nil
}

// UnregisterUserFromSubspace unregister the user from the subspace with the given ID.
// It returns error when the user is not registered.
func (k Keeper) UnregisterUserFromSubspace(ctx sdk.Context, subspaceID, user, admin string) error {
	// check if the subspace exists and the admin is an actual admin
	subspace, err := k.checkSubspaceAndAdmin(ctx, subspaceID, admin)
	if err != nil {
		return err
	}

	// check if the user is already registered inside the subspace
	if !subspace.IsRegistered(user) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the user with address : %s is not registered inside the subspace: %s", user, subspaceID)
	}

	subspace.RegisteredUsers = types.RemoveUser(subspace.RegisteredUsers, user)

	store := ctx.KVStore(k.storeKey)
	key := types.SubspaceStoreKey(subspace.ID)
	store.Set(key, k.cdc.MustMarshalBinaryBare(&subspace))
	return nil
}

// BanUserInSubspace block the given user inside the given subspace.
// It returns error if the user is already blocked inside the subspace.
func (k Keeper) BanUserInSubspace(ctx sdk.Context, subspaceID, user, admin string) error {
	// check if the subspace exists and the admin is an actual admin
	subspace, err := k.checkSubspaceAndAdmin(ctx, subspaceID, admin)
	if err != nil {
		return err
	}

	// check if the user is already registered inside the subspace
	if subspace.IsBanned(user) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the user with address : %s is already blocked inside the subspace: %s", user, subspaceID)
	}

	subspace.BannedUsers = append(subspace.BannedUsers, user)

	store := ctx.KVStore(k.storeKey)
	key := types.SubspaceStoreKey(subspace.ID)
	store.Set(key, k.cdc.MustMarshalBinaryBare(&subspace))
	return nil
}

// UnbanUserInSubspace unblock the given user inside the given subspace.
// It returns error if the user is not blocked inside the subspace.
func (k Keeper) UnbanUserInSubspace(ctx sdk.Context, subspaceID, user, admin string) error {
	// check if the subspace exists and the admin is an actual admin
	subspace, err := k.checkSubspaceAndAdmin(ctx, subspaceID, admin)
	if err != nil {
		return err
	}

	// check if the user is already registered inside the subspace
	if !subspace.IsBanned(user) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the user with address : %s is not blocked inside the subspace: %s", user, subspaceID)
	}

	subspace.BannedUsers = types.RemoveUser(subspace.BannedUsers, user)

	store := ctx.KVStore(k.storeKey)
	key := types.SubspaceStoreKey(subspace.ID)
	store.Set(key, k.cdc.MustMarshalBinaryBare(&subspace))
	return nil
}
