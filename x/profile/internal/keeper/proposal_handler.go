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
		case models.NameSurnameParamsEditProposal:
			return handleNameSurnameEditParamsProposal(ctx, k, c)
		case models.MonikerParamsEditProposal:
			return handleMonikerEditParamsProposal(ctx, k, c)
		case models.BioParamsEditProposal:
			return handleBioEditParamsProposal(ctx, k, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized profiles module proposal type %T", c)
		}
	}
}

// handleNameSurnameEditParamsProposal handles the edit of name surname parameters
func handleNameSurnameEditParamsProposal(ctx sdk.Context, k Keeper, proposal models.NameSurnameParamsEditProposal) error {
	actualParams := proposal.NameSurnameParams

	if proposal.NameSurnameParams.MinNameSurnameLen == nil {
		savedParams := k.GetNameSurnameLenParams(ctx)
		actualParams = models.NewNameSurnameLenParams(savedParams.MinNameSurnameLen, proposal.NameSurnameParams.MaxNameSurnameLen)
	} else if proposal.NameSurnameParams.MaxNameSurnameLen == nil {
		savedParams := k.GetNameSurnameLenParams(ctx)
		actualParams = models.NewNameSurnameLenParams(proposal.NameSurnameParams.MinNameSurnameLen, savedParams.MaxNameSurnameLen)
	}

	k.SetNameSurnameLenParams(ctx, actualParams)
	return nil
}

// handleMonikerEditParamsProposal handles the edit of moniker parameters
func handleMonikerEditParamsProposal(ctx sdk.Context, k Keeper, proposal models.MonikerParamsEditProposal) error {
	actualParams := proposal.MonikerParams

	if proposal.MonikerParams.MinMonikerLen == nil {
		savedParams := k.GetMonikerLenParams(ctx)
		actualParams = models.NewMonikerLenParams(savedParams.MinMonikerLen, proposal.MonikerParams.MaxMonikerLen)
	} else if proposal.MonikerParams.MaxMonikerLen == nil {
		savedParams := k.GetMonikerLenParams(ctx)
		actualParams = models.NewMonikerLenParams(proposal.MonikerParams.MinMonikerLen, savedParams.MaxMonikerLen)
	}

	k.SetMonikerLenParams(ctx, actualParams)
	return nil
}

// handleBioEditParamsProposal handles the edit of biography parameter
func handleBioEditParamsProposal(ctx sdk.Context, k Keeper, proposal models.BioParamsEditProposal) error {
	k.SetBioLenParams(ctx, proposal.BioParams)
	return nil
}
