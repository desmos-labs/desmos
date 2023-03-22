package simulation

// DONTCOVER

import (
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/desmos-labs/desmos/v4/testutil/simtesting"
	feeskeeper "github.com/desmos-labs/desmos/v4/x/fees/keeper"
	"github.com/desmos-labs/desmos/v4/x/posts/keeper"
	"github.com/desmos-labs/desmos/v4/x/posts/types"
	subspaceskeeper "github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	subspacessim "github.com/desmos-labs/desmos/v4/x/subspaces/simulation"
	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"
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
			return simtypes.NoOpMsg(types.ModuleName, "create post", "skip"), nil, nil
		}

		msg := types.NewMsgCreatePost(
			data.SubspaceID,
			data.SectionID,
			data.ExternalID,
			data.Text,
			data.ConversationID,
			data.ReplySettings,
			data.Entities,
			data.Tags,
			nil,
			data.ReferencedPosts,
			author.Address.String(),
		)
		txCtx, err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, author)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, "create post", "invalid"), nil, nil
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
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
	sections := sk.GetAllSections(ctx)
	if len(sections) == 0 {
		// Skip because there are no sections
		skip = true
		return
	}
	section := subspacessim.RandomSection(r, sections)

	// Get an author
	users := sk.GetUsersWithRootPermissions(ctx, section.SubspaceID, subspacestypes.NewPermissions(types.PermissionWrite))
	acc := subspacessim.GetAccount(subspacessim.RandomAddress(r, users), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	author = *acc

	if !k.HasProfile(ctx, author.Address.String()) {
		// Skip because the author does not have a profile
		skip = true
		return
	}

	postID, err := k.GetNextPostID(ctx, section.SubspaceID)
	if err != nil {
		panic(err)
	}

	post = GenerateRandomPost(r, accs, section.SubspaceID, section.ID, postID, k.GetParams(ctx))
	err = k.ValidatePost(ctx, post)
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
			return simtypes.NoOpMsg(types.ModuleName, "edit post", "skip"), nil, nil
		}

		msg := types.NewMsgEditPost(subspaceID, postID, data.Text, data.Entities, data.Tags, editor.Address.String())
		txCtx, err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, editor)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, "edit post", "invalid"), nil, nil
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// randomPostEditFields returns the data needed to edit a post
func randomPostEditFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, sk subspaceskeeper.Keeper,
) (subspaceID uint64, postID uint64, update types.PostUpdate, editor simtypes.Account, skip bool) {
	if len(accs) == 0 {
		// Skip because there are no accounts
		skip = true
		return
	}

	// Get a random post
	posts := k.GetPosts(ctx)
	if len(posts) == 0 {
		// Skip because there are no posts
		skip = true
		return
	}
	post := RandomPost(r, posts)
	subspaceID = post.SubspaceID
	sectionID := post.SectionID
	postID = post.ID

	// Get the post author
	authorAcc := subspacessim.GetAccount(post.Author, accs)
	if authorAcc == nil {
		// Skip because the author is not an account we have access to
		skip = true
		return
	}
	editor = *authorAcc

	// Check the permissions
	if !sk.HasPermission(ctx, subspaceID, sectionID, post.Author, types.PermissionEditOwnContent) {
		// Skip because the user has not the permissions
		skip = true
		return
	}

	// Generate a random update
	update = types.NewPostUpdate(
		GenerateRandomText(r, k.GetParams(ctx).MaxTextLength),
		nil,
		GenerateRandomTags(r, 4),
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
			return simtypes.NoOpMsg(types.ModuleName, "delete post", "skip"), nil, nil
		}

		msg := types.NewMsgDeletePost(subspaceID, postID, editor.Address.String())
		txCtx, err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, editor)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, "delete post", "invalid"), nil, nil
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// randomPostEditFields returns the data needed to delete a post
func randomPostDeleteFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, sk subspaceskeeper.Keeper,
) (subspaceID uint64, postID uint64, user simtypes.Account, skip bool) {
	if len(accs) == 0 {
		// Skip because there are no accounts
		skip = true
		return
	}

	// Get a random post
	posts := k.GetPosts(ctx)
	if len(posts) == 0 {
		// Skip because there are no posts
		skip = true
		return
	}
	post := RandomPost(r, posts)
	subspaceID = post.SubspaceID
	sectionID := post.SectionID
	postID = post.ID

	// Get the user
	authorAddr := post.Author
	if r.Intn(101) < 50 {
		// 50% of a moderator deleting the post
		moderators := sk.GetUsersWithRootPermissions(ctx, subspaceID, subspacestypes.NewPermissions(types.PermissionModerateContent))
		authorAddr = subspacessim.RandomAddress(r, moderators)
	} else if !sk.HasPermission(ctx, subspaceID, sectionID, authorAddr, types.PermissionEditOwnContent) {
		// Skip because the user has not the permissions
		skip = true
		return
	}

	userAcc := subspacessim.GetAccount(authorAddr, accs)
	if userAcc == nil {
		// Skip because the author is not an account we have access to
		skip = true
		return
	}
	user = *userAcc

	return subspaceID, postID, user, false
}
