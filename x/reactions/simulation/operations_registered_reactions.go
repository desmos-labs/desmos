package simulation

// DONTCOVER

import (
	"math/rand"

	"github.com/desmos-labs/desmos/v4/x/reactions/keeper"

	feeskeeper "github.com/desmos-labs/desmos/v4/x/fees/keeper"
	subspaceskeeper "github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	subspacessim "github.com/desmos-labs/desmos/v4/x/subspaces/simulation"
	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"

	"github.com/desmos-labs/desmos/v4/testutil/simtesting"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/desmos-labs/desmos/v4/x/reactions/types"
)

// SimulateMsgAddRegisteredReaction tests and runs a single MsgAddRegisteredReaction
func SimulateMsgAddRegisteredReaction(
	sk subspaceskeeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		data, signer, skip := randomAddRegisteredReactionFields(r, ctx, accs, sk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgAddRegisteredReaction", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgAddRegisteredReaction(
			data.SubspaceID,
			data.ShorthandCode,
			data.DisplayValue,
			signer.Address.String(),
		)

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, fk, types.RouterKey, msg, ctx, signer)
	}
}

// randomAddRegisteredReactionFields returns the data used to build a random MsgAddRegisteredReaction
func randomAddRegisteredReactionFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, sk subspaceskeeper.Keeper,
) (reaction types.RegisteredReaction, user simtypes.Account, skip bool) {
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

	// Get a user
	users := sk.GetUsersWithRootPermissions(ctx, subspace.ID, subspacestypes.NewPermissions(types.PermissionManageRegisteredReactions))
	acc := subspacessim.GetAccount(subspacessim.RandomAddress(r, users), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	user = *acc

	// Generate a random reaction
	reaction = types.NewRegisteredReaction(
		subspaceID,
		0,
		GenerateRandomShorthandCode(r),
		GenerateRandomDisplayValue(r),
	)
	return reaction, user, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgEditRegisteredReaction tests and runs a single MsgEditRegisteredReaction
func SimulateMsgEditRegisteredReaction(
	k keeper.Keeper, sk subspaceskeeper.Keeper,
	ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		data, signer, skip := randomEditRegisteredReactionFields(r, ctx, accs, k, sk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgEditRegisteredReaction", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgEditRegisteredReaction(
			data.SubspaceID,
			data.ID,
			data.ShorthandCode,
			data.DisplayValue,
			signer.Address.String(),
		)

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, fk, types.RouterKey, msg, ctx, signer)
	}
}

// randomEditRegisteredReactionFields returns the data used to build a random MsgEditRegisteredReaction
func randomEditRegisteredReactionFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account,
	k keeper.Keeper, sk subspaceskeeper.Keeper,
) (reaction types.RegisteredReaction, user simtypes.Account, skip bool) {
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

	// Get a random reaction
	reactions := k.GetSubspaceRegisteredReactions(ctx, subspaceID)
	if len(reactions) == 0 {
		// Skip because there are no registered reactions
		skip = true
		return
	}
	existingReaction := RandomRegisteredReaction(r, reactions)

	// Generate new random data
	reaction = types.NewRegisteredReaction(
		existingReaction.SubspaceID,
		existingReaction.ID,
		GenerateRandomShorthandCode(r),
		GenerateRandomDisplayValue(r),
	)

	// Get a user
	users := sk.GetUsersWithRootPermissions(ctx, subspace.ID, subspacestypes.NewPermissions(types.PermissionManageRegisteredReactions))
	acc := subspacessim.GetAccount(subspacessim.RandomAddress(r, users), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	user = *acc

	return reaction, user, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgRemoveRegisteredReaction tests and runs a single MsgRemoveRegisteredReaction
func SimulateMsgRemoveRegisteredReaction(
	k keeper.Keeper, sk subspaceskeeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		data, signer, skip := randomRemoveRegisteredReactionFields(r, ctx, accs, k, sk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgRemoveRegisteredReaction", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgRemoveRegisteredReaction(
			data.SubspaceID,
			data.ID,
			signer.Address.String(),
		)

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, fk, types.RouterKey, msg, ctx, signer)
	}
}

// randomRemoveRegisteredReactionFields returns the data used to build a random MsgRemoveRegisteredReaction
func randomRemoveRegisteredReactionFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account,
	k keeper.Keeper, sk subspaceskeeper.Keeper,
) (reaction types.RegisteredReaction, user simtypes.Account, skip bool) {
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

	// Get a random reaction
	reactions := k.GetSubspaceRegisteredReactions(ctx, subspaceID)
	if len(reactions) == 0 {
		// Skip because there are no registered reactions
		skip = true
		return
	}
	reaction = RandomRegisteredReaction(r, reactions)

	// Get a user
	users := sk.GetUsersWithRootPermissions(ctx, subspace.ID, subspacestypes.NewPermissions(types.PermissionManageRegisteredReactions))
	acc := subspacessim.GetAccount(subspacessim.RandomAddress(r, users), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	user = *acc

	return reaction, user, false
}
