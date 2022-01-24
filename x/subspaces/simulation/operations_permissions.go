package simulation

import (
	"math/rand"

	"github.com/desmos-labs/desmos/v2/testutil/simtesting"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/desmos-labs/desmos/v2/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

// SimulateMsgSetPermissions tests and runs a single MsgSetPermissions
func SimulateMsgSetPermissions(k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, target, permissions, creator, skip := randomSetPermissionsFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgSetPermissions"), nil, nil
		}

		// Build the message
		msg := types.NewMsgSetPermissions(subspaceID, target, permissions, creator.Address.String())

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{creator.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgSetPermissions"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgSetPermissions", nil), nil, nil
	}
}

// randomSetPermissionsFields returns the data used to build a random MsgSetPermissions
func randomSetPermissionsFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (subspaceID uint64, target string, permissions types.Permission, account simtypes.Account, skip bool) {
	// Get a subspace id
	subspaces := k.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace, _ := RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get a target
	targets := make([]string, len(accs))
	for i, acc := range accs {
		targets[i] = acc.Address.String()
	}
	targets = append(targets, k.GetSubspaceGroups(ctx, subspace.ID)...)
	target = RandomString(r, targets)

	// Get a permission
	permissions = RandomPermission(r, []types.Permission{
		types.PermissionWrite,
		types.PermissionModerateContent,
		types.PermissionChangeInfo,
		types.PermissionManageGroups,
	})

	// Get a signer
	signers, _ := k.GetUsersWithPermission(ctx, subspace.ID, types.PermissionSetPermissions)
	acc := GetAccount(RandomAddress(r, signers), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	account = *acc

	return subspaceID, target, permissions, account, false
}
