package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MinFeeDecorator struct {
	feesKeeper Keeper
}

func NewMinFeeDecorator(feesKeeper Keeper) MinFeeDecorator {
	return MinFeeDecorator{
		feesKeeper: feesKeeper,
	}
}

func (mfd MinFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	// all transactions must be of type feeTx
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return newCtx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	// skip block with height 0, otherwise no chain initialization could happen!
	//if ctx.BlockHeight() == 0 {
	//	return next(ctx, tx, simulate)
	//}

	// Check the minimum fees of the transaction
	err = mfd.feesKeeper.CheckFees(ctx, feeTx.GetFee(), feeTx.GetMsgs())
	if err != nil {
		return ctx, err
	}

	return next(ctx, tx, simulate)
}
