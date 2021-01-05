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

// SimulateMsgCreatePost tests and runs a single msg create post where the post creator account already exists
// nolint: funlen
func SimulateMsgCreatePost(k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		data, skip := randomPostCreateFields(r, ctx, accs, k, ak)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgCreatePost"), nil, nil
		}

		msg := types.NewMsgCreatePost(
			data.Message,
			data.ParentID,
			data.AllowsComments,
			data.Subspace,
			data.OptionalData,
			data.CreatorAccount.Address.String(),
			data.Attachments,
			data.PollData,
		)

		err := sendMsgCreatePost(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{data.CreatorAccount.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgCreatePost"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgCreatePost"), nil, nil
	}
}

// sendMsgCreatePost sends a transaction with a MsgCreatePost from a provided random account.
func sendMsgCreatePost(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgCreatePost, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
) error {
	addr, _ := sdk.AccAddressFromBech32(msg.Creator)
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

// randomPostCreateFields returns the creator of the post as well as the parent id
func randomPostCreateFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, ak authkeeper.AccountKeeper,
) (*PostData, bool) {

	postData := RandomPostData(r, accs)
	acc := ak.GetAccount(ctx, postData.CreatorAccount.Address)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return nil, true
	}

	// Skip the operation as the poll is closed
	if postData.PollData != nil && postData.PollData.EndDate.Before(ctx.BlockTime()) {
		return nil, true
	}

	// Check to make sure none that is tagged, is also blocked
	for _, attachment := range postData.Attachments {
		for _, tag := range attachment.Tags {
			if k.IsUserBlocked(ctx, tag, postData.CreatorAccount.Address.String(), postData.Subspace) {
				return nil, true
			}
		}
	}

	// Set the parent id properly
	postData.ParentID = ""
	posts := k.GetPosts(ctx)
	if posts != nil {
		if parent, _ := RandomPost(r, posts); parent.AllowsComments {
			postData.ParentID = parent.PostID
		}
	}

	return &postData, false
}

// ___________________________________________________________________________________________________________________

// SimulateMsgEditPost tests and runs a single msg edit post where the post creator account already exists
// nolint: funlen
func SimulateMsgEditPost(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		account, id, message, attachments, pollData, skip := randomPostEditFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgEditPost"), nil, nil
		}

		msg := types.NewMsgEditPost(id, message, attachments, pollData, account.Address.String())

		err := sendMsgEditPost(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{account.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgEditPost"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgEditPost"), nil, nil
	}
}

// sendMsgEditPost sends a transaction with a MsgEditPost from a provided random account.
func sendMsgEditPost(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgEditPost, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
) error {
	addr, _ := sdk.AccAddressFromBech32(msg.Editor)
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

// randomPostEditFields returns the data needed to edit a post
func randomPostEditFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (simtypes.Account, string, string, types.Attachments, *types.PollData, bool) {
	posts := k.GetPosts(ctx)
	if len(posts) == 0 {
		// Skip cause there are no posts
		return simtypes.Account{}, "", "", nil, nil, true
	}

	post, _ := RandomPost(r, posts)
	addr, _ := sdk.AccAddressFromBech32(post.Creator)
	acc := GetAccount(addr, accs)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return simtypes.Account{}, "", "", nil, nil, true
	}

	editedAttachments := RandomAttachments(r, accs)

	for _, attachment := range editedAttachments {
		for _, tag := range attachment.Tags {
			if k.IsUserBlocked(ctx, tag, post.Creator, post.Subspace) {
				return simtypes.Account{}, "", "", nil, nil, true
			}
		}
	}

	return *acc, post.PostID, RandomMessage(r), editedAttachments, RandomPollData(r), false
}
