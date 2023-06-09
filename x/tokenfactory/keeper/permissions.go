package keeper

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	tokenfactorytypes "github.com/osmosis-labs/osmosis/v15/x/tokenfactory/types"

	subspacestypes "github.com/desmos-labs/desmos/v5/x/subspaces/types"
	"github.com/desmos-labs/desmos/v5/x/tokenfactory/types"
)

// ValidateManageTokenPermission validates the sender has the manage denom permission to the subspace tokens inside the given subspace
func (k Keeper) ValidateManageTokenPermission(ctx sdk.Context, subspace subspacestypes.Subspace, sender string, denom string) error {

	// Check the permission to manage the subspace tokens
	if !k.sk.HasPermission(ctx, subspace.ID, subspacestypes.RootSectionID, sender, types.PermissionManageSubspaceTokens) {
		return errors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot manage the subspace tokens inside this subspace")
	}

	// Check if the denom exists
	_, denomExists := k.bk.GetDenomMetaData(ctx, denom)
	if !denomExists {
		return tokenfactorytypes.ErrDenomDoesNotExist.Wrapf("denom: %s", denom)
	}

	authorityMetadata, err := k.tfk.GetAuthorityMetadata(ctx, denom)
	if err != nil {
		return err
	}

	// Check if the subspace treasury is the admin of the denom
	if subspace.Treasury != authorityMetadata.GetAdmin() {
		return tokenfactorytypes.ErrUnauthorized
	}

	return nil
}
