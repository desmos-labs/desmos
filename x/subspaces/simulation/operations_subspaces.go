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

	"github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

// SimulateMsgCreateSubspace tests and runs a single MsgCreateSubspace
func SimulateMsgCreateSubspace(
	ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspace, creator, skip := randomSubspaceCreateFields(r, accs)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgCreateSubspace", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgCreateSubspace(
			subspace.Name,
			subspace.Description,
			subspace.Owner,
			creator.Address.String(),
		)

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, creator)
	}
}

// randomSubspaceCreateFields returns the data used to build a random MsgCreateSubspace
func randomSubspaceCreateFields(
	r *rand.Rand, accs []simtypes.Account,
) (subspace types.Subspace, creator simtypes.Account, skip bool) {
	// Get the creator
	if len(accs) == 0 {
		// Skip because there are no accounts
		skip = true
		return
	}
	creator, _ = simtypes.RandomAcc(r, accs)

	// Get the subspace data
	subspace = GenerateRandomSubspace(r, accs)

	// Get the creator
	return subspace, creator, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgEditSubspace tests and runs a single msg edit subspace
func SimulateMsgEditSubspace(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, update, editor, skip := randomEditSubspaceFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgEditSubspace", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgEditSubspace(
			subspaceID,
			update.Name,
			update.Description,
			update.Owner,
			editor.Address.String(),
		)

		// Send the data
		return simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, editor)
	}
}

// randomEditSubspaceFields returns the data needed to edit a subspace
func randomEditSubspaceFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (subspaceID uint64, update types.SubspaceUpdate, account simtypes.Account, skip bool) {
	// Get a subspace id
	subspaces := k.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get an editor
	editors := k.GetUsersWithRootPermissions(ctx, subspace.ID, types.NewPermissions(types.PermissionEditSubspace))
	acc := GetAccount(RandomAddress(r, editors), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	account = *acc

	// Build the update data
	update = types.NewSubspaceUpdate(
		RandomName(r),
		RandomDescription(r),
		account.Address.String(),
	)
	if r.Intn(101) < 50 {
		// 50% of chance of not editing the name
		update.Name = types.DoNotModify
	}
	if r.Intn(101) < 50 {
		// 50% of chance of not editing the description
		update.Description = types.DoNotModify
	}
	if r.Intn(101) < 50 {
		// 50% of chance of not editing the owner
		update.Owner = types.DoNotModify
	}

	return subspaceID, update, account, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgDeleteSubspace tests and runs a single msg delete subspace
func SimulateMsgDeleteSubspace(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, editor, skip := randomDeleteSubspaceFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgDeleteSubspace", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgDeleteSubspace(subspaceID, editor.Address.String())

		// Send the data
		return simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, editor)
	}
}

// randomDeleteSubspaceFields returns the data needed to delete a subspace
func randomDeleteSubspaceFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (subspaceID uint64, account simtypes.Account, skip bool) {
	// Get a subspace id
	subspaces := k.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get an editor
	editors := k.GetUsersWithRootPermissions(ctx, subspace.ID, types.NewPermissions(types.PermissionDeleteSubspace))
	acc := GetAccount(RandomAddress(r, editors), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	account = *acc

	return subspaceID, account, false
}
