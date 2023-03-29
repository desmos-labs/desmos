package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/desmos-labs/desmos/v4/testutil/simtesting"
	feeskeeper "github.com/desmos-labs/desmos/v4/x/fees/keeper"
	"github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

// DONTCOVER

// SimulateMsgGrantTreasuryAuthorization tests and runs a single MsgGrantTreasuryAuthorization
func SimulateMsgGrantTreasuryAuthorization(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, grantee, granter, skip := randomGrantTreasuryAuthorizationFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgGrantTreasuryAuthorization", "skip"), nil, nil
		}

		// Build the message
		expiration := ctx.BlockTime().AddDate(1, 0, 0)
		msg := types.NewMsgGrantTreasuryAuthorization(subspaceID, granter.Address.String(), grantee, authz.NewGenericAuthorization(sdk.MsgTypeURL(&banktypes.MsgSend{})), &expiration)

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, granter)
	}
}

// randomGrantTreasuryAuthorizationFields returns the data used to build a random MsgGrantTreasuryAuthorization
func randomGrantTreasuryAuthorizationFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (subspaceID uint64, grantee string, granter simtypes.Account, skip bool) {
	// Get a subspace id
	subspaces := k.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get a grantee
	granteeAcc, _ := simtypes.RandomAcc(r, accs)
	grantee = granteeAcc.Address.String()

	// Get a granter
	granters := k.GetUsersWithRootPermissions(ctx, subspace.ID, types.NewPermissions(types.PermissionManageTreasuryAuthorization))
	acc := GetAccount(RandomAddress(r, granters), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	granter = *acc

	// Make sure the signer and the user are not the same
	if acc.Address.String() == grantee {
		skip = true
		return
	}

	return subspaceID, grantee, granter, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgRevokeTreasuryAuthorization tests and runs a single MsgRevokeTreasuryAuthorization
func SimulateMsgRevokeTreasuryAuthorization(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper, authzk authzkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, grantee, granter, msgTypeUrl, skip := randomRevokeTreasuryAuthorizationFields(r, ctx, accs, k, authzk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgRevokeTreasuryAuthorization", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgRevokeTreasuryAuthorization(subspaceID, granter.Address.String(), grantee, msgTypeUrl)

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, granter)
	}
}

// randomRevokeTreasuryAuthorizationFields returns the data used to build a random MsgRevokeTreasuryAuthorization
func randomRevokeTreasuryAuthorizationFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, authzk authzkeeper.Keeper,
) (subspaceID uint64, grantee string, granter simtypes.Account, msgTypeUrl string, skip bool) {
	// Get a subspace id
	subspaces := k.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get an authorization and a grantee
	var authorizations []authz.Authorization
	var grantees []string
	authzk.IterateGrants(ctx, func(granterAddr sdk.AccAddress, granteeAddr sdk.AccAddress, grant authz.Grant) bool {
		if granterAddr.String() == subspace.Treasury {
			authorization, err := grant.GetAuthorization()
			if err != nil {
				panic(err)
			}

			authorizations = append(authorizations, authorization)
			grantees = append(grantees, granteeAddr.String())
		}
		return false
	})
	if len(authorizations) == 0 {
		skip = true
		return
	}
	grantIndex := r.Intn(len(authorizations))
	authorization, grantee := authorizations[grantIndex], grantees[grantIndex]

	// Get a granter
	granters := k.GetUsersWithRootPermissions(ctx, subspace.ID, types.NewPermissions(types.PermissionManageTreasuryAuthorization))
	acc := GetAccount(RandomAddress(r, granters), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	granter = *acc

	return subspaceID, grantee, granter, authorization.MsgTypeURL(), false
}
