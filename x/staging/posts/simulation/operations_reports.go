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
	subspaceskeeper "github.com/desmos-labs/desmos/x/staging/subspaces/keeper"
)

// SimulateMsgReportPost tests and runs a single MsgReportPost created by a random account.
// nolint: funlen
func SimulateMsgReportPost(
	pk postskeeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, sk subspaceskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		data, skip := randomReportPostFields(r, ctx, accs, ak, pk, sk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgReportPost"), nil, nil
		}

		msg := types.NewMsgReportPost(
			data.PostID,
			data.Type,
			data.Message,
			data.Creator.Address.String(),
		)

		err := sendMsgReportPost(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{data.Creator.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgReportPost"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgReportPost"), nil, nil
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
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, ak authkeeper.AccountKeeper, pk postskeeper.Keeper, sk subspaceskeeper.Keeper,
) (*ReportsData, bool) {
	posts := pk.GetPosts(ctx)
	if posts == nil {
		return nil, true
	}

	subspaceData := RandomSubspace(r, accs)

	if err := sk.SaveSubspace(ctx, subspaceData.Subspace, subspaceData.Subspace.Creator); err != nil {
		return nil, true
	}

	reportsData := RandomReportsData(r, posts, accs)

	post, found := pk.GetPost(ctx, reportsData.PostID)
	if !found {
		return nil, true
	}
	post.Subspace = subspaceData.Subspace.ID
	pk.SavePost(ctx, post)

	acc := ak.GetAccount(ctx, reportsData.Creator.Address)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return nil, true
	}

	return &reportsData, false
}
