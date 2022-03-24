package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	distributionkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	"github.com/desmos-labs/desmos/v3/x/supply/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	cdc codec.BinaryCodec
	ak  authkeeper.AccountKeeper
	bk  bankkeeper.Keeper
	dk  distributionkeeper.Keeper
}

// NewKeeper creates new instances of the supply keeper
func NewKeeper(cdc codec.BinaryCodec,
	ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, dk distributionkeeper.Keeper) Keeper {
	return Keeper{
		cdc: cdc,
		ak:  ak,
		bk:  bk,
		dk:  dk,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// GetConvertedTotalSupply returns the total supply converted by dividing it for the given divider factor
func (k Keeper) GetConvertedTotalSupply(ctx sdk.Context, coinDenom string, divider sdk.Int) sdk.Int {
	totalSupply := k.bk.GetSupply(ctx, coinDenom)
	return totalSupply.Amount.Quo(divider)
}

// CalculateCirculatingSupply calculates the current circulating supply by:
// 1. Getting the total supply
// 2. Subtract the community pool amount from it
// 3. Subtract the vesting accounts locked tokens from it
// It then returns it converted by diving it with the divider factor
func (k Keeper) CalculateCirculatingSupply(ctx sdk.Context, coinDenom string, divider sdk.Int) sdk.Int {
	var circulatingSupply sdk.Int

	// Get total supply
	totalSupply := k.bk.GetSupply(ctx, coinDenom)

	// Get the community pool denom amount
	communityPoolDenomAmount := k.dk.GetFeePoolCommunityCoins(ctx).AmountOf(coinDenom)

	// Subtract community pool amount from the total supply
	circulatingSupply = totalSupply.Amount.Sub(communityPoolDenomAmount.RoundInt())

	// Subtract all vesting account locked tokens amount from the circulating supply
	k.ak.IterateAccounts(ctx, func(account authtypes.AccountI) (stop bool) {
		if vestingAcc, ok := account.(exported.VestingAccount); ok {
			circulatingSupply = k.subtractVestingAccountDenomAmounts(circulatingSupply, vestingAcc, coinDenom)
		}
		return false
	})

	// Convert the circulating supply with the divider factor
	convertedCirculatingSupply := circulatingSupply.Quo(divider)

	return convertedCirculatingSupply
}

// subtractVestingAccountDenomAmounts subtract the given vesting account denom amount from the
// circulating supply
func (k Keeper) subtractVestingAccountDenomAmounts(circulatingSupply sdk.Int,
	vestingAccount exported.VestingAccount, denom string) sdk.Int {
	originalVesting := vestingAccount.GetOriginalVesting()
	delegatedFree := vestingAccount.GetDelegatedFree()

	originalVestingAmount := originalVesting.AmountOf(denom)
	delegatedFreeAmount := delegatedFree.AmountOf(denom)

	return circulatingSupply.Sub(originalVestingAmount).Sub(delegatedFreeAmount)
}
