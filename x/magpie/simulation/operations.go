package simulation

// DONTCOVER

import (
	"math/rand"

	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/tendermint/tendermint/crypto"

	"github.com/desmos-labs/desmos/app/params"
	"github.com/desmos-labs/desmos/x/magpie/types"
)

// Simulation operation weights constants
const (
	OpWeightMsgCreatePost = "op_weight_msg_create_session"
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONMarshaler, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) sim.WeightedOperations {

	var weightMsgCreatePost int
	appParams.GetOrGenerate(cdc, OpWeightMsgCreatePost, &weightMsgCreatePost, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePost = params.DefaultWeightMsgCreatePost
		},
	)

	return sim.WeightedOperations{
		sim.NewWeightedOperation(
			weightMsgCreatePost,
			SimulateMsgCreateSession(ak, bk),
		),
	}
}

// SimulateMsgCreateSession tests and runs a single msg create session where the post creator
// account already exists
// nolint: funlen
func SimulateMsgCreateSession(ak authkeeper.AccountKeeper, bk bankkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		data, skip := randomSessionFields(r, ctx, accs, ak)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ActionCreationSession, ""), nil, nil
		}

		msg := types.NewMsgCreateSession(
			data.Owner.Address.String(),
			data.Namespace,
			data.ExternalOwner,
			data.PubKey,
			data.Signature,
		)

		err := sendMsgCreateSession(r, app, ak, bk, msg, ctx, chainID, []crypto.PrivKey{data.Owner.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ActionCreationSession, ""), nil, nil
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgCreateSession sends a transaction with a MsgCreateSession from a provided random account.
func sendMsgCreateSession(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgCreateSession, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
) error {

	address, _ := sdk.AccAddressFromBech32(msg.Owner)
	account := ak.GetAccount(ctx, address)
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
		helpers.DefaultGenTxGas,
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

// randomSessionFields returns all the random fields that are needed to create a MsgCreateSession
func randomSessionFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, ak authkeeper.AccountKeeper,
) (*SessionData, bool) {

	simAccount, _ := simtypes.RandomAcc(r, accs)
	acc := ak.GetAccount(ctx, simAccount.Address)
	if acc == nil {
		return nil, true
	}

	sessionData := RandomSessionData(simAccount, r)
	return &sessionData, false
}
