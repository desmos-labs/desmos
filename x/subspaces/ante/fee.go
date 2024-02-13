package ante

import (
	"fmt"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	antetypes "github.com/desmos-labs/desmos/v7/x/subspaces/ante/types"
	"github.com/desmos-labs/desmos/v7/x/subspaces/types"
)

var _ sdk.AnteDecorator = &DeductFeeDecorator{}

// DeductFeeDecorator represents the decorator used to deduct fee
type DeductFeeDecorator struct {
	authDeductFeeDecorator antetypes.AuthDeductFeeDecorator
	ak                     antetypes.AccountKeeper
	bk                     antetypes.BankKeeper
	sk                     antetypes.SubspacesKeeper

	txFeeChecker ante.TxFeeChecker
}

// NewDeductFeeDecorator returns a new DeductFeeDecorator instance
func NewDeductFeeDecorator(
	authDeductFeeDecorator antetypes.AuthDeductFeeDecorator,
	ak antetypes.AccountKeeper,
	bk antetypes.BankKeeper,
	sk antetypes.SubspacesKeeper,
	txFeeChecker ante.TxFeeChecker,
) DeductFeeDecorator {
	if txFeeChecker == nil {
		txFeeChecker = CheckTxFeeWithSubspaceMinPrices(ante.CheckTxFeeWithValidatorMinGasPrices, sk)
	}

	return DeductFeeDecorator{
		authDeductFeeDecorator: authDeductFeeDecorator,
		ak:                     ak,
		bk:                     bk,
		sk:                     sk,
		txFeeChecker:           txFeeChecker,
	}
}

// AnteHandle implements AnteDecorator
func (dfd DeductFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	subspaceID, isSubspaceTx := GetTxSubspaceID(tx)
	if !isSubspaceTx {
		return dfd.authDeductFeeDecorator.AnteHandle(ctx, tx, simulate, next)
	}

	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, errors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	if !simulate && ctx.BlockHeight() > 0 && feeTx.GetGas() == 0 {
		return ctx, errors.Wrap(sdkerrors.ErrInvalidGasLimit, "must provide positive gas")
	}

	fees := feeTx.GetFee()
	var priority int64
	if !simulate {
		fees, priority, err = dfd.txFeeChecker(ctx, tx)
		if err != nil {
			return ctx, err
		}
	}

	newCtx, success, err := dfd.tryHandleSubspaceTx(ctx, feeTx, subspaceID, fees)
	if err != nil {
		return newCtx, err
	}

	// Move to next ante if the process was successful
	if success {
		newCtx = newCtx.WithPriority(priority)
		return next(newCtx, tx, simulate)
	}

	// If the custom handling was not successful, fallback to the default handling
	return dfd.authDeductFeeDecorator.AnteHandle(ctx, tx, simulate, next)
}

// GetTxSubspaceID  returns the valid subspace id, returns false if it is invalid
func GetTxSubspaceID(tx sdk.Tx) (uint64, bool) {
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

// tryHandleSubspaceTx handles the fee deduction for a single-subspace transaction,
// and returns if the process succeeded or not
func (dfd DeductFeeDecorator) tryHandleSubspaceTx(ctx sdk.Context, tx sdk.FeeTx, subspaceID uint64, fees sdk.Coins) (newCtx sdk.Context, success bool, err error) {
	feePayer := tx.FeePayer()
	feeGranter := tx.FeeGranter()
	deductFeesFrom := feePayer

	// If the fee granter is not set, or it's not equal to the subspace treasury,
	// then use auth.DeductFeeDecorator to deal with fees
	if feeGranter == nil || !feeGranter.Equals(types.GetTreasuryAddress(subspaceID)) {
		return ctx, false, nil
	}

	if addr := dfd.ak.GetModuleAddress(authtypes.FeeCollectorName); addr == nil {
		return ctx, false, fmt.Errorf("fee collector module account (%s) has not been set", authtypes.FeeCollectorName)
	}

	// Try using the granted fee grant to deduct the fee. If we can't, it means no grant exists
	used := dfd.sk.UseGrantedFees(ctx, subspaceID, feePayer, fees, tx.GetMsgs())
	if !used {
		return ctx, false, nil
	}

	deductFeesFrom = feeGranter
	deductFeesFromAcc := dfd.ak.GetAccount(ctx, deductFeesFrom)
	if deductFeesFromAcc == nil {
		return ctx, false, errors.Wrapf(sdkerrors.ErrUnknownAddress, "fee payer address: %s does not exist", deductFeesFrom)
	}

	// Deduct the fees
	if !fees.IsZero() {
		err = ante.DeductFees(dfd.bk, ctx, deductFeesFromAcc, fees)
		if err != nil {
			return ctx, false, err
		}
	}

	// Emit the fee deduction events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeTx,
			sdk.NewAttribute(sdk.AttributeKeyFee, fees.String()),
			sdk.NewAttribute(sdk.AttributeKeyFeePayer, deductFeesFrom.String()),
		),
	})
	return ctx, true, err
}
