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
)

// SimulateMsgAddReaction tests and runs a single msg add reaction where the reacting user account already exists
// nolint: funlen
func SimulateMsgAddReaction(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (sim.OperationMsg, []sim.FutureOperation, error) {

		data, skip, err := randomAddReactionFields(r, ctx, accs, k, ak)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgAddPostReaction(data.PostID, data.Value, data.User.Address)
		err = sendMsgAddReaction(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{data.User.PrivKey})
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgAddReaction sends a transaction with a MsgAddReaction from a provided random account.
func sendMsgAddReaction(
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

// randomAddReactionFields returns the data used to create a MsgAddReaction message
func randomAddReactionFields(
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper, ak auth.AccountKeeper,
) (*ReactionData, bool, error) {

	reactionData := RandomReactionData(r, accs, k.GetPosts(ctx))
	acc := ak.GetAccount(ctx, reactionData.User.Address)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return nil, true, nil
	}

	// Skip if the reaction already exists
	reactions := k.GetPostReactions(ctx, reactionData.PostID)
	if reactions.ContainsReactionFrom(reactionData.User.Address, reactionData.Value) {
		return nil, true, nil
	}

	return &reactionData, false, nil
}

// SimulateMsgRemoveReaction tests and runs a single msg remove reaction where the reacting user account already exists
// nolint: funlen
func SimulateMsgRemoveReaction(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (sim.OperationMsg, []sim.FutureOperation, error) {

		data, skip, err := randomRemoveReactionFields(r, ctx, accs, k, ak)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgRemovePostReaction(data.PostID, data.User.Address, data.Value)
		err = sendMsgRemoveReaction(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{data.User.PrivKey})
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgRemoveReaction sends a transaction with a MsgRemoveReaction from a provided random account.
func sendMsgRemoveReaction(
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
func randomRemoveReactionFields(
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper, ak auth.AccountKeeper,
) (*ReactionData, bool, error) {

	post, _ := RandomPost(r, k.GetPosts(ctx))

	reactions := k.GetPostReactions(ctx, post.PostID)

	// Skip if the post has no reactions
	if len(reactions) == 0 {
		return nil, true, nil
	}

	reaction := reactions[r.Intn(len(reactions))]

	acc := ak.GetAccount(ctx, reaction.Owner)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return nil, true, nil
	}

	user := GetAccount(reaction.Owner, accs)
	data := ReactionData{Value: reaction.Value, User: *user, PostID: post.PostID}
	return &data, false, nil
}
