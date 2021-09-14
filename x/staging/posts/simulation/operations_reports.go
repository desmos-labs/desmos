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

	postskeeper "github.com/desmos-labs/desmos/x/staging/posts/keeper"
	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

// SimulateMsgReportPost tests and runs a single MsgReportPost created by a random account.
// nolint: funlen
func SimulateMsgReportPost(
	pk postskeeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		data, skip := randomReportPostFields(r, ctx, accs, ak, pk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgReportPost"), nil, nil
		}

		msg := types.NewMsgReportPost(
			data.PostID,
			data.Reasons,
			data.Message,
			data.Creator.Address.String(),
		)

		err := sendMsgReportPost(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{data.Creator.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgReportPost"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgReportPost", nil), nil, nil
	}
}

// sendMsgReportPost sends a transaction with a MsgReportPost from a provided random account.
func sendMsgReportPost(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgReportPost, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
) error {
	addr, _ := sdk.AccAddressFromBech32(msg.User)
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

func randomReportPostFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, ak authkeeper.AccountKeeper, pk postskeeper.Keeper,
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

// ___________________________________________________________________________________________________________________

// SimulateMsgRemovePostReport tests and runs a single msg remove report where the reporting user account already exists
// nolint: funlen
func SimulateMsgRemovePostReport(
	pk postskeeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		data, skip := randomRemovePostReportFields(r, ctx, accs, ak, pk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgRemovePostReport"), nil, nil
		}

		msg := types.NewMsgRemovePostReport(
			data.PostID,
			data.Creator.Address.String(),
		)

		err := sendMsgRemovePostReport(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{data.Creator.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgRemovePostReport"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgRemovePostReport", nil), nil, nil
	}
}

// sendMsgRemovePostReport sends a transaction with a MsgRemovePostReport from a provided random account.
func sendMsgRemovePostReport(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgRemovePostReport, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
) error {
	addr, _ := sdk.AccAddressFromBech32(msg.User)
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

// randomRemovePostReportFields returns the data used to create a MsgRemovePostReport message
func randomRemovePostReportFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, ak authkeeper.AccountKeeper, pk postskeeper.Keeper,
) (*ReportsData, bool) {
	posts := pk.GetPosts(ctx)
	if len(posts) == 0 {
		// Skip cause there are no posts
		return nil, true
	}

	post, _ := RandomPost(r, posts)
	var reports []types.Report
	pk.IteratePostReports(ctx, post.PostID, func(_ int64, report types.Report) bool {
		reports = append(reports, report)
		return false
	})

	// Skip if the post has no reports
	if len(reports) == 0 {
		return nil, true
	}

	report := reports[r.Intn(len(reports))]
	addr, _ := sdk.AccAddressFromBech32(report.User)
	acc := ak.GetAccount(ctx, addr)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return nil, true
	}

	user := GetAccount(addr, accs)
	data := ReportsData{PostID: post.PostID, Creator: *user}
	return &data, false
}
