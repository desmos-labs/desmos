package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		k.getSubspaceDataEntries(ctx),
		k.getPostData(ctx),
		k.getAllAttachments(ctx),
		k.getAllUserAnswers(ctx),
		k.GetParams(ctx),
	)
}

// getSubspaceDataEntries returns the subspaces data entries stored in the given context
func (k Keeper) getSubspaceDataEntries(ctx sdk.Context) []types.SubspaceDataEntry {
	var entries []types.SubspaceDataEntry
	k.sk.IterateSubspaces(ctx, func(index int64, subspace subspacestypes.Subspace) (stop bool) {
		nextPostID, err := k.GetNextPostID(ctx, subspace.ID)
		if err != nil {
			panic(err)
		}

		entries = append(entries, types.NewSubspaceDataEntry(subspace.ID, nextPostID))
		return false
	})
	return entries
}

// getPostData returns the posts data stored in the given context
func (k Keeper) getPostData(ctx sdk.Context) []types.GenesisPost {
	var posts []types.GenesisPost
	k.IteratePosts(ctx, func(index int64, post types.Post) (stop bool) {
		attachmentID, err := k.GetNextAttachmentID(ctx, post.SubspaceID, post.ID)
		if err != nil {
			panic(err)
		}

		posts = append(posts, types.NewGenesisPost(attachmentID, post))
		return false
	})
	return posts
}

// getAllAttachments returns all the attachments stored inside the given context
func (k Keeper) getAllAttachments(ctx sdk.Context) []types.Attachment {
	var attachments []types.Attachment
	k.IterateAttachments(ctx, func(index int64, attachment types.Attachment) (stop bool) {
		attachments = append(attachments, attachment)
		return false
	})
	return attachments
}

// getAllUserAnswers returns all the user answers stored inside the given context
func (k Keeper) getAllUserAnswers(ctx sdk.Context) []types.UserAnswer {
	var answers []types.UserAnswer
	k.IterateUserAnswers(ctx, func(index int64, answer types.UserAnswer) (stop bool) {
		answers = append(answers, answer)
		return false
	})
	return answers
}

// --------------------------------------------------------------------------------------------------------------------

// InitGenesis initializes the chain state based on the given GenesisState
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	// Initialize the initial post id for each subspace
	for _, entry := range data.SubspacesData {
		k.SetNextPostID(ctx, entry.SubspaceID, entry.InitialPostID)
	}

	// Initialize all the posts
	for _, post := range data.GenesisPosts {
		k.SetNextAttachmentID(ctx, post.SubspaceID, post.ID, post.InitialAttachmentID)
		k.SavePost(ctx, post.Post)
	}

	// Initialize the attachments
	for _, attachment := range data.Attachments {
		k.SaveAttachment(ctx, attachment)
		if poll, ok := attachment.Content.GetCachedValue().(*types.Poll); ok {
			if poll.EndDate.After(ctx.BlockTime()) {
				k.InsertActivePollQueue(ctx, attachment)
			}
		}
	}

	// Initialize the user answers
	for _, answer := range data.UserAnswers {
		k.SaveUserAnswer(ctx, answer)
	}

	// Initialize the params
	k.SetParams(ctx, data.Params)
}
