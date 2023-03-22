package simulation

// DONTCOVER

import (
	"math/rand"

	feeskeeper "github.com/desmos-labs/desmos/v4/x/fees/keeper"

	"github.com/desmos-labs/desmos/v4/testutil/simtesting"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/profiles/keeper"
	"github.com/desmos-labs/desmos/v4/x/profiles/types"
)

// SimulateMsgRequestDTagTransfer tests and runs a single MsgRequestDTagTransfer
func SimulateMsgRequestDTagTransfer(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {
		sender, receiver, skip := randomDTagRequestTransferFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.ModuleName, "MsgRequestDTagTransfer", "skip"), nil, nil
		}

		msg := types.NewMsgRequestDTagTransfer(sender.Address.String(), receiver.GetAddress().String())

		txCtx, err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, sender)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, "MsgRequestDTagTransfer", "invalid"), nil, nil
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// randomDTagRequestTransferFields returns random dTagRequest data
func randomDTagRequestTransferFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (sender simtypes.Account, receiver *types.Profile, skip bool) {
	if len(accs) == 0 {
		return simtypes.Account{}, nil, true
	}

	// Get random sender account
	sender, _ = simtypes.RandomAcc(r, accs)

	profiles := k.GetProfiles(ctx)
	if profiles == nil {
		return simtypes.Account{}, nil, true
	}

	// Get a random Profile
	receiverProfile := RandomProfile(r, profiles)
	receiverAddress := receiverProfile.GetAddress()
	if receiverAddress.Equals(sender.Address) {
		return simtypes.Account{}, nil, true
	}

	// Skip if the sender is blocked
	if k.IsUserBlocked(ctx, receiverAddress.String(), sender.Address.String()) {
		return simtypes.Account{}, nil, true
	}

	// Skip if requests already exists
	_, found, err := k.GetDTagTransferRequest(ctx, sender.Address.String(), receiverAddress.String())
	if err != nil || found {
		return simtypes.Account{}, nil, true
	}

	return sender, receiverProfile, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgAcceptDTagTransfer tests and runs a single MsgAcceptDTagTransfer
func SimulateMsgAcceptDTagTransfer(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {
		acc, request, dTag, skip := randomDTagAcceptRequestTransferFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.ModuleName, "MsgAcceptDTagTransferRequest", "skip"), nil, nil
		}

		msg := types.NewMsgAcceptDTagTransferRequest(dTag, request.Sender, request.Receiver)
		txCtx, err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, acc)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, "MsgAcceptDTagTransferRequest", "invalid"), nil, nil
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// randomDTagAcceptRequestTransferFields returns random dTagRequest data and a random dTag
func randomDTagAcceptRequestTransferFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (simtypes.Account, types.DTagTransferRequest, string, bool) {
	if len(accs) == 0 {
		return simtypes.Account{}, types.DTagTransferRequest{}, "", true
	}

	// Get random accounts
	sender, _ := simtypes.RandomAcc(r, accs)
	receiver, _ := simtypes.RandomAcc(r, accs)

	// skip if the two addresses are equals
	if sender.Address.Equals(receiver.Address) {
		return simtypes.Account{}, types.DTagTransferRequest{}, "", true
	}

	// Skip if requests doesn't exists
	req, found, err := k.GetDTagTransferRequest(ctx, sender.Address.String(), receiver.Address.String())
	if err != nil || !found {
		return simtypes.Account{}, types.DTagTransferRequest{}, "", true
	}

	return receiver, req, RandomDTag(r), false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgRefuseDTagTransfer tests and runs a single MsgRefuseDTagTransfer
func SimulateMsgRefuseDTagTransfer(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {
		sender, receiver, skip := randomRefuseDTagTransferFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.ModuleName, "MsgRefuseDTagTransferRequest", "skip"), nil, nil
		}

		msg := types.NewMsgRefuseDTagTransferRequest(sender.Address.String(), receiver.Address.String())
		txCtx, err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, receiver)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, "MsgRefuseDTagTransferRequest", "invalid"), nil, nil
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// randomRefuseDTagTransferFields returns random refuse DTag transfer fields
func randomRefuseDTagTransferFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (simtypes.Account, simtypes.Account, bool) {
	if len(accs) == 0 {
		return simtypes.Account{}, simtypes.Account{}, true
	}

	// Get a random request
	requests := k.GetDTagTransferRequests(ctx)
	if len(requests) == 0 {
		return simtypes.Account{}, simtypes.Account{}, true
	}
	request := RandomDTagTransferRequest(r, requests)

	// Get the sender account
	senderAddr, _ := sdk.AccAddressFromBech32(request.Sender)
	senderAcc := GetSimAccount(senderAddr, accs)
	if senderAcc == nil {
		return simtypes.Account{}, simtypes.Account{}, true
	}
	sender := *senderAcc

	// Get the receiver account
	receiverAddr, _ := sdk.AccAddressFromBech32(request.Receiver)
	receiverAcc := GetSimAccount(receiverAddr, accs)
	if receiverAcc == nil {
		return simtypes.Account{}, simtypes.Account{}, true
	}
	receiver := *receiverAcc

	return sender, receiver, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgCancelDTagTransfer tests and runs a single MsgCancelDTagTransfer
func SimulateMsgCancelDTagTransfer(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {
		sender, receiver, skip := randomCancelDTagTransferFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.ModuleName, "MsgCancelDTagTransfer", "skip"), nil, nil
		}

		msg := types.NewMsgCancelDTagTransferRequest(
			sender.Address.String(),
			receiver.Address.String(),
		)

		txCtx, err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, sender)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, "MsgCancelDTagTransfer", "invalid"), nil, nil
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// randomCancelDTagTransferFields returns random refuse DTag transfer fields
func randomCancelDTagTransferFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (simtypes.Account, simtypes.Account, bool) {
	if len(accs) == 0 {
		return simtypes.Account{}, simtypes.Account{}, true
	}

	// Get a random request
	requests := k.GetDTagTransferRequests(ctx)
	if len(requests) == 0 {
		return simtypes.Account{}, simtypes.Account{}, true
	}
	request := RandomDTagTransferRequest(r, requests)

	// Get the sender account
	senderAddr, _ := sdk.AccAddressFromBech32(request.Sender)
	senderAcc := GetSimAccount(senderAddr, accs)
	if senderAcc == nil {
		return simtypes.Account{}, simtypes.Account{}, true
	}
	sender := *senderAcc

	// Get the receiver account
	receiverAddr, _ := sdk.AccAddressFromBech32(request.Receiver)
	receiverAcc := GetSimAccount(receiverAddr, accs)
	if receiverAcc == nil {
		return simtypes.Account{}, simtypes.Account{}, true
	}
	receiver := *receiverAcc

	return sender, receiver, false
}
