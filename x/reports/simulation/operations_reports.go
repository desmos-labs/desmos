package simulation

// DONTCOVER

import (
	"math/rand"

	"github.com/desmos-labs/desmos/x/reports/keeper"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/tendermint/tendermint/crypto"

	postskeeper "github.com/desmos-labs/desmos/x/posts/keeper"
	"github.com/desmos-labs/desmos/x/reports/types"
)

// SimulateMsgReportPost tests and runs a single msg reports post created by a random account.
// nolint: funlen
func SimulateMsgReportPost(ak auth.AccountKeeper, k keeper.Keeper, pk postskeeper.Keeper) sim.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (sim.OperationMsg, []sim.FutureOperation, error) {
		data, skip := randomReportPostFields(r, ctx, accs, ak, pk)
		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgReportPost(
			data.PostID,
			data.Type,
			data.Message,
			data.Creator.Address,
		)

		err := sendMsgReportPost(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{data.Creator.PrivKey})
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgReportPost sends a transaction with a MsgReportPost from a provided random account.
func sendMsgReportPost(
	r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgReportPost, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
) error {
	account := ak.GetAccount(ctx, msg.Report.User)
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

func randomReportPostFields(
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, ak auth.AccountKeeper, pk postskeeper.Keeper,
) (*ReportsData, bool) {
	posts := pk.GetPosts(ctx)
	if posts == nil {
		return nil, true
	}

	reportsData := RandomReportsData(r, posts, accs)
	acc := ak.GetAccount(ctx, reportsData.Creator.Address)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return nil, true
	}

	return &reportsData, false
}
