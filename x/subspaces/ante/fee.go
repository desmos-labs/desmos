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
	authDeductFeeDecorator AuthDeductFeeDecorator
	ak                     AccountKeeper
	bk                     BankKeeper
	sk                     SubspacesKeeper
}

// NewDeductFeeDecorator returns a new DeductFeeDecorator instance
func NewDeductFeeDecorator(authDeductFeeDecorator AuthDeductFeeDecorator, ak AccountKeeper, bk BankKeeper, sk SubspacesKeeper) DeductFeeDecorator {
	return DeductFeeDecorator{
		authDeductFeeDecorator: authDeductFeeDecorator,
		ak:                     ak,
		bk:                     bk,
		sk:                     sk,
	}
}

// AnteHandle implements AnteDecorator
func (dfd DeductFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	if id, ok := isValidSubspaceTx(tx); ok {
		feeTx, ok := tx.(sdk.FeeTx)
		if !ok {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
		}

		newCtx, success, err := dfd.tryHandleSubspaceTx(ctx, feeTx, id)
		if err != nil {
			return newCtx, err
		}
		// move to next ante if tryHandleSubspaceTx is success, or using auth.DeductFeeDecorator instead
		if success {
			return next(newCtx, tx, simulate)
		}
	}
	return dfd.authDeductFeeDecorator.AnteHandle(ctx, tx, simulate, next)
}

// isValidSubspaceTx returns the valid subspace id, returns false if it is invalid
func isValidSubspaceTx(tx sdk.Tx) (uint64, bool) {
	subspaceID := uint64(0)
	for _, msg := range tx.GetMsgs() {
		if subspaceMsg, ok := msg.(types.SubspaceMsg); ok {
			if subspaceID == 0 {
				subspaceID = subspaceMsg.GetSubspaceID()
			}

			if subspaceMsg.GetSubspaceID() == subspaceID {
				continue
			}
		}

		return 0, false
	}
	return subspaceID, true
}

// tryHandleSubspaceTx handles the fee deduction for subspace transaction, returns false if the process is failed
func (dfd DeductFeeDecorator) tryHandleSubspaceTx(ctx sdk.Context, tx sdk.FeeTx, subspaceID uint64) (newCtx sdk.Context, success bool, err error) {
	fees := tx.GetFee()
	feePayer := tx.FeePayer()
	feeGranter := tx.FeeGranter()
	deductFeesFrom := feePayer

	// if feegranter is not set set or fee granter equals to payer, then use auth.DeductFeeDecorator to deal with fees
	if feeGranter == nil || !feeGranter.Equals(types.GetTreasuryAddress(subspaceID)) {
		return ctx, false, nil
	}

	if addr := dfd.ak.GetModuleAddress(authtypes.FeeCollectorName); addr == nil {
		return ctx, false, fmt.Errorf("fee collector module account (%s) has not been set", authtypes.FeeCollectorName)
	}

	used := dfd.sk.UseGrantedFees(ctx, subspaceID, feePayer, fees, tx.GetMsgs())
	if !used {
		return ctx, false, nil
	}

	deductFeesFrom = feeGranter
	deductFeesFromAcc := dfd.ak.GetAccount(ctx, deductFeesFrom)
	if deductFeesFromAcc == nil {
		return ctx, false, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "fee payer address: %s does not exist", deductFeesFrom)
	}

	// deduct the fees
	if !fees.IsZero() {
		err = ante.DeductFees(dfd.bk, ctx, deductFeesFromAcc, fees)
		if err != nil {
			return ctx, false, err
		}
	}

	events := sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeTx,
			sdk.NewAttribute(sdk.AttributeKeyFee, fees.String()),
			sdk.NewAttribute(sdk.AttributeKeyFeePayer, deductFeesFrom.String()),
		),
	}
	ctx.EventManager().EmitEvents(events)
	return ctx, true, err
}
