package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/desmos-labs/desmos/v3/x/fees/types"
	"github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	// The reference to the ParamsStore to get and set params
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
func (k Keeper) CheckFees(ctx sdk.Context, fees sdk.Coins, msgs []sdk.Msg) error {
	feesParams := k.GetParams(ctx)

	// calculate required fees for the given messages
	requiredFees := sdk.NewCoins()
	for _, msg := range msgs {
		for _, minFee := range feesParams.MinFees {
			if proto.MessageName(msg) == minFee.MessageType {
				requiredFees = requiredFees.Add(minFee.Amount...)
			}
		}
	}

	if !requiredFees.IsZero() && (requiredFees.IsAnyGT(fees) || fees.IsZero()) {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFee,
			fmt.Sprintf("Expected at least %s, got %s", requiredFees, fees))
	}

	return nil
}
