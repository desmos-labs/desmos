package ante

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

var _ sdk.AnteDecorator = &DeductFeeDecorator{}

// DeductFeeDecorator represents the decorator used to deduct fee
type DeductFeeDecorator struct {
	authDeductAnte ante.DeductFeeDecorator
	ak             AccountKeeper
	bk             BankKeeper
	sk             SubspacesKeeper
}

// NewDeductFeeDecorator returns a new DeductFeeDecorator instance
func NewDeductFeeDecorator(ak AccountKeeper, bk BankKeeper, fk FeegrantKeeper, sk SubspacesKeeper) DeductFeeDecorator {
	return DeductFeeDecorator{
		authDeductAnte: ante.NewDeductFeeDecorator(ak, bk, fk),
		ak:             ak,
		bk:             bk,
		sk:             sk,
	}
}

// AnteHandle implements AnteDecorator
func (dfd DeductFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	if id, ok := isValidSubspaceTx(tx); ok {
		feeTx, ok := tx.(sdk.FeeTx)
		if !ok {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
		}

		newCtx, success, err := dfd.handleSubspaceTx(ctx, feeTx, id)
		if err != nil {
			return newCtx, err
		}
		if success {
			return next(newCtx, tx, simulate)
		}
	}
	return dfd.authDeductAnte.AnteHandle(ctx, tx, simulate, next)
}

// isValidSubspaceTx returns the valid subspace id, returns false if it is invalid
func isValidSubspaceTx(tx sdk.Tx) (uint64, bool) {
	subspaceId := uint64(0)
	for _, msg := range tx.GetMsgs() {
		subspaceMsg, ok := msg.(types.SubspaceMsg)
		if !ok {
			return 0, false
		}
		if subspaceId == 0 {
			subspaceId = subspaceMsg.GetSubspaceID()
		} else if subspaceId != subspaceMsg.GetSubspaceID() {
			return 0, false
		}
	}
	return subspaceId, true
}

// handleSubspaceTx handles the fee deduction for subspace transaction, return false if the process is failed
func (dfd DeductFeeDecorator) handleSubspaceTx(ctx sdk.Context, tx sdk.FeeTx, subspaceID uint64) (newCtx sdk.Context, success bool, err error) {
	fee := tx.GetFee()
	feePayer := tx.FeePayer()
	feeGranter := tx.FeeGranter()
	deductFeesFrom := feePayer

	// if feegranter does not set or fee granter equals to payer, then use auth.DeductFeeDecorator to deal with fees
	if feeGranter == nil || feeGranter.Equals(feePayer) {
		return ctx, false, nil
	}

	if addr := dfd.ak.GetModuleAddress(authtypes.FeeCollectorName); addr == nil {
		return ctx, false, fmt.Errorf("fee collector module account (%s) has not been set", authtypes.FeeCollectorName)
	}

	used := dfd.sk.UseGrantedFees(ctx, subspaceID, feeGranter, feePayer, fee, tx.GetMsgs())
	if !used {
		return ctx, false, nil
	}
	deductFeesFrom = feeGranter

	deductFeesFromAcc := dfd.ak.GetAccount(ctx, deductFeesFrom)
	if deductFeesFromAcc == nil {
		return ctx, false, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "fee payer address: %s does not exist", deductFeesFrom)
	}

	// deduct the fees
	if !fee.IsZero() {
		err = ante.DeductFees(dfd.bk, ctx, deductFeesFromAcc, fee)
		if err != nil {
			return ctx, false, err
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
	return ctx, true, err
}
