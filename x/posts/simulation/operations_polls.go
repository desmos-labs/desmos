package simulation

// DONTCOVER

import (
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/desmos-labs/desmos/v3/testutil/simtesting"
	feeskeeper "github.com/desmos-labs/desmos/v3/x/fees/keeper"
	"github.com/desmos-labs/desmos/v3/x/posts/keeper"
	"github.com/desmos-labs/desmos/v3/x/posts/types"
	subspaceskeeper "github.com/desmos-labs/desmos/v3/x/subspaces/keeper"
	subspacessim "github.com/desmos-labs/desmos/v3/x/subspaces/simulation"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// SimulateMsgAnswerPoll tests and runs a single msg answer poll post
func SimulateMsgAnswerPoll(
	k keeper.Keeper, sk subspaceskeeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		answer, user, skip := randomAnswerPollFields(r, ctx, accs, k, sk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "answer poll"), nil, nil
		}

		msg := types.NewMsgAnswerPoll(answer.SubspaceID, answer.PostID, answer.PollID, answer.AnswersIndexes, user.Address.String())
		err = simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{user.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "answer poll"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "answer poll", nil), nil, nil
	}
}

// randomAnswerPollFields returns the data needed to answer a user poll
func randomAnswerPollFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, sk subspaceskeeper.Keeper,
) (answer types.UserAnswer, user simtypes.Account, skip bool) {
	if len(accs) == 0 {
		// Skip because there are no accounts
		skip = true
		return
	}

	// Get a poll
	var polls []types.Attachment
	k.IterateActivePollsQueue(ctx, time.Now(), func(poll types.Attachment) (stop bool) {
		polls = append(polls, poll)
		return false
	})
	if len(polls) == 0 {
		// Skip because there are no active polls
		skip = true
		return
	}

	// Get a random poll
	poll := RandomAttachment(r, polls)

	// Get a user
	users, _ := sk.GetUsersWithRootPermission(ctx, poll.SubspaceID, subspacestypes.PermissionInteractWithContent)
	acc := subspacessim.GetAccount(subspacessim.RandomAddress(r, users), accs)
	if acc == nil {
		// Skip the operation withofut error as the account is not valid
		skip = true
		return
	}
	user = *acc

	// Get some answers
	answersIndexes := RandomAnswersIndexes(r, poll.Content.GetCachedValue().(*types.Poll))
	userAnswer := types.NewUserAnswer(poll.SubspaceID, poll.PostID, poll.ID, answersIndexes, user.Address.String())
	return userAnswer, user, false
}
