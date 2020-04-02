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

// SimulateMsgCreateProfile tests and runs a single msg create profile where the creator already exists
// nolint: funlen
func SimulateMsgCreateProfile(ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {
		data, skip, err := randomCreateProfileFields(r, ctx, accs, ak)
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

		err = sendMsgCreateProfile(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{data.Creator.PrivKey})
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgCreateProfile sends a transaction with a MsgCreateProfile from a provided random profile.
func sendMsgCreateProfile(
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

// randomCreateProfileFields returns random profile data
func randomCreateProfileFields(r *rand.Rand, ctx sdk.Context, accs []sim.Account, ak auth.AccountKeeper,
) (*AccountData, bool, error) {

	accountData := RandomAccountData(r, accs)
	acc := ak.GetAccount(ctx, accountData.Creator.Address)

	// Skip the operation without error as the profile is not valid
	if acc == nil {
		return nil, true, nil
	}

	return &accountData, false, nil
}

// SimulateMsgEditProfile tests and runs a single msg edit profile where the creator already exists
// nolint: funlen
func SimulateMsgEditProfile(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {

		acc, data, newMoniker, skip, err := randomProfileEditFields(r, ctx, accs, k, ak)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgEditProfile(
			data.Moniker,
			newMoniker,
			data.Name,
			data.Surname,
			data.Bio,
			data.Pictures,
			acc.Address,
		)

		err = sendMsgEditProfile(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{acc.PrivKey})
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgEditProfile sends a transaction with a MsgEditProfile from a provided random profile.
func sendMsgEditProfile(
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

// randomProfileEditFields returns random profile data
func randomProfileEditFields(
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper, ak auth.AccountKeeper,
) (sim.Account, types.Profile, string, bool, error) {
	if len(accs) == 0 {
		return sim.Account{}, types.Profile{}, "", true, nil
	}
	accounts := k.GetProfiles(ctx)
	if len(accounts) == 0 {
		return sim.Account{}, types.Profile{}, "", true, nil
	}
	account := RandomAccount(r, accounts)
	acc := GetSimAccount(account.Creator, accs)

	// Skip the operation without error as the profile is not valid
	if acc == nil {
		return sim.Account{}, types.Profile{}, "", true, nil
	}

	return *acc, account, RandomMoniker(r), false, nil
}

// SimulateMsgDeleteProfile tests and runs a single msg delete profile where the creator already exists
// nolint: funlen
func SimulateMsgDeleteProfile(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {
		acc, moniker, skip, err := randomProfileDeleteFields(r, ctx, accs, k, ak)
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgDeleteProfile(moniker, acc.Address)

		err = sendMsgDeleteProfile(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{acc.PrivKey})
		if err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgDeleteProfile sends a transaction with a MsgDeleteProfile from a provided random profile.
func sendMsgDeleteProfile(
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

// randomProfileDeleteFields returns random profile data
func randomProfileDeleteFields(
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper, ak auth.AccountKeeper,
) (sim.Account, string, bool, error) {

	accounts := k.GetProfiles(ctx)

	if len(accounts) == 0 {
		return sim.Account{}, "", true, nil
	}
	account := RandomAccount(r, accounts)

	acc := GetSimAccount(account.Creator, accs)

	// Skip the operation without error as the profile is not valid
	if acc == nil {
		return sim.Account{}, "", true, nil
	}

	return *acc, account.Moniker, false, nil
}
