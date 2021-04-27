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
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace already exists")
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
// If no subspace is associated with the given ID the function will return an error.
func (k Keeper) GetSubspace(ctx sdk.Context, subspaceId string) (subspace types.Subspace, found bool) {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(types.SubspaceStoreKey(subspaceId)) {
		return types.Subspace{}, false
	}

	k.cdc.MustUnmarshalBinaryBare(store.Get(types.SubspaceStoreKey(subspaceId)), &subspace)
	return subspace, true
}

// AddAdminToSubspace insert the given admin inside the admins list if its not present.
// Returns an error if the admin is already present.
func (k Keeper) AddAdminToSubspace(ctx sdk.Context, subspaceId, newAdmin string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.AdminsStoreKey(subspaceId)

	wrapped := types.MustUnmarshalAdmins(k.cdc, store.Get(key))
	for _, admin := range wrapped.Admins {
		if admin == newAdmin {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"the user is already an admin of subspace: %s", subspaceId)
		}
	}

	wrapped.Admins = append(wrapped.Admins, newAdmin)
	store.Set(key, types.MustMarshalAdmins(k.cdc, wrapped))
	return nil
}

// GetAllSubspaceAdmins returns a list of all the subspace admins
func (k Keeper) GetAllSubspaceAdmins(ctx sdk.Context, subspaceId string) types.Admins {
	store := ctx.KVStore(k.storeKey)
	key := types.AdminsStoreKey(subspaceId)

	return types.MustUnmarshalAdmins(k.cdc, store.Get(key))
}

func (k Keeper) RemoveAdminFromSubspace(ctx sdk.Context, subspaceId, admin string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.AdminsStoreKey(subspaceId)

	wrapped := types.MustUnmarshalAdmins(k.cdc, store.Get(key))
	admins, found := types.RemoveAdmin(wrapped.Admins, admin)

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
		store.Set(key, types.MustMarshalAdmins(k.cdc, types.Admins{Admins: admins}))
	}

	return nil
}
