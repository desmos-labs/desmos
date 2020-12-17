package simulation

// DONTCOVER

import (
	"math/rand"

	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/posts/keeper"
	"github.com/desmos-labs/desmos/x/posts/types"
)

// SimulateMsgAnswerToPoll tests and runs a single msg poll answer where the answering user account already exists
// nolint: funlen
func SimulateMsgAnswerToPoll(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		acc, answers, postID, skip := randomPollAnswerFields(r, ctx, accs, k, ak)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgAnswerToPoll"), nil, nil
		}

		msg := types.NewMsgAnswerPoll(postID, answers, acc.Address.String())
		err := sendMsgAnswerPoll(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{acc.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgAnswerToPoll"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgAnswerToPoll"), nil, nil
	}
}

// sendMsgAnswerPoll sends a transaction with a MsgAnswerPoll from a provided random account.
func sendMsgAnswerPoll(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgAnswerPoll, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
) error {
	addr, _ := sdk.AccAddressFromBech32(msg.Answerer)
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

// randomPollAnswerFields returns the data used to create a MsgAnswerPoll message
func randomPollAnswerFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, ak authkeeper.AccountKeeper,
) (simtypes.Account, []string, string, bool) {
	posts := k.GetPosts(ctx)
	if len(posts) == 0 {
		// Skip cause there are no posts
		return simtypes.Account{}, nil, "", true
	}

	post, _ := RandomPost(r, posts)

	// Skip the operation without any error if there is no poll, or the poll is closed
	if post.PollData == nil || post.PollData.EndDate.Before(ctx.BlockTime()) {
		return simtypes.Account{}, nil, "", true
	}

	simAccount, _ := simtypes.RandomAcc(r, accs)
	acc := ak.GetAccount(ctx, simAccount.Address)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return simtypes.Account{}, nil, "", true
	}

	// Skip the operation without err as the poll does not allow to edit answers
	currentAnswers := k.GetPollAnswersByUser(ctx, post.PostID, acc.GetAddress().String())
	if len(currentAnswers) > 0 && !post.PollData.AllowsAnswerEdits {
		return simtypes.Account{}, nil, "", true
	}

	providedAnswers := post.PollData.ProvidedAnswers

	answersLength := 1
	if post.PollData.AllowsMultipleAnswers {
		answersLength = r.Intn(len(post.PollData.ProvidedAnswers)) + 1 // At least one answer is necessary
	}

	answers := make([]string, answersLength)
	for i := 0; i < answersLength; i++ {
		answers[i] = providedAnswers[i].ID
	}

	return simAccount, answers, post.PostID, false
}
