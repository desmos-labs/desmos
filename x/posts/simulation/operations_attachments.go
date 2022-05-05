package simulation

// DONTCOVER

import (
	"math/rand"

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

// SimulateMsgAddPostAttachment tests and runs a single msg add post attachment
func SimulateMsgAddPostAttachment(
	k keeper.Keeper, sk subspaceskeeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		subspaceID, postID, content, editor, skip := randomAddPostAttachmentFields(r, ctx, accs, k, sk, ak)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "add post attachment"), nil, nil
		}

		msg := types.NewMsgAddPostAttachment(subspaceID, postID, content, editor.Address.String())
		err = simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{editor.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "add post attachment"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "add post attachment", nil), nil, nil
	}
}

// randomAddPostAttachmentFields returns the data needed to add an attachment to an existing post
func randomAddPostAttachmentFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, sk subspaceskeeper.Keeper, ak authkeeper.AccountKeeper,
) (subspaceID uint64, postID uint64, content types.AttachmentContent, editor simtypes.Account, skip bool) {
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

	// Generate a random attachment content
	content = GenerateRandomAttachmentContent(r)

	return subspaceID, postID, content, editor, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgRemovePostAttachment tests and runs a single msg remove post attachment
func SimulateMsgRemovePostAttachment(
	k keeper.Keeper, sk subspaceskeeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		subspaceID, postID, attachmentID, editor, skip := randomRemovePostAttachmentFields(r, ctx, accs, k, sk, ak)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "remove post attachment"), nil, nil
		}

		msg := types.NewMsgRemovePostAttachment(subspaceID, postID, attachmentID, editor.Address.String())
		err = simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{editor.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "remove post attachment"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "remove post attachment", nil), nil, nil
	}
}

// randomRemovePostAttachmentFields returns the data needed to remove an attachment from an existing post
func randomRemovePostAttachmentFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, sk subspaceskeeper.Keeper, ak authkeeper.AccountKeeper,
) (subspaceID uint64, postID uint64, attachmentID uint32, editor simtypes.Account, skip bool) {
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
