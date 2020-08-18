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

// SimulateMsgCreateMonoDirectionalRelationship tests and runs a single msg create monoDirectional relationship
// nolint: funlen
func SimulateMsgCreateMonoDirectionalRelationship(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {

		acc, data, skip := randomMonoRelationshipFields(r, ctx, accs, k)
		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgCreateMonoDirectionalRelationship(data.Sender, data.Receiver)
		if err := sendMsgCreateMonoDirectionalRelationship(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{acc.PrivKey}); err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgCreateMonoDirectionalRelationship sends a transaction with a MsgCreateMonoDirectionalRelationship from a provided random account
func sendMsgCreateMonoDirectionalRelationship(r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgCreateMonoDirectionalRelationship, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
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

// randomMonoRelationshipFields returns random monoDirectional relationship fields
func randomMonoRelationshipFields(
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper,
) (sim.Account, types.MonodirectionalRelationship, bool) {
	if len(accs) == 0 {
		return sim.Account{}, types.MonodirectionalRelationship{}, true
	}

	// Get random accounts
	sender, _ := sim.RandomAcc(r, accs)
	receiver, _ := sim.RandomAcc(r, accs)

	// skip if the two address are equals
	if sender.Equals(receiver) {
		return sim.Account{}, types.MonodirectionalRelationship{}, true
	}

	relationship := types.NewMonodirectionalRelationship(sender.Address, receiver.Address)

	// skip if relationship already exists
	if k.DoesRelationshipExist(ctx, relationship.ID) {
		return sim.Account{}, types.MonodirectionalRelationship{}, true
	}

	return sender, relationship, false
}

// SimulateMsgRequestBidirectionalRelationship tests and runs a single msg request biDirectional relationship
// nolint: funlen
func SimulateMsgRequestBidirectionalRelationship(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {

		acc, data, skip := randomBiDirectionalRelationshipFields(r, ctx, accs, k)
		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgRequestBidirectionalRelationship(data.Sender, data.Receiver, sim.RandStringOfLength(r, 5))
		if err := sendMsgRequestBiDirectionalRelationship(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{acc.PrivKey}); err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgRequestBiDirectionalRelationship sends a transaction with a MsgRequestBidirectionalRelationship from a provided random account
func sendMsgRequestBiDirectionalRelationship(r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgRequestBidirectionalRelationship, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
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

// randomBiDirectionalRelationshipFields returns random biDirectional relationship fields
func randomBiDirectionalRelationshipFields(
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper,
) (sim.Account, types.BidirectionalRelationship, bool) {
	if len(accs) == 0 {
		return sim.Account{}, types.BidirectionalRelationship{}, true
	}

	// Get random accounts
	sender, _ := sim.RandomAcc(r, accs)
	receiver, _ := sim.RandomAcc(r, accs)

	// skip if the two address are equals
	if sender.Equals(receiver) {
		return sim.Account{}, types.BidirectionalRelationship{}, true
	}

	relationship := types.NewBiDirectionalRelationship(sender.Address, receiver.Address, types.Sent)

	// skip if relationship already exists
	if k.DoesRelationshipExist(ctx, relationship.ID) {
		return sim.Account{}, types.BidirectionalRelationship{}, true
	}

	return sender, relationship, false
}

// SimulateMsgAcceptBidirectionalRelationship tests and runs a single msg accept biDirectional relationship
// nolint: funlen
func SimulateMsgAcceptBidirectionalRelationship(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {

		acc, data, skip := randomBiDirectionalRelationshipSentFields(r, ctx, accs, k)
		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgAcceptBidirectionalRelationship(data.ID, data.Receiver)
		if err := sendMsgAcceptBidirectionalRelationship(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{acc.PrivKey}); err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgAcceptBidirectionalRelationship sends a transaction with a MsgAcceptBidirectionalRelationship from a provided random account
func sendMsgAcceptBidirectionalRelationship(r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgAcceptBidirectionalRelationship, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
) error {
	account := ak.GetAccount(ctx, msg.Receiver)
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

// SimulateMsgDenyBidirectionalRelationship tests and runs a single msg deny biDirectional relationship
// nolint: funlen
func SimulateMsgDenyBidirectionalRelationship(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {

		acc, data, skip := randomBiDirectionalRelationshipSentFields(r, ctx, accs, k)
		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgAcceptBidirectionalRelationship(data.ID, data.Receiver)
		if err := sendMsgAcceptBidirectionalRelationship(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{acc.PrivKey}); err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgDenyBidirectionalRelationship sends a transaction with a MsgDenyBidirectionalRelationship from a provided random account
func sendMsgDenyBidirectionalRelationship(r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgAcceptBidirectionalRelationship, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
) error {
	account := ak.GetAccount(ctx, msg.Receiver)
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

// randomBiDirectionalRelationshipSentFields returns random biDirectional relationship fields that has already been sent
func randomBiDirectionalRelationshipSentFields(r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper,
) (sim.Account, types.BidirectionalRelationship, bool) {
	if len(accs) == 0 {
		return sim.Account{}, types.BidirectionalRelationship{}, true
	}

	// Get random accounts
	user, _ := sim.RandomAcc(r, accs)

	relationships := k.GetUserRelationships(ctx, user.Address)

	// skip the test if the user has no relationships
	if len(relationships) == 0 {
		return sim.Account{}, types.BidirectionalRelationship{}, true
	}

	for _, relationship := range relationships {
		if rel, ok := relationship.(types.BidirectionalRelationship); ok && rel.Status == types.Sent {
			return user, rel, false
		}
	}

	return sim.Account{}, types.BidirectionalRelationship{}, true
}

// SimulateMsgDeleteRelationship tests and runs a single msg delete relationship
// nolint: funlen
func SimulateMsgDeleteRelationship(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {

		acc, relationshipID, skip := randomDeleteRelationshipFields(r, ctx, accs, k)
		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgDeleteRelationship(relationshipID, acc.Address)
		if err := sendMsgDeleteRelationship(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{acc.PrivKey}); err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgDeleteRelationship sends a transaction with a MsgDenyBidirectionalRelationship from a provided random account
func sendMsgDeleteRelationship(r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgDeleteRelationship, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
) error {
	account := ak.GetAccount(ctx, msg.User)
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

// randomDeleteRelationshipFields returns random delete relationship fields
func randomDeleteRelationshipFields(r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper,
) (sim.Account, types.RelationshipID, bool) {
	if len(accs) == 0 {
		return sim.Account{}, "", true
	}

	// Get random accounts
	user, _ := sim.RandomAcc(r, accs)

	relationships := k.GetUserRelationships(ctx, user.Address)

	// skip the test if the user has no relationships
	if len(relationships) == 0 {
		return sim.Account{}, "", true
	}

	return sim.Account{}, RandomRelationship(r, relationships).RelationshipID(), false
}
