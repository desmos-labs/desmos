package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/desmos-labs/desmos/v3/x/fees/types"
)

type Keeper struct {
	paramSubspace paramstypes.Subspace
	cdc           codec.BinaryCodec
}

func NewKeeper(cdc codec.BinaryCodec, paramSpace paramstypes.Subspace) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:           cdc,
		paramSubspace: paramSpace,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// SetParams sets module parameters
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSubspace.SetParamSet(ctx, &params)
}

// GetParams returns the module parameters
func (k Keeper) GetParams(ctx sdk.Context) (p types.Params) {
	k.paramSubspace.GetParamSet(ctx, &p)
	return p
}

// CheckFees checks whether the given fees are sufficient to pay for all the given messages.
// The check is performed considering the minimum fee amount specified inside the module parameters
// for each message and by summing them up.
func (k Keeper) CheckFees(ctx sdk.Context, msgs []sdk.Msg, fees sdk.Coins) error {
	feesParams := k.GetParams(ctx)

	// calculate required fees for the given messages
	requiredFees := sdk.NewCoins()
	for _, msg := range msgs {
		for _, minFee := range feesParams.MinFees {
			if sdk.MsgTypeURL(msg) == minFee.MessageType {
				requiredFees = requiredFees.Add(minFee.Amount...)
			}
		}
	}

	for _, coin := range requiredFees {
		amt := fees.AmountOf(coin.Denom)
		if coin.Amount.GT(amt) {
			return sdkerrors.Wrap(sdkerrors.ErrInsufficientFee,
				fmt.Sprintf("expected at least %s, got %s", requiredFees, fees))
		}
	}

	return nil
}
