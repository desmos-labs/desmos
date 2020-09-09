package simulation

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/profiles/keeper"
	"github.com/desmos-labs/desmos/x/profiles/types"
	"github.com/tendermint/tendermint/crypto"
	"math/rand"
)

// SimulateMsgRequestDTagTransfer tests and runs a single MsgRequestDTagTransfer
// nolint: funlen
func SimulateMsgRequestDTagTransfer(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {
		acc, request, skip := randomDtagRequestTransferFields(r, ctx, accs, k, ak)
		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgRequestDTagTransfer(request.CurrentOwner, request.ReceivingUser)

		err = sendMsgRequestDTagTransfer(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{acc.PrivKey})
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgRequestDTagTransfer sends a transaction with a MsgRequestDTagTransfer from a provided random account.
func sendMsgRequestDTagTransfer(
	r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgRequestDTagTransfer, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
) error {

	account := ak.GetAccount(ctx, msg.ReceivingUser)
	coins := account.SpendableCoins(ctx.BlockTime())

	fees, err := sim.RandomFees(r, ctx, coins)
	if err != nil {
		return err
	}

	tx := helpers.GenTx(
		[]sdk.Msg{msg},
		fees,
		DefaultGasValue,
		chainID,
		[]uint64{account.GetAccountNumber()},
		[]uint64{account.GetSequence()},
		privkeys...,
	)

	_, _, err = app.Deliver(tx)
	if err != nil {
		return err
	}

	return nil
}

// randomDtagRequestTransferFields returns random dTagRequest data
func randomDtagRequestTransferFields(
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper, _ auth.AccountKeeper,
) (sim.Account, types.DTagTransferRequest, bool) {
	if len(accs) == 0 {
		return sim.Account{}, types.DTagTransferRequest{}, true
	}

	// Get random accounts
	currentOwner, _ := sim.RandomAcc(r, accs)
	receivingUser, _ := sim.RandomAcc(r, accs)

	// skip if the two addresses are equals
	if currentOwner.Equals(receivingUser) {
		return sim.Account{}, types.DTagTransferRequest{}, true
	}

	req := types.NewDTagTransferRequest(currentOwner.Address, receivingUser.Address)

	// skip if requests already exists
	requests := k.GetUserDTagTransferRequests(ctx, currentOwner.Address)
	for _, request := range requests {
		if request.Equals(req) {
			return sim.Account{}, types.DTagTransferRequest{}, true
		}
	}

	return receivingUser, req, false
}

// SimulateMsgAcceptDTagTransfer tests and runs a single MsgAcceptDTagTransfer
func SimulateMsgAcceptDTagTransfer(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {
		acc, request, dtag, skip := randomDtagAcceptRequestTransferFields(r, ctx, accs, k, ak)
		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgAcceptDTagTransfer(dtag, request.CurrentOwner, request.ReceivingUser)

		err = sendMsgMsgAcceptDTagTransfer(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{acc.PrivKey})
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgMsgAcceptDTagTransfer sends a transaction with a MsgAcceptDTagTransfer from a provided random account.
func sendMsgMsgAcceptDTagTransfer(
	r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgAcceptDTagTransfer, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
) error {

	account := ak.GetAccount(ctx, msg.CurrentOwner)
	coins := account.SpendableCoins(ctx.BlockTime())

	fees, err := sim.RandomFees(r, ctx, coins)
	if err != nil {
		return err
	}

	tx := helpers.GenTx(
		[]sdk.Msg{msg},
		fees,
		DefaultGasValue,
		chainID,
		[]uint64{account.GetAccountNumber()},
		[]uint64{account.GetSequence()},
		privkeys...,
	)

	_, _, err = app.Deliver(tx)
	if err != nil {
		return err
	}

	return nil
}

// randomDtagAcceptRequestTransferFields returns random dTagRequest data and a random dTag
func randomDtagAcceptRequestTransferFields(r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper, _ auth.AccountKeeper,
) (sim.Account, types.DTagTransferRequest, string, bool) {
	if len(accs) == 0 {
		return sim.Account{}, types.DTagTransferRequest{}, "", true
	}

	// Get random accounts
	currentOwner, _ := sim.RandomAcc(r, accs)
	receivingUser, _ := sim.RandomAcc(r, accs)

	// skip if the two addresses are equals
	if currentOwner.Equals(receivingUser) {
		return sim.Account{}, types.DTagTransferRequest{}, "", true
	}

	req := types.NewDTagTransferRequest(currentOwner.Address, receivingUser.Address)

	// skip if requests already exists
	requests := k.GetUserDTagTransferRequests(ctx, currentOwner.Address)
	found := false
	for _, request := range requests {
		if request.Equals(req) {
			found = true
			break
		}
	}

	if !found {
		return sim.Account{}, types.DTagTransferRequest{}, "", true
	}

	return currentOwner, req, RandomDTag(r), false
}
