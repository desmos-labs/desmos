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

// SimulateMsgCreateRelationship tests and runs a single msg create relationships
// nolint: funlen
func SimulateMsgCreateRelationship(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {

		sender, relationship, skip := randomRelationshipFields(r, ctx, accs, k)
		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgCreateRelationship(sender.Address, relationship.Recipient, relationship.Subspace)
		if err := sendMsgCreateRelationship(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{sender.PrivKey}); err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgCreateRelationship sends a transaction with a Relationship from a provided random account
func sendMsgCreateRelationship(r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgCreateRelationship, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
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

// randomRelationshipFields returns random relationships fields
func randomRelationshipFields(
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper,
) (sim.Account, types.Relationship, bool) {
	if len(accs) == 0 {
		return sim.Account{}, types.Relationship{}, true
	}

	// Get random accounts
	sender, _ := sim.RandomAcc(r, accs)
	receiver, _ := sim.RandomAcc(r, accs)

	subspace := RandomSubspace(r)

	// skip if the two relationship are equals
	if sender.Equals(receiver) {
		return sim.Account{}, types.Relationship{}, true
	}

	rel := types.NewRelationship(receiver.Address, subspace)

	// skip if relationships already exists
	relationships := k.GetUserRelationships(ctx, sender.Address)
	for _, relationship := range relationships {
		if relationship.Equals(rel) {
			return sim.Account{}, types.Relationship{}, true
		}
	}

	return sender, rel, false
}

// SimulateMsgDeleteRelationship tests and runs a single msg delete relationships
// nolint: funlen
func SimulateMsgDeleteRelationship(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {

		sender, relationship, skip := randomDeleteRelationshipFields(r, ctx, accs, k)
		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgDeleteRelationship(sender.Address, relationship.Recipient, relationship.Subspace)
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
) (sim.Account, types.Relationship, bool) {
	if len(accs) == 0 {
		return sim.Account{}, types.Relationship{}, true
	}

	// Get random accounts
	user, _ := sim.RandomAcc(r, accs)

	relationships := k.GetUserRelationships(ctx, user.Address)

	// skip the test if the user has no relationships
	if len(relationships) == 0 {
		return sim.Account{}, types.Relationship{}, true
	}

	return user, RandomRelationship(r, relationships), false
}
