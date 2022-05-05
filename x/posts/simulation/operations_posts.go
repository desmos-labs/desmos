package simulation

// DONTCOVER

import (
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/desmos-labs/desmos/v3/testutil/simtesting"
	feeskeeper "github.com/desmos-labs/desmos/v3/x/fees/keeper"
	"github.com/desmos-labs/desmos/v3/x/posts/keeper"
	"github.com/desmos-labs/desmos/v3/x/posts/types"
	subspaceskeeper "github.com/desmos-labs/desmos/v3/x/subspaces/keeper"
	subspacessim "github.com/desmos-labs/desmos/v3/x/subspaces/simulation"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// SimulateMsgCreatePost tests and runs a single msg create post
func SimulateMsgCreatePost(
	k keeper.Keeper, sk subspaceskeeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		data, author, skip := randomPostCreateFields(r, ctx, accs, k, sk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "create post"), nil, nil
		}

		msg := types.NewMsgCreatePost(
			data.SubspaceID,
			data.ExternalID,
			data.Text,
			data.ConversationID,
			data.ReplySettings,
			data.Entities,
			nil,
			data.ReferencedPosts,
			author.Address.String(),
		)
		err = simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{author.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "create post"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "create post", nil), nil, nil
	}
}

// randomPostCreateFields returns the data needed to create a post
func randomPostCreateFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, sk subspaceskeeper.Keeper,
) (post types.Post, author simtypes.Account, skip bool) {
	if len(accs) == 0 {
		// Skip because there are no accounts
		skip = true
		return
	}

	// Get a subspace id
	subspaces := sk.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := subspacessim.RandomSubspace(r, subspaces)

	// Get an author
	editors, _ := sk.GetUsersWithPermission(ctx, subspace.ID, subspacestypes.PermissionWrite)
	acc := subspacessim.GetAccount(subspacessim.RandomAddress(r, editors), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	author = *acc

	post = GenerateRandomPost(r, accs, subspace.ID, 0, k.GetParams(ctx))
	err := k.ValidatePost(ctx, post)
	if err != nil {
		// Skip the operation because the post is not valid (there are too many reasons why it might be)
		skip = true
		return
	}

	return post, author, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgEditPost tests and runs a single msg edit post
func SimulateMsgEditPost(
	k keeper.Keeper, sk subspaceskeeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		subspaceID, postID, data, editor, skip := randomPostEditFields(r, ctx, accs, k, sk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "edit post"), nil, nil
		}

		msg := types.NewMsgEditPost(subspaceID, postID, data.Text, data.Entities, editor.Address.String())
		err = simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{editor.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "edit post"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "edit post", nil), nil, nil
	}
}

// randomPostEditFields returns the data needed to edit a post
func randomPostEditFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, sk subspaceskeeper.Keeper,
) (subspaceID uint64, postID uint64, update *types.PostUpdate, editor simtypes.Account, skip bool) {
	if len(accs) == 0 {
		// Skip because there are no accounts
		skip = true
		return
	}

	// Get a subspace id
	subspaces := sk.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := subspacessim.RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get an editor
	editors, _ := sk.GetUsersWithPermission(ctx, subspace.ID, subspacestypes.PermissionEditOwnContent)
	acc := subspacessim.GetAccount(subspacessim.RandomAddress(r, editors), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	editor = *acc

	// Get a post
	k.IterateSubspacePosts(ctx, subspaceID, func(index int64, post types.Post) (stop bool) {
		if post.Author == editor.Address.String() {
			postID = post.ID
			return true
		}
		return false
	})

	if postID == 0 {
		// Skip because we didn't find any post from the editor inside the given subspace
		skip = true
		return
	}

	// Generate a random update
	update = types.NewPostUpdate(
		GenerateRandomText(r, k.GetParams(ctx).MaxTextLength),
		nil,
		time.Now(),
	)
	return subspaceID, postID, update, editor, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgDeletePost tests and runs a single msg delete post
func SimulateMsgDeletePost(
	k keeper.Keeper, sk subspaceskeeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		subspaceID, postID, editor, skip := randomPostDeleteFields(r, ctx, accs, k, sk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "delete post"), nil, nil
		}

		msg := types.NewMsgDeletePost(subspaceID, postID, editor.Address.String())
		err = simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{editor.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "delete post"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "delete post", nil), nil, nil
	}
}

// randomPostEditFields returns the data needed to delete a post
func randomPostDeleteFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, sk subspaceskeeper.Keeper,
) (subspaceID uint64, postID uint64, editor simtypes.Account, skip bool) {
	if len(accs) == 0 {
		// Skip because there are no accounts
		skip = true
		return
	}

	// Get a subspace id
	subspaces := sk.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := subspacessim.RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get an editor
	selfDelete := r.Intn(101) < 50 // 50% chance of moderator deletion
	requiredPermission := subspacestypes.PermissionEditOwnContent
	if !selfDelete {
		requiredPermission = subspacestypes.PermissionModerateContent
	}

	users, _ := sk.GetUsersWithPermission(ctx, subspace.ID, requiredPermission)
	acc := subspacessim.GetAccount(subspacessim.RandomAddress(r, users), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	editor = *acc

	// Get a post
	k.IterateSubspacePosts(ctx, subspaceID, func(index int64, post types.Post) (stop bool) {
		// If is a moderator deleting the post, get a 30% of getting this post
		if !selfDelete && r.Intn(101) < 30 {
			postID = post.ID
			return true
		}

		// If self delete, get the first post by the author
		if selfDelete && post.Author == editor.Address.String() {
			postID = post.ID
			return true
		}
		return false
	})

	if postID == 0 {
		// Skip because we didn't find any post from the editor inside the given subspace
		skip = true
		return
	}

	return subspaceID, postID, editor, false
}
