package simulation

// DONTCOVER

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/relationships/keeper"
	"github.com/desmos-labs/desmos/x/relationships/types"
	"github.com/tendermint/tendermint/crypto"
)

// SimulateMsgCreateMonoDirectionalRelationship tests and runs a single msg create monoDirectional relationships
// nolint: funlen
func SimulateMsgCreateMonoDirectionalRelationship(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {

		sender, receiver, skip := randomMonoRelationshipFields(r, ctx, accs, k)
		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgCreateMonoDirectionalRelationship(sender.Address, receiver)
		if err := sendMsgCreateMonoDirectionalRelationship(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{sender.PrivKey}); err != nil {
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

// randomMonoRelationshipFields returns random monoDirectional relationships fields
func randomMonoRelationshipFields(
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper,
) (sim.Account, sdk.AccAddress, bool) {
	if len(accs) == 0 {
		return sim.Account{}, nil, true
	}

	// Get random accounts
	sender, _ := sim.RandomAcc(r, accs)
	receiver, _ := sim.RandomAcc(r, accs)

	// skip if the two address are equals
	if sender.Equals(receiver) {
		return sim.Account{}, nil, true
	}

	// skip if relationships already exists
	relationships := k.GetUserRelationships(ctx, sender.Address)
	for _, address := range relationships {
		if address.Equals(receiver.Address) {
			return sim.Account{}, nil, true
		}
	}

	return sender, receiver.Address, false
}

// SimulateMsgDeleteRelationship tests and runs a single msg delete relationships
// nolint: funlen
func SimulateMsgDeleteRelationship(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {

		sender, receiver, skip := randomDeleteRelationshipFields(r, ctx, accs, k)
		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgDeleteRelationship(sender.Address, receiver)
		if err := sendMsgDeleteRelationship(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{sender.PrivKey}); err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgDeleteRelationship sends a transaction with a MsgDenyBidirectionalRelationship from a provided random account
func sendMsgDeleteRelationship(r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgDeleteRelationship, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
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

// randomDeleteRelationshipFields returns random delete relationships fields
func randomDeleteRelationshipFields(r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper,
) (sim.Account, sdk.AccAddress, bool) {
	if len(accs) == 0 {
		return sim.Account{}, nil, true
	}

	// Get random accounts
	user, _ := sim.RandomAcc(r, accs)

	relationships := k.GetUserRelationships(ctx, user.Address)

	// skip the test if the user has no relationships
	if len(relationships) == 0 {
		return sim.Account{}, nil, true
	}

	return user, RandomRelationship(r, relationships), false
}
