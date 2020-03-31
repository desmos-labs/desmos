package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/desmos-labs/desmos/x/profile/internal/keeper"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
)

// SimulateMsgCreateAccount tests and runs a single msg create profile where the creator already exists
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

		msg := types.NewMsgCreateProfile(
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

// sendMsgCreateAccount sends a transaction with a MsgCreateProfile from a provided random profile.
func sendMsgCreateAccount(
	r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgCreateProfile, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
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

// randomCreateAccountFields returns random profile data
func randomCreateAccountFields(r *rand.Rand, ctx sdk.Context, accs []sim.Account, ak auth.AccountKeeper,
) (*AccountData, bool, error) {

	accountData := RandomAccountData(r, accs)
	acc := ak.GetAccount(ctx, accountData.Creator.Address)

	// Skip the operation without error as the profile is not valid
	if acc == nil {
		return nil, true, nil
	}

	return &accountData, false, nil
}

// SimulateMsgEditAccount tests and runs a single msg edit profile where the creator already exists
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

		msg := types.NewMsgEditProfile(
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

// sendMsgEditAccount sends a transaction with a MsgEditProfile from a provided random profile.
func sendMsgEditAccount(
	r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgEditProfile, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
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

// randomAccountEditFields returns random profile data
func randomAccountEditFields(
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper, ak auth.AccountKeeper,
) (sim.Account, types.Profile, bool, error) {
	account := RandomAccount(r, k.GetAccounts(ctx))
	acc := GetSimAccount(account.Creator, accs)

	// Skip the operation without error as the profile is not valid
	if acc == nil {
		return sim.Account{}, types.Profile{}, true, nil
	}

	return *acc, account, false, nil
}

// SimulateMsgDeleteAccount tests and runs a single msg delete profile where the creator already exists
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

		msg := types.NewMsgDeleteProfile(moniker, acc.Address)

		err = sendMsgDeleteAccount(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{acc.PrivKey})
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgDeleteAccount sends a transaction with a MsgDeleteProfile from a provided random profile.
func sendMsgDeleteAccount(
	r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgDeleteProfile, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
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

// randomAccountDeleteFields returns random profile data
func randomAccountDeleteFields(
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper, ak auth.AccountKeeper,
) (sim.Account, string, bool, error) {
	account := RandomAccount(r, k.GetAccounts(ctx))

	acc := GetSimAccount(account.Creator, accs)

	// Skip the operation without error as the profile is not valid
	if acc == nil {
		return sim.Account{}, "", true, nil
	}

	return *acc, account.Moniker, false, nil
}
