package ante

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	cosmosante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/desmos-labs/desmos/x/fees"
)

// NewAnteHandler returns a custom AnteHandler that besides all the default checks
//(sequence number increment, signature and account number checks, fee deduction) make sure that each
// transaction has a minimum fee based on messages types
func NewAnteHandler(
	ak keeper.AccountKeeper,
	supplyKeeper types.SupplyKeeper,
	sigGasConsumer cosmosante.SignatureVerificationGasConsumer,
	feesKeeper fees.Keeper,
) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		cosmosante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		cosmosante.NewMempoolFeeDecorator(),
		cosmosante.NewValidateBasicDecorator(),
		cosmosante.NewValidateMemoDecorator(ak),
		NewMinFeeDecorator(feesKeeper),
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
	feesKeeper fees.Keeper
}

func NewMinFeeDecorator(feesKeeper fees.Keeper) MinFeeDecorator {
	return MinFeeDecorator{
		feesKeeper: feesKeeper,
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

	// Check the minimum fees of the transaction
	if err := mfd.feesKeeper.CheckFees(ctx, stdTx.Fee, stdTx.Msgs); err != nil {
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
