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

		var stop = false
		stop = fn(i, subspace)

		if stop {
			break
		}

		i++
	}
}

// CheckSubspaceExistenceAndAdminValidity checks if the subspace with the given id exists and
// if the address belongs to one of its admins
func (k Keeper) CheckSubspaceExistenceAndAdminValidity(ctx sdk.Context, address, subspaceId string) error {
	if !k.DoesSubspaceExists(ctx, subspaceId) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the subspace with id %s doesn't exist", subspaceId,
		)
	}

	if !k.IsAdmin(ctx, address, subspaceId) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the user: %s is not an admin and can't perform this operation on the subspace: %s",
			address, subspaceId)
	}

	return nil
}

// CheckSubspaceExistenceAndOwnerValidity check if the subspace with the given id exists and
// if the address belongs to its creator
func (k Keeper) CheckSubspaceExistenceAndOwnerValidity(ctx sdk.Context, address, subspaceId string) error {
	subspace, exist := k.GetSubspace(ctx, subspaceId)
	if !exist {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the subspace with id %s doesn't exist", subspaceId,
		)
	}

	if subspace.Owner != address {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the user: %s is not the subspace owner and can't perform this operation on the subspace: %s",
			address, subspaceId,
		)
	}

	return nil
}
