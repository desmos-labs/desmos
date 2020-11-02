package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/relationships/keeper"
	"github.com/desmos-labs/desmos/x/relationships/types"
	"github.com/desmos-labs/desmos/x/relationships/types/msgs"
	"github.com/tendermint/tendermint/crypto"
)

//DONTCOVER

// SimulateMsgBlockUser tests and runs a single msg block user
// nolint: funlen
func SimulateMsgBlockUser(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {

		blocker, blocked, skip := randomUserBlocksFields(r, ctx, accs, k)
		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgBlockUser(blocker.Address, blocked, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")
		if err := sendMsgBlockUser(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{blocker.PrivKey}); err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgBlockUser sends a transaction with a MsgBlockUser from a provided random account
func sendMsgBlockUser(r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg msgs.MsgBlockUser, ctx sdk.Context, chainID string, privKeys []crypto.PrivKey) error {
	account := ak.GetAccount(ctx, msg.Blocker)
	coins := account.SpendableCoins(ctx.BlockTime())

	fees, err := sim.RandomFees(r, ctx, coins)
	if err != nil {
		return err
	}

	if fees.IsAllLT(minRequiredFee) {
		return nil
	}

	tx := helpers.GenTx(
		[]sdk.Msg{msg},
		fees,
		DefaultGasValue,
		chainID,
		[]uint64{account.GetAccountNumber()},
		[]uint64{account.GetSequence()},
		privKeys...,
	)

	_, _, err = app.Deliver(tx)
	if err != nil {
		return err
	}

	return nil
}

// randomUserBlocksFields returns random block user fields
func randomUserBlocksFields(r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper) (sim.Account, sdk.AccAddress, bool) {
	if len(accs) == 0 {
		return sim.Account{}, nil, true
	}

	// Get random accounts
	blocker, _ := sim.RandomAcc(r, accs)
	blocked, _ := sim.RandomAcc(r, accs)

	// skip if the two address are equals
	if blocker.Equals(blocked) {
		return sim.Account{}, nil, true
	}

	// skip if user block already exists
	userBlocks := k.GetUserBlocks(ctx, blocker.Address)
	for _, userBlock := range userBlocks {
		if userBlock.Blocked.Equals(blocked.Address) {
			return sim.Account{}, nil, true
		}
	}

	return blocker, blocked.Address, false
}

// SimulateMsgUnblockUser tests and runs a single msg unblock user
// nolint: funlen
func SimulateMsgUnblockUser(k keeper.Keeper, ak auth.AccountKeeper) sim.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []sim.Account, chainID string) (OperationMsg sim.OperationMsg, futureOps []sim.FutureOperation, err error) {

		blocker, userBlock, skip := randomUnblockUserFields(r, ctx, accs, k)
		if skip {
			return sim.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgUnblockUser(blocker.Address, userBlock.Blocked, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")
		if err := sendMsgUnblockUser(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{blocker.PrivKey}); err != nil {
			return sim.NoOpMsg(types.ModuleName), nil, err
		}

		return sim.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// randomUnblockUserFields returns random unblock user fields
func randomUnblockUserFields(r *rand.Rand, ctx sdk.Context, accs []sim.Account, k keeper.Keeper) (sim.Account, types.UserBlock, bool) {
	if len(accs) == 0 {
		return sim.Account{}, types.UserBlock{}, true
	}

	// Get random accounts
	user, _ := sim.RandomAcc(r, accs)

	userBlocks := k.GetUserBlocks(ctx, user.Address)

	// skip the test if the user has no userBlocks
	if len(userBlocks) == 0 {
		return sim.Account{}, types.UserBlock{}, true
	}

	return user, RandomUserBlock(r, userBlocks), false
}

// sendMsgUnblockUser sends a transaction with a MsgUnblockUser from a provided random account
func sendMsgUnblockUser(r *rand.Rand, app *baseapp.BaseApp, ak auth.AccountKeeper,
	msg msgs.MsgUnblockUser, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey) error {
	account := ak.GetAccount(ctx, msg.Blocker)
	coins := account.SpendableCoins(ctx.BlockTime())

	fees, err := sim.RandomFees(r, ctx, coins)
	if err != nil {
		return err
	}

	if fees.IsAllLT(minRequiredFee) {
		return nil
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
