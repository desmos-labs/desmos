package ante

import (
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	cosmosante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// fixedRequiredFee is the amount of fee desmos requires to process each tx. Actually set to 0.01 daric
var fixedRequiredFee = sdk.NewDecWithPrec(1, 2)

// NewAnteHandler returns a custom AnteHandler that besides all the default checks
//(sequence number increment, signature and account number checks, fee deduction) make sure that each
// transaction has a minimum fee of 0.01 daric
func NewAnteHandler(
	ak keeper.AccountKeeper,
	supplyKeeper types.SupplyKeeper,
	sigGasConsumer cosmosante.SignatureVerificationGasConsumer,
	feeTokenDenom string,
) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		cosmosante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		cosmosante.NewMempoolFeeDecorator(),
		cosmosante.NewValidateBasicDecorator(),
		cosmosante.NewValidateMemoDecorator(ak),
		NewMinFeeDecorator(feeTokenDenom),
		cosmosante.NewConsumeGasForTxSizeDecorator(ak),
		cosmosante.NewSetPubKeyDecorator(ak), // SetPubKeyDecorator must be called before all signature verification decorators
		cosmosante.NewValidateSigCountDecorator(ak),
		cosmosante.NewDeductFeeDecorator(ak, supplyKeeper),
		cosmosante.NewSigGasConsumeDecorator(ak, sigGasConsumer),
		cosmosante.NewSigVerificationDecorator(ak),
		cosmosante.NewIncrementSequenceDecorator(ak),
	)
}

type MinFeeDecorator struct {
	tokenDenom string
}

func NewMinFeeDecorator(tokenDenom string) MinFeeDecorator {
	return MinFeeDecorator{
		tokenDenom: tokenDenom,
	}
}

func (mfd MinFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	// all transactions must be of type auth.StdTx
	stdTx, ok := tx.(types.StdTx)
	if !ok {
		// Set a gas meter with limit 0 as to prevent an infinite gas meter attack
		// during runTx.
		newCtx = setGasMeter(simulate, ctx, 0)
		return newCtx, errors.New("tx must be StdTx")
	}

	// skip block with height 0, otherwise no chain initialization could happen!
	if ctx.BlockHeight() == 0 {
		return next(ctx, tx, simulate)
	}

	// calculate required fees for this transaction as (number of messages * fixed required fees)
	requiredFees := fixedRequiredFee.MulInt64(int64(len(stdTx.Msgs)))

	// Check the minimum fees
	if err := checkMinimumFees(stdTx, mfd.tokenDenom, requiredFees); err != nil {
		return ctx, err
	}

	return next(ctx, tx, simulate)
}

func checkMinimumFees(
	stdTx types.StdTx,
	tokenDenom string,
	requiredFees sdk.Dec,
) error {
	// Each message should cost 0.01 daric
	// Token quantity is always set as millionth of units
	stableRequiredQty := requiredFees.MulInt64(10000)
	feeAmount := sdk.NewDecFromInt(stdTx.Fee.Amount.AmountOf(tokenDenom))
	if stableRequiredQty.IsZero() || stableRequiredQty.LTE(feeAmount) {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFee,
			fmt.Sprintf("Insufficient fees. Expected %s %s amount, got %s", requiredFees, tokenDenom, feeAmount))
	}

	return nil
}

// setGasMeter returns a new context with a gas meter set from a given context.
func setGasMeter(simulate bool, ctx sdk.Context, gasLimit uint64) sdk.Context {
	// In various cases such as simulation and during the genesis block, we do not
	// meter any gas utilization.
	if simulate || ctx.BlockHeight() == 0 {
		return ctx.WithGasMeter(sdk.NewInfiniteGasMeter())
	}

	return ctx.WithGasMeter(sdk.NewGasMeter(gasLimit))
}
