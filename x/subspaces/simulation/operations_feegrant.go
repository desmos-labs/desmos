package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/desmos-labs/desmos/v4/testutil/simtesting"
	feeskeeper "github.com/desmos-labs/desmos/v4/x/fees/keeper"
	"github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

// DONTCOVER

// SimulateMsgGrantUserAllowance tests and runs a single MsgGrantUserAllowance
func SimulateMsgGrantUserAllowance(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Get the data
		subspaceID, granter, grantee, signer, skip := randomGrantUserAllowanceFields(r, ctx, accs, k, ak)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgGrantUserAllowance"), nil, nil
		}

		// Build the message
		msg := types.NewMsgGrantUserAllowance(subspaceID, granter, grantee, &feegrant.BasicAllowance{})

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{signer.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgGrantUserAllowance"), nil, err
		}
		return simtypes.NewOperationMsg(nil, true, "MsgGrantUserAllowance", nil), nil, nil
	}
}

// randomGrantUserAllowanceFields returns the data used to build a random MsgGrantUserAllowance
func randomGrantUserAllowanceFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, ak authkeeper.AccountKeeper,
) (subspaceID uint64, granter string, grantee string, signer simtypes.Account, skip bool) {
	// Get a subspace id
	subspaces := k.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get a granter and grantee
	accounts := ak.GetAllAccounts(ctx)
	granter = RandomAuthAccount(r, accounts).GetAddress().String()
	grantee = RandomAuthAccount(r, accounts).GetAddress().String()
	if granter == grantee {
		skip = true
		return
	}

	// Get a signer account
	acc := GetAccount(granter, accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	signer = *acc

	return subspaceID, granter, grantee, signer, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgRevokeUserAllowance tests and runs a single MsgRevokeUserAllowance
func SimulateMsgRevokeUserAllowance(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Get the data
		subspaceID, granter, grantee, signer, skip := randomRevokeUserAllowanceFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgRevokeUserAllowance"), nil, nil
		}

		// Build the message
		msg := types.NewMsgRevokeUserAllowance(subspaceID, granter, grantee)

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{signer.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgRevokeUserAllowance"), nil, err
		}
		return simtypes.NewOperationMsg(nil, true, "MsgRevokeUserAllowance", nil), nil, nil
	}
}

// randomRevokeUserAllowanceFields returns the data used to build a random MsgRevokeUserAllowance
func randomRevokeUserAllowanceFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (subspaceID uint64, granter string, grantee string, signer simtypes.Account, skip bool) {
	// Get a grant
	grants := k.GetAllUserGrants(ctx)
	if len(grants) == 0 {
		// Skip if there are no grants
		skip = true
		return
	}
	grant := RandomUserGrant(r, grants)
	subspaceID = grant.SubspaceID
	granter = grant.Granter
	grantee = grant.Grantee

	// Get a signer account
	acc := GetAccount(granter, accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	signer = *acc

	return subspaceID, granter, grantee, signer, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgGrantGroupAllowance tests and runs a single MsgGrantGroupAllowance
func SimulateMsgGrantGroupAllowance(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Get the data
		subspaceID, granter, groupID, signer, skip := randomGrantGroupAllowanceFields(r, ctx, accs, k, ak)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgGrantGroupAllowance"), nil, nil
		}

		// Build the message
		msg := types.NewMsgGrantGroupAllowance(subspaceID, granter, groupID, &feegrant.BasicAllowance{})

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{signer.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgGrantGroupAllowance"), nil, err
		}
		return simtypes.NewOperationMsg(nil, true, "MsgGrantGroupAllowance", nil), nil, nil
	}
}

// randomGrantGroupAllowanceFields returns the data used to build a random MsgGrantGroupAllowance
func randomGrantGroupAllowanceFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, ak authkeeper.AccountKeeper,
) (subspaceID uint64, granter string, groupID uint32, signer simtypes.Account, skip bool) {
	// Get a group
	groups := k.GetAllUserGroups(ctx)
	if len(groups) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	group := RandomGroup(r, groups)
	if group.ID == 0 {
		// Skip because we cannot remove users from the group with ID 0 since it's the default one
		skip = true
		return
	}
	subspaceID = group.SubspaceID
	groupID = group.ID

	// Get a granter and grantee
	accounts := ak.GetAllAccounts(ctx)
	granter = RandomAuthAccount(r, accounts).GetAddress().String()

	// Get a signer account
	acc := GetAccount(granter, accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	signer = *acc

	return subspaceID, granter, groupID, signer, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgRevokeGroupAllowance tests and runs a single MsgRevokeGroupAllowance
func SimulateMsgRevokeGroupAllowance(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Get the data
		subspaceID, granter, groupID, signer, skip := randomRevokeGroupAllowanceFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgRevokeGroupAllowance"), nil, nil
		}

		// Build the message
		msg := types.NewMsgRevokeGroupAllowance(subspaceID, granter, groupID)

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{signer.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgRevokeGroupAllowance"), nil, err
		}
		return simtypes.NewOperationMsg(nil, true, "MsgRevokeGroupAllowance", nil), nil, nil
	}
}

// randomRevokeGroupAllowanceFields returns the data used to build a random MsgRevokeGroupAllowance
func randomRevokeGroupAllowanceFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (subspaceID uint64, granter string, groupID uint32, signer simtypes.Account, skip bool) {
	// Get a grant
	grants := k.GetAllGroupGrants(ctx)
	if len(grants) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	grant := RandomGroupGrant(r, grants)
	subspaceID = grant.SubspaceID
	granter = grant.Granter
	groupID = grant.GroupID

	// Get a signer account
	acc := GetAccount(granter, accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	signer = *acc

	return subspaceID, granter, groupID, signer, false
}
