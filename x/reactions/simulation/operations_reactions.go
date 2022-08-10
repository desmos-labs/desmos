package simulation

// DONTCOVER

import (
	"math/rand"

	postskeeper "github.com/desmos-labs/desmos/v4/x/posts/keeper"
	postssim "github.com/desmos-labs/desmos/v4/x/posts/simulation"
	"github.com/desmos-labs/desmos/v4/x/reactions/keeper"

	feeskeeper "github.com/desmos-labs/desmos/v4/x/fees/keeper"
	subspaceskeeper "github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	subspacessim "github.com/desmos-labs/desmos/v4/x/subspaces/simulation"
	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"

	"github.com/desmos-labs/desmos/v4/testutil/simtesting"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/desmos-labs/desmos/v4/x/reactions/types"
)

// SimulateMsgAddReaction tests and runs a single MsgAddReaction
func SimulateMsgAddReaction(
	k keeper.Keeper, sk subspaceskeeper.Keeper, pk postskeeper.Keeper,
	ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		data, signer, skip := randomAddReactionFields(r, ctx, accs, k, sk, pk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgAddReaction"), nil, nil
		}

		// Build the message
		msg := types.NewMsgAddReaction(
			data.SubspaceID,
			data.PostID,
			data.Value.GetCachedValue().(types.ReactionValue),
			signer.Address.String(),
		)

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{signer.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgAddReaction"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgAddReaction", nil), nil, nil
	}
}

// randomAddReactionFields returns the data used to build a random MsgAddReaction
func randomAddReactionFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account,
	k keeper.Keeper, sk subspaceskeeper.Keeper, pk postskeeper.Keeper,
) (reaction types.Reaction, user simtypes.Account, skip bool) {
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

	// Get the posts
	posts := pk.GetSubspacePosts(ctx, subspaceID)
	if len(posts) == 0 {
		// Skip because there are no posts
		skip = true
		return
	}
	post := postssim.RandomPost(r, posts)
	postID := post.ID

	// Get the subspaces params
	params, err := k.GetSubspaceReactionsParams(ctx, subspaceID)
	if err != nil {
		// Skip because there are some errors
		skip = true
		return
	}

	var value types.ReactionValue
	if r.Intn(101) < 50 {
		// 50% change of a RegisteredReactionValue
		if !params.RegisteredReaction.Enabled {
			// Skip because the registered reactions are not enabled
			skip = true
			return
		}

		registeredReactions := k.GetSubspaceRegisteredReactions(ctx, subspaceID)
		if len(registeredReactions) == 0 {
			// Skip because there are no registered reactions
			skip = true
			return
		}
		reaction := RandomRegisteredReaction(r, registeredReactions)
		value = types.NewRegisteredReactionValue(reaction.ID)
	} else {
		// 50% change of FreeTextValue
		if !params.FreeText.Enabled {
			// Skip because the free text reactions are not enabled
			skip = true
			return
		}

		value = types.NewFreeTextValue(GetRandomFreeTextValue(r, params.FreeText.MaxLength))
	}

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
	reaction = types.NewReaction(subspaceID, postID, 0, value, user.Address.String())
	return reaction, user, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgRemoveReaction tests and runs a single MsgRemoveReaction
func SimulateMsgRemoveReaction(
	k keeper.Keeper, sk subspaceskeeper.Keeper,
	ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		reaction, signer, skip := randomRemoveReactionFields(r, ctx, accs, k, sk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgRemoveReaction"), nil, nil
		}

		// Build the message
		msg := types.NewMsgRemoveReaction(
			reaction.SubspaceID,
			reaction.PostID,
			reaction.ID,
			signer.Address.String(),
		)

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{signer.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgRemoveReaction"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgRemoveReaction", nil), nil, nil
	}
}

// randomRemoveReactionFields returns the data used to build a random MsgRemoveReaction
func randomRemoveReactionFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account,
	k keeper.Keeper, sk subspaceskeeper.Keeper,
) (reaction types.Reaction, user simtypes.Account, skip bool) {
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

	// Get the reactions
	reactions := k.GetSubspaceReactions(ctx, subspaceID)
	if len(reactions) == 0 {
		// Skip because there are no reactions to be removed
		skip = true
		return
	}
	reaction = RandomReaction(r, reactions)

	// Get a user
	users := sk.GetUsersWithRootPermissions(ctx, subspace.ID, subspacestypes.NewPermissions(types.PermissionsReact))
	acc := subspacessim.GetAccount(subspacessim.RandomAddress(r, users), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	user = *acc

	return reaction, user, false
}
