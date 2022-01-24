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

// SimulateMsgCreateUserGroup tests and runs a single MsgCreateUserGroup
func SimulateMsgCreateUserGroup(k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, groupName, permissions, creator, skip := randomCreateUserGroupFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgCreateUserGroup"), nil, nil
		}

		// Build the message
		msg := types.NewMsgCreateUserGroup(subspaceID, groupName, permissions, creator.Address.String())

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{creator.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgCreateUserGroup"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgCreateUserGroup", nil), nil, nil
	}
}

// randomCreateUserGroupFields returns the data used to build a random MsgCreateUserGroup
func randomCreateUserGroupFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (subspaceID uint64, groupName string, permissions types.Permission, account simtypes.Account, skip bool) {
	// Get a subspace id
	subspaces := k.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace, _ := RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get a group name
	groupName = RandomName(r)
	if k.HasUserGroup(ctx, subspaceID, groupName) {
		// Skip if the group already exists
		skip = true
		return
	}

	// Get a default permission
	permissions = RandomPermission(r, []types.Permission{
		types.PermissionWrite,
		types.PermissionChangeInfo,
		types.PermissionEverything,
	})

	// Get a signer
	signers, _ := k.GetUsersWithPermission(ctx, subspace.ID, types.PermissionManageGroups)
	acc := GetAccount(RandomAddress(r, signers), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	account = *acc

	return subspaceID, groupName, permissions, account, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgDeleteUserGroup tests and runs a single MsgDeleteUserGroup
func SimulateMsgDeleteUserGroup(k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, groupName, signer, skip := randomDeleteUserGroupFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgDeleteUserGroup"), nil, nil
		}

		// Build the message
		msg := types.NewMsgDeleteUserGroup(subspaceID, groupName, signer.Address.String())

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{signer.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgDeleteUserGroup"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgDeleteUserGroup", nil), nil, nil
	}
}

// randomDeleteUserGroupFields returns the data used to build a random MsgDeleteUserGroup
func randomDeleteUserGroupFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (subspaceID uint64, groupName string, account simtypes.Account, skip bool) {
	// Get a subspace id
	subspaces := k.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace, _ := RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get a group name
	groups := k.GetSubspaceGroups(ctx, subspace.ID)
	groupName = RandomString(r, groups)

	// Get a signer
	signers, _ := k.GetUsersWithPermission(ctx, subspace.ID, types.PermissionManageGroups)
	acc := GetAccount(RandomAddress(r, signers), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	account = *acc

	return subspaceID, groupName, account, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgAddUserToUserGroup tests and runs a single MsgAddUserToUserGroup
func SimulateMsgAddUserToUserGroup(k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, groupName, user, signer, skip := randomAddUserToUserGroupFields(r, ctx, accs, k, ak)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgAddUserToUserGroup"), nil, nil
		}

		// Build the message
		msg := types.NewMsgAddUserToUserGroup(subspaceID, groupName, user, signer.Address.String())

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{signer.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgAddUserToUserGroup"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgAddUserToUserGroup", nil), nil, nil
	}
}

// randomAddUserToUserGroupFields returns the data used to build a random MsgAddUserToUserGroup
func randomAddUserToUserGroupFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, ak authkeeper.AccountKeeper,
) (subspaceID uint64, groupName string, user string, account simtypes.Account, skip bool) {
	// Get a subspace id
	subspaces := k.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace, _ := RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get a group
	groups := k.GetSubspaceGroups(ctx, subspace.ID)
	groupName = RandomString(r, groups)

	// Get a user
	accounts := ak.GetAllAccounts(ctx)
	userAccount := RandomAuthAccount(r, accounts)
	if k.IsMemberOfGroup(ctx, subspace.ID, groupName, userAccount.GetAddress()) {
		// Skip if the user is already part of group
		skip = true
		return
	}
	user = userAccount.GetAddress().String()

	// Get a signer
	signers, _ := k.GetUsersWithPermission(ctx, subspace.ID, types.PermissionSetPermissions)
	acc := GetAccount(RandomAddress(r, signers), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	account = *acc

	return subspaceID, groupName, user, account, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgRemoveUserFromUserGroup tests and runs a single MsgRemoveUserFromUserGroup
func SimulateMsgRemoveUserFromUserGroup(k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, groupName, user, signer, skip := randomRemoveUserFromUserGroupFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgRemoveUserFromUserGroup"), nil, nil
		}

		// Build the message
		msg := types.NewMsgRemoveUserFromUserGroup(subspaceID, groupName, user, signer.Address.String())

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{signer.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgRemoveUserFromUserGroup"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgRemoveUserFromUserGroup", nil), nil, nil
	}
}

// randomRemoveUserFromUserGroupFields returns the data used to build a random MsgRemoveUserFromUserGroup
func randomRemoveUserFromUserGroupFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (subspaceID uint64, groupName string, user string, account simtypes.Account, skip bool) {
	// Get a subspace id
	subspaces := k.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace, _ := RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get a group
	groups := k.GetSubspaceGroups(ctx, subspace.ID)
	groupName = RandomString(r, groups)

	// Get a user
	members := k.GetGroupMembers(ctx, subspace.ID, groupName)
	user = RandomAddress(r, members).String()

	// Get a signer
	signers, _ := k.GetUsersWithPermission(ctx, subspace.ID, types.PermissionSetPermissions)
	acc := GetAccount(RandomAddress(r, signers), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	account = *acc

	return subspaceID, groupName, user, account, false
}
