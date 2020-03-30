package simulation

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/profile/internal/keeper"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/tendermint/tendermint/crypto"
	"math/rand"
)

// SimulateMsgCreateAccount tests and runs a single msg create account where the creator already exists
// nolint: funlen
func SimulateMsgCreateAccount(ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {
		data, skip, err := randomCreateAccountFields(r, ctx, accs, ak)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgCreateAccount(
			data.Name,
			data.Surname,
			data.Moniker,
			data.Bio,
			&data.Picture,
			data.Creator.Address,
		)

		err = sendMsgCreateAccount(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{data.Creator.PrivKey})
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgCreateAccount sends a transaction with a MsgCreateAccount from a provided random account.
func sendMsgCreateAccount(
	r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgCreateAccount, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
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

// randomCreateAccountFields returns random account data
func randomCreateAccountFields(r *rand.Rand, ctx sdk.Context, accs []sim.Account, ak auth.AccountKeeper,
) (*AccountData, bool, error) {

	accountData := RandomAccountData(r, accs)
	acc := ak.GetAccount(ctx, accountData.Creator.Address)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return nil, true, nil
	}

	return &accountData, false, nil
}

// SimulateMsgEditAccount tests and runs a single msg edit account where the creator already exists
// nolint: funlen
func SimulateMsgEditAccount(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {

		acc, data, skip, err := randomAccountEditFields(r, ctx, accs, k, ak)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgEditAccount(
			data.Name,
			data.Surname,
			data.Moniker,
			data.Bio,
			data.Pictures,
			acc.Address,
		)

		err = sendMsgEditAccount(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{acc.PrivKey})
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgEditAccount sends a transaction with a MsgEditAccount from a provided random account.
func sendMsgEditAccount(
	r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgEditAccount, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
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

// randomAccountEditFields returns random account data
func randomAccountEditFields(
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper, ak auth.AccountKeeper,
) (sim.Account, types.Account, bool, error) {
	account := RandomAccount(r, k.GetAccounts(ctx))
	acc := GetSimAccount(account.Creator, accs)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return sim.Account{}, types.Account{}, true, nil
	}

	return *acc, account, false, nil
}

// SimulateMsgDeleteAccount tests and runs a single msg delete account where the creator already exists
// nolint: funlen
func SimulateMsgDeleteAccount(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {
		acc, moniker, skip, err := randomAccountDeleteFields(r, ctx, accs, k, ak)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgDeleteAccount(moniker, acc.Address)

		err = sendMsgDeleteAccount(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{acc.PrivKey})
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgDeleteAccount sends a transaction with a MsgDeleteAccount from a provided random account.
func sendMsgDeleteAccount(
	r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgDeleteAccount, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
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

// randomAccountDeleteFields returns random account data
func randomAccountDeleteFields(
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper, ak auth.AccountKeeper,
) (sim.Account, string, bool, error) {
	account := RandomAccount(r, k.GetAccounts(ctx))

	acc := GetSimAccount(account.Creator, accs)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return sim.Account{}, "", true, nil
	}

	return *acc, account.Moniker, false, nil
}
