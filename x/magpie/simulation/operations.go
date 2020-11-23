package simulation

// DONTCOVER

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/fees"
	"github.com/tendermint/tendermint/crypto"

	"github.com/desmos-labs/desmos/app/params"
	"github.com/desmos-labs/desmos/x/magpie/keeper"
	"github.com/desmos-labs/desmos/x/magpie/types"
)

// Simulation operation weights constants
const (
	OpWeightMsgCreatePost = "op_weight_msg_create_session"
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(appParams sim.AppParams, cdc *codec.Codec, k keeper.Keeper, ak auth.AccountKeeper,
	fk fees.Keeper) sim.WeightedOperations {

	var weightMsgCreatePost int
	appParams.GetOrGenerate(cdc, OpWeightMsgCreatePost, &weightMsgCreatePost, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePost = params.DefaultWeightMsgCreatePost
		},
	)

	return sim.WeightedOperations{
		sim.NewWeightedOperation(
			weightMsgCreatePost,
			SimulateMsgCreateSession(ak, fk),
		),
	}
}

// SimulateMsgCreateSession tests and runs a single msg create session where the post creator
// account already exists
// nolint: funlen
func SimulateMsgCreateSession(ak auth.AccountKeeper, fk fees.Keeper) sim.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (sim.OperationMsg, []sim.FutureOperation, error) {

		data, skip := randomSessionFields(r, ctx, accs, ak)
		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgCreateSession(
			data.Owner.Address,
			data.Namespace,
			data.ExternalOwner,
			data.PubKey,
			data.Signature,
		)

		err := sendMsgCreateSession(r, app, ak, fk, msg, ctx, chainID, []crypto.PrivKey{data.Owner.PrivKey})
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgCreateSession sends a transaction with a MsgCreateSession from a provided random account.
func sendMsgCreateSession(
	r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper, fk fees.Keeper,
	msg types.MsgCreateSession, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
) error {
	account := ak.GetAccount(ctx, msg.Owner)
	coins := account.SpendableCoins(ctx.BlockTime())

	randFees, err := sim.RandomFees(r, ctx, coins)
	if err != nil {
		return err
	}

	msgs := []sdk.Msg{msg}

	if err := fk.CheckFees(ctx, authtypes.NewStdFee(helpers.DefaultGenTxGas, randFees), msgs); err != nil {
		return nil
	}

	tx := helpers.GenTx(
		msgs,
		randFees,
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

// randomSessionFields returns all the random fields that are needed to create a MsgCreateSession
func randomSessionFields(
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, ak auth.AccountKeeper,
) (*SessionData, bool) {

	simAccount, _ := sim.RandomAcc(r, accs)
	acc := ak.GetAccount(ctx, simAccount.Address)
	if acc == nil {
		return nil, true
	}

	sessionData := RandomSessionData(simAccount, r)
	return &sessionData, false
}
