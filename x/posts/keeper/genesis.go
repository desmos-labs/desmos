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
		k.getAllPosts(ctx),
		k.getAllAttachments(ctx),
		k.getAllActivePollsData(ctx),
		k.getAllUserAnswers(ctx),
		k.GetParams(ctx),
	)
}

// getSubspaceDataEntries returns the subspaces data entries stored in the given context
func (k Keeper) getSubspaceDataEntries(ctx sdk.Context) []types.SubspaceDataEntry {
	var entries []types.SubspaceDataEntry
	k.sk.IterateSubspaces(ctx, func(subspace subspacestypes.Subspace) (stop bool) {
		nextPostID, err := k.GetNextPostID(ctx, subspace.ID)
		if err != nil {
			panic(err)
		}

		entries = append(entries, types.NewSubspaceDataEntry(subspace.ID, nextPostID))
		return false
	})
	return entries
}

// getAllPosts returns the posts data stored in the given context
func (k Keeper) getAllPosts(ctx sdk.Context) []types.GenesisPost {
	var posts []types.GenesisPost
	k.IteratePosts(ctx, func(post types.Post) (stop bool) {
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
	k.IterateAttachments(ctx, func(attachment types.Attachment) (stop bool) {
		attachments = append(attachments, attachment)
		return false
	})
	return attachments
}

// getAllActivePollsData returns the active polls data
func (k Keeper) getAllActivePollsData(ctx sdk.Context) []types.ActivePollData {
	var data []types.ActivePollData
	k.IterateActivePolls(ctx, func(poll types.Attachment) (stop bool) {
		data = append(data, types.NewActivePollData(
			poll.SubspaceID,
			poll.PostID,
			poll.ID,
			poll.Content.GetCachedValue().(*types.Poll).EndDate,
		))
		return false
	})
	return data
}

// getAllUserAnswers returns all the user answers stored inside the given context
func (k Keeper) getAllUserAnswers(ctx sdk.Context) []types.UserAnswer {
	var answers []types.UserAnswer
	k.IterateUserAnswers(ctx, func(answer types.UserAnswer) (stop bool) {
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
	}

	// Initialize the active polls
	for _, pollData := range data.ActivePolls {
		k.setPollAsActive(ctx, pollData.SubspaceID, pollData.PostID, pollData.PollID, pollData.EndDate)
	}

	// Initialize the user answers
	for _, answer := range data.UserAnswers {
		k.SaveUserAnswer(ctx, answer)
	}

	// Initialize the params
	k.SetParams(ctx, data.Params)
}
