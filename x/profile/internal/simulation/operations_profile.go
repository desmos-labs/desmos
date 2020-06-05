package simulation

import (
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
	"github.com/desmos-labs/desmos/x/profile/internal/types/msgs"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/desmos-labs/desmos/x/profile/internal/keeper"
	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
)

// SimulateMsgSaveProfile tests and runs a single msg save profile where the creator already exists
// nolint: funlen
func SimulateMsgSaveProfile(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {

		acc, data, newMoniker, skip, err := randomProfileSaveFields(r, ctx, accs, k)
		if err != nil {
			return sim.NoOpMsg(models.ModuleName), nil, err
		}

		if skip {
			return sim.NoOpMsg(models.ModuleName), nil, nil
		}

		msg := msgs.NewMsgSaveProfile(
			newMoniker,
			data.Name,
			data.Surname,
			data.Bio,
			nil,
			nil,
			acc.Address,
		)

		err = sendMsgSaveProfile(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{acc.PrivKey})
		if err != nil {
			return sim.NoOpMsg(models.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgSaveProfile sends a transaction with a MsgSaveProfile from a provided random profile.
func sendMsgSaveProfile(
	r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg msgs.MsgSaveProfile, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
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

// randomProfileSaveFields returns random profile data
func randomProfileSaveFields(
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper,
) (sim.Account, models.Profile, string, bool, error) {
	if len(accs) == 0 {
		return sim.Account{}, models.Profile{}, "", true, nil
	}
	accounts := k.GetProfiles(ctx)
	if len(accounts) == 0 {
		return sim.Account{}, models.Profile{}, "", true, nil
	}
	account := RandomProfile(r, accounts)
	acc := GetSimAccount(account.Creator, accs)

	// Skip the operation without error as the profile is not valid
	if acc == nil {
		return sim.Account{}, models.Profile{}, "", true, nil
	}

	return *acc, account, RandomMoniker(r), false, nil
}

// SimulateMsgDeleteProfile tests and runs a single msg delete profile where the creator already exists
// nolint: funlen
func SimulateMsgDeleteProfile(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {
		acc, skip, err := randomProfileDeleteFields(r, ctx, accs, k, ak)
		if err != nil {
			return sim.NoOpMsg(models.ModuleName), nil, err
		}

		if skip {
			return sim.NoOpMsg(models.ModuleName), nil, nil
		}

		msg := msgs.NewMsgDeleteProfile(acc.Address)

		err = sendMsgDeleteProfile(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{acc.PrivKey})
		if err != nil {
			return sim.NoOpMsg(models.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgDeleteProfile sends a transaction with a MsgDeleteProfile from a provided random profile.
func sendMsgDeleteProfile(
	r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg msgs.MsgDeleteProfile, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
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
) (sim.Account, bool, error) {
	if len(accs) == 0 {
		return sim.Account{}, true, nil
	}

	accounts := k.GetProfiles(ctx)

	if len(accounts) == 0 {
		return sim.Account{}, true, nil
	}
	account := RandomProfile(r, accounts)

	acc := GetSimAccount(account.Creator, accs)

	// Skip the operation without error as the profile is not valid
	if acc == nil {
		return sim.Account{}, true, nil
	}

	return *acc, false, nil
}
