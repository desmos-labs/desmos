package simulation

//DONTCOVER

import (
	"math/rand"

	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/relationships/keeper"
	"github.com/desmos-labs/desmos/x/relationships/types"
)

// SimulateMsgBlockUser tests and runs a single msg block user
// nolint: funlen
func SimulateMsgBlockUser(k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		blocker, blocked, skip := randomUserBlocksFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, ""), nil, nil
		}

		msg := types.NewMsgBlockUser(
			blocker.Address.String(),
			blocked.String(),
			"reason",
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		)
		err = sendMsgBlockUser(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{blocker.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.QuerierRoute, types.ModuleName, ""), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgBlockUser sends a transaction with a MsgBlockUser from a provided random account
func sendMsgBlockUser(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgBlockUser, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
) error {
	addr, _ := sdk.AccAddressFromBech32(msg.Blocker)
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

// randomUserBlocksFields returns random block user fields
func randomUserBlocksFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (simtypes.Account, sdk.AccAddress, bool) {
	if len(accs) == 0 {
		return simtypes.Account{}, nil, true
	}

	// Get random accounts
	blocker, _ := simtypes.RandomAcc(r, accs)
	blocked, _ := simtypes.RandomAcc(r, accs)

	// skip if the two address are equals
	if blocker.Equals(blocked) {
		return simtypes.Account{}, nil, true
	}

	// skip if user block already exists
	userBlocks := k.GetUserBlocks(ctx, blocker.Address.String())
	for _, userBlock := range userBlocks {
		if userBlock.Blocked == blocked.Address.String() {
			return simtypes.Account{}, nil, true
		}
	}

	return blocker, blocked.Address, false
}

// SimulateMsgUnblockUser tests and runs a single msg unblock user
// nolint: funlen
func SimulateMsgUnblockUser(k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		blocker, userBlock, skip := randomUnblockUserFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.QuerierRoute, types.ModuleName, ""), nil, nil
		}

		msg := types.NewMsgUnblockUser(
			blocker.Address.String(),
			userBlock.Blocked,
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		)
		if err := sendMsgUnblockUser(r, app, ak, bk, msg, ctx, chainID, []cryptotypes.PrivKey{blocker.PrivKey}); err != nil {
			return simtypes.NoOpMsg(types.QuerierRoute, types.ModuleName, ""), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// randomUnblockUserFields returns random unblock user fields
func randomUnblockUserFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (simtypes.Account, types.UserBlock, bool) {
	if len(accs) == 0 {
		return simtypes.Account{}, types.UserBlock{}, true
	}

	// Get random accounts
	user, _ := simtypes.RandomAcc(r, accs)
	userBlocks := k.GetUserBlocks(ctx, user.Address.String())

	// skip the test if the user has no userBlocks
	if len(userBlocks) == 0 {
		return simtypes.Account{}, types.UserBlock{}, true
	}

	return user, RandomUserBlock(r, userBlocks), false
}

// sendMsgUnblockUser sends a transaction with a MsgUnblockUser from a provided random account
func sendMsgUnblockUser(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	msg *types.MsgUnblockUser, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
) error {
	addr, _ := sdk.AccAddressFromBech32(msg.Blocker)
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
