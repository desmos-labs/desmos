package simulation

// DONTCOVER

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/desmos-labs/desmos/v7/testutil/simtesting"

	subspacessim "github.com/desmos-labs/desmos/v7/x/subspaces/simulation"
	subspacestypes "github.com/desmos-labs/desmos/v7/x/subspaces/types"

	"github.com/desmos-labs/desmos/v7/x/reports/keeper"
	"github.com/desmos-labs/desmos/v7/x/reports/types"
)

// SimulateMsgSupportStandardReason tests and runs a single MsgSupportStandardReason
func SimulateMsgSupportStandardReason(
	k keeper.Keeper, sk types.SubspacesKeeper,
	ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, standardReasonID, signer, skip := randomSupportStandardReasonFields(r, ctx, accs, sk, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgSupportStandardReason", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgSupportStandardReason(subspaceID, standardReasonID, signer.Address.String())

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, signer)
	}
}

// randomSupportStandardReasonFields returns the data used to build a random MsgSupportStandardReason
func randomSupportStandardReasonFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, sk types.SubspacesKeeper, k keeper.Keeper,
) (subspaceID uint64, standardReasonID uint32, user simtypes.Account, skip bool) {
	// Get the user
	if len(accs) == 0 {
		// Skip because there are no accounts
		skip = true
		return
	}
	user, _ = simtypes.RandomAcc(r, accs)

	// Get a subspace id
	subspaces := sk.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := subspacessim.RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get a reason
	reasons := k.GetParams(ctx).StandardReasons
	if len(reasons) == 0 {
		// Skip because there are no standard reasons to support
		skip = true
		return
	}
	reason := RandomStandardReason(r, reasons)
	standardReasonID = reason.ID

	// Get a user
	users := sk.GetUsersWithRootPermissions(ctx, subspace.ID, subspacestypes.NewPermissions(types.PermissionManageReasons))
	acc := subspacessim.GetAccount(subspacessim.RandomAddress(r, users), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	user = *acc

	return subspaceID, standardReasonID, user, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgAddReason tests and runs a single MsgAddReason
func SimulateMsgAddReason(
	sk types.SubspacesKeeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		data, signer, skip := randomAddReasonFields(r, ctx, accs, sk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgAddReason", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgAddReason(data.SubspaceID, data.Title, data.Description, signer.Address.String())

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, signer)
	}
}

// randomAddReasonFields returns the data used to build a random MsgAddReason
func randomAddReasonFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, sk types.SubspacesKeeper,
) (data types.Reason, user simtypes.Account, skip bool) {
	// Get the user
	if len(accs) == 0 {
		// Skip because there are no accounts
		skip = true
		return
	}
	user, _ = simtypes.RandomAcc(r, accs)

	// Get a subspace id
	subspaces := sk.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := subspacessim.RandomSubspace(r, subspaces)
	subspaceID := subspace.ID

	// Generate a random reason
	reason := types.NewReason(
		subspaceID,
		0,
		GetRandomReasonTitle(r),
		GetRandomReasonDescription(r),
	)

	// Get a user
	users := sk.GetUsersWithRootPermissions(ctx, subspace.ID, subspacestypes.NewPermissions(types.PermissionManageReasons))
	acc := subspacessim.GetAccount(subspacessim.RandomAddress(r, users), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	user = *acc

	return reason, user, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgRemoveReason tests and runs a single MsgRemoveReason
func SimulateMsgRemoveReason(
	k keeper.Keeper, sk types.SubspacesKeeper,
	ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, reasonID, signer, skip := randomRemoveReason(r, ctx, accs, k, sk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgRemoveReason", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgRemoveReason(subspaceID, reasonID, signer.Address.String())

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, signer)
	}
}

// randomRemoveReason returns the data used to build a random MsgRemoveReason
func randomRemoveReason(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, sk types.SubspacesKeeper,
) (subspaceID uint64, reasonID uint32, user simtypes.Account, skip bool) {
	// Get the user
	if len(accs) == 0 {
		// Skip because there are no accounts
		skip = true
		return
	}
	user, _ = simtypes.RandomAcc(r, accs)

	// Get a subspace id
	subspaces := sk.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := subspacessim.RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get a random reason
	reasons := k.GetSubspaceReasons(ctx, subspaceID)
	if len(reasons) == 0 {
		// Skip because there are no reasons to delete
		skip = true
		return
	}
	reason := RandomReason(r, reasons)
	reasonID = reason.ID

	// Get a user
	users := sk.GetUsersWithRootPermissions(ctx, subspace.ID, subspacestypes.NewPermissions(types.PermissionManageReasons))
	acc := subspacessim.GetAccount(subspacessim.RandomAddress(r, users), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	user = *acc

	return subspaceID, reasonID, user, false
}
