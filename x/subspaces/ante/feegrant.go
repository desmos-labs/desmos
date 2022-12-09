package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"

	"github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

func NewDeductFeeDecorator(ak AccountKeeper, bk BankKeeper, fk FeegrantKeeper, sk keeper.Keeper) DeductFeeDecorator {
	return DeductFeeDecorator{
		authDeductAnte: ante.NewDeductFeeDecorator(ak, bk, fk),
		ak:             ak,
		bk:             bk,
		sk:             sk,
	}
}

func (dfd DeductFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	if id, ok := isValidSubspaceMsgs(tx.GetMsgs()); ok {
		newCtx, success, err := dfd.anteHandle(ctx, feeTx, simulate, next, id)
		if success {
			return newCtx, err
		}
	}
	return dfd.authDeductAnte.AnteHandle(ctx, tx, simulate, next)
}

func isValidSubspaceMsgs(msgs []sdk.Msg) (uint64, bool) {
	subspaceId := uint64(0)
	for _, msg := range msgs {
		subspaceMsg, ok := msg.(types.SubspaceMsg)
		if !ok {
			return 0, false
		}
		if subspaceId == 0 {
			subspaceId = subspaceMsg.GetSubspaceID()
		} else if subspaceId == subspaceMsg.GetSubspaceID() {
			return 0, false
		}
	}
	return subspaceId, true
}

func (dfd DeductFeeDecorator) anteHandle(ctx sdk.Context, tx sdk.FeeTx, simulate bool, next sdk.AnteHandler, subspaceID uint64) (newCtx sdk.Context, used bool, err error) {
	fee := tx.GetFee()
	feePayer := tx.FeePayer()
	feeGranter := tx.FeeGranter()
	deductFeesFrom := feePayer

	// if feegranter set deduct fee from feegranter account.
	// this works with only when feegrant enabled.
	if feeGranter != nil {
		if !feeGranter.Equals(feePayer) {
			used = dfd.sk.UseGrantedFees(ctx, subspaceID, feeGranter, feePayer, fee, tx.GetMsgs())
			if !used {
				return ctx, false, sdkerrors.Wrapf(err, "%s not allowed to pay fees from %s", feeGranter, feePayer)
			}
		}

		deductFeesFrom = feeGranter
	}

	deductFeesFromAcc := dfd.ak.GetAccount(ctx, deductFeesFrom)
	if deductFeesFromAcc == nil {
		return ctx, used, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "fee payer address: %s does not exist", deductFeesFrom)
	}

	// deduct the fees
	if !tx.GetFee().IsZero() {
		err = ante.DeductFees(dfd.bk, ctx, deductFeesFromAcc, tx.GetFee())
		if err != nil {
			return ctx, used, err
		}
	}

	events := sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeTx,
			sdk.NewAttribute(sdk.AttributeKeyFee, fee.String()),
			sdk.NewAttribute(sdk.AttributeKeyFeePayer, deductFeesFrom.String()),
		),
	}
	ctx.EventManager().EmitEvents(events)
	newCtx, err = next(ctx, tx, simulate)
	return newCtx, used, err
}
