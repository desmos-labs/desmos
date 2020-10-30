package simulation

// DONTCOVER

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/tendermint/tendermint/crypto"

	"github.com/desmos-labs/desmos/x/posts/keeper"
	"github.com/desmos-labs/desmos/x/posts/types"
)

// SimulateMsgCreatePost tests and runs a single msg create post where the post creator account already exists
// nolint: funlen
func SimulateMsgCreatePost(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (sim.OperationMsg, []sim.FutureOperation, error) {

		data, skip := randomPostCreateFields(r, ctx, accs, k, ak)
		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgCreatePost(
			data.Message,
			data.ParentID,
			data.AllowsComments,
			data.Subspace,
			data.OptionalData,
			data.Creator.Address,
			data.Attachments,
			data.PollData,
		)

		err := sendMsgCreatePost(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{data.Creator.PrivKey})
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgCreatePost sends a transaction with a MsgCreatePost from a provided random account.
func sendMsgCreatePost(
	r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgCreatePost, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
) error {
	account := ak.GetAccount(ctx, msg.Creator)

	tx := helpers.GenTx(
		[]sdk.Msg{msg},
		sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10000))),
		DefaultGasValue,
		chainID,
		[]uint64{account.GetAccountNumber()},
		[]uint64{account.GetSequence()},
		privkeys...,
	)

	_, _, err := app.Deliver(tx)
	if err != nil {
		return err
	}

	return nil
}

// randomPostCreateFields returns the creator of the post as well as the parent id
func randomPostCreateFields(
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper, ak auth.AccountKeeper,
) (*PostData, bool) {

	postData := RandomPostData(r, accs)
	acc := ak.GetAccount(ctx, postData.Creator.Address)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return nil, true
	}

	// Skip the operation as the poll is closed
	if postData.PollData != nil && postData.PollData.EndDate.Before(ctx.BlockTime()) {
		return nil, true
	}

	for _, attachment := range postData.Attachments {
		for _, tag := range attachment.Tags {
			if k.IsUserBlocked(ctx, tag, postData.Creator.Address) {
				return nil, true
			}
		}
	}

	posts := k.GetPosts(ctx)
	if posts != nil {
		if parent, _ := RandomPost(r, posts); parent.AllowsComments {
			postData.ParentID = parent.PostID
		}
	}

	return &postData, false
}

// SimulateMsgEditPost tests and runs a single msg edit post where the post creator account already exists
// nolint: funlen
func SimulateMsgEditPost(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (sim.OperationMsg, []sim.FutureOperation, error) {

		account, id, message, attachments, pollData, skip := randomPostEditFields(r, ctx, accs, k, ak)
		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgEditPost(id, message, attachments, pollData, account.Address)

		err := sendMsgEditPost(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{account.PrivKey})
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgEditPost sends a transaction with a MsgEditPost from a provided random account.
func sendMsgEditPost(
	r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgEditPost, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
) error {

	account := ak.GetAccount(ctx, msg.Editor)

	tx := helpers.GenTx(
		[]sdk.Msg{msg},
		sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10000))),
		DefaultGasValue,
		chainID,
		[]uint64{account.GetAccountNumber()},
		[]uint64{account.GetSequence()},
		privkeys...,
	)

	_, _, err := app.Deliver(tx)
	if err != nil {
		return err
	}

	return nil
}

// randomPostEditFields returns the data needed to edit a post
func randomPostEditFields(
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper, _ auth.AccountKeeper,
) (sim.Account, types.PostID, string, types.Attachments, *types.PollData, bool) {

	post, _ := RandomPost(r, k.GetPosts(ctx))
	acc := GetAccount(post.Creator, accs)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return sim.Account{}, "", "", nil, nil, true
	}

	editedAttachments := RandomAttachments(r, accs)

	for _, attachment := range editedAttachments {
		for _, tag := range attachment.Tags {
			if k.IsUserBlocked(ctx, tag, post.Creator) {
				return sim.Account{}, "", "", nil, nil, true
			}
		}
	}

	return *acc, post.PostID, RandomMessage(r), editedAttachments, RandomPollData(r), false
}
