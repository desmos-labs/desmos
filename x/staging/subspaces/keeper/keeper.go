package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/libs/log"

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

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
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

	k.Logger(ctx).Info("saved subspace", "id", subspace.ID, "owner", subspace.Owner)
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

// IsAdmin returns true iff the given user is an admin of the subspace with the given id
func (k Keeper) IsAdmin(ctx sdk.Context, subspaceID string, user string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.SubspaceAdminKey(subspaceID, user))
}

// AddAdminToSubspace sets the given user as an admin of the subspace having the given id.
// Returns an error if the user is already an admin or the subspace does not exist.
func (k Keeper) AddAdminToSubspace(ctx sdk.Context, subspaceID, user, owner string) error {
	// Check if the subspace exists and the owner is valid
	err := k.checkSubspaceOwner(ctx, subspaceID, owner)
	if err != nil {
		return err
	}

	// Check if the user we want to set as admin is already an admin
	if k.IsAdmin(ctx, subspaceID, user) {
		return sdkerrors.Wrapf(types.ErrDuplicatedAdmin, user)
	}

	// Store the admin
	store := ctx.KVStore(k.storeKey)
	store.Set(types.SubspaceAdminKey(subspaceID, user), []byte(user))

	k.Logger(ctx).Info("added admin", "subspace-id", subspaceID, "admin", user)
	return nil
}

// RemoveAdminFromSubspace removes the given user from the admin set of the subspace having the given id.
// It returns an error if the user was not an admin or the subspace does not exist.
func (k Keeper) RemoveAdminFromSubspace(ctx sdk.Context, subspaceID, user, owner string) error {
	// Check if the subspace exists and the owner is valid
	err := k.checkSubspaceOwner(ctx, subspaceID, owner)
	if err != nil {
		return err
	}

	// Check if the user is not an admin
	if !k.IsAdmin(ctx, subspaceID, user) {
		return sdkerrors.Wrapf(types.ErrInvalidAdmin, user)
	}

	// Delete the admin
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.SubspaceAdminKey(subspaceID, user))

	k.Logger(ctx).Info("removed admin", "subspace-id", subspaceID, "admin", user)
	return nil
}

// IsRegistered returns true iff the given user is registered inside the subspace with the given id
func (k Keeper) IsRegistered(ctx sdk.Context, subspaceID, user string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.SubspaceRegisteredUserKey(subspaceID, user))
}

// RegisterUserInSubspace registers the given user inside the subspace with the given ID.
// It returns error if the user is already registered or the subspace does not exist.
func (k Keeper) RegisterUserInSubspace(ctx sdk.Context, subspaceID, user, admin string) error {
	// Check if the subspace exists and the admin is an actual admin
	err := k.checkSubspaceAdmin(ctx, subspaceID, admin)
	if err != nil {
		return err
	}

	// Check if the user is already registered inside the subspace
	if k.IsRegistered(ctx, subspaceID, user) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "user already registered")
	}

	// Store the new user
	store := ctx.KVStore(k.storeKey)
	store.Set(types.SubspaceRegisteredUserKey(subspaceID, user), []byte(user))
	return nil
}

// UnregisterUserFromSubspace unregisters the given user from the subspace with the given ID.
// It returns error if the user is not registered or the subspace does not exist.
func (k Keeper) UnregisterUserFromSubspace(ctx sdk.Context, subspaceID, user, admin string) error {
	// Check if the subspace exists and the admin is an actual admin
	err := k.checkSubspaceAdmin(ctx, subspaceID, admin)
	if err != nil {
		return err
	}

	// Check if the user is already registered inside the subspace
	if !k.IsRegistered(ctx, subspaceID, user) {
		return sdkerrors.Wrapf(types.ErrUserNotFound,
			"user: %s, subspace: %s", user, subspaceID)
	}

	// Remove the user
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.SubspaceRegisteredUserKey(subspaceID, user))
	return nil
}

// IsBanned returns true iff the given user is banned inside the subspace with the given id
func (k Keeper) IsBanned(ctx sdk.Context, subspaceID, user string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.SubspaceBannedUserKey(subspaceID, user))
}

// BanUserInSubspace bans the given user inside the subspace with the given ID.
// It returns and error if the user is already blocked inside the subspace or the subspace does not exist.
func (k Keeper) BanUserInSubspace(ctx sdk.Context, subspaceID, user, admin string) error {
	// Check if the subspace exists and the admin is an actual admin
	err := k.checkSubspaceAdmin(ctx, subspaceID, admin)
	if err != nil {
		return err
	}

	// Check if the user is already banned inside the subspace
	if k.IsBanned(ctx, subspaceID, user) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "user already banned")
	}

	// Store the banned user
	store := ctx.KVStore(k.storeKey)
	store.Set(types.SubspaceBannedUserKey(subspaceID, user), []byte(user))
	return nil
}

// UnbanUserInSubspace unbans the given user inside the subspace with the given ID.
// It returns error if the user is not banned inside the subspace or the subspace does not exist.
func (k Keeper) UnbanUserInSubspace(ctx sdk.Context, subspaceID, user, admin string) error {
	// Check if the subspace exists and the admin is an actual admin
	err := k.checkSubspaceAdmin(ctx, subspaceID, admin)
	if err != nil {
		return err
	}

	// Check if the user is already banned inside the subspace
	if !k.IsBanned(ctx, subspaceID, user) {
		return sdkerrors.Wrapf(types.ErrUserNotFound, "%s is not banned inside the subspace", user)
	}

	// Remove the banned user
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.SubspaceBannedUserKey(subspaceID, user))
	return nil
}

// CheckSubspaceUserPermission checks the permission of the given user inside the subspace with the
// given id to make sure they are able to perform operations inside it
func (k Keeper) CheckSubspaceUserPermission(ctx sdk.Context, subspaceID string, user string) error {
	subspace, found := k.GetSubspace(ctx, subspaceID)
	if !found {
		return sdkerrors.Wrapf(types.ErrInvalidSubspaceID, subspaceID)
	}

	if k.IsBanned(ctx, subspaceID, user) {
		return sdkerrors.Wrapf(types.ErrPermissionDenied, user)
	}

	if subspace.Type == types.SubspaceTypeClosed && !k.IsRegistered(ctx, subspaceID, user) {
		return sdkerrors.Wrapf(types.ErrPermissionDenied, user)
	}

	return nil
}
