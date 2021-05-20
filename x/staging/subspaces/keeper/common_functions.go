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

// CheckSubspaceAndAdmin check if the subspaces with the given id exists and
// if the address belongs to the admin of the subspace or one of its admins.
// It returns an error or the subspace itself if everything's fine.
func (k Keeper) CheckSubspaceAndAdmin(subspaceBytes []byte, subspaceID, address string) (types.Subspace, error) {
	var subspace types.Subspace
	if subspaceBytes == nil {
		return types.Subspace{}, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the subspace with id %s doesn't exist", subspaceID,
		)
	}

	k.cdc.MustUnmarshalBinaryBare(subspaceBytes, &subspace)

	if subspace.Owner != address {
		if !subspace.Admins.IsPresent(address) {
			return types.Subspace{}, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"the user: %s is not the subspace admin or an admin and can't perform this operation on the subspaces: %s",
				address, subspaceID,
			)
		}
	}

	return subspace, nil
}

// CheckSubspaceAndOwner check if the subspaces with the given id exists and
// if the address belongs to the admin of the subspace.
// It returns an error or the subspace itself if everything's fine.
func (k Keeper) CheckSubspaceAndOwner(subspaceBytes []byte, subspaceID, address string) (types.Subspace, error) {
	var subspace types.Subspace
	if subspaceBytes == nil {
		return types.Subspace{}, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the subspace with id %s doesn't exist", subspaceID,
		)
	}

	k.cdc.MustUnmarshalBinaryBare(subspaceBytes, &subspace)

	if subspace.Owner != address {
		return types.Subspace{}, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the user: %s is not the subspace admin and can't perform this operation on the subspaces: %s",
			address, subspaceID,
		)
	}

	return subspace, nil
}
