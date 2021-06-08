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

	"github.com/desmos-labs/desmos/x/staging/subspaces/keeper"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
)

// SimulateMsgAddAdmin tests and runs a single msg add admin
// nolint: funlen
func SimulateMsgAddAdmin(k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		owner, subspaceID, address, skip := randomAddAdminFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgAddAdmin"), nil, nil
		}

		msg := types.NewMsgAddAdmin(
			subspaceID,
			address,
			owner.Address.String(),
		)
		err = sendMsgAddAdmin(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{owner.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.QuerierRoute, types.ModuleName, "MsgAddAdmin"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgAddAdmin"), nil, nil
	}
}

// sendMsgAddAdmin sends a transaction with a MsgAddAdmin from a provided random account
func sendMsgAddAdmin(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgAddAdmin, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
) error {
	addr, _ := sdk.AccAddressFromBech32(msg.Owner)
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

// randomAddAdminFields returns random add admin fields
func randomAddAdminFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (simtypes.Account, string, string, bool) {
	if len(accs) == 0 {
		return simtypes.Account{}, "", "", true
	}

	subspaces := k.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip cause there are no subspaces
		return simtypes.Account{}, "", "", true
	}

	subspace, _ := RandomSubspace(r, subspaces)
	// Get random account
	user, _ := simtypes.RandomAcc(r, accs)

	// skip if the owner and the user are equals
	if subspace.Owner == user.Address.String() {
		return simtypes.Account{}, "", "", true
	}

	// skip if the user is already an admin
	if subspace.IsAdmin(user.Address.String()) {
		return simtypes.Account{}, "", "", true
	}

	// Get the owner account
	acc, _ := sdk.AccAddressFromBech32(subspace.Owner)
	owner := GetAccount(acc, accs)

	return *owner, subspace.ID, user.Address.String(), false
}

// ___________________________________________________________________________________________________________________

// SimulateMsgRemoveAdmin tests and runs a single msg remove admin
// nolint: funlen
func SimulateMsgRemoveAdmin(k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		owner, subspaceID, address, skip := randomRemoveAdminFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgRemoveAdmin"), nil, nil
		}

		msg := types.NewMsgRemoveAdmin(
			subspaceID,
			address,
			owner.Address.String(),
		)
		err = sendMsgRemoveAdmin(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{owner.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.QuerierRoute, types.ModuleName, "MsgRemoveAdmin"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgRemoveAdmin"), nil, nil
	}
}

// randomRemoveAdminFields returns random remove admin fields
func randomRemoveAdminFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (simtypes.Account, string, string, bool) {
	if len(accs) == 0 {
		return simtypes.Account{}, "", "", true
	}

	subspaces := k.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip cause there are no subspaces
		return simtypes.Account{}, "", "", true
	}

	subspace, _ := RandomSubspace(r, subspaces)
	// Get random account
	user, _ := simtypes.RandomAcc(r, accs)

	// skip if the owner and the user are equals
	if subspace.Owner == user.Address.String() {
		return simtypes.Account{}, "", "", true
	}

	// skip if the user is not an admin
	if !subspace.IsAdmin(user.Address.String()) {
		return simtypes.Account{}, "", "", true
	}

	// Get the owner account
	acc, _ := sdk.AccAddressFromBech32(subspace.Owner)
	owner := GetAccount(acc, accs)

	return *owner, subspace.ID, user.Address.String(), false
}

// sendMsgRemoveAdmin sends a transaction with a MsgRemoveAdmin from a provided random account
func sendMsgRemoveAdmin(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgRemoveAdmin, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
) error {
	addr, _ := sdk.AccAddressFromBech32(msg.Owner)
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
