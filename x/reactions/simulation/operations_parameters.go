package simulation

// DONTCOVER

import (
	"math/rand"

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

// SimulateMsgSetReactionsParams tests and runs a single MsgSetReactionsParams
func SimulateMsgSetReactionsParams(
	sk subspaceskeeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		data, signer, skip := randomSetReactionsParamsFields(r, ctx, accs, sk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgSetReactionsParams", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgSetReactionsParams(
			data.SubspaceID,
			data.RegisteredReaction,
			data.FreeText,
			signer.Address.String(),
		)

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, signer)
	}
}

// randomSetReactionsParamsFields returns the data used to build a random MsgSetReactionsParams
func randomSetReactionsParamsFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, sk subspaceskeeper.Keeper,
) (data types.SubspaceReactionsParams, user simtypes.Account, skip bool) {
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
	users := sk.GetUsersWithRootPermissions(ctx, subspace.ID, subspacestypes.NewPermissions(types.PermissionsReact))
	acc := subspacessim.GetAccount(subspacessim.RandomAddress(r, users), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	user = *acc

	// Generate a random reaction
	data = GenerateRandomSubspaceReactionsParams(r, subspaceID)
	return data, user, false
}
