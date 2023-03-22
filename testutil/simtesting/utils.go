package simtesting

import (
	"math/rand"

	"github.com/desmos-labs/desmos/v4/app/params"
	feeskeeper "github.com/desmos-labs/desmos/v4/x/fees/keeper"
	"github.com/desmos-labs/desmos/v4/x/fees/types"

	"github.com/cosmos/cosmos-sdk/baseapp"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// SendMsg sends a transaction with the specified message.
func SendMsg(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
	msg interface {
		sdk.Msg
		Type() string
	}, ctx sdk.Context,
	simAccount simtypes.Account,
) (simulation.OperationInput, error) {
	addr := msg.GetSigners()[0]
	account := ak.GetAccount(ctx, addr)
	coins := bk.SpendableCoins(ctx, account.GetAddress())

	fees, sendTx, err := computeFees(r, ctx, fk, msg, coins)
	if err != nil {
		return simulation.OperationInput{}, err
	}

	if !sendTx {
		return simulation.OperationInput{}, nil
	}

	txGen := params.MakeEncodingConfig().TxConfig
	return simulation.OperationInput{
		R:               r,
		App:             app,
		TxGen:           txGen,
		Cdc:             nil,
		Msg:             msg,
		MsgType:         msg.Type(),
		Context:         ctx,
		SimAccount:      simAccount,
		AccountKeeper:   ak,
		Bankkeeper:      bk,
		ModuleName:      types.ModuleName,
		CoinsSpentInMsg: fees,
	}, nil
}

// computeFees computes the fees that should be used to send a transaction with the given message,
// considering the max spendable amount provided and the min fees set inside the fees module
func computeFees(
	r *rand.Rand, ctx sdk.Context, fk feeskeeper.Keeper, msg sdk.Msg, max sdk.Coins,
) (fees sdk.Coins, sendTx bool, err error) {
	minFees := fk.GetParams(ctx).MinFees
	for _, minFee := range minFees {
		if sdk.MsgTypeURL(msg) == minFee.MessageType {
			fees = minFee.Amount
			sendTx = fees.IsAllLT(max)
			return
		}
	}

	fees, err = simtypes.RandomFees(r, ctx, max)
	return
}
