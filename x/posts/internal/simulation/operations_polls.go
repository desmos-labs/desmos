package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/posts/internal/keeper"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/tendermint/tendermint/crypto"
)

// SimulateMsgAnswerToPoll tests and runs a single msg poll answer where the answering user account already exists
// nolint: funlen
func SimulateMsgAnswerToPoll(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (sim.OperationMsg, []sim.FutureOperation, error) {

		acc, answers, postID, skip, err := randomPollAnswerFields(r, ctx, accs, k, ak)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgAnswerPoll(postID, answers, acc.Address)
		err = sendMsgAnswerPoll(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{acc.PrivKey})
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgAnswerPoll sends a transaction with a MsgAnswerPoll from a provided random account.
func sendMsgAnswerPoll(
	r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgAnswerPoll, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
) error {

	account := ak.GetAccount(ctx, msg.Answerer)
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

// randomPollAnswerFields returns the data used to create a MsgAnswerPoll message
func randomPollAnswerFields(
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper, ak auth.AccountKeeper,
) (sim.Account, []types.AnswerID, types.PostID, bool, error) {

	post, _ := RandomPost(r, k.GetPosts(ctx))

	// Skip the operation without any error if there is no poll, or the poll is closed
	if post.PollData == nil || !post.PollData.Open {
		return sim.Account{}, nil, "", true, nil
	}

	simAccount, _ := sim.RandomAcc(r, accs)
	acc := ak.GetAccount(ctx, simAccount.Address)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return sim.Account{}, nil, "", true, nil
	}

	// Skip the operation without err as the poll does not allow to edit answers
	currentAnswers := k.GetPollAnswersByUser(ctx, post.PostID, acc.GetAddress())
	if len(currentAnswers) > 0 && !post.PollData.AllowsAnswerEdits {
		return sim.Account{}, nil, "", true, nil
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

	return simAccount, answers, post.PostID, false, nil
}
