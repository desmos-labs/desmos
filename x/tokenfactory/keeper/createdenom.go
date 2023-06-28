package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/desmos-labs/desmos/v5/x/tokenfactory/types"
)

// ConvertToBaseToken converts a fee amount in a whitelisted fee token to the base fee token amount
func (k Keeper) CreateDenom(ctx sdk.Context, creator string, subdenom string) (newTokenDenom string, err error) {
	denom, err := k.validateCreateDenom(ctx, creator, subdenom)
	if err != nil {
		return "", err
	}

	err = k.chargeForCreateDenom(ctx, creator)
	if err != nil {
		return "", err
	}

	err = k.createDenomAfterValidation(ctx, creator, denom)
	return denom, err
}

// Runs CreateDenom logic after the charge and all denom validation has been handled.
// Made into a second function for genesis initialization.
func (k Keeper) createDenomAfterValidation(ctx sdk.Context, creator string, denom string) (err error) {
	_, exists := k.bk.GetDenomMetaData(ctx, denom)
	if !exists {
		denomMetaData := banktypes.Metadata{
			DenomUnits: []*banktypes.DenomUnit{{
				Denom:    denom,
				Exponent: 0,
			}},
			Base: denom,
		}

		k.bk.SetDenomMetaData(ctx, denomMetaData)
	}

	authorityMetadata := types.DenomAuthorityMetadata{
		Admin: creator,
	}
	err = k.SetAuthorityMetadata(ctx, denom, authorityMetadata)
	if err != nil {
		return err
	}

	k.AddDenomFromCreator(ctx, creator, denom)
	return nil
}

func (k Keeper) validateCreateDenom(ctx sdk.Context, creator string, subdenom string) (newTokenDenom string, err error) {
	// Temporary check until IBC bug is sorted out
	if k.bk.HasSupply(ctx, subdenom) {
		return "", fmt.Errorf("temporary error until IBC bug is sorted out, " +
			"can't create subdenoms that are the same as a native denom")
	}

	denom, err := types.GetTokenDenom(creator, subdenom)
	if err != nil {
		return "", err
	}

	_, found := k.bk.GetDenomMetaData(ctx, denom)
	if found {
		return "", types.ErrDenomExists
	}

	return denom, nil
}

func (k Keeper) chargeForCreateDenom(ctx sdk.Context, creator string) (err error) {
	creationFee := k.GetParams(ctx).DenomCreationFee

	// Burn creation fee
	if creationFee != nil {
		err := k.bk.SendCoinsFromAccountToModule(ctx,
			sdk.MustAccAddressFromBech32(creator),
			types.ModuleName,
			creationFee)
		if err != nil {
			return err
		}

		return k.bk.BurnCoins(ctx, types.ModuleName, creationFee)
	}

	return nil
}
