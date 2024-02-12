package simulation

// DONTCOVER

import (
	"math/rand"

	poststypes "github.com/desmos-labs/desmos/v7/x/posts/types"

	"github.com/desmos-labs/desmos/v7/testutil/simtesting"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/desmos-labs/desmos/v7/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v7/x/subspaces/types"
)

// SimulateMsgCreateUserGroup tests and runs a single MsgCreateUserGroup
func SimulateMsgCreateUserGroup(
	k *keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, update, permissions, creator, skip := randomCreateUserGroupFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgCreateUserGroup", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgCreateUserGroup(
			subspaceID,
			0,
			update.Name,
			update.Description,
			permissions,
			nil,
			creator.Address.String(),
		)

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, creator)
	}
}

// randomCreateUserGroupFields returns the data used to build a random MsgCreateUserGroup
func randomCreateUserGroupFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k *keeper.Keeper,
) (subspaceID uint64, update types.GroupUpdate, permissions types.Permissions, account simtypes.Account, skip bool) {
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
	permissions = RandomPermission(r, []types.Permissions{
		types.NewPermissions(poststypes.PermissionWrite),
		types.NewPermissions(types.PermissionEditSubspace),
	})

	// Get a signer
	signers := k.GetUsersWithRootPermissions(ctx, subspace.ID, types.CombinePermissions(types.PermissionManageGroups, types.PermissionSetPermissions))
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
	k *keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, groupID, update, signer, skip := randomEditUserGroupFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgEditUserGroup", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgEditUserGroup(subspaceID, groupID, update.Name, update.Description, signer.Address.String())

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, signer)
	}
}

// randomEditUserGroupFields returns the data used to build a random MsgEditUserGroup
func randomEditUserGroupFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k *keeper.Keeper,
) (subspaceID uint64, groupID uint32, update types.GroupUpdate, account simtypes.Account, skip bool) {
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
	signers := k.GetUsersWithRootPermissions(ctx, subspaceID, types.NewPermissions(types.PermissionManageGroups))
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

// SimulateMsgMoveUserGroup tests and runs a single MsgMoveUserGroup
func SimulateMsgMoveUserGroup(
	k *keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, groupID, newSectionID, signer, skip := randomMoveUserGroupFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgMoveUserGroup", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgMoveUserGroup(subspaceID, groupID, newSectionID, signer.Address.String())

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, signer)
	}
}

// randomMoveUserGroupFields returns the data used to build a random MsgMoveUserGroup
func randomMoveUserGroupFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k *keeper.Keeper,
) (subspaceID uint64, groupID uint32, newSectionID uint32, account simtypes.Account, skip bool) {
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
	groups := k.GetSubspaceUserGroups(ctx, subspaceID)
	if len(groups) == 0 {
		// Skip if there are no groups
		skip = true
		return
	}
	group := RandomGroup(r, groups)
	groupID = group.ID
	if groupID == 0 {
		// Skip because default group can not be moved
		skip = true
		return
	}

	// Get a section
	sections := k.GetSubspaceSections(ctx, subspaceID)
	section := RandomSection(r, sections)
	newSectionID = section.ID

	// Get a signer
	signers := k.GetUsersWithRootPermissions(ctx, subspace.ID, types.NewPermissions(types.PermissionEverything))
	acc := GetAccount(RandomAddress(r, signers), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	account = *acc

	return subspaceID, groupID, newSectionID, account, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgSetUserGroupPermissions tests and runs a single MsgSetUserGroupPermissions
func SimulateMsgSetUserGroupPermissions(
	k *keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, groupID, permissions, signer, skip := randomSetUserGroupPermissionsFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgSetUserGroupPermissions", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgSetUserGroupPermissions(subspaceID, groupID, permissions, signer.Address.String())

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, signer)
	}
}

// randomSetUserGroupPermissionsFields returns the data used to build a random MsgSetUserGroupPermissions
func randomSetUserGroupPermissionsFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k *keeper.Keeper,
) (subspaceID uint64, groupID uint32, permissions types.Permissions, account simtypes.Account, skip bool) {
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
	groups := k.GetSubspaceUserGroups(ctx, subspaceID)
	if len(groups) == 0 {
		// Skip if there are no groups
		skip = true
		return
	}
	groupID = RandomGroup(r, groups).ID

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

	// Make sure the user can change this group's validPermissions
	if subspace.Owner != account.Address.String() && k.IsMemberOfGroup(ctx, subspaceID, groupID, account.Address.String()) {
		// If the user is not the subspace owner and it's part of the user group they cannot edit the group validPermissions
		skip = true
		return
	}

	return subspaceID, groupID, permissions, account, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgDeleteUserGroup tests and runs a single MsgDeleteUserGroup
func SimulateMsgDeleteUserGroup(
	k *keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, groupID, signer, skip := randomDeleteUserGroupFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgDeleteUserGroup", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgDeleteUserGroup(subspaceID, groupID, signer.Address.String())

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, signer)
	}
}

// randomDeleteUserGroupFields returns the data used to build a random MsgDeleteUserGroup
func randomDeleteUserGroupFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k *keeper.Keeper,
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
	signers := k.GetUsersWithRootPermissions(ctx, subspaceID, types.NewPermissions(types.PermissionManageGroups))
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
	k *keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, groupID, user, signer, skip := randomAddUserToUserGroupFields(r, ctx, accs, k, ak)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgAddUserToUserGroup", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgAddUserToUserGroup(subspaceID, groupID, user, signer.Address.String())

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, signer)
	}
}

// randomAddUserToUserGroupFields returns the data used to build a random MsgAddUserToUserGroup
func randomAddUserToUserGroupFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k *keeper.Keeper, ak authkeeper.AccountKeeper,
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
	if k.IsMemberOfGroup(ctx, subspaceID, groupID, userAccount.GetAddress().String()) {
		// Skip if the user is already part of group
		skip = true
		return
	}
	user = userAccount.GetAddress().String()

	// Get a signer
	signers := k.GetUsersWithRootPermissions(ctx, subspaceID, types.NewPermissions(types.PermissionSetPermissions))
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
	k *keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, groupID, user, signer, skip := randomRemoveUserFromUserGroupFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgRemoveUserFromUserGroup", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgRemoveUserFromUserGroup(subspaceID, groupID, user, signer.Address.String())

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, signer)
	}
}

// randomRemoveUserFromUserGroupFields returns the data used to build a random MsgRemoveUserFromUserGroup
func randomRemoveUserFromUserGroupFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k *keeper.Keeper,
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
	members := k.GetUserGroupMembers(ctx, subspaceID, groupID)
	if len(members) == 0 {
		// Skip if there are no member groups to remove
		skip = true
		return
	}

	user = RandomAddress(r, members)

	// Get a signer
	signers := k.GetUsersWithRootPermissions(ctx, subspaceID, types.NewPermissions(types.PermissionSetPermissions))
	acc := GetAccount(RandomAddress(r, signers), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	account = *acc

	return subspaceID, groupID, user, account, false
}
