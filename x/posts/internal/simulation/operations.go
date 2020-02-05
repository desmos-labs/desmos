package simulation

import (
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/desmos-labs/desmos/app/params"
	"github.com/desmos-labs/desmos/x/posts/internal/keeper"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/tendermint/tendermint/crypto"

	sim "github.com/cosmos/cosmos-sdk/x/simulation"
)

// Simulation operation weights constants
const (
	OpWeightMsgCreatePost = "op_weight_msg_create_post"
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(appParams sim.AppParams, cdc *codec.Codec, k keeper.Keeper, ak auth.AccountKeeper) sim.WeightedOperations {

	var weightMsgSend int
	appParams.GetOrGenerate(cdc, OpWeightMsgCreatePost, &weightMsgSend, nil,
		func(_ *rand.Rand) {
			weightMsgSend = params.DefaultWeightMsgCreatePost
		},
	)

	return sim.WeightedOperations{
		sim.NewWeightedOperation(
			weightMsgSend,
			SimulateMsgCreatePost(k, ak),
		),
	}
}

// SimulateMsgCreatePost tests and runs a single msg create post where the post creator
// account already exists
// nolint: funlen
func SimulateMsgCreatePost(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (sim.OperationMsg, []sim.FutureOperation, error) {

		data, skip, err := randomPostFields(r, ctx, accs, k, ak)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

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
			data.CreationDate,
			data.Medias,
			data.PollData,
		)

		err = sendMsgCreatePost(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{data.Creator.PrivKey})
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
	coins := account.SpendableCoins(ctx.BlockTime())

	fees, err := sim.RandomFees(r, ctx, coins)
	if err != nil {
		return err
	}

	tx := helpers.GenTx(
		[]sdk.Msg{msg},
		fees,
		helpers.DefaultGenTxGas,
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

// postData contains the randomly generated data of a post
type postData struct {
	Creator        sim.Account
	ParentID       types.PostID
	Message        string
	AllowsComments bool
	Subspace       string
	CreationDate   time.Time
	OptionalData   map[string]string
	Medias         types.PostMedias
	PollData       *types.PollData
}

// randomPostFields returns the creator of the post as well as the parent id
func randomPostFields(
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper, ak auth.AccountKeeper,
) (*postData, bool, error) {

	simAccount, _ := sim.RandomAcc(r, accs)
	acc := ak.GetAccount(ctx, simAccount.Address)
	if acc == nil {
		return nil, true, nil // skip the operation without error as the account is not valid
	}

	pollData := RandomPollData(r)
	if pollData != nil && !pollData.Open {
		return nil, true, nil // skip the operation as the poll is closed
	}

	postData := postData{
		Creator:        simAccount,
		ParentID:       types.PostID(0),
		Message:        RandomMessage(r),
		AllowsComments: r.Intn(101) <= 50, // 50% chance of allowing comments
		Subspace:       RandomSubspace(r),
		CreationDate:   time.Now().UTC(),
		Medias:         RandomMedias(r),
		PollData:       pollData,
	}

	posts := k.GetPosts(ctx)
	if posts != nil {
		if parent, _ := RandomPost(r, posts); parent.AllowsComments {
			postData.ParentID = parent.PostID
		}
	}

	return &postData, false, nil
}
