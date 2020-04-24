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
	emoji "github.com/tmdvs/Go-Emoji-Utils"
)

// ---------------
// --- PostReaction
// ---------------

// SimulateMsgAddPostReaction tests and runs a single msg add reaction where the reacting user account already exists
// nolint: funlen
func SimulateMsgAddPostReaction(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (sim.OperationMsg, []sim.FutureOperation, error) {

		data, skip, err := randomAddPostReactionFields(r, ctx, accs, k, ak)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgAddPostReaction(data.PostID, data.Value, data.User.Address)
		err = sendMsgAddPostReaction(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{data.User.PrivKey})
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgAddPostReaction sends a transaction with a MsgAddReaction from a provided random account.
func sendMsgAddPostReaction(
	r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgAddPostReaction, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
) error {

	account := ak.GetAccount(ctx, msg.User)
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

// randomAddPostReactionFields returns the data used to create a MsgAddReaction message
func randomAddPostReactionFields(
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper, ak auth.AccountKeeper,
) (*PostReactionData, bool, error) {

	postData := RandomPostData(r, accs)
	postID := types.ComputeID(postData.CreationDate, postData.Creator.Address, postData.Subspace)

	post := types.NewPost(
		postID,
		postData.ParentID,
		postData.Message,
		postData.AllowsComments,
		postData.Subspace,
		postData.OptionalData,
		postData.CreationDate,
		postData.Creator.Address,
	)

	k.SavePost(ctx, post)

	data := RandomReactionData(r, accs)
	reaction := types.NewReaction(data.Creator.Address, data.ShortCode, data.Value, post.Subspace)
	k.RegisterReaction(ctx, reaction)

	reactionData := RandomPostReactionData(r, accs, postID, types.NewReactions(reaction))
	acc := ak.GetAccount(ctx, reactionData.User.Address)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return nil, true, nil
	}

	// Skip if the reaction already exists
	postReactions := k.GetPostReactions(ctx, reactionData.PostID)
	if postReactions.ContainsReactionFrom(reactionData.User.Address, reactionData.Value) {
		return nil, true, nil
	}

	return &reactionData, false, nil
}

// SimulateMsgRemovePostReaction tests and runs a single msg remove reaction where the reacting user account already exists
// nolint: funlen
func SimulateMsgRemovePostReaction(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (sim.OperationMsg, []sim.FutureOperation, error) {

		data, skip, err := randomRemovePostReactionFields(r, ctx, accs, k, ak)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgRemovePostReaction(data.PostID, data.User.Address, data.Value)
		err = sendMsgRemovePostReaction(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{data.User.PrivKey})
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgRemovePostReaction sends a transaction with a MsgRemoveReaction from a provided random account.
func sendMsgRemovePostReaction(
	r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgRemovePostReaction, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
) error {

	account := ak.GetAccount(ctx, msg.User)
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

// randomReactionFields returns the data used to create a MsgAddReaction message
func randomRemovePostReactionFields(
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper, ak auth.AccountKeeper,
) (*PostReactionData, bool, error) {

	post, _ := RandomPost(r, k.GetPosts(ctx))

	reactions := k.GetPostReactions(ctx, post.PostID)

	// Skip if the post has no reactions
	if len(reactions) == 0 {
		return nil, true, nil
	}

	reaction := reactions[r.Intn(len(reactions))]
	reactionValue := reaction.Value

	// 50 % of chance of using the shortcode, if it's a valid emoji
	if r.Intn(101) <= 50 {
		if e, err := emoji.LookupEmojiByCode(reactionValue); err == nil {
			reactionValue = e.Value
		}
	}

	acc := ak.GetAccount(ctx, reaction.Owner)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return nil, true, nil
	}

	user := GetAccount(reaction.Owner, accs)
	data := PostReactionData{Value: reactionValue, User: *user, PostID: post.PostID}
	return &data, false, nil
}

// ---------------
// --- Reaction
// ---------------

// SimulateMsgRegisterReaction tests and runs a single msg register reaction where the registering user account already exist
// nolint: funlen
func SimulateMsgRegisterReaction(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (sim.OperationMsg, []sim.FutureOperation, error) {
		reactionData, skip, err := randomRegisteredReactionFields(r, ctx, accs, k, ak)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgRegisterReaction(reactionData.Creator.Address,
			reactionData.ShortCode, reactionData.Value, reactionData.Subspace)

		err = sendMsgRegisterReaction(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{reactionData.Creator.PrivKey})
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgRegisterReaction sends a transaction with a MsgRegisterReaction from a provided random account.
func sendMsgRegisterReaction(r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgRegisterReaction, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
) error {
	account := ak.GetAccount(ctx, msg.Creator)
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

// randomRegisteredReactionFields returns the data used to create a MsgRegisterReaction message
func randomRegisteredReactionFields(r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper, ak auth.AccountKeeper,
) (*ReactionData, bool, error) {
	reactionData := RandomReactionData(r, accs)
	acc := ak.GetAccount(ctx, reactionData.Creator.Address)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return nil, true, nil
	}

	// Skip if the reaction already exists
	_, registered := k.GetRegisteredReaction(ctx, reactionData.ShortCode, reactionData.Subspace)
	if registered {
		return nil, true, nil
	}

	return &reactionData, false, nil
}
