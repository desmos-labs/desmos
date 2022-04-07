package simtesting

import (
	"math/rand"

	feeskeeper "github.com/desmos-labs/desmos/v3/x/fees/keeper"

	"github.com/cosmos/cosmos-sdk/baseapp"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

// SendMsg sends a transaction with the specified message.
func SendMsg(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
	msg sdk.Msg, ctx sdk.Context, chainID string, gasValue uint64, privkeys []cryptotypes.PrivKey,
) error {
	addr := msg.GetSigners()[0]
	account := ak.GetAccount(ctx, addr)
	coins := bk.SpendableCoins(ctx, account.GetAddress())

	fees, sendTx, err := computeFees(r, ctx, fk, msg, coins)
	if err != nil {
		return err
	}

	if !sendTx {
		return nil
	}

	txGen := simappparams.MakeTestEncodingConfig().TxConfig
	tx, err := helpers.GenTx(
		txGen,
		[]sdk.Msg{msg},
		fees,
		gasValue,
		chainID,
		[]uint64{account.GetAccountNumber()},
		[]uint64{account.GetSequence()},
		privkeys...,
	)
	if err != nil {
		return err
	}

	_, _, err = app.Deliver(txGen.TxEncoder(), tx)
	if err != nil {
		return err
	}

	return nil
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
