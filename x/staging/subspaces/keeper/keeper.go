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
func (k Keeper) SaveSubspace(ctx sdk.Context, subspace types.Subspace, user string) error {
	err := subspace.Validate()
	if err != nil {
		return err
	}

	storedSubspace, found := k.GetSubspace(ctx, subspace.ID)

	// Check the editor when the user is trying to edit the subspace
	if found && storedSubspace.Owner != user {
		return sdkerrors.Wrapf(types.ErrInvalidSubspaceOwner, user)
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

// AddAdminToSubspace sets the given user as an admin of the subspace having the given id.
// Returns an error if the user is already an admin or the subspace does not exist.
func (k Keeper) AddAdminToSubspace(ctx sdk.Context, subspaceID, user, owner string) error {
	// Check if the subspace exists and the owner is valid
	_, err := k.checkSubspaceAndOwner(ctx, subspaceID, owner)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	key := types.SubspaceAdminKey(subspaceID, user)

	// Check if the user we want to set as admin is already an admin
	if store.Has(key) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "the user with address %s is already an admin", user)
	}

	// Store the admin
	store.Set(key, []byte(user))
	return nil
}

// RemoveAdminFromSubspace removes the given user from the admin set of the subspace having the given id.
// It returns an error if the user was not an admin or the subspace does not exist.
func (k Keeper) RemoveAdminFromSubspace(ctx sdk.Context, subspaceID, user, owner string) error {
	// Check if the subspace exists and the owner is valid
	_, err := k.checkSubspaceAndOwner(ctx, subspaceID, owner)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	key := types.SubspaceAdminKey(subspaceID, user)

	// Check if the user is not an admin
	if !store.Has(key) {
		return sdkerrors.Wrapf(types.ErrInvalidSubspaceAdmin, user)
	}

	// Delete the admin
	store.Delete(key)
	return nil
}

// RegisterUserInSubspace registers the given user inside the subspace with the given ID.
// It returns error if the user is already registered or the subspace does not exist.
func (k Keeper) RegisterUserInSubspace(ctx sdk.Context, subspaceID, user, admin string) error {
	// Check if the subspace exists and the admin is an actual admin
	_, err := k.checkSubspaceAndAdmin(ctx, subspaceID, admin)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	key := types.SubspaceRegisteredUserKey(subspaceID, user)

	// Check if the user is already registered inside the subspace
	if store.Has(key) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "user already registered")
	}

	// Store the new user
	store.Set(key, []byte(user))
	return nil
}

// UnregisterUserFromSubspace unregisters the given user from the subspace with the given ID.
// It returns error if the user is not registered or the subspace does not exist.
func (k Keeper) UnregisterUserFromSubspace(ctx sdk.Context, subspaceID, user, admin string) error {
	// Check if the subspace exists and the admin is an actual admin
	_, err := k.checkSubspaceAndAdmin(ctx, subspaceID, admin)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	key := types.SubspaceRegisteredUserKey(subspaceID, user)

	// Check if the user is already registered inside the subspace
	if !store.Has(key) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the user with address : %s is not registered inside the subspace: %s", user, subspaceID)
	}

	// Remove the user
	store.Delete(key)
	return nil
}

// BanUserInSubspace bans the given user inside the subspace with the given ID.
// It returns and error if the user is already blocked inside the subspace or the subspace does not exist.
func (k Keeper) BanUserInSubspace(ctx sdk.Context, subspaceID, user, admin string) error {
	// Check if the subspace exists and the admin is an actual admin
	_, err := k.checkSubspaceAndAdmin(ctx, subspaceID, admin)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	key := types.SubspaceBannedUserKey(subspaceID, user)

	// Check if the user is already banned inside the subspace
	if store.Has(key) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "user already banned")
	}

	// Store the banned user
	store.Set(key, []byte(user))
	return nil
}

// UnbanUserInSubspace unbans the given user inside the subspace with the given ID.
// It returns error if the user is not banned inside the subspace or the subspace does not exist.
func (k Keeper) UnbanUserInSubspace(ctx sdk.Context, subspaceID, user, admin string) error {
	// Check if the subspace exists and the admin is an actual admin
	_, err := k.checkSubspaceAndAdmin(ctx, subspaceID, admin)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	key := types.SubspaceBannedUserKey(subspaceID, user)

	// Check if the user is already banned inside the subspace
	if !store.Has(key) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "the user is not blocked inside the subspace")
	}

	// Remove the banned user
	store.Delete(key)
	return nil
}
