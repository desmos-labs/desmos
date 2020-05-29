package simulation

import (
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/posts"
	"github.com/desmos-labs/desmos/x/reports/internal/keeper"
	"github.com/desmos-labs/desmos/x/reports/internal/types"
	"github.com/tendermint/tendermint/crypto"
)

// SimulateMsgReportPost tests and runs a single msg reports post created by a random account.
// nolint: funlen
func SimulateMsgReportPost(ak auth.AccountKeeper, k keeper.Keeper, pk posts.Keeper) sim.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (sim.OperationMsg, []sim.FutureOperation, error) {
		data, skip, err := randomReportPostFields(r, ctx, accs, ak)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgReportPost(
			data.PostID,
			data.Type,
			data.Message,
			data.Creator.Address,
		)

		post := posts.NewPost(
			data.PostID,
			"",
			"message",
			true,
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			map[string]string{},
			time.Now(),
			data.Creator.Address,
		)

		// Save the post
		k.PostKeeper.SavePost(ctx, post)

		err = sendMsgReportPost(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{data.Creator.PrivKey})
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
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, ak auth.AccountKeeper,
) (*ReportsData, bool, error) {
	reportsData := RandomReportsData(r, accs)
	acc := ak.GetAccount(ctx, reportsData.Creator.Address)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return nil, true, nil
	}

	return &reportsData, false, nil
}
