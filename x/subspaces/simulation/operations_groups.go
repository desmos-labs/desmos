package simulation

// DONTCOVER

import (
	"math/rand"

	feeskeeper "github.com/desmos-labs/desmos/v3/x/fees/keeper"

	"github.com/desmos-labs/desmos/v3/testutil/simtesting"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/desmos-labs/desmos/v3/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// SimulateMsgCreateUserGroup tests and runs a single MsgCreateUserGroup
func SimulateMsgCreateUserGroup(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, update, permissions, creator, skip := randomCreateUserGroupFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgCreateUserGroup"), nil, nil
		}

		// Build the message
		msg := types.NewMsgCreateUserGroup(
			subspaceID,
			update.Name,
			update.Description,
			permissions,
			creator.Address.String(),
		)

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{creator.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgCreateUserGroup"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgCreateUserGroup", nil), nil, nil
	}
}

// randomCreateUserGroupFields returns the data used to build a random MsgCreateUserGroup
func randomCreateUserGroupFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (subspaceID uint64, update *types.GroupUpdate, permissions types.Permission, account simtypes.Account, skip bool) {
	// Get a subspace id
	subspaces := k.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get a group name
	groupName := RandomName(r)
	groupDescription := RandomDescription(r)

	// Get a default permission
	permissions = RandomPermission(r, []types.Permission{
		types.PermissionWrite,
		types.PermissionChangeInfo,
		types.PermissionEverything,
	})

	// Get a signer
	signers, _ := k.GetUsersWithPermission(ctx, subspace.ID, types.CombinePermissions(types.PermissionManageGroups, types.PermissionSetPermissions))
	acc := GetAccount(RandomAddress(r, signers), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	account = *acc

	return subspaceID, types.NewGroupUpdate(groupName, groupDescription), permissions, account, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgEditUserGroup tests and runs a single MsgEditUserGroup
func SimulateMsgEditUserGroup(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, groupID, update, signer, skip := randomEditUserGroupFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgEditUserGroup"), nil, nil
		}

		// Build the message
		msg := types.NewMsgEditUserGroup(subspaceID, groupID, update.Name, update.Description, signer.Address.String())

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{signer.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgEditUserGroup"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgEditUserGroup", nil), nil, nil
	}
}

// randomEditUserGroupFields returns the data used to build a random MsgEditUserGroup
func randomEditUserGroupFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (subspaceID uint64, groupID uint32, update *types.GroupUpdate, account simtypes.Account, skip bool) {
	// Get a group
	groups := k.GetAllUserGroups(ctx)
	if len(groups) == 0 {
		// Skip if there are no groups
		skip = true
		return
	}
	group := RandomGroup(r, groups)
	subspaceID = group.SubspaceID
	groupID = group.ID

	// Build the update
	update = types.NewGroupUpdate(RandomName(r), RandomDescription(r))
	if r.Intn(101) < 50 {
		// 50% of chance of not editing the name
		update.Name = types.DoNotModify
	}
	if r.Intn(101) < 50 {
		// 50% of chance of not editing the description
		update.Description = types.DoNotModify
	}

	// Get a signer
	signers, _ := k.GetUsersWithPermission(ctx, subspaceID, types.PermissionManageGroups)
	acc := GetAccount(RandomAddress(r, signers), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	account = *acc

	return subspaceID, groupID, update, account, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgSetUserGroupPermissions tests and runs a single MsgSetUserGroupPermissions
func SimulateMsgSetUserGroupPermissions(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, groupID, permissions, signer, skip := randomSetUserGroupPermissionsFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgSetUserGroupPermissions"), nil, nil
		}

		// Build the message
		msg := types.NewMsgSetUserGroupPermissions(subspaceID, groupID, permissions, signer.Address.String())

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{signer.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgSetUserGroupPermissions"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgSetUserGroupPermissions", nil), nil, nil
	}
}

// randomSetUserGroupPermissionsFields returns the data used to build a random MsgSetUserGroupPermissions
func randomSetUserGroupPermissionsFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (subspaceID uint64, groupID uint32, permissions types.Permission, account simtypes.Account, skip bool) {
	// Get a subspace id
	subspaces := k.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get a group
	groups := k.GetSubspaceGroups(ctx, subspaceID)
	if len(groups) == 0 {
		// Skip if there are no groups
		skip = true
		return
	}
	groupID = RandomGroup(r, groups).ID

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

	// Make sure the user can change this group's permissions
	if subspace.Owner != account.Address.String() && k.IsMemberOfGroup(ctx, subspaceID, groupID, account.Address) {
		// If the user is not the subspace owner and it's part of the user group they cannot edit the group permissions
		skip = true
		return
	}

	return subspaceID, groupID, permissions, account, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgDeleteUserGroup tests and runs a single MsgDeleteUserGroup
func SimulateMsgDeleteUserGroup(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, groupID, signer, skip := randomDeleteUserGroupFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgDeleteUserGroup"), nil, nil
		}

		// Build the message
		msg := types.NewMsgDeleteUserGroup(subspaceID, groupID, signer.Address.String())

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{signer.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgDeleteUserGroup"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgDeleteUserGroup", nil), nil, nil
	}
}

