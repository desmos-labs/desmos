package simulation

// DONTCOVER

import (
	"math/rand"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/keeper"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

// SimulateMsgRequestDTagTransfer tests and runs a single MsgRequestDTagTransfer
// nolint: funlen
func SimulateMsgRequestDTagTransfer(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {
		sender, request, skip := randomDtagRequestTransferFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, ""), nil, nil
		}

		msg := types.NewMsgRequestDTagTransfer(request.Sender, request.Receiver)

		err = sendMsgRequestDTagTransfer(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{sender.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgRequestDTagTransfer"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgRequestDTagTransfer sends a transaction with a MsgRequestDTagTransfer from a provided random account.
func sendMsgRequestDTagTransfer(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgRequestDTagTransfer, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
) error {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	account := ak.GetAccount(ctx, addr)
	coins := bk.SpendableCoins(ctx, account.GetAddress())

	fees, err := simtypes.RandomFees(r, ctx, coins)
	if err != nil {
		return err
	}

	txGen := simappparams.MakeTestEncodingConfig().TxConfig
	tx, err := helpers.GenTx(
		txGen,
		[]sdk.Msg{msg},
		fees,
		DefaultGasValue,
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

// randomDtagRequestTransferFields returns random dTagRequest data
func randomDtagRequestTransferFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (simtypes.Account, types.DTagTransferRequest, bool) {
	if len(accs) == 0 {
		return simtypes.Account{}, types.DTagTransferRequest{}, true
	}

	// Get random accounts
	receiver, _ := simtypes.RandomAcc(r, accs)
	sender, _ := simtypes.RandomAcc(r, accs)

	// skip if the two addresses are equals
	if receiver.Equals(sender) {
		return simtypes.Account{}, types.DTagTransferRequest{}, true
	}

	if k.IsUserBlocked(ctx, receiver.Address.String(), sender.Address.String()) {
		return simtypes.Account{}, types.DTagTransferRequest{}, true
	}

	randomDTag := RandomDTag(r)
	req := types.NewDTagTransferRequest(randomDTag, sender.Address.String(), receiver.Address.String())
	_ = k.StoreProfile(ctx, types.NewProfile(
		randomDTag,
		"",
		"",
		types.NewPictures("", ""),
		ctx.BlockTime(),
		receiver.Address.String(),
	))

	// skip if requests already exists
	requests := k.GetUserIncomingDTagTransferRequests(ctx, receiver.Address.String())
	for _, request := range requests {
		if request.Sender == req.Sender {
			return simtypes.Account{}, types.DTagTransferRequest{}, true
		}
	}

	return sender, req, false
}

// ___________________________________________________________________________________________________________________

// SimulateMsgAcceptDTagTransfer tests and runs a single MsgAcceptDTagTransfer
func SimulateMsgAcceptDTagTransfer(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {
		acc, request, dtag, skip := randomDtagAcceptRequestTransferFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, ""), nil, nil
		}

		msg := types.NewMsgAcceptDTagTransfer(
			dtag,
			request.Sender,
			request.Receiver,
		)

		err = sendMsgMsgAcceptDTagTransfer(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{acc.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgAcceptDTagTransfer"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgMsgAcceptDTagTransfer sends a transaction with a MsgAcceptDTagTransfer from a provided random account.
func sendMsgMsgAcceptDTagTransfer(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgAcceptDTagTransfer, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
) error {
	addr, _ := sdk.AccAddressFromBech32(msg.Receiver)
	account := ak.GetAccount(ctx, addr)
	coins := bk.SpendableCoins(ctx, account.GetAddress())

	fees, err := simtypes.RandomFees(r, ctx, coins)
	if err != nil {
		return err
	}

	txGen := simappparams.MakeTestEncodingConfig().TxConfig
	tx, err := helpers.GenTx(
		txGen,
		[]sdk.Msg{msg},
		fees,
		DefaultGasValue,
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

// randomDtagAcceptRequestTransferFields returns random dTagRequest data and a random dTag
func randomDtagAcceptRequestTransferFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (simtypes.Account, types.DTagTransferRequest, string, bool) {
	if len(accs) == 0 {
		return simtypes.Account{}, types.DTagTransferRequest{}, "", true
	}

	// Get random accounts
	sender, _ := simtypes.RandomAcc(r, accs)
	receiver, _ := simtypes.RandomAcc(r, accs)

	// skip if the two addresses are equals
	if sender.Equals(receiver) {
		return simtypes.Account{}, types.DTagTransferRequest{}, "", true
	}

	req := types.NewDTagTransferRequest(
		"dtag",
		receiver.Address.String(),
		sender.Address.String(),
	)

	// skip if requests doesnt exists
	requests := k.GetUserIncomingDTagTransferRequests(ctx, sender.Address.String())
	found := false
	for _, request := range requests {
		if request.Equal(req) {
			found = true
			break
		}
	}

	if !found {
		return simtypes.Account{}, types.DTagTransferRequest{}, "", true
	}

	profile := NewRandomProfile(r, sender.Address)
	profile.Dtag = "dtag"

	err := k.ValidateProfile(ctx, profile)
	if err != nil {
		return simtypes.Account{}, types.DTagTransferRequest{}, "", true
	}

	err = k.StoreProfile(ctx, profile)

	if err != nil {
		return simtypes.Account{}, types.DTagTransferRequest{}, "", true
	}

	return receiver, req, RandomDTag(r), false
}

// ___________________________________________________________________________________________________________________

// SimulateMsgRefuseDTagTransfer tests and runs a single MsgRefuseDTagTransfer
// nolint: funlen
func SimulateMsgRefuseDTagTransfer(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {
		sender, receiver, skip := randomRefuseDTagTransferFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, ""), nil, nil
		}

		msg := types.NewMsgRefuseDTagTransferRequest(
			sender.Address.String(),
			receiver.Address.String(),
		)

		err = sendMsgMsgRefuseDTagTransfer(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{receiver.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgRefuseDTagTransfer"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgMsgRefuseDTagTransfer sends a transaction with a MsgRefuseDTagTransfer from a provided random account.
func sendMsgMsgRefuseDTagTransfer(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgRefuseDTagTransfer, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
) error {
	addr, _ := sdk.AccAddressFromBech32(msg.Receiver)
	account := ak.GetAccount(ctx, addr)
	coins := bk.SpendableCoins(ctx, account.GetAddress())

	fees, err := simtypes.RandomFees(r, ctx, coins)
	if err != nil {
		return err
	}

	txGen := simappparams.MakeTestEncodingConfig().TxConfig
	tx, err := helpers.GenTx(
		txGen,
		[]sdk.Msg{msg},
		fees,
		DefaultGasValue,
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

// randomRefuseDTagTransferFields returns random refuse DTag transfer fields
func randomRefuseDTagTransferFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (simtypes.Account, simtypes.Account, bool) {
	if len(accs) == 0 {
		return simtypes.Account{}, simtypes.Account{}, true
	}

	// Get random accounts
	sender, _ := simtypes.RandomAcc(r, accs)
	receiver, _ := simtypes.RandomAcc(r, accs)

	// skip if the two addresses are equals
	if sender.Equals(receiver) {
		return simtypes.Account{}, simtypes.Account{}, true
	}

	req := types.NewDTagTransferRequest(
		"dtag",
		sender.Address.String(),
		receiver.Address.String(),
	)
	err := k.SaveDTagTransferRequest(ctx, req)
	if err != nil {
		return simtypes.Account{}, simtypes.Account{}, true
	}

	return sender, receiver, false
}

// ___________________________________________________________________________________________________________________

// SimulateMsgCancelDTagTransfer tests and runs a single MsgCancelDTagTransfer
// nolint: funlen
func SimulateMsgCancelDTagTransfer(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {
		sender, receiver, skip := randomCancelDTagTransferFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, ""), nil, nil
		}

		msg := types.NewMsgCancelDTagTransferRequest(
			sender.Address.String(),
			receiver.Address.String(),
		)

		err = sendMsgMsgCancelDTagTransfer(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{sender.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, ""), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgCancelDTagTransfer"), nil, nil
	}
}

// sendMsgMsgCancelDTagTransfer sends a transaction with a MsgCancelDTagTransfer from a provided random account.
func sendMsgMsgCancelDTagTransfer(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgCancelDTagTransfer, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
) error {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	account := ak.GetAccount(ctx, addr)
	coins := bk.SpendableCoins(ctx, account.GetAddress())

	fees, err := simtypes.RandomFees(r, ctx, coins)
	if err != nil {
		return err
	}

	txGen := simappparams.MakeTestEncodingConfig().TxConfig
	tx, err := helpers.GenTx(
		txGen,
		[]sdk.Msg{msg},
		fees,
		DefaultGasValue,
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

// randomCancelDTagTransferFields returns random refuse DTag transfer fields
func randomCancelDTagTransferFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (simtypes.Account, simtypes.Account, bool) {
	if len(accs) == 0 {
		return simtypes.Account{}, simtypes.Account{}, true
	}

	// Get random accounts
	sender, _ := simtypes.RandomAcc(r, accs)
	receiver, _ := simtypes.RandomAcc(r, accs)

	// skip if the two addresses are equals
	if receiver.Equals(sender) {
		return simtypes.Account{}, simtypes.Account{}, true
	}

	req := types.NewDTagTransferRequest("dtag", sender.Address.String(), receiver.Address.String())
	err := k.SaveDTagTransferRequest(ctx, req)
	if err != nil {
		return simtypes.Account{}, simtypes.Account{}, true
	}

	return sender, receiver, false
}
