package simulation

// DONTCOVER

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

// SimulateMsgCreateSubspace tests and runs a single MsgCreateSubspace
func SimulateMsgCreateSubspace(ak authkeeper.AccountKeeper, bk bankkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspace, creator, skip, err := randomSubspaceCreateFields(r, accs)
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgCreateSubspace"), nil, err
		}

		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgCreateSubspace"), nil, nil
		}

		// Build the message
		msg := types.NewMsgCreateSubspace(
			subspace.Name,
			subspace.Description,
			subspace.Treasury,
			subspace.Owner,
			subspace.Creator,
		)

		// Send the message
		err = simtesting.SendMsg(r, app, ak, bk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{creator.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgCreateSubspace"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgCreateSubspace", nil), nil, nil
	}
}

// randomSubspaceCreateFields returns the data used to build a random MsgCreateSubspace
func randomSubspaceCreateFields(
	r *rand.Rand, accs []simtypes.Account,
) (subspace types.Subspace, creator simtypes.Account, skip bool, err error) {
	// Get the subspace data
	subspace = GenerateRandomSubspace(r, accs)

	// Get the creator
	sdkAddr, err := sdk.AccAddressFromBech32(subspace.Creator)
	if err != nil {
		return
	}
	account := GetAccount(sdkAddr, accs)
	if account == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	creator = *account

	return
}

// ___________________________________________________________________________________________________________________

// SimulateMsgEditSubspace tests and runs a single msg edit subspace
func SimulateMsgEditSubspace(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, update, editor, skip := randomEditSubspaceFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgEditSubspace"), nil, nil
		}

		// Build the message
		msg := types.NewMsgEditSubspace(
			subspaceID,
			update.Name,
			update.Description,
			update.Treasury,
			update.Owner,
			editor.Address.String(),
		)

		// Send the data
		err := simtesting.SendMsg(r, app, ak, bk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{editor.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgEditSubspace"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgEditSubspace", nil), nil, nil
	}
}

// randomEditSubspaceFields returns the data needed to edit a subspace
func randomEditSubspaceFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (subspaceID uint64, update *types.SubspaceUpdate, account simtypes.Account, skip bool) {
	// Get a subspace id
	subspaces := k.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace, _ := RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Build the update data
	update = types.NewSubspaceUpdate(
		RandomName(r),
		RandomDescription(r),
		account.Address.String(),
		account.Address.String(),
	)

	// Get an editor
	editors, _ := k.GetUsersWithPermission(ctx, subspace.ID, types.PermissionChangeInfo)
	acc := GetAccount(RandomAddress(r, editors), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	account = *acc

	return subspaceID, update, account, false
}
