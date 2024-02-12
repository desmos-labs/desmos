package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v7/x/tokenfactory/types"
)

func (k Keeper) mintTo(ctx sdk.Context, amount sdk.Coin, mintTo string) error {
	// verify that denom is an x/tokenfactory denom
	_, _, err := types.DeconstructDenom(amount.Denom)
	if err != nil {
		return err
	}

	err = k.bk.MintCoins(ctx, types.ModuleName, sdk.NewCoins(amount))
	if err != nil {
		return err
	}

	addr := sdk.MustAccAddressFromBech32(mintTo)

	return k.bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(amount))
}

func (k Keeper) burnFrom(ctx sdk.Context, amount sdk.Coin, burnFrom string) error {
	// verify that denom is an x/tokenfactory denom
	_, _, err := types.DeconstructDenom(amount.Denom)
	if err != nil {
		return err
	}

	addr := sdk.MustAccAddressFromBech32(burnFrom)

	if err := k.bk.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sdk.NewCoins(amount)); err != nil {
		return err
	}

	return k.bk.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(amount))
}
