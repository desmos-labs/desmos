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

// SimulateMsgCreateSubspace tests and runs a single msg create subspace
// nolint: funlen
func SimulateMsgCreateSubspace(k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		data, skip := randomSubspaceCreateFields(r, ctx, accs, k, ak)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgCreateSubspace"), nil, nil
		}

		msg := types.NewMsgCreateSubspace(data.Subspace.ID, data.Subspace.Name, data.Subspace.Creator, data.Subspace.Type)

		err := sendMsgCreateSubspace(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{data.CreatorAccount.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgCreateSubspace"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgCreateSubspace"), nil, nil
	}
}

// sendMsgCreateSubspace sends a transaction with a MsgCreateSubspace from a provided random account.
func sendMsgCreateSubspace(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgCreateSubspace, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
) error {
	addr, _ := sdk.AccAddressFromBech32(msg.Creator)
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

// randomSubspaceCreateFields returns the subspace random fields
func randomSubspaceCreateFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, ak authkeeper.AccountKeeper,
) (*SubspaceData, bool) {

	subspaceData := RandomSubspaceData(r, accs)
	acc := ak.GetAccount(ctx, subspaceData.CreatorAccount.Address)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return nil, true
	}

	// Skip the operation if the subspace already exists
	if k.DoesSubspaceExist(ctx, subspaceData.Subspace.ID) {
		return nil, true
	}

	return &subspaceData, false
}

// ___________________________________________________________________________________________________________________

// SimulateMsgEditSubspace tests and runs a single msg edit subspace
// nolint: funlen
func SimulateMsgEditSubspace(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		account, id, newName, newOwner, newType, skip := randomEditSubspaceFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgEditSubspace"), nil, nil
		}

		msg := types.NewMsgEditSubspace(id, newOwner, newName, account.Address.String(), newType)

		err := sendMsgEditSubspace(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{account.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgEditSubspace"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgEditSubspace"), nil, nil
	}
}

// sendMsgEditSubspace sends a transaction with a MsgEditSubspace from a provided random account.
func sendMsgEditSubspace(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgEditSubspace, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
) error {
	addr, _ := sdk.AccAddressFromBech32(msg.Editor)
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

// randomEditSubspaceFields returns the data needed to edit a subspace
func randomEditSubspaceFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (simtypes.Account, string, string, string, types.SubspaceType, bool) {
	subspaces := k.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip cause there are no subspaces
		return simtypes.Account{}, "", "", "", types.SubspaceTypeUnspecified, true
	}

	subspace, _ := RandomSubspace(r, subspaces)
	addr, _ := sdk.AccAddressFromBech32(subspace.Owner)
	acc := GetAccount(addr, accs)

	// Skip the operation without error as the account is not valid
	if acc == nil {
		return simtypes.Account{}, "", "", "", types.SubspaceTypeUnspecified, true
	}

	randomOwner, _ := simtypes.RandomAcc(r, accs)

	// Skip the operation without error if the new owner is equal to the actual one
	if randomOwner.Address.String() == acc.Address.String() {
		return simtypes.Account{}, "", "", "", types.SubspaceTypeUnspecified, true
	}

	return *acc, subspace.ID, RandomName(r), randomOwner.Address.String(), RandomSubspaceType(r), false
}
