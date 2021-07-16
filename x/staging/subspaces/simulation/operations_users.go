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

// SimulateMsgRegisterUser tests and runs a single msg register user
// nolint: funlen
func SimulateMsgRegisterUser(k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		owner, subspaceID, address, skip := randomRegisterUserFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgRegisterUser"), nil, nil
		}

		msg := types.NewMsgRegisterUser(
			subspaceID,
			address,
			owner.Address.String(),
		)
		err = sendMsgRegisterUser(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{owner.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.QuerierRoute, types.ModuleName, "MsgRegisterUser"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgRegisterUser"), nil, nil
	}
}

// sendMsgRegisterUser sends a transaction with a MsgRegisterUser from a provided random account
func sendMsgRegisterUser(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgRegisterUser, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
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

// randomRegisterUserFields returns random register user fields
func randomRegisterUserFields(
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

	// skip if the user is already registered
	if k.IsRegistered(ctx, subspace.ID, user.Address.String()) {
		return simtypes.Account{}, "", "", true
	}

	// Get the owner account
	acc, _ := sdk.AccAddressFromBech32(subspace.Owner)
	owner := GetAccount(acc, accs)

	return *owner, subspace.ID, user.Address.String(), false
}

// ___________________________________________________________________________________________________________________

// SimulateMsgUnregisterUser tests and runs a single msg unregister user
// nolint: funlen
func SimulateMsgUnregisterUser(k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		owner, subspaceID, address, skip := randomUnregisterUserFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgUnregisterUser"), nil, nil
		}

		msg := types.NewMsgUnregisterUser(
			subspaceID,
			address,
			owner.Address.String(),
		)
		err = sendMsgUnregisterUser(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{owner.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.QuerierRoute, types.ModuleName, "MsgUnregisterUser"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgUnregisterUser"), nil, nil
	}
}

// randomUnregisterUserFields returns random unregister user fields
func randomUnregisterUserFields(
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

	// skip if the user is not registered
	if !k.IsRegistered(ctx, subspace.ID, user.Address.String()) {
		return simtypes.Account{}, "", "", true
	}

	// Get the owner account
	acc, _ := sdk.AccAddressFromBech32(subspace.Owner)
	owner := GetAccount(acc, accs)

	return *owner, subspace.ID, user.Address.String(), false
}

// sendMsgUnregisterUser sends a transaction with a MsgUnregisterUser from a provided random account
func sendMsgUnregisterUser(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgUnregisterUser, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
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

// ___________________________________________________________________________________________________________________

// SimulateMsgBanUser tests and runs a single msg ban user
// nolint: funlen
func SimulateMsgBanUser(k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		owner, subspaceID, address, skip := randomBanUserFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgBanUser"), nil, nil
		}

		msg := types.NewMsgBanUser(
			subspaceID,
			address,
			owner.Address.String(),
		)
		err = sendMsgBanUser(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{owner.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.QuerierRoute, types.ModuleName, "MsgBanUser"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgBanUser"), nil, nil
	}
}

// sendMsgBanUser sends a transaction with a MsgBanUser from a provided random account
func sendMsgBanUser(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgBanUser, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
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

// randomBanUserFields returns random ban user fields
func randomBanUserFields(
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

	// skip if the user is already banned
	if k.IsBanned(ctx, subspace.ID, user.Address.String()) {
		return simtypes.Account{}, "", "", true
	}

	// Get the owner account
	acc, _ := sdk.AccAddressFromBech32(subspace.Owner)
	owner := GetAccount(acc, accs)

	return *owner, subspace.ID, user.Address.String(), false
}

// ___________________________________________________________________________________________________________________

// SimulateMsgUnbanUser tests and runs a single msg unban user
// nolint: funlen
func SimulateMsgUnbanUser(k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		owner, subspaceID, address, skip := randomUnbanUserFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgUnbanUser"), nil, nil
		}

		msg := types.NewMsgUnbanUser(
			subspaceID,
			address,
			owner.Address.String(),
		)
		err = sendMsgUnbanUser(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{owner.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.QuerierRoute, types.ModuleName, "MsgUnbanUser"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgUnbanUser"), nil, nil
	}
}

// randomUnbanUserFields returns random unban user fields
func randomUnbanUserFields(
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

	// skip if the user is not banned
	if !k.IsBanned(ctx, subspace.ID, user.Address.String()) {
		return simtypes.Account{}, "", "", true
	}

	// Get the owner account
	acc, _ := sdk.AccAddressFromBech32(subspace.Owner)
	owner := GetAccount(acc, accs)

	return *owner, subspace.ID, user.Address.String(), false
}

// sendMsgUnbanUser sends a transaction with a MsgUnbanUser from a provided random account
func sendMsgUnbanUser(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgUnbanUser, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
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
