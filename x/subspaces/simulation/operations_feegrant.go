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

// SimulateMsgGrantAllowance tests and runs a single MsgGrantAllowance
func SimulateMsgGrantAllowance(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Get the data
		subspaceID, granter, grantee, signer, skip := randomGrantAllowanceFields(r, ctx, accs, k, ak)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgGrantAllowance"), nil, nil
		}

		// Build the message
		msg := types.NewMsgGrantAllowance(subspaceID, granter, grantee, &feegrant.BasicAllowance{})

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{signer.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgGrantAllowance"), nil, err
		}
		return simtypes.NewOperationMsg(msg, true, "MsgGrantAllowance", nil), nil, nil
	}
}

// randomGrantAllowanceFields returns the data used to build a random MsgGrantAllowance
func randomGrantAllowanceFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, ak authkeeper.AccountKeeper,
) (subspaceID uint64, granter string, grantee types.Grantee, signer simtypes.Account, skip bool) {
	// Get a subspace id
	subspaces := k.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}

	subspaceID = RandomSubspace(r, subspaces).ID

	// Get a granter
	granters := k.GetUsersWithRootPermissions(ctx, subspaceID, types.NewPermissions(types.PermissionManageTreasuryAuthorization))
	granter = RandomAddress(r, granters)

	// 50% of having a user grantee, otherwise a group grantee
	if r.Intn(100) < 50 {
		accounts := ak.GetAllAccounts(ctx)
		userAddr := RandomAuthAccount(r, accounts).GetAddress().String()
		if granter == userAddr {
			skip = true
			return
		}
		grantee = types.NewUserGrantee(userAddr)
	} else {
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
		grantee = types.NewGroupGrantee(group.ID)
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

// SimulateMsgRevokeAllowance tests and runs a single MsgRevokeAllowance
func SimulateMsgRevokeAllowance(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// Get the data
		subspaceID, granter, grantee, signer, skip := randomRevokeAllowanceFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgRevokeAllowance"), nil, nil
		}

		// Build the message
		msg := types.NewMsgRevokeAllowance(subspaceID, granter, grantee)

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{signer.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgRevokeAllowance"), nil, err
		}
		return simtypes.NewOperationMsg(msg, true, "MsgRevokeAllowance", nil), nil, nil
	}
}

// randomRevokeAllowanceFields returns the data used to build a random MsgRevokeAllowance
func randomRevokeAllowanceFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (subspaceID uint64, granter string, grantee types.Grantee, signer simtypes.Account, skip bool) {
	// 50% of having user grants, otherwise group grants
	var grants []types.Grant
	if r.Intn(100) < 50 {
		grants = k.GetAllUserGrants(ctx)
	} else {
		grants = k.GetAllUserGroupsGrants(ctx)
	}

	if len(grants) == 0 {
		// Skip because there are no grants
		skip = true
		return
	}
	grant := RandomGrant(r, grants)
	subspaceID = grant.SubspaceID
	grantee = grant.Grantee.GetCachedValue().(types.Grantee)

	granters := k.GetUsersWithRootPermissions(ctx, subspaceID, types.NewPermissions(types.PermissionManageTreasuryAuthorization))
	granter = RandomAddress(r, granters)

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
