package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/crypto"

	"github.com/desmos-labs/desmos/x/profiles/keeper"
	"github.com/desmos-labs/desmos/x/profiles/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
)

// SimulateMsgSaveProfile tests and runs a single msg save profile where the creator already exists
// nolint: funlen
func SimulateMsgSaveProfile(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {

		acc, data, skip := randomProfileSaveFields(r, ctx, accs, k)
		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		var profilePic, coverPic = "", ""
		if data.Pictures != nil {
			profilePic = *data.Pictures.Profile
			coverPic = *data.Pictures.Cover
		}

		msg := types.NewMsgSaveProfile(data.DTag, data.Moniker, data.Bio, &profilePic, &coverPic, acc.Address)
		if err := sendMsgSaveProfile(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{acc.PrivKey}); err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgSaveProfile sends a transaction with a MsgSaveProfile from a provided random profile.
func sendMsgSaveProfile(
	r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg types.MsgSaveProfile, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
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
) (sim.Account, types.Profile, bool) {
	if len(accs) == 0 {
		return sim.Account{}, types.Profile{}, true
	}

	// Get a random account
	account, _ := sim.RandomAcc(r, accs)

	// See if there is already the profile, otherwise create it from scratch
	var profile types.Profile
	existing, found := k.GetProfile(ctx, account.Address)
	if found {
		profile = existing
	} else {
		profile = NewRandomProfile(r, account.Address)
	}

	// 50% chance of changing something
	if r.Intn(101) <= 50 {
		profile = profile.
			WithMoniker(RandomMoniker(r)).
			WithBio(RandomBio(r)).
			WithPictures(RandomProfilePic(r), RandomProfileCover(r))
	}

	return account, profile, false
}

// SimulateMsgDeleteProfile tests and runs a single msg delete profile where the creator already exists
// nolint: funlen
func SimulateMsgDeleteProfile(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string,
	) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {
		acc, skip := randomProfileDeleteFields(r, ctx, accs, k, ak)
		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgDeleteProfile(acc.Address)

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
	r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper, _ auth.AccountKeeper,
) (sim.Account, bool) {
	if len(accs) == 0 {
		return sim.Account{}, true
	}

	accounts := k.GetProfiles(ctx)

	if len(accounts) == 0 {
		return sim.Account{}, true
	}
	account := RandomProfile(r, accounts)

	acc := GetSimAccount(account.Creator, accs)

	// Skip the operation without error as the profile is not valid
	if acc == nil {
		return sim.Account{}, true
	}

	return *acc, false
}
