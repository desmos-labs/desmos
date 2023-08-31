package simulation

// DONTCOVER

import (
	"math/rand"

	subspacestypes "github.com/desmos-labs/desmos/v6/x/subspaces/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/desmos-labs/desmos/v6/testutil/simtesting"

	subspacessim "github.com/desmos-labs/desmos/v6/x/subspaces/simulation"

	"github.com/desmos-labs/desmos/v6/x/posts/keeper"
	"github.com/desmos-labs/desmos/v6/x/posts/types"
)

// SimulateMsgAddPostAttachment tests and runs a single msg add post attachment
func SimulateMsgAddPostAttachment(
	k *keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		subspaceID, postID, content, editor, skip := randomAddPostAttachmentFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "add post attachment", "skip"), nil, nil
		}

		msg := types.NewMsgAddPostAttachment(subspaceID, postID, content, editor.Address.String())
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, editor)
	}
}

// randomAddPostAttachmentFields returns the data needed to add an attachment to an existing post
func randomAddPostAttachmentFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k *keeper.Keeper,
) (subspaceID uint64, postID uint64, content types.AttachmentContent, editor simtypes.Account, skip bool) {
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
	postID = post.ID

	// Get an editor
	acc := subspacessim.GetAccount(post.Owner, accs)
	if acc == nil {
		// Skip because the author is not an account we have access to
		skip = true
		return
	}
	if !k.HasPermission(ctx, subspaceID, post.SectionID, acc.Address.String(), types.PermissionEditOwnContent) {
		// Skip because the author does not have permission
		skip = true
		return
	}

	editor = *acc

	// Generate a random attachment content
	content = GenerateRandomAttachmentContent(r, ctx.BlockTime())

	return subspaceID, postID, content, editor, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgRemovePostAttachment tests and runs a single msg remove post attachment
func SimulateMsgRemovePostAttachment(
	k *keeper.Keeper, sk types.SubspacesKeeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		subspaceID, postID, attachmentID, editor, skip := randomRemovePostAttachmentFields(r, ctx, accs, k, sk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "remove post attachment", "skip"), nil, nil
		}

		msg := types.NewMsgRemovePostAttachment(subspaceID, postID, attachmentID, editor.Address.String())

		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, editor)
	}
}

// randomRemovePostAttachmentFields returns the data needed to remove an attachment from an existing post
func randomRemovePostAttachmentFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k *keeper.Keeper, sk types.SubspacesKeeper,
) (subspaceID uint64, postID uint64, attachmentID uint32, editor simtypes.Account, skip bool) {
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

	// Get an editor
	editorAddr := post.Owner
	if r.Intn(101) < 50 {
		// 50% of a moderator removing an attachment
		moderators := sk.GetUsersWithRootPermissions(ctx, subspaceID, subspacestypes.NewPermissions(types.PermissionModerateContent))
		editorAddr = subspacessim.RandomAddress(r, moderators)
	} else if !sk.HasPermission(ctx, subspaceID, sectionID, editorAddr, types.PermissionEditOwnContent) {
		// Skip because the user has not the permission to edit their own content
		skip = true
		return
	}

	acc := subspacessim.GetAccount(editorAddr, accs)
	if acc == nil {
		// Skip because the author is not an account we have access to
		skip = true
		return
	}
	editor = *acc

	// Get a random attachment
	attachments := k.GetPostAttachments(ctx, subspaceID, postID)
	if len(attachments) == 0 {
		// Skip because the post has no attachment
		skip = true
		return
	}

	attachment := RandomAttachment(r, attachments)
	attachmentID = attachment.ID

	return subspaceID, postID, attachmentID, editor, false
}
