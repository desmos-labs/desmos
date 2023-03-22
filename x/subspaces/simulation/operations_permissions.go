package simulation

// DONTCOVER

import (
	"math/rand"

	feeskeeper "github.com/desmos-labs/desmos/v4/x/fees/keeper"

	"github.com/desmos-labs/desmos/v4/testutil/simtesting"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

// SimulateMsgSetUserPermissions tests and runs a single MsgSetUserPermissions
func SimulateMsgSetUserPermissions(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, user, permissions, creator, skip := randomSetUserPermissionsFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.ModuleName, "MsgSetUserPermissions", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgSetUserPermissions(subspaceID, 0, user, permissions, creator.Address.String())

		// Send the message
		txCtx, err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, creator)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, "MsgSetUserPermissions", "invalid"), nil, nil
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// randomSetUserPermissionsFields returns the data used to build a random MsgSetUserPermissions
func randomSetUserPermissionsFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (subspaceID uint64, target string, permissions types.Permissions, account simtypes.Account, skip bool) {
	// Get a subspace id
	subspaces := k.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get a target
	targetAcc, _ := simtypes.RandomAcc(r, accs)
	target = targetAcc.Address.String()

	// Get a permission
	permissions = RandomPermission(r, validPermissions)

	// Get a signer
	signers := k.GetUsersWithRootPermissions(ctx, subspace.ID, types.NewPermissions(types.PermissionSetPermissions))
	acc := GetAccount(RandomAddress(r, signers), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	account = *acc

	// Make sure the signer and the user are not the same
	if acc.Address.String() == target {
		skip = true
		return
	}

	return subspaceID, target, permissions, account, false
}
