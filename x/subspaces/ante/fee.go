package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/desmos-labs/desmos/v4/x/subspaces/keeper"

	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

type DeductFeeDecorator struct {
	authDeductFeeDecorator ante.DeductFeeDecorator
	subspaceKeeper         keeper.Keeper
}

func NewDeductFeeDecorator(authDeductFeeDecorator ante.DeductFeeDecorator, subspaceKeeper keeper.Keeper) DeductFeeDecorator {
	return DeductFeeDecorator{
		authDeductFeeDecorator,
		subspaceKeeper,
	}
}

func (dfd DeductFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}
	if isValidSubspaseMsgs(feeTx.GetMsgs()) {
		return dfd.anteHandle(ctx, tx, simulate, next)
	}
	return dfd.authDeductFeeDecorator.AnteHandle(ctx, tx, simulate, next)
}

func isValidSubspaseMsgs(msgs []sdk.Msg) bool {
	subspaceId := uint64(0)
	for _, msg := range msgs {
		subsbaseMsg, ok := msg.(types.SubspaceMsg)
		if !ok {
			return false
		}
		if subspaceId == 0 {
			subspaceId = subsbaseMsg.GetSubspaceID()
		} else if subspaceId == subsbaseMsg.GetSubspaceID() {
			return false
		}
	}
	return true
}

func (dfd DeductFeeDecorator) anteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	return next(ctx, tx, simulate)
}
