package simtesting

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/desmos-labs/desmos/v7/x/posts/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// SendMsg sends a transaction with the specified message.
func SendMsg(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg interface {
		sdk.Msg
		Route() string
		Type() string
	}, ctx sdk.Context,
	simAccount simtypes.Account,
) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

	deposit := sdk.Coins{}
	spendableCoins := bk.SpendableCoins(ctx, simAccount.Address)
	for _, v := range spendableCoins {
		if bk.IsSendEnabledCoin(ctx, v) {
			deposit = deposit.Add(simtypes.RandSubsetCoins(r, sdk.NewCoins(v))...)
		}
	}

	if deposit.IsZero() {
		return simtypes.NoOpMsg(msg.Route(), msg.Type(), "skip because of broke account"), nil, nil
	}

	interfaceRegistry := codectypes.NewInterfaceRegistry()
	txConfig := tx.NewTxConfig(codec.NewProtoCodec(interfaceRegistry), tx.DefaultSignModes)
	txCtx := simulation.OperationInput{
		R:               r,
		App:             app,
		TxGen:           txConfig,
		Cdc:             nil,
		Msg:             msg,
		MsgType:         msg.Type(),
		Context:         ctx,
		SimAccount:      simAccount,
		AccountKeeper:   ak,
		Bankkeeper:      bk,
		ModuleName:      types.ModuleName,
		CoinsSpentInMsg: deposit,
	}
	return simulation.GenAndDeliverTxWithRandFees(txCtx)
}
