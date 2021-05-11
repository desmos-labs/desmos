package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

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

// CheckSubspaceExistenceAndCreatorValidity check if the subspace with the given id exists and
// if the address belongs to its creator
func (k Keeper) CheckSubspaceExistenceAndCreatorValidity(ctx sdk.Context, subspaceId, address string) error {
	subspace, exist := k.GetSubspace(ctx, subspaceId)
	if !exist {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the subspace with id %s doesn't exist", subspaceId,
		)
	}

	if subspace.Creator != address {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the user: %s is not the subspace creator and can't perform this operation on the subspace: %s",
			address, subspaceId,
		)
	}

	return nil
}
