package simulation

// DONTCOVER

import (
	"math/rand"

	feeskeeper "github.com/desmos-labs/desmos/v4/x/fees/keeper"

	subspaceskeeper "github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	subspacessim "github.com/desmos-labs/desmos/v4/x/subspaces/simulation"

	"github.com/desmos-labs/desmos/v4/testutil/simtesting"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/relationships/keeper"
	"github.com/desmos-labs/desmos/v4/x/relationships/types"
)

// SimulateMsgBlockUser tests and runs a single MsgBlockUser
func SimulateMsgBlockUser(
	k keeper.Keeper, sk subspaceskeeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {
		acc, blocked, subspaceID, skip := randomUserBlocksFields(r, ctx, accs, k, sk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgBlockUser", "skip"), nil, nil
		}

		msg := types.NewMsgBlockUser(acc.Address.String(), blocked, "", subspaceID)
		return simtesting.SendMsg(r, app, ak, bk, fk, types.RouterKey, msg, ctx, acc)
	}
}

// randomUserBlocksFields returns the data used to build a random MsgBlockUser
func randomUserBlocksFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, sk subspaceskeeper.Keeper,
) (blocker simtypes.Account, blockedAddr string, subspaceID uint64, skip bool) {
	// Get a random blocker
	if len(accs) == 0 {
		skip = true
		return
	}
	blocker, _ = simtypes.RandomAcc(r, accs)
	blockerAddr := blocker.Address.String()

	// Get a random blocked account
	blockedAcc, _ := simtypes.RandomAcc(r, accs)
	blockedAddr = blockedAcc.Address.String()

	// Skip if the blocker and blocked user are equals
	if blockerAddr == blockedAddr {
		skip = true
		return
	}

	// Get a random subspace
	subspaces := sk.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := subspacessim.RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Skip if user block already exists
	if k.HasUserBlocked(ctx, blockerAddr, blockedAddr, subspaceID) {
		skip = true
		return
	}

	return blocker, blockedAddr, subspaceID, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgUnblockUser tests and runs a single MsgUnblockUser
func SimulateMsgUnblockUser(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {
		acc, blocked, subspaceID, skip := randomUnblockUserFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgUnblockUser", "skip"), nil, nil
		}

		msg := types.NewMsgUnblockUser(acc.Address.String(), blocked, subspaceID)
		return simtesting.SendMsg(r, app, ak, bk, fk, types.RouterKey, msg, ctx, acc)
	}
}

// randomUnblockUserFields returns the data used to build a random MsgUnblockUser
func randomUnblockUserFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (blocker simtypes.Account, blockedAddr string, subspaceID uint64, skip bool) {
	// Get a random blocker
	if len(accs) == 0 {
		skip = true
		return
	}
	blocker, _ = simtypes.RandomAcc(r, accs)
	blockerAddr := blocker.Address.String()

	// Get a random userBlock
	var userBlocks []types.UserBlock
	k.IterateUsersBlocks(ctx, func(_ int64, block types.UserBlock) (stop bool) {
		if block.Blocker == blockerAddr {
			userBlocks = append(userBlocks, block)
		}
		return false
	})
	if len(userBlocks) == 0 {
		// Skip because there are no blocks
		skip = true
		return
	}

	userBlock := RandomUserBlock(r, userBlocks)
	return blocker, userBlock.Blocked, userBlock.SubspaceID, false
}
