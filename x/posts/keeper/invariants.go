package keeper

import (
	"fmt"

	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
)

// RegisterInvariants registers all posts invariants
func RegisterInvariants(ir sdk.InvariantRegistry, keeper Keeper) {
	ir.RegisterRoute(types.ModuleName, "valid-subspaces",
		ValidSubspacesInvariant(keeper))
	ir.RegisterRoute(types.ModuleName, "valid-posts",
		ValidPostsInvariant(keeper))
	ir.RegisterRoute(types.ModuleName, "valid-attachments",
		ValidAttachmentsInvariant(keeper))
	ir.RegisterRoute(types.ModuleName, "valid-user-answers",
		ValidUserAnswersInvariant(keeper))
	ir.RegisterRoute(types.ModuleName, "valid-active-polls",
		ValidActivePollsInvariant(keeper))
}

// --------------------------------------------------------------------------------------------------------------------

// ValidSubspacesInvariant checks that all the subspaces have a valid post id to them
func ValidSubspacesInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (message string, broken bool) {
		var invalidSubspaces []subspacestypes.Subspace
		k.sk.IterateSubspaces(ctx, func(subspace subspacestypes.Subspace) (stop bool) {

			// Make sure the next post id exists for the subspace
			if !k.HasNextPostID(ctx, subspace.ID) {
				invalidSubspaces = append(invalidSubspaces, subspace)
			}

			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid subspaces",
			fmt.Sprintf("the following subspaces are invalid:\n %s", formatOutputSubspaces(invalidSubspaces)),
		), invalidSubspaces != nil
	}
}

// formatOutputPosts concatenates the given subspaces information into a string
func formatOutputSubspaces(subspaces []subspacestypes.Subspace) (output string) {
	for _, subspace := range subspaces {
		output += fmt.Sprintf("%d\n", subspace.ID)
	}
	return output
}

// --------------------------------------------------------------------------------------------------------------------

// ValidPostsInvariant checks that all the posts are valid
func ValidPostsInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (message string, broken bool) {
		var invalidPosts []types.Post
		k.IteratePosts(ctx, func(post types.Post) (stop bool) {
			invalid := false

			// The only check we need to perform here is if the subspace still exists.
			// All referenced posts might have been deleted, and params might have changed, so we can't use k.ValidatePost
			if !k.HasSubspace(ctx, post.SubspaceID) {
				invalid = true
			}

			// Make sure the section exists
			if !k.HasSection(ctx, post.SubspaceID, post.SectionID) {
				invalid = true
			}

			nextPostID, err := k.GetNextPostID(ctx, post.SubspaceID)
			if err != nil {
				invalid = true
			}

			// Make sure the post id is always less than the next one
			if post.ID >= nextPostID {
				invalid = true
			}

			// Make sure the attachment id exists
			if !k.HasNextAttachmentID(ctx, post.SubspaceID, post.ID) {
				invalid = true
			}

			// Validate the post
			err = post.Validate()
			if err != nil {
				invalid = true
			}

			if invalid {
				invalidPosts = append(invalidPosts, post)
			}

			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid posts",
			fmt.Sprintf("the following posts are invalid:\n%s", formatOutputPosts(invalidPosts)),
		), invalidPosts != nil
	}
}

// formatOutputPosts concatenates the given posts information into a string
func formatOutputPosts(posts []types.Post) (output string) {
	for _, post := range posts {
		output += fmt.Sprintf("SubspaceID: %d, PostID: %d\n",
			post.SubspaceID, post.ID)
	}
	return output
}

// --------------------------------------------------------------------------------------------------------------------

// ValidAttachmentsInvariant checks that all the attachments are valid
func ValidAttachmentsInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (message string, broken bool) {
		var invalidAttachments []types.Attachment
		k.IterateAttachments(ctx, func(attachment types.Attachment) (stop bool) {
			invalid := false

			// Check subspace
			if !k.HasSubspace(ctx, attachment.SubspaceID) {
				invalid = true
			}

			// Check associated post
			if !k.HasPost(ctx, attachment.SubspaceID, attachment.PostID) {
				invalid = true
			}

			nextAttachmentID, err := k.GetNextAttachmentID(ctx, attachment.SubspaceID, attachment.PostID)
			if err != nil {
				invalid = true
			}

			// Make sure the attachment id is always less than the next one
			if attachment.ID >= nextAttachmentID {
				invalid = true
			}

			// Validate attachment
			err = attachment.Validate()
			if err != nil {
				invalid = true
			}

			if invalid {
				invalidAttachments = append(invalidAttachments, attachment)
			}

			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid attachments",
			fmt.Sprintf("the following attachments are invalid:\n%s", formatOutputAttachments(invalidAttachments)),
		), invalidAttachments != nil
	}
}

// formatOutputAttachments concatenates the given attachment information into a string
func formatOutputAttachments(attachments []types.Attachment) (output string) {
	for _, attachment := range attachments {
		output += fmt.Sprintf("SubspaceID: %d, PostID: %d, AttachmentID: %d\n",
			attachment.SubspaceID, attachment.PostID, attachment.ID)
	}
	return output
}

// --------------------------------------------------------------------------------------------------------------------

// ValidUserAnswersInvariant checks that all the user answers are valid
func ValidUserAnswersInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (message string, broken bool) {
		var invalidUserAnswers []types.UserAnswer
		k.IterateUserAnswers(ctx, func(answer types.UserAnswer) (stop bool) {
			invalid := false

			// Check subspace
			if !k.HasSubspace(ctx, answer.SubspaceID) {
				invalid = true
			}

			// Check post
			if !k.HasPost(ctx, answer.SubspaceID, answer.PostID) {
				invalid = true
			}

			// Check the associated poll
			if !k.HasPoll(ctx, answer.SubspaceID, answer.PostID, answer.PollID) {
				invalid = true
			}

			// Validate the answer
			err := answer.Validate()
			if err != nil {
				invalid = true
			}

			if invalid {
				invalidUserAnswers = append(invalidUserAnswers, answer)
			}

			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid user answers",
			fmt.Sprintf("the following user answers are invalid:\n%s", formatOutputUserAnswers(invalidUserAnswers)),
		), invalidUserAnswers != nil
	}
}

// formatOutputUserAnswers concatenates the given user answers information into a string
func formatOutputUserAnswers(answers []types.UserAnswer) (output string) {
	for _, answer := range answers {
		output += fmt.Sprintf("SubspaceID: %d, PostID: %d, PollID: %d, User: %s\n",
			answer.SubspaceID, answer.PostID, answer.PollID, answer.User)
	}
	return output
}

// --------------------------------------------------------------------------------------------------------------------

// ValidActivePollsInvariant checks that all the active polls are valid
func ValidActivePollsInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (message string, broken bool) {
		var invalidActivePolls []types.Attachment
		k.IterateActivePolls(ctx, func(attachment types.Attachment) (stop bool) {
			poll := attachment.Content.GetCachedValue().(*types.Poll)

			// Make sure active polls do not have tally results yet
			if poll.FinalTallyResults != nil {
				invalidActivePolls = append(invalidActivePolls, attachment)
			}

			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid active polls",
			fmt.Sprintf("the following active polls are invalid:\n%s", formatOutputActivePolls(invalidActivePolls)),
		), invalidActivePolls != nil
	}
}

// formatOutputActivePolls concatenates the given polls information into a string
func formatOutputActivePolls(polls []types.Attachment) (output string) {
	for _, poll := range polls {
		output += fmt.Sprintf("SubspaceID: %d, PostID: %d, PollID: %d",
			poll.SubspaceID, poll.PostID, poll.ID)
	}
	return output
}
