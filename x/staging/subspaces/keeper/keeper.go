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
func (k Keeper) SaveSubspace(ctx sdk.Context, subspace types.Subspace) {
	store := ctx.KVStore(k.storeKey)
	key := types.SubspaceStoreKey(subspace.ID)
	store.Set(key, k.cdc.MustMarshalBinaryBare(&subspace))
}

// DoesSubspaceExists returns true if the subspaces with the given id exists inside the store.
func (k Keeper) DoesSubspaceExists(ctx sdk.Context, subspaceID string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.SubspaceStoreKey(subspaceID))
}

// GetSubspace returns the subspaces associated with the given ID.
// If there is no subspace associated with the given ID the function will return an empty subspace and false.
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
	var subspaces []types.Subspace
	k.IterateSubspaces(ctx, func(_ int64, subspace types.Subspace) (stop bool) {
		subspaces = append(subspaces, subspace)
		return false
	})

	return subspaces
}

// AddAdminToSubspace insert the newAdmin inside the admins list of the given subspaces if its not present.
// Returns an error if the admin is already present.
func (k Keeper) AddAdminToSubspace(ctx sdk.Context, subspaceID, user, owner string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.SubspaceStoreKey(subspaceID)

	subspaceBytes := store.Get(key)
	// check if the subspace exists and the owner is the actual owner of it
	subspace, err := k.CheckSubspaceAndOwner(subspaceBytes, subspaceID, owner)
	if err != nil {
		return err
	}

	// check if the user we want to set as admin is already an admin
	if subspace.Admins.IsPresent(user) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "the user with address : %s is already an admin", user)
	}

	subspace.Admins = subspace.Admins.AppendUser(user)

	store.Set(key, k.cdc.MustMarshalBinaryBare(&subspace))
	return nil
}

// RemoveAdminFromSubspace remove the given admin from the given subspaces.
// It returns error when the admin is not present inside the subspaces.
func (k Keeper) RemoveAdminFromSubspace(ctx sdk.Context, subspaceID, user, owner string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.SubspaceStoreKey(subspaceID)

	subspaceBytes := store.Get(key)
	// check if the subspace exists and the owner is the actual owner of it
	subspace, err := k.CheckSubspaceAndOwner(subspaceBytes, subspaceID, owner)
	if err != nil {
		return err
	}

	// check if the user is not an admin
	if !subspace.Admins.IsPresent(user) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "the user with address : %s is not an admin", user)
	}

	subspace.Admins = subspace.Admins.RemoveUser(user)
	store.Set(key, k.cdc.MustMarshalBinaryBare(&subspace))
	return nil
}

// RegisterUserInSubspace register the user in the subspace with the given ID.
// It returns error when the user is already registered.
func (k Keeper) RegisterUserInSubspace(ctx sdk.Context, subspaceID, user, admin string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.SubspaceStoreKey(subspaceID)

	subspaceBytes := store.Get(key)
	// check if the subspace exists and the admin is an actual admin
	subspace, err := k.CheckSubspaceAndAdmin(subspaceBytes, subspaceID, admin)
	if err != nil {
		return err
	}

	// check if the user is already registered inside the subspace
	if subspace.RegisteredUsers.IsPresent(user) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the user with address : %s is already registered inside the subspace: %s", user, subspaceID)
	}

	subspace.RegisteredUsers = subspace.RegisteredUsers.AppendUser(user)

	store.Set(key, k.cdc.MustMarshalBinaryBare(&subspace))
	return nil
}

// BlockUserInSubspace block the given user inside the given subspace.
// It returns error if the user is already blocked inside the subspace.
func (k Keeper) BlockUserInSubspace(ctx sdk.Context, subspaceID, user, admin string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.SubspaceStoreKey(subspaceID)

	subspaceBytes := store.Get(key)
	// check if the subspace exists and the admin is an actual admin
	subspace, err := k.CheckSubspaceAndAdmin(subspaceBytes, subspaceID, admin)
	if err != nil {
		return err
	}

	// check if the user is already registered inside the subspace
	if subspace.BlockedUsers.IsPresent(user) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the user with address : %s is already blocked inside the subspace: %s", user, subspaceID)
	}

	subspace.RegisteredUsers = subspace.BlockedUsers.AppendUser(user)

	store.Set(key, k.cdc.MustMarshalBinaryBare(&subspace))
	return nil
}
