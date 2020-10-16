package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/profiles/keeper"
	"github.com/desmos-labs/desmos/x/profiles/types"
	"github.com/tendermint/tendermint/crypto"
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

	randomDTag := RandomDTag(r)
	req := types.NewDTagTransferRequest(randomDTag, currentOwner.Address, receivingUser.Address)
	_ = k.SaveProfile(ctx, types.NewProfile(randomDTag, currentOwner.Address, ctx.BlockTime()))

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

	req := types.NewDTagTransferRequest("dtag", currentOwner.Address, receivingUser.Address)

	// skip if requests doesnt exists
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

	profile := NewRandomProfile(r, currentOwner.Address).WithDTag("dtag")
	err := keeper.ValidateProfile(ctx, k, profile)
	if err != nil {
		return sim.Account{}, types.DTagTransferRequest{}, "", true
	}

	err = k.SaveProfile(ctx, profile)

	if err != nil {
		return sim.Account{}, types.DTagTransferRequest{}, "", true
	}

	return currentOwner, req, RandomDTag(r), false
}

// SimulateMsgRefuseDTagTransfer tests and runs a single MsgRefuseDTagTransfer
// nolint: funlen
func SimulateMsgRefuseDTagTransfer(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {
		acc, sender, skip := randomRefuseDTagTransferFields(r, ctx, accs, k, ak)
		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgRefuseDTagTransferRequest(sender, acc.Address)

		err = sendMsgMsgRefuseDTagTransfer(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{acc.PrivKey})
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgMsgRefuseDTagTransfer sends a transaction with a MsgRefuseDTagTransfer from a provided random account.
func sendMsgMsgRefuseDTagTransfer(
	r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgRefuseDTagTransferRequest, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
) error {

	account := ak.GetAccount(ctx, msg.Owner)
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

// randomRefuseDTagTransferFields returns random refuse DTag transfer fields
func randomRefuseDTagTransferFields(r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper, _ auth.AccountKeeper,
) (sim.Account, sdk.AccAddress, bool) {
	if len(accs) == 0 {
		return sim.Account{}, sdk.AccAddress{}, true
	}

	// Get random accounts
	currentOwner, _ := sim.RandomAcc(r, accs)
	receivingUser, _ := sim.RandomAcc(r, accs)

	// skip if the two addresses are equals
	if currentOwner.Equals(receivingUser) {
		return sim.Account{}, sdk.AccAddress{}, true
	}

	req := types.NewDTagTransferRequest("dtag", currentOwner.Address, receivingUser.Address)
	err := k.SaveDTagTransferRequest(ctx, req)
	if err != nil {
		return sim.Account{}, sdk.AccAddress{}, true
	}

	return currentOwner, receivingUser.Address, false
}

// SimulateMsgCancelDTagTransfer tests and runs a single MsgCancelDTagTransfer
// nolint: funlen
func SimulateMsgCancelDTagTransfer(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {
		acc, owner, skip := randomCancelDTagTransferFields(r, ctx, accs, k, ak)
		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgCancelDTagRequest(acc.Address, owner)

		err = sendMsgMsgCancelDTagTransfer(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{acc.PrivKey})
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgMsgCancelDTagTransfer sends a transaction with a MsgCancelDTagTransfer from a provided random account.
func sendMsgMsgCancelDTagTransfer(
	r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgCancelDTagTransferRequest, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
) error {

	account := ak.GetAccount(ctx, msg.Sender)
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

// randomCancelDTagTransferFields returns random refuse DTag transfer fields
func randomCancelDTagTransferFields(r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper, _ auth.AccountKeeper,
) (sim.Account, sdk.AccAddress, bool) {
	if len(accs) == 0 {
		return sim.Account{}, sdk.AccAddress{}, true
	}

	// Get random accounts
	currentOwner, _ := sim.RandomAcc(r, accs)
	receivingUser, _ := sim.RandomAcc(r, accs)

	// skip if the two addresses are equals
	if currentOwner.Equals(receivingUser) {
		return sim.Account{}, sdk.AccAddress{}, true
	}

	req := types.NewDTagTransferRequest("dtag", currentOwner.Address, receivingUser.Address)
	err := k.SaveDTagTransferRequest(ctx, req)
	if err != nil {
		return sim.Account{}, sdk.AccAddress{}, true
	}

	return receivingUser, currentOwner.Address, false
}
