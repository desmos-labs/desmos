package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
)

func NewEditParamsProposalHandler(k Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case models.EditNameSurnameParamsProposal:
			return handleNameSurnameEditParamsProposal(ctx, k, c)
		case models.EditMonikerParamsProposal:
			return handleMonikerEditParamsProposal(ctx, k, c)
		case models.EditBioParamsProposal:
			return handleBioEditParamsProposal(ctx, k, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized profiles module proposal type %T", c)
		}
	}
}

// handleNameSurnameEditParamsProposal handles the edit of name surname parameters
func handleNameSurnameEditParamsProposal(ctx sdk.Context, k Keeper, proposal models.EditNameSurnameParamsProposal) error {
	params := k.GetParams(ctx)
	params = params.WithNameSurnameParams(proposal.NameSurnameParams)
	k.SetParams(ctx, params)
	return nil
}

// handleMonikerEditParamsProposal handles the edit of moniker parameters
func handleMonikerEditParamsProposal(ctx sdk.Context, k Keeper, proposal models.EditMonikerParamsProposal) error {
	params := k.GetParams(ctx)
	params = params.WithMonikerParams(proposal.MonikerParams)
	k.SetParams(ctx, params)
	return nil
}

// handleBioEditParamsProposal handles the edit of biography parameter
func handleBioEditParamsProposal(ctx sdk.Context, k Keeper, proposal models.EditBioParamsProposal) error {
	params := k.GetParams(ctx)
	params.BiographyLengths = proposal.BioParams
	k.SetParams(ctx, params)
	return nil
}
