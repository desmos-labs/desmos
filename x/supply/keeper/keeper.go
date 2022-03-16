package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	distributionkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	"github.com/desmos-labs/desmos/v2/x/supply/types"
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

// GetConvertedTotalSupply returns the total supply converted from millionth amount to normal amount
func (k Keeper) GetConvertedTotalSupply(ctx sdk.Context, coinDenom string) sdk.Int {
	totalSupply := k.bk.GetSupply(ctx, coinDenom)
	return totalSupply.Amount.QuoRaw(6)
}

// CalculateCirculatingSupply calculates the current circulating supply by:
// 1. Getting the total supply
// 2. Subtract the community pool amount from it
// 3. Subtract the vesting accounts locked tokens from it
func (k Keeper) CalculateCirculatingSupply(ctx sdk.Context, coinDenom string) sdk.Coin {
	var circulatingSupply sdk.Int

	// Get total supply
	totalSupply := k.bk.GetSupply(ctx, coinDenom)

	// Get the community pool denom amount
	communityPoolDenomAmount := k.dk.GetFeePoolCommunityCoins(ctx).AmountOf(coinDenom)

	// Subtract community pool amount from the total supply
	circulatingSupply = totalSupply.Amount.Sub(communityPoolDenomAmount.RoundInt())

	// Subtract all vesting account locked tokens amount from the circulating supply
	accounts := k.ak.GetAllAccounts(ctx)
	for _, account := range accounts {
		if vestingAcc, ok := account.(exported.VestingAccount); ok {
			circulatingSupply = k.subtractVestingAccountDenomAmounts(circulatingSupply, vestingAcc, coinDenom)
		}
	}

	return sdk.NewCoin(coinDenom, circulatingSupply)
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
