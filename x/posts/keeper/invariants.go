package keeper

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/posts/types"
)

// RegisterInvariants registers all posts invariants
func RegisterInvariants(ir sdk.InvariantRegistry, keeper Keeper) {
	ir.RegisterRoute(types.ModuleName, "valid-post",
		ValidPostsInvariant(keeper))
	ir.RegisterRoute(types.ModuleName, "comments-date",
		ValidCommentsDateInvariant(keeper))
	ir.RegisterRoute(types.ModuleName, "post-reactions",
		ValidPostForReactionsInvariant(keeper))
	ir.RegisterRoute(types.ModuleName, "post-poll-answers",
		ValidPollForPollAnswersInvariant(keeper))
}

// AllInvariants runs all invariants of the module
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		if res, stop := ValidPostsInvariant(k)(ctx); stop {
			return res, stop
		}

		if res, stop := ValidCommentsDateInvariant(k)(ctx); stop {
			return res, stop
		}

		if res, stop := ValidPollForPollAnswersInvariant(k)(ctx); stop {
			return res, stop
		}

		if res, stop := ValidPostForReactionsInvariant(k)(ctx); stop {
			return res, stop
		}

		return "Every invariant condition is fulfilled correctly", true
	}
}

//____________________________________________________________________________

// formatOutputIDs concatenate the ids given into a unique string
func formatOutputIDs(ids []string) (outputIDs string) {
	return strings.Join(ids, "\n")
}

// ValidPostsInvariant checks that the all posts are valid
func ValidPostsInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidPostIDs []string
		k.IteratePosts(ctx, func(_ int64, post types.Post) (stop bool) {
			if k.ValidatePost(ctx, post) != nil {
				invalidPostIDs = append(invalidPostIDs, post.PostID)
			}
			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid posts IDs",
			fmt.Sprintf("The following posts are invalid:\n %s", formatOutputIDs(invalidPostIDs)),
		), invalidPostIDs != nil
	}
}

//____________________________________________________________________________

// ValidCommentsDateInvariant checks that comments creation date is always greater than parent creation date
func ValidCommentsDateInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidCommentsIDs []string
		k.IteratePosts(ctx, func(_ int64, post types.Post) (stop bool) {
			if types.IsValidPostID(post.ParentID) {
				parentPost, _ := k.GetPost(ctx, post.ParentID)
				if post.Created.Before(parentPost.Created) {
					invalidCommentsIDs = append(invalidCommentsIDs, post.PostID)
				}
			}
			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "comments dates before parent post date",
			fmt.Sprintf("The following post IDs referred to posts which are invalid comments "+
				"having creation date before parent post creation date:\n %s",
				formatOutputIDs(invalidCommentsIDs)),
		), invalidCommentsIDs != nil
	}
}

//____________________________________________________________________________

// formatOutputReactions concatenate the reactions given into a unique string
func formatOutputReactions(reactions []types.PostReaction) (outputReactions string) {
	for _, reaction := range reactions {
		outputReactions += reaction.String() + "\n"
	}
	return outputReactions
}

// ValidPostForReactionsInvariant checks that the post related to the reactions is valid and exists
func ValidPostForReactionsInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidReactions []types.PostReaction
		reactions := k.GetPostReactionsEntries(ctx)
		for _, entry := range reactions {
			if !k.DoesPostExist(ctx, entry.PostId) {
				invalidReactions = append(invalidReactions, entry.Reactions...)
			}
		}

		return sdk.FormatInvariant(types.ModuleName, "posts reactions refers to non existing posts",
			fmt.Sprintf("The following reactions refer to posts that do not exist:\n %s",
				formatOutputReactions(invalidReactions)),
		), invalidReactions != nil
	}
}

//____________________________________________________________________________

// formatOutputPollAnswers concatenate the poll answers given into a unique string
func formatOutputPollAnswers(pollAnswers []types.UserAnswer) (outputAnswers string) {
	for _, answer := range pollAnswers {
		outputAnswers += answer.String() + "\n"
	}
	return outputAnswers
}

// ValidPollForPollAnswersInvariant check that the poll answers are referred to a valid post's poll
func ValidPollForPollAnswersInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidPollAnswers []types.UserAnswer
		answers := k.GetUserAnswersEntries(ctx)
		for _, entry := range answers {
			if post, found := k.GetPost(ctx, entry.PostId); !found || (found && post.PollData == nil) {
				invalidPollAnswers = append(invalidPollAnswers, entry.UserAnswers...)
			}
		}

		return sdk.FormatInvariant(types.ModuleName, "poll answers refers to posts without poll",
			fmt.Sprintf("The following answers refer to a post that either does not exist or has no poll associated to it:\n %s",
				formatOutputPollAnswers(invalidPollAnswers)),
		), invalidPollAnswers != nil
	}
}
