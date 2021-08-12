package simulation

// DONTCOVER

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/desmos-labs/desmos/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/x/subspaces/types"
)

// SimulateMsgSaveTokenomics tests and runs a single msg save tokenomics
// nolint: funlen
func SimulateMsgSaveTokenomics(k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		data, skip := randomSaveTokenomicsFields(r, ctx, accs, k, ak)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgSaveTokenomics"), nil, nil
		}

		msg := types.NewMsgSaveTokenomics(
			data.Tokenomics.SubspaceID,
			data.Tokenomics.ContractAddress,
			data.Tokenomics.Admin,
			data.Tokenomics.Message,
		)

		err := sendMsgSaveTokenomics(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{data.AdminAccount.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgSaveTokenomics"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgSaveTokenomics"), nil, nil
	}
}

// sendMsgSaveTokenomics sends a transaction with a MsgSaveTokenomics from a provided random account.
func sendMsgSaveTokenomics(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgSaveTokenomics, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
) error {
	addr, _ := sdk.AccAddressFromBech32(msg.Admin)
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

// randomSaveTokenomicsFields returns the tokenomics random fields
func randomSaveTokenomicsFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, ak authkeeper.AccountKeeper,
) (*TokenomicsData, bool) {

	tokenomicsData := RandomTokenomicsData(r, accs)
	acc := ak.GetAccount(ctx, tokenomicsData.AdminAccount.Address)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return nil, true
	}

	// Skip the operation if admin and contract addresses are equals
	if tokenomicsData.Tokenomics.Admin == tokenomicsData.Tokenomics.ContractAddress {
		return nil, true
	}

	// Skip the operation if the subspace doesn't exists
	if !k.DoesSubspaceExist(ctx, tokenomicsData.Tokenomics.SubspaceID) {
		return nil, true
	}

	// skip if the user is not an admin
	if !k.IsAdmin(ctx, tokenomicsData.Tokenomics.SubspaceID, tokenomicsData.AdminAccount.Address.String()) {
		return nil, true
	}

	return &tokenomicsData, false
}
