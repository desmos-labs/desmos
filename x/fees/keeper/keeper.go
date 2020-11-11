package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/desmos-labs/desmos/x/fees/types"
)

type Keeper struct {
	// The reference to the ParamsStore to get and set params
	paramSubspace params.Subspace
	Cdc           *codec.Codec
}

func NewKeeper(cdc *codec.Codec, paramSpace params.Subspace) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		Cdc:           cdc,
		paramSubspace: paramSpace,
	}
}

// SetParams sets params on the store
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSubspace.SetParamSet(ctx, &params)
}

// GetParams returns the params from the store
func (k Keeper) GetParams(ctx sdk.Context) (p types.Params) {
	k.paramSubspace.GetParamSet(ctx, &p)
	return p
}

// CheckFees checks whether the given fees are sufficient to pay for all the given messages.
// The check is performed considering the minimum fee amounts specified inside the module parameters.
func (k Keeper) CheckFees(ctx sdk.Context, fees authtypes.StdFee, msgs []sdk.Msg) error {
	feesParams := k.GetParams(ctx)

	// calculate required fees for the given messages
	requiredFees := sdk.NewCoins()
	for _, msg := range msgs {
		for _, minFee := range feesParams.MinFees {
			if msg.Type() == minFee.MessageType {
				requiredFees = requiredFees.Add(minFee.Amount...)
			}
		}
	}

	if !requiredFees.IsZero() && (requiredFees.IsAnyGT(fees.Amount) || fees.Amount.IsZero()) {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFee,
			fmt.Sprintf("Expected at least %s, got %s", requiredFees, fees.Amount))
	}

	return nil
}
