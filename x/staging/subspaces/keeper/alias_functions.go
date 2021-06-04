package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
)

// IterateSubspaces iterates through the subspaces set and performs the given function
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

// GetAllSubspaces returns a list of all the subspaces that have been store inside the given context
func (k Keeper) GetAllSubspaces(ctx sdk.Context) []types.Subspace {
	var subspaces []types.Subspace
	k.IterateSubspaces(ctx, func(_ int64, subspace types.Subspace) (stop bool) {
		subspaces = append(subspaces, subspace)
		return false
	})

	return subspaces
}

// checkSubspaceAndAdmin check if the subspaces with the given id exists and
// if the address belongs to the admin of the subspace or one of its admins.
// It returns an error or the subspace itself if everything's fine.
func (k Keeper) checkSubspaceAndAdmin(ctx sdk.Context, id, address string) (types.Subspace, error) {
	subspace, found := k.GetSubspace(ctx, id)
	if !found {
		return types.Subspace{}, sdkerrors.Wrapf(types.ErrInvalidSubspaceID,
			"the subspace with id %s doesn't exist", id,
		)
	}

	if subspace.Owner != address {
		if subspace.IsAdmin(address) {
			return types.Subspace{}, sdkerrors.Wrapf(types.ErrInvalidSubspaceAdmin,
				"%s is not a subspace admin and can't perform this operation on the subspaces: %s",
				address, id,
			)
		}
	}

	return subspace, nil
}

// checkSubspaceAndOwner check if the subspaces with the given id exists and
// if the address belongs to the admin of the subspace.
// It returns an error or the subspace itself if everything's fine.
func (k Keeper) checkSubspaceAndOwner(ctx sdk.Context, id, address string) (types.Subspace, error) {
	subspace, found := k.GetSubspace(ctx, id)
	if !found {
		return types.Subspace{}, sdkerrors.Wrapf(types.ErrInvalidSubspaceID,
			"the subspace with id %s doesn't exist", id,
		)
	}

	if subspace.Owner != address {
		return types.Subspace{}, sdkerrors.Wrapf(types.ErrInvalidSubspaceOwner,
			"%s is not the subspace owner and can't perform this operation on the subspaces: %s",
			address, id,
		)
	}

	return subspace, nil
}
