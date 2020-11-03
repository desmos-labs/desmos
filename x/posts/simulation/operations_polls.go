package simulation

// DONTCOVER

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/fees"
	"github.com/tendermint/tendermint/crypto"

	"github.com/desmos-labs/desmos/x/posts/keeper"
	"github.com/desmos-labs/desmos/x/posts/types"
)

// SimulateMsgAnswerToPoll tests and runs a single msg poll answer where the answering user account already exists
// nolint: funlen
func SimulateMsgAnswerToPoll(k keeper.Keeper, ak auth.AccountKeeper, fk fees.Keeper) sim.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (sim.OperationMsg, []sim.FutureOperation, error) {

		acc, answers, postID, skip := randomPollAnswerFields(r, ctx, accs, k, ak)
		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgAnswerPoll(postID, answers, acc.Address)
		err := sendMsgAnswerPoll(r, app, ak, fk, msg, ctx, chainID, []crypto.PrivKey{acc.PrivKey})
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgAnswerPoll sends a transaction with a MsgAnswerPoll from a provided random account.
func sendMsgAnswerPoll(
	r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper, fk fees.Keeper,
	msg types.MsgAnswerPoll, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
) error {
	account := ak.GetAccount(ctx, msg.Answerer)
	coins := account.SpendableCoins(ctx.BlockTime())

	randFees, err := sim.RandomFees(r, ctx, coins)
	if err != nil {
		return err
	}

	msgs := []sdk.Msg{msg}

	if err := fk.CheckFees(ctx, authtypes.NewStdFee(helpers.DefaultGenTxGas, randFees), msgs); err != nil {
		return nil
	}

	tx := helpers.GenTx(
		msgs,
		randFees,
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

// randomPollAnswerFields returns the data used to create a MsgAnswerPoll message
func randomPollAnswerFields(
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper, ak auth.AccountKeeper,
) (sim.Account, []types.AnswerID, types.PostID, bool) {

	posts := k.GetPosts(ctx)
	if posts == nil {
		return sim.Account{}, nil, "", true
	}

	post, _ := RandomPost(r, posts)

	// Skip the operation without any error if there is no poll, or the poll is closed
	if post.PollData == nil || post.PollData.EndDate.Before(ctx.BlockTime()) {
		return sim.Account{}, nil, "", true
	}

	simAccount, _ := sim.RandomAcc(r, accs)
	acc := ak.GetAccount(ctx, simAccount.Address)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return sim.Account{}, nil, "", true
	}

	// Skip the operation without err as the poll does not allow to edit answers
	currentAnswers := k.GetPollAnswersByUser(ctx, post.PostID, acc.GetAddress())
	if len(currentAnswers) > 0 && !post.PollData.AllowsAnswerEdits {
		return sim.Account{}, nil, "", true
	}

	providedAnswers := post.PollData.ProvidedAnswers

	answersLength := 1
	if post.PollData.AllowsMultipleAnswers {
		answersLength = r.Intn(len(post.PollData.ProvidedAnswers)) + 1 // At least one answer is necessary
	}

	answers := make([]types.AnswerID, answersLength)
	for i := 0; i < answersLength; i++ {
		answers[i] = providedAnswers[i].ID
	}

	return simAccount, answers, post.PostID, false
}
