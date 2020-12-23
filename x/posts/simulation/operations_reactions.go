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

// ---------------
// --- PostReaction
// ---------------

// SimulateMsgAddPostReaction tests and runs a single msg add reaction where the reacting user account already exists
// nolint: funlen
func SimulateMsgAddPostReaction(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		data, skip := randomAddPostReactionFields(r, ctx, accs, k, ak)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgAddPostReaction"), nil, nil
		}

		msg := types.NewMsgAddPostReaction(data.PostID, data.Shortcode, data.User.Address.String())
		err := sendMsgAddPostReaction(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{data.User.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgAddPostReaction"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgAddPostReaction"), nil, nil
	}
}

// sendMsgAddPostReaction sends a transaction with a MsgAddReaction from a provided random account.
func sendMsgAddPostReaction(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgAddPostReaction, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
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

// randomAddPostReactionFields returns the data used to create a MsgAddReaction message
func randomAddPostReactionFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, ak authkeeper.AccountKeeper,
) (*PostReactionData, bool) {

	posts := k.GetPosts(ctx)
	if posts == nil {
		return nil, true
	}
	post, _ := RandomPost(r, posts)

	var reaction types.RegisteredReaction
	data := RandomReactionData(r, accs)

	reaction = types.NewRegisteredReaction(data.Creator.Address.String(), data.ShortCode, data.Value, post.Subspace)
	k.SaveRegisteredReaction(ctx, reaction)

	reactionData := RandomPostReactionData(r, accs, post.PostID, reaction.ShortCode, reaction.Value)
	acc := ak.GetAccount(ctx, reactionData.User.Address)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return nil, true
	}

	// Skip if the reaction already exists
	reactions := types.NewPostReactions(k.GetPostReactions(ctx, post.PostID)...)
	if reactions.ContainsReactionFrom(reactionData.User.Address.String(), reactionData.Value) {
		return nil, true
	}

	return &reactionData, false
}

// ___________________________________________________________________________________________________________________

// SimulateMsgRemovePostReaction tests and runs a single msg remove reaction where the reacting user account already exists
// nolint: funlen
func SimulateMsgRemovePostReaction(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		data, skip := randomRemovePostReactionFields(r, ctx, accs, k, ak)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgRemovePostReaction"), nil, nil
		}

		msg := types.NewMsgRemovePostReaction(data.PostID, data.User.Address.String(), data.Shortcode)
		err := sendMsgRemovePostReaction(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{data.User.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgRemovePostReaction"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgRemovePostReaction"), nil, nil
	}
}

// sendMsgRemovePostReaction sends a transaction with a MsgRemoveReaction from a provided random account.
func sendMsgRemovePostReaction(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgRemovePostReaction, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
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

// randomReactionFields returns the data used to create a MsgAddReaction message
func randomRemovePostReactionFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, ak authkeeper.AccountKeeper,
) (*PostReactionData, bool) {
	posts := k.GetPosts(ctx)
	if len(posts) == 0 {
		// Skip cause there are no posts
		return nil, true
	}

	post, _ := RandomPost(r, posts)
	reactions := k.GetPostReactions(ctx, post.PostID)

	// Skip if the post has no reactions
	if len(reactions) == 0 {
		return nil, true
	}

	reaction := reactions[r.Intn(len(reactions))]
	addr, _ := sdk.AccAddressFromBech32(reaction.Owner)
	acc := ak.GetAccount(ctx, addr)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return nil, true
	}

	user := GetAccount(addr, accs)
	data := PostReactionData{Shortcode: reaction.ShortCode, Value: reaction.Value, User: *user, PostID: post.PostID}
	return &data, false
}

// ___________________________________________________________________________________________________________________

// SimulateMsgRegisterReaction tests and runs a single msg register reaction where the registering user account already exist
// nolint: funlen
func SimulateMsgRegisterReaction(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		reactionData, skip := randomRegisteredReactionFields(r, ctx, accs, k, ak)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgRegisterReaction"), nil, nil
		}

		msg := types.NewMsgRegisterReaction(
			reactionData.Creator.Address.String(),
			reactionData.ShortCode,
			reactionData.Value,
			reactionData.Subspace,
		)

		err := sendMsgRegisterReaction(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{reactionData.Creator.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgRegisterReaction"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgRegisterReaction"), nil, nil
	}
}

// sendMsgRegisterReaction sends a transaction with a MsgRegisterReaction from a provided random account.
func sendMsgRegisterReaction(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgRegisterReaction, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
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

// randomRegisteredReactionFields returns the data used to create a MsgRegisterReaction message
func randomRegisteredReactionFields(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, ak authkeeper.AccountKeeper,
) (*ReactionData, bool) {
	reactionData := RandomReactionData(r, accs)
	acc := ak.GetAccount(ctx, reactionData.Creator.Address)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return nil, true
	}

	// Skip if the reaction already exists
	_, registered := k.GetRegisteredReaction(ctx, reactionData.ShortCode, reactionData.Subspace)
	if registered {
		return nil, true
	}

	return &reactionData, false
}
