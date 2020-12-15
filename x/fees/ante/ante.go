package ante

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/types"

	feeskeeper "github.com/desmos-labs/desmos/x/fees/keeper"
)

// NewAnteHandler returns a custom AnteHandler that, besides all the default checks
// (sequence number increment, signature and account number checks, fee deduction),
// makes sure that each transaction has a minimum fee based on the contained messages.
func NewAnteHandler(
	ak authante.AccountKeeper,
	bankKeeper types.BankKeeper,
	sigGasConsumer authante.SignatureVerificationGasConsumer,
	feesKeeper feeskeeper.Keeper,
	signModeHandler signing.SignModeHandler,
) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		authante.NewSetUpContextDecorator(),
		authante.NewMempoolFeeDecorator(),
		authante.NewValidateBasicDecorator(),
		authante.NewValidateMemoDecorator(ak),
		NewMinFeeDecorator(feesKeeper),
		authante.NewConsumeGasForTxSizeDecorator(ak),
		authante.NewSetPubKeyDecorator(ak),
		authante.NewValidateSigCountDecorator(ak),
		authante.NewDeductFeeDecorator(ak, bankKeeper),
		authante.NewSigGasConsumeDecorator(ak, sigGasConsumer),
		authante.NewSigVerificationDecorator(ak, signModeHandler),
		authante.NewIncrementSequenceDecorator(ak),
	)
}

type MinFeeDecorator struct {
	feesKeeper feeskeeper.Keeper
}

func NewMinFeeDecorator(feesKeeper feeskeeper.Keeper) MinFeeDecorator {
	return MinFeeDecorator{
		feesKeeper: feesKeeper,
	}
}

func (mfd MinFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	// all transactions must be of type feeTx
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		// Set a gas meter with limit 0 as to prevent an infinite gas meter attack
		// during runTx.
		newCtx = setGasMeter(simulate, ctx, 0)
		return newCtx, errors.New("tx must be FeeTx")
	}

	// skip block with height 0, otherwise no chain initialization could happen!
	if ctx.BlockHeight() == 0 {
		return next(ctx, tx, simulate)
	}

	// Check the minimum fees of the transaction
	err = mfd.feesKeeper.CheckFees(ctx, feeTx.GetFee(), feeTx.GetMsgs())
	if err != nil {
		return ctx, err
	}

	return next(ctx, tx, simulate)
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