// randomDeleteUserGroupFields returns the data used to build a random MsgDeleteUserGroup
func randomDeleteUserGroupFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (subspaceID uint64, groupID uint32, account simtypes.Account, skip bool) {
	// Get a group
	groups := k.GetAllUserGroups(ctx)
	if len(groups) == 0 {
		// Skip if there are no groups
		skip = true
		return
	}
	group := RandomGroup(r, groups)
	if group.ID == 0 {
		// Skip because we cannot delete the group with ID 0 since it's the default one
		skip = true
		return
	}

	subspaceID = group.SubspaceID
	groupID = group.ID

	// Get a signer
	signers, _ := k.GetUsersWithPermission(ctx, subspaceID, types.PermissionManageGroups)
	acc := GetAccount(RandomAddress(r, signers), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	account = *acc

	return subspaceID, groupID, account, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgAddUserToUserGroup tests and runs a single MsgAddUserToUserGroup
func SimulateMsgAddUserToUserGroup(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, groupID, user, signer, skip := randomAddUserToUserGroupFields(r, ctx, accs, k, ak)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgAddUserToUserGroup"), nil, nil
		}

		// Build the message
		msg := types.NewMsgAddUserToUserGroup(subspaceID, groupID, user, signer.Address.String())

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{signer.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgAddUserToUserGroup"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgAddUserToUserGroup", nil), nil, nil
	}
}

// randomAddUserToUserGroupFields returns the data used to build a random MsgAddUserToUserGroup
func randomAddUserToUserGroupFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, ak authkeeper.AccountKeeper,
) (subspaceID uint64, groupID uint32, user string, account simtypes.Account, skip bool) {
	// Get a group
	groups := k.GetAllUserGroups(ctx)
	if len(groups) == 0 {
		// Skip if there are no groups
		skip = true
		return
	}
	group := RandomGroup(r, groups)
	if group.ID == 0 {
		// Skip because we cannot add users to the group with ID 0 since it's the default one
		skip = true
		return
	}

	subspaceID = group.SubspaceID
	groupID = group.ID

	// Get a user
	accounts := ak.GetAllAccounts(ctx)
	userAccount := RandomAuthAccount(r, accounts)
	if k.IsMemberOfGroup(ctx, subspaceID, groupID, userAccount.GetAddress()) {
		// Skip if the user is already part of group
		skip = true
		return
	}
	user = userAccount.GetAddress().String()

	// Get a signer
	signers, _ := k.GetUsersWithPermission(ctx, subspaceID, types.PermissionSetPermissions)
	acc := GetAccount(RandomAddress(r, signers), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	account = *acc

	return subspaceID, groupID, user, account, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgRemoveUserFromUserGroup tests and runs a single MsgRemoveUserFromUserGroup
func SimulateMsgRemoveUserFromUserGroup(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, groupID, user, signer, skip := randomRemoveUserFromUserGroupFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgRemoveUserFromUserGroup"), nil, nil
		}

		// Build the message
		msg := types.NewMsgRemoveUserFromUserGroup(subspaceID, groupID, user, signer.Address.String())

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{signer.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgRemoveUserFromUserGroup"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgRemoveUserFromUserGroup", nil), nil, nil
	}
}

// randomRemoveUserFromUserGroupFields returns the data used to build a random MsgRemoveUserFromUserGroup
func randomRemoveUserFromUserGroupFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (subspaceID uint64, groupID uint32, user string, account simtypes.Account, skip bool) {
	// Get a group
	groups := k.GetAllUserGroups(ctx)
	if len(groups) == 0 {
		// Skip if there are no groups
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

	// Get a user
	members := k.GetGroupMembers(ctx, subspaceID, groupID)
	if len(members) == 0 {
		// Skip if there are no member groups to remove
		skip = true
		return
	}

	user = RandomAddress(r, members).String()

	// Get a signer
	signers, _ := k.GetUsersWithPermission(ctx, subspaceID, types.PermissionSetPermissions)
	acc := GetAccount(RandomAddress(r, signers), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	account = *acc

	return subspaceID, groupID, user, account, false
}
