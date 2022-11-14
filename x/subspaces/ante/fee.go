package ante

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	accountkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"

	"github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

type DeductFeeDecorator struct {
	adfd ante.DeductFeeDecorator
	ak   accountkeeper.AccountKeeper
	bk   bankkeeper.Keeper
	sk   keeper.Keeper
}

func NewDeductFeeDecorator(ak accountkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feegrantkeeper.Keeper, sk keeper.Keeper) DeductFeeDecorator {
	return DeductFeeDecorator{
		adfd: ante.NewDeductFeeDecorator(ak, bk, fk),
		ak:   ak,
		bk:   bk,
		sk:   sk,
	}
}

func (dfd DeductFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	if id, ok := isValidSubspaseMsgs(tx.GetMsgs()); ok {
		return dfd.anteHandle(ctx, tx, simulate, next, id)
	}
	return dfd.adfd.AnteHandle(ctx, tx, simulate, next)
}

func isValidSubspaseMsgs(msgs []sdk.Msg) (uint64, bool) {
	subspaceId := uint64(0)
	for _, msg := range msgs {
		subsbaseMsg, ok := msg.(types.SubspaceMsg)
		if !ok {
			return 0, false
		}
		if subspaceId == 0 {
			subspaceId = subsbaseMsg.GetSubspaceID()
		} else if subspaceId == subsbaseMsg.GetSubspaceID() {
			return 0, false
		}
	}
	return subspaceId, true
}

func (dfd DeductFeeDecorator) anteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler, subspaceId uint64) (newCtx sdk.Context, err error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}
	if addr := dfd.ak.GetModuleAddress(authtypes.FeeCollectorName); addr == nil {
		return ctx, fmt.Errorf("fee collector module account (%s) has not been set", authtypes.FeeCollectorName)
	}

	fee := feeTx.GetFee()
	feePayer := feeTx.FeePayer()
	feeGranter := feeTx.FeeGranter()
	deductFeesFrom := feePayer

	// if feegranter set deduct fee from feegranter account.
	// this works with only when feegrant enabled.
	if feeGranter != nil {
		if !feeGranter.Equals(feePayer) {
			err := dfd.sk.UseGrantedFees(ctx, feeGranter, feePayer, subspaceId, fee, tx.GetMsgs())
			if err != nil {
				return ctx, sdkerrors.Wrapf(err, "%s not allowed to pay fees from %s", feeGranter, feePayer)
			}
		}

		deductFeesFrom = feeGranter
	}

	deductFeesFromAcc := dfd.ak.GetAccount(ctx, deductFeesFrom)
	if deductFeesFromAcc == nil {
		return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "fee payer address: %s does not exist", deductFeesFrom)
	}

	// deduct the fees
	if !feeTx.GetFee().IsZero() {
		err = ante.DeductFees(dfd.bk, ctx, deductFeesFromAcc, feeTx.GetFee())
		if err != nil {
			return ctx, err
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
	return next(ctx, tx, simulate)
}
